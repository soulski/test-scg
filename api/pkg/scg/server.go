package scg

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/soulski/test-scg/pkg/scg/adapter"
	"github.com/soulski/test-scg/pkg/scg/config"
	handler "github.com/soulski/test-scg/pkg/scg/handlers"
	"github.com/spf13/viper"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	var config *config.Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	container := DefaultContainer(config)
	scgHandler := getInstance(container, "handler.scg").(*handler.SCGHandler)
	messageHandler := getInstance(container, "handler.message").(*handler.MessageHandler)
	messageAdapter := getInstance(container, "adapter.message").(*adapter.MessageAdapter)

	router.GET("/api/xyz", scgHandler.FindXYZ)
	router.GET("/api/restaurant", scgHandler.FindRestaurants)
	router.GET("/api/line/users", messageHandler.ListUser)
	router.GET("/api/line/conversations", messageHandler.ListConversation)
	router.GET("/api/line/conversations/:cid", messageHandler.GetConversation)
	router.POST("/api/line/conversations/:cid", messageHandler.CreateConversationMessage)

	router.POST("/hook/line", func(c *gin.Context) {
		handler := messageAdapter.Handler()
		eventCh := handler(c.Writer, c.Request)

		go func() {
			for event := range eventCh {
				switch event.Type {
				case linebot.EventTypeMessage:
					messageHandler.EventMessage(event)
				default:
				}
			}
		}()
	})

	return &Server{
		router: router,
	}
}

func getInstance(container *Container, name string) interface{} {
	instance, err := container.GetInstance(name)
	if err != nil {
		panic(err.Error())
	}

	return instance
}

func (s *Server) Start() {
	s.router.Run()
}
