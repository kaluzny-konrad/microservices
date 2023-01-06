package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", clockHandler)
	http.Handle("/css/", http.FileServer(http.Dir("static")))
	http.Handle("/js/", http.FileServer(http.Dir("static")))

	http.HandleFunc("/api/clockWidget/clockHtml", clockWidgetHtmlHandler)
	http.HandleFunc("/api/clockWidget/clockJs", clockWidgetJsHandler)

	http.HandleFunc("/api/clockWidget/clockText", clockWidgetTextHandler)

	fmt.Printf("Starting server at port 7001\n")
	if err := http.ListenAndServe(":7001", nil); err != nil {
		log.Fatal(err)
	}
}

func clockHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "index.html")
	} else {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
	}
}

func clockWidgetHtmlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Printf("Serving clock widget.html\n")
		http.ServeFile(w, r, "api/clockWidget.html")
	} else {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
	}
}

func clockWidgetJsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Printf("Serving clock widget.js\n")
		http.ServeFile(w, r, "api/clockWidget.js")
	} else {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
	}
}

func clockWidgetTextHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)

		result := getTime()
		resp["time"] = result

		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	} else {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
	}
}

func getTime() string {
	date := time.Now()
	h := date.Hour()
	fmt.Printf("%d - ", h)
	m := int(date.Month())
	s := date.Second()
	session := "AM"

	if h == 0 {
		h = 12
	}

	if h > 12 {
		h = h - 12
		session = "PM"
	}

	hStr := getClockVariable(h)
	mStr := getClockVariable(m)
	sStr := getClockVariable(s)

	time := hStr + ":" + mStr + ":" + sStr + " " + session
	return time
}

func getClockVariable(v int) (result string) {
	result = fmt.Sprint(v)
	if v < 10 {
		result = "0" + result
	}
	return
}
