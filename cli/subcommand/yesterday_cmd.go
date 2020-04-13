package subcommand

import (
	"github.com/spf13/cobra"
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
	tglCl, err := readTogglClientConfig()
	if err != nil { return err }

	from := tglCl.Yesterday()
	till := tglCl.After24Hours(from, 1)

	return tglCl.ProcessDay(from, till, cmd.OutOrStdout())
}