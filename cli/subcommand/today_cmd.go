package subcommand

import (
	"github.com/spf13/cobra"
	"github.com/masaki-linkode/tglo/pkg/tglo_core"
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
	from := tglo_core.Today()
	till := tglo_core.After24Hours(from, 1)

	return processDay(from, till)
}