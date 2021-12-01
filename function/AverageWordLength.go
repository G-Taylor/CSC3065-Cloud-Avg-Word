package function

import (
	"log"
	"regexp"
	"strings"
)

func AverageWordLength(sentence string) float32 {
	var cleanedSentence []string
	var characters int
	var wordCount float32
	var averageWordLength float32

	// set up the regular expression
	reg, err := regexp.Compile("[^a-zA-Z]+")

	if err != nil {
		log.Fatal(err)
	}

	// split the sentence into separate words based on whitespace
	s := strings.Fields(sentence)

	// loop through each word in the sentence
	for i, word := range s {
		log.Println(i, word, len(word))
		// remove special characters and numbers using regex
		cleanedWord := reg.ReplaceAllString(word, "")

		// if the cleaned word is not an empty string after removing numbers and special characters, add it to the cleanedSentence array
		if cleanedWord != "" {
			cleanedSentence = append(cleanedSentence, cleanedWord)
			// increment character count and word count
			characters += len(cleanedWord)
			wordCount += 1
		}
	}

	// get average word length and out put it as a float
	averageWordLength = float32(characters) / float32(wordCount)

	return averageWordLength
}
