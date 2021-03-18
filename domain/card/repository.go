package card

// CardRepository provides an abstraction on top of the card data source
type Repository interface {
  CreateCard(*Card) (*Card, error)
  ReadCard(int) (*Card, error)
  ListCards() ([]Card, error)
  RandomCard() (*Card, error)
}
