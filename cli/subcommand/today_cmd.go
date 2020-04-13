package subcommand

import (
	"github.com/spf13/cobra"
)

func newTodayCommand() *cobra.Command {
	me := &cobra.Command{
		Use: "today",
		Short: "本日のtogglエントリを出力する",
		Long:  `本日のtogglエントリを出力する`,
		RunE: todayCommand,
		SilenceUsage: true,
		SilenceErrors: true,
	}
	return me
}

func todayCommand(cmd *cobra.Command, args []string) (err error) {
	tglCl, err := readTogglClientConfig()
	if err != nil { return err }

	from := tglCl.Today()
	till := tglCl.After24Hours(from, 1)

	return tglCl.ProcessDay(from, till, cmd.OutOrStdout())
}