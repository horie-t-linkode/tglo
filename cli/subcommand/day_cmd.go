package subcommand

import (
	"github.com/spf13/cobra"
	"github.com/masaki-linkode/tglo/pkg/tglo_core/template"
	"github.com/masaki-linkode/tglo/pkg/tglo_core/time_util"
	"time"
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

	from, err := time_util.Date(dateS)
	if err != nil { return err }

	till := time_util.After24Hours(from, 1)

	return processDay(from, till)
}

func processDay(from time.Time, till time.Time) (err error) {
	tglCl, err := readTogglClientConfig()
	if err != nil { return err }

	content, err := tglCl.Process(from, till, true)
	if err != nil { return err }

	return template.TemplateExecute(template.DayTemplate(), commandOut_, content)
}