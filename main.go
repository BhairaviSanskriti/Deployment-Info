package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var namespace string
var kubeconfig string

var rootCmd = &cobra.Command{
	Use:   "k8s-info",
	Short: "Retrieve deployment data from a Kubernetes cluster",
	Run:   run,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "default", "the namespace to retrieve deployments from")
	rootCmd.PersistentFlags().StringVarP(&kubeconfig, "kubeconfig", "k", "", "absolute path to the kubeconfig file")
}

type DeploymentData struct {
	Name              string `json:"name"`
	HealthyReplicas   int32  `json:"healthy_replicas"`
	UnhealthyReplicas int32  `json:"unhealthy_replicas"`
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func run(cmd *cobra.Command, args []string) {
	config, err := getClientConfig(kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.Background(), v1.ListOptions{})

	if err != nil {
		log.Fatal(err)
	}

	data := []DeploymentData{}

	for _, deployment := range deployments.Items {
		healthyReplicas := deployment.Status.AvailableReplicas
		unhealthyReplicas := *deployment.Spec.Replicas - deployment.Status.AvailableReplicas

		data = append(data, DeploymentData{
			Name: deployment.Name,

			UnhealthyReplicas: unhealthyReplicas,
			HealthyReplicas:   healthyReplicas,
		})
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
}

func getClientConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	//check config file path in $KUBECONFIG variable
	kubeconfig = os.Getenv("KUBECONFIG")
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	//check config file in ~/.kube/config
	kubeconfig = filepath.Join(os.Getenv("HOME"), ".kube", "config")
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	//check config file in the current directory
	kubeconfig = filepath.Join(os.Getenv("PWD"), "kubeconfig")
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	// set KUBERNETES_SERVICE_HOST and KUBERNETES_SERVICE_PORT in case the function uses in-cluster info
	return rest.InClusterConfig()

}
