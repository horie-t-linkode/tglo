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

var postDocbase_ bool
var supressDetail_ bool

func newWeekCommand() *cobra.Command {
	me := &cobra.Command{
		Use: "week",
		Short: "指定日を含む週分のtogglエントリのサマリを出力する",
		Long:  `指定日を含む週分のtogglエントリのサマリを出力する`,
		RunE: weekCommand,
		SilenceUsage: true,
		SilenceErrors: true,
	}
	me.Flags().StringP("date", "d", "", "指定日。yyyy-mm-dd形式。")
	me.MarkFlagRequired("date")
	me.Flags().BoolVarP(&supressDetail_, "supressDetail", "s", false, "詳細出力を抑制")
	me.Flags().BoolVarP(&postDocbase_, "postDocbase", "", false, "docbaseにポスト")
	return me
}

func weekCommand(cmd *cobra.Command, args []string) (err error) {
	dateS, _ := cmd.Flags().GetString("date")

	date, err := time_util.Date(dateS)
	if err != nil { return err }
	from := time_util.StartDayOfWeek(date)
	till := time_util.After24Hours(from, 7)

	return processWeek(from, till, postDocbase_, !supressDetail_)
}

func processWeek(from time.Time, till time.Time, postDocbase bool, showDetail bool) (err error) {

	tglCl, err := readTogglClientConfig(verboseOut_)
	if err != nil { return err }

	content, err := tglCl.Process(from, till, showDetail)
	if err != nil { return err }

	var buffer bytes.Buffer
	err = template.TemplateExecute(template.WeekTemplate(), &buffer, content)
	if err != nil { return err }


	writers := []io.Writer{commandOut_}
	if postDocbase {
		docbaseCl, err := readDocbaseClientConfig(verboseOut_)
		docbaseCl.PostingTitle = fmt.Sprintf("%s [%s 〜 %s]", docbaseCl.PostingTitle, content.From, content.Till)
		if err != nil { return err }

		writers = append(writers, docbaseCl)
	}
	mw := io.MultiWriter(writers...)
	_, err = buffer.WriteTo(mw)
	return err
}