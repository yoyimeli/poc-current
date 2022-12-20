package main

import (
	"fmt"
	"log"

	"github.com/yoyimeli/poc-current/current"
)

func main() {
	manyStatements, err := current.GetOpenStatement[[]current.StatementSummary]("many")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(manyStatements)

	oneStatements, err := current.GetOpenStatement[current.StatementSummary]("one")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(oneStatements)

	person, err := current.GetOpenStatement[current.Person]("person")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(person)
}
