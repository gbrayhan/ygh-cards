package cards

import (
  "net/http"
  "strconv"

  "github.com/gin-gonic/gin"

  "github.com/gbrayhan/academy-go-q12021/domain/card"
  domainErrors "github.com/gbrayhan/academy-go-q12021/domain/errors"
)

// NewRoutesFactory create and returns a factory to create routes for the card
func NewRoutesFactory(group *gin.RouterGroup) func(service card.CardService) {
  cardRoutesFactory := func(service card.CardService) {
    group.GET("/", func(c *gin.Context) {
      results, err := service.ListCards()
      if err != nil {
        _ = c.Error(err)
        return
      }

      var responseItems = make([]CardResponse, len(results))

      for i := range results {
        responseItems[i] = *toResponseModel(&results[i])
      }

      response := &ListResponse{
        Data: responseItems,
      }

      c.JSON(http.StatusOK, response)
    })

    group.POST("/", func(c *gin.Context) {
      card, err := Bind(c)
      if err != nil {
        appError := domainErrors.NewAppError(err, domainErrors.ValidationError)
        c.Error(appError)
        return
      }

      newCard, err := service.CreateCard(card)
      if err != nil {
        _ = c.Error(err)
        return
      }

      c.JSON(http.StatusCreated, *toResponseModel(newCard))
    })



    group.GET("/:cardId", func(c *gin.Context) {
      id := c.Param("cardId")
      var i, err = strconv.Atoi(id)
      if err != nil {
        err = domainErrors.NewAppErrorWithType(domainErrors.ValidationError)
        _ = c.Error(err)
        return
      }

      result, err := service.ReadCard(i)
      if err != nil {
        _ = c.Error(err)
        return
      }

      c.JSON(http.StatusOK, *toResponseModel(result))
    })
  }

  return cardRoutesFactory
}
