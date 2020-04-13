package subcommand

import (
  "os"
  "github.com/spf13/cobra"
  "tglo_cli/mylogger"
)

var logger = mylogger.GetLogger()
var verbose bool

func NewRootCommand() *cobra.Command {
  me := &cobra.Command{
    Use: "tglo",
    Long:  `togglエントリ/サマリを出力する`,
    RunE: func(cmd *cobra.Command, args []string) error {
      return cmd.Help()
    },
    SilenceUsage: true,
		SilenceErrors: true,
  }

  me.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "開発者用デバッグ出力")

  cobra.OnInitialize(func() {
    if verbose {
      mylogger.SetLevelDebug()
    }
  })

  me.AddCommand(newVersionCommand())
  me.AddCommand(newDayCommand())
  me.AddCommand(newTodayCommand())
  me.AddCommand(newYesterdayCommand())
  me.AddCommand(newThisWeekCommand())
  me.AddCommand(newLastWeekCommand())
  
  return me
} 

var myversion string
var myrevision string
var mybuildDate string

func Execute(version string, revision string, buildDate string) {
  myversion = version
  myrevision = revision
  mybuildDate = buildDate
  rootCmd := NewRootCommand()
  rootCmd.SetOutput(os.Stdout)
  Exit(rootCmd.Execute())
}

func Exit(err error) {
  code := 0
  if err != nil {
    logger.Error(err)
    code = 1
  }
  os.Exit(code)
}