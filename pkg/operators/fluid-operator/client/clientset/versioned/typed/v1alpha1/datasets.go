package v1alpha1

import (
	v1alpha1 "github.com/kubeflow/arena/pkg/operators/fluid-operator/apis/v1alpha1"
	"github.com/kubeflow/arena/pkg/operators/fluid-operator/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

// DatasetsGetter has a method to return a DatasetInterface.
// A group's client should implement this interface.
type DatasetsGetter interface {
	Datasets(namespace string) DatasetInterface
}

// DatasetInterface has methods to work wth Dataset resouces.
type DatasetInterface interface {
	Create(*v1alpha1.Dataset) (*v1alpha1.Dataset, error)
	Update(*v1alpha1.Dataset) (*v1alpha1.Dataset, error)
	Delete(name string, options *v1.DeleteOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Dataset, error)
	List(opts v1.ListOptions) (*v1alpha1.DatasetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	DataExpansion
}

// datasets implement DatasetInterface
type datasets struct {
	client rest.Interface
	ns     string
}

// newDataset returns a Datasets
func newDatasets(c *V1alpha1Client, namespace string) *datasets {
	return &datasets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name if the dataset, and returns the corresponding dataset object, and an errir if there is any.
func (c *datasets) Get(name string, options v1.GetOptions) (result *v1alpha1.Dataset, err error) {
	result = &v1alpha1.Dataset{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datasets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Datasets that match those selectors.
func (c *datasets) List(opts v1.ListOptions) (result *v1alpha1.DatasetList, err error) {
	result = &v1alpha1.DatasetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datasets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)

	return
}

// Watch returns a watch.Interface that watches the requested datasets.
func (c *datasets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("datasets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a dataset and creates it.  Returns the server's representation of the dataset, and an error, if there is any.
func (c *datasets) Create(dataset *v1alpha1.Dataset) (result *v1alpha1.Dataset, err error) {
	result = &v1alpha1.Dataset{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("datasets").
		Body(dataset).
		Do().
		Into(result)
	return
}

// Delete takes name of the dataset and deletes it. Returns an error if one occurs.
func (c *datasets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datasets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// Update takes the representation of a dataset and updates it. Returns the server's representation of the dataset, and an error, if there is any.
func (c *datasets) Update(dataset *v1alpha1.Dataset) (result *v1alpha1.Dataset, err error) {
	result = &v1alpha1.Dataset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datasets").
		Name(dataset.Name).
		Body(dataset).
		Do().
		Into(result)
	return
}
