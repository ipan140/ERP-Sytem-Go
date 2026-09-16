package blog

import (
	"log"
	
	"ERP-System/pkg/rabbitmq"
)

func CreateBlogPostService(data *BlogPost) error {
	return CreateBlogPost(data)
}

func GetAllBlogPostService() ([]BlogPost, error) {
	if rabbitmq.Channel != nil {
		_ = rabbitmq.PublishEvent(rabbitmq.Channel, "core_outbound_webhook", []byte(`{"event":"page_view","page":"blog_index"}`))
		log.Println("🌐 Event RabbitMQ: Kunjungan halaman Blog direkam (Activity Tracker)!")
	}
	return GetAllBlogPost()
}

func GetPaginatedBlogPostService(offset, limit int, search string) ([]BlogPost, int64, error) {
	return GetPaginatedBlogPosts(offset, limit, search)
}

func GetBlogPostByIDService(id uint) (*BlogPost, error) {
	return GetBlogPostByID(id)
}

func UpdateBlogPostService(data *BlogPost) error {
	return UpdateBlogPost(data)
}

func DeleteBlogPostService(id uint) error {
	return DeleteBlogPost(id)
}
