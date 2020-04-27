package tglo_core

import (
	"io"
	"fmt"
	"bytes"
	"github.com/slack-go/slack"
)

type SlackClient struct {
	OAuthAccessToken string
	PostingChannelID string
	PostingTitle string
	VerboseOut io.Writer
}

func (me *SlackClient) Write(bs []byte) (n int, err error) {
	me.verbose(fmt.Sprintf("%s %s %s\n", me.OAuthAccessToken, me.PostingChannelID, me.PostingTitle))

	api := slack.New(me.OAuthAccessToken)

	s := string(bs)

	var buf bytes.Buffer
	buf.WriteString(me.PostingTitle)
	buf.WriteString(s)
	message := buf.String()

	params := slack.NewPostMessageParameters()
	//fmt.Printf("%#v", params)
	params.Parse = "full"
	channelID, timestamp, err := api.
		PostMessage(me.PostingChannelID, slack.MsgOptionText(message, false), slack.MsgOptionPostMessageParameters(params))
	if err != nil {
		return 0, err
	}
	me.verbose(fmt.Sprintf("Message successfully sent to channel %s at %s", channelID, timestamp))

	return len(s), nil
}

func (me *SlackClient) verbose(s string) {
	if me.VerboseOut != nil {
		me.VerboseOut.Write([]byte(s))
	}
}