package repository

import (
	"database/sql"
	"fmt"
	"time"

	"rxdrag.com/entify-schema-registry/config"
)

type Service struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Url         string    `json:"url"`
	TypeDefs    string    `json:"typeDefs"`
	IsAlive     bool      `json:"isAlive"`
	Version     string    `json:"version"`
	AddedTime   time.Time `json:"addedTime"`
	UpdatedTime time.Time `json:"updatedTime"`
}

type ServiceOutput struct {
	Id          int
	Name        string
	Url         string
	ServiceType sql.NullString
	TypeDefs    sql.NullString
	IsAlive     sql.NullBool
	Version     sql.NullString
	AddedTime   sql.NullTime
	UpdatedTime sql.NullTime
}

func (service ServiceOutput) covertService() *Service {
	return &Service{
		Name:        service.Name,
		Url:         service.Url,
		TypeDefs:    service.TypeDefs.String,
		IsAlive:     service.IsAlive.Bool,
		Version:     service.Version.String,
		AddedTime:   service.AddedTime.Time,
		UpdatedTime: service.UpdatedTime.Time,
	}
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
			typeDefs,
			version,
			isAlive,
			addedTime,
			updatedTime
	`

func serviceScanValues(service *ServiceOutput) []interface{} {
	return []interface{}{
		&service.Url,
		&service.Name,
		&service.TypeDefs,
		&service.Version,
		&service.IsAlive,
		&service.AddedTime,
		&service.UpdatedTime,
	}
}

func GetServices() []*Service {
	var services = []*Service{}
	db := openDb()
	sqlStr := fmt.Sprintf(`	SELECT %s	FROM services `, fieldStr)

	rows, err := db.Query(sqlStr)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	for rows.Next() {
		var service ServiceOutput
		err = rows.Scan(serviceScanValues(&service)...)
		services = append(services, service.covertService())
	}
	return services
}

func IsInstalled() bool {
	sqlStr := fmt.Sprintf(
		"SELECT COUNT(*) FROM information_schema.TABLES WHERE table_name ='services' AND table_schema ='%s'",
		config.GetDbConfig().Database,
	)
	db := openDb()

	var count int
	err := db.QueryRow(sqlStr).Scan(&count)
	switch {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		panic(err.Error())
	}
	return count > 0
}

func Install() {
	cfg := config.GetDbConfig()
	db, err := sql.Open(cfg.Driver, DbString(cfg))
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	sqlStr := `CREATE TABLE services (
		id int unsigned NOT NULL AUTO_INCREMENT,
		url varchar(500) NOT NULL,
		name varchar(500) DEFAULT NULL,
		typeDefs longtext,
		version varchar(100) DEFAULT NULL,
		isAlive tinyint(1) DEFAULT NULL,
		addedTime datetime DEFAULT NULL,
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
		(url,
		name,
		typeDefs,
		version,
		isAlive,
		addedTime,
		updatedTime)
		VALUES
		(?,?,,?,?,?,?,?,?)
	`

	_, err := db.Exec(sqlStr,
		service.Url,
		service.Name,
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
		typeDefs = ?,
		version = ?,
		isAlive = ?,
		updatedTime = ?
		WHERE id = ?
	`

	_, err := db.Exec(sqlStr,
		service.Url,
		service.Name,
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

func GetService(id int) *Service {
	var service ServiceOutput
	db := openDb()
	sqlStr := fmt.Sprintf(`	SELECT %s	FROM services WHERE id = ?`, fieldStr)

	err := db.QueryRow(sqlStr, id).Scan(serviceScanValues(&service)...)

	switch {
	case err == sql.ErrNoRows:
		return nil
	case err != nil:
		panic(err.Error())
	}

	return service.covertService()

}
