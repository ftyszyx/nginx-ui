package sqlite

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "path"
)

type DBSettings interface {
    GetName() string
}

func Open(dir string, dbs DBSettings) gorm.Dialector {
    return sqlite.Open(path.Join(dir, fmt.Sprintf("%s.db", dbs.GetName())))
}
