using System.Collections.Generic;
using System.Text.Json.Serialization;

namespace ApiGateway.Models
{
    public class GetTodosResponse
    {
        [JsonPropertyName("meta")]
        public GetTodosResponseMeta Meta { get; set; }

        [JsonPropertyName("data")]
        public List<Todo> Data { get; set; }
    }
}