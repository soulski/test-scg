package req

type CreateConversationMessage struct {
	Message string `form:"message" json:"message" xml:"message"  binding:"required"`
}
