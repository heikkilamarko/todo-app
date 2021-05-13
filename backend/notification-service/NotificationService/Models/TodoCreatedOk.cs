using System.ComponentModel.DataAnnotations;
using System.Text.Json.Serialization;

namespace NotificationService.Models
{
    public class TodoCreatedOk
    {
        [JsonPropertyName("todo")]
        [Required]
        public Todo Todo { get; set; }
    }
}