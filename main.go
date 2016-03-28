package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	
	// Custom package with models
	"_models"
	
	// Contrib packages
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

var DB *gorm.DB
var CFG *models.Config
const COUNT_PAGE int = 20

func main() {
	InitConfig()
	db_params := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", CFG.Username, CFG.Password, CFG.Name)
	db, err := gorm.Open("mysql", db_params)
	if err != nil {
		fmt.Println(err.Error())
	}

	DB = db

	// Get database connection handle [*sql.DB](http://golang.org/pkg/database/sql/#DB)
	DB.DB()
	DB.LogMode(true)
	// Then you could invoke `*sql.DB`'s functions with it
	DB.DB().Ping()
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	// Disable table name's pluralization
	DB.SingularTable(false)
	// DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&models.Anek{})

	router := InitRouter()

	host := fmt.Sprintf(":%d", CFG.AppPort)
	http.ListenAndServe(host, router)
}

func InitConfig() {
	config, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err.Error() + "\nRestore default config")
	}
	err = yaml.Unmarshal(config, &CFG)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/aneks", AnekIndex)
	router.HandleFunc("/aneks/random", AnekRandom)
	router.HandleFunc("/aneks/{anekId}", AnekShow)
	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func AnekIndex(w http.ResponseWriter, r *http.Request) {
	var aneks []models.Anek

	if len(r.URL.Query().Get("page")) > 0 {
		var count int = COUNT_PAGE
		if len(r.URL.Query().Get("count")) > 0 {
			count, _ = strconv.Atoi(r.URL.Query().Get("count"))
			
			if count >= 100 {
				count = COUNT_PAGE
			}
		}
		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			fmt.Println(err.Error())
			page = 0
		}
		DB.Limit(count).Offset(count * page).Find(&aneks)
	}

	w.Header().Set("Content-Type", "application/json")
	
	if len(aneks) > 0 {
		w.WriteHeader(http.StatusOK)
		a, err := json.Marshal(aneks)
		if err != nil {
			fmt.Println(err.Error())
		}

		w.Write(a)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func AnekRandom(w http.ResponseWriter, r *http.Request) {
	var anek models.Anek
	DB.Order("RAND()").First(&anek)

	a, err := json.Marshal(anek)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(a)
}

func AnekShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	anekId, err := strconv.Atoi(vars["anekId"])
	if err != nil {
		fmt.Println(err.Error())
	}

	var anek models.Anek
	DB.Where(&models.Anek{ID: anekId}).First(&anek)

	w.Header().Set("Content-Type", "application/json")
	if anek.ID != 0 {
		w.WriteHeader(http.StatusOK)
		a, err := json.Marshal(anek)
		if err != nil {
			fmt.Println(err.Error())
		}

		w.Write(a)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
