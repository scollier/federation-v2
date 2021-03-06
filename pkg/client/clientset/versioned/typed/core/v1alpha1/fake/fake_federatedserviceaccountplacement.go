/*
Copyright 2018 The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/apis/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFederatedServiceAccountPlacements implements FederatedServiceAccountPlacementInterface
type FakeFederatedServiceAccountPlacements struct {
	Fake *FakeCoreV1alpha1
	ns   string
}

var federatedserviceaccountplacementsResource = schema.GroupVersionResource{Group: "core.federation.k8s.io", Version: "v1alpha1", Resource: "federatedserviceaccountplacements"}

var federatedserviceaccountplacementsKind = schema.GroupVersionKind{Group: "core.federation.k8s.io", Version: "v1alpha1", Kind: "FederatedServiceAccountPlacement"}

// Get takes name of the federatedServiceAccountPlacement, and returns the corresponding federatedServiceAccountPlacement object, and an error if there is any.
func (c *FakeFederatedServiceAccountPlacements) Get(name string, options v1.GetOptions) (result *v1alpha1.FederatedServiceAccountPlacement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(federatedserviceaccountplacementsResource, c.ns, name), &v1alpha1.FederatedServiceAccountPlacement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedServiceAccountPlacement), err
}

// List takes label and field selectors, and returns the list of FederatedServiceAccountPlacements that match those selectors.
func (c *FakeFederatedServiceAccountPlacements) List(opts v1.ListOptions) (result *v1alpha1.FederatedServiceAccountPlacementList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(federatedserviceaccountplacementsResource, federatedserviceaccountplacementsKind, c.ns, opts), &v1alpha1.FederatedServiceAccountPlacementList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FederatedServiceAccountPlacementList{}
	for _, item := range obj.(*v1alpha1.FederatedServiceAccountPlacementList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested federatedServiceAccountPlacements.
func (c *FakeFederatedServiceAccountPlacements) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(federatedserviceaccountplacementsResource, c.ns, opts))

}

// Create takes the representation of a federatedServiceAccountPlacement and creates it.  Returns the server's representation of the federatedServiceAccountPlacement, and an error, if there is any.
func (c *FakeFederatedServiceAccountPlacements) Create(federatedServiceAccountPlacement *v1alpha1.FederatedServiceAccountPlacement) (result *v1alpha1.FederatedServiceAccountPlacement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(federatedserviceaccountplacementsResource, c.ns, federatedServiceAccountPlacement), &v1alpha1.FederatedServiceAccountPlacement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedServiceAccountPlacement), err
}

// Update takes the representation of a federatedServiceAccountPlacement and updates it. Returns the server's representation of the federatedServiceAccountPlacement, and an error, if there is any.
func (c *FakeFederatedServiceAccountPlacements) Update(federatedServiceAccountPlacement *v1alpha1.FederatedServiceAccountPlacement) (result *v1alpha1.FederatedServiceAccountPlacement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(federatedserviceaccountplacementsResource, c.ns, federatedServiceAccountPlacement), &v1alpha1.FederatedServiceAccountPlacement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedServiceAccountPlacement), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFederatedServiceAccountPlacements) UpdateStatus(federatedServiceAccountPlacement *v1alpha1.FederatedServiceAccountPlacement) (*v1alpha1.FederatedServiceAccountPlacement, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(federatedserviceaccountplacementsResource, "status", c.ns, federatedServiceAccountPlacement), &v1alpha1.FederatedServiceAccountPlacement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedServiceAccountPlacement), err
}

// Delete takes name of the federatedServiceAccountPlacement and deletes it. Returns an error if one occurs.
func (c *FakeFederatedServiceAccountPlacements) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(federatedserviceaccountplacementsResource, c.ns, name), &v1alpha1.FederatedServiceAccountPlacement{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFederatedServiceAccountPlacements) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(federatedserviceaccountplacementsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.FederatedServiceAccountPlacementList{})
	return err
}

// Patch applies the patch and returns the patched federatedServiceAccountPlacement.
func (c *FakeFederatedServiceAccountPlacements) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FederatedServiceAccountPlacement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(federatedserviceaccountplacementsResource, c.ns, name, data, subresources...), &v1alpha1.FederatedServiceAccountPlacement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedServiceAccountPlacement), err
}
