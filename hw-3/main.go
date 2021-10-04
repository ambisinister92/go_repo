package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/pflag"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)


type Player struct {
	place int
	name string
	score int
	inGame time.Duration
}


type cities struct {
	cities [] city
}

type city struct {
	cityId int `json:"city_id"`
	countryId int `json:"country_id"`
	regionId int `json:"region_id"`
	name string `json:"name"`
}

var usedNames []string
const defaultFilePath = "./city.json"
var filePath =defaultFilePath
var citiesList cities

func readJson() {

	jsonFile, err:=os.Open(filePath)
	if err !=nil{
		log.Fatalf("failed to read path: %v\n", err)
	}
	defer jsonFile.Close()
	byteValue,_:=ioutil.ReadAll(jsonFile)
	if err:=json.Unmarshal(byteValue, &citiesList); err!=nil{
		log.Fatalf("failed to unmarshal:%v\n", err)
	}
}


func (p *Player)playGame(input chan string, output chan string,  done chan struct{}) {

	for  {
		select{
		case inputWord:= <-input:
			inputWord=strings.ToLower(inputWord)
			var lastChar string
			charNum:=len(inputWord)-1
			for {
				lastChar=inputWord[:charNum]
				switch lastChar {
				case "ь","ъ","ы","й":
					charNum--
				default:
					break
				}
			}
			for _,city:=range citiesList.cities{

				if lastChar==city.name[:0]{

					for _,usedName:=range usedNames{

						if city.name != usedName{
							usedNames = append(usedNames, city.name)
							p.score++
							output<-city.name
						}else {
							return
						}
					}
				}
			}
			case <-done:

				return

		}
	}
}
func turnOrder()  {

}

func getGameStatus(p []Player)  {

	sort.Slice(p, func(i, j int) bool {
		return p[i].score>p[j].score
	})
	for i,player:= range p{
		player.place=i+1
	}
	fmt.Printf(
		"| Place | Name  | Score | In-Game |\n" +
			"| ----- | ----- | ----- | ------- |\n")
	for _, player:= range p{
		fmt.Printf("|   %d   |   %s   |   %v   |    %v    |\n", player.place, player.name, player.score, player.inGame)
	}
}

func main()  {

	inp1:=pflag.StringSliceP("playerslist", "l",nil,"add list of players")
	inp2:=pflag.StringP("jsonpath","p","./city.json", "path to json file")
	pflag.Parse()
	if *inp1 == nil {
		pflag.Usage()
		return
	}
	if len(*inp2) != 0{
		filePath=*inp2
	}
	order:=*inp1

	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(order), func(i, j int) {
		order[i], order[j] = order[j], order[i]
	})



	done := make(chan struct{})
	word:=make(chan string)//рандомно пишем слово в канал word
	readjson:=make(chan struct{})


	readJson()

	start := time.Now()
	wg:=sync.WaitGroup{}
	wg.Add(len(players))
	for _, player:=range players{
		go func() {
			player.playGame()
			finish := time.Now()
			elapsed := finish.Sub(start)
			player.inGame=elapsed
			fmt.Printf("Player %s finished game with score %v in %v",p.name,p.score,p.inGame)
		}()
	}









}