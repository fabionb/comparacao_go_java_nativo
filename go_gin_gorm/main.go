package main

import (
    "os"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "runtime/pprof"
    "fmt"
    "time"
)

type Test1Entity struct {
    Id int64  `gorm:"primaryKey,column:id,type:autoIncrement"`
}

func (Test1Entity) TableName() string {
  return "test1entity"
}

func main() {
    db, err := gorm.Open(mysql.Open("root:root@/sys"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    DB, err := db.DB()
    if err != nil {
        panic("failed to connect database")
    }
    
    DB.SetMaxIdleConns(0)
    DB.SetMaxOpenConns(10)

    db.AutoMigrate(&Test1Entity{})

    gin.SetMode(gin.ReleaseMode)
    router := gin.Default()

    router.GET("/", func(c *gin.Context) {
        err := db.Transaction(func(tx *gorm.DB) error {
            if err := tx.Create(&Test1Entity{}).Error; err != nil {
                return err
            }
            return nil
        })
        if err != nil {
            panic(err)
        }
        
        c.AbortWithStatus(200)
    })

    ticker := time.NewTicker(500 * time.Second)
    go func() {
        for {
           select {
            case <- ticker.C:
                fmt.Println("DUMP")
                pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
            }
        }
    }()
    router.Run(":8081")
}
