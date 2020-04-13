package subcommand

import (
	"github.com/spf13/cobra"
)

var supressDetail bool

func newThisWeekCommand() *cobra.Command {
	me := &cobra.Command{
		Use: "thisweek",
		Short: "今週分のtogglエントリのサマリを出力する",
		Long:  `今週分のtogglエントリのサマリを出力する`,
		RunE: thisWeekCommand,
		SilenceUsage: true,
		SilenceErrors: true,
	}
	me.Flags().BoolVarP(&supressDetail, "supressDetail", "s", false, "詳細出力を抑制")
	return me
}

func thisWeekCommand(cmd *cobra.Command, args []string) (err error) {
	tglCl, err := readTogglClientConfig()
	if err != nil { return err }

	from := tglCl.StartDayOfThisWeek()
	till := tglCl.After24Hours(from, 7)

	return tglCl.ProcessWeek(from, till, cmd.OutOrStdout(), !supressDetail)
}