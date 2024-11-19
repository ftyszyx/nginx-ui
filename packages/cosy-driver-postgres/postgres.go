package postgres

import (
    "fmt"
    "gorm.io/driver/postgres"
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
    return postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
        dbs.GetHost(), dbs.GetUser(), dbs.GetPassword(), dbs.GetName(), dbs.GetPort()))
}
