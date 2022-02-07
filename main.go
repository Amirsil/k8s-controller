// Note: the example only works with the code within the same release/branch.
package main

import (
	"context"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/pkg/api/v1"
        "k8s.io/client-go/tools/clientcmd"
    	"k8s.io/client-go/tools/cache"
     	"k8s.io/client-go/pkg/fields"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	watchlist := cache.NewListWatchFromClient(
		clientset.Core().RESTClient(), 
		"services", v1.NamespaceDefault,
        	fields.Everything())
	
        _, controller := cache.NewInformer(
		watchlist,
		&v1.Service{},
		time.Second * 0,
		cache.ResourceEventHandlerFuncs{
		    AddFunc: func(obj interface{}) {
			fmt.Printf("service added: %s \n", obj)
		    },
		    DeleteFunc: func(obj interface{}) {
			fmt.Printf("service deleted: %s \n", obj)
		    },
		    UpdateFunc:func(oldObj, newObj interface{}) {
			fmt.Printf("service changed \n")
		    },
	        }
	)
        stop := make(chan struct{})
        go controller.Run(stop)
        for {
	    time.Sleep(time.Second)
        }
}
