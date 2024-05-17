package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"math/rand"
	"strconv"
)

type Movie struct{
	ID string `json: "id"`
	Isbn string `json: "isbn"`
	title string `json: "title"`
	Director *Director `json:"director"`

}
type Director struct{
	Firstname: string `json: "firstname"`
	Lastname: string `json: "Lastname"`
}

func main() {

}