package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", adderHandler)
	http.HandleFunc("/api/adder", adderApiHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func getAdderResult(firstNumber int, secondNumber int) int {
	return firstNumber + secondNumber
}

func adderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		adderHandlerGet(w, r)
	} else if r.Method == "POST" {
		adderHandlerPost(w, r)
	} else {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
	}
}

func adderApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		adderApiHandlerPost(w, r)
	} else {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
	}
}

func adderHandlerGet(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func adderHandlerPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}

	firstNumber, err := strconv.Atoi(r.FormValue("firstNumber"))
	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}

	secondNumber, err := strconv.Atoi(r.FormValue("secondNumber"))
	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}

	fmt.Fprintf(w, "First number = %d\n", firstNumber)
	fmt.Fprintf(w, "Second number = %d\n", secondNumber)

	result := getAdderResult(firstNumber, secondNumber)
	fmt.Fprintf(w, "Result = %d\n", result)
}

func adderApiHandlerPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "ParseForm problem", http.StatusNotFound)
		return
	}

	firstNumber, err := strconv.Atoi(r.FormValue("firstNumber"))
	if err != nil {
		http.Error(w, "firstNumber is wrong", http.StatusNotFound)
		return
	}

	secondNumber, err := strconv.Atoi(r.FormValue("secondNumber"))
	if err != nil {
		http.Error(w, "secondNumber is wrong", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]int)

	result := getAdderResult(firstNumber, secondNumber)
	resp["result"] = result

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}
