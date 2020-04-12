package subcommand

import (
	"github.com/spf13/cobra"
)

func newDayCommand() *cobra.Command {
	me := &cobra.Command{
		Use: "day",
		Short: "hogehoge",
		Long:  `hogehahahahah`,
		RunE: dayCommand,
		SilenceUsage: true,
		SilenceErrors: true,
	}
	me.Flags().StringP("date", "d", "", "yyyy-mm-dd")
	me.MarkFlagRequired("date")
	return me
}

func dayCommand(cmd *cobra.Command, args []string) (err error) {
	dateS, _ := cmd.Flags().GetString("date")

	tglCl, err := readConfig()
	if err != nil { return err }

	from, err := tglCl.Date(dateS)
	if err != nil { return err }

	till := tglCl.After24Hours(from, 1)

	return tglCl.ProcessDay(from, till, cmd.OutOrStdout())
}