package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

//Create a struct to hold my movie elements
type Movie struct {
	Title string `json:"Title"`
}

//Add default welcome page
func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the default page!")
	fmt.Println("Endpoint Hit: root")
}

//Add routing to the http server
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/movies", getAllMoviesSortedDetails)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

//List of movies
var Movies = []Movie{
	Movie{Title: "Fargo"},
	Movie{Title: "Terminator"},
	Movie{Title: "Equalizer"},
	Movie{Title: "Tenet"},
	Movie{Title: "Mortal"},
}

func main() {
	fmt.Println("Server started:")
	handleRequests()
}

func getAllMoviesSortedDetails(w http.ResponseWriter, r *http.Request){
	//Get the API Key from env
	apiKey := os.Getenv("apiKey")

	//Print to stdOut
	fmt.Println("Endpoint Hit GET: movies")

	//Sort initial array of movies
	sort.Slice(Movies, func(i, j int) bool {
		return Movies[i].Title < Movies[j].Title
	})

	//Create an array with movie details from omdbapi
	var arrMovieDetails[]string
	for k := range Movies {
		title := Movies[k].Title
		response, _ := http.Get("http://www.omdbapi.com/?apikey=" + apiKey + "&t=" + title)
		responseBody, _ := ioutil.ReadAll(response.Body)
		responseString := string(responseBody)
		arrMovieDetails = append(arrMovieDetails,responseString)
	}

	//Print the array as json to the http get request
	fmt.Fprint(w,"["+strings.Join(arrMovieDetails,",")+"]")

}
