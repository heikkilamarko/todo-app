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
        private readonly ApiGatewayClient _apiGatewayClient;
        private readonly ILogger<Worker> _logger;

        public Worker(
            ILogger<Worker> logger,
            IOptions<NatsOptions> natsOptions,
            ApiGatewayClient apiGatewayClient)
        {
            _logger = logger;
            _natsOptions = natsOptions.Value;
            _apiGatewayClient = apiGatewayClient;
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

                connection.SubscribeAsync(Constants.MessageTodoCreatedOk, async (_, args) =>
                {
                    _logger.LogInformation("message received ({Subject})", args.Message.Subject);

                    try
                    {
                        var data = JsonSerializer.Deserialize<TodoCreatedOk>(args.Message.Data);
                        await _apiGatewayClient.SendNotification(new Notification
                        {
                            Type = Constants.MessageTodoCreatedOk,
                            Data = data
                        });
                    }
                    catch (Exception exception)
                    {
                        _logger.LogError(exception, "message handling failed");
                    }
                });

                connection.SubscribeAsync(Constants.MessageTodoCompletedOk, async (_, args) =>
                {
                    _logger.LogInformation("message received ({Subject})", args.Message.Subject);

                    try
                    {
                        var data = JsonSerializer.Deserialize<TodoCompletedOk>(args.Message.Data);
                        await _apiGatewayClient.SendNotification(new Notification
                        {
                            Type = Constants.MessageTodoCompletedOk,
                            Data = data
                        });
                    }
                    catch (Exception exception)
                    {
                        _logger.LogError(exception, "message handling failed");
                    }
                });

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
    }
}