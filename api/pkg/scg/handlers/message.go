package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/soulski/test-scg/pkg/scg/adapter"
	"github.com/soulski/test-scg/pkg/scg/handlers/req"
	"github.com/soulski/test-scg/pkg/scg/handlers/res"
	"github.com/soulski/test-scg/pkg/scg/model"
	"github.com/soulski/test-scg/pkg/scg/repository"
)

type MessageHandler struct {
	messageAdapter   *adapter.MessageAdapter
	userRepository   repository.UserRepository
	converRepository repository.ConversationRepository
}

func NewMessageHandler(
	messageAdapter *adapter.MessageAdapter,
	userRepository repository.UserRepository,
	converRepository repository.ConversationRepository,
) *MessageHandler {
	return &MessageHandler{messageAdapter, userRepository, converRepository}
}

func (m *MessageHandler) EventMessage(event *linebot.Event) {
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		userIdStr := event.Source.UserID

		var (
			user   *model.User
			conver *model.Conversation
		)

		fmt.Printf("%+v \n", event)

		userId := model.NewDefinedId(userIdStr)
		if tmpModel, err := m.userRepository.FindById(userId); err != nil {
			user = model.NewUser(userIdStr)
			m.userRepository.Create(user)
		} else {
			user = tmpModel.(*model.User)
		}

		if tmpModel, err := m.converRepository.FindById(userId); err != nil {
			conver = model.NewConversation(user)
			conver.AddMessage(model.UserSender, message.Text, event.Timestamp, event.ReplyToken)
			m.converRepository.Create(conver)
		} else {
			conver = tmpModel.(*model.Conversation)
			conver.AddMessage(model.UserSender, message.Text, event.Timestamp, event.ReplyToken)
			m.converRepository.Update(conver)
		}

		fmt.Printf("%+v\n", message)
	default:
		fmt.Println("Unknow message.")
	}
}

func (m *MessageHandler) ListUser(c *gin.Context) {
	models := m.userRepository.FindAll()
	users := make([]*res.User, len(models))

	for index, mod := range models {
		users[index] = res.NewUser(mod.(*model.User))
	}

	c.JSON(http.StatusOK, users)
}

func (m *MessageHandler) ListConversation(c *gin.Context) {
	models := m.converRepository.FindAll()
	conversations := make([]*res.Conversation, len(models))

	for index, mod := range models {
		conversations[index] = res.NewConversation(mod.(*model.Conversation))
	}

	c.JSON(http.StatusOK, conversations)
}

func (m *MessageHandler) GetConversation(c *gin.Context) {
	conversation, err := m.converRepository.FindById(model.NewDefinedId(c.Param("cid")))
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, res.NewConversation(conversation.(*model.Conversation)))
}

func (m *MessageHandler) CreateConversationMessage(c *gin.Context) {
	mod, err := m.converRepository.FindById(model.NewDefinedId(c.Param("cid")))
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	var createConversationMsg req.CreateConversationMessage
	if err := c.ShouldBindJSON(&createConversationMsg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msgStr := createConversationMsg.Message
	conversation := mod.(*model.Conversation)
	//lastestMsg := conversation.GetLastestMessages()
	to := conversation.GetUser().GetId().String()

	m.messageAdapter.PushMessage(to, msgStr)

	conversation.AddMessage(model.SystemSender, msgStr, time.Now(), "")
	m.converRepository.Update(conversation)

	c.JSON(http.StatusOK, res.NewConversation(conversation))
}
