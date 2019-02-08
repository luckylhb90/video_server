package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func AddUserCredential(loginName string, pwd string) error {
	stmsIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) values (?, ?)")
	if err != nil {
		return err
	}
	stmsIns.Exec(loginName, pwd)
	stmsIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmsOut, err := dbConn.Prepare("SELECT Pwd from users where login_name=?")
	if err != nil {
		log.Panicln("%s", err)
		return "", err
	}
	var pwd string
	stmsOut.QueryRow(loginName).Scan(&pwd)
	stmsOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmsDel, err := dbConn.Prepare("DELETE from users where login_name=? and pwd=?")
	if err != nil {
		log.Panicln("%s", err)
		return err
	}
	stmsDel.Exec(loginName, pwd)
	stmsDel.Close()
	return nil
}
