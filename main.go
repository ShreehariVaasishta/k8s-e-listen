package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1" // Import for metav1 types
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	namespace := flag.String("namespace", "default", "namespace to listen for events")
	eventType := flag.String("event-type", "Normal", "type of event to filter")
	reason := flag.String("reason", "", "specific reason to filter events")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	watchEvents(clientset, *namespace, *eventType, *reason)
}

func watchEvents(clientset *kubernetes.Clientset, namespace, eventType, reason string) {
	watchlist := clientset.CoreV1().Events(namespace)
	options := metav1.ListOptions{ // Updated to metav1
		FieldSelector: fields.OneTermEqualSelector("type", eventType).String(),
	}

	if reason != "" {
		options.FieldSelector += "," + fields.OneTermEqualSelector("reason", reason).String()
	}

	watchInterface, err := watchlist.Watch(context.Background(), options)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Listening for events in namespace %s\n", namespace)

	for eventObj := range watchInterface.ResultChan() {
		event, ok := eventObj.Object.(*corev1.Event) // Updated to corev1.Event
		if !ok {
			fmt.Printf("unexpected type %T\n", eventObj.Object)
			continue
		}

		// Calculating the age of the event
		age := time.Since(event.LastTimestamp.Time).Round(time.Second)

		fmt.Printf("Event Type: %s\n", event.Type)
		fmt.Printf("Message: %s\n", event.Message)
		fmt.Printf("Namespace: %s\n", event.Namespace)
		fmt.Printf("Source: %s\n", event.Source.Component)
		fmt.Printf("Age: %v\n\n", age)
	}
}
