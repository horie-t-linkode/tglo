package subcommand

import (
	"github.com/spf13/cobra"
	"bytes"
	"io"
)

var postDocbase bool

func newLastWeekCommand() *cobra.Command {
	me := &cobra.Command{
		Use: "lastweek",
		Short: "先週分のtogglエントリのサマリを出力する",
		Long:  `先週分のtogglエントリのサマリを出力する`,
		RunE: lastWeekCommand,
		SilenceUsage: true,
		SilenceErrors: true,
	}
	me.Flags().BoolVarP(&supressDetail, "supressDetail", "s", false, "詳細出力を抑制")
	me.Flags().BoolVarP(&postDocbase, "postDocbase", "", false, "docbaseにポスト")
	return me
}

func lastWeekCommand(cmd *cobra.Command, args []string) (err error) {
	tglCl, err := readTogglClientConfig()
	if err != nil { return err }

	from := tglCl.StartDayOfLastWeek()
	till := tglCl.After24Hours(from, 7)


	var buffer bytes.Buffer
	err = tglCl.ProcessWeek(from, till, &buffer, !supressDetail)
	if err != nil { return err }

	writers := []io.Writer{cmd.OutOrStdout()}
	if postDocbase {
		docbaseCl, err := readDocbaseClientConfig()
		if err != nil { return err }

		writers = append(writers, docbaseCl)
	}
	mw := io.MultiWriter(writers...)
	_, err = buffer.WriteTo(mw)
	return err
}