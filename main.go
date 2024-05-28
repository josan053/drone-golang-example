package main

import "fmt"

func HelloWorld() string {
	return "Trigger the GitHub Actions, traefik workshop!"
}
func main() {
	fmt.Println(HelloWorld())

}

//測試指令  go test -v .
