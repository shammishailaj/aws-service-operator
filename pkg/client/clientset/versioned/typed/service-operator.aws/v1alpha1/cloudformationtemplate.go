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

package v1alpha1

import (
	v1alpha1 "github.com/awslabs/aws-service-operator/pkg/apis/service-operator.aws/v1alpha1"
	scheme "github.com/awslabs/aws-service-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CloudFormationTemplatesGetter has a method to return a CloudFormationTemplateInterface.
// A group's client should implement this interface.
type CloudFormationTemplatesGetter interface {
	CloudFormationTemplates(namespace string) CloudFormationTemplateInterface
}

// CloudFormationTemplateInterface has methods to work with CloudFormationTemplate resources.
type CloudFormationTemplateInterface interface {
	Create(*v1alpha1.CloudFormationTemplate) (*v1alpha1.CloudFormationTemplate, error)
	Update(*v1alpha1.CloudFormationTemplate) (*v1alpha1.CloudFormationTemplate, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CloudFormationTemplate, error)
	List(opts v1.ListOptions) (*v1alpha1.CloudFormationTemplateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudFormationTemplate, err error)
	CloudFormationTemplateExpansion
}

// cloudFormationTemplates implements CloudFormationTemplateInterface
type cloudFormationTemplates struct {
	client rest.Interface
	ns     string
}

// newCloudFormationTemplates returns a CloudFormationTemplates
func newCloudFormationTemplates(c *ServiceoperatorV1alpha1Client, namespace string) *cloudFormationTemplates {
	return &cloudFormationTemplates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cloudFormationTemplate, and returns the corresponding cloudFormationTemplate object, and an error if there is any.
func (c *cloudFormationTemplates) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudFormationTemplate, err error) {
	result = &v1alpha1.CloudFormationTemplate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudformationtemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudFormationTemplates that match those selectors.
func (c *cloudFormationTemplates) List(opts v1.ListOptions) (result *v1alpha1.CloudFormationTemplateList, err error) {
	result = &v1alpha1.CloudFormationTemplateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudformationtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudFormationTemplates.
func (c *cloudFormationTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cloudformationtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a cloudFormationTemplate and creates it.  Returns the server's representation of the cloudFormationTemplate, and an error, if there is any.
func (c *cloudFormationTemplates) Create(cloudFormationTemplate *v1alpha1.CloudFormationTemplate) (result *v1alpha1.CloudFormationTemplate, err error) {
	result = &v1alpha1.CloudFormationTemplate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cloudformationtemplates").
		Body(cloudFormationTemplate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cloudFormationTemplate and updates it. Returns the server's representation of the cloudFormationTemplate, and an error, if there is any.
func (c *cloudFormationTemplates) Update(cloudFormationTemplate *v1alpha1.CloudFormationTemplate) (result *v1alpha1.CloudFormationTemplate, err error) {
	result = &v1alpha1.CloudFormationTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudformationtemplates").
		Name(cloudFormationTemplate.Name).
		Body(cloudFormationTemplate).
		Do().
		Into(result)
	return
}

// Delete takes name of the cloudFormationTemplate and deletes it. Returns an error if one occurs.
func (c *cloudFormationTemplates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudformationtemplates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudFormationTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudformationtemplates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cloudFormationTemplate.
func (c *cloudFormationTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudFormationTemplate, err error) {
	result = &v1alpha1.CloudFormationTemplate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cloudformationtemplates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
