package subcommand

import (
	"github.com/spf13/cobra"
)

func newThisWeekCommand() *cobra.Command {
	me := &cobra.Command{
		Use: "thisweek",
		Short: "hogehoge",
		Long:  `hogehahahahah`,
		RunE: thisWeekCommand,
		SilenceUsage: true,
		SilenceErrors: true,
	}
	return me
}

func thisWeekCommand(cmd *cobra.Command, args []string) (err error) {
	tglCl, err := readConfig()
	if err != nil { return err }

	from := tglCl.StartDayOfThisWeek()
	till := tglCl.After24Hours(from, 7)

	return tglCl.ProcessWeek(from, till, cmd.OutOrStdout())
}