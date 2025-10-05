package main

var justString string

func createHugeString(size int) string {
	return string(make([]byte, size))
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]rune(v)[:100])
}

func main() {
	someFunc()
}
