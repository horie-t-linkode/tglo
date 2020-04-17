package subcommand

import (
	"github.com/spf13/cobra"
	"github.com/masaki-linkode/tglo/pkg/tglo_core"
)

func newYesterdayCommand() *cobra.Command {
	me := &cobra.Command{
		Use: "yesterday",
		Short: "昨日のtogglエントリを出力する",
		Long:  `昨日のtogglエントリを出力する`,
		RunE: yesterdayCommand,
		SilenceUsage: true,
		SilenceErrors: true,
	}
	return me
}

func yesterdayCommand(cmd *cobra.Command, args []string) (err error) {

	from := tglo_core.Yesterday()
	till := tglo_core.After24Hours(from, 1)

	return processDay(from, till)
}