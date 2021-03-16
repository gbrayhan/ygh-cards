package csvstore

import (
  "github.com/pkg/errors"

  dataCard "github.com/gbrayhan/academy-go-q12021/data/card"
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
  csv       FileCSV
  cardsRepo domainCard.CardRepository
}

func New() *Store {
  return &Store{
  }
}

//
func (s *Store) CreateCard(cardDom *domainCard.Card) (card *domainCard.Card, err error) {
 //cardData := dataCard.ToDataModel(cardDom)
 //dataKeys, _ = s.csv.mapKeysExistData()
 //
 //
 //



 return
}

func (s *Store) ReadCard(id int) (card *domainCard.Card, err error) {
  dCard, err := s.csv.FindCardByID(id)
  if err != nil {
    err = domainErrors.NewAppError(errors.Wrap(err, readError), domainErrors.RepositoryError)
    return
  }

  if dCard.ID == 0 {
    err = domainErrors.NewAppErrorWithType(domainErrors.NotFound)
    return
  }

  card = dataCard.ToDomainModel(&dCard)
  return
}

func (s *Store) ListCards() (cards []domainCard.Card, err error) {
  var results []dataCard.Card
  err = s.csv.FindAllCards(&results)

  if err != nil {
    err = domainErrors.NewAppError(errors.Wrap(err, listError), domainErrors.RepositoryError)
    return
  }

  cards = make([]domainCard.Card, len(results))

  for i := range results {
    cards[i] = *dataCard.ToDomainModel(&results[i])
  }

  return
}

func (s *Store) RandomCard() (card *domainCard.Card, err error) {
  dCard, err := s.csv.RandCard()
  if err != nil {
    err = domainErrors.NewAppError(errors.Wrap(err, readError), domainErrors.RepositoryError)
    return
  }

  if dCard.ID == 0 {
    err = domainErrors.NewAppErrorWithType(domainErrors.NotFound)
    return
  }

  card = dataCard.ToDomainModel(&dCard)
  return
}
