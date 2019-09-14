package examples

import (
	"fmt"
	"log"

	"github.com/suzuki-shunsuke/goarg"
)

func Example() {
	parser := goarg.NewParser()
	_, b, err := parser.
		Add(func(args ...string) (interface{}, error) {
			fmt.Println(args)
			fmt.Println("git status")
			return nil, nil
		}, "git", "status").
		Add(func(args ...string) (interface{}, error) {
			fmt.Println(args)
			fmt.Println("git add")
			return nil, nil
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
