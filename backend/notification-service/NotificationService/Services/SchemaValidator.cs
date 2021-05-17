using System.Collections.Generic;
using System.Text.Json;
using System.Threading.Tasks;
using Json.Schema;
using Microsoft.Extensions.Caching.Memory;
using Microsoft.Extensions.FileProviders;
using Microsoft.Extensions.Logging;
using Serilog;

namespace NotificationService.Services
{
    public class SchemaValidator
    {
        private readonly Dictionary<string, JsonSchema> _schemas = new();

        private readonly IFileProvider _fileProvider;
        private readonly IMemoryCache _memoryCache;
        private readonly ILogger<SchemaValidator> _logger;

        public SchemaValidator(IFileProvider fileProvider, IMemoryCache memoryCache, ILogger<SchemaValidator> logger)
        {
            _fileProvider = fileProvider;
            _memoryCache = memoryCache;
            _logger = logger;
        }

        public async Task<ValidationResults> ValidateAsync(string schemaName, JsonElement data)
        {
            var schema = await GetSchemaAsync(schemaName);

            if (schema == null)
            {
                _logger.LogWarning("schema not found ({SchemaName})", schemaName);
                return null;
            }

            return schema.Validate(data,
                new ValidationOptions
                {
                    ValidateAs = Draft.Draft7,
                    OutputFormat = OutputFormat.Detailed
                });
        }

        private async Task<JsonSchema> GetSchemaAsync(string schemaName)
        {
            var cacheKey = $"schema:{schemaName}";

            return await _memoryCache.GetOrCreateAsync(cacheKey, async _ =>
            {
                var info = _fileProvider.GetFileInfo($"Schemas/{schemaName}.json");
                return info.Exists
                    ? await JsonSchema.FromStream(info.CreateReadStream())
                    : null;
            });
        }
    }
}