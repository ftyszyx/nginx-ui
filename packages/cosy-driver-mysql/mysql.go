package mysql

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type DBSettings interface {
    GetHost() string
    GetUser() string
    GetPassword() string
    GetName() string
    GetPort() uint
}

func Open(dbs DBSettings) gorm.Dialector {
    return mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        dbs.GetUser(), dbs.GetPassword(), dbs.GetHost(), dbs.GetPort(), dbs.GetName()))
}
