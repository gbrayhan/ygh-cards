package externalcsvdeck

type Card struct {
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
  } `json:"card_sets"`
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
  } `json:"card_prices"`
}

