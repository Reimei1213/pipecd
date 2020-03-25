// Copyright 2020 The PipeCD Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/kapetaniosci/pipe/pkg/crd/apis/pipeline/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePipelines implements PipelineInterface
type FakePipelines struct {
	Fake *FakePipecdV1beta1
	ns   string
}

var pipelinesResource = schema.GroupVersionResource{Group: "pipecd.dev", Version: "v1beta1", Resource: "pipelines"}

var pipelinesKind = schema.GroupVersionKind{Group: "pipecd.dev", Version: "v1beta1", Kind: "Pipeline"}

// Get takes name of the pipeline, and returns the corresponding pipeline object, and an error if there is any.
func (c *FakePipelines) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Pipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pipelinesResource, c.ns, name), &v1beta1.Pipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Pipeline), err
}

// List takes label and field selectors, and returns the list of Pipelines that match those selectors.
func (c *FakePipelines) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.PipelineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pipelinesResource, pipelinesKind, c.ns, opts), &v1beta1.PipelineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.PipelineList{ListMeta: obj.(*v1beta1.PipelineList).ListMeta}
	for _, item := range obj.(*v1beta1.PipelineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pipelines.
func (c *FakePipelines) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pipelinesResource, c.ns, opts))

}

// Create takes the representation of a pipeline and creates it.  Returns the server's representation of the pipeline, and an error, if there is any.
func (c *FakePipelines) Create(ctx context.Context, pipeline *v1beta1.Pipeline, opts v1.CreateOptions) (result *v1beta1.Pipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pipelinesResource, c.ns, pipeline), &v1beta1.Pipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Pipeline), err
}

// Update takes the representation of a pipeline and updates it. Returns the server's representation of the pipeline, and an error, if there is any.
func (c *FakePipelines) Update(ctx context.Context, pipeline *v1beta1.Pipeline, opts v1.UpdateOptions) (result *v1beta1.Pipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pipelinesResource, c.ns, pipeline), &v1beta1.Pipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Pipeline), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePipelines) UpdateStatus(ctx context.Context, pipeline *v1beta1.Pipeline, opts v1.UpdateOptions) (*v1beta1.Pipeline, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pipelinesResource, "status", c.ns, pipeline), &v1beta1.Pipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Pipeline), err
}

// Delete takes name of the pipeline and deletes it. Returns an error if one occurs.
func (c *FakePipelines) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pipelinesResource, c.ns, name), &v1beta1.Pipeline{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePipelines) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pipelinesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.PipelineList{})
	return err
}

// Patch applies the patch and returns the patched pipeline.
func (c *FakePipelines) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Pipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pipelinesResource, c.ns, name, pt, data, subresources...), &v1beta1.Pipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Pipeline), err
}
