package examples

import (
	"fmt"
	"log"

	"github.com/suzuki-shunsuke/goarg"
)

func Example() {
	parser := goarg.NewParser()
	b, err := parser.
		Add(func(args ...string) error {
			fmt.Println(args)
			fmt.Println("git status")
			return nil
		}, "git", "status").
		Add(func(args ...string) error {
			fmt.Println(args)
			fmt.Println("git add")
			return nil
		}, "git", "add").
		Parse("git", "add", "foo.txt")
	if !b {
		fmt.Println("No command matches")
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	// Output:
	// [foo.txt]
	// git add
}
