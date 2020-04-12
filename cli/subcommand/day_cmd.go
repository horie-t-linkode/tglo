package subcommand

import (
	"github.com/spf13/cobra"
)

func newDayCommand() *cobra.Command {
	me := &cobra.Command{
		Use: "day",
		Short: "指定日のtogglエントリを出力する",
		Long:  `指定日のtogglエントリを出力する`,
		RunE: dayCommand,
		SilenceUsage: true,
		SilenceErrors: true,
	}
	me.Flags().StringP("date", "d", "", "指定日。yyyy-mm-dd形式。")
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