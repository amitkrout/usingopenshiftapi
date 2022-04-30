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
	)

	var k8sClient client.Client

	route := &routev1.Route{}
	err := k8sClient.Get(context.TODO(), types.NamespacedName{Name: routeName, Namespace: namespace}, route)
	if err != nil {
		fmt.Println("Some issue")
	}
	fmt.Println("Api call completed")
}
