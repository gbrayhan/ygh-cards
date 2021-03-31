package card

// CardService defines card service behavior.
type CardService interface {
  CreateCard(*Card) (*Card, error)
  ReadCard(id int) (*Card, error)
  ListCards() ([]Card, error)
  RandomCard() (*Card, error)
  ConcurrencyCards(string, int, int) ([]Card, error)
}

// Service struct handles card business logic tasks.
type Service struct {
  Repository Repository
}

func (svc *Service) CreateCard(card *Card) (*Card, error) {
  return svc.Repository.CreateCard(card)
}

func (svc *Service) ReadCard(id int) (*Card, error) {
  return svc.Repository.ReadCard(id)
}

func (svc *Service) ListCards() ([]Card, error) {
  return svc.Repository.ListCards()
}

func (svc *Service) RandomCard() (*Card, error) {
  return svc.Repository.RandomCard()
}

func (svc *Service) ConcurrencyCards(typeQuery string, items int, workers int) ([]Card, error) {
  return svc.Repository.ConcurrencyCards(typeQuery, items, workers)
}

// NewService creates a new service struct
func NewService(repository Repository) *Service {
  return &Service{Repository: repository}
}
