package main

import (
        "fmt"
            "net/http"
            
)

func main() {
        http.Handle("/", http.FileServer(http.Dir("./static"))) // Serves files from the "static" folder
            http.HandleFunc("/api", apiHandler)                    // Custom API endpoint for game logic
                fmt.Println("Server running on http://localhost:8080")
                    http.ListenAndServe(":8080", nil)
                    
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
                // Process incoming game data
                        fmt.Fprintf(w, "Game data received")
                            
                    } else {
                                http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
                                    
                    }
                    
}

