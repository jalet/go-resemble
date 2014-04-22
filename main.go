package main

import "os"
import "fmt"
import "sync"
import "strings"
import "os/exec"
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


	install()
	setup()

	numSites := len(config.Sites)
	numSelectors := len(config.Selectors)

	var wg sync.WaitGroup

	for i := 0; i < numSites; i++ {
		for j := 0; j < numSelectors; j++ {
			wg.Add(1)
			go process(&wg, config.Sites[i], config.Selectors[j])
		}
	}

	wg.Wait()
	resemble()
}

func install() {
	fmt.Println("-- Install node packages")
	cmd := exec.Command("npm", "install")
	err := cmd.Run()

	if err != nil {
		panic("Could not install node packages")
	}
}

func setup() {
	fmt.Println("-- Cleanup (Screenshots and failures)")
	cmd1 := exec.Command("rm", "-rf", "screenshots")
	cmd1.Stdin = strings.NewReader("y")
	cmd1.Run()

	cmd2 := exec.Command("rm", "-rf", "failures")
	cmd2.Stdin = strings.NewReader("y")
	cmd2.Run()

	configFile, err := os.Open("config.json")
	if err != nil {
		panic("Could not open config.json")
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
}

func process(wg *sync.WaitGroup, site WebSite, selector string) {

	fmt.Println("-- Working ", selector, " at ", site.Url)
	exec.Command("node_modules/casperjs/bin/casperjs", "test", "lib/a.js",
			"--url="+site.Url,
			"--selector="+selector,
			"--ignore-ssl-errors=yes").Run()


	wg.Done()
}

func resemble() {
	fmt.Println("-- Resemble")
	out, err := exec.Command("node_modules/casperjs/bin/casperjs", "test", "lib/b.js").Output()

	if err != nil {}

	fmt.Printf("%s\n", out)
}
