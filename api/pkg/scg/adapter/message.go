package adapter

import (
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/soulski/test-scg/pkg/scg/config"
)

type MessageAdapter struct {
	channel *linebot.Client
}

func NewMessageAdapter(config *config.MessageConfig) *MessageAdapter {
	channel, _ := linebot.New(config.Secret, config.AccessToken)
	return &MessageAdapter{channel}
}

func (m *MessageAdapter) Handler() func(http.ResponseWriter, *http.Request) chan *linebot.Event {
	return func(w http.ResponseWriter, req *http.Request) chan *linebot.Event {
		events, err := m.channel.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
		} else {
			w.WriteHeader(200)
		}

		outCh := make(chan *linebot.Event)
		go func(outCh chan *linebot.Event) {
			for _, event := range events {
				if event.ReplyToken == "00000000000000000000000000000000" || event.ReplyToken == "ffffffffffffffffffffffffffffffff" {
					continue
				}

				outCh <- event
			}
			close(outCh)
		}(outCh)

		return outCh
	}
}

func (m *MessageAdapter) ReplyMessage(replyToken string, message string) (*linebot.BasicResponse, error) {
	return m.channel.ReplyMessage(replyToken, linebot.NewTextMessage(message)).Do()
}

func (m *MessageAdapter) PushMessage(to string, message string) (*linebot.BasicResponse, error) {
	return m.channel.PushMessage(to, linebot.NewTextMessage(message)).Do()
}
