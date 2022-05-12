package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"time"
)

func main() {
	timeResource := TimeResource{}
	http.ListenAndServe("localhost:8080", timeResource.Router())

}

func (t TimeResource) Router() chi.Router {
	r := chi.NewRouter()

	r.Get("/time/current", t.CurrentTime)

	return r
}

func (t TimeResource) CurrentTime(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, AmazTime{time.Now()})
}

type TimeResource struct {
}

type AmazTime struct {
	CurrentTime time.Time `json:"current_time"`
}
