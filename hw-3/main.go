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
	"unicode"
)

const defaultFilePath = "./city.json"
var filePath =defaultFilePath
var citiesList cities
var startTime time.Time
var activeWord string


type cities struct {
	cities [] city
}

type city struct {
	CityId    string `json:"city_id"`
	CountryId string    `json:"country_id"`
	RegionId string    `json:"region_id"`
	Name     string `json:"name"`
}


type player struct {
	playerName string
	playerScore int
	playerTime time.Duration
	playerPlace int
	isFinished bool
}


func (c *cities)removeCity(cityName string, mutex *sync.Mutex)  {
	mutex.Lock()
	for i, city:=range c.cities{
		if city.Name ==cityName{
			c.cities=append(c.cities[:i], c.cities[i+1:]...)
			break
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
	if err:=json.Unmarshal(byteValue, &citiesList.cities); err!=nil{
		log.Fatalf("failed to unmarshal:%v\n", err)
	}
}

func findWord(input string) string {
	var lastChar string
	var nxtWord string

	f:= func(c rune) bool{return !unicode.IsLetter(c)}
	inputStringArr := strings.FieldsFunc(strings.ToLower(input), f)
	inputWord:=[]rune(strings.Join(inputStringArr, " "))
	var buffArr []string
	var buffWord []rune

	for charNum:=len(inputWord)-1; charNum>=0; charNum--{
		if lastChar=string(inputWord[charNum]); lastChar!="ь"||lastChar!="ъ"{
			for _,city:= range citiesList.cities{
				buffArr=strings.FieldsFunc(strings.ToLower(city.Name),f)
				buffWord=[]rune(strings.Join(buffArr, " "))
				if lastChar == string(buffWord[0]){
					nxtWord=city.Name
					return nxtWord
				}
			}
		}
	}
	return nxtWord
}

func (receiver *player) playerFinish()  {

	receiver.playerTime=time.Now().Sub(startTime)
	receiver.isFinished=true
	fmt.Printf("Player '%s' finished game in %v, with score %d\n", receiver.playerName,receiver.playerTime,receiver.playerScore)

}

func (receiver *player) playTurn(turn chan string,word chan string,endChan chan struct{},done chan struct{})  {
	for{
		select {
		case <-done:
			receiver.playerFinish()
			return
		case name:=<-turn:
			if name == receiver.playerName{
				var nxtWord string
				inputWord:=<-word
				if inputWord=="" {
					receiver.playerFinish()
					endChan<- struct{}{}
					return
				}

				nxtWord=findWord(inputWord)

				if nxtWord == ""{
					receiver.playerFinish()
					endChan<- struct{}{}
					return
				}
				fmt.Printf("%v says: %s\n",receiver.playerName, nxtWord)
				receiver.playerScore++
				word<- nxtWord
			}else {
				turn <-name
			}

		}
	}
}

func getGameStatus(p *[]player)  {

	sort.Slice(*p, func(i, j int) bool {
		return (*p)[i].playerTime>(*p)[j].playerTime
	})
	sort.Slice(*p, func(i, j int) bool {
		return (*p)[i].playerTime>(*p)[j].playerTime
	})
	for i:= range *p{
		(*p)[i].playerPlace=i+1
	}

	fmt.Printf(
		"| Place | Name  | Score | In-Game |\n" +
			"| ----- | ----- | ----- | ------- |\n")
	for _, player:= range *p{
		fmt.Printf("|   %d   |   %s   |   %v   |    %v    |\n", player.playerPlace, player.playerName, player.playerScore, player.playerTime)
	}
}


func playGame(nextPlayer chan string,word chan string,endChan chan struct{},done chan struct{}, order[]string,mutex *sync.Mutex)  {

	var index int
	rand.Seed(time.Now().UnixNano())
	activeWord=citiesList.cities[rand.Intn(len(citiesList.cities))].Name
	citiesList.removeCity(activeWord, mutex)
	fmt.Printf("Game Start\n")

	for{
		select {
		case  <-done:
			fmt.Printf("Game Over\n")
			return
		default:

			if len(order)==1{
				nextPlayer<-order[0]
				word<-""
				<-endChan
				fmt.Printf("Game Over\n")
				done<- struct{}{}
				return
			}
			if index==len(order){
				index=0
			}
			fmt.Printf("Next player:%s\n",order[index])
			nextPlayer<-order[index]
			fmt.Printf("New City:%s\n",activeWord)
			word<-activeWord
			select {
			case activeWord=<-word:
				citiesList.removeCity(activeWord, mutex)
				index++
			case <-endChan:
				rand.Seed(time.Now().UnixNano())
				activeWord=citiesList.cities[rand.Intn(len(citiesList.cities))].Name
				citiesList.removeCity(activeWord, mutex)
				order=append(order[:index],order[index+1:]...)
			}

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
	if *inp1 == nil {
		pflag.Usage()
		return
	}
	if *inp2 != ""{
		filePath=*inp2
	}

	order:=*inp1
	players:=make([]player,len(*inp1))

	for i,playerName:=range *inp1{
		players[i].playerName=playerName
	}

	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(order), func(i, j int) {
		order[i], order[j] = order[j], order[i]
	})



	nextPlayer:=make(chan string)
	word:=make(chan string)
	endChan:=make(chan struct{})
	done:=make(chan struct{})

	readJson()

	startTime=time.Now()

	wg := sync.WaitGroup{}
	var mutex sync.Mutex
	for i:=range players{
		wg.Add(1)
		go func(p *player) {
			defer wg.Done()
			p.playTurn(nextPlayer,word,endChan,done)
		}(&players[i])
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		playGame(nextPlayer,word,endChan,done,order,&mutex)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		graceful(done)
	}()

	wg.Wait()


	getGameStatus(&players)


}
