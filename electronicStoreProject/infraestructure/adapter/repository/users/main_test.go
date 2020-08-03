package users

import(
     "testing"
     "log"
     "os"
     "github.com/jinzhu/gorm"
)

var testEntityManager *UserMysqlRepository

const (
	dbDriver= "mysql"
	dbSource = "root:root@tcp(localhost:3308)/market_electronics_tests?charset=utf8&parseTime=True&loc=UTC"
)

func TestMain(m *testing.M){
     conn, err := gorm.Open(dbDriver,dbSource)
     if err != nil{
     	log.Fatal("error with connection:",err)
     }
     testEntityManager = &UserMysqlRepository{
		Db: conn,
	}

     os.Exit(m.Run())
}