package dao

type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
	Image   string `json:"image"`
	Type    int    `json:"type"`
	Views   int    `json:"views"`
}

func (Article) TableName() string {
	return "go_article"
}

type ArticleValidator struct {
	UserId   int `binding:"required" validate:"required"`
	Page     int `binding:"required" validate:"required,gte=1,lte=100"`
	PageSize int `binding:"required" validate:"required,gte=1,lte=100"`
}

type HotValidator struct {
	Type     int `validate:"required"`
	Page     int `validate:"required,gte=1,lte=100"`
	PageSize int `validate:"required,gte=1,lte=100"`
}
