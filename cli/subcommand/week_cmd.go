package subcommand

import (
	//"github.com/spf13/cobra"
	//"github.com/masaki-linkode/tglo/pkg/tglo_core"
	"github.com/masaki-linkode/tglo/pkg/tglo_core/template"
	"time"
	"io"
	"bytes"
	"fmt"
)

var postDocbase_ bool
var supressDetail_ bool

func processWeek(from time.Time, till time.Time, postDocbase bool, showDetail bool) (err error) {

	tglCl, err := readTogglClientConfig()
	if err != nil { return err }

	content, err := tglCl.Process(from, till, showDetail)
	if err != nil { return err }

	var buffer bytes.Buffer
	err = template.TemplateExecute(template.WeekTemplate(), &buffer, content)
	if err != nil { return err }


	writers := []io.Writer{commandOut_}
	if postDocbase {
		docbaseCl, err := readDocbaseClientConfig()
		docbaseCl.PostingTitle = fmt.Sprintf("%s [%s 〜 %s]", docbaseCl.PostingTitle, content.From, content.Till)
		if err != nil { return err }

		writers = append(writers, docbaseCl)
	}
	mw := io.MultiWriter(writers...)
	_, err = buffer.WriteTo(mw)
	return err
}