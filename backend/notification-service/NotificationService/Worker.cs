using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;
using System;
using System.Text.Json;
using System.Threading;
using System.Threading.Tasks;
using Microsoft.Extensions.Options;
using NATS.Client;
using NotificationService.Models;
using NotificationService.Services;

namespace NotificationService
{
    public class Worker : BackgroundService
    {
        private readonly NatsOptions _natsOptions;
        private readonly SchemaValidator _schemaValidator;
        private readonly ApiGatewayClient _apiGatewayClient;
        private readonly ILogger<Worker> _logger;

        public Worker(
            IOptions<NatsOptions> natsOptions,
            SchemaValidator schemaValidator,
            ApiGatewayClient apiGatewayClient,
            ILogger<Worker> logger)
        {
            _natsOptions = natsOptions.Value;
            _schemaValidator = schemaValidator;
            _apiGatewayClient = apiGatewayClient;
            _logger = logger;
        }

        protected override async Task ExecuteAsync(CancellationToken stoppingToken)
        {
            try
            {
                _logger.LogInformation("application is starting up...");

                var options = ConnectionFactory.GetDefaultOptions();
                options.Url = _natsOptions.Url;
                options.Token = _natsOptions.Token;

                var connection = new ConnectionFactory()
                    .CreateConnection(options);

                connection.SubscribeAsync(
                    Constants.MessageTodo,
                    async (_, args) =>
                    {
                        _logger.LogInformation("message received ({Subject})", args.Message.Subject);
                        await HandleMessageAsync(args.Message);
                        _logger.LogInformation("message handled ({Subject})", args.Message.Subject);
                    }
                );

                _logger.LogInformation("application is running");

                while (!stoppingToken.IsCancellationRequested)
                {
                    try
                    {
                        await Task.Delay(5000, stoppingToken);
                    }
                    catch
                    {
                        // ignored
                    }
                }

                _logger.LogInformation("application is shutting down...");

                await connection.DrainAsync();

                _logger.LogInformation("application is shut down");
            }
            catch (Exception exception)
            {
                _logger.LogError(exception, "critical error");
                throw;
            }
        }

        private async Task HandleMessageAsync(Msg message)
        {
            try
            {
                var data = JsonSerializer.Deserialize<JsonElement>(message.Data);

                var results = await _schemaValidator.ValidateAsync(message.Subject, data);

                if (!results.IsValid)
                {
                    _logger.LogWarning("invalid message");
                    return;
                }

                await _apiGatewayClient.SendNotification(new Notification
                {
                    Type = message.Subject,
                    Data = data
                });
            }
            catch (Exception exception)
            {
                _logger.LogError(exception, "message handling failed");
            }
        }
    }
}