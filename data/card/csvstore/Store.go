package csvstore

import (
  "github.com/pkg/errors"

  domainCard "github.com/gbrayhan/academy-go-q12021/domain/card"
  domainErrors "github.com/gbrayhan/academy-go-q12021/domain/errors"
)

const (
  alreadyExistsError = "Error card already exists"
  createError        = "Error in creating new card"
  readError          = "Error in finding card in the database"
  listError          = "Error in getting cards from the database"
)

// Store struct manages interactions with cards store
type Store struct {
  csv       FileCSV
  cardsRepo domainCard.Repository
}

func New() *Store {
  return &Store{
  }
}

//
func (s *Store) CreateCard(cardDom *domainCard.Card) (card *domainCard.Card, err error) {
  cardData := ToDataModel(cardDom)
  err = s.csv.SaveCard(cardData)
  if errors.Is(err, ErrAlreadyExists) {
    err = domainErrors.NewAppErrorWithType(domainErrors.ResourceAlreadyExists)
  }
  card = ToDomainModel(cardData)
  return
}

func (s *Store) ReadCard(id int) (card *domainCard.Card, err error) {
  dCard, err := s.csv.FindCardByID(id)

  if errors.Is(err, ErrNotFound) {
    err = domainErrors.NewAppErrorWithType(domainErrors.NotFound)
    return
  }

  if err != nil {
    err = domainErrors.NewAppError(errors.Wrap(err, readError), domainErrors.RepositoryError)
    return
  }

  card = ToDomainModel(&dCard)
  return
}

func (s *Store) ListCards() (cards []domainCard.Card, err error) {
  var results []Card
  err = s.csv.FindAllCards(&results)

  if err != nil {
    err = domainErrors.NewAppError(errors.Wrap(err, listError), domainErrors.RepositoryError)
    return
  }

  cards = make([]domainCard.Card, len(results))

  for i := range results {
    cards[i] = *ToDomainModel(&results[i])
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

  card = ToDomainModel(&dCard)
  return
}

func (s *Store) ConcurrencyCards(typeQuery string, items int, workers int) (cards []domainCard.Card, err error) {
  cardsSource, err := s.csv.concurrencyReadQuery(typeQuery, items, workers)
  cards = make([]domainCard.Card, len(cardsSource))

  for i := range cardsSource {
    cards[i] = *ToDomainModel(&cardsSource[i])
  }
  return
}
