package subcommand

import (
	"github.com/spf13/cobra"
)

func newLastWeekCommand() *cobra.Command {
	me := &cobra.Command{
		Use: "lastweek",
		Short: "hogehoge",
		Long:  `hogehahahahah`,
		RunE: lastWeekCommand,
		SilenceUsage: true,
		SilenceErrors: true,
	}
	return me
}

func lastWeekCommand(cmd *cobra.Command, args []string) (err error) {
	tglCl, err := readConfig()
	if err != nil { return err }

	from := tglCl.StartDayOfLastWeek()
	till := tglCl.After24Hours(from, 7)

	return tglCl.Process(from, till, cmd.OutOrStdout())
}