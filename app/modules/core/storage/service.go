package storage

func CreateAttachmentService(data *Attachment) error {
	return CreateAttachment(data)
}

func GetAllAttachmentService() ([]Attachment, error) {
	return GetAllAttachment()
}

func GetPaginatedAttachmentService(offset, limit int, search string) ([]Attachment, int64, error) {
	return GetPaginatedAttachments(offset, limit, search)
}

func GetAttachmentByIDService(id uint) (*Attachment, error) {
	return GetAttachmentByID(id)
}

func UpdateAttachmentService(data *Attachment) error {
	return UpdateAttachment(data)
}

func DeleteAttachmentService(id uint) error {
	return DeleteAttachment(id)
}
