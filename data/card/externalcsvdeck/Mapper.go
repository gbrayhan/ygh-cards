package externalcsvdeck

import (
  domain "github.com/gbrayhan/academy-go-q12021/domain/card"
)

func ToDataModel(entity *domain.Card) *Card {
  return &Card{
    ID:        entity.ID,
    Name:      entity.Name,
    Type:      entity.Type,
    Level:     entity.Level,
    Race:      entity.Race,
    Attribute: entity.Attribute,
    Atk:       entity.ATK,
    Def:       entity.DEF,
  }
}

func toDomainModel(entity *Card) *domain.Card {
  var image string

  if len(entity.CardImages) > 0 {
    image = entity.CardImages[0].ImageURL
  }

  return &domain.Card{
    ID:        entity.ID,
    Name:      entity.Name,
    Type:      entity.Type,
    Level:     entity.Level,
    Race:      entity.Race,
    Attribute: entity.Attribute,
    ATK:       entity.Atk,
    DEF:       entity.Def,
    Img:       image,
  }
}

