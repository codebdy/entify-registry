package repository

import (
	"database/sql"
	"fmt"
	"time"

	"rxdrag.com/entify-schema-registry/config"
	"rxdrag.com/entify-schema-registry/consts"
)

type Service struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Url         string    `json:"url"`
	ServiceType string    `json:"serviceType"`
	TypeDefs    string    `json:"tpeDefs"`
	IsAlive     bool      `json:"isAlive"`
	Version     string    `json:"version"`
	AddedTime   time.Time `json:"addedTime"`
	UpdatedTime time.Time `json:"updatedTime"`
}

var openedDB *sql.DB

func openDb() *sql.DB {
	if openedDB != nil {
		return openedDB
	}
	cfg := config.GetDbConfig()
	openedDB, err := sql.Open(cfg.Driver, DbString(cfg))

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	openedDB.SetMaxOpenConns(10)
	openedDB.SetMaxIdleConns(5)

	return openedDB
}

var fieldStr = `
			id,
			url,
			name,
			serviceType,
			typeDefs,
			version,
			isAlive,
			addedTime,
			updatedTime
	`

func serviceScanValues(service *Service) []interface{} {
	return []interface{}{
		&service.Id,
		&service.Name,
		&service.ServiceType,
		&service.TypeDefs,
		&service.Version,
		&service.IsAlive,
		&service.AddedTime,
		&service.UpdatedTime,
	}
}

func GetServices() []Service {
	var services = []Service{}
	db := openDb()
	sqlStr := fmt.Sprintf(`	SELECT %s	FROM services `, fieldStr)

	rows, err := db.Query(sqlStr)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	for rows.Next() {
		var service Service
		err = rows.Scan(serviceScanValues(&service)...)
		services = append(services, service)
	}
	return services
}

func Install(cfg config.DbConfig) {
	db, err := sql.Open(cfg.Driver, DbString(cfg))
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	sqlStr := `CREATE TABLE services (
		id int unsigned NOT NULL,
		url varchar(500) NOT NULL,
		name varchar(500) DEFAULT NULL,
		serviceType varchar(100) DEFAULT NULL,
		typeDefs longtext,
		version varchar(100) DEFAULT NULL,
		isAlive tinyint(1) DEFAULT NULL,
		addedTime varchar(45) DEFAULT NULL,
		updatedTime datetime DEFAULT NULL,
		PRIMARY KEY (id),
		UNIQUE KEY id_UNIQUE (id),
		UNIQUE KEY url_UNIQUE (url)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
	`

	_, err = db.Exec(sqlStr)
	if err != nil {
		panic(err.Error())
	}
}

func AddService(service Service) {
	db := openDb()

	sqlStr := `
	INSERT INTO services
		(id,
		url,
		name,
		serviceType,
		typeDefs,
		version,
		isAlive,
		addedTime,
		updatedTime)
		VALUES
		(?,?,?,?,?,?,?,?,?)
	`

	_, err := db.Exec(sqlStr,
		service.Id,
		service.Url,
		service.Name,
		service.ServiceType,
		service.TypeDefs,
		service.Version,
		service.IsAlive,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func UpdateService(service Service) {
	db := openDb()

	sqlStr := `
		UPDATE services
		SET
		url = ?,
		name = ?,
		serviceType = ?,
		typeDefs = ?,
		version = ?,
		isAlive = ?,
		updatedTime = ?
		WHERE id = ?
	`

	_, err := db.Exec(sqlStr,
		service.Url,
		service.Name,
		service.ServiceType,
		service.TypeDefs,
		service.Version,
		service.IsAlive,
		time.Now(),
		service.Id,
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func RemoveService(id int) {
	db := openDb()

	sqlStr := `
		DELETE FROM services
		WHERE id = ?;
	`

	_, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func GetService(id int) Service {
	var service Service
	db := openDb()
	sqlStr := fmt.Sprintf(`	SELECT %s	FROM services WHERE id = ?`, fieldStr)

	db.QueryRow(sqlStr, id).Scan(serviceScanValues(&service)...)
	return service

}

func GetAuthService() Service {
	var service Service
	db := openDb()
	sqlStr := fmt.Sprintf(`	SELECT %s	FROM services WHERE serviceType = ?`, fieldStr)

	db.QueryRow(sqlStr, consts.AUTH_SERVICE).Scan(serviceScanValues(&service)...)
	return service
}
