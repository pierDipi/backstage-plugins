/*
Copyright 2021 The Knative Authors

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

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1 "knative.dev/eventing/pkg/apis/eventing/v1"
)

// FakeTriggers implements TriggerInterface
type FakeTriggers struct {
	Fake *FakeEventingV1
	ns   string
}

var triggersResource = v1.SchemeGroupVersion.WithResource("triggers")

var triggersKind = v1.SchemeGroupVersion.WithKind("Trigger")

// Get takes name of the trigger, and returns the corresponding trigger object, and an error if there is any.
func (c *FakeTriggers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Trigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(triggersResource, c.ns, name), &v1.Trigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Trigger), err
}

// List takes label and field selectors, and returns the list of Triggers that match those selectors.
func (c *FakeTriggers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.TriggerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(triggersResource, triggersKind, c.ns, opts), &v1.TriggerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.TriggerList{ListMeta: obj.(*v1.TriggerList).ListMeta}
	for _, item := range obj.(*v1.TriggerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested triggers.
func (c *FakeTriggers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(triggersResource, c.ns, opts))

}

// Create takes the representation of a trigger and creates it.  Returns the server's representation of the trigger, and an error, if there is any.
func (c *FakeTriggers) Create(ctx context.Context, trigger *v1.Trigger, opts metav1.CreateOptions) (result *v1.Trigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(triggersResource, c.ns, trigger), &v1.Trigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Trigger), err
}

// Update takes the representation of a trigger and updates it. Returns the server's representation of the trigger, and an error, if there is any.
func (c *FakeTriggers) Update(ctx context.Context, trigger *v1.Trigger, opts metav1.UpdateOptions) (result *v1.Trigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(triggersResource, c.ns, trigger), &v1.Trigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Trigger), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTriggers) UpdateStatus(ctx context.Context, trigger *v1.Trigger, opts metav1.UpdateOptions) (*v1.Trigger, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(triggersResource, "status", c.ns, trigger), &v1.Trigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Trigger), err
}

// Delete takes name of the trigger and deletes it. Returns an error if one occurs.
func (c *FakeTriggers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(triggersResource, c.ns, name, opts), &v1.Trigger{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTriggers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(triggersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.TriggerList{})
	return err
}

// Patch applies the patch and returns the patched trigger.
func (c *FakeTriggers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Trigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(triggersResource, c.ns, name, pt, data, subresources...), &v1.Trigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Trigger), err
}