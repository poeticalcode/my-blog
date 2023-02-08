package service_test

import (
	"log"
	"testing"

	"github.com/he-wen-yao/my-blog/server/model/entity"
	"github.com/he-wen-yao/my-blog/server/service"
)

func TestTagService_AddTag(t *testing.T) {
	tag := &entity.Tag{Name: "Java"}
	service.TagService.AddTag(tag)
	log.Printf("ID = %d", tag.ID)
}
