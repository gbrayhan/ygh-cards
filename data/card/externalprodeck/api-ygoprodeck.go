package externalprodeck

import (
  "context"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  "time"

  "github.com/spf13/viper"
)


var apiPro ApiProDec

// FileCSV
type ApiProDec struct {
  Url   string
  Error error
}

func init() {
  viper.SetConfigFile(os.Getenv("APP_YGH_CONFIG_FILE"))
  err := viper.ReadInConfig()
  if err != nil {
    panic(err.Error)
    return
  }

  apiPro.Url = viper.GetString("Endpoints.YGOProDeck")
}

func RandApiCard() (apiCard Card, err error) {
  request, err := http.NewRequest("GET", fmt.Sprintf("%s/randomcard.php", apiPro.Url), nil)
  if err != nil {
    return
  }

  ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(900))
  defer cancel()
  request = request.WithContext(ctx)

  client := &http.Client{}

  resp, err := client.Do(request)
  if err != nil {
    return
  }
  defer resp.Body.Close()

  if resp.StatusCode == http.StatusOK {
    bodyBytes, errRead := ioutil.ReadAll(resp.Body)
    if errRead != nil {
      err = errRead
      return
    }
    if errJSON := json.Unmarshal(bodyBytes, &apiCard); errJSON != nil {
      err = errJSON
      return
    }
  }

  return
}

// AllAPICards
func AllAPICards() (apiCards []Card, err error) {
  viper.SetConfigFile(os.Getenv("APP_YGH_CONFIG_FILE"))
  err = viper.ReadInConfig()
  if err != nil {
    return
  }

  apiUrl := viper.GetString("Endpoints.YGOProDeck")
  request, err := http.NewRequest("GET", fmt.Sprintf("%s/cardinfo.php", apiUrl), nil)
  if err != nil {
    return
  }
  ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(9000))
  defer cancel()
  request = request.WithContext(ctx)

  client := &http.Client{}

  resp, err := client.Do(request)
  if err != nil {
    return
  }
  defer resp.Body.Close()
  if resp.StatusCode == http.StatusOK {
    response := struct {
      Data []Card `json:"data"`
    }{}

    bodyBytes, errRead := ioutil.ReadAll(resp.Body)
    if errRead != nil {
      err = errRead
      return
    }
    if errJSON := json.Unmarshal(bodyBytes, &response); errJSON != nil {
      err = errJSON
      return
    }
    apiCards = response.Data
  }
  return
}
