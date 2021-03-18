package externalprodeck

import (
  "github.com/pkg/errors"

  domainCard "github.com/gbrayhan/academy-go-q12021/domain/card"
  domainErrors "github.com/gbrayhan/academy-go-q12021/domain/errors"
)

const (
  createError = "Error in creating new card"
  readError   = "Error in finding card in the database"
  listError   = "Error in getting cards from the database"
)

// Store struct manages interactions with cards store
type Store struct {
  api       ApiProDec
  cardsRepo domainCard.Repository
}

func (s *Store) CreateCard(card *domainCard.Card) (*domainCard.Card, error) {
  panic("implement me")
}

func (s *Store) ReadCard(i int) (*domainCard.Card, error) {
  panic("implement me")
}

func (s *Store) ListCards() (cards []domainCard.Card, err error) {
  cardsAPI, err := AllAPICards()
  if err != nil {
    return
  }

  for _, c := range cardsAPI {
    var image string

    if len(c.CardImages) > 0 {
      image = c.CardImages[0].ImageURL
    }

    cardDomain := domainCard.Card{
      ID:        c.ID,
      Name:      c.Name,
      Type:      c.Type,
      Level:     c.Level,
      Race:      c.Race,
      Attribute: c.Attribute,
      ATK:       c.Atk,
      DEF:       c.Def,
      Img:       image}
    cards = append(cards, cardDomain)

  }

  return
}

func New() *Store {
  return &Store{
  }
}

func (s *Store) RandomCard() (card *domainCard.Card, err error) {
  apiCard, err := RandApiCard()
  if err != nil {
    err = domainErrors.NewAppError(errors.Wrap(err, readError), domainErrors.RepositoryError)
    return
  }

  if apiCard.Name == "" {
    err = domainErrors.NewAppErrorWithType(domainErrors.NotFound)
    return
  }

  card = &domainCard.Card{
    ID:        apiCard.ID,
    Name:      apiCard.Name,
    Type:      apiCard.Type,
    Level:     apiCard.Level,
    Race:      apiCard.Race,
    Attribute: apiCard.Attribute,
    ATK:       apiCard.Atk,
    DEF:       apiCard.Def}

  return
}
