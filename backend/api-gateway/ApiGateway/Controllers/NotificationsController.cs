using System.Threading;
using ApiGateway.Models;
using Microsoft.AspNetCore.Mvc;
using System.Threading.Tasks;
using ApiGateway.Hubs;
using Microsoft.AspNetCore.SignalR;

namespace ApiGateway.Controllers
{
    [ApiController]
    [Route("notifications")]
    public class NotificationsController : ControllerBase
    {
        private readonly IHubContext<NotificationsHub> _hub;

        public NotificationsController(IHubContext<NotificationsHub> hub)
        {
            _hub = hub;
        }

        [HttpPost]
        public async Task<ActionResult> PostNotification([FromBody] Todo todo, CancellationToken token)
        {
            await _hub.Clients.All.SendAsync("ReceiveNotification", todo, token);

            return Ok();
        }
    }
}