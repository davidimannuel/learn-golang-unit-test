package helper

func HelloWorld(name string) string {
	if name != "" {
		return "hello " + name
	} else {
		return "hello world"
	}
}
