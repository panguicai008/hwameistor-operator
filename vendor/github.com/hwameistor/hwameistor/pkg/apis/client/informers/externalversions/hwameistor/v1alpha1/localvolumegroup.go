// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	versioned "github.com/hwameistor/hwameistor/pkg/apis/client/clientset/versioned"
	internalinterfaces "github.com/hwameistor/hwameistor/pkg/apis/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/hwameistor/hwameistor/pkg/apis/client/listers/hwameistor/v1alpha1"
	hwameistorv1alpha1 "github.com/hwameistor/hwameistor/pkg/apis/hwameistor/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// LocalVolumeGroupInformer provides access to a shared informer and lister for
// LocalVolumeGroups.
type LocalVolumeGroupInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.LocalVolumeGroupLister
}

type localVolumeGroupInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewLocalVolumeGroupInformer constructs a new informer for LocalVolumeGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewLocalVolumeGroupInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredLocalVolumeGroupInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredLocalVolumeGroupInformer constructs a new informer for LocalVolumeGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredLocalVolumeGroupInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HwameistorV1alpha1().LocalVolumeGroups().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HwameistorV1alpha1().LocalVolumeGroups().Watch(context.TODO(), options)
			},
		},
		&hwameistorv1alpha1.LocalVolumeGroup{},
		resyncPeriod,
		indexers,
	)
}

func (f *localVolumeGroupInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredLocalVolumeGroupInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *localVolumeGroupInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&hwameistorv1alpha1.LocalVolumeGroup{}, f.defaultInformer)
}

func (f *localVolumeGroupInformer) Lister() v1alpha1.LocalVolumeGroupLister {
	return v1alpha1.NewLocalVolumeGroupLister(f.Informer().GetIndexer())
}
