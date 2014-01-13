# Compiling for x64, adjust accordingly for alternative platforms.

all: prefix

clean:
	rm algoholic.6 prefix.6 prefix

algoholic.6: ../trie.go
	go tool 6g -o algoholic.6 ../trie.go

prefix.6: algoholic.6 prefix.go
	go tool 6g -o prefix.6 prefix.go

prefix: prefix.6
	go tool 6l -o prefix prefix.6

.PHONY: all clean
