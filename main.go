package main

import (
        "fmt"
        "net/http"
        "strings"
        "the-saturn-project/api"            
)

func main() {
        http.Handle("/", http.FileServer(http.Dir("./static"))) // Serves files from the "static" folder
        http.HandleFunc("/api/", apiHandler)                    // Custom API endpoint for game logic
        fmt.Println("Server running on http://localhost:8080")
        http.ListenAndServe(":8080", nil)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
    // Extract the path after /api/
    path := strings.TrimPrefix(r.URL.Path, "/api/")
    switch {
        case strings.HasPrefix(path, "serverstatus"):
           api.ServerStatusHandle(w, r)
        case strings.HasPrefix(path, "gamestate"):
            api.GetGameStateHandle(w, r)
        default:
            http.Error(w, "Endpoint not found", http.StatusNotFound)
    }
}

