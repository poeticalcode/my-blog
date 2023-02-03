package entity

type Article struct {
	baseModel
	Title       string `json:"title"`       // 文章标题
	Cover       string `json:"cover"`       // 文章封面
	MdText      string `json:"md_text"`     // markdowm 文本
	Description string `json:"description"` // 文章描述
}

func (Article) TableName() string {
	return "tb_article"
}
