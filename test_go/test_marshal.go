package main

import (
    "fmt"
    "encoding/json"
)

type ShirtSize string

const (
    XS ShirtSize = "XS"
    S = "S"
    M
    L
    XL
)

type input struct {
    Name string  `json:"name"`
    Born string  `json:"birthdate"`
    Size ShirtSize     `json:"shirt-size"`
}

func (ss *ShirtSize) UnmarshalJSON(data []byte) error {
    var s string
    if err := json.Unmarshal(data, &s); err != nil {
        return fmt.Errorf("shirt-size should be a string, got %s", data)
    }
    fmt.Println("输入的字符串为111:",s)
    got, ok := map[string]ShirtSize{"XS": XS, "S": S, "M": M, "L": L, "X": XL}[s]
    if !ok {
        return fmt.Errorf("invalid ShirtSize %q", s)
    }
    *ss = got
    return nil
}

func main() {

    //type ShirtSize uint
    s :=input{"Gopher","2009/11/10","S"}
    b , _ :=json.Marshal(s)
    fmt.Println("转化的json串为:", b)

    var data input

    if err := json.Unmarshal(b, &data); err != nil {
        fmt.Println("error occurs")
    }
    fmt.Println(data);
}
