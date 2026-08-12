package live_chat

func CreateChatSessionService(data *ChatSession) error {
	return CreateChatSession(data)
}

func GetAllChatSessionService() ([]ChatSession, error) {
	return GetAllChatSession()
}

func GetChatSessionByIDService(id uint) (*ChatSession, error) {
	return GetChatSessionByID(id)
}

func UpdateChatSessionService(data *ChatSession) error {
	return UpdateChatSession(data)
}

func DeleteChatSessionService(id uint) error {
	return DeleteChatSession(id)
}
