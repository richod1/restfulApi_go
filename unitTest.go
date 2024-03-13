package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetItem(t *testing.T){
	req,err:=http.NewRequest("GET","/",nil)
	if err!=nil{
		t.Fatal(err)
	}

	rr:=httptest.NewRecorder()
	handler :=http.HandlerFunc(getItem)

	handler.ServeHTTP(rr,req)

}