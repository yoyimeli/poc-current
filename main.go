package main

import (
	"fmt"
	"log"

	"github.com/yoyimeli/poc-current/current"
)

func main() {

	fmt.Println("Many Statements")
	manyStatements, err := current.GetOpenStatement[[]current.StatementSummary]("many")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(manyStatements)

	fmt.Println("One Statements")
	oneStatements, err := current.GetOpenStatement[current.StatementSummary]("one")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(oneStatements)

	fmt.Println("Person")
	person, err := current.GetOpenStatement[current.Person]("person")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(person)
}
