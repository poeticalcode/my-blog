package do

type Article struct {
	baseModel
	Title  string `json:"title"`
	Cover  string `json:"cover"`
	MdText string `json:"md_text"`
}

func (Article) TableName() string {
	return "tb_article"
}
