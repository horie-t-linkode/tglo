package subcommand

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version = "unknown"
var revision = "unknown"
var buildDate = "unknown"

func newVersionCommand() *cobra.Command {
	me := &cobra.Command{
		Use: "version",
		Short: "バージョンを表示",
		Long:  `本コマンドのバージョンを表示する。`,
		RunE: VersionCommand,
		SilenceUsage: true,
		SilenceErrors: true,
	}
	return me
}

func VersionCommand(cmd *cobra.Command, args []string) error {
	cmd.Println(fmt.Sprintf("%s.%s %s", version, revision, buildDate))
	return nil
}