package handlers

import (
	"banner-service/pkg/models"
	"database/sql"
	"encoding/json"
	"github.com/lib/pq"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// GetBanners обрабатывает запросы на получение списка баннеров
func GetBanners(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var banners []models.Banner
		query := `SELECT id, feature_id, tag_id, content, active FROM banners`
		rows, err := db.Query(query)
		if err != nil {
			http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var b models.Banner
			if err := rows.Scan(&b.ID, &b.FeatureID, pq.Array(&b.TagIDs), &b.Content, &b.Active); err != nil {
				http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
				return
			}
			banners = append(banners, b)
		}

		json.NewEncoder(w).Encode(banners)
	}
}

// GetUserBanner обрабатывает запросы на получение баннера для пользователя
func GetUserBanner(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		featureID, err := strconv.Atoi(r.URL.Query().Get("feature_id"))
		if err != nil {
			http.Error(w, "Invalid feature_id parameter", http.StatusBadRequest)
			return
		}

		tagID, err := strconv.Atoi(r.URL.Query().Get("tag_id"))
		if err != nil {
			http.Error(w, "Invalid tag_id parameter", http.StatusBadRequest)
			return
		}

		useLastRevision := r.URL.Query().Get("use_last_revision") == "true"

		var banners []models.Banner
		var rows *sql.Rows

		if useLastRevision {
			// Запрос на получение последнего активного баннера для данной фичи и тега
			rows, err = db.Query("SELECT id, feature_id, tag_id, content, active FROM banners WHERE feature_id = $1 AND $2 = ANY(tag_id) AND active = true ORDER BY updated_at DESC LIMIT 1", featureID, tagID)
		} else {
			// Показываем данные, которые могут быть старше на 5 минут
			fiveMinutesAgo := time.Now().Add(-5 * time.Minute)
			rows, err = db.Query("SELECT id, feature_id, tag_id, content, active FROM banners WHERE feature_id = $1 AND $2 = ANY(tag_id) AND active = true AND updated_at <= $3 ORDER BY updated_at DESC LIMIT 1", featureID, tagID, fiveMinutesAgo)
		}

		if err != nil {
			http.Error(w, "Error querying the database: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var b models.Banner
			err := rows.Scan(&b.ID, &b.FeatureID, pq.Array(&b.TagIDs), &b.Content, &b.Active)
			if err != nil {
				http.Error(w, "Database scan error: "+err.Error(), http.StatusInternalServerError)
				return
			}
			banners = append(banners, b)
		}

		json.NewEncoder(w).Encode(banners)
	}
}

func CreateBanner(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var banner models.Banner
		if err := json.NewDecoder(r.Body).Decode(&banner); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		query := `INSERT INTO banners (feature_id, tag_id, content, active) VALUES ($1, $2, $3, $4)`
		_, err := db.Exec(query, banner.FeatureID, pq.Array(banner.TagIDs), banner.Content, banner.Active)
		if err != nil {
			http.Error(w, "Failed to create banner", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(banner)
	}
}

// UpdateBanner обрабатывает обновление существующего баннера
func UpdateBanner(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		bannerID, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid banner ID", http.StatusBadRequest)
			return
		}

		var banner models.Banner
		if err := json.NewDecoder(r.Body).Decode(&banner); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		query := `UPDATE banners SET feature_id = $1, tag_id = $2, content = $3, active = $4 WHERE id = $5`
		_, err = db.Exec(query, banner.FeatureID, pq.Array(banner.TagIDs), banner.Content, banner.Active, bannerID)
		if err != nil {
			http.Error(w, "Failed to update banner", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(banner)
	}
}

// DeleteBanner обрабатывает удаление баннера
func DeleteBanner(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		bannerID, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid banner ID", http.StatusBadRequest)
			return
		}

		query := `DELETE FROM banners WHERE id = $1`
		_, err = db.Exec(query, bannerID)
		if err != nil {
			http.Error(w, "Failed to delete banner", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
