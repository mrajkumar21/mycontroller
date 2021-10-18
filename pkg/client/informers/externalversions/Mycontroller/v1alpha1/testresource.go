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
	Mycontrollerv1alpha1 "mycontroller/pkg/apis/Mycontroller/v1alpha1"
	versioned "mycontroller/pkg/client/clientset/versioned"
	internalinterfaces "mycontroller/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "mycontroller/pkg/client/listers/Mycontroller/v1alpha1"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TestResourceInformer provides access to a shared informer and lister for
// TestResources.
type TestResourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.TestResourceLister
}

type testResourceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTestResourceInformer constructs a new informer for TestResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTestResourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTestResourceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTestResourceInformer constructs a new informer for TestResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTestResourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MycontrollerV1alpha1().TestResources(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MycontrollerV1alpha1().TestResources(namespace).Watch(context.TODO(), options)
			},
		},
		&Mycontrollerv1alpha1.TestResource{},
		resyncPeriod,
		indexers,
	)
}

func (f *testResourceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTestResourceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *testResourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&Mycontrollerv1alpha1.TestResource{}, f.defaultInformer)
}

func (f *testResourceInformer) Lister() v1alpha1.TestResourceLister {
	return v1alpha1.NewTestResourceLister(f.Informer().GetIndexer())
}
