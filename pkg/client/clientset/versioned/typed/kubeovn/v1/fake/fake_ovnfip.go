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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/kubeovn/kube-ovn/pkg/apis/kubeovn/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOvnFips implements OvnFipInterface
type FakeOvnFips struct {
	Fake *FakeKubeovnV1
}

var ovnfipsResource = v1.SchemeGroupVersion.WithResource("ovn-fips")

var ovnfipsKind = v1.SchemeGroupVersion.WithKind("OvnFip")

// Get takes name of the ovnFip, and returns the corresponding ovnFip object, and an error if there is any.
func (c *FakeOvnFips) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.OvnFip, err error) {
	emptyResult := &v1.OvnFip{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(ovnfipsResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.OvnFip), err
}

// List takes label and field selectors, and returns the list of OvnFips that match those selectors.
func (c *FakeOvnFips) List(ctx context.Context, opts metav1.ListOptions) (result *v1.OvnFipList, err error) {
	emptyResult := &v1.OvnFipList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(ovnfipsResource, ovnfipsKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.OvnFipList{ListMeta: obj.(*v1.OvnFipList).ListMeta}
	for _, item := range obj.(*v1.OvnFipList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ovnFips.
func (c *FakeOvnFips) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(ovnfipsResource, opts))
}

// Create takes the representation of a ovnFip and creates it.  Returns the server's representation of the ovnFip, and an error, if there is any.
func (c *FakeOvnFips) Create(ctx context.Context, ovnFip *v1.OvnFip, opts metav1.CreateOptions) (result *v1.OvnFip, err error) {
	emptyResult := &v1.OvnFip{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(ovnfipsResource, ovnFip, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.OvnFip), err
}

// Update takes the representation of a ovnFip and updates it. Returns the server's representation of the ovnFip, and an error, if there is any.
func (c *FakeOvnFips) Update(ctx context.Context, ovnFip *v1.OvnFip, opts metav1.UpdateOptions) (result *v1.OvnFip, err error) {
	emptyResult := &v1.OvnFip{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(ovnfipsResource, ovnFip, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.OvnFip), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOvnFips) UpdateStatus(ctx context.Context, ovnFip *v1.OvnFip, opts metav1.UpdateOptions) (result *v1.OvnFip, err error) {
	emptyResult := &v1.OvnFip{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(ovnfipsResource, "status", ovnFip, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.OvnFip), err
}

// Delete takes name of the ovnFip and deletes it. Returns an error if one occurs.
func (c *FakeOvnFips) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(ovnfipsResource, name, opts), &v1.OvnFip{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOvnFips) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(ovnfipsResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.OvnFipList{})
	return err
}

// Patch applies the patch and returns the patched ovnFip.
func (c *FakeOvnFips) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.OvnFip, err error) {
	emptyResult := &v1.OvnFip{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(ovnfipsResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.OvnFip), err
}
