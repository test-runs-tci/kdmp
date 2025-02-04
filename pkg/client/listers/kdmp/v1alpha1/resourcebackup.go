/*

LICENSE

*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/portworx/kdmp/pkg/apis/kdmp/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ResourceBackupLister helps list ResourceBackups.
// All objects returned here must be treated as read-only.
type ResourceBackupLister interface {
	// List lists all ResourceBackups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ResourceBackup, err error)
	// ResourceBackups returns an object that can list and get ResourceBackups.
	ResourceBackups(namespace string) ResourceBackupNamespaceLister
	ResourceBackupListerExpansion
}

// resourceBackupLister implements the ResourceBackupLister interface.
type resourceBackupLister struct {
	indexer cache.Indexer
}

// NewResourceBackupLister returns a new ResourceBackupLister.
func NewResourceBackupLister(indexer cache.Indexer) ResourceBackupLister {
	return &resourceBackupLister{indexer: indexer}
}

// List lists all ResourceBackups in the indexer.
func (s *resourceBackupLister) List(selector labels.Selector) (ret []*v1alpha1.ResourceBackup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ResourceBackup))
	})
	return ret, err
}

// ResourceBackups returns an object that can list and get ResourceBackups.
func (s *resourceBackupLister) ResourceBackups(namespace string) ResourceBackupNamespaceLister {
	return resourceBackupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ResourceBackupNamespaceLister helps list and get ResourceBackups.
// All objects returned here must be treated as read-only.
type ResourceBackupNamespaceLister interface {
	// List lists all ResourceBackups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ResourceBackup, err error)
	// Get retrieves the ResourceBackup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ResourceBackup, error)
	ResourceBackupNamespaceListerExpansion
}

// resourceBackupNamespaceLister implements the ResourceBackupNamespaceLister
// interface.
type resourceBackupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ResourceBackups in the indexer for a given namespace.
func (s resourceBackupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ResourceBackup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ResourceBackup))
	})
	return ret, err
}

// Get retrieves the ResourceBackup from the indexer for a given namespace and name.
func (s resourceBackupNamespaceLister) Get(name string) (*v1alpha1.ResourceBackup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("resourcebackup"), name)
	}
	return obj.(*v1alpha1.ResourceBackup), nil
}
