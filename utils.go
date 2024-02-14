package main

func DirtyIndicator() string {
	if IsDirty {
		return " !"
	}

	return ""
}
