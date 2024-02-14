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
