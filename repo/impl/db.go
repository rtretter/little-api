package impl

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"little-api/repo"
)

type Db struct {
	conn *gorm.DB
}

func (d *Db) Update(object interface{}) {
    d.conn.Save(object)
}

func (d *Db) Create(object interface{}) {
	d.conn.Create(object)
}

func (d *Db) DeleteById(object interface{}, id string) {
	d.conn.Delete(object, id)
}

func (d *Db) FindAll(object interface{}) {
	d.conn.Find(object)
}

func (d *Db) Find(object interface{}, id string) error {
    result := d.conn.First(object, id)
    return result.Error
}

func CreateDB() repo.Db {
	dsn := "host=localhost user=postgres password=r00tpwd dbname=little-api port=5432 sslmode=disable TimeZone=Europe/Berlin"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	database := new(Db)
	database.conn = db
	return database
}
