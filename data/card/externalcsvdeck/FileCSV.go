package externalcsvdeck

import (
  b64 "encoding/base64"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "os"
  "strconv"
  "strings"

  "github.com/mitchellh/mapstructure"
  "github.com/spf13/viper"
)

// FileCSV
type FileCSV struct {
  Error error
}


type StructureCSVFile struct {
  ID          int `json:"ID"`
  Name        int `json:"Name"`
  Type        int `json:"Type"`
  Description int `json:"Description"`
  ATK         int `json:"ATK"`
  DEF         int `json:"DEF"`
  Level       int `json:"Level"`
  Race        int `json:"Race"`
  Attribute   int `json:"Attribute"`
  Archetype   int `json:"Archetype"`
  CardSets    int `json:"card_sets"`
  CardImages  int `json:"card_images"`
  CardPrices  int `json:"card_prices"`
}

// MapCSVFile
func MapCSVFile() (data map[int]Card, err error) {
  var structureCSVCards StructureCSVFile
  viper.SetConfigFile(os.Getenv("APP_YGH_CONFIG_FILE"))
  err = viper.ReadInConfig()
  if err != nil {
    return
  }

  err = mapstructure.Decode(viper.GetStringMap("Databases.CSV.ProDeckYGO.Structure"), &structureCSVCards)
  if err != nil {
    return
  }

  data = make(map[int]Card)
  fileName := viper.GetString("Databases.CSV.ProDeckYGO.File")
  contentBytes, err := ioutil.ReadFile(fileName)
  if err != nil {
    return
  }

  for _, line := range strings.Split(string(contentBytes), "\n") {
    var row []string
    if line != "" {
      row = strings.Split(line, "|")
    }

    if len(row) != 0 {
      var key, errConv = strconv.Atoi(row[structureCSVCards.ID])
      if errConv != nil {
        continue
      }
      var currentCard Card

      descDec, _ := b64.StdEncoding.DecodeString(row[structureCSVCards.Description])
      level, _ := strconv.Atoi(row[structureCSVCards.Level])
      atk, _ := strconv.Atoi(row[structureCSVCards.ATK])
      def, _ := strconv.Atoi(row[structureCSVCards.DEF])
      archetype, _ := b64.StdEncoding.DecodeString(row[structureCSVCards.Archetype])
      cardSetsDec, _ := b64.StdEncoding.DecodeString(row[structureCSVCards.CardSets])
      cardImagesDec, _ := b64.StdEncoding.DecodeString(row[structureCSVCards.CardImages])

      currentCard = Card{
        ID:        key,
        Name:      row[structureCSVCards.Name],
        Type:      row[structureCSVCards.Type],
        Desc:      string(descDec),
        Atk:       atk,
        Def:       def,
        Level:     level,
        Race:      row[structureCSVCards.Race],
        Attribute: row[structureCSVCards.Attribute],
        Archetype: string(archetype),
      }

      if errJSON := json.Unmarshal(cardSetsDec, &currentCard.CardSets); errJSON != nil {
        err = errJSON
        return
      }
      if errJSON := json.Unmarshal(cardImagesDec, &currentCard.CardImages); errJSON != nil {
        err = errJSON
        return
      }

      data[key] = currentCard
    }
  }
  return
}

func mapKeysCSV() (data map[string]map[string]bool, err error) {

  return
}

//GetAllCards
func GetAllCards() (cards []Card, err error) {
  mapData, err := MapCSVFile()
  if err != nil {
    return
  }
  for _, v := range mapData {
    cards = append(cards, v)
  }
  return
}

// Save
func (card *Card) Save() (err error) {
  viper.SetConfigFile(os.Getenv("APP_YGH_CONFIG_FILE"))
  err = viper.ReadInConfig()
  if err != nil {
    return
  }

  dataMap, err := MapCSVFile()
  if err != nil {
    return
  }

  for _, v := range dataMap {
    if v.ID == card.ID {
      return
    }

    if v.Name == card.Name {
      return
    }
  }

  descriptionEnc := b64.StdEncoding.EncodeToString([]byte(card.Desc))
  archetypeEnc := b64.StdEncoding.EncodeToString([]byte(card.Archetype))

  cardSets, err := json.Marshal(card.CardSets)
  if err != nil {
    return
  }
  cardSetsEnc := b64.StdEncoding.EncodeToString(cardSets)

  cardImages, err := json.Marshal(card.CardImages)
  if err != nil {
    return
  }
  cardImagesEnc := b64.StdEncoding.EncodeToString(cardImages)

  cardPrices, err := json.Marshal(card.CardPrices)
  if err != nil {
    return
  }
  cardPricesEnc := b64.StdEncoding.EncodeToString(cardPrices)

  // id_api|name|type|desc_enc|atk|def|level|race|attribute|archetype_enc|card_sets_enc|card_images_enc|card_prices_enc
  rowCsvString := fmt.Sprintf("%d|%s|%s|%s|%d|%d|%d|%s|%s|%s|%s|%s|%s",
    card.ID, card.Name, card.Type, descriptionEnc, card.Atk, card.Def, card.Level, card.Race, card.Attribute, archetypeEnc, cardSetsEnc, cardImagesEnc, cardPricesEnc)

  f, err := os.OpenFile(viper.GetString("Databases.CSV.ProDeckYGO.File"),
    os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    log.Println(err)
  }

  defer f.Close()
  _, err = f.WriteString(fmt.Sprintf("%s\n", rowCsvString))

  return
}
