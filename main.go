package main

import (
	"flag"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/fluhus/gostuff/nlp/wordnet"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	pattern := flag.String("pattern", "noun|verb|number", "Pattern for generating username")
	numberlength := flag.Int("numbermax", 100, "Maximum number")
	cleansymbols := flag.Bool("clean", true, "Remove symbols")
	capitalize := flag.Bool("capitalize", true, "Capitalize each word in user")

	flag.Parse()

	wn, err := wordnet.Parse("wordnet-dict")
	if err != nil {
		fmt.Println("Error parsing WordNet:", err)
		return
	}

	username := generateUsername(*pattern, *numberlength, *cleansymbols, *capitalize, wn)
	fmt.Println("Random Username:", username)
}

func generateUsername(pattern string, numberlength int, cleansymbols bool, capitalize bool, wn *wordnet.WordNet) string {
	parts := strings.Split(pattern, "|")
	var username string

	var usernameParts []string
	for _, partType := range parts {
		var part string

		switch partType {
		case "noun":
			part = getRandomWord(wn, "n")
		case "verb":
			part = getRandomWord(wn, "v")
		case "adjective":
			part = getRandomWord(wn, "a")
		case "adverb":
			part = getRandomWord(wn, "r")
		case "number":
			part = fmt.Sprintf("%d", rand.Intn(numberlength))
		default:
			fmt.Println("Unknown pattern partType:", partType)
		}

		if capitalize {
			part = cases.Title(language.English, cases.Compact).String(part)
		}

		usernameParts = append(usernameParts, part)
	}

	username = strings.Join(usernameParts, "")

	if cleansymbols {
		regex := regexp.MustCompile("[^a-zA-Z0-9]+")
		username = regex.ReplaceAllString(username, "")
	}

	return username
}

func getRandomWord(wn *wordnet.WordNet, pos string) string {
	words := make([]string, 0)
	for word, idList := range wn.Lemma {
		if strings.HasPrefix(word, pos) {
			words = append(words, idList...)
		}
	}
	if len(words) == 0 {
		return ""
	}
	selectedID := words[rand.Intn(len(words))]
	selectedSynset := wn.Synset[selectedID]
	return selectedSynset.Word[0]
}
