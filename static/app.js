document.getElementById("actionButton").addEventListener("click", () => {
    fetch("/api/serverstatus", {
        method: "GET"
    })
    .then(response => response.text())
    .then(data => console.log(data))
    .catch(error => console.error("Error:", error));
    
});

