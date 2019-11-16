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
        Watcher<V1Role> namespacedRoleWatcher;
        Watcher<V1ServiceAccount> namespacedServiceAccountWatcher;
        Watcher<V1RoleBinding> namespacedRoleBindingWatcher;
        Watcher<V1Role> roleWatcher;
        Watcher<V1ClusterRoleBinding> clusterRoleBindingWatcher;
        public KubernetesService(IHubContext<KubernetesHub> hubContext)
        {
            this.hubContext = hubContext;
        }

        public async Task GetNamespacedRole(string @namespace)
        {
            var config = KubernetesClientConfiguration.BuildConfigFromConfigFile();

            IKubernetes client = new Kubernetes(config);

            var namespacedRoleListResp = await client.ListNamespacedRoleWithHttpMessagesAsync(@namespace, watch: true);
            namespacedRoleWatcher = namespacedRoleListResp.Watch<V1Role, V1RoleList>(async (watchEventType, v1Role) =>
            {
                switch (watchEventType)
                {
                    case WatchEventType.Added:
                        await hubContext.Clients.Group(KubernetesHub.WatchRolesString).SendAsync("NotifyNamespacedRoleAddEvent", v1Role.Metadata.Name);
                        break;
                    case WatchEventType.Modified:
                        break;
                    case WatchEventType.Deleted:
                        await hubContext.Clients.Group(KubernetesHub.WatchRolesString).SendAsync("NotifyNamespacedRoleDeleteEvent", v1Role.Metadata.Name);
                        break;
                    case WatchEventType.Error:
                        break;
                    default:
                        break;
                }
            });
        }

        public void StopNamespacedRoleWatch()
        {
            namespacedRoleWatcher.Dispose();
        }

        public async Task GetNamespacedServiceAccount(string @namespace)
        {
            var config = KubernetesClientConfiguration.BuildConfigFromConfigFile();

            IKubernetes client = new Kubernetes(config);

            var serviceAccountlistResp = await client.ListNamespacedServiceAccountWithHttpMessagesAsync(@namespace, watch: true);
            namespacedServiceAccountWatcher = serviceAccountlistResp.Watch<V1ServiceAccount, V1ServiceAccountList>(async (watchEventType, v1ServiceAccount) =>
            {
                switch (watchEventType)
                {
                    case WatchEventType.Added:
                        await hubContext.Clients.Group(KubernetesHub.WatchServiceAccountsString).SendAsync("NotifyNamespacedServiceAccountAddEvent", v1ServiceAccount.Metadata.Name);
                        break;
                    case WatchEventType.Modified:
                        break;
                    case WatchEventType.Deleted:
                        await hubContext.Clients.Group(KubernetesHub.WatchServiceAccountsString).SendAsync("NotifyNamespacedServiceAccountDeleteEvent", v1ServiceAccount.Metadata.Name);
                        break;
                    case WatchEventType.Error:
                        break;
                    default:
                        break;
                }
            });
        }

        public void StopNamespacedServiceAccountWatch()
        {
            namespacedServiceAccountWatcher.Dispose();
        }


        public async Task GetNamespacedRoleBinding(string @namespace)
        {
            var config = KubernetesClientConfiguration.BuildConfigFromConfigFile();

            IKubernetes client = new Kubernetes(config);

            var namespacedRoleBindingListResp = await client.ListNamespacedRoleBindingWithHttpMessagesAsync(@namespace, watch: true);
            namespacedRoleBindingWatcher = namespacedRoleBindingListResp.Watch<V1RoleBinding, V1RoleBindingList>(async (watchEventType, v1RoleBinding) =>
            {
                switch (watchEventType)
                {
                    case WatchEventType.Added:
                        await hubContext.Clients.Group(KubernetesHub.WatchRoleBindingsString).SendAsync("NotifyNamespacedRoleBindingAddEvent", v1RoleBinding.Metadata.Name);
                        break;
                    case WatchEventType.Modified:
                        break;
                    case WatchEventType.Deleted:
                        await hubContext.Clients.Group(KubernetesHub.WatchRoleBindingsString).SendAsync("NotifyNamespacedRoleBindingDeleteEvent", v1RoleBinding.Metadata.Name);
                        break;
                    case WatchEventType.Error:
                        break;
                    default:
                        break;
                }
            });
        }

        public void StopNamespacedRoleBindingWatch()
        {
            namespacedRoleBindingWatcher.Dispose();
        }

        public async Task GetRoleForAllNamespaces()
        {
            var config = KubernetesClientConfiguration.BuildConfigFromConfigFile();

            IKubernetes client = new Kubernetes(config);

            var roleListResp = await client.ListRoleForAllNamespacesWithHttpMessagesAsync(watch: true);
            roleWatcher = roleListResp.Watch<V1Role, V1RoleList>(async (watchEventType, v1Role) =>
            {
                switch (watchEventType)
                {
                    case WatchEventType.Added:
                        await hubContext.Clients.Group(KubernetesHub.WatchRolesString).SendAsync("NotifyRoleAddEvent", v1Role.Metadata.Name);
                        break;
                    case WatchEventType.Modified:
                        break;
                    case WatchEventType.Deleted:
                        await hubContext.Clients.Group(KubernetesHub.WatchRolesString).SendAsync("NotifyRoleDeleteEvent", v1Role.Metadata.Name);
                        break;
                    case WatchEventType.Error:
                        break;
                    default:
                        break;
                }
            });
        }

        public async Task GetClusterRoleBinding()
        {
            var config = KubernetesClientConfiguration.BuildConfigFromConfigFile();

            IKubernetes client = new Kubernetes(config);

            var clusterRoleListResp = await client.ListClusterRoleBindingWithHttpMessagesAsync(watch: true);
            clusterRoleBindingWatcher = clusterRoleListResp.Watch<V1ClusterRoleBinding, V1ClusterRoleBindingList>(async (watchEventType, v1ClusterRoleBinding) =>
            {
                switch (watchEventType)
                {
                    case WatchEventType.Added:
                        await hubContext.Clients.Group(KubernetesHub.WatchClusterRoleBindingsString).SendAsync("NotifyClusterRoleBindingAddEvent", v1ClusterRoleBinding.Metadata.Name);
                        break;
                    case WatchEventType.Modified:
                        break;
                    case WatchEventType.Deleted:
                        await hubContext.Clients.Group(KubernetesHub.WatchClusterRoleBindingsString).SendAsync("NotifyClusterRoleBindingDeleteEvent", v1ClusterRoleBinding.Metadata.Name);
                        break;
                    case WatchEventType.Error:
                        break;
                    default:
                        break;
                }
            });
        }

    }
}
