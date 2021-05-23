package pg

import (
	"hackz-api/pkg/api/infrastructure/persistence/gorm/model"

	"github.com/jinzhu/gorm"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

var db *gorm.DB

type Env struct {
	Sslmode  string `default:"disable"`
	Port     string `default:"5432"`
	Host     string `default:"localhost"`
	Dbname   string `default:"test"`
	User     string `default:"postgres"`
	Password string `default:"password"`
}

func Connect() *gorm.DB {
	var err error
	var pgenv Env
	if err := envconfig.Process("pg", &pgenv); err != nil {
		panic(err)
	}

	db, err = gorm.Open("postgres", "host="+pgenv.Host+" port="+pgenv.Port+" user="+pgenv.User+" dbname="+pgenv.Dbname+" password="+pgenv.Password+" sslmode="+pgenv.Sslmode)

	if err != nil {
		panic(err)
	}

	if !db.HasTable(&model.Task{}) {
		if err := db.Table("tasks").CreateTable(&model.Task{}).Error; err != nil {
			panic(err)
		}
	}

	db.LogMode(true)

	return db
}

func CloseConn() {
	db.Close()
}
