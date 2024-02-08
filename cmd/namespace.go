package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/atarte/kubectl-display/ressource"
	"gitlab.com/atarte/kubectl-display/utils"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

var NamespaceCmd = &cobra.Command{
	Use:     "namespace",
	Aliases: []string{"namespace"},
	Short:   "Display kubernetes namespaces",
	Long:    `Display kubernetes namespaces, refreshing the display every time an update is detected`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Namespace subcommande")
		DisplayNamespace()
	},
}

func DisplayNamespace() {
	utils.ClearScreen()

	client, err := utils.GetKubeClient()
	if err != nil {
		os.Exit(1)
	}
	go WatchNamespaces(client)

	fmt.Scanln()
}

// Will need to define the list of ressource suported
// WatchNamespaces
func WatchNamespaces(client *kubernetes.Clientset) {
	timeOut := int64(60)
	namespacesWatcher, err := client.CoreV1().Namespaces().Watch(context.Background(), metav1.ListOptions{TimeoutSeconds: &timeOut})
	if err != nil {
		fmt.Println("fail for somme reason")
		return
	}

	namespaces := ressource.NamespacesRessource{}

	for event := range namespacesWatcher.ResultChan() {
		item := event.Object.(*corev1.Namespace)

		switch event.Type {
		case watch.Modified:
			namespaces.UpdateNamespace(item)
		case watch.Bookmark:
		case watch.Error:
		case watch.Deleted:
			namespaces.RemoveNamespace(string(item.UID))
		case watch.Added:
			namespaces.InsertNamespace(item)
		}

		// fmt.Printf("%T\n", item)
		// fmt.Printf("%#v\n", event)
		// fmt.Printf("%#v\n", item)
		// fmt.Println("")i

		namespaces.Display()

		// DisplayNamespace(namespaces)
	}

}

// DisplayNamespace
// Later the param will be a struct dedicated to the ressource
// func DisplayNamespace(namespaces NamespacesRessource) {
// 	utils.ClearScreen()

// 	for _, v := range namespaces.NamespacesList {
// 		fmt.Println(v.UID, v.Name)

// 	}

// 	utils.ExitLine()
// }
