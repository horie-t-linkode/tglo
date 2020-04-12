package subcommand

import (
  "errors"
  "fmt"
  "os"
  "strconv"
  "github.com/spf13/cobra"
  "tglo_cli/mylogger"
  "github.com/joho/godotenv"
  "github.com/masaki-linkode/tglo/pkg"
)

var logger = mylogger.GetLogger()
var verbose bool

func readConfig() (me *pkg.TogglClient, err error) {
  _ = godotenv.Load()
  apiToken := os.Getenv("TGLO_APITOKEN")
  if apiToken == "" {
    return nil, errors.New(fmt.Sprintf("TGLO_APITOKEN is empty"))
  }
  workSpaceIdS := os.Getenv("TGLO_WORKSPACEID")
  if workSpaceIdS == "" {
    return nil, errors.New(fmt.Sprintf("TGLO_WORKSPACEID is empty"))
  }
  workspaceId, err := strconv.Atoi(workSpaceIdS)
  if err != nil { 
    return nil, errors.New(fmt.Sprintf("TGLO_WORKSPACEID: %s", err.Error()))
  }

  return &pkg.TogglClient{ApiToken: apiToken, WorkSpaceId: workspaceId, Verbose: verbose}, nil
}


func NewRootCommand() *cobra.Command {
  me := &cobra.Command{
    Use: "tglo_cli",
    Long:  `hogehoge`,
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

func Execute() {
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