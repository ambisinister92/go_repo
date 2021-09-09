package main

import (
	"fmt"
	"hw-2/jsontoyaml"
	"hw-2/trafficlights"
)

func main() {

	trafficlights.GoTrafficLights()
	fmt.Println(jsontoyaml.GoJsonToYaml("./lesson.json"))
	fmt.Println("Done!")

}
