package wordlist

import (
	"bufio"
	"os"
	"strings"
)

// WordList contains a list of words and the supported list operations.
type WordList interface {
	// AddWord persists a new word to the existing list.
	AddWord(word string) (err error)

	// GetWords returns all of the words in the existing list.
	GetWords() (words []string, err error)
}

// TODO: revisit this to see if there is a more efficient data structure to use for
// the word list.
type wordListImpl struct {
	filename string
	words    []string
}

// New instantiates a new WordList.
func New(filename string) WordList {
	return &wordListImpl{filename: filename}
}

// AddWord persists a new word to the existing list.
func (w *wordListImpl) AddWord(word string) (err error) {
	var f *os.File
	if f, err = os.OpenFile(w.filename, os.O_APPEND, 0644); err != nil {
		return
	}
	defer f.Close()

	if _, err = f.Write([]byte(word)); err != nil {
		return
	}

	return
}

// TODO: we should cache the words in memory to avoid
// GetWords returns all of the words in the existing list.
func (w *wordListImpl) GetWords() (words []string, err error) {
	var f *os.File
	if f, err = os.Open(w.filename); err != nil {
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)

	// TODO: there is the case where a word that is being checked for existince
	// is outside of the bounds. we should handle this case.
	// for now we will hack in a limit thats the same size of the word list.
	// there is also an issue when doing adds as that increases the number of words.

	// TODO: use len of (words)
	// TOOD: there is a bug with bounds
	for i := 0; i < 100000; i++ {
		var s string
		if s, err = r.ReadString('\n'); err != nil {
			return
		}

		words = append(words, strings.TrimSpace(s))
	}

	return
}

func WordExists(word string, words []string) (exists bool, err error) {
	wordLower := strings.ToLower(word)

	for _, w := range words {
		if strings.ToLower(w) == wordLower {
			return true, nil
		}
	}

	return false, nil
}
