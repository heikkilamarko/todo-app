using System.Text.Json.Serialization;

namespace ApiGateway.Models
{
    public class GetTodosResponseMeta
    {
        [JsonPropertyName("offset")] 
        public int? Offset { get; set; }

        [JsonPropertyName("limit")] 
        public int? Limit { get; set; }
    }
}