using System;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.FileProviders;
using Microsoft.Extensions.Hosting;
using NotificationService.Models;
using NotificationService.Services;
using Serilog;
using Serilog.Events;
using Serilog.Formatting.Compact;

namespace NotificationService
{
    public class Program
    {
        public static int Main(string[] args)
        {
            Log.Logger = new LoggerConfiguration()
                .MinimumLevel.Override("Microsoft", LogEventLevel.Warning)
                .Enrich.WithProperty("app", "notification-service")
                .Enrich.FromLogContext()
                .WriteTo.Console(new RenderedCompactJsonFormatter())
                .CreateBootstrapLogger();

            try
            {
                Log.Information("starting host...");
                CreateHostBuilder(args).Build().Run();
                return 0;
            }
            catch (Exception ex)
            {
                Log.Fatal(ex, "host terminated unexpectedly");
                return 1;
            }
            finally
            {
                Log.CloseAndFlush();
            }
        }

        public static IHostBuilder CreateHostBuilder(string[] args) =>
            Host.CreateDefaultBuilder(args)
                .UseSerilog((_, _, configuration) =>
                {
                    configuration
                        .MinimumLevel.Override("Microsoft", LogEventLevel.Warning)
                        .Enrich.WithProperty("app", "notification-service")
                        .Enrich.FromLogContext()
                        .WriteTo.Console(new RenderedCompactJsonFormatter());
                })
                .ConfigureServices((hostContext, services) =>
                {
                    var c = hostContext.Configuration;

                    services.Configure<NatsOptions>(o => c.Bind("Nats", o));

                    services.AddApiGatewayClient(o => c.Bind("ApiGateway", o));

                    services.AddSingleton<SchemaValidator>();

                    services.AddSingleton<IFileProvider>(
                        new ManifestEmbeddedFileProvider(typeof(Program).Assembly));

                    services.AddHostedService<Worker>();
                });
    }
}