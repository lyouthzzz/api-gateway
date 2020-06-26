package api

import (
	"github.com/lyouthzzz/poseidon-api-gateway/pkg/config"
	"github.com/lyouthzzz/poseidon-api-gateway/pkg/database"
	"github.com/lyouthzzz/poseidon-api-gateway/pkg/model"
	"log"
)

func main() {

	var conf config.Config
	if err := config.LoadConfig("./pkg/config/config.yaml", &conf); err != nil {
		log.Fatal(err)
	}
	dbConf := database.Config{
		Name:     conf.Database.Name,
		Host:     conf.Database.Host,
		Port:     conf.Database.Port,
		User:     conf.Database.User,
		Password: conf.Database.Password,
		Debug:    conf.Debug,
	}
	db, err := database.New(dbConf)
	if err != nil {
		log.Fatal(err)
	}
	err = database.MigrateTables(db, model.Project{}, model.Group{}, model.APIBackend{}, model.APIFrontend{}, model.APIResult{}, model.API{})
	if err != nil {
		log.Fatal(err)
	}
}
