package cards

// CardResponse struct defines response fields
type CardResponse struct {
  ID        int    `json:"id"`
  Name      string `json:"name"`
  Type      string `json:"type"`
  Level     int    `json:"level"`
  Race      string `json:"race"`
  Attribute string `json:"attribute"`
  ATK       int    `json:"atk"`
  DEF       int    `json:"def"`
  Img       string `json:"img,omitempty"`
}

// ListResponse struct defines cards list response structure
type ListResponse struct {
  Data []CardResponse `json:"data"`
}
