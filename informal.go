// A simple command line utility to process text
// returning a more informal version with contractions.
package main

import (
  "fmt"
  "strings"
  "io/ioutil"
  "os"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  contractor := strings.NewReplacer(
    "will not ",  "won't ",
    "can not ",   "can't ",
    "ould not ",  "ouldn't ",
    "do not ",    "don't ",
    "is not ",    "isn't ",
    "are not ",   "aren't ",
    "did not ",   "didn't ",
    "has not ",   "hasn't ",
    "have not ",  "haven't ",
    "were not ",  "weren't ",
    "do not ",    "don't ",
    "does not ",  "doesn't ",
    "will not ",  "won't ",
    "are not ",   "aren't ",
    "can not ",   "can't ",
    "did not ",   "didn't ",
    "has not ",   "hasn't ",
    "Do not ",    "Don't ",
    "Is not ",    "Isn't ",
    "Are not ",   "Aren't ",
    "Did not ",   "Didn't ",
    "Can not ",   "Can't ",
    "Did not ",   "Didn't ",
    "Has not ",   "Hasn't ",
    "Have not ",  "Haven't ",
    "Were not ",  "Weren't ")

  informalizer := strings.NewReplacer(
    "I am ",      "I'm ",
    "he is ",     "he's ",
    "she is ",    "she's ",
    "it is ",     "it's ",
    "we are ",    "we're ",
    "they are ",  "they're ",
    "I will ",    "I'll ",
    "I shall ",   "I'll ",
    "you will ",  "you'll ",
    "he will ",   "he'll ",
    "we will ",   "we'll ",
    "they will ", "they'll ",
    "it will ",   "it'll ",
    "is not ",    "isn't ",
    "you are ",   "you're ",
    "what is ",   "what's ",
    "who is ",    "who's ",
    "that is ",   "that's ",
    "there is ",  "there's ",
    "He is ",     "He's ",
    "She is ",    "She's ",
    "It is ",     "It's ",
    "We are ",    "We're ",
    "They are ",  "They're ",
    "Do not ",    "Don't ",
    "Does not ",  "Doesn't ",
    "Will not ",  "Won't ",
    "Are not ",   "Aren't ",
    "You will ",  "You'll ",
    "He will ",   "He'll ",
    "We will ",   "We'll ",
    "They will ", "They'll ",
    "It will ",   "It'll ",
    "Is not ",    "Isn't ",
    "Let us ",    "Let's ",
    "You are ",   "You're ",
    "What is ",   "What's ",
    "Who is ",    "Who's ",
    "That is ",   "That's ",
    "There is ",  "There's ")

  doc, err := ioutil.ReadFile(os.Args[1])
  check(err)

  contracted := contractor.Replace(string(doc))
  informalized := informalizer.Replace(contracted)

  fmt.Println(informalized)
}
