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

package v1alpha1

import (
	"context"
	v1alpha1 "mycontroller/pkg/apis/Mycontroller/v1alpha1"
	scheme "mycontroller/pkg/client/clientset/versioned/scheme"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TestResourcesGetter has a method to return a TestResourceInterface.
// A group's client should implement this interface.
type TestResourcesGetter interface {
	TestResources(namespace string) TestResourceInterface
}

// TestResourceInterface has methods to work with TestResource resources.
type TestResourceInterface interface {
	Create(ctx context.Context, testResource *v1alpha1.TestResource, opts v1.CreateOptions) (*v1alpha1.TestResource, error)
	Update(ctx context.Context, testResource *v1alpha1.TestResource, opts v1.UpdateOptions) (*v1alpha1.TestResource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.TestResource, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.TestResourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TestResource, err error)
	TestResourceExpansion
}

// testResources implements TestResourceInterface
type testResources struct {
	client rest.Interface
	ns     string
}

// newTestResources returns a TestResources
func newTestResources(c *MycontrollerV1alpha1Client, namespace string) *testResources {
	return &testResources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the testResource, and returns the corresponding testResource object, and an error if there is any.
func (c *testResources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TestResource, err error) {
	result = &v1alpha1.TestResource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("testresources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TestResources that match those selectors.
func (c *testResources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TestResourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.TestResourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("testresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested testResources.
func (c *testResources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("testresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a testResource and creates it.  Returns the server's representation of the testResource, and an error, if there is any.
func (c *testResources) Create(ctx context.Context, testResource *v1alpha1.TestResource, opts v1.CreateOptions) (result *v1alpha1.TestResource, err error) {
	result = &v1alpha1.TestResource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("testresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(testResource).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a testResource and updates it. Returns the server's representation of the testResource, and an error, if there is any.
func (c *testResources) Update(ctx context.Context, testResource *v1alpha1.TestResource, opts v1.UpdateOptions) (result *v1alpha1.TestResource, err error) {
	result = &v1alpha1.TestResource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("testresources").
		Name(testResource.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(testResource).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the testResource and deletes it. Returns an error if one occurs.
func (c *testResources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("testresources").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *testResources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("testresources").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched testResource.
func (c *testResources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TestResource, err error) {
	result = &v1alpha1.TestResource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("testresources").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
