package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type CoffeeNowServer struct {
	Db *CoffeeNowDatabase
}

func NewCoffeeNowServer() (*CoffeeNowServer, error) {
	db, err := NewCoffeeNowDatabase()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	return &CoffeeNowServer{db}, nil
}

func (cns *CoffeeNowServer) handleGet(w http.ResponseWriter, r *http.Request) {
	lat, err := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	if err != nil {
		fmt.Fprintf(w, "Error: Could not parse latitude.\n")
		w.WriteHeader(500)
		return
	}
	lng, err := strconv.ParseFloat(r.URL.Query().Get("lng"), 64)
	if err != nil {
		fmt.Fprintf(w, "Error: Could not parse longitude.\n")
		w.WriteHeader(500)
		return
	}
	reqs, err := cns.Db.GetRequests(lat, lng, 1000)
	if err != nil {
		fmt.Fprintf(w, "Error: Could not query database.\n%s\n", err)
		w.WriteHeader(500)
		return
	}
	for _, req := range reqs {
		fmt.Fprintf(w, "%s\n", req.Host)
	}
}
func (cns *CoffeeNowServer) handlePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	host, ok := r.Form["host"]
	if !ok {
		fmt.Fprintf(w, "Error: Could not parse host.\n")
		w.WriteHeader(500)
		return
	}
	latStr, ok := r.Form["lat"]
	if !ok {
		fmt.Fprintf(w, "Error: Could not find latitude.\n")
		w.WriteHeader(500)
		return
	}
	lat, err := strconv.ParseFloat(latStr[0], 64)
	if err != nil {
		fmt.Fprintf(w, "Error: Could not parse latitude.\n")
		w.WriteHeader(500)
		return
	}
	lngStr, ok := r.Form["lng"]
	if !ok {
		fmt.Fprintf(w, "Error: Could not find longitude.\n")
		w.WriteHeader(500)
		return
	}
	lng, err := strconv.ParseFloat(lngStr[0], 64)
	if err != nil {
		fmt.Fprintf(w, "Error: Could not parse longitude.\n")
		w.WriteHeader(500)
		return
	}

	id, err := cns.Db.PostRequest(host[0], lat, lng)
	if err != nil {
		fmt.Fprintf(w, "Error: Could not post request.\n")
		w.WriteHeader(500)
		return
	}
	fmt.Fprintf(w, "%d\n", id)
}

func (cns *CoffeeNowServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		cns.handleGet(w, r)
		return
	case r.Method == "POST":
		cns.handlePost(w, r)
		return
	}
}
