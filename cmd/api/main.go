package main

import (
	"github.com/bezaeel/gorray/internal/api"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func main() {
	api.Run("")
}
