package subcommand

import (
	"github.com/spf13/cobra"
	"github.com/masaki-linkode/tgl/pkg"
)

func newDayCommand() *cobra.Command {
	me := &cobra.Command{
		Use: "day",
		Short: "hogehoge",
		Long:  `hogehahahahah`,
		RunE: DayCommand,
		SilenceUsage: true,
		SilenceErrors: true,
	}
	me.Flags().StringP("date", "d", "", "yyyy-mm-dd")
	me.MarkFlagRequired("date")
	return me
}

func DayCommand(cmd *cobra.Command, args []string) (err error) {
	dateS, _ := cmd.Flags().GetString("date")

	config, err := readConfig()
	if err != nil { return err }

	return pkg.ProcessDay(config.ApiToken, config.WorkSpaceId, dateS, cmd.OutOrStdout())
}