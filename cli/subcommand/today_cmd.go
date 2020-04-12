package subcommand

import (
	"github.com/spf13/cobra"
)

func newTodayCommand() *cobra.Command {
	me := &cobra.Command{
		Use: "today",
		Short: "hogehoge",
		Long:  `hogehahahahah`,
		RunE: todayCommand,
		SilenceUsage: true,
		SilenceErrors: true,
	}
	return me
}

func todayCommand(cmd *cobra.Command, args []string) (err error) {
	tglCl, err := readConfig()
	if err != nil { return err }

	from := tglCl.Today()
	till := tglCl.After24Hours(from, 1)

	return tglCl.ProcessDay(from, till, cmd.OutOrStdout())
}