package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/api/sample", sample)

	port := os.Getenv("PORT")

	http.ListenAndServe(":" + port, nil)
}

func sample(w http.ResponseWriter, _r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	ping := map[string]string{"message": "Hello World!"}

	res, _ := json.Marshal(ping)

	w.Write(res)
}
