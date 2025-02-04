package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

	models "luisferllub230.com/github_user_activity/Models"
)

func GithubEvents(GithubUserName string, GithubPagesQuantity int) []models.Event {
	var url string = "https://api.github.com/users/" + GithubUserName + "/events?per_page=" + strconv.Itoa(GithubPagesQuantity)
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println("status code error: " + res.Status)
		os.Exit(1)
	}

	var events []models.Event
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&events)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return events
}

func main() {

	var Commands []models.Command = []models.Command{
		{
			Name:      "Github events",
			Flag:      "github",
			ShortFlag: "gh",
			Help:      "Usage: go run main.go -gh <github_user_name>",
		},
		{
			Name:      "Pages quantity",
			Flag:      "quantity",
			ShortFlag: "pq",
			Help:      "Usage: go run main.go -pq <github_pages_quantity>",
		},
		{
			Name:      "Help",
			Flag:      "help",
			ShortFlag: "h",
			Help:      "Usage: go run main.go -h",
		},
	}

	var GithubUserName string
	var PageQuantity int
	var ShowHelp bool

	flag.StringVar(&GithubUserName, Commands[0].ShortFlag, "", Commands[0].Help)
	flag.IntVar(&PageQuantity, Commands[1].ShortFlag, 30, Commands[1].Help)
	flag.BoolVar(&ShowHelp, Commands[2].ShortFlag, false, Commands[2].Help)
	flag.Parse()

	if ShowHelp {
		for _, cmd := range Commands {
			fmt.Println(cmd.Help)
		}
		os.Exit(0)
	}

	if GithubUserName != "" {
		var events []models.Event = GithubEvents(GithubUserName, PageQuantity)
		for _, event := range events {
			fmt.Println(event.Repo.Name)
		}
		os.Exit(0)
	}
}
