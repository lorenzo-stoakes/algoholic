package main

import (
	"./algoholic"
	"bufio"
	"fmt"
	"io"
	"os"
)

const DICT_PATH = "/usr/share/dict/words"

func readDict() []string {
	var ret []string

	if file, err := os.Open(DICT_PATH); err == nil {
		defer file.Close()
		reader := bufio.NewReader(file)

		for word, err := reader.ReadString('\n'); err != io.EOF; word, err = reader.ReadString('\n') {
			if err != nil {
				panic(err)
			}

			// Strip newline.
			word = word[:len(word)-1]
			ret = append(ret, word)
		}
	} else {
		panic(err)
	}

	return ret
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}()

	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: prefix [prefix]\n")
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "Reading dictionary... ")
	words := readDict()
	fmt.Fprintln(os.Stderr, "Done.")

	fmt.Fprintf(os.Stderr, "Creating trie... ")
	trie := algoholic.NewTrieFromStrings(words)
	fmt.Fprintln(os.Stderr, "Done.")

	prefix := os.Args[1]

	for _, suffix := range trie.FindSuffixes(prefix) {
		fmt.Println(suffix)
	}
}
