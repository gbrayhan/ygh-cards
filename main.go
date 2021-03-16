package main

import (
  "fmt"
  "net/http"

  "github.com/spf13/viper"

  cardsStore "github.com/gbrayhan/academy-go-q12021/data/card/csvstore"
  apiDeckStore "github.com/gbrayhan/academy-go-q12021/data/card/externalprodeck"
  "github.com/gbrayhan/academy-go-q12021/domain/card"
  routerHttp "github.com/gbrayhan/academy-go-q12021/router/http"
)

func main() {
  cardsRepoCSV := cardsStore.New()
  cardsSvc := card.NewService(cardsRepoCSV)


  cardsRepoAPIDeck := apiDeckStore.New()
  cardsAPIDeckSvc := card.NewService(cardsRepoAPIDeck)


  // router.ApplicationV1Router(router)
  httpRouter := routerHttp.NewHTTPHandler(cardsSvc, cardsAPIDeckSvc)

  viper.SetConfigFile("config.json")
  if err := viper.ReadInConfig(); err != nil {
    panic(fmt.Errorf("fatal error in config file: %s \n", err))
  }
  err := http.ListenAndServe(":"+viper.GetString("ServerPort"), httpRouter)
  if err != nil {
    panic(err)
  }
}
