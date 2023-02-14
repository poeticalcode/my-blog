package service

import (
	"github.com/he-wen-yao/my-blog/server/db"
	"github.com/he-wen-yao/my-blog/server/model/entity"
	"github.com/he-wen-yao/my-blog/server/model/req"
)

// 标签服务
type tagService struct{}

// FetchTagList 获取标签列表
func (tagService) FetchTagList(param *req.PagingParam) {

}

// AddTag 添加标签
func (tagService) AddTag(tag *entity.Tag) bool {
	return db.DB.Create(tag).RowsAffected != 0
}

// DeleteTagById 删除标签
func (tagService) DeleteTagById(id int64) (bool, error) {
	// 根据主键删除
	res := db.DB.Delete(&entity.Article{}, id)
	if res.Error != nil {
		return false, res.Error
	}
	return res.RowsAffected != 0, nil
}

// UpdateTagById 更新标签
func (tagService) UpdateTagById() {

}
