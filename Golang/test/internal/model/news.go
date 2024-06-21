package model

type News struct {
	Id      int64  `reform:"id,pk"`
	Title   string `reform:"title"`
	Content string `reform:"content"`
}

type NewsCategory struct {
	NewsId     int64 `reform:"news_id,pk"`
	CategoryId int64 `reform:"category_id,pk"`
}
