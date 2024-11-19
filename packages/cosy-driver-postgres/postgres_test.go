package postgres

import (
    "gorm.io/driver/postgres"
    "testing"
)

type DataBase struct {
    User     string
    Password string
    Host     string
    Port     uint
    Name     string
}

func (d *DataBase) GetName() string {
    return d.Name
}

func (d *DataBase) GetHost() string {
    return d.Host
}

func (d *DataBase) GetPassword() string {
    return d.Password
}

func (d *DataBase) GetPort() uint {
    return d.Port
}

func (d *DataBase) GetUser() string {
    return d.User
}

func TestOpen(t *testing.T) {
    dbs := &DataBase{
        User:     "cosy",
        Password: "cosy",
        Host:     "127.0.0.1",
        Port:     5432,
        Name:     "cosy",
    }

    dialector := Open(dbs)

    d, ok := dialector.(*postgres.Dialector)
    if !ok {
        t.Fatal("dialector is not *Dialector")
    }
    if d.DSN == "" {
        t.Error("dialector.DSN is empty")
    }
}
