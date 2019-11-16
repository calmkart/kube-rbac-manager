using Microsoft.AspNetCore.SignalR;
using System;
using System.Threading.Tasks;

namespace KubeRbacMangerBlazor.Hubs
{
    public class KubernetesHub : Hub
    {
        public static string ClientsString = "Clients";
        public static string WatchRolesString = "WatchRoles";
        public static string WatchServiceAccountsString = "WatchServiceAccounts";
        public static string WatchRoleBindingsString = "WatchRoleBindings";
        public static string WatchClusterRoleBindingsString = "WatchClusterRoleBindingsString";
        private readonly KubernetesService kubernetesService;

        public KubernetesHub(KubernetesService kubernetesService)
        {
            this.kubernetesService = kubernetesService;
        }
        public override async Task OnConnectedAsync()
        {
            await Groups.AddToGroupAsync(Context.ConnectionId, ClientsString);
        }

        public override async Task OnDisconnectedAsync(Exception exception)
        {
            await Groups.RemoveFromGroupAsync(Context.ConnectionId, ClientsString);
            await Groups.RemoveFromGroupAsync(Context.ConnectionId, WatchRolesString);
            await Groups.RemoveFromGroupAsync(Context.ConnectionId, WatchServiceAccountsString);
            await Groups.RemoveFromGroupAsync(Context.ConnectionId, WatchRoleBindingsString);
            await Groups.RemoveFromGroupAsync(Context.ConnectionId, WatchClusterRoleBindingsString);
        }

        public async Task GetNamespacedRole(string @namespace)
        {
            await Groups.AddToGroupAsync(Context.ConnectionId, WatchRolesString);

            await kubernetesService.GetNamespacedRole(@namespace);
        }

        public async Task GetNamespacedServiceAccount(string @namespace)
        {
            await Groups.AddToGroupAsync(Context.ConnectionId, WatchServiceAccountsString);
            await kubernetesService.GetNamespacedServiceAccount(@namespace);
        }

        public async Task GetNamespacedRoleBinding(string @namespace)
        {
            await Groups.AddToGroupAsync(Context.ConnectionId, WatchRoleBindingsString);
            await kubernetesService.GetNamespacedRoleBinding(@namespace);
        }

        public async Task GetClusterRoleBinding()
        {
            await Groups.AddToGroupAsync(Context.ConnectionId, WatchClusterRoleBindingsString);
            await kubernetesService.GetClusterRoleBinding();
        }
    }
}
