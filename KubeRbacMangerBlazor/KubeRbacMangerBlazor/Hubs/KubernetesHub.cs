using Microsoft.AspNetCore.SignalR;
using System.Threading.Tasks;

namespace KubeRbacMangerBlazor.Hubs
{
    public class KubernetesHub : Hub
    {
        public static string ClientsString = "Clients";
        private readonly KubernetesService kubernetesService;

        public KubernetesHub(KubernetesService kubernetesService)
        {
            this.kubernetesService = kubernetesService;
        }
        public override async Task OnConnectedAsync()
        {
            await Groups.AddToGroupAsync(Context.ConnectionId, ClientsString);
        }

        public async Task GetRoles(string @namespace)
        {
            await kubernetesService.GetRoles(@namespace);
        }
    }
}
