package dataset

import (
	"os"

	"github.com/kubeflow/arena/pkg/apis/arenaclient"
	"github.com/kubeflow/arena/pkg/apis/types"
	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewDatasetsListCommand() *cobra.Command {
	var allNamespaces bool
	var command = &cobra.Command{
		Use:   "list",
		Short: "list all cache engine",
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := arenaclient.NewArenaClient(types.ArenaClientArgs{
				Kubeconfig:     viper.GetString("config"),
				LogLevel:       viper.GetString("loglevel"),
				Namespace:      viper.GetString("namespace"),
				ArenaNamespace: viper.GetString("arena-namespace"),
				IsDaemonMode:   false,
			})
			if err != nil {
				log.Errorf("failed to build arena client,reason: %v", err)
				os.Exit(1)
			}

			return client.Dataset().ListAndPrintCache(viper.GetString("namespace"), allNamespaces)
		},
	}

	command.Flags().BoolVar(&allNamespaces, "allNamespaces", false, "show all the namespaces")

	return command
}
