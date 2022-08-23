package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Dollar float64

func (d Dollar) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type Database struct {
	mu   sync.Mutex
	data map[string]Dollar
}

func (db *Database) List() map[string]Dollar {
	db.mu.Lock()
	defer db.mu.Unlock()

	prices := map[string]Dollar{}
	for i, p := range db.data {
		prices[i] = p
	}
	return prices
}

func (db *Database) Get(key string) (Dollar, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	value, ok := db.data[key]
	if !ok {
		return 0, fmt.Errorf("%v not found", key)
	}
	return value, nil
}

func (db *Database) Insert(key string, value Dollar) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[key] = value
}

func (db *Database) Update(key string, value Dollar) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.data[key]; !ok {
		return fmt.Errorf("%v not found", key)
	}

	db.data[key] = value
	return nil
}

func (db *Database) Delete(key string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.data[key]; !ok {
		return fmt.Errorf("%v not found", key)
	}

	delete(db.data, key)
	return nil
}

type Handler struct {
	db           *Database
	listTemplate *template.Template
}

func (h *Handler) list(w http.ResponseWriter, _ *http.Request) {
	prices := h.db.List()
	err := h.listTemplate.Execute(w, prices)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintln(w, err)
		return
	}
}

func (h *Handler) get(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, err := h.db.Get(item)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintln(w, err)
		return
	}
	_, _ = fmt.Fprintln(w, price)
}

func (h *Handler) insert(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	qPrice := req.URL.Query().Get("price")
	fPrice, err := strconv.ParseFloat(qPrice, 64)
	price := Dollar(fPrice)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "parsing price failed: %v", err)
		return
	}

	h.db.Insert(item, price)
	_, _ = fmt.Fprintf(w, "Added item: %s price: %s sucessfully", item, price)
}

func (h *Handler) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	qPrice := req.URL.Query().Get("price")
	fPrice, err := strconv.ParseFloat(qPrice, 64)
	price := Dollar(fPrice)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "parsing price failed: %v", err)
		return
	}

	if err := h.db.Update(item, price); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "updating failed: %v", err)
		return
	}
	_, _ = fmt.Fprintf(w, "Updated item: %s price: %s sucessfully", item, price)
}

func (h *Handler) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if err := h.db.Delete(item); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "deleting failed: %v", err)
		return
	}
	_, _ = fmt.Fprintf(w, "Deleted item: %s sucessfully", item)
}

func main() {
	tableTemplate, err := template.ParseFiles("ch7/sec7.7/table.html")
	if err != nil {
		log.Fatal("Parsing template file failed", err)
	}

	db := &Database{
		data: map[string]Dollar{
			"shoes": 50,
			"socks": 5,
		},
	}
	handler := Handler{
		db:           db,
		listTemplate: tableTemplate,
	}

	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(handler.list))
	mux.Handle("/price", http.HandlerFunc(handler.get))
	mux.Handle("/insert", http.HandlerFunc(handler.insert))
	mux.Handle("/update", http.HandlerFunc(handler.update))
	mux.Handle("/delete", http.HandlerFunc(handler.delete))

	log.Fatal(http.ListenAndServe(":8000", mux))
}
