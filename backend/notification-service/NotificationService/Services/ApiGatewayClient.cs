using System;
using System.Net.Http;
using System.Net.Http.Json;
using System.Threading.Tasks;
using Microsoft.Extensions.Logging;
using Microsoft.Extensions.Options;
using NotificationService.Models;

namespace NotificationService.Services
{
    public class ApiGatewayClient
    {
        private readonly HttpClient _httpClient;
        private readonly ILogger<ApiGatewayClient> _logger;

        public ApiGatewayClient(
            HttpClient httpClient,
            IOptions<ApiGatewayClientOptions> options,
            ILogger<ApiGatewayClient> logger)
        {
            _httpClient = httpClient;
            _httpClient.BaseAddress = new Uri(options.Value.Url);
            _logger = logger;
        }

        public async Task SendNotification(Notification notification)
        {
            var response = await _httpClient.PostAsJsonAsync("notifications", notification);

            if (!response.IsSuccessStatusCode)
            {
                var body = await response.Content.ReadAsStringAsync();
                _logger.LogInformation("Error body: {Body}", body);
            }

            response.EnsureSuccessStatusCode();
        }
    }
}