package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

const (
	defaultDictionaryFolder string = "dicts"
	defaultDictionaryFile   string = "polish.txt"
	defaultPasswordsCount   uint   = 10
	defaultWordsCount       uint   = 3
)

func main() {
	dictionary := flag.String("d", defaultDictionaryPath(), "path to dictionary file")
	passwordsCount := flag.Uint("p", defaultPasswordsCount, "number of passwords to generate")
	wordsCount := flag.Uint("w", defaultWordsCount, "number of words in each password")
	flag.Parse()

	passphrase := NewPassphrase()
	passwords, err := passphrase.GeneratePasswords(*dictionary, *passwordsCount, *wordsCount)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	for _, password := range passwords {
		fmt.Println(password)
	}
}

// defaultDictionaryPath returns full path to the default dictionary.
func defaultDictionaryPath() string {
	p, err := os.Executable()
	if err != nil {
		return path.Join(defaultDictionaryFolder, defaultDictionaryFile)
	}
	d := filepath.Dir(p)
	return path.Join(d, defaultDictionaryFolder, defaultDictionaryFile)
}
