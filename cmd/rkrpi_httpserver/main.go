package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Msg struct {
	Id  int    `db:"id"`
	Msg string `db:"msg"`
}

type Handler struct {
	db *sqlx.DB
}

func (h *Handler) GetIndex(w http.ResponseWriter, r *http.Request) {
	msgs := []Msg{}
	err := h.db.Select(&msgs, "SELECT * FROM msgs", 100)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(msgs)
	}
}

func main() {
	db, err := sqlx.Connect("sqlite3", "rkrpi.db")
	if err != nil {
		log.Fatalln(err)
	}

	handler := Handler{db}

	http.HandleFunc("/", handler.GetIndex)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
