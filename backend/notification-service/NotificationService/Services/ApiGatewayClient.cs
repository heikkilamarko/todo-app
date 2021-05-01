using System;
using System.Net.Http;
using System.Net.Http.Json;
using System.Threading.Tasks;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Options;
using NotificationService.Models;
using Polly;

namespace NotificationService.Services
{
    public class ApiGatewayClient
    {
        private readonly ApiGatewayOptions _options;
        private readonly HttpClient _httpClient;

        public ApiGatewayClient(IOptions<ApiGatewayOptions> options, HttpClient httpClient)
        {
            _options = options.Value;

            httpClient.BaseAddress = new Uri(_options.Url);

            _httpClient = httpClient;
        }

        public async Task SendNotification(Todo todo)
        {
            var response = await _httpClient.PostAsJsonAsync("notifications", todo);
            response.EnsureSuccessStatusCode();
        }
    }

    public static class ApiGatewayClientServiceCollectionExtensions
    {
        public static IServiceCollection AddApiGatewayClient(this IServiceCollection services,
            Action<ApiGatewayOptions> setupAction)
        {
            _ = services ?? throw new ArgumentNullException(nameof(services));
            _ = setupAction ?? throw new ArgumentNullException(nameof(setupAction));

            services.AddOptions();
            services.Configure(setupAction);

            services.AddHttpClient<ApiGatewayClient>()
                .AddTransientHttpErrorPolicy(p =>
                    p.WaitAndRetryAsync(3, i =>
                        TimeSpan.FromSeconds(Math.Pow(2, i))));

            return services;
        }
    }
}