/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	pispicodevv1alpha1 "github.com/pispico/digitalocean-operator/pkg/apis/pispico.dev/v1alpha1"
	versioned "github.com/pispico/digitalocean-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/pispico/digitalocean-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/pispico/digitalocean-operator/pkg/client/listers/pispico.dev/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KlusterInformer provides access to a shared informer and lister for
// Klusters.
type KlusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.KlusterLister
}

type klusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKlusterInformer constructs a new informer for Kluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKlusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKlusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKlusterInformer constructs a new informer for Kluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKlusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PispicoV1alpha1().Klusters(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PispicoV1alpha1().Klusters(namespace).Watch(context.TODO(), options)
			},
		},
		&pispicodevv1alpha1.Kluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *klusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKlusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *klusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&pispicodevv1alpha1.Kluster{}, f.defaultInformer)
}

func (f *klusterInformer) Lister() v1alpha1.KlusterLister {
	return v1alpha1.NewKlusterLister(f.Informer().GetIndexer())
}
