using System.ComponentModel.DataAnnotations;
using System.Text.Json;

namespace ApiGateway.Models
{
    public class Notification
    {
        [Required]
        public string Type { get; set; }
        
        public JsonElement Data { get; set; }
    }
}