package csvstore

import (
  "errors"
  "io/ioutil"
  "math/rand"
  "strconv"
  "strings"
  "time"

  "github.com/mitchellh/mapstructure"
  "github.com/spf13/viper"

  dataCard "github.com/gbrayhan/academy-go-q12021/data/card"
)

// Exportable errors
var ErrAlreadyExists = errors.New("element already exists")
var ErrNotFound = errors.New("element not found")

var fileCSV FileCSV

type FileCSV struct {
  Name      string
  Error     error
  Structure StructureCSVFile
}

type StructureCSVFile struct {
  ID        int `json:"ID"`
  Name      int `json:"Name"`
  Type      int `json:"Type"`
  Level     int `json:"Level"`
  Race      int `json:"Race"`
  Attribute int `json:"Attribute"`
  ATK       int `json:"ATK"`
  DEF       int `json:"DEF"`
}

func init() {
  viper.SetConfigFile("config.json")
  err := viper.ReadInConfig()
  if err != nil {
    return
  }

  err = mapstructure.Decode(viper.GetStringMap("Databases.CSV.CardsYGO.Structure"), &fileCSV.Structure)
  if err != nil {
    return
  }

  fileCSV.Name = viper.GetString("Databases.CSV.CardsYGO.File")
}

func (f *FileCSV) mapCSVFile() (data map[int]dataCard.Card, err error) {
  data = make(map[int]dataCard.Card)
  contentBytes, err := ioutil.ReadFile(f.Name)
  if err != nil {
    return
  }

  for _, line := range strings.Split(string(contentBytes), "\n") {
    var row []string
    if line != "" {
      row = strings.Split(line, ",")
    }

    if len(row) != 0 {
      var key, errConv = strconv.Atoi(row[f.Structure.ID])
      if errConv != nil {
        continue
      }
      level, _ := strconv.Atoi(row[f.Structure.Level])
      atk, _ := strconv.Atoi(row[f.Structure.ATK])
      def, _ := strconv.Atoi(row[f.Structure.DEF])

      data[key] = dataCard.Card{
        ID:        key,
        Name:      row[f.Structure.Name],
        Type:      row[f.Structure.Type],
        Level:     level,
        Race:      row[f.Structure.Race],
        Attribute: row[f.Structure.Attribute],
        ATK:       atk,
        DEF:       def}
    }
  }
  return
}

func (f *FileCSV) mapKeysExistData() (data map[string]map[string]bool, err error) {
  data = make(map[string]map[string]bool)
  contentBytes, err := ioutil.ReadFile(fileCSV.Name)
  if err != nil {
    return
  }

  for _, line := range strings.Split(string(contentBytes), "\n") {
    var row []string
    if line != "" {
      row = strings.Split(line, ",")
    }

    if len(row) != 0 {
      data["id"][row[fileCSV.Structure.ID]] = true
      data["name"][strings.ToLower(strings.ReplaceAll(row[fileCSV.Structure.Name], " ", ""))] = true
    }
  }

  return
}

func (f *FileCSV) isDuplicate(card dataCard.Card) (isDuplicate bool, err error) {
  dataKeys, _ := f.mapKeysExistData()
  if dataKeys["name"][strings.ToLower(strings.ReplaceAll(card.Name, " ", ""))] {
    isDuplicate = true
    return
  }
  return
}

func (f *FileCSV) SaveCard(card dataCard.Card) (err error) {
  isDuplicate, err := fileCSV.isDuplicate(card)
  if err != nil {
    return
  }
  if isDuplicate {
    err = ErrAlreadyExists
    return
  }







  return
}

// FindCardByID
func (f *FileCSV) FindCardByID(id int) (card dataCard.Card, err error) {
  dataMap, err := fileCSV.mapCSVFile()
  if err != nil {
    return
  }

  if _, ok := dataMap[id]; !ok {
    err = ErrNotFound
    return
  }

  card = dataMap[id]
  return
}

func (f *FileCSV) FindAllCards(cards *[]dataCard.Card, ) (err error) {
  dataMap, err := fileCSV.mapCSVFile()
  if err != nil {
    return
  }

  for _, card := range dataMap {
    *cards = append(*cards, card)
  }

  return
}

func (f *FileCSV) RandCard() (card dataCard.Card, err error) {
  dataMap, err := fileCSV.mapCSVFile()
  if err != nil {
    return
  }

  rand.Seed(time.Now().UnixNano())
  if _, ok := dataMap[rand.Intn(len(dataMap))]; ok {
    card = dataMap[rand.Intn(len(dataMap))]
  }

  return
}
