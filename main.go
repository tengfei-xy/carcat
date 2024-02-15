package main

import (
	"database/sql"
	"env"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	pnt "print"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

//
// init
//

func mainInit() {
	pnt.Init("init start")
	appCode := flag.String("c", "", "请指定AppCode")
	// c8401f02e0a84693a880f835ea2986cc
	flag.Parse()
	if *appCode == "" {
		pnt.ErrorString("no appCode")
		os.Exit(-1)
	}
	env.AppCode = *appCode
	DB = initMySQL()
	pnt.Init("init over")

}

func initMySQL() *sql.DB {

	Username := `root`
	Password := `if(hdc==MYSQL)`
	UnixSocket := `/tmp/mysql.sock`
	Database := `carcat`

	linkAddress := fmt.Sprintf("%s:%s@%s(%s)/%s", Username, Password, "unix", UnixSocket, Database)

	// 启动连接
	db, err := sql.Open("mysql", linkAddress)
	if err != nil {
		panic(err)
	}
	// 连接测试
	if err = db.Ping(); err != nil {
		panic(err)
	}
	pnt.Init("MySQL connection successful")

	return db
}

//
// main
//
func main() {
	mainInit()
	http.HandleFunc("/", index)
	pnt.Info("start!")

	go pnt.Info(http.ListenAndServe("0.0.0.0:1768", nil))

}
func index(w http.ResponseWriter, r *http.Request) {
	msg, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		pnt.Error(err)
	}
	w.Write(msgMain(msg))
}
