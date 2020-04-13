package tglo_core

import (
	"fmt"
	"context"
	"encoding/json"
	"github.com/kyoh86/go-docbase/v2/docbase"
)

type DocbaseClient struct {
	AccessToken string
	Domain string
}

func (me *DocbaseClient) Write(bs []byte) (n int, err error) {	
  fmt.Printf("%s %s\n", me.Domain, me.AccessToken)
  client := docbase.NewAuthClient(me.Domain, me.AccessToken)

	body := string(bs)
	post, res, err := client.
		Post.
		Create("testTitle", body).
		Scope(docbase.ScopePrivate).
		Notice(false).
		Tags([]string{"go-docbase-test"}).
		Do(context.Background())
	if err != nil {
		return 0, err
	}
	fmt.Println(res.Response.StatusCode)
	fmt.Println(jsonify(post))
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