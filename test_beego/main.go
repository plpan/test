package main

import (
    "fmt"
    "log"
    orm2 "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

type User struct {
    Id int
    Name string
}

func init() {
    err := orm2.RegisterDriver("mysql", orm2.DRMySQL)
    if err != nil {
        log.Panic(err)
    }
    err = orm2.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&timeout=1s&readTimeout=1s&writeTimeout=500ms", 1, 1)
    if err != nil {
        log.Panic(err)
    }
    orm2.RegisterModel(new(User))
}

func main() {
    orm := orm2.NewOrm()
    user := User{Id:12}
    fmt.Println(orm.Read(&user))
    fmt.Println(user)
}
