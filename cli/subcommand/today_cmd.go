package subcommand

import (
	"github.com/spf13/cobra"
	"github.com/masaki-linkode/tglo/pkg"
)

func newTodayCommand() *cobra.Command {
	me := &cobra.Command{
		Use: "today",
		Short: "hogehoge",
		Long:  `hogehahahahah`,
		RunE: TodayCommand,
		SilenceUsage: true,
		SilenceErrors: true,
	}
	return me
}

func TodayCommand(cmd *cobra.Command, args []string) (err error) {
	config, err := readConfig()
	if err != nil { return err }

	from := pkg.Today()
	till := pkg.NextDay(from)

	return pkg.Process(config.ApiToken, config.WorkSpaceId, from, till, cmd.OutOrStdout())
}