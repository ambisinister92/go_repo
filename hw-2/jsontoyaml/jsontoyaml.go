package jsontoyaml

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"

	"encoding/json"
)
type (
	Sweets struct {
		SweetsId string `json:"id" yaml:"id"`
		SweetsType string `json:"type" yaml:"type"`
		SweetsName     string    `json:"name" yaml:"name"`
		Ppu float64 `json:"ppu" yaml:"ppu"`
		Batters     Batters  `json:"batters" yaml:"batters"`
		Topping []Topping `json:"topping" yaml:"topping"`

	}

	Batters struct{
		Batter []Batter `json:"batter" yaml:"batter"`
	}

	Batter struct {
		BatterId string `json:"id" yaml:"id"`
		BatterType      string `json:"type" yaml:"type"`
	}
	Topping struct {
		ToppingId string `json:"id" yaml:"id"`
		ToppingType string `json:"type" yaml:"type"`
	}
)


func GoJsonToYaml(s string) string{


	jsonFile, err:=os.Open(s)
	if err !=nil{
		log.Fatalf("failed to read: %v\n", err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var sw1 Sweets

	if err := json.Unmarshal(byteValue, &sw1); err != nil {
		log.Fatalf("failed to unmarshal: %v\n", err)
	}
	

	enc, err := yaml.Marshal(sw1)
	if err != nil {
		log.Fatalf("failed to marshal: %v\n", err)
	}
	yamlFile, _ := os.Create("lesson.yaml")
	yamlFile.Write(enc)
	yamlFile.Close()
	return string(enc)

}
