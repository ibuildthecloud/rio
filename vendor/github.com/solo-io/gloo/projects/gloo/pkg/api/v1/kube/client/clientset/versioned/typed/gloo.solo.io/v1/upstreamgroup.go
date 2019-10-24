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

package v1

import (
	"time"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/kube/apis/gloo.solo.io/v1"
	scheme "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/kube/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// UpstreamGroupsGetter has a method to return a UpstreamGroupInterface.
// A group's client should implement this interface.
type UpstreamGroupsGetter interface {
	UpstreamGroups(namespace string) UpstreamGroupInterface
}

// UpstreamGroupInterface has methods to work with UpstreamGroup resources.
type UpstreamGroupInterface interface {
	Create(*v1.UpstreamGroup) (*v1.UpstreamGroup, error)
	Update(*v1.UpstreamGroup) (*v1.UpstreamGroup, error)
	UpdateStatus(*v1.UpstreamGroup) (*v1.UpstreamGroup, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.UpstreamGroup, error)
	List(opts metav1.ListOptions) (*v1.UpstreamGroupList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.UpstreamGroup, err error)
	UpstreamGroupExpansion
}

// upstreamGroups implements UpstreamGroupInterface
type upstreamGroups struct {
	client rest.Interface
	ns     string
}

// newUpstreamGroups returns a UpstreamGroups
func newUpstreamGroups(c *GlooV1Client, namespace string) *upstreamGroups {
	return &upstreamGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the upstreamGroup, and returns the corresponding upstreamGroup object, and an error if there is any.
func (c *upstreamGroups) Get(name string, options metav1.GetOptions) (result *v1.UpstreamGroup, err error) {
	result = &v1.UpstreamGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("upstreamgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of UpstreamGroups that match those selectors.
func (c *upstreamGroups) List(opts metav1.ListOptions) (result *v1.UpstreamGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.UpstreamGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("upstreamgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested upstreamGroups.
func (c *upstreamGroups) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("upstreamgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a upstreamGroup and creates it.  Returns the server's representation of the upstreamGroup, and an error, if there is any.
func (c *upstreamGroups) Create(upstreamGroup *v1.UpstreamGroup) (result *v1.UpstreamGroup, err error) {
	result = &v1.UpstreamGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("upstreamgroups").
		Body(upstreamGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a upstreamGroup and updates it. Returns the server's representation of the upstreamGroup, and an error, if there is any.
func (c *upstreamGroups) Update(upstreamGroup *v1.UpstreamGroup) (result *v1.UpstreamGroup, err error) {
	result = &v1.UpstreamGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("upstreamgroups").
		Name(upstreamGroup.Name).
		Body(upstreamGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *upstreamGroups) UpdateStatus(upstreamGroup *v1.UpstreamGroup) (result *v1.UpstreamGroup, err error) {
	result = &v1.UpstreamGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("upstreamgroups").
		Name(upstreamGroup.Name).
		SubResource("status").
		Body(upstreamGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the upstreamGroup and deletes it. Returns an error if one occurs.
func (c *upstreamGroups) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("upstreamgroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *upstreamGroups) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("upstreamgroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched upstreamGroup.
func (c *upstreamGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.UpstreamGroup, err error) {
	result = &v1.UpstreamGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("upstreamgroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}