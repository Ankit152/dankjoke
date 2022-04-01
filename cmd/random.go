/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random dank joke 😆",
	Long:  `This commad fetches you with a random joke on its own!`,
	Run: func(cmd *cobra.Command, args []string) {
		jokeTerm, _ := cmd.Flags().GetString("term")
		if jokeTerm != "" {
			getRandomJokeWithTerm(jokeTerm)
		} else {
			getRandomJoke()
		}
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
	rootCmd.PersistentFlags().String("term", "", "A search result for a dank joke.")
}

// dank joke structure
type DankJoke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

// search result structure of dank jokes
type SearchResult struct {
	Results    json.RawMessage `json:"results"`
	Status     int             `json:"status"`
	SearchTerm string          `json:"search_term"`
	TotalJokes int             `json:"total_jokes"`
}

// function to return random joke
func getRandomJoke() {
	url := "https://icanhazdadjoke.com/"
	responseBytes := getJokeData(url)
	joke := DankJoke{}
	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		log.Printf("Couldnot unmarshall response %v", err)
		os.Exit(1)
	}
	fmt.Println(string(joke.Joke))
}

// gets the joke from the API
func getJokeData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)
	if err != nil {
		log.Printf("Could not get a dank joke. Error %v", err)
		os.Exit(1)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dankjoke CLI (github.com/Ankit152/dankjoke)")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request. Error %v", err)
		os.Exit(1)
	}
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Couldnot read the response. Error occured %v", err)
		os.Exit(1)
	}
	return responseBytes
}

// function that prints a random joke having a specified term
func getRandomJokeWithTerm(term string) {
	total, jokes := getRandomJokeDataWithTerm(term)
	randomJokeList(total, jokes)
}

// function that prints a random joke from the list of jokes
func randomJokeList(length int, jokeList []DankJoke) {
	rand.Seed(time.Now().Unix())
	min, max := 0, length-1
	if length <= 0 {
		err := fmt.Errorf("No joke found with this term")
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		index := min + rand.Intn(max-min)
		fmt.Println(jokeList[index].Joke)
	}
}

// get a random joke with with a specified term
func getRandomJokeDataWithTerm(jokeTerm string) (int, []DankJoke) {
	url := fmt.Sprintf("https://icanhazdadjoke.com/search?term=%s", jokeTerm)
	responseBytes := getJokeData(url)
	jokeWithTerm := SearchResult{}
	if err := json.Unmarshal(responseBytes, &jokeWithTerm); err != nil {
		log.Printf("Couldnot unmarshall response %v", err)
		os.Exit(1)
	}

	jokes := []DankJoke{}
	if err := json.Unmarshal(jokeWithTerm.Results, &jokes); err != nil {
		log.Printf("Couldnot unrmashall Joke with Terms")
		os.Exit(1)
	}
	return jokeWithTerm.TotalJokes, jokes
}
