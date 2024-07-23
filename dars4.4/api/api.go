package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "github.com/go-redis/redis/v8"
    "github.com/gorilla/mux"
)

var rdb *redis.Client

func main() {
    rdb = redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })
    defer rdb.Close()

    r := mux.NewRouter()
    r.HandleFunc("/stock/{company}/{extreme}", getStockExtreme).Methods("GET")

    log.Fatal(http.ListenAndServe(":8080", r))
}

func getStockExtreme(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    company := vars["company"]
    extreme := vars["extreme"]

    ctx := context.Background()

    var key string
    if extreme == "high" {
        key = fmt.Sprintf("%s_high", company)
    } else if extreme == "low" {
        key = fmt.Sprintf("%s_low", company)
    } else {
        http.Error(w, "Invalid extreme value. Use 'high' or 'low'.", http.StatusBadRequest)
        return
    }

    val, err := rdb.Get(ctx, key).Result()
    if err == redis.Nil {
        http.Error(w, "No data found", http.StatusNotFound)
        return
    } else if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := map[string]string{
        "company": company,
        "extreme": extreme,
        "value":   val,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}