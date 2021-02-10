package dataset

import (
	"github.com/spf13/cobra"
)

var (
	datasetLong = `manage cache for cache.
Avaliable Commands:
	submit				Submit a cache dataset for training.
	list, ls			List the cache dataset engine.
	`
)

func NewDatasetCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "dataset",
		Short: "manage dataset cache.",
		Long:  datasetLong,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	command.AddCommand(NewDatasetsListCommand())
	return command
}
