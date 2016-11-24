package main

import (
  "fmt"
  "github.com/jinzhu/gorm"
  "github.com/kataras/iris"
  _ "github.com/mattn/go-sqlite3"
)

type User struct {
  Id   int
  Name string `sql:"unique"`
  Age  int
}

func main() {
  db, _ := gorm.Open("sqlite3", "data/store.db")

  if (!db.HasTable(new(User))) {
    fmt.Print("Table doesn't exist. Creating new...\n")

    db.CreateTable(new(User))

    fmt.Print("Inserting two users into the table...\n")
    userOne := User{Name: "Ann", Age: 101}
    userTwo := User{Name: "Bob", Age: 102}

    db.Save(&userOne)
    db.Save(&userTwo)
  }

  var bob User

  db.Where("name = ?", "Bob").Find(&bob)
  body := fmt.Sprintf("<h1>Dockerized Go server is running!</h1><p>Age of %s is %d.\n</p>", bob.Name, bob.Age)

  iris.Get("/", func(c *iris.Context) {
    c.HTML(iris.StatusOK, body)
  })

  iris.Listen(":8080")

}