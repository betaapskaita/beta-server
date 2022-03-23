package databases

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MariaDB struct {
	connection *gorm.DB
	connStr    string
}

func NewMariaDB(connStr string) Database {
	db := &MariaDB{
		connStr: connStr,
	}

	return db
}

func (db *MariaDB) Connect() (*gorm.DB, error) {
	var err error

	if db.connection == nil {
		db.connection, err = db.getConnection()
	}

	return db.connection, err
}

func (db *MariaDB) GetConnection() *gorm.DB {
	conn, err := db.Connect()
	if err != nil {
		log.Println(err)
	}

	return conn
}

func (db *MariaDB) getConnection() (*gorm.DB, error) {
	dialect := mysql.New(mysql.Config{
		DSN:                       db.connStr, // data source name
		DefaultStringSize:         256,        // default size for string fields
		DisableDatetimePrecision:  true,       // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,       // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,       // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,      // auto configure based on currently MySQL version
	})

	conn, err := gorm.Open(dialect, &gorm.Config{})

	if err != nil {
		log.Println(err, db.connStr)
		return conn, err
	}

	return conn, err
}
