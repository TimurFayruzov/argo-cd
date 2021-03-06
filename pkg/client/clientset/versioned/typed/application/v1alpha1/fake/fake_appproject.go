package fake

import (
	v1alpha1 "github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAppProjects implements AppProjectInterface
type FakeAppProjects struct {
	Fake *FakeArgoprojV1alpha1
	ns   string
}

var appprojectsResource = schema.GroupVersionResource{Group: "argoproj.io", Version: "v1alpha1", Resource: "appprojects"}

var appprojectsKind = schema.GroupVersionKind{Group: "argoproj.io", Version: "v1alpha1", Kind: "AppProject"}

// Get takes name of the appProject, and returns the corresponding appProject object, and an error if there is any.
func (c *FakeAppProjects) Get(name string, options v1.GetOptions) (result *v1alpha1.AppProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(appprojectsResource, c.ns, name), &v1alpha1.AppProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppProject), err
}

// List takes label and field selectors, and returns the list of AppProjects that match those selectors.
func (c *FakeAppProjects) List(opts v1.ListOptions) (result *v1alpha1.AppProjectList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(appprojectsResource, appprojectsKind, c.ns, opts), &v1alpha1.AppProjectList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AppProjectList{}
	for _, item := range obj.(*v1alpha1.AppProjectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested appProjects.
func (c *FakeAppProjects) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(appprojectsResource, c.ns, opts))

}

// Create takes the representation of a appProject and creates it.  Returns the server's representation of the appProject, and an error, if there is any.
func (c *FakeAppProjects) Create(appProject *v1alpha1.AppProject) (result *v1alpha1.AppProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(appprojectsResource, c.ns, appProject), &v1alpha1.AppProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppProject), err
}

// Update takes the representation of a appProject and updates it. Returns the server's representation of the appProject, and an error, if there is any.
func (c *FakeAppProjects) Update(appProject *v1alpha1.AppProject) (result *v1alpha1.AppProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(appprojectsResource, c.ns, appProject), &v1alpha1.AppProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppProject), err
}

// Delete takes name of the appProject and deletes it. Returns an error if one occurs.
func (c *FakeAppProjects) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(appprojectsResource, c.ns, name), &v1alpha1.AppProject{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAppProjects) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(appprojectsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AppProjectList{})
	return err
}

// Patch applies the patch and returns the patched appProject.
func (c *FakeAppProjects) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(appprojectsResource, c.ns, name, data, subresources...), &v1alpha1.AppProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppProject), err
}
