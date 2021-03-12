package csvstore

import (
  "io/ioutil"
  "strconv"
  "strings"

  "github.com/mitchellh/mapstructure"
  "github.com/spf13/viper"

  dataCard "github.com/gbrayhan/academy-go-q12021/data/card"
)

// FileCSV
type FileCSV struct {
  Error error
}

type StructureCSV struct {
  ID        int `json:"ID"`
  Name      int `json:"Name"`
  Type      int `json:"Type"`
  Level     int `json:"Level"`
  Race      int `json:"Race"`
  Attribute int `json:"Attribute"`
  ATK       int `json:"ATK"`
  DEF       int `json:"DEF"`
}

func (FileCSV) mapCSVFile() (data map[int]dataCard.Card, err error) {
  var structureCSVCards StructureCSV
  err = mapstructure.Decode(viper.GetStringMap("Databases.CSV.CardsYGO.Structure"), &structureCSVCards)
  if err != nil {
    return
  }

  data = make(map[int]dataCard.Card)

  viper.SetConfigFile("config.json")
  err = viper.ReadInConfig()
  if err != nil {
    return
  }

  fileName := viper.GetString("Databases.CSV.CardsYGO.File")
  contentBytes, err := ioutil.ReadFile(fileName)
  if err != nil {
    return
  }

  for _, line := range strings.Split(string(contentBytes), "\n") {
    var row []string
    if line != "" {
      row = strings.Split(line, ",")
    }

    if len(row) != 0 {
      var key, errConv = strconv.Atoi(row[structureCSVCards.ID])
      if errConv != nil {
        continue
      }
      level, _ := strconv.Atoi(row[structureCSVCards.Level])
      atk, _ := strconv.Atoi(row[structureCSVCards.ATK])
      def, _ := strconv.Atoi(row[structureCSVCards.DEF])

      data[key] = dataCard.Card{
        ID:        key,
        Name:      row[structureCSVCards.Name],
        Type:      row[structureCSVCards.Type],
        Level:     level,
        Race:      row[structureCSVCards.Race],
        Attribute: row[structureCSVCards.Attribute],
        ATK:       atk,
        DEF:       def}
    }
  }
  return
}

func (f *FileCSV) FindCardByID(id int) (card dataCard.Card, err error) {
  dataMap, err := f.mapCSVFile()
  if err != nil {
    return
  }
  card = dataMap[id]
  return
}

func (f *FileCSV) FindAllCards(cards *[]dataCard.Card, ) (err error) {
  dataMap, err := f.mapCSVFile()
  if err != nil {
    return
  }

  for _, card := range dataMap {
    *cards = append(*cards, card)
  }

  return
}
