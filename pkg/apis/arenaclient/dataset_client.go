package arenaclient

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/kubeflow/arena/pkg/apis/config"
	"github.com/kubeflow/arena/pkg/operators/fluid-operator/apis/v1alpha1"
	"github.com/kubeflow/arena/pkg/operators/fluid-operator/client/clientset/versioned"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var errNotFoundOperator = errors.New("the server could not find the requested resource")

type DatasetClient struct {
	namespace     string
	configer      *config.ArenaConfiger
	datasetClient *versioned.Clientset
}

func NewDatasetClient(namespace string, configer *config.ArenaConfiger) *DatasetClient {
	log.Debugf("Init Dataset Client")
	// this step is used to check operator is installed or not
	datasetClient := versioned.NewForConfigOrDie(config.GetArenaConfiger().GetRestConfig())
	_, err := datasetClient.V1alpha1().Datasets(namespace).Get("test", metav1.GetOptions{})
	if err != nil && strings.Contains(err.Error(), errNotFoundOperator.Error()) {
		log.Debugf("not found fluid dataset operator")
	}
	return &DatasetClient{
		namespace:     namespace,
		configer:      configer,
		datasetClient: datasetClient,
	}
}

func (d *DatasetClient) getDataset(name, namespace string) (*v1alpha1.Dataset, error) {
	dataset, err := d.datasetClient.V1alpha1().Datasets(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(`datasets.data.fluid.io "%v" not found`, name)) {
			errMsg := fmt.Sprint("%s dataset in %s namespace %s not found", name, namespace)
			return nil, errors.New(errMsg)
		}
		return nil, err
	}
	return dataset, nil
}

func (d *DatasetClient) listDatasetByNamespace(namespace string) (*v1alpha1.DatasetList, error) {
	datasets, err := d.datasetClient.V1alpha1().Datasets(namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return datasets, nil
}

func (d *DatasetClient) ListAndPrintCache(namespace string, allNamespaces bool) error {
	if allNamespaces {
		namespace = metav1.NamespaceAll
	}
	datasets, err := d.listDatasetByNamespace(namespace)
	if err != nil {
		return errors.New("dataset not found")
	}
	displayDataset(datasets, allNamespaces)
	return nil
}

func displayDataset(datasets *v1alpha1.DatasetList, allNamespaces bool) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	if allNamespaces {
		fmt.Fprintf(w, "NAME\tUFS TOTAL SIZE\tCACHED\tCACHE CAPACITY\tCACHED PERCENTAGE\tPHASE\tOWNER\tAGE\n")
	} else {
		fmt.Fprintf(w, "NAME\tUFS TOTAL SIZE\tCACHED\tCACHE CAPACITY\tCACHED PERCENTAGE\tPHASE\tAGE\n")
	}
	for _, item := range datasets.Items {
		if allNamespaces {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
				item.DatasetName(),
				item.UFSTotalSize(),
				item.CacheSize(),
				item.CacheCapacity(),
				item.CachePercentage(),
				item.Phase(),
				item.Namespace,
				item.Age())
		} else {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
				item.DatasetName(),
				item.UFSTotalSize(),
				item.CacheSize(),
				item.CacheCapacity(),
				item.CachePercentage(),
				item.Phase(),
				item.Age())
		}
	}
	_ = w.Flush()
}
