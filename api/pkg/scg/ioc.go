package scg

import (
	"fmt"

	"github.com/soulski/test-scg/pkg/scg/adapter"
	"github.com/soulski/test-scg/pkg/scg/config"
	handler "github.com/soulski/test-scg/pkg/scg/handlers"
	"github.com/soulski/test-scg/pkg/scg/repository"
)

type Container struct {
	instances map[string]interface{}
}

func DefaultContainer(config *config.Config) *Container {
	instances := make(map[string]interface{})

	//	googleAPIKey := "AIzaSyBVuTyYukncj6QYilJauk0DdmVe6Un-TUQ"
	placeApdater := adapter.NewPlaceAdapter(config.Place)

	//	lineChannelSecret := "81c8f4347a9ac9361de1f13d86a1c471"
	//	lineChannelAccessToken := "R/HukWph4QFWrSEXuJlYfVaRmukUwq1dkK8VYDZ+tbTP5WrRDckTDZB05Ihaezs0uDjfkvzpjq4oguIgHD0Hk+l6oi7+w4/f4Cg2tGsZcJnlUcV/gBV787HrDBVj1q68tOZ2t+T/c0mRWxxXRj/U4QdB04t89/1O/w1cDnyilFU="
	messageAdapter := adapter.NewMessageAdapter(config.Message)

	userRepository := repository.NewUserRepository()
	conversationRepository := repository.NewConversationRepository()

	scgHandler := handler.NewSCGHandler(placeApdater, userRepository, conversationRepository)
	messageHandler := handler.NewMessageHandler(messageAdapter, userRepository, conversationRepository)

	instances["repository.user"] = userRepository
	instances["repository.conversation"] = conversationRepository
	instances["adapter.place"] = placeApdater
	instances["adapter.message"] = messageAdapter
	instances["handler.scg"] = scgHandler
	instances["handler.message"] = messageHandler

	return &Container{
		instances,
	}
}

func (c *Container) GetInstance(name string) (interface{}, error) {
	instance, ok := c.instances[name]
	if !ok {
		return nil, fmt.Errorf("Instance name '%s' doens't exists", name)
	}

	return instance, nil
}
