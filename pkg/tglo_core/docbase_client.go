package tglo_core

import (
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
}

func (me *DocbaseClient) Write(bs []byte) (n int, err error) {	
  	fmt.Printf("%s %s\n", me.Domain, me.AccessToken)
  	client := docbase.NewAuthClient(me.Domain, me.AccessToken)

	  var postingGroupIds []docbase.GroupID
	  From(me.PostingGroupIds).
	  SelectT(func(n int) docbase.GroupID {
		  return docbase.GroupID(n)
	  }).
	  ToSlice(&postingGroupIds)
	  fmt.Printf("%v\n", postingGroupIds)

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
	fmt.Println(jsonify(post))
	//fmt.Println(jsonify(res))
	fmt.Println(res.Response.StatusCode)

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