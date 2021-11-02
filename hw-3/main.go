package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/pflag"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sort"
	"strings"
	"sync"
	"syscall"
	"time"
)

const defaultFilePath = "./city.json"
var filePath =defaultFilePath
var citiesList cities
var startTime time.Time


type cities struct {
	cities [] city
}

type city struct {
	cityId int `json:"city_id"`
	countryId int `json:"country_id"`
	regionId int `json:"region_id"`
	name string `json:"name"`
}


type player struct {
	playerName string
	playerScore int
	playerTime time.Duration
	playerPlace int
}


func (c *cities)removeCity(cityName string, mutex *sync.Mutex)  {
	mutex.Lock()
	for i, city:=range c.cities{
		if city.name==cityName{
			c.cities=append(c.cities[:i], c.cities[i+1:]...)
		}
	}
	mutex.Unlock()
}

func readJson() {

	jsonFile, err:=os.Open(filePath)
	if err !=nil{
		log.Fatalf("failed to read path: %v\n", err)
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			log.Fatalf("failed to close: %v\n", err)
		}
	}(jsonFile)
	byteValue,_:=ioutil.ReadAll(jsonFile)
	if err:=json.Unmarshal(byteValue, &citiesList); err!=nil{
		log.Fatalf("failed to unmarshal:%v\n", err)
	}
}

func (receiver *player) playerFinish()  {

	receiver.playerTime=time.Now().Sub(startTime)
	fmt.Printf("Player '%s' finished game in %v, with score %d\n", receiver.playerName,receiver.playerTime,receiver.playerScore)

}

func (receiver *player) playTurn(word chan string, turn chan string, endTurn chan struct{}, endGame chan string,done chan struct{}, mutex *sync.Mutex)  {
	for{
		select {
		case name:=<-turn:
			if name == receiver.playerName{
				for {
					var nxtWord string
					select {
					case inputWord:=<-word:
						inputWord=strings.ToLower(inputWord)
						var lastChar string
						for charNum:=len(inputWord)-1; charNum>=0; charNum--{
							if lastChar!="ь"||lastChar!="ъ"{
								for _,city:= range citiesList.cities{
									if lastChar == city.name[:0]{
										nxtWord=city.name
										break
									}
								}
								if nxtWord !=""{
									citiesList.removeCity(nxtWord,mutex)
									receiver.playerScore++
									word<-nxtWord
									break
								}
							}
						}
						if nxtWord == ""{
							endGame<-receiver.playerName
							receiver.playerFinish()
							return
						}
						endTurn<- struct{}{}
					}
				}
			}
		case <-done:
			receiver.playerFinish()
			return
		}
	}
}

func getGameStatus(p []player)  {

	sort.Slice(p, func(i, j int) bool {
		return p[i].playerScore>p[j].playerScore
	})
	for i,player:= range p{
		player.playerPlace=i+1
	}
	fmt.Printf(
		"| Place | Name  | Score | In-Game |\n" +
			"| ----- | ----- | ----- | ------- |\n")
	for _, player:= range p{
		fmt.Printf("|   %d   |   %s   |   %v   |    %v    |\n", player.playerPlace, player.playerName, player.playerScore, player.playerTime)
	}
}


func playGame(word chan string,nextPlayer chan string,endTurn chan struct{},endGame chan string,done chan struct{}, mutex *sync.Mutex)  {

	rand.Seed(time.Now().UnixNano())
	firstWord:=citiesList.cities[rand.Intn(len(citiesList.cities))].name
	word<-firstWord
	citiesList.removeCity(firstWord, mutex)
	endTurn<- struct{}{}
	for{
		select {
		case <-endTurn:
			nextPlayer<-"next"
		case finishedPlayer:=<-endGame:
			rmplayer(finishedPlayer)
		case <-done:
			fmt.Printf("Game Over\n")
			return
		}
	}
	
}


func graceful(done chan struct{}) {
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case <-s:
			fmt.Println("terminating app")
			close(done)
		case <-done:
			fmt.Println("DONE SIGNAL: graceful")
			return
		}
	}

}


func main(){

	inp1:=pflag.StringSliceP("playersList", "l",nil,"add list of players")
	inp2:=pflag.StringP("jsonpath","p","./city.json", "path to json file")
	pflag.Parse()
	if inp1 == nil {
		pflag.Usage()
		return
	}
	if *inp2 != ""{
		filePath=*inp2
	}

	players:=make([]player,len(*inp1))

	for i,playerName:=range *inp1{
		players[i].playerName=playerName
	}

	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(players), func(i, j int) {
		players[i], players[j] = players[j], players[i]
	})



	word:=make(chan string)
	nextPlayer:=make(chan string)
	endTurn:=make(chan struct{})
	endGame:=make(chan string)
	done:=make(chan struct{})

	readJson()

	startTime=time.Now()

	wg := sync.WaitGroup{}
	var mutex sync.Mutex
	for _,p:=range players{
		wg.Add(1)
		go func(p player) {
			defer wg.Done()
			p.playTurn(word,nextPlayer,endTurn,endGame,done,&mutex)
		}(p)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		playGame(word,nextPlayer,endTurn,endGame,done,&mutex)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		graceful(done)
	}()

	wg.Wait()

	getGameStatus(players)

}
