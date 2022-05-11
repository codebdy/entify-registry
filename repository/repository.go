package repository

import (
	"database/sql"
	"fmt"
	"time"

	"rxdrag.com/entify-schema-registry/config"
	"rxdrag.com/entify-schema-registry/consts"
)

type Service struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	Url         string    `json:"url"`
	ServiceType string    `json:"serviceType"`
	TypeDefs    string    `json:"tpeDefs"`
	IsAlive     bool      `json:"isAlive"`
	Version     string    `json:"version"`
	AddedTime   time.Time `json:"addedTime"`
	UpdatedTime time.Time `json:"updatedTime"`
}

func GetServices() []Service {
	var services = []Service{}

	return services
}

func Install(cfg config.DbConfig) {
	db, err := sql.Open(config.GetString(consts.DB_DRIVER), DbString(cfg))
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	sql := `CREATE TABLE services (
		id int unsigned NOT NULL,
		name varchar(500) DEFAULT NULL,
		url varchar(500) DEFAULT NULL,
		serviceType varchar(100) DEFAULT NULL,
		typeDefs longtext,
		version varchar(100) DEFAULT NULL,
		isAlive tinyint(1) DEFAULT NULL,
		addedTime varchar(45) DEFAULT NULL,
		updatedTime datetime DEFAULT NULL,
		PRIMARY KEY (id),
		UNIQUE KEY id_UNIQUE (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
	`

	_, err = db.Exec(sql)
	if err != nil {
		panic(err.Error())
	}
}
