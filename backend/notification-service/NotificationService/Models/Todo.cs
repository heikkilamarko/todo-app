using System;
using System.ComponentModel.DataAnnotations;
using System.Text.Json.Serialization;

namespace NotificationService.Models
{
    public class Todo
    {
        [JsonPropertyName("id")] 
        [Required] 
        public int? Id { get; set; }

        [JsonPropertyName("name")]
        [Required]
        [MaxLength(100)]
        public string Name { get; set; }

        [JsonPropertyName("description")]
        [MaxLength(1000)]
        public string Description { get; set; }

        [JsonPropertyName("created_at")]
        [Required]
        public DateTimeOffset? CreatedAt { get; set; }

        [JsonPropertyName("updated_at")]
        [Required]
        public DateTimeOffset? UpdatedAt { get; set; }
    }
}