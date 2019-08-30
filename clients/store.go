package clients

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "Store > ", log.LstdFlags)

const (
	OREGON_LEGISLATURE_API string = "oregonLegislatureAPI"
)

type Item struct {
	key string
	payload interface{}
}

type Store struct {
	Context map[string]interface{}
	Set func(i Item)
	Get func(key string)
	Vi viper.Viper
}

type options func(store *Store) *Store

func (s *Store) initialize(options ...options) {
	for _, opt := range options {
		s = opt(s)
	}
}

func (s *Store) InitializeViperEndpoints() {

	viper.SetConfigType("json")

	viper.SetConfigName("endpoints")

	viper.AddConfigPath("./store/")

	err := viper.ReadInConfig()

	if err != nil {
		logger.Panic(err)
	}
}