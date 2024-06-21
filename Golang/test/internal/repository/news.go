package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"test/internal/model"
	"test/pkg/db"
)

type NewsInput struct {
	Title      *string `json:"Title"`
	Content    *string `json:"Content"`
	Categories []int64 `json:"Categories"`
}

func UpdateNews(newsID string, input *NewsInput) error {
	database := db.DB
	news := &model.News{}

	tx, err := database.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %v", err)
	}
	defer tx.Rollback()

	err = tx.QueryRow("SELECT Id, Title, Content FROM News WHERE Id = ?", newsID).Scan(&news.Id, &news.Title, &news.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("news not found")
		}
		return fmt.Errorf("query news: %v", err)
	}

	if input.Title != nil {
		news.Title = *input.Title
	}
	if input.Content != nil {
		news.Content = *input.Content
	}
	_, err = tx.Exec("UPDATE News SET Title = ?, Content = ? WHERE Id = ?", news.Title, news.Content, news.Id)
	if err != nil {
		return fmt.Errorf("update news: %v", err)
	}

	if input.Categories != nil {
		_, err = tx.Exec("DELETE FROM NewsCategories WHERE NewsId = ?", news.Id)
		if err != nil {
			return fmt.Errorf("delete categories: %v", err)
		}

		for _, catID := range input.Categories {
			_, err = tx.Exec("INSERT INTO NewsCategories (NewsId, CategoryId) VALUES (?, ?)", news.Id, catID)
			if err != nil {
				return fmt.Errorf("insert category: %v", err)
			}
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("commit transaction: %v", err)
	}

	return nil
}

func GetNewsList(page int, limit int) ([]model.News, error) {
	database := db.DB
	offset := (page - 1) * limit

	rows, err := database.Query("SELECT Id, Title, Content FROM News LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, fmt.Errorf("query news: %v", err)
	}
	defer rows.Close()

	var newsList []model.News
	for rows.Next() {
		var news model.News
		err := rows.Scan(&news.Id, &news.Title, &news.Content)
		if err != nil {
			return nil, fmt.Errorf("scan news: %v", err)
		}
		newsList = append(newsList, news)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows final error: %v", err)
	}

	return newsList, nil
}
