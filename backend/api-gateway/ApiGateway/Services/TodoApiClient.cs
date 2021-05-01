using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Options;
using Polly;
using System;
using System.Net.Http;
using System.Net.Http.Json;
using System.Threading.Tasks;
using ApiGateway.Models;

namespace ApiGateway.Services
{
    public class TodoApiClientOptions
    {
        public string Url { get; set; }
    }

    public class TodoApiClient
    {
        private readonly TodoApiClientOptions _options;
        private readonly HttpClient _httpClient;

        public TodoApiClient(IOptions<TodoApiClientOptions> options, HttpClient httpClient)
        {
            _options = options.Value;

            httpClient.BaseAddress = new Uri(_options.Url);

            _httpClient = httpClient;
        }

        public async Task<GetTodosResponse> GetTodos(string queryString)
        {
            return await _httpClient.GetFromJsonAsync<GetTodosResponse>($"todos{queryString}");
        }

        public async Task CreateTodo(NewTodo todo)
        {
            var response = await _httpClient.PostAsJsonAsync("todos", todo);
            response.EnsureSuccessStatusCode();
        }
    }

    public static class TodoApiClientServiceCollectionExtensions
    {
        public static IServiceCollection AddTodoApiClient(this IServiceCollection services,
            Action<TodoApiClientOptions> setupAction)
        {
            _ = services ?? throw new ArgumentNullException(nameof(services));
            _ = setupAction ?? throw new ArgumentNullException(nameof(setupAction));

            services.AddOptions();
            services.Configure(setupAction);

            services.AddHttpClient<TodoApiClient>()
                .AddTransientHttpErrorPolicy(p =>
                    p.WaitAndRetryAsync(3, i =>
                        TimeSpan.FromSeconds(Math.Pow(2, i))));

            return services;
        }
    }
}