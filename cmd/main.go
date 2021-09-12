package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"log"
	"github.com/joho/godotenv"
)

func main() {
	Env_load()
	http.HandleFunc("/api/sample", sample)

	port := os.Getenv("PORT")

	http.ListenAndServe(":" + port, nil)
}

func Env_load() {
	log.Print("================")
	log.Print(os.Getenv("GO_ENV"))
	log.Print(os.Getenv("CORS_ORIGIN"))
	log.Print("================")
	err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GO_ENV")))

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func sample(w http.ResponseWriter, _r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("CORS_ORIGIN"))

	ping := map[string]string{"message": "Hello World!"}

	res, _ := json.Marshal(ping)

	w.Write(res)
}
