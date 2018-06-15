package main

import "os"
import "fmt"
import "strings"
import "io/ioutil"
import "strconv"

func main() {

  // command arguments
  args := os.Args
  word := args[1]
  maxNumberOfWords := atoiOrPanic(args[2])

  // init state
  ngrams := buildNgramsFromData("data1.txt")
  wordIndex := 0

  // reduce qualifying ngrams until we run out or have reached the max words requested
  for  {

    nextWord, filteredNgrams := filterNgramsForWordAtIndex(ngrams, word, wordIndex)
    if len(filteredNgrams) == 0 || wordIndex + 1 >= maxNumberOfWords {
      break
    }

    word = nextWord
    ngrams = filteredNgrams
    wordIndex += 1
  }

  fmt.Println("Result:", ngrams)
}

func filterNgramsForWordAtIndex(ngrams []Ngram, word string, wordIndex int) (string, []Ngram) {

  matchingNgramsByNextWord := map[string][]Ngram{}
  probabilityByNextWord := map[string]int{}

  // sort ngrams into map of ngram list keyed by next word
  // along the way, keep track of the probability sum by for each word
  for _, ngram := range ngrams {

    words := ngram.Words
    nextWordIndex := wordIndex + 1

    // only consider ngrams that have the correct starting word and have a next word
    if len(words) > nextWordIndex && words[wordIndex] == word {
      nextWord := words[nextWordIndex]
      _, hasValue := matchingNgramsByNextWord[nextWord]
      if !hasValue {
        matchingNgramsByNextWord[nextWord] = []Ngram{}
      }
      matchingNgramsByNextWord[nextWord] = append(matchingNgramsByNextWord[nextWord], ngram)
      probabilityByNextWord[nextWord] = probabilityByNextWord[nextWord] + ngram.MatchCount
    }
  }

  // select ngram list with highest probability
  highestProbability := 0
  highestProbabilityNextWord := ""
  for key, value := range probabilityByNextWord {
    if value > highestProbability {
      highestProbabilityNextWord = key
      highestProbability = value
    }
  }

  return highestProbabilityNextWord, matchingNgramsByNextWord[highestProbabilityNextWord]
}

// parse a text file of data of arbitrary ngram length, followed columns for match_count, page_count, volume_count
func buildNgramsFromData(filename string) []Ngram {

  b, err := ioutil.ReadFile(filename)
  if (err != nil) {
    panic(err)
  }

  dataStr := string(b)
  lines := strings.Split(dataStr, "\n")
  ngrams := []Ngram{}
  for _, line := range lines {

    cols := strings.Fields(line)
    if len(cols) < 5 {
      continue
    }

    matchCount := atoiOrPanic(cols[len(cols) - 3])
    pageCount := atoiOrPanic(cols[len(cols) - 2])
    volumeCount := atoiOrPanic(cols[len(cols) - 1])

    ngram := Ngram{cols[:len(cols) - 4], cols[len(cols) - 4], matchCount, pageCount, volumeCount}
    ngrams = append(ngrams, ngram)
  }

  return ngrams
}

// helper that assumes all strings we parse to int will succeed or we'll panic
func atoiOrPanic(stringValue string) int {

  intValue, error := strconv.Atoi(stringValue)
  if error != nil {
    panic(error)
  }
  return intValue
}
