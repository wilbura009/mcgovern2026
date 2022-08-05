package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Define a template.
	const letter = `
Dear {{.Name}},
{{if .Attended}}
	It was a pleasure to battle you.
{{- else}}
	It is a shame you couldn't make it to the wedding.
{{- end}}
{{- with .Gift}}
	Thank you for the lovely {{.}}.
{{end}}
Best wishes,
{{if .Name}}
Gorr
{{else}}
{{end}}
`
	// Prepare some data to insert into the template.
	type Recipient struct {
		Name, Gift string
		Attended bool
	}
	var recipients = []Recipient{
		{"Thor Odinson", "Stormbreaker", true},
		{"Gorr", "", false},
	}
	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	// Execute the template.
	err := t.Execute(os.Stdout, recipients[0])
	if err != nil {
		log.Println("executing template:", err)
	}
/*
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
*/

/* ----------------------------------------------------------------------------
	// Prepare some data to insert into the template.
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	var recipients = []Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin pants", false},
		{"Cousin Rodney", "", false},
	}
	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
*/
}
