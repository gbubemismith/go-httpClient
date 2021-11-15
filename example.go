package main

import (
	"fmt"

	"github.com/gbubemismith/go-httpClient/gohttp"
)

//define variables in bulk
var (
	//singleton http client
	githubHttpClient = getGithubClient()
)

func getGithubClient() gohttp.Client {
	client := gohttp.NewBuilder().Build()

	return client
}

func main() {
	getUrls()

}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func getUrls() {
	response, err := githubHttpClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	//using custom response
	var user User
	if err := response.UnmarshalJson(&user); err != nil {
		panic(err)
	}

	fmt.Println(response.Status())
	fmt.Println(response.StatusCode())
	fmt.Println(response.String())

}

// func createUser(user User) {
// 	response, err := githubHttpClient.Post("https://api.github.com", nil, user)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(response.StatusCode)
// 	bytes, _ := ioutil.ReadAll(response.Body)
// 	fmt.Println(string(bytes))
// }
