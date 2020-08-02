package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"


	"github.com/gocolly/colly"
)

func main() {
	prbPtr := flag.Int("p", 10, "retrieves the problem from ProjectEuler.net")

	flag.Parse()

	p := *prbPtr
	t, txt := visitWebsite(p)
	h, _ := os.UserHomeDir()
	d := fmt.Sprintf("%s/Code/ProjectEuler/Problems/%04d", h, p)
	if _, err := os.Stat(d); os.IsNotExist(err) {
		_ = os.MkdirAll(d, 0777)
		_ = os.MkdirAll(d+"/solution", 0777)
		gf := []byte(fmt.Sprintf("package main\n\nimport (\n\t\"log\"\n\t\"time\"\n\n\thelper \"github.com/cuken/ProjectEuler/Helper\"\n)\n\nfunc main() {\ndefer helper.TrackTime(time.Now(), \"Problem %v\")\n\n}", p))
		err := ioutil.WriteFile(fmt.Sprintf("%s/problem%v.go", d, p), gf, 0777)
		if err != nil {
			panic(err)
		}
		rm := []byte(fmt.Sprintf("# Problem %v\n\n[Link to Problem](https://projecteuler.net/problem=%v)\n\n## %s\n\n```%s```\n\n### Thought Process\n\n", p, p, t, txt))
		err = ioutil.WriteFile(fmt.Sprintf("%s/readme.md", d), rm, 0777)
		if err != nil {
			panic(err)
		}
		srm := []byte(fmt.Sprintf("# Solution\n\n![solution](problem%v.png)", p))
		err = ioutil.WriteFile(fmt.Sprintf("%s/solution/readme.md", d), srm, 0777)
		if err != nil {
			panic(err)
		}
	}
}

func visitWebsite(p int) (title string, text string) {
	// response, err := http.Get(fmt.Sprintf("https://projecteuler.net/problem=%v", p))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer response.Body.Close()

	// buf := new(bytes.Buffer)
	// buf.ReadFrom(response.Body)
	// s := buf.String()
	var t, txt string

	// log.Printf("Website Content ->\n%s", s)

	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("h2", func(e *colly.HTMLElement) {
		s, _ := e.DOM.Html()
		t = s
	})
	c.OnHTML(".problem_content", func(e *colly.HTMLElement) {
		s, _ := e.DOM.Html()
		txt = s
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err := c.Visit(fmt.Sprintf("https://projecteuler.net/problem=%v", p))
	if err != nil {
		fmt.Printf("Error when visiting site: %v", err)
	}

	txt = strings.Replace(txt, "<p>", "", -1)
	txt = strings.Replace(txt, "</p>", "", -1)

	return t, txt
}
