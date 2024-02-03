package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gitlab.com/atarte/kubectl-display/cmd"
	"gitlab.com/atarte/kubectl-display/utils"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func getKubeClient() (*kubernetes.Clientset, error) {
	// Get user home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	// Get client based on the config
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homeDir, ".kube", "config"))
	if err != nil {
		return nil, err

	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return client, nil
}

type Namespace struct {
	UID  string
	Name string
}

// Probably will need to make a resource struct generic or taht will containt every body
type NamespacesRessource struct {
	NamespacesList []Namespace
}

// make a function convert v1.Namespace to namespace
func (n *NamespacesRessource) InsertNamespace(namespace *v1.Namespace) {
	n.NamespacesList = append(n.NamespacesList, Namespace{
		UID:  string(namespace.UID),
		Name: namespace.Name,
	})
}
func (n *NamespacesRessource) RemoveNamespace(UID string) {
	namespaceIndex, _ := n.GetRessourceIndex(UID)
	n.NamespacesList = append(n.NamespacesList[:namespaceIndex], n.NamespacesList[namespaceIndex+1:]...)
}
func (n *NamespacesRessource) UpdateNamespace(namespace *v1.Namespace) {
	namespaceIndex, _ := n.GetRessourceIndex(string(namespace.UID))
	n.NamespacesList[namespaceIndex] = Namespace{
		UID:  string(namespace.UID),
		Name: namespace.Name,
	}
}

func (n *NamespacesRessource) GetRessourceIndex(UID string) (int, error) {
	for i, ressource := range n.NamespacesList {
		if (ressource.UID) == UID {
			return i, nil
		}
	}
	return -1, errors.New("UID" + UID + "not found")
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

	namespaces := NamespacesRessource{}

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
		// fmt.Println("")

		DisplayNamespace(namespaces)
	}

}

// DisplayNamespace
// Later the param will be a struct dedicated to the ressource
func DisplayNamespace(namespaces NamespacesRessource) {
	utils.ClearScreen()

	for _, v := range namespaces.NamespacesList {
		fmt.Println(v.UID, v.Name)

	}

	utils.ExitLine()
}

func WatchPods() {

}

func main() {
	// utils.ClearScreen()

	fmt.Println("start")

	// client, err := getKubeClient()
	// if err != nil {
	// 	os.Exit(1)
	// }
	// go WatchNamespaces(client)

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

	// fmt.Scanln()
	// client, err := getKubeClient()
	// if err != nil {
	// 	os.Exit(1)
	// }

	// namespace := apiv1.NamespaceDefault

	// // podclient.CoreV1().Pods(namespace)

	// pods, err := client.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	// if err != nil {
	// 	err = fmt.Errorf("error getting pods: %v\n", err)
	// 	// return nil, err
	// }

	// for _, pod := range pods.Items {
	// 	fmt.Printf("Pod name: %v\n", pod.Name)
	// }

}
