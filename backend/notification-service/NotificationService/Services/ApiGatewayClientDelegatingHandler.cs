using System.Net.Http;
using System.Net.Http.Headers;
using System.Threading;
using System.Threading.Tasks;
using IdentityModel.Client;
using Microsoft.Extensions.Options;

namespace NotificationService.Services
{
    public class ApiGatewayClientDelegatingHandler : DelegatingHandler
    {
        private readonly ApiGatewayClientOptions _options;
        private readonly IHttpClientFactory _clientFactory;

        public ApiGatewayClientDelegatingHandler(
            IOptions<ApiGatewayClientOptions> options,
            IHttpClientFactory clientFactory)
        {
            _options = options.Value;
            _clientFactory = clientFactory;
        }

        protected override async Task<HttpResponseMessage> SendAsync(HttpRequestMessage request,
            CancellationToken cancellationToken)
        {
            if (request.Headers.Authorization == null)
            {
                var accessToken = await GetAccessTokenAsync();
                request.Headers.Authorization = new AuthenticationHeaderValue("Bearer", accessToken);
            }

            return await base.SendAsync(request, cancellationToken);
        }

        private async Task<string> GetAccessTokenAsync()
        {
            var response = await _clientFactory
                .CreateClient("auth")
                .RequestClientCredentialsTokenAsync(
                    new ClientCredentialsTokenRequest
                    {
                        Address = _options.TokenUrl,
                        ClientId = _options.ClientId,
                        ClientSecret = _options.ClientSecret,
                        Scope = _options.Scope
                    });

            return response.AccessToken;
        }
    }
}