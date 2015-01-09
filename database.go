package main

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type CoffeeRequest struct {
	Id             uint64
	Host           string
	Date           time.Time
	Lat, Lng, Dist float64
}

type CoffeeNowDatabase struct {
	Connection               *sql.DB
	StmtPostReq, StmtGetReqs *sql.Stmt
}

func NewCoffeeNowDatabase() (*CoffeeNowDatabase, error) {
	conn, err := sql.Open("postgres", "port=5433 user=postgres dbname=coffeenow")
	if err != nil {
		return nil, err
	}
	stmtPost, err := conn.Prepare("INSERT INTO coffeerequests (host, date, lat, lng) VALUES ($1, current_timestamp + 30 * interval '1 minute', $2, $3) RETURNING id;")
	if err != nil {
		return nil, err
	}
	stmtGet, err := conn.Prepare("SELECT *, earth_distance(ll_to_earth($1, $2), ll_to_earth(lat, lng)) as distance FROM coffeerequests WHERE earth_box(ll_to_earth($1, $2), $3) @> ll_to_earth(lat, lng) AND date > current_timestamp ORDER BY distance;")
	if err != nil {
		return nil, err
	}
	return &CoffeeNowDatabase{conn, stmtPost, stmtGet}, nil
}

func (cndb CoffeeNowDatabase) Close() {
	// This does not work as it should...
	//defer cndb.StmtPostReq.Close()
	//defer cndb.StmtGetReqs.Close()
	//defer cndb.Connection.Close()
}

func (cns *CoffeeNowDatabase) PostRequest(host string, lat float64, lng float64) (uint64, error) {
	var id uint64
	err := cns.StmtPostReq.QueryRow(host, lat, lng).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (cns *CoffeeNowDatabase) GetRequests(lat float64, lng float64, dist float64) ([]CoffeeRequest, error) {
	rows, err := cns.StmtGetReqs.Query(lat, lng, dist)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reqs := make([]CoffeeRequest, 0, 10)
	for rows.Next() {
		var req CoffeeRequest
		err := rows.Scan(&(req.Id), &(req.Host), &(req.Date), &(req.Lat), &(req.Lng), &(req.Dist))
		if err != nil {
			return nil, err
		}
		reqs = append(reqs, req)
	}
	return reqs, nil
}
