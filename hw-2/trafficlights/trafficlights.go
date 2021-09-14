package trafficlights

import (
	"fmt"
	"sync"
	"time"
)

const (
	red    = 10
	green  = 7
	yellow = 2
)

type light struct {
	color string
	time  int
}

func (l light) light() {
	ticker := time.NewTicker(1 * time.Second)
	time:=l.time
	for {
		select {
		case <-ticker.C:
			fmt.Printf("traffic light: %s: %v seconds left\n", l.color, time)
		}
		if time==0	{
			return
		}
		time--
	}
}

func GoTrafficLights() {
	lights := []light{

		{"red", red},
		{"yellow", yellow},
		{"green", green},
	}
	next:= make(chan string)
	steady:=make (chan  string)

	cicles:=3
	var wg sync.WaitGroup

	

	for {
		for _, l := range lights {
			wg.Add(1)
			go func(l light) {
				l.light()
				wg.Done()
			}(l)
			wg.Wait()
		}
		if cicles==0{
			return
		}
		cicles--
	}

}
