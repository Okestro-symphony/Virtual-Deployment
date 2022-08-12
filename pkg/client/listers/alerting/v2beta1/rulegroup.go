/*
Copyright 2020 The KubeSphere Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v2beta1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v2beta1 "kubesphere.io/api/alerting/v2beta1"
)

// RuleGroupLister helps list RuleGroups.
// All objects returned here must be treated as read-only.
type RuleGroupLister interface {
	// List lists all RuleGroups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2beta1.RuleGroup, err error)
	// RuleGroups returns an object that can list and get RuleGroups.
	RuleGroups(namespace string) RuleGroupNamespaceLister
	RuleGroupListerExpansion
}

// ruleGroupLister implements the RuleGroupLister interface.
type ruleGroupLister struct {
	indexer cache.Indexer
}

// NewRuleGroupLister returns a new RuleGroupLister.
func NewRuleGroupLister(indexer cache.Indexer) RuleGroupLister {
	return &ruleGroupLister{indexer: indexer}
}

// List lists all RuleGroups in the indexer.
func (s *ruleGroupLister) List(selector labels.Selector) (ret []*v2beta1.RuleGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2beta1.RuleGroup))
	})
	return ret, err
}

// RuleGroups returns an object that can list and get RuleGroups.
func (s *ruleGroupLister) RuleGroups(namespace string) RuleGroupNamespaceLister {
	return ruleGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RuleGroupNamespaceLister helps list and get RuleGroups.
// All objects returned here must be treated as read-only.
type RuleGroupNamespaceLister interface {
	// List lists all RuleGroups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2beta1.RuleGroup, err error)
	// Get retrieves the RuleGroup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v2beta1.RuleGroup, error)
	RuleGroupNamespaceListerExpansion
}

// ruleGroupNamespaceLister implements the RuleGroupNamespaceLister
// interface.
type ruleGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RuleGroups in the indexer for a given namespace.
func (s ruleGroupNamespaceLister) List(selector labels.Selector) (ret []*v2beta1.RuleGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v2beta1.RuleGroup))
	})
	return ret, err
}

// Get retrieves the RuleGroup from the indexer for a given namespace and name.
func (s ruleGroupNamespaceLister) Get(name string) (*v2beta1.RuleGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2beta1.Resource("rulegroup"), name)
	}
	return obj.(*v2beta1.RuleGroup), nil
}
