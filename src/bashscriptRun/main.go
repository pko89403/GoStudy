package main

import "github.com/pko89403/runner"

func main() {
	const scriptName = "./script/script.sh"
	runner.Handle(scriptName, "a", "b", "c")

}
