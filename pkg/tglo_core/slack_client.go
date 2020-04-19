package tglo_core

import (
	"io"
	"fmt"
	"github.com/slack-go/slack"
)

type SlackClient struct {
	OAuthAccessToken string
	PostingChannelID string
	VerboseOut io.Writer
}

func (me *SlackClient) Write(bs []byte) (n int, err error) {
	me.verbose(fmt.Sprintf("%s %s\n", me.OAuthAccessToken, me.PostingChannelID))

	api := slack.New(me.OAuthAccessToken)
	message := string(bs)

	channelID, timestamp, err := api.
		PostMessage(me.PostingChannelID, slack.MsgOptionText(message, false))
	if err != nil {
		return 0, err
	}
	me.verbose(fmt.Sprintf("Message successfully sent to channel %s at %s", channelID, timestamp))

	return len(message), nil
}

func (me *SlackClient) verbose(s string) {
	if me.VerboseOut != nil {
		me.VerboseOut.Write([]byte(s))
	}
}