package main

import "strconv"

type CharType int

const (
	Undefined	CharType = 0
	Number		CharType = 1
	UpperCase 	CharType = 2
	LowerCase	CharType = 3
	WhiteSpace  CharType = 4
	WildCard 	CharType = 5
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
		if (isUpperCase(char)) {
			if (prevChar != UpperCase) {
				regex += addRuleWithLength(prevChar, (i-prevIndex), input[prevIndex:i], openLengthPattern)
				prevIndex = i
			}
			prevChar = UpperCase
		} else if (isLowerCase(char)) {
			if (prevChar != LowerCase) {
				regex += addRuleWithLength(prevChar, (i-prevIndex),input[prevIndex:i], openLengthPattern)
				prevIndex = i
			}
			prevChar = LowerCase
		} else if (isNumber(char)) {
			if (prevChar != Number) {
				regex += addRuleWithLength(prevChar, (i-prevIndex), input[prevIndex:i], openLengthPattern)
				prevIndex = i
			}
			prevChar = Number
		} else if (isWhiteSpace(char)) {
			if (prevChar != WhiteSpace) {
				regex += addRuleWithLength(prevChar, (i-prevIndex), input[prevIndex:i], openLengthPattern)
				prevIndex = i
			}
			prevChar = WhiteSpace
		} else if (isWildCard(char)) {
			if (prevChar != WildCard) {
				regex += addRuleWithLength(prevChar, (i-prevIndex), input[prevIndex:i], openLengthPattern)
				prevIndex = i
			}
			prevChar = WildCard
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
		return "[a-z]" + addFixedPatternLength(length, openLengthPattern)
	} else if (charType == Number) {
		return "[0-9]" + addFixedPatternLength(length, openLengthPattern)
	} else if (charType == WhiteSpace) {
		return "\\s" + addFixedPatternLength(length, openLengthPattern)
	} else if (charType == WildCard) {
		return "." + addFixedPatternLength(length, openLengthPattern)
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

func isUpperCase(charValue rune) bool {
    return (charValue > 64 && charValue < 91)
}

func isLowerCase(charValue rune) bool {
    return (charValue > 96 && charValue < 123)
}

func isNumber(charValue rune) bool {
    return (charValue > 47 && charValue < 58)
}

func isWhiteSpace(charValue rune) bool {
	return charValue == 32
}

func isWildCard(charValue rune) bool {
	return charValue == 42
}
