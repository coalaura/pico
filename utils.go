package main

func DirtyIndicator() string {
	if IsDirty {
		return " !"
	}

	return ""
}

func InvertColors(str string) string {
	return "\033[30;47m" + str + "\033[0m"
}

func ReplaceCharacterAt(in, repl string, x int) string {
	runes := []rune(in)

	before := runes[:x]

	after := []rune{}

	if x+1 < len(runes) {
		after = runes[x+1:]
	}

	return string(before) + repl + string(after)
}
