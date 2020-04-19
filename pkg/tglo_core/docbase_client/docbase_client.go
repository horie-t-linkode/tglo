package docbase_client

import (
	"io"
	"fmt"
	"context"
	"encoding/json"
	"github.com/kyoh86/go-docbase/v2/docbase"
	. "github.com/ahmetb/go-linq"
)

type DocbaseClient struct {
	AccessToken string
	Domain string
	PostingTitle string
	PostingTags []string
	PostingGroupIds []int
	VerboseOut io.Writer
}

func (me *DocbaseClient) Write(bs []byte) (n int, err error) {	
	me.verbose(fmt.Sprintf("%s %s\n", me.Domain, me.AccessToken))
  	client := docbase.NewAuthClient(me.Domain, me.AccessToken)

	  var postingGroupIds []docbase.GroupID
	  From(me.PostingGroupIds).
	  SelectT(func(n int) docbase.GroupID {
		  return docbase.GroupID(n)
	  }).
	  ToSlice(&postingGroupIds)
	  me.verbose(fmt.Sprintf("%v\n", postingGroupIds))

	body := string(bs)
	post, res, err := client.
		Post.
		Create(me.PostingTitle, body).
		Scope(docbase.ScopeGroup).
		Notice(false).
		Tags(me.PostingTags).
		Groups(postingGroupIds).
		Do(context.Background())
	if err != nil {
		return 0, err
	}
	me.verbose(fmt.Sprintln(jsonify(post)))
	//fmt.Println(jsonify(res))
	me.verbose(fmt.Sprintln(res.Response.StatusCode))

	//samplePost = *post

	return len(body), nil
}

func jsonify(o interface{}) string {
	buf, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(buf)
}

func (me *DocbaseClient) verbose(s string) {
	if me.VerboseOut != nil {
		me.VerboseOut.Write([]byte(s))
	}
}