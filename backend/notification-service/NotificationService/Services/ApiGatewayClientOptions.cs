namespace NotificationService.Services
{
    public class ApiGatewayClientOptions
    {
        public string Url { get; set; }
        public string TokenUrl { get; set; }
        public string ClientId { get; set; }
        public string ClientSecret { get; set; }
        public string Scope { get; set; }
    }
}