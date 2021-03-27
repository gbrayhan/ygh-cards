package card

import (
  "reflect"
  "testing"

  cardsStore "github.com/gbrayhan/academy-go-q12021/data/card/csvstore"
  "github.com/gbrayhan/academy-go-q12021/domain/card"
)

func TestService_CreateCard(t *testing.T) {
  cardsRepoCSV := cardsStore.New()
  cardsSvc := card.NewService(cardsRepoCSV)

  type fields struct {
    repository card.Repository
  }
  type args struct {
    card *card.Card
  }
  tests := []struct {
    name    string
    fields  fields
    args    args
    want    *card.Card
    wantErr bool
  }{
    {
      name:   "csv test success",
      fields: fields{repository: cardsSvc},
      args: args{
        card: &card.Card{
          Name:      "test",
          Type:      "test",
          Level:     777,
          Race:      "test",
          Attribute: "test",
          ATK:       777,
          DEF:       777,
          Img:       "some",
        },
      },
      want:    nil,
      wantErr: false,
    },
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      svc := &card.Service{
        Repository: tt.fields.repository,
      }
      got, err := svc.CreateCard(tt.args.card)
      if (err != nil) != tt.wantErr {
        t.Errorf("Service.CreateCard() error = %v, wantErr %v", err, tt.wantErr)
        return
      }
      if !reflect.DeepEqual(got, tt.want) {
        t.Errorf("Service.CreateCard() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestService_ReadCard(t *testing.T) {
  type fields struct {
    repository card.Repository
  }
  type args struct {
    id int
  }
  tests := []struct {
    name    string
    fields  fields
    args    args
    want    *card.Card
    wantErr bool
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      svc := &card.Service{
        Repository: tt.fields.repository,
      }
      got, err := svc.ReadCard(tt.args.id)
      if (err != nil) != tt.wantErr {
        t.Errorf("Service.ReadCard() error = %v, wantErr %v", err, tt.wantErr)
        return
      }
      if !reflect.DeepEqual(got, tt.want) {
        t.Errorf("Service.ReadCard() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestService_ListCards(t *testing.T) {
  type fields struct {
    repository card.Repository
  }
  tests := []struct {
    name    string
    fields  fields
    want    []card.Card
    wantErr bool
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      svc := &card.Service{
        Repository: tt.fields.repository,
      }
      got, err := svc.ListCards()
      if (err != nil) != tt.wantErr {
        t.Errorf("Service.ListCards() error = %v, wantErr %v", err, tt.wantErr)
        return
      }
      if !reflect.DeepEqual(got, tt.want) {
        t.Errorf("Service.ListCards() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestService_RandomCard(t *testing.T) {
  type fields struct {
    repository card.Repository
  }
  tests := []struct {
    name    string
    fields  fields
    want    *card.Card
    wantErr bool
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      svc := &card.Service{
        Repository: tt.fields.repository,
      }
      got, err := svc.RandomCard()
      if (err != nil) != tt.wantErr {
        t.Errorf("Service.RandomCard() error = %v, wantErr %v", err, tt.wantErr)
        return
      }
      if !reflect.DeepEqual(got, tt.want) {
        t.Errorf("Service.RandomCard() = %v, want %v", got, tt.want)
      }
    })
  }
}
