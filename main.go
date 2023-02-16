/*
Copyright 2017 The Kubernetes Authors.

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

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	doclient "github.com/pispico/digitalocean-operator/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Printf("Building config from flags, %s", err.Error())
	}

	doclientset, err := doclient.NewForConfig(config)
	if err != nil {
		log.Printf("getting doclientset %s\n", err.Error())
	}

	fmt.Println(doclientset)

	klusters, err := doclientset.PispicoV1alpha1().Klusters("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Printf("listing klusters%s\n", err.Error())
	}
	fmt.Printf("Number of klusters is %d\n", len(klusters.Items))
	fmt.Printf("Name of klusters is %s\n", klusters.Items[0].Name)

}
