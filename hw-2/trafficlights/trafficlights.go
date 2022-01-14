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
	circles =3
)

type light struct {
	color string
	duration  int
	input chan struct{}
	output chan struct{}
	done chan struct{}
}

func (l light) makeLight() {
	duration:=l.duration
	ticker := time.NewTicker(time.Second)
	ticker.Stop()

	for {
		select {
		case <-l.done:
			fmt.Printf("%s end working\n",l.color)
			return
		case <-ticker.C:
			fmt.Printf("traffic light: %s: %v seconds left\n", l.color, duration)
			if duration==0	{
				ticker.Stop()
				duration=l.duration
				l.output<- struct{}{}
				continue
			}
			duration--
		case <-l.input:
			ticker.Reset(time.Second)
		}

	}
}

func lightsControl(lights []light,done chan struct{})  {

	var toggleSwitch bool
	cycle:=circles
	for{

		if toggleSwitch {
			toggleSwitch=false
			lights[2].input<- struct {}{}
			<-lights[2].output
		}else {
			if cycle==0{
				close(done)
				fmt.Println("lights control go off")
				return
			}
			cycle--
			toggleSwitch=true
			lights[0].input<- struct {}{}
			<-lights[0].output
		}
		lights[1].input<- struct {}{}
		<-lights[1].output

	}

}

func GoTrafficLights() {
	wg:=sync.WaitGroup{}
	done:=make(chan struct{})
	lights := []light{

		{"red", red,make(chan struct{}),make(chan struct{}),done},
		{"yellow", yellow,make(chan struct{}),make(chan struct{}),done},
		{"green", green,make(chan struct{}),make(chan struct{}),done},
	}
	for _,l:=range lights{
		wg.Add(1)
		go func(l light) {
			defer wg.Done()
			l.makeLight()
		}(l)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		lightsControl(lights,done)
	}()
	wg.Wait()
	fmt.Println("GoTrafficLights go off")
}
