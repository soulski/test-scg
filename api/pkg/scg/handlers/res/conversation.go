package res

import (
	"time"

	"github.com/soulski/test-scg/pkg/scg/model"
)

type Sender struct {
}

type Message struct {
	Sender     string
	Text       string
	CreateDate time.Time
}

func NewMessage(msgModel *model.Message) *Message {
	return &Message{
		Sender:     string(msgModel.Sender),
		Text:       msgModel.Text,
		CreateDate: msgModel.CreateDate,
	}
}

type Conversation struct {
	Id       string
	User     *User
	Messages []*Message
}

func NewConversation(converModel *model.Conversation) *Conversation {
	messsages := make([]*Message, len(converModel.GetMessages()))

	for index, msgModel := range converModel.GetMessages() {
		messsages[index] = NewMessage(msgModel)
	}

	return &Conversation{
		Id:       converModel.GetId().String(),
		User:     NewUser(converModel.GetUser()),
		Messages: messsages,
	}
}
