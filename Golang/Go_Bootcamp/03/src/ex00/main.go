package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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

type Store interface {
	GetPlaces(limit int, offset int) ([]Place, int, error)
}

type ElasticsearchStore struct {
	client *elasticsearch.Client
}

// NewElasticsearchStore создает новый экземпляр хранилища Elasticsearch с указанным адресом
func NewElasticsearchStore(address string) (*ElasticsearchStore, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			address,
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	req := esapi.IndicesExistsRequest{
		Index: []string{"places"},
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error checking if index exists: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
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
	req := esapi.SearchRequest{
		Index: []string{"places"},
		From:  &offset,
		Size:  &limit,
	}

	res, err := req.Do(context.Background(), s.client)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	var r map[string]interface{}
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

	return places, int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)), nil
}

func main() {
	store, err := NewElasticsearchStore("http://localhost:9200")
	if err != nil {
		log.Fatalf("Error creating the store: %s", err)
	}

	f, err := os.Open("data.csv")
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = '\t'

	_, err = r.Read()
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	records, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	// Создание буфера для данных Bulk API
	var buf strings.Builder

	for _, record := range records {
		if record[0] == "" {
			log.Printf("Skipping record with empty ID: %v", record)
			continue
		}

		// Преобразование широты и долготы в числа с плавающей точкой
		lat, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			log.Fatalf("Error converting latitude to float: %s", err)
		}

		lon, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			log.Fatalf("Error converting longitude to float: %s", err)
		}

		place := Place{
			ID:      record[0],
			Name:    record[1],
			Address: record[2],
			Phone:   record[3],
			Location: struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			}{
				Lat: lat,
				Lon: lon,
			},
		}

		meta := map[string]interface{}{
			"index": map[string]interface{}{
				"_id": place.ID,
			},
		}

		data, err := json.Marshal(meta)
		if err != nil {
			log.Fatalf("Error marshalling metadata: %s", err)
		}
		buf.WriteString(string(data) + "\n")

		data, err = json.Marshal(place)
		if err != nil {
			log.Fatalf("Error marshalling place: %s", err)
		}
		buf.WriteString(string(data) + "\n")
	}

	// Отправка запроса Bulk API
	req := esapi.BulkRequest{
		Body:    strings.NewReader(buf.String()),
		Refresh: "true",
	}

	// Выполнение запроса Bulk API
	res, err := req.Do(context.Background(), store.client)
	if err != nil {
		log.Fatalf("Error sending bulk request: %s", err)
	}
	defer res.Body.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pageParam := r.URL.Query().Get("page")
		page := 1
		if pageParam != "" {
			var err error
			page, err = strconv.Atoi(pageParam)
			if err != nil || page < 1 {
				http.Error(w, "Invalid page number", http.StatusBadRequest)
				return
			}
		}

		places, _, err := store.GetPlaces(10, (page-1)*10)
		if err != nil {
			http.Error(w, "Error getting places", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "<!doctype html>\n<html>\n<body>\n<h1>Places</h1>\n<ul>\n")
		for _, place := range places {
			fmt.Fprintf(w, "<li><div>%s</div><div>%s</div><div>%s</div></li>\n", place.Name, place.Address, place.Phone)
		}
		fmt.Fprintf(w, "</ul>\n</body>\n</html>\n")
	})

	http.ListenAndServe(":8888", nil)
}
