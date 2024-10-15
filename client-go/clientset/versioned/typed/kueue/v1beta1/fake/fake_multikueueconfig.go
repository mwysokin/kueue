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
	json "encoding/json"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1beta1 "sigs.k8s.io/kueue/apis/kueue/v1beta1"
	kueuev1beta1 "sigs.k8s.io/kueue/client-go/applyconfiguration/kueue/v1beta1"
)

// FakeMultiKueueConfigs implements MultiKueueConfigInterface
type FakeMultiKueueConfigs struct {
	Fake *FakeKueueV1beta1
}

var multikueueconfigsResource = v1beta1.SchemeGroupVersion.WithResource("multikueueconfigs")

var multikueueconfigsKind = v1beta1.SchemeGroupVersion.WithKind("MultiKueueConfig")

// Get takes name of the multiKueueConfig, and returns the corresponding multiKueueConfig object, and an error if there is any.
func (c *FakeMultiKueueConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.MultiKueueConfig, err error) {
	emptyResult := &v1beta1.MultiKueueConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(multikueueconfigsResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.MultiKueueConfig), err
}

// List takes label and field selectors, and returns the list of MultiKueueConfigs that match those selectors.
func (c *FakeMultiKueueConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.MultiKueueConfigList, err error) {
	emptyResult := &v1beta1.MultiKueueConfigList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(multikueueconfigsResource, multikueueconfigsKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.MultiKueueConfigList{ListMeta: obj.(*v1beta1.MultiKueueConfigList).ListMeta}
	for _, item := range obj.(*v1beta1.MultiKueueConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested multiKueueConfigs.
func (c *FakeMultiKueueConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(multikueueconfigsResource, opts))
}

// Create takes the representation of a multiKueueConfig and creates it.  Returns the server's representation of the multiKueueConfig, and an error, if there is any.
func (c *FakeMultiKueueConfigs) Create(ctx context.Context, multiKueueConfig *v1beta1.MultiKueueConfig, opts v1.CreateOptions) (result *v1beta1.MultiKueueConfig, err error) {
	emptyResult := &v1beta1.MultiKueueConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(multikueueconfigsResource, multiKueueConfig, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.MultiKueueConfig), err
}

// Update takes the representation of a multiKueueConfig and updates it. Returns the server's representation of the multiKueueConfig, and an error, if there is any.
func (c *FakeMultiKueueConfigs) Update(ctx context.Context, multiKueueConfig *v1beta1.MultiKueueConfig, opts v1.UpdateOptions) (result *v1beta1.MultiKueueConfig, err error) {
	emptyResult := &v1beta1.MultiKueueConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(multikueueconfigsResource, multiKueueConfig, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.MultiKueueConfig), err
}

// Delete takes name of the multiKueueConfig and deletes it. Returns an error if one occurs.
func (c *FakeMultiKueueConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(multikueueconfigsResource, name, opts), &v1beta1.MultiKueueConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMultiKueueConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(multikueueconfigsResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.MultiKueueConfigList{})
	return err
}

// Patch applies the patch and returns the patched multiKueueConfig.
func (c *FakeMultiKueueConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.MultiKueueConfig, err error) {
	emptyResult := &v1beta1.MultiKueueConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(multikueueconfigsResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.MultiKueueConfig), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied multiKueueConfig.
func (c *FakeMultiKueueConfigs) Apply(ctx context.Context, multiKueueConfig *kueuev1beta1.MultiKueueConfigApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.MultiKueueConfig, err error) {
	if multiKueueConfig == nil {
		return nil, fmt.Errorf("multiKueueConfig provided to Apply must not be nil")
	}
	data, err := json.Marshal(multiKueueConfig)
	if err != nil {
		return nil, err
	}
	name := multiKueueConfig.Name
	if name == nil {
		return nil, fmt.Errorf("multiKueueConfig.Name must be provided to Apply")
	}
	emptyResult := &v1beta1.MultiKueueConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(multikueueconfigsResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.MultiKueueConfig), err
}