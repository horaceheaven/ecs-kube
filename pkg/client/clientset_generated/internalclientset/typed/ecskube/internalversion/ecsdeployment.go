/*
Copyright 2017 The Kubernetes Authors.

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
package internalversion

import (
	ecskube "github.com/robinpercy/ecs-kube/pkg/apis/ecskube"
	scheme "github.com/robinpercy/ecs-kube/pkg/client/clientset_generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ECSDeploymentsGetter has a method to return a ECSDeploymentInterface.
// A group's client should implement this interface.
type ECSDeploymentsGetter interface {
	ECSDeployments(namespace string) ECSDeploymentInterface
}

// ECSDeploymentInterface has methods to work with ECSDeployment resources.
type ECSDeploymentInterface interface {
	Create(*ecskube.ECSDeployment) (*ecskube.ECSDeployment, error)
	Update(*ecskube.ECSDeployment) (*ecskube.ECSDeployment, error)
	UpdateStatus(*ecskube.ECSDeployment) (*ecskube.ECSDeployment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*ecskube.ECSDeployment, error)
	List(opts v1.ListOptions) (*ecskube.ECSDeploymentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *ecskube.ECSDeployment, err error)
	ECSDeploymentExpansion
}

// eCSDeployments implements ECSDeploymentInterface
type eCSDeployments struct {
	client rest.Interface
	ns     string
}

// newECSDeployments returns a ECSDeployments
func newECSDeployments(c *EcskubeClient, namespace string) *eCSDeployments {
	return &eCSDeployments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the eCSDeployment, and returns the corresponding eCSDeployment object, and an error if there is any.
func (c *eCSDeployments) Get(name string, options v1.GetOptions) (result *ecskube.ECSDeployment, err error) {
	result = &ecskube.ECSDeployment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ecsdeployments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ECSDeployments that match those selectors.
func (c *eCSDeployments) List(opts v1.ListOptions) (result *ecskube.ECSDeploymentList, err error) {
	result = &ecskube.ECSDeploymentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ecsdeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested eCSDeployments.
func (c *eCSDeployments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ecsdeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a eCSDeployment and creates it.  Returns the server's representation of the eCSDeployment, and an error, if there is any.
func (c *eCSDeployments) Create(eCSDeployment *ecskube.ECSDeployment) (result *ecskube.ECSDeployment, err error) {
	result = &ecskube.ECSDeployment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ecsdeployments").
		Body(eCSDeployment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a eCSDeployment and updates it. Returns the server's representation of the eCSDeployment, and an error, if there is any.
func (c *eCSDeployments) Update(eCSDeployment *ecskube.ECSDeployment) (result *ecskube.ECSDeployment, err error) {
	result = &ecskube.ECSDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ecsdeployments").
		Name(eCSDeployment.Name).
		Body(eCSDeployment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *eCSDeployments) UpdateStatus(eCSDeployment *ecskube.ECSDeployment) (result *ecskube.ECSDeployment, err error) {
	result = &ecskube.ECSDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ecsdeployments").
		Name(eCSDeployment.Name).
		SubResource("status").
		Body(eCSDeployment).
		Do().
		Into(result)
	return
}

// Delete takes name of the eCSDeployment and deletes it. Returns an error if one occurs.
func (c *eCSDeployments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ecsdeployments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *eCSDeployments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ecsdeployments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched eCSDeployment.
func (c *eCSDeployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *ecskube.ECSDeployment, err error) {
	result = &ecskube.ECSDeployment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ecsdeployments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}