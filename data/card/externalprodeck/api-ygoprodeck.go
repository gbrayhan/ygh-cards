package externalprodeck

import (
  "context"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "time"

  "github.com/spf13/viper"
)

type CardProDeckApi struct {
  ID        int    `json:"id"`
  Name      string `json:"name"`
  Type      string `json:"type"`
  Desc      string `json:"desc"`
  Atk       int    `json:"atk"`
  Def       int    `json:"def"`
  Level     int    `json:"level"`
  Race      string `json:"race"`
  Attribute string `json:"attribute"`
  Archetype string `json:"archetype"`
  CardSets  []struct {
    SetName       string `json:"set_name"`
    SetCode       string `json:"set_code"`
    SetRarity     string `json:"set_rarity"`
    SetRarityCode string `json:"set_rarity_code"`
    SetPrice      string `json:"set_price"`
  } `json:"card_sets,omitempty"`
  CardImages []struct {
    ID            int    `json:"id"`
    ImageURL      string `json:"image_url"`
    ImageURLSmall string `json:"image_url_small"`
  } `json:"card_images"`
  CardPrices []struct {
    CardmarketPrice   string `json:"cardmarket_price"`
    TcgplayerPrice    string `json:"tcgplayer_price"`
    EbayPrice         string `json:"ebay_price"`
    AmazonPrice       string `json:"amazon_price"`
    CoolstuffincPrice string `json:"coolstuffinc_price"`
  } `json:"card_prices,omitempty"`
}

// FileCSV
type ApiProDec struct {
  Url   string
  Error error
}

func RandApiCard() (apiCard CardProDeckApi, err error) {
  viper.SetConfigFile("config.json")
  err = viper.ReadInConfig()
  if err != nil {
    return
  }

  apiUrl := viper.GetString("Endpoints.YGOProDeck")
  request, err := http.NewRequest("GET", fmt.Sprintf("%s/randomcard.php", apiUrl), nil)
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
func AllAPICards() (apiCards []CardProDeckApi, err error) {
  viper.SetConfigFile("config.json")
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
      Data []CardProDeckApi `json:"data"`
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
