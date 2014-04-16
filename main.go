package main

import "os"
import "fmt"
import "sync"
import "time"
import "encoding/json"

type Config struct {
	Sites     []WebSite `json:"sites"`
	Selectors []string  `json:"Selectors"`
}

type WebSite struct {
	Url string
}

var config = new(Config)

func main() {

	setup()

	var wg sync.WaitGroup
	wg.Add(2)

	cs := make(chan WebSite)
	defer close(cs)

	for i := 0; i < len(config.Sites); i++ {
		go fetch(cs, config.Sites[i])
		go process(cs, &wg, i)
	}

	wg.Wait()
	resemble()
}

func setup() {
	configFile, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
}

func fetch(cs chan WebSite, site WebSite) {
	fmt.Println("Fetching: ", site.Url)

	// Simulate network lag
	time.Sleep(time.Second * 4)
	cs <- site
}

func process(cs chan WebSite, wg *sync.WaitGroup, i int) {
	s := <-cs

	for j := 0; j < len(config.Selectors); j++ {
		fmt.Println("-- Working ", config.Selectors[j], " with key ", i, " at ", s.Url);
	}

	wg.Done()
}

func resemble() {
	fmt.Println("Resemble")
}
