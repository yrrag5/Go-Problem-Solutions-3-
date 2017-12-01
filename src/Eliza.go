// Author: Garry Cummins

// Resource: https://gist.github.com/ianmcloughlin/c4c2b8dc586d06943f54b75d9e2250fe
package main

import (
    "fmt"
    "math/rand"
    "time"
    // Regular Expressions import
    "regexp"
    // String import
    "strings"
)
func ElizaResponse(input string) string {
    // Q2
    //[Ff] matches every father inputted, (?i). for flag, \b before and after 
    // father to check for grandfather 
    if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
        return "Why don’t you tell me more about your father?"
    }

    // Q3
    // [] - outside, * - none of these chars, any char, outside ? - 0 or 1
    re:= regexp.MustCompile(`I am ([^.!?]*)[.!?]?`)
    if matched := re.MatchString(input); matched {
        return re.ReplaceAllString(input, "How do you know you are $1?")
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

// Q4 Reflection

func reflection(input string) string {

    // List of reflections
    pronouns := [][]string {
        {`am`, `are`},
        {`my`, `your`},
        {`your`, `my`},
        {`I`, `you`},
        {`you`, `I`},
        {`me`, `you`},
        
    }

    boundaries := regexp.MustCompile(`\b`)

    values := boundaries.Split(input, -1)

    // Loops through tokens until reflect finds a match
    for i, token := range values {
		for _, reflection := range pronouns {
			if matched, _ := regexp.MatchString(reflection[0], token); matched {
				values[i] = reflection[1]
				break
			}
		}
    }

    join := strings.Join(values, ``)
    // Q5 Reflect pronouns
    reResp := []string {
        "Why do ",
        "How do you know that ",
        "I find it interesting that ",
    }
    return(reResp[rand.Intn(len(reResp))] + join)
}

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

    fmt.Println("Input: " + " I'm looking forward to the weekend.")
    fmt.Println("Output: " + ElizaResponse(" I’m looking forward to the weekend."))
    fmt.Println()

    fmt.Println("Input: " + " My grandfather was French!")
    fmt.Println("Output: " + ElizaResponse(" My grandfather was French!"))
    fmt.Println()

    // Q3
    //  I am inputs and outputs

    fmt.Println("Input: " + " I am happy")
    fmt.Println("Output: " + ElizaResponse("I am happy"))
    fmt.Println()

    fmt.Println("Input: " + " I am not happy with your responses.")
    fmt.Println("Output: " + ElizaResponse("I am not happy with your responses."))
    fmt.Println()

    
    fmt.Println("Input: " + " I am not sure that you understand the effect that your questions are having on me.")
    fmt.Println("Output: " + ElizaResponse("I am not sure that you understand the effect that your questions are having on me."))
    fmt.Println()

    
    fmt.Println("Input: " + " I am supposed to just take what you're saying at face value?")
    fmt.Println("Output: " + ElizaResponse("I am supposed to just take what your saying at face value?"))
    fmt.Println()

     // Q4 Reflection
    fmt.Println("Input: " + " I am not happy with your responses.")
    fmt.Println("Output: " + reflection("I am not happy with your responses."))
    fmt.Println()


} // Main