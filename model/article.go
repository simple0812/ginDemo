package model

type Article struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Status    int    `json:"status"`
	UserId    int    `json:"userId" sql:"foreignkey:user(id)"`
	CreatedAt string `json:"createdAt"`

	Comments []Comment `json:"_"`
}
