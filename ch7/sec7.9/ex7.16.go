package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Request struct {
	Expr string
	Env  Env
}

type Response struct {
	Value float64
}

func calculate(w http.ResponseWriter, r *http.Request) {
	defer func() {
		_ = r.Body.Close()
	}()

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}

	expr, err := Parse(req.Expr)
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}

	resp := Response{
		Value: expr.Eval(req.Env),
	}
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(resp)
}

func main() {
	http.HandleFunc("/calculate", calculate)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
