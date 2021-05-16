using System.Collections.Generic;
using System.Text.Json;
using System.Threading.Tasks;
using Json.Schema;
using Microsoft.Extensions.FileProviders;

namespace NotificationService.Services
{
    public class SchemaValidator
    {
        private readonly Dictionary<string, JsonSchema> _schemas = new();

        private readonly IFileProvider _fileProvider;

        public SchemaValidator(IFileProvider fileProvider)
        {
            _fileProvider = fileProvider;
        }

        public async Task<ValidationResults> ValidateAsync(string schemaName, JsonElement data)
        {
            if (!_schemas.TryGetValue(schemaName, out var schema))
            {
                var info = _fileProvider.GetFileInfo($"Schemas/{schemaName}.json");
                schema = await JsonSchema.FromStream(info.CreateReadStream());
                _schemas.Add(schemaName, schema);
            }

            return schema.Validate(data,
                new ValidationOptions
                {
                    ValidateAs = Draft.Draft7,
                    OutputFormat = OutputFormat.Detailed
                });
        }
    }
}