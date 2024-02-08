package main

import (
	"fmt"
	"log"

	"gitlab.com/atarte/kubectl-display/cmd"
)

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
