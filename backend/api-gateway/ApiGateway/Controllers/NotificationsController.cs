using System.Threading;
using System.Threading.Tasks;
using ApiGateway.Hubs;
using ApiGateway.Models;
using Microsoft.AspNetCore.Mvc;
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
        public async Task<ActionResult> PostNotification([FromBody] Notification notification, CancellationToken token)
        {
            await _hub.Clients.All.SendAsync("ReceiveNotification", notification, token);

            return Ok();
        }
    }
}