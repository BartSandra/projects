package models

import "encoding/json"

// Banner описывает структуру баннера для управления через API
type Banner struct {
	ID        int             `json:"id"`
	FeatureID int             `json:"feature_id"`
	TagIDs    []int           `json:"tag_id"`
	Content   json.RawMessage `json:"content"`
	Active    bool            `json:"active"`
}
