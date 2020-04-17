package subcommand

import (
	"github.com/spf13/cobra"
	"github.com/masaki-linkode/tglo/pkg/tglo_core"
)



func newThisWeekCommand() *cobra.Command {
	me := &cobra.Command{
		Use: "thisweek",
		Short: "今週分のtogglエントリのサマリを出力する",
		Long:  `今週分のtogglエントリのサマリを出力する`,
		RunE: thisWeekCommand,
		SilenceUsage: true,
		SilenceErrors: true,
	}
	me.Flags().BoolVarP(&supressDetail_, "supressDetail", "s", false, "詳細出力を抑制")
	me.Flags().BoolVarP(&postDocbase_, "postDocbase", "", false, "docbaseにポスト")
	return me
}

func thisWeekCommand(cmd *cobra.Command, args []string) (err error) {

	from := tglo_core.StartDayOfThisWeek()
	till := tglo_core.After24Hours(from, 7)

	return processWeek(from, till, postDocbase_, !supressDetail_)
}