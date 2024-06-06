package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) Initialize(DBUser string, DBPass string, DBName string) error {

	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1)/%v", DBUser, DBPass, DBName)
	var err error
	app.DB, err = sql.Open("mysql", connectionString)

	if err != nil {
		return err
	}

	app.Router = mux.NewRouter().StrictSlash(true)
	app.handleRoutes()
	return nil
}

func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {

	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func sendError(w http.ResponseWriter, statusCode int, err string) {

	err_msg := map[string]string{"error": err}
	sendResponse(w, statusCode, err_msg)
}

func (app *App) getProducts(w http.ResponseWriter, r *http.Request) {

	products, err := getProducts(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error()) // Convert err to string using err.Error()
		return
	}
	sendResponse(w, http.StatusOK, products)
}

func (app *App) getProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	// product, err := getProduct(app.DB, id)
	if err != nil {
		sendError(w, http.StatusNotFound, "invalid product id")
		return
	}
	prod := Product{ID: id}
	err = prod.getProduct(app.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			sendError(w, http.StatusNotFound, "Product not found")
		default:
			sendError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	sendResponse(w, http.StatusOK, prod)

}

func (app *App) createProduct(w http.ResponseWriter, r *http.Request) {

	var prod Product
	err := json.NewDecoder(r.Body).Decode(&prod)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	if err := prod.createProduct(app.DB); err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusCreated, prod)
}

func (app *App) updateProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusNotFound, "invalid product id")
		return
	}
	var prod Product
	err = json.NewDecoder(r.Body).Decode(&prod)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	prod.ID = id
	if err := prod.updateProduct(app.DB); err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, prod)
}

func (app *App) deleteProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusNotFound, "invalid product id")
		return
	}
	defer r.Body.Close()
	var prod Product = Product{ID: id}
	if err := prod.deleteProduct(app.DB); err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, map[string]string{"result": "Successful deletion"})
}

func (app *App) handleRoutes() {

	app.Router.HandleFunc("/products", app.getProducts).Methods("GET")
	app.Router.HandleFunc("/product/{id}", app.getProduct).Methods("GET")
	app.Router.HandleFunc("/product", app.createProduct).Methods("POST")
	app.Router.HandleFunc("/product/{id}", app.updateProduct).Methods("PUT")
	app.Router.HandleFunc("/product/{id}", app.deleteProduct).Methods("DELETE")
}
