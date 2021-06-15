/** Copyright 2020-2021 Alibaba Group Holding Limited.

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
	"time"

	v1alpha1 "github.com/v6d-io/v6d/k8s/api/k8s/v1alpha1"
	scheme "github.com/v6d-io/v6d/k8s/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// LocalObjectsGetter has a method to return a LocalObjectInterface.
// A group's client should implement this interface.
type LocalObjectsGetter interface {
	LocalObjects(namespace string) LocalObjectInterface
}

// LocalObjectInterface has methods to work with LocalObject resources.
type LocalObjectInterface interface {
	Create(ctx context.Context, localObject *v1alpha1.LocalObject, opts v1.CreateOptions) (*v1alpha1.LocalObject, error)
	Update(ctx context.Context, localObject *v1alpha1.LocalObject, opts v1.UpdateOptions) (*v1alpha1.LocalObject, error)
	UpdateStatus(ctx context.Context, localObject *v1alpha1.LocalObject, opts v1.UpdateOptions) (*v1alpha1.LocalObject, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.LocalObject, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.LocalObjectList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LocalObject, err error)
	LocalObjectExpansion
}

// localObjects implements LocalObjectInterface
type localObjects struct {
	client rest.Interface
	ns     string
}

// newLocalObjects returns a LocalObjects
func newLocalObjects(c *K8sV1alpha1Client, namespace string) *localObjects {
	return &localObjects{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the localObject, and returns the corresponding localObject object, and an error if there is any.
func (c *localObjects) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LocalObject, err error) {
	result = &v1alpha1.LocalObject{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("localobjects").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LocalObjects that match those selectors.
func (c *localObjects) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LocalObjectList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LocalObjectList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("localobjects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested localObjects.
func (c *localObjects) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("localobjects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a localObject and creates it.  Returns the server's representation of the localObject, and an error, if there is any.
func (c *localObjects) Create(ctx context.Context, localObject *v1alpha1.LocalObject, opts v1.CreateOptions) (result *v1alpha1.LocalObject, err error) {
	result = &v1alpha1.LocalObject{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("localobjects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(localObject).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a localObject and updates it. Returns the server's representation of the localObject, and an error, if there is any.
func (c *localObjects) Update(ctx context.Context, localObject *v1alpha1.LocalObject, opts v1.UpdateOptions) (result *v1alpha1.LocalObject, err error) {
	result = &v1alpha1.LocalObject{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("localobjects").
		Name(localObject.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(localObject).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *localObjects) UpdateStatus(ctx context.Context, localObject *v1alpha1.LocalObject, opts v1.UpdateOptions) (result *v1alpha1.LocalObject, err error) {
	result = &v1alpha1.LocalObject{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("localobjects").
		Name(localObject.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(localObject).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the localObject and deletes it. Returns an error if one occurs.
func (c *localObjects) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("localobjects").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *localObjects) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("localobjects").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched localObject.
func (c *localObjects) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LocalObject, err error) {
	result = &v1alpha1.LocalObject{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("localobjects").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}