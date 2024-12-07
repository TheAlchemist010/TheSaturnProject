package api

import (
    "encoding/json"
    "net/http"
    
)

func ServerStatusHandle(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        // Example: Send server status
        response := map[string]string{
            "status": "running",
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
        
    } else {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        
    }
    
}

