package subcommand

import (
	"github.com/spf13/cobra"
)

func newYesterdayCommand() *cobra.Command {
	me := &cobra.Command{
		Use: "yesterday",
		Short: "hogehoge",
		Long:  `hogehahahahah`,
		RunE: yesterdayCommand,
		SilenceUsage: true,
		SilenceErrors: true,
	}
	return me
}

func yesterdayCommand(cmd *cobra.Command, args []string) (err error) {
	tglCl, err := readConfig()
	if err != nil { return err }

	from := tglCl.Yesterday()
	till := tglCl.After24Hours(from, 1)

	return tglCl.Process(from, till, cmd.OutOrStdout())
}