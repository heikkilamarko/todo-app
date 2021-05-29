using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Cors;
using Microsoft.AspNetCore.SignalR;

namespace ApiGateway.Hubs
{
    [Authorize("app")]
    [EnableCors("app")]
    public class NotificationsHub : Hub
    {
    }
}