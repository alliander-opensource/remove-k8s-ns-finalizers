/*
Copyright 2016 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This code is based on https://github.com/kubernetes/client-go/tree/master/examples/out-of-cluster-client-configuration
package main

import (
	"flag"
	"fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %v <NAMESPACE>\n", os.Args[0])
		return
	}
	namespace := os.Args[1]

	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	nsClient := clientset.CoreV1().Namespaces()
	namespaceObj, err := nsClient.Get(namespace, metav1.GetOptions{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	namespaceObj.Spec.Finalizers = []v1.FinalizerName{}

	_, err = nsClient.Finalize(namespaceObj)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("OK, namespace should be removed")
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
