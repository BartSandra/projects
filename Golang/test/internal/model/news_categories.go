package model

import (
	"database/sql"
	"fmt"
)

func UpdateNewsCategories(db *sql.DB, newsID int64, categoryIDs []int64) error {

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %v", err)
	}

	if _, err := tx.Exec("DELETE FROM NewsCategories WHERE NewsID = ?", newsID); err != nil {
		tx.Rollback()
		return fmt.Errorf("delete existing categories: %v", err)
	}

	for _, catID := range categoryIDs {
		_, err := tx.Exec("INSERT INTO NewsCategories (NewsID, CategoryID) VALUES (?, ?)", newsID, catID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("insert new category: %v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit transaction: %v", err)
	}

	return nil
}

func GetNewsCategories(db *sql.DB, newsID int64) ([]int64, error) {
	rows, err := db.Query("SELECT CategoryID FROM NewsCategories WHERE NewsID = ?", newsID)
	if err != nil {
		return nil, fmt.Errorf("query categories: %v", err)
	}
	defer rows.Close()

	var categories []int64
	for rows.Next() {
		var catID int64
		if err := rows.Scan(&catID); err != nil {
			return nil, fmt.Errorf("scan category: %v", err)
		}
		categories = append(categories, catID)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows final error: %v", err)
	}

	return categories, nil
}
