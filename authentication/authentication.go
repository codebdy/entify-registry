package authentication

import (
	"database/sql"
	"errors"
	"fmt"

	"rxdrag.com/entify-schema-registry/repository"
)

func Login(loginName, pwd string) (string, error) {
	db, err := sql.Open("mysql", repository.DbConfig())
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	var password string
	err = db.QueryRow("select password from rx_user where loginName = ?", loginName).Scan(&password)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("Login failed!")
	}

	// err = bcrypt.CompareHashAndPassword([]byte(pwd), []byte(password)) //验证（对比）
	// if err != nil {
	// 	fmt.Println(err)
	// 	return "", errors.New("Password error!")
	// }
	return loginName, err
}

func Logout() {

}
