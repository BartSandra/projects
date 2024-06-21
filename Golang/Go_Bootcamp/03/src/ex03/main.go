package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"src/db"
	"strconv"
)

type Server struct {
	store db.Store
}

// NewServer создает новый экземпляр сервера с указанным хранилищем
func NewServer(store db.Store) *Server {
	return &Server{store: store}
}

// handler обрабатывает запросы к корневому эндпоинту, возвращая список мест с пагинацией
func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	pageParam := r.URL.Query().Get("page")
	if pageParam == "" {
		pageParam = "1"
	}
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		http.Error(w, "Invalid 'page' value: "+pageParam, http.StatusBadRequest)
		return
	}

	limit := 10
	offset := (page - 1) * limit
	places, total, err := s.store.GetPlaces(limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lastPage := (total + limit - 1) / limit
	if page > lastPage {
		http.Error(w, fmt.Sprintf("Invalid 'page' value: %d. Total pages: %d.", page, lastPage), http.StatusBadRequest)
		return
	}

	data := struct {
		Places []db.Place
		Total  int
		Page   int
	}{
		Places: places,
		Total:  total,
		Page:   page,
	}

	funcs := template.FuncMap{
		"sub": func(a, b int) int {
			return a - b
		},
		"add": func(a, b int) int {
			return a + b
		},
		"ceilDiv": func(a, b int) int {
			return int(math.Ceil(float64(a) / float64(b)))
		},
	}

	tmpl := template.Must(template.New("").Funcs(funcs).Parse(`<!doctype html>
<html>
<head>
    <meta charset="utf-8">
    <title>Places</title>
    <meta name="description" content="">
    <meta name="viewport" content="width=device-width, initial-scale=1">
</head>
<body>
<h5>Total: {{.Total}}</h5>
<ul>
    {{range .Places}}
    <li>
        <div>{{.Name}}</div>
        <div>{{.Address}}</div>
        <div>{{.Phone}}</div>
    </li>
    {{end}}
</ul>
{{if gt .Page 1}}<a href="/?page={{sub .Page 1}}">Previous</a>{{end}}
{{if lt .Page 1365}}<a href="/?page={{add .Page 1}}">Next</a>{{end}}
<a href="/?page=1365">Last</a>
</body>
</html>`))

	tmpl.Execute(w, data)
}

// apiHandler обрабатывает запросы к /api/places, возвращая список мест в формате JSON с пагинацией
func (s *Server) apiHandler(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		http.Error(w, "Invalid 'page' value", http.StatusBadRequest)
		return
	}

	limit := 10
	offset := (page - 1) * limit
	places, total, err := s.store.GetPlaces(limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Name     string      `json:"name"`
		Total    int         `json:"total"`
		Places   []db.Place  `json:"places"`
		PrevPage interface{} `json:"prev_page"`
		NextPage interface{} `json:"next_page"`
		LastPage int         `json:"last_page"`
	}{
		Name:     "Places",
		Total:    total,
		Places:   places,
		PrevPage: nil,
		NextPage: nil,
		LastPage: (total + limit - 1) / limit,
	}

	if page > 1 {
		data.PrevPage = page - 1
	}
	if page < data.LastPage {
		data.NextPage = page + 1
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// recommendHandler обрабатывает запросы к /api/recommend, возвращая список ближайших мест в формате JSON
func (s *Server) recommendHandler(w http.ResponseWriter, r *http.Request) {
	lat, err := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	if err != nil {
		http.Error(w, "Invalid 'lat' value", http.StatusBadRequest)
		return
	}

	lon, err := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)
	if err != nil {
		http.Error(w, "Invalid 'lon' value", http.StatusBadRequest)
		return
	}

	places, err := s.store.GetClosestPlaces(lat, lon, 3)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Name   string     `json:"name"`
		Places []db.Place `json:"places"`
	}{
		Name:   "Recommendation",
		Places: places,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// main инициализирует сервер и начинает слушать входящие HTTP-запросы
func main() {
	store, err := db.NewElasticStore("http://localhost:9200")
	if err != nil {
		log.Fatalf("Error creating the store: %s", err)
	}
	server := NewServer(store)

	http.HandleFunc("/", server.handler)
	http.HandleFunc("/api/places", server.apiHandler)
	http.HandleFunc("/api/recommend", server.recommendHandler)
	http.ListenAndServe(":8888", nil)
}
