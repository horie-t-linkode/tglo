package subcommand

import (
	"fmt"
	"time"
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
	me.Flags().StringP("date", "d", "", "yyyymmdd")
	me.MarkFlagRequired("date")
	return me
}

func DayCommand(cmd *cobra.Command, args []string) (err error) {
	dateS, _ := cmd.Flags().GetString("date")

	date, err := time.Parse("2006-01-02 MST", fmt.Sprintf("%s JST", dateS))
	if err != nil { return err }

	config, err := readConfig()
	if err != nil { return err }

	return pkg.Process(config.ApiToken, config.WorkSpaceId, date, cmd.OutOrStdout())
}