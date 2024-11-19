package sqlite

import (
    "gorm.io/driver/sqlite"
    "testing"
)

type DataBase struct {
    Name string
}

func (d *DataBase) GetName() string {
    return d.Name
}

func TestOpen(t *testing.T) {
    dbs := &DataBase{
        Name: "cosy",
    }

    dialector := Open("", dbs)

    d, ok := dialector.(*sqlite.Dialector)
    if !ok {
        t.Fatal("dialector is not *Dialector")
    }
    if d.DSN == "" {
        t.Error("dialector.DSN is empty")
    }
}
