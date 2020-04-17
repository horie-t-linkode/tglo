package subcommand

import (
	"github.com/spf13/cobra"
	"github.com/masaki-linkode/tglo/pkg/tglo_core/time_util"
)

func newLastWeekCommand() *cobra.Command {
	me := &cobra.Command{
		Use: "lastweek",
		Short: "先週分のtogglエントリのサマリを出力する",
		Long:  `先週分のtogglエントリのサマリを出力する`,
		RunE: lastWeekCommand,
		SilenceUsage: true,
		SilenceErrors: true,
	}
	me.Flags().BoolVarP(&supressDetail_, "supressDetail", "s", false, "詳細出力を抑制")
	me.Flags().BoolVarP(&postDocbase_, "postDocbase", "", false, "docbaseにポスト")
	return me
}

func lastWeekCommand(cmd *cobra.Command, args []string) (err error) {

	from := time_util.StartDayOfLastWeek()
	till := time_util.After24Hours(from, 7)

	return processWeek(from, till, postDocbase_, !supressDetail_)
}