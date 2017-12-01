// Author: Garry Cummins
package main

import (
    "fmt"
    "math/rand"
    "time"
    // Regular expressions import 
    "regexp"
)
func ElizaResponse(input string) string {
    // Q2
    //[Ff] matches every father inputted, (?i). for flag, \b before and after 
    // father to check for grandfather 
    if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
        return "Why don’t you tell me more about your father?"
    }

  
    responses := []string{
      "I'm not sure what you’re trying to say. Could you explain it to me?",
      "How does that make you feel?",
      "Why do you say that?",
      "Could you go into more detail?",
      "I see",
      "You've put alot of thought into this",
    }
    return responses[rand.Intn(len(responses))]
}// Function

func main(){
    rand.Seed(time.Now().UTC().UnixNano())
    fmt.Println("Input: " + " People say I look like both my mother and father.")
    fmt.Println("Output: " + ElizaResponse(" People say I look like both my mother and father."))
    fmt.Println()

    fmt.Println("Input: " + " Father was a teacher.")
    fmt.Println("Output: " + ElizaResponse(" Father was a teacher."))
    fmt.Println()

    fmt.Println("Input: " + " I was my father’s favourite.")
    fmt.Println("Output: " + ElizaResponse(" I was my father’s favourite."))
    fmt.Println()

    fmt.Println("Input: " + " I’m looking forward to the weekend.")
    fmt.Println("Output: " + ElizaResponse(" I’m looking forward to the weekend."))
    fmt.Println()

    fmt.Println("Input: " + " My grandfather was French!")
    fmt.Println("Output: " + ElizaResponse(" My grandfather was French!"))
    fmt.Println()

    // add I am inputs
} // Main