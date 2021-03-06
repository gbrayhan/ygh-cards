package csvstore

import (
  "bufio"
  "errors"
  "fmt"
  "io/ioutil"
  "math/rand"
  "os"
  "strconv"
  "strings"
  "time"

  "github.com/mitchellh/mapstructure"
  "github.com/spf13/viper"
)

// Exportable errors
var ErrAlreadyExists = errors.New("element already exists")
var ErrNotFound = errors.New("element not found")
var LastIdError = errors.New("last id error")

var fileCSV FileCSV

type FileCSV struct {
  Name        string
  Error       error
  Structure   StructureCSVFile
  MapCSVData  map[int]Card
  MapKeysData map[string]map[string]bool
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
  viper.SetConfigFile(os.Getenv("APP_YGH_CONFIG_FILE"))
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

func (f *FileCSV) nextLastID() (nextLastID int, err error) {

  file, err := os.Open(fileCSV.Name)
  if err != nil {
    return
  }
  defer file.Close()

  buf := make([]byte, 62)
  stat, err := os.Stat(fileCSV.Name)
  if err != nil {
    return
  }
  start := stat.Size() - 62
  _, err = file.ReadAt(buf, start)
  if err != nil {
    return
  }
  lastLine := fmt.Sprintf("%s\n", buf)
  lastLine = strings.Trim(lastLine, "\r\n")

  row := strings.Split(lastLine, ",")
  lastID, err := strconv.Atoi(strings.TrimSpace(row[fileCSV.Structure.ID]))

  if err != nil {
    return
  }

  nextLastID = lastID + 1
  err = f.mapKeysExistData()
  if err != nil {
    return
  }

  for i := 0; i <= 1000; i++ {
    if !f.MapKeysData["id"][string(rune(nextLastID))] {
      return
    }
    nextLastID++
  }
  err = LastIdError
  return
}

func (f *FileCSV) mapCSVFile() (err error) {
  if f.MapCSVData != nil {
    return
  }

  f.MapCSVData = make(map[int]Card)

  file, err := os.Open(fileCSV.Name)
  if err != nil {
    return
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()
    var row []string
    if line != "" {
      row = strings.Split(line, ",")
    }

    if len(row) != 0 {
      var key, errConv = strconv.Atoi(row[fileCSV.Structure.ID])
      if errConv != nil {
        continue
      }
      level, _ := strconv.Atoi(row[fileCSV.Structure.Level])
      atk, _ := strconv.Atoi(row[fileCSV.Structure.ATK])
      def, _ := strconv.Atoi(row[fileCSV.Structure.DEF])
      f.MapCSVData[key] = Card{
        ID:        key,
        Name:      row[fileCSV.Structure.Name],
        Type:      row[fileCSV.Structure.Type],
        Level:     level,
        Race:      row[fileCSV.Structure.Race],
        Attribute: row[fileCSV.Structure.Attribute],
        ATK:       atk,
        DEF:       def}
    }

  }

  if err = scanner.Err(); err != nil {
    return err
  }

  return
}

func (f *FileCSV) mapKeysExistData() (err error) {
  if f.MapKeysData != nil {
    return
  }

  f.MapKeysData = make(map[string]map[string]bool)
  contentBytes, err := ioutil.ReadFile(fileCSV.Name)
  if err != nil {
    return
  }

  f.MapKeysData["id"] = make(map[string]bool)
  f.MapKeysData["name"] = make(map[string]bool)

  for _, line := range strings.Split(string(contentBytes), "\n") {
    var row []string
    if line != "" {
      row = strings.Split(line, ",")
    }

    if len(row) != 0 {

      var _, errConv = strconv.Atoi(row[fileCSV.Structure.ID])
      if errConv != nil {
        continue
      }

      f.MapKeysData["id"][row[fileCSV.Structure.ID]] = true
      f.MapKeysData["name"][strings.ToLower(strings.ReplaceAll(row[fileCSV.Structure.Name], " ", ""))] = true
    }
  }

  return
}

func (f *FileCSV) isDuplicate(card *Card) (isDuplicate bool, err error) {
  err = f.mapKeysExistData()
  if err != nil {
    return
  }

  if f.MapKeysData["name"][strings.ToLower(strings.ReplaceAll(card.Name, " ", ""))] {
    isDuplicate = true
    return
  }
  return
}

func (f *FileCSV) SaveCard(card *Card) (err error) {
  isDuplicate, err := f.isDuplicate(card)
  if err != nil {
    return
  }
  if isDuplicate {
    err = ErrAlreadyExists
    return
  }
  card.ID, err = f.nextLastID()

  // Id,Name,Type,Level,Race,Attribute,ATK,DEF
  line := fmt.Sprintf("%d,%s,%s,%d,%s,%s,%d,%d", card.ID, card.Name, card.Type, card.Level, card.Race, card.Attribute, card.ATK, card.DEF)

  file, err := os.OpenFile(fileCSV.Name,
    os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    return
  }

  defer file.Close()
  _, err = file.WriteString(fmt.Sprintf("\n%s", line))

  return
}

// FindCardByID
func (f *FileCSV) FindCardByID(id int) (card Card, err error) {
  err = f.mapCSVFile()
  if err != nil {
    return
  }

  if _, ok := f.MapCSVData[id]; !ok {
    err = ErrNotFound
    return
  }

  card = f.MapCSVData[id]
  return
}

func (f *FileCSV) FindAllCards(cards *[]Card, ) (err error) {
  err = f.mapCSVFile()
  if err != nil {
    return
  }

  for _, card := range f.MapCSVData {
    *cards = append(*cards, card)
  }

  return
}

func (f *FileCSV) RandCard() (card Card, err error) {
  err = f.mapCSVFile()
  if err != nil {
    return
  }

  rand.Seed(time.Now().UnixNano())
  if _, ok := f.MapCSVData[rand.Intn(len(f.MapCSVData))]; ok {
    card = f.MapCSVData[rand.Intn(len(f.MapCSVData))]
  }

  return
}
