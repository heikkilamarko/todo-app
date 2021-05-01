using System.ComponentModel.DataAnnotations;
using System.Text.Json.Serialization;

namespace ApiGateway.Models
{
    public class GetTodosQuery
    {
        [JsonPropertyName("offset")]
        [Range(0, int.MaxValue)]
        public int? Offset { get; set; } = 0;

        [JsonPropertyName("limit")]
        [Range(1, 100)]
        public int? Limit { get; set; } = 100;
    }
}