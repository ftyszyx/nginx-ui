package cosy

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/uozi-tech/cosy/logger"
	"github.com/uozi-tech/cosy/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

func (c *Ctx[T]) SetFussy(keys ...string) *Ctx[T] {
	c.fussy = append(c.fussy, keys...)
	for _, key := range keys {
		c.gormScopes = append(c.gormScopes, func(tx *gorm.DB) *gorm.DB {
			return QueryToFussySearch(c.Context, tx, key)
		})
	}
	return c
}

func (c *Ctx[T]) SetSearchFussyKeys(keys ...string) *Ctx[T] {
	c.search = append(c.search, keys...)
	c.gormScopes = append(c.gormScopes, func(tx *gorm.DB) *gorm.DB {
		return QueryToFussyKeysSearch(c.Context, tx, keys...)
	})
	return c
}

func (c *Ctx[T]) SetEqual(keys ...string) *Ctx[T] {
	c.eq = append(c.eq, keys...)
	c.gormScopes = append(c.gormScopes, func(tx *gorm.DB) *gorm.DB {
		return QueryToEqualSearch(c.Context, tx, keys...)
	})
	return c
}

func (c *Ctx[T]) SetIn(keys ...string) *Ctx[T] {
	c.in = append(c.in, keys...)
	c.gormScopes = append(c.gormScopes, func(tx *gorm.DB) *gorm.DB {
		return QueriesToInSearch(c.Context, tx, keys...)
	})
	return c
}

func (c *Ctx[T]) SetInWithKey(value string, key string) *Ctx[T] {
	c.gormScopes = append(c.gormScopes, func(tx *gorm.DB) *gorm.DB {
		return QueryToInSearch(c.Context, tx, value, key)
	})
	return c
}

func (c *Ctx[T]) SetOrFussy(keys ...string) *Ctx[T] {
	c.orFussy = append(c.orFussy, keys...)
	c.gormScopes = append(c.gormScopes, func(tx *gorm.DB) *gorm.DB {
		return QueryToOrFussySearch(c.Context, tx, keys...)
	})
	return c
}

func (c *Ctx[T]) SetOrEqual(keys ...string) *Ctx[T] {
	c.orEq = append(c.orEq, keys...)
	c.gormScopes = append(c.gormScopes, func(tx *gorm.DB) *gorm.DB {
		return QueryToOrEqualSearch(c.Context, tx, keys...)
	})
	return c
}

func (c *Ctx[T]) SetBetween(keys ...string) *Ctx[T] {
	c.between = append(c.between, keys...)
	c.gormScopes = append(c.gormScopes, func(tx *gorm.DB) *gorm.DB {
		return QueriesToBetweenSearch(c.Context, tx, keys...)
	})
	return c
}

func (c *Ctx[T]) SetBetweenWithKey(value string, key string) *Ctx[T] {
	c.gormScopes = append(c.gormScopes, func(tx *gorm.DB) *gorm.DB {
		return QueryToBetweenSearch(c.Context, tx, value, key)
	})
	return c
}

func (c *Ctx[T]) SetOrIn(keys ...string) *Ctx[T] {
	c.orIn = append(c.orIn, keys...)
	c.gormScopes = append(c.gormScopes, func(tx *gorm.DB) *gorm.DB {
		return QueryToOrInSearch(c.Context, tx, keys...)
	})
	return c
}

func QueriesToInSearch(c *gin.Context, db *gorm.DB, keys ...string) *gorm.DB {
	for _, v := range keys {
		QueryToInSearch(c, db, v)
	}
	return db
}

func QueryToInSearch(c *gin.Context, db *gorm.DB, value string, key ...string) *gorm.DB {
	queryArray := c.QueryArray(value + "[]")
	if len(queryArray) == 0 {
		queryArray = c.QueryArray(value)
	}
	if len(queryArray) == 1 && queryArray[0] == "" {
		return db
	}

	if len(queryArray) >= 1 {
		var builder strings.Builder
		stmt := db.Statement

		column := value
		if len(key) != 0 {
			column = key[0]
		}

		stmt.QuoteTo(&builder, clause.Column{Table: stmt.Table, Name: column})
		builder.WriteString(" IN ?")

		return db.Where(builder.String(), queryArray)
	}
	return db
}

func QueryToEqualSearch(c *gin.Context, db *gorm.DB, keys ...string) *gorm.DB {
	for _, v := range keys {
		if c.Query(v) != "" {
			var sb strings.Builder
			stmt := db.Statement

			stmt.QuoteTo(&sb, clause.Column{Table: stmt.Table, Name: v})
			sb.WriteString(" = ?")

			db = db.Where(sb.String(), c.Query(v))
		}
	}
	return db
}

func QueryToFussySearch(c *gin.Context, db *gorm.DB, key string) *gorm.DB {
	if qArr := c.QueryArray(key + "[]"); qArr != nil {
		db = applyFuzzyCondition(db, key, qArr)
	} else if q := c.Query(key); q != "" {
		db = applyFuzzyCondition(db, key, []string{q})
	}
	return db
}

func applyFuzzyCondition(tx *gorm.DB, column string, values []string) *gorm.DB {
	stmt := tx.Statement

	// build column name (column LIKE ?)
	var colBuilder strings.Builder
	stmt.QuoteTo(&colBuilder, clause.Column{Table: stmt.Table, Name: column})
	colBuilder.WriteString(" LIKE ?")

	db := model.UseDB()
	var valueBuilder strings.Builder

	for _, value := range values {
		// build value for query (%value%)
		valueBuilder.Reset()
		valueBuilder.WriteString("%")
		valueBuilder.WriteString(value)
		valueBuilder.WriteString("%")

		db = db.Or(colBuilder.String(), valueBuilder.String())
	}

	return tx.Where(db)
}

func QueryToFussyKeysSearch(c *gin.Context, tx *gorm.DB, keys ...string) *gorm.DB {
	value := c.Query("search")
	if value == "" {
		return tx
	}

	// build value for query (%value%)
	var valueBuilder strings.Builder
	valueBuilder.WriteString("%")
	valueBuilder.WriteString(value)
	valueBuilder.WriteString("%")
	likeValue := valueBuilder.String()

	db := model.UseDB()
	var colBuilder strings.Builder

	for _, v := range keys {
		// build column name (column LIKE ?)
		colBuilder.Reset()
		colBuilder.WriteString(v)
		colBuilder.WriteString(" LIKE ?")

		db = db.Or(colBuilder.String(), likeValue)
	}

	return tx.Where(db)
}

func QueryToOrInSearch(c *gin.Context, db *gorm.DB, keys ...string) *gorm.DB {
	for _, v := range keys {
		queryArray := c.QueryArray(v + "[]")
		if len(queryArray) == 0 {
			queryArray = c.QueryArray(v)
		}
		if len(queryArray) == 1 && queryArray[0] == "" {
			continue
		}
		if len(queryArray) >= 1 {
			var sb strings.Builder
			stmt := db.Statement

			stmt.QuoteTo(&sb, clause.Column{Table: stmt.Table, Name: v})
			sb.WriteString(" IN ?")

			db = db.Or(sb.String(), queryArray)
		}
	}
	return db
}

func QueryToOrEqualSearch(c *gin.Context, db *gorm.DB, keys ...string) *gorm.DB {
	for _, v := range keys {
		if c.Query(v) != "" {
			var sb strings.Builder
			stmt := db.Statement

			stmt.QuoteTo(&sb, clause.Column{Table: stmt.Table, Name: v})
			sb.WriteString(" = ?")

			db = db.Or(sb.String(), c.Query(v))
		}
	}
	return db
}

func QueryToOrFussySearch(c *gin.Context, db *gorm.DB, keys ...string) *gorm.DB {
	for _, v := range keys {
		if c.Query(v) != "" {
			var sb strings.Builder
			stmt := db.Statement

			stmt.QuoteTo(&sb, clause.Column{Table: stmt.Table, Name: v})

			sb.WriteString(" LIKE ?")

			var sbValue strings.Builder

			_, err := fmt.Fprintf(&sbValue, "%%%s%%", c.Query(v))

			if err != nil {
				logger.Error(err)
				continue
			}

			db = db.Or(sb.String(), sbValue.String())
		}
	}
	return db
}

func QueriesToBetweenSearch(c *gin.Context, db *gorm.DB, keys ...string) *gorm.DB {
	for _, v := range keys {
		db = QueryToBetweenSearch(c, db, v)
	}
	return db
}

func QueryToBetweenSearch(c *gin.Context, db *gorm.DB, value string, key ...string) *gorm.DB {
	queryArray := c.QueryArray(value + "[]")
	if len(queryArray) == 0 {
		queryArray = c.QueryArray(value)
	}
	if len(queryArray) <= 1 {
		return db
	}

	if len(queryArray) == 2 && queryArray[0] != "" && queryArray[1] != "" {
		var builder strings.Builder
		stmt := db.Statement

		column := value
		if len(key) != 0 {
			column = key[0]
		}

		stmt.QuoteTo(&builder, clause.Column{Table: stmt.Table, Name: column})
		builder.WriteString(" BETWEEN ? AND ?")

		return db.Where(builder.String(), queryArray[0], queryArray[1])
	}
	return db
}
