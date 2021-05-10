using System.ComponentModel.DataAnnotations;
using System.Text.Json.Serialization;

namespace NotificationService.Models
{
    public class TodoCompletedOk
    {
        [JsonPropertyName("id")]
        [Required]
        public int? Id { get; set; }
    }
}