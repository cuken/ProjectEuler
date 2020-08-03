package helper

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

// GenerateProblem will create the folder structure and starting code for the given (p) Project Euler Problem.
func GenerateProblem(p int, d string) (t string, txt string, err error) {
	t, txt = visitWebsite(p)
	if _, err := os.Stat(d); os.IsNotExist(err) {
		_ = os.MkdirAll(d+"/solution", 0777)
		gf := []byte(fmt.Sprintf("package main\n\nimport (\n\t\"log\"\n\t\"time\"\n\n\thelper \"github.com/cuken/ProjectEuler/Helper\"\n)\n\nfunc main() {\ndefer helper.TrackTime(time.Now(), \"Problem %v\")\n\n}", p))
		err := ioutil.WriteFile(fmt.Sprintf("%s/problem%v.go", d, p), gf, 0777)
		if err != nil {

			return t, txt, nil
		}
		rm := []byte(fmt.Sprintf("# Problem %v\n\n[Link to Problem](https://projecteuler.net/problem=%v)\n\n## %s\n\n```%s```\n\n### Thought Process\n\n", p, p, t, txt))
		err = ioutil.WriteFile(fmt.Sprintf("%s/readme.md", d), rm, 0777)
		if err != nil {

			return t, txt, nil
		}
		srm := []byte(fmt.Sprintf("# Solution\n\n![solution](problem%v.png)", p))
		err = ioutil.WriteFile(fmt.Sprintf("%s/solution/readme.md", d), srm, 0777)
		if err != nil {

			return t, txt, nil
		}
		err = saveWebsite(p, d+"/problem.html")
		if err != nil {

			return t, txt, nil
		}
	}
	return t, txt, nil
}

// SaveWebsite visits the Project Euler website, and saves an html file for the given file.
func saveWebsite(p int, f string) (err error) {
	c := colly.NewCollector()

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Received %vbyte response\n", len(r.Body))
		err := r.Save(f)
		if err != nil {
			fmt.Printf("Error when saving response from website: %s", err)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err = c.Visit(fmt.Sprintf("https://projecteuler.net/problem=%v", p))
	if err != nil {
		fmt.Printf("Error when visiting site: %v", err)
	}

	return nil

}

// VisitWebsite vists the Project Euler website for the provided problem and returns the problem as a text and title component.
func visitWebsite(p int) (title string, text string) {
	var t, txt string

	c := colly.NewCollector()

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
