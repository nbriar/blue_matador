package chuck_norris

import (
	"math/rand"
	"time"

	"github.com/dnaeon/go-chucknorris/api"
)

// JokeByCategory - Finds a joke in the specified category
func JokeByCategory(c string) string {
	client := api.NewClient(nil)
	category := api.Category(c)
	joke, err := client.RandomJoke(category)
	if err != nil {
		return err.Error()
	}

	return joke.Value
}

// JokeSearch - Find jokes which include the query string
func JokeSearch(q string) string {
	client := api.NewClient(nil)
	result, err := client.Search(q)
	if err != nil {
		return err.Error()
	}

	if len(result.Result) == 0 {
		return ""
	}

	if len(result.Result) == 1 {
		return result.Result[0].Value
	}

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	i := r.Intn(len(result.Result))

	return result.Result[i].Value
}
