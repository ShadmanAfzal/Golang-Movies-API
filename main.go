package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/shadmanAfzal/movies_api/router"
)

const PORT = 8000

func main() {
	r := router.GetRouter()
	fmt.Println("Serving at http://localhost:" + strconv.Itoa(PORT))
	log.Fatal(http.ListenAndServe(":8000", r))
}
