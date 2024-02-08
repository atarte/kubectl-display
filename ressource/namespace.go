package ressource

import (
	"errors"
	"fmt"

	"gitlab.com/atarte/kubectl-display/utils"
	v1 "k8s.io/api/core/v1"
)

// Probably will need to make a resource struct generic or taht will containt every body
type NamespacesRessource struct {
	NamespacesList []Namespace
}

type Namespace struct {
	UID               string
	Name              string
	Status            string
	CreationTimestamp int
}

// make a function convert v1.Namespace to namespace
func (n *NamespacesRessource) InsertNamespace(namespace *v1.Namespace) {
	n.NamespacesList = append(n.NamespacesList, Namespace{
		UID:               string(namespace.UID),
		Name:              namespace.Name,
		Status:            string(namespace.Status.Phase),
		CreationTimestamp: namespace.CreationTimestamp.Nanosecond(),
	})
}
func (n *NamespacesRessource) RemoveNamespace(UID string) {
	namespaceIndex, _ := n.GetRessourceIndex(UID)
	n.NamespacesList = append(n.NamespacesList[:namespaceIndex], n.NamespacesList[namespaceIndex+1:]...)
}
func (n *NamespacesRessource) UpdateNamespace(namespace *v1.Namespace) {
	namespaceIndex, _ := n.GetRessourceIndex(string(namespace.UID))
	n.NamespacesList[namespaceIndex] = Namespace{
		UID:               string(namespace.UID),
		Name:              namespace.Name,
		Status:            string(namespace.Status.Phase),
		CreationTimestamp: namespace.CreationTimestamp.Nanosecond(),
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

// DisplayTable
func (n *NamespacesRessource) Display() {
	utils.ClearScreen()

	// need to make this pretty but to tired for now
	fmt.Println("NAME   STATUS   AGE")
	for _, namespace := range n.NamespacesList {
		fmt.Printf("%s   %s   %dnano-seconde\n", namespace.Name, namespace.Status, namespace.CreationTimestamp)
	}

	utils.ExitLine()
}
