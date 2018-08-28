package main

import (
	"bufio"
	"os"
	"strings"
)

const (
	maxPasswordsCount uint   = 128
	maxWordsCount     uint   = 16
	separators        string = " +-=:.,!#%&*_"
)

// Passphrase implements password generator.
type Passphrase struct {
}

// NewPassphrase returns initialized password generator object.
func NewPassphrase() *Passphrase {
	return &Passphrase{}
}

// GeneratePasswords returns list of random passwords.
func (p *Passphrase) GeneratePasswords(dictionary string, passwordsCount uint, wordsCount uint) ([]string, error) {
	if passwordsCount < 1 {
		passwordsCount = 1
	} else if passwordsCount > maxPasswordsCount {
		passwordsCount = maxPasswordsCount
	}

	if wordsCount < 1 {
		wordsCount = 1
	} else if wordsCount > maxWordsCount {
		wordsCount = maxWordsCount
	}

	d, err := p.loadDictionary(dictionary)
	if err != nil {
		return nil, err
	}

	var passwords []string

	r := NewRandIndex()
	for i := uint(0); i < passwordsCount; i++ {
		var words []string
		for j := uint(0); j < wordsCount; j++ {
			ndx := r.RandInt32(uint32(len(d)))
			words = append(words, d[ndx])
		}
		sep := separators[r.RandInt32(uint32(len(separators))-1)]
		passwords = append(passwords, strings.Join(words, string(sep)))
	}

	return passwords, nil
}

func (p *Passphrase) loadDictionary(dictionary string) ([]string, error) {
	file, err := os.Open(dictionary)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
