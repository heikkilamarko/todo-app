using System.ComponentModel.DataAnnotations;
using System.Text.Json.Serialization;

namespace ApiGateway.Models
{
    public class NewTodo
    {
        [JsonPropertyName("name")]
        [Required]
        [MaxLength(100)]
        public string Name { get; set; }

        [JsonPropertyName("description")]
        [MaxLength(1000)]
        public string Description { get; set; }
    }
}
