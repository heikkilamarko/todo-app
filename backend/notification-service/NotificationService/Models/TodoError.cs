using System.ComponentModel.DataAnnotations;
using System.Text.Json.Serialization;

namespace NotificationService.Models
{
    public class TodoError
    {
        [JsonPropertyName("code")]
        [Required]
        public string Code { get; set; }

        [JsonPropertyName("message")]
        [Required]
        public string Message { get; set; }
    }
}