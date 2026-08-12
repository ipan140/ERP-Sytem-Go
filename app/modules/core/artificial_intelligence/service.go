package artificial_intelligence

func CreateAIPromptService(data *AIPrompt) error {
	return CreateAIPrompt(data)
}

func GetAllAIPromptService() ([]AIPrompt, error) {
	return GetAllAIPrompt()
}

func GetAIPromptByIDService(id uint) (*AIPrompt, error) {
	return GetAIPromptByID(id)
}

func UpdateAIPromptService(data *AIPrompt) error {
	return UpdateAIPrompt(data)
}

func DeleteAIPromptService(id uint) error {
	return DeleteAIPrompt(id)
}
