package v1alpha1

import (
	v1alpha1 "github.com/kubeflow/arena/pkg/operators/fluid-operator/apis/v1alpha1"
	"github.com/kubeflow/arena/pkg/operators/fluid-operator/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type V1alpha1Interface interface {
	RESTClient() rest.Interface
	DatasetsGetter
}

// V1alpha1Client is used to interact with features provided by the  group.
type V1alpha1Client struct {
	restClient rest.Interface
}

func (c *V1alpha1Client) Datasets(namespace string) DatasetInterface {
	return newDatasets(c, namespace)
}

// NewForConfig creates a new V1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*V1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &V1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new V1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *V1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new V1alpha1Client for the given RESTClient.
func New(c rest.Interface) *V1alpha1Client {
	return &V1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.GroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *V1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
