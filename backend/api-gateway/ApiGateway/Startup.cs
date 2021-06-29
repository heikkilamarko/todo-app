using System.Threading.Tasks;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;

namespace ApiGateway
{
    public class Startup
    {
        public Startup(IConfiguration configuration)
        {
            Configuration = configuration;
        }

        public IConfiguration Configuration { get; }

        public void ConfigureServices(IServiceCollection services)
        {
            services.AddCors(options =>
            {
                options.AddPolicy("app",
                    builder =>
                    {
                        builder
                            .WithOrigins(Configuration["Cors:Origins"].Split(","))
                            .AllowAnyHeader()
                            .AllowCredentials();
                    });
            });

            services.AddReverseProxy().LoadFromConfig(Configuration.GetSection("ReverseProxy"));

            services.AddAuthentication(JwtBearerDefaults.AuthenticationScheme).AddJwtBearer(o =>
            {
                o.Audience = Configuration["Auth:Audience"];
                o.Authority = Configuration["Auth:Authority"];
                o.MetadataAddress = Configuration["Auth:MetadataAddress"];
                o.TokenValidationParameters.ValidIssuers = Configuration["Auth:ValidIssuers"].Split(",");
                o.RequireHttpsMetadata = false;
                o.Events = new JwtBearerEvents
                {
                    // SignalR
                    OnMessageReceived = context =>
                    {
                        var accessToken = context.Request.Query["access_token"];

                        if (!string.IsNullOrWhiteSpace(accessToken))
                        {
                            context.Token = accessToken;
                        }

                        return Task.CompletedTask;
                    }
                };
            });

            services.AddAuthorization(o => o.AddPolicy("app", b => b.RequireAuthenticatedUser()));
        }

        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            app.UseRouting();
            app.UseCors();
            app.UseAuthentication();
            app.UseAuthorization();
            app.UseEndpoints(endpoints => { endpoints.MapReverseProxy(); });
        }
    }
}