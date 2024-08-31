package main

import (
  "fmt"
  "database/sql"
  "time"

  _ "github.com/go-sql-driver/mysql" 
)

type User struct {
  Uuid string
  Firstname string
  Lastname string
  Email string
  password string
}

func main() {
  fmt.Printf("Hello world")
  users, err := retrieveUsers()
  if err != nil {
    panic(err)
  }
  fmt.Printf("%v", users)
}

func retrieveUsers() ([]User, error) {
  db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/sport_track_go")
  if err != nil {
	  panic(err)
  }
  // See "Important settings" section.
  fmt.Println(db)
  db.SetConnMaxLifetime(time.Minute * 3)
  db.SetMaxOpenConns(10)
  db.SetMaxIdleConns(10)
  pingErr := db.Ping()
    if pingErr != nil {
        panic(pingErr)
    }
  fmt.Println("Connected!")
  var users []User
  rows, err := db.Query("SELECT * from Users")
  if err != nil {
    panic(err)
  }

  defer rows.Close()

  for rows.Next() {
    var user User
    if err := rows.Scan(&user.Firstname, &user.Lastname, &user.password, &user.Email, &user.Uuid); err != nil {
      return nil, fmt.Errorf("%v", err)
    }
    users = append(users, user)
  }

  if err := rows.Err(); err != nil {
    return nil, fmt.Errorf("%v", err)
  } 
  return users, nil
}
