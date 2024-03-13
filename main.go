package main

import (
	// "fmt"
	"net/http"
	"strconv"
	"encoding/json"
	"log"

	"github.com/gorilla/mux"
)

type Item struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
}

type Response struct{
	Message string `json:"name"`
}

var items []Item

func main(){
	router:=mux.NewRouter();

	router.HandleFunc("/",getItems).Methods("GET")
	router.HandleFunc("/items/{id}",getItem).Methods("GET")
	router.HandleFunc("/items",createItem).Methods("POST")
	router.HandleFunc("/items/{id}",updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}",deleteItem).Methods("DELETE")
	// checking router test
	router.HandleFunc("/test",getItemss).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080",router))


	// log.Output(http.ListenAndServe("server is up on port :8080",router))
	


}

// checking json response
func getItemss(w http.ResponseWriter,r *http.Request){
	response:=Response{Message:"Hello Items"}
	json.NewEncoder(w).Encode(response)
}

// function to get items
func getItems(w http.ResponseWriter,r *http.Request){
	json.NewEncoder(w).Encode(items)

}

// function to get item
func getItem(w http.ResponseWriter,r *http.Request){
	params:=mux.Vars(r)
	id,err:=strconv.Atoi(params["id"])
	if err !=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	for _,item:=range items{
		if item.ID==id{
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.NotFound(w,r)
}

// function to createItem
func createItem(w http.ResponseWriter,r *http.Request){
	var item Item
	err:=json.NewDecoder(r.Body).Decode(&item)
	if err !=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	item.ID=len(items) +1;
	items=append(items,item)
	json.NewEncoder(w).Encode(item)
}

// function to update item
func updateItem(w http.ResponseWriter,r *http.Request){
	params:=mux.Vars(r)
	id,err:=strconv.Atoi(params["id"])
	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	var updatedItem Item
	err=json.NewDecoder(r.Body).Decode(&updatedItem)
	if err !=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	for i,item:=range items{
		if item.ID==id{
			items[i]=updatedItem
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}
	http.NotFound(w,r)
}

// function ro delete item
func deleteItem(w http.ResponseWriter,r *http.Request){
	params:=mux.Vars(r)

	id,err:=strconv.Atoi(params["id"])
	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	for i,item:=range items{
		if item.ID==id{
			items=append(items[:i],items[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
		}
	}

	http.NotFound(w,r)
}

