using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.FileProviders;
using NotificationService.Hubs;
using NotificationService.Models;
using NotificationService.Services;

namespace NotificationService
{
    public class Startup
    {
        public Startup(IConfiguration configuration)
        {
            Configuration = configuration;
        }

        public IConfiguration Configuration { get; }

        public void ConfigureServices(IServiceCollection services)
        {
            services.AddMemoryCache();
            services.AddSignalR();

            services.Configure<NatsOptions>(o => Configuration.Bind("Nats", o));

            services.AddSingleton<SchemaValidator>();

            services.AddSingleton<IFileProvider>(
                new ManifestEmbeddedFileProvider(typeof(Program).Assembly));

            services.AddHostedService<Worker>();
        }

        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            app.UseRouting();
            app.UseEndpoints(endpoints => { endpoints.MapHub<NotificationsHub>("/notifications"); });
        }
    }
}