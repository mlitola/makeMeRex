package main

import "strconv"

type CharType int

const (
	Number		CharType = 0
	UpperCase 	CharType = 1
	LowerCase	CharType = 2
	Undefined	CharType = 3
)

func AnalyzeAndMakePattern(input string, fixedLength *bool, onePerLine *bool) string {
	var prevChar CharType = Undefined
	var prevIndex int = 0
	var regex string = ""
	var openLengthPattern = "{1,}"

	if *fixedLength {
		openLengthPattern = ""
	}

	for i, char := range input {		
		if (isCharAUpperCase(char)) {
			if (prevChar != UpperCase) {
				regex += addRuleWithLength(prevChar, (i-prevIndex), input[prevIndex:i], openLengthPattern)
				prevIndex = i
			}
			prevChar = UpperCase
		} else if (isCharALowerCase(char)) {
			if (prevChar != LowerCase) {
				regex += addRuleWithLength(prevChar, (i-prevIndex),input[prevIndex:i], openLengthPattern)
				prevIndex = i
			}
			prevChar = LowerCase
		} else if (isCharANumber(char)) {
			if (prevChar != Number) {
				regex += addRuleWithLength(prevChar, (i-prevIndex), input[prevIndex:i], openLengthPattern)
				prevIndex = i
			}
			prevChar = Number
		} else {
			if (prevChar != Undefined) {
				regex += addRuleWithLength(prevChar, (i-prevIndex), input[prevIndex:i], openLengthPattern)
				prevIndex = i
			}
			prevChar = Undefined
		}
	}

	if (prevIndex < len(input)) {
		i := len(input)
		regex += addRuleWithLength(prevChar, (i-prevIndex), input[prevIndex:i], openLengthPattern)
	}

    if *onePerLine {
        regex = addStartAndEndOfLine(regex)
	}

    return regex
}

func addRuleWithLength(charType CharType, length int, charValues string, openLengthPattern string) string {
	if (charType == UpperCase) {
		return "[A-Z]" + addFixedPatternLength(length, openLengthPattern)
	} else if (charType == LowerCase) {
		return "[a-z]{" + addFixedPatternLength(length, openLengthPattern)
	} else if (charType == Number) {
		return "[0-9]{" + addFixedPatternLength(length, openLengthPattern)
	}
	return charValues
}

func addFixedPatternLength(length int, openLengthPattern string) string {
	if len(openLengthPattern) > 0 {
		return openLengthPattern
	}
    return "{" + strconv.Itoa(length) + "}"
}

func addStartAndEndOfLine(input string) string {
    return "^" + input + "$"
}

func isCharAUpperCase(charValue rune) bool {
    return (charValue > 64 && charValue < 91)
}

func isCharALowerCase(charValue rune) bool {
    return (charValue > 96 && charValue < 123)
}

func isCharANumber(charValue rune) bool {
    return (charValue > 47 && charValue < 58)
}
