package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		if _, err := fmt.Fprintln(w, "ping"); err != nil {
			log.Println("ping failed", err)
			return
		}
	}).Methods(http.MethodGet)

	r.HandleFunc("/carts/{id}", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			log.Println("carts failed", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		cart, ok := carts[id]
		if !ok {
			if err := json.NewEncoder(w).Encode(Response{
				Message: "not exist cart",
			}); err != nil {
				log.Println("carts failed", err)
				return
			}
			return
		}

		if err := json.NewEncoder(w).Encode(Response{
			Message: "ok",
			Data:    cart,
		}); err != nil {
			log.Println("carts failed", err)
		}
	}).Methods(http.MethodGet)

	r.HandleFunc("/carts", func(w http.ResponseWriter, req *http.Request) {
		var cart Cart
		if err := json.NewDecoder(req.Body).Decode(&cart); err != nil {
			if err := json.NewEncoder(w).Encode(Response{
				Message: "not correct cart",
			}); err != nil {
				log.Println("carts failed", err)
				return
			}
			return
		}

		savedCart, ok := carts[cart.ID]
		if !ok {
			if err := json.NewEncoder(w).Encode(Response{
				Message: "not exist cart",
			}); err != nil {
				log.Println("carts failed", err)
				return
			}
			return
		}

		if savedCart.Account != cart.Account {
			if err := json.NewEncoder(w).Encode(Response{
				Message: "not correct account",
			}); err != nil {
				log.Println("carts failed", err)
				return
			}
			return
		}

		carts[cart.ID] = cart

		if err := json.NewEncoder(w).Encode(Response{
			Message: "success",
		}); err != nil {
			log.Println("carts failed", err)
			return
		}
		return
	}).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":42069", r))
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
