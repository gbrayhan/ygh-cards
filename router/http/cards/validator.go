package cards

// CardValidator is a struct used to validate the JSON payload representing a card.
type CardValidator struct {
  Name string `binding:"required" json:"name"`
  Type string `binding:"required" json:"type"`
}

type SaveCardValidator struct {
  ID        int    `json:"id"`
  Name      string `json:"name"`
  Type      string `json:"type"`
  Level     int    `json:"level"`
  Race      string `json:"race"`
  Attribute string `json:"attribute"`
  Atk       int    `json:"atk"`
  Def       int    `json:"def"`
}


