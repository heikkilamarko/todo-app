using System;
using Microsoft.Extensions.DependencyInjection;
using Polly;

namespace NotificationService.Services
{
    public static class ApiGatewayClientServiceCollectionExtensions
    {
        public static IServiceCollection AddApiGatewayClient(
            this IServiceCollection services,
            Action<ApiGatewayClientOptions> setupAction)
        {
            _ = services ?? throw new ArgumentNullException(nameof(services));
            _ = setupAction ?? throw new ArgumentNullException(nameof(setupAction));

            services.Configure(setupAction);

            services.AddHttpClient<ApiGatewayClient>()
                .AddTransientHttpErrorPolicy(p =>
                    p.WaitAndRetryAsync(3, i =>
                        TimeSpan.FromSeconds(Math.Pow(2, i))));

            return services;
        }
    }
}