package subcommand

import (
	"github.com/spf13/cobra"
	"github.com/masaki-linkode/tglo/pkg/tglo_core/template"
	"github.com/masaki-linkode/tglo/pkg/tglo_core/time_util"
	"time"
	"bytes"
	"io"
)

var postSlack_ bool

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
	me.Flags().BoolVarP(&postSlack_, "postSlack", "", false, "slackにポスト")
	return me
}

func dayCommand(cmd *cobra.Command, args []string) (err error) {
	dateS, _ := cmd.Flags().GetString("date")

	from, err := time_util.Date(dateS)
	if err != nil { return err }

	till := time_util.After24Hours(from, 1)

	return processDay(from, till, postSlack_)
}

func processDay(from time.Time, till time.Time, postSlack bool) (err error) {
	tglCl, err := readTogglClientConfig(verboseOut_)
	if err != nil { return err }

	content, err := tglCl.Process(from, till, true)
	if err != nil { return err }

	var buffer bytes.Buffer
	err = template.TemplateExecute(template.DayTemplate(), &buffer, content)
	if err != nil { return err }

	writers := []io.Writer{commandOut_}
	if postSlack {
		slackCl, err := readSlackConfig(verboseOut_)
		if err != nil { return err }
	
		writers = append(writers, slackCl)
	}
	mw := io.MultiWriter(writers...)
	_, err = buffer.WriteTo(mw)
	return err
}