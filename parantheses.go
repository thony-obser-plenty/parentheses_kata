package main

import "fmt"

func main() {
	fmt.Println(isValid("{test(test[2])}"))

	fmt.Println(isValid("()"))     // Output: true
	fmt.Println(isValid("()[]{}")) // Output: true
	fmt.Println(isValid("(]"))     // Output: false
	fmt.Println(isValid("([)]"))   // Output: false
	fmt.Println(isValid("{[]}"))   // Output: true
}

func isValid(s string) bool {
	parenthesesMap := &map[string]string{"(": ")", "{": "}", "[": "]"}
	parenthesesStack := &[]string{}

	for _, char := range s {
		for startingParenthesis, closingParenthesis := range *parenthesesMap {
			unfittingMatch := CheckForFittingMatches(char, closingParenthesis, parenthesesStack, parenthesesMap)

			if unfittingMatch {
				return false
			}

			CheckForNewParentheses(char, startingParenthesis, parenthesesStack)
		}
	}

	if len(*parenthesesStack) > 0 {
		return false
	}

	return true
}

func CheckForFittingMatches(char int32, closingParenthesis string, parenthesisStack *[]string, parenthesisMap *map[string]string) bool {
	parenthesesStackIsNotEmpty := len(*parenthesisStack) > 0
	unfittingMatch := false

	if string(char) == closingParenthesis && parenthesesStackIsNotEmpty {
		stack := *parenthesisStack
		lastParenthesis := stack[len(stack)-1]
		matchingParenthesis := (*parenthesisMap)[lastParenthesis]

		if string(char) == matchingParenthesis {
			*parenthesisStack = stack[:len(stack)-1]
		} else {
			unfittingMatch = true
		}
	}

	return unfittingMatch
}

func CheckForNewParentheses(char int32, startingParenthesis string, parenthesesStack *[]string) {
	if string(char) == startingParenthesis {
		*parenthesesStack = append(*parenthesesStack, startingParenthesis)
	}
}
