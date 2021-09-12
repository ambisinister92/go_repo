package main

import (
	"fmt"

	"flag"

	"hw-2/jsontoyaml"
	"hw-2/trafficlights"
)

func main() {

	tlights:=flag.Bool("lightsUp", false,  "enable traffic lights")
	var s string
	flag.StringVar(&s,"toYaml","","make yaml from specified json")
	flag.Parse()


	if *tlights{
		trafficlights.GoTrafficLights()
	}
	if s!=""{
		jsontoyaml.GoJsonToYaml(s)
	}

	fmt.Println("Done!")

}
