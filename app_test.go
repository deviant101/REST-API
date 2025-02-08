package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var app App

func TestMain(m *testing.M) {

	err := app.Initialize()
	if err != nil {
		log.Fatal("Database Initialization Error")
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
	app.DB.Exec("ALTER TABLE items AUTO_INCREMENT = 1")
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

func TestCreateProduct(t *testing.T) {
	clearTable()
	var payload = []byte(`{"name": "Product2", "quantity": 20, "price": 250.00}`)
	request, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(payload))
	request.Header.Set("Content-Type", "application/json")

	response := sendRequest(request)
	checkStatusCode(t, http.StatusCreated, response.Code)

	var responseMap map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &responseMap)

	if responseMap["name"] != "Product2" {
		t.Errorf("Expected: Product2, Received: %v", responseMap["name"])
	}
	if responseMap["quantity"] != 20.0 {
		t.Errorf("Expected: 20, Received: %v", responseMap["quantity"])
	}
}

func TestDeleteProduct(t *testing.T) {
	clearTable()
	addProduct("newProduct", 10, 100.00)

	request, _ := http.NewRequest("GET", "/product/1", nil)
	response := sendRequest(request)
	checkStatusCode(t, http.StatusOK, response.Code)

	request, _ = http.NewRequest("DELETE", "/product/1", response.Body)
	response = sendRequest(request)
	checkStatusCode(t, http.StatusOK, response.Code)

	request, _ = http.NewRequest("GET", "/product/1", nil)
	response = sendRequest(request)
	checkStatusCode(t, http.StatusNotFound, response.Code)

}
func TestUpdateProduct(t *testing.T) {
	clearTable()
	addProduct("Product1", 10, 100.00)

	request, _ := http.NewRequest("GET", "/product/1", nil)
	response := sendRequest(request)

	var oldProduct map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &oldProduct)

	var payload = []byte(`{"name": "Product1", "quantity": 10, "price": 255.00}`)
	request, _ = http.NewRequest("PUT", "/product/1", bytes.NewBuffer(payload))
	request.Header.Set("Content-Type", "application/json")

	response = sendRequest(request)
	var newProduct map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &newProduct)

	if newProduct["id"] != oldProduct["id"] {
		t.Errorf("Expected: %v, Received: %v", newProduct["id"], oldProduct["id"])
	}
	if newProduct["name"] != oldProduct["name"] {
		t.Errorf("Expected: %v, Received: %v", newProduct["name"], oldProduct["name"])
	}
	if newProduct["price"] == oldProduct["price"] { //because we have updated the price only
		t.Errorf("Expected: %v, Received: %v", newProduct["price"], oldProduct["price"])
	}
	if newProduct["quantity"] != oldProduct["quantity"] {
		t.Errorf("Expected: %v, Received: %v", newProduct["quantity"], oldProduct["quantity"])
	}
}
