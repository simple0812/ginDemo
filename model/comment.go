package model

type Comment struct {
	ID        int `json:"id"`
	UserId    int `json:"userId" sql:"index"`
	ArticleId int `json:"atricleId" sql:"index" gorm:"foreign_key:article(id)"`
	Content   int `json:"content"`
	Status    int `json:"status"`
	CreatedAt int `json:"createdAt"`
}
