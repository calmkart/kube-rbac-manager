using k8s;
using k8s.Models;
using KubeRbacMangerBlazor.Hubs;
using Microsoft.AspNetCore.SignalR;
using System.Threading.Tasks;

namespace KubeRbacMangerBlazor
{
    public class KubernetesService
    {
        private readonly IHubContext<KubernetesHub> hubContext;

        public KubernetesService(IHubContext<KubernetesHub> hubContext)
        {
            this.hubContext = hubContext;
        }

        public async Task GetRoles(string @namespace)
        {
            var config = KubernetesClientConfiguration.BuildConfigFromConfigFile();

            IKubernetes client = new Kubernetes(config);

            var podlistResp = await client.ListNamespacedRoleWithHttpMessagesAsync(@namespace, watch: true);
            podlistResp.Watch<V1Role, V1RoleList>(Watcher_OnEvent);
        }

        private async void Watcher_OnEvent(WatchEventType watchEventType, V1Role v1Role)
        {
            switch (watchEventType)
            {
                case WatchEventType.Added:
                    await hubContext.Clients.Group(KubernetesHub.ClientsString).SendAsync("NotifyRoleAddEvent", v1Role.Metadata.Name);
                    break;
                case WatchEventType.Modified:
                    break;
                case WatchEventType.Deleted:
                    await hubContext.Clients.Group(KubernetesHub.ClientsString).SendAsync("NotifyRoleDeleteEvent", v1Role.Metadata.Name);
                    break;
                case WatchEventType.Error:
                    break;
                default:
                    break;
            }
        }
    }
}
