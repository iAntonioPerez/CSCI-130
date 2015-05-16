//Antonio Perez
package main

import (
	//"fmt"
	"log"
	"os"
	"text/template"
)

type guest struct {
	Honorific string
	Last      string
	Attended  bool
	Donated   bool
	Upcomming []string
}

func main() {
	const formletter = `
Dear {{.Honorific}}{{.Last}},
{{if .Attended}}
It was a pleasure to see you at the fundraiser.{{else}}
It is a shame you couldn't attend the fundraiser.{{end}}
{{if .Donated}}
Thank you for donating.{{end}}
{{range .Upcomming}}
    Remember the upcoming events: {{.}}
{{end}}
Best wishes,
The Pretender
`
	var guests = []guest{
		{"Mr. ", "Grohl", true, true, []string{"event1, event2, event3"}},
		{"Mrs. ", "Hawkins", true, false, []string{"event1, event2, event3"}},
		{"Ms. ", "Shiftlett", false, false, []string{"event1, event2, event3"}},
	}

	temp := template.Must(template.New("formletter").Parse(formletter))

	for _, g := range guests {
		err := temp.Execute(os.Stdout, g)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

}
