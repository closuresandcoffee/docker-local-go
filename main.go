package main
 
import (
    "fmt"
    "io"
    "net/http"
    "time"
)
 
func MainHandler(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, time.Now().Format("2006-01-02 15:04:05"))
}
 
func main() {
    http.HandleFunc("/", MainHandler)
 
    fmt.Println("Listening on port 8000...")
 
    http.ListenAndServe(":8000", nil)
}