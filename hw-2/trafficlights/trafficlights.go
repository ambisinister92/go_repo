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
	for i := l.time; i >= 0; i-- {
		select {
		case <-ticker.C:
			fmt.Printf("traffic light: %s: %v seconds left\n", l.color, i)
		}
	}
}

func GoTrafficLights() {
	lights := []light{

		{"red", red},
		{"yellow", yellow},
		{"green", green},
	}
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		for _, l := range lights {
			wg.Add(1)
			go func(l light) {
				l.light()
				wg.Done()
			}(l)
			wg.Wait()
		}
	}

}
