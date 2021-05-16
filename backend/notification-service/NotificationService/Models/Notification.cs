using System.Text.Json;
using System.Text.Json.Serialization;

namespace NotificationService.Models
{
    public class Notification
    {
        [JsonPropertyName("type")]
        public string Type { get; set; }

        [JsonPropertyName("data")]
        public JsonElement Data { get; set; }
    }
}