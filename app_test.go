package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var app App

func TestMain(m *testing.M) {

	err := app.Initialize(DBUser, DBPass, "test_db")
	if err != nil {
		log.Fatal("Database Initializtion Error")
	}
	createTables()
	m.Run()
}

func createTables() {
	createTableQuery := `CREATE TABLE IF NOT EXISTS items (
		id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		quantity INT,
		price FLOAT(10,7)
	);`

	_, err := app.DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	app.DB.Exec("DELETE FROM items")
	log.Println("Table Cleared")
}

func addProduct(name string, quantity int, price float64) {
	_, err := app.DB.Exec("INSERT INTO items(name, quantity, price) VALUES (?, ?, ?)", name, quantity, price)
	if err != nil {
		log.Println(err)
	}
}

func TestGetProduct(t *testing.T) {
	clearTable()
	addProduct("Product1", 10, 100.00)
	request, _ := http.NewRequest("GET", "/products", nil)
	response := sendRequest(request)
	checkStatusCode(t, http.StatusOK, response.Code)

}

func checkStatusCode(t *testing.T, expectedCode int, actualCode int) {
	if expectedCode != actualCode {
		t.Errorf("Expected status : %v, Received %v", expectedCode, actualCode)
	}
}

func sendRequest(request *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, request)
	return rr

}
