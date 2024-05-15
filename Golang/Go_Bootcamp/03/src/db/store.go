package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type Place struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Location struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"location"`
}

// Store интерфейс определяет методы для работы с хранилищем данных
type Store interface {
	GetPlaces(limit int, offset int) ([]Place, int, error)
	GetClosestPlaces(lat float64, lon float64, limit int) ([]Place, error)
}

// ElasticsearchStore структура представляет хранилище данных Elasticsearch
type ElasticsearchStore struct {
	client *elasticsearch.Client
}

// NewElasticStore создает новый экземпляр хранилища Elasticsearch с указанным адресом
func NewElasticStore(address string) (*ElasticsearchStore, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			address,
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	// Check if index exists
	req := esapi.IndicesExistsRequest{
		Index: []string{"places"},
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error checking if index exists: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		// Index does not exist, create it
		mappings := `{
			"mappings": {
				"properties": {
					"name": {"type": "text"},
					"address": {"type": "text"},
					"phone": {"type": "text"},
					"location": {"type": "geo_point"}
				}
			}
		}`

		req := esapi.IndicesCreateRequest{
			Index: "places",
			Body:  strings.NewReader(mappings),
		}

		res, err := req.Do(context.Background(), es)
		if err != nil {
			log.Fatalf("Error creating index: %s", err)
		}
		defer res.Body.Close()

		if res.IsError() {
			log.Printf("[%s] Error creating index", res.Status())
		} else {
			log.Printf("[%s] Index created", res.Status())
		}
	} else if res.IsError() {
		log.Printf("[%s] Error checking if index exists", res.Status())
	}

	return &ElasticsearchStore{client: es}, nil
}

// GetPlaces возвращает список мест из хранилища Elasticsearch
func (s *ElasticsearchStore) GetPlaces(limit int, offset int) ([]Place, int, error) {
	var r map[string]interface{}

	req := esapi.SearchRequest{
		Index: []string{"places"},
		From:  &offset,
		Size:  &limit,
		Body: strings.NewReader(`{
			"track_total_hits": true
		}`),
	}

	res, err := req.Do(context.Background(), s.client)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, 0, err
	}

	hits := r["hits"].(map[string]interface{})["hits"].([]interface{})
	places := make([]Place, len(hits))
	for i, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		places[i] = Place{
			ID:      source["id"].(string),
			Name:    source["name"].(string),
			Address: source["address"].(string),
			Phone:   source["phone"].(string),
			Location: struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			}{
				Lat: source["location"].(map[string]interface{})["lat"].(float64),
				Lon: source["location"].(map[string]interface{})["lon"].(float64),
			},
		}
	}

	total := int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))

	return places, total, nil
}

// GetClosestPlaces возвращает список ближайших мест из хранилища Elasticsearch
func (s *ElasticsearchStore) GetClosestPlaces(lat float64, lon float64, limit int) ([]Place, error) {
	var r map[string]interface{}

	query := fmt.Sprintf(`{
		"size": %d,
		"sort": [
			{
				"_geo_distance": {
					"location": {
						"lat": %f,
						"lon": %f
					},
					"order": "asc",
					"unit": "km",
					"mode": "min",
					"distance_type": "arc",
					"ignore_unmapped": true
				}
			}
		]
	}`, limit, lat, lon)

	req := esapi.SearchRequest{
		Index: []string{"places"},
		Body:  strings.NewReader(query),
	}

	res, err := req.Do(context.Background(), s.client)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	hits := r["hits"].(map[string]interface{})["hits"].([]interface{})
	places := make([]Place, len(hits))
	for i, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		places[i] = Place{
			ID:      source["id"].(string),
			Name:    source["name"].(string),
			Address: source["address"].(string),
			Phone:   source["phone"].(string),
			Location: struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			}{
				Lat: source["location"].(map[string]interface{})["lat"].(float64),
				Lon: source["location"].(map[string]interface{})["lon"].(float64),
			},
		}
	}

	return places, nil
}
