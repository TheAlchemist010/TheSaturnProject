document.getElementById("actionButton").addEventListener("click", () => {
    fetch("/api", {
                method: "POST",
        headers: {
                        "Content-Type": "application/json"
                    
        },
                body: JSON.stringify({ action: "move"  })
            
    })
        .then(response => response.text())
        .then(data => console.log(data))
        .catch(error => console.error("Error:", error));
    
});

