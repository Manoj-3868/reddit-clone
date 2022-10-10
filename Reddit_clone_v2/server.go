package main

//go mod init 31Jul2022GolangRestfulApi
//go get -u github.com/gorilla/mux
//go get go.mongodb.org/mongo-driver
//go mod tidy
import (
	//"encoding/json"
	"fmt"
	"log"
	"net/http"
	router "reddit_clone_v2/Router"
)

func main() {
	//go run .
	router := router.GetRouter()
	fmt.Println("in main")

	log.Fatal(http.ListenAndServe(":3000", router))
	//log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("server started on port 3000")
}
