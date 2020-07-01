package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/marceloaguero/serverless-api/users"
)

const timeout = time.Second * 5

// Delivery -
type delivery struct {
	usecase users.Usecase
}

// NewDelivery -
func newDelivery(uc users.Usecase) *delivery {
	return &delivery{
		usecase: uc,
	}
}

func writeErr(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func (d *delivery) Create(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	decoder := json.NewDecoder(r.Body)
	user := &users.User{}
	if err := decoder.Decode(&user); err != nil {
		writeErr(w, err)
		return
	}

	if err := d.usecase.Create(ctx, user); err != nil {
		writeErr(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Ok"))
}

func (d *delivery) Get(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	vars := mux.Vars(r)
	id := vars["id"]

	user, err := d.usecase.Get(ctx, id)
	if err != nil {
		writeErr(w, err)
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		writeErr(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (d *delivery) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	users, err := d.usecase.GetAll(ctx)
	if err != nil {
		writeErr(w, err)
		return
	}

	data, err := json.Marshal(users)
	if err != nil {
		writeErr(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (d *delivery) Update(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	decoder := json.NewDecoder(r.Body)
	user := &users.UpdateUser{}
	if err := decoder.Decode(&user); err != nil {
		writeErr(w, err)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	if err := d.usecase.Update(ctx, id, user); err != nil {
		writeErr(w, err)
		return
	}
}

func (d *delivery) Delete(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	vars := mux.Vars(r)
	id := vars["id"]

	if err := d.usecase.Delete(ctx, id); err != nil {
		writeErr(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Deleted"))
}

// Routes -
func Routes() (*mux.Router, error) {
	usecase, err := users.Init()
	if err != nil {
		log.Panic(err)
	}

	delivery := newDelivery(usecase)

	r := mux.NewRouter()
	r.HandleFunc("/users", delivery.Create).Methods("POST")
	r.HandleFunc("/users/{id}", delivery.Get).Methods("GET")
	r.HandleFunc("/users", delivery.GetAll).Methods("GET")
	r.HandleFunc("/users/{id}", delivery.Update).Methods("PUT")
	r.HandleFunc("/users/{id}", delivery.Delete).Methods("DELETE")

	return r, nil
}
