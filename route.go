package main

import (
	"context"
	"fmt"

	routev1 "github.com/openshift/api/route/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
)

func main() {

	const (
		namespace = "test"
		routeName = "nodejs-basic"
	)

	var cfg *rest.Config
	var testEnv *envtest.Environment
	var k8sClient client.Client

	useActualCluster := true
	testEnv = &envtest.Environment{
		UseExistingCluster: &useActualCluster, // use an actual OpenShift cluster specified in kubeconfig
	}

	cfg, err := testEnv.Start()

	if err != nil || cfg != nil {
		fmt.Println("Some issue with the test env")
	}

	k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})

	if err != nil || k8sClient != nil {
		fmt.Println("Some issue with initializing the client")
	}

	route := &routev1.Route{}
	err = k8sClient.Get(context.TODO(), types.NamespacedName{Name: routeName, Namespace: namespace}, route)
	if err != nil {
		fmt.Println("Some issue")
	}
	fmt.Println("Api call completed")
}
