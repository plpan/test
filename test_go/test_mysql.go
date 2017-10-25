package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func main() {
    db, err := sql.Open("mysql", "root:123456@/test")
    checkErr(err)
    defer db.Close()

    err = db.Ping()
    checkErr(err)

    // Prepare insert statement
    stmt, err := db.Prepare("INSERT INTO square VALUES(?, ?)")
    checkErr(err)
    defer stmt.Close()

    // insert square number into table
    for i := 1; i < 5; i++ {
        res, err := stmt.Exec(i, (i * i))
        _ = res
        checkErr(err)
    }

    // Prepare select statement
    stmt, err = db.Prepare("SELECT value FROM square WHERE id=?")
    checkErr(err)

    var num int // scan the result in here

    err = stmt.QueryRow(1).Scan(&num)
    checkErr(err)

    fmt.Printf("The square number of %d is : %d\n", 1, num)

    // Prepare delete statemenet
    stmt, err = db.Prepare("DELETE FROM square")
    checkErr(err)
    res, err := stmt.Exec()
    _ = res
    checkErr(err)
}


func checkErr(err error) {
    if err != nil {
        panic(err.Error())
    }
}
