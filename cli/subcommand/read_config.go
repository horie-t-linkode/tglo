package subcommand

import (
  "errors"
  "fmt"
  "os"
  "strconv"
  "strings"
  "github.com/joho/godotenv"
  "github.com/masaki-linkode/tglo/pkg/tglo_core"
)

func readTogglClientConfig() (me *tglo_core.TogglClient, err error) {
  _ = godotenv.Load()
  
  apiToken := os.Getenv("TGLO_TOGGL_APITOKEN")
  if apiToken == "" {
    return nil, errors.New(fmt.Sprintf("TGLO_TOGGL_APITOKEN is empty"))
  }
  
  workSpaceIdS := os.Getenv("TGLO_TOGGL_WORKSPACEID")
  if workSpaceIdS == "" {
    return nil, errors.New(fmt.Sprintf("TGLO_TOGGL_WORKSPACEID is empty"))
  }
  workspaceId, err := strconv.Atoi(workSpaceIdS)
  if err != nil { 
    return nil, errors.New(fmt.Sprintf("TGLO_TOGGL_WORKSPACEID: %s", err.Error()))
  }

  return &tglo_core.TogglClient{ApiToken: apiToken, WorkSpaceId: workspaceId, VerboseOut: verboseOut_}, nil
}

func readDocbaseClientConfig() (me *tglo_core.DocbaseClient, err error) {
  _ = godotenv.Load()

  domain := os.Getenv("TGLO_DOCBASE_DOMAIN")
  if domain == "" {
    return nil, errors.New(fmt.Sprintf("TGLO_DOCBASE_DOMAIN is empty"))
  }

  accessToken := os.Getenv("TGLO_DOCBASE_ACCESSTOKEN")
  if accessToken == "" {
    return nil, errors.New(fmt.Sprintf("TGLO_DOCBASE_ACCESSTOKEN is empty"))
  }

  postingTitle := os.Getenv("TGLO_DOCBASE_POSTING_TITLE")
  if postingTitle == "" {
    return nil, errors.New(fmt.Sprintf("TGLO_DOCBASE_POST_TAGS is empty"))
  }

  postingTags := os.Getenv("TGLO_DOCBASE_POSTING_TAGS")
  //if postingTags == "" {
  //  return nil, errors.New(fmt.Sprintf("TGLO_DOCBASE_POST_TAGS is empty"))
  //}

  postingGroups := os.Getenv("TGLO_DOCBASE_POSTING_GROUPS")
  if postingGroups == "" {
    return nil, errors.New(fmt.Sprintf("TGLO_DOCBASE_POST_GROUPS is empty"))
  }
  
  var postingGroupIds []int
  for _, s := range strings.Split(postingGroups, ","){
    var n, err = strconv.Atoi(s)
    if err != nil { return nil, errors.New(fmt.Sprintf("TGLO_DOCBASE_POST_GROUPS not convert")) }
    postingGroupIds = append(postingGroupIds, n)
  }

  return &tglo_core.DocbaseClient{
    AccessToken: accessToken, 
    Domain: domain, 
    PostingTitle: postingTitle,
    PostingTags: strings.Split(postingTags, ","),
    PostingGroupIds: postingGroupIds,
    }, nil
}