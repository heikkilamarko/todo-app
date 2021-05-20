using System.Threading;
using System.Threading.Tasks;

namespace NotificationService.Extensions
{
    public static class CancellationTokenExtensions
    {
        public static async Task WaitShutdownAsync(this CancellationToken cancellationToken)
        {
            if (!cancellationToken.IsCancellationRequested)
            {
                try
                {
                    await Task.Delay(-1, cancellationToken);
                }
                catch
                {
                    // ignored
                }
            }
        }
    }
}