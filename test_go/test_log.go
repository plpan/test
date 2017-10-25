package main

import (
    "os"
    "log"
)

func main() {
    f()
    f()
    f()
}

func f() {
    logFile, err := os.OpenFile("./test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalln("open log file error")
    }
    logger := log.New(logFile, "[tag]", log.Llongfile)
    logger.Println("d", "c")
}
