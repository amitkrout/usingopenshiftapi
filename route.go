package main

import (
	"context"
	"fmt"

	routev1 "github.com/openshift/api/route/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func main() {

	const (
		namespace = "test"
		routeName = "nodejs-basic"
		// consoleLinkName = "argocd"
	)

	var k8sClient client.Client

	route := &routev1.Route{}
	// consoleLink := &console.ConsoleLink{}

	// err := k8sClient.Get(context.TODO(), types.NamespacedName{Name: consoleLinkName}, consoleLink)
	// if err != nil {
	// 	fmt.Println("link not found")
	// }
	err := k8sClient.Get(context.TODO(), types.NamespacedName{Name: routeName, Namespace: namespace}, route)
	if err != nil {
		fmt.Println("Route not found")
	}
	// clLink := strings.TrimLeft(consoleLink.Spec.Href, "https://")
	// routeLink := route.Spec.Host
	// if clLink != routeLink {
	// 	fmt.Println("URL mismatch, route: %s, consoleLink: %s", routeLink, clLink)
	// }
}
