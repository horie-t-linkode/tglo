package subcommand

import (
	"github.com/spf13/cobra"
	"github.com/masaki-linkode/tglo/pkg/tglo_core/time_util"
	"github.com/masaki-linkode/tglo/pkg/tglo_core/template"
	"time"
	"io"
	"bytes"
	"fmt"
)

func newDurationCommand() *cobra.Command {
	me := &cobra.Command{
		Use: "duration",
		Short: "指定開始日から指定終了日までのtogglエントリのサマリを出力する",
		Long:  `指定開始日から指定終了日までのtogglエントリのサマリを出力する`,
		RunE: durationCommand,
		SilenceUsage: true,
		SilenceErrors: true,
	}
	me.Flags().StringP("start", "", "", "指定日開始日。yyyy-mm-dd形式。")
	me.MarkFlagRequired("start")
	me.Flags().StringP("end", "", "", "指定日終了日。yyyy-mm-dd形式。")
	me.MarkFlagRequired("end")

	me.Flags().BoolVarP(&supressDetail_, "supressDetail", "s", false, "詳細出力を抑制")
	me.Flags().BoolVarP(&postDocbase_, "postDocbase", "", false, "docbaseにポスト")
	me.Flags().BoolVarP(&postSlack_, "postSlack", "", false, "slackにポスト")
	return me
}

func durationCommand(cmd *cobra.Command, args []string) (err error) {
	startS, _ := cmd.Flags().GetString("start")
	endS, _ := cmd.Flags().GetString("end")

	start, err := time_util.Date(startS)
	if err != nil { return err }
	end, err := time_util.Date(endS)
	if err != nil { return err }

	from := start
	till := time_util.After24Hours(end, 1)

	return processDuration(from, till, postDocbase_, postSlack_, !supressDetail_)
}

func processDuration(from time.Time, till time.Time, postDocbase bool, postSlack bool, showDetail bool) (err error) {

	tglCl, err := readTogglClientConfig(verboseOut_)
	if err != nil { return err }

	content, err := tglCl.Process(from, till)
	if err != nil { return err }

	var buffer bytes.Buffer
	if showDetail {
		err = template.TemplateExecute(template.WeekTemplate(), &buffer, content)
	} else {
		err = template.TemplateExecute(template.WeekTemplateSupressDetail(), &buffer, content)
	}
	if err != nil { return err }

	writers := []io.Writer{commandOut_}
	if postDocbase {
		docbaseCl, err := readDocbaseClientConfig(verboseOut_)
		docbaseCl.PostingTitle = fmt.Sprintf("%s [%s 〜 %s]", docbaseCl.PostingTitle, content.From, content.Till)
		if err != nil { return err }

		writers = append(writers, docbaseCl)
	}
	if postSlack {
		slackCl, err := readSlackConfig(verboseOut_)
		if err != nil { return err }
	
		writers = append(writers, slackCl)
	}
	mw := io.MultiWriter(writers...)
	_, err = buffer.WriteTo(mw)
	return err
}