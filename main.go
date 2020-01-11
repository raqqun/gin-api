package main

import (
    "os"
    "log"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"

    "github.com/raqqun/gin-api/routes"
    "github.com/raqqun/gin-api/config"
    "github.com/raqqun/gin-api/models"
)

func initDB() *gorm.DB {
    db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
    if err != nil {
        log.Printf("Cannot connect to %s database\n", "sqlite3")
        log.Panicln("Error", err)
    }

    gormDebug, err := strconv.ParseBool(os.Getenv("GORM_DEBUG"));
    if err != nil {
        log.Println(
            "GORM_DEBUG expects boolean 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.",
            "Used false by default.")
        gormDebug = true
    }

    db.LogMode(gormDebug)

    db.Debug().AutoMigrate(
        &models.Users{},
    )

    return db
}



func main() {
    router := gin.Default()
    port := config.Port()
    // router.RedirectTrailingSlash = false

    // init database
    db := initDB()
    defer db.Close()


    // register middlewares

    // register routes
    routes.API(router, db)

    log.Printf("Starting on port %s", port)

    router.Run(port)
}
