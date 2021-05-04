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

                var connectionFactory = new ConnectionFactory();
                var connection = connectionFactory.CreateConnection(_natsOptions.Url);

                connection.SubscribeAsync("todo.processed", async (_, args) =>
                {
                    _logger.LogInformation("message received ({Subject})", args.Message.Subject);

                    try
                    {
                        var todo = JsonSerializer.Deserialize<Todo>(args.Message.Data);
                        await _apiGatewayClient.SendNotification(todo);
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