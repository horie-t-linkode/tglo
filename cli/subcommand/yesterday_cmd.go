package subcommand

import (
	"github.com/spf13/cobra"
	"github.com/masaki-linkode/tglo/pkg/tglo_core/time_util"
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
	me.Flags().BoolVarP(&postSlack_, "postSlack", "", false, "slackにポスト")
	return me
}

func yesterdayCommand(cmd *cobra.Command, args []string) (err error) {

	from := time_util.Yesterday()
	till := time_util.After24Hours(from, 1)

	return processDay(from, till, postSlack_)
}