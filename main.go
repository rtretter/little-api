package main

import (
    "little-api/model"
    "little-api/routing"

    "github.com/labstack/echo"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var (
    db  *gorm.DB
    err error
	dsn = "host=localhost user=postgres password=r00tpwd dbname=little-api port=5432 sslmode=disable TimeZone=Europe/Berlin"
)

func main() {
    dbinit()
    e := echo.New()
    routing.Init(e)
    e.Logger.Fatal(e.Start(":1323"))
}

func dbinit() {
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
    }
    db.AutoMigrate(model.Note{})
}
