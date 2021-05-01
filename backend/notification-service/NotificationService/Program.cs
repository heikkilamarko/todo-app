using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using NotificationService.Models;
using NotificationService.Services;

namespace NotificationService
{
    public class Program
    {
        public static void Main(string[] args)
        {
            CreateHostBuilder(args).Build().Run();
        }

        public static IHostBuilder CreateHostBuilder(string[] args) =>
            Host.CreateDefaultBuilder(args)
                .ConfigureServices((hostContext, services) =>
                {
                    var c = hostContext.Configuration;

                    services.Configure<NatsOptions>(o => c.Bind("Nats", o));

                    services.AddApiGatewayClient(o => c.Bind("ApiGateway", o));

                    services.AddHostedService<Worker>();
                });
    }
}