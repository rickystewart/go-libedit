package main

import (
	"fmt"
	"io"
	"log"
	"strings"

	libedit "github.com/knz/go-libedit"
)

type example struct{}

func (_ example) GetLeftPrompt(_ libedit.EditLine) string {
	return "hello> "
}

func (_ example) GetRightPrompt(_ libedit.EditLine) string {
	return "(-:"
}

func (_ example) GetCompletions(word string) []string {
	if strings.HasPrefix(word, "he") {
		return []string{"hello!"}
	}
	return nil
}

func main() {
	el, err := libedit.Init("example")
	if err != nil {
		log.Fatal(err)
	}
	defer el.Close()

	el.UseHistory(-1, true)
	el.SetCompleter(example{})
	el.SetLeftPrompt(example{})
	el.SetRightPrompt(example{})
	for {
		s, err := el.GetLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			if err == libedit.ErrInterrupted {
				fmt.Println("interrupted!")
				continue
			}
			log.Fatal(err)
		}
		fmt.Println("echo", s)
		if err := el.AddHistory(s); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("goodbye!")
}
