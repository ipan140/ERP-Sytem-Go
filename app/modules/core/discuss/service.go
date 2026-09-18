package discuss

func CreateChannelService(data *Channel) error {
	return CreateChannel(data)
}

func GetAllChannelService() ([]Channel, error) {
	return GetAllChannel()
}

func GetChannelByIDService(id uint) (*Channel, error) {
	return GetChannelByID(id)
}

func UpdateChannelService(data *Channel) error {
	return UpdateChannel(data)
}

func DeleteChannelService(id uint) error {
	return DeleteChannel(id)
}

func GetMessagesByChannelService(channelID string) ([]DiscussMessage, error) {
	return GetMessagesByChannel(channelID)
}

func CreateDiscussMessageService(data *DiscussMessage) error {
	return CreateDiscussMessage(data)
}
