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
            _logger.LogInformation("Application running");

            var cf = new ConnectionFactory();
            var c = cf.CreateConnection(_natsOptions.Url);

            EventHandler<MsgHandlerEventArgs> h = async (sender, args) =>
            {
                _logger.LogInformation("Message received ({subject})", args.Message.Subject);

                try
                {
                    var todo = JsonSerializer.Deserialize<Todo>(args.Message.Data);
                    await _apiGatewayClient.SendNotification(todo);
                }
                catch (Exception exception)
                {
                    _logger.LogError(exception, "Todo handling failed");
                }
            };

            var s = c.SubscribeAsync("todo.processed", h);

            while (!stoppingToken.IsCancellationRequested)
            {
                await Task.Delay(5000, stoppingToken);
            }

            _logger.LogInformation("Application shutdown.");
        }
    }
}