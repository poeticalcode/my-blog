package entity

type Tag struct {
	baseModel
	Name string `json:"name"` // 标签名称
}

func (Tag) TableName() string {
	return "tb_tag"
}
