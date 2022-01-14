package dao

type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

func (Article) TableName() string {
	return "go_article"
}
