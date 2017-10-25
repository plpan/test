package main

import (
    "fmt"
    "net"
    "io"
    "errors"
    "strings"
    "encoding/base64"
    "crypto/sha1"
)

func main() {
    ln, err := net.Listen("tcp", ":8088")
    if err != nil {
        panic(err)
    }

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }

        for {
            handleConnection(conn)
        }
    }
}

func handleConnection(conn net.Conn) {
    content := make([]byte, 1024)
    _, err := conn.Read(content)
    if err != nil {
        fmt.Println(err)
        return
    }

    isHttp := false
    if string(content[0:3]) == "GET" {
        isHttp = true
    }
    fmt.Println("isHttp: ", isHttp)

    if isHttp {
        header := parseHandShake(string(content))
        fmt.Println("header: ", header)

        secWebsocketKey := header["Sec-WebSocket-Key"]
        // ignore other validation
        guid := "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
        h := sha1.New()
        fmt.Println("accept raw: ", secWebsocketKey + guid)
        io.WriteString(h, secWebsocketKey + guid)
        accept := make([]byte, 28)
        base64.StdEncoding.Encode(accept, h.Sum(nil))
        fmt.Println(accept)

        response := "HTTP/1.1 101 Switching Protocols\r\n"
        response = response + "Sec-WebSocket-Accept: " + string(accept) + "\r\n"
        response = response + "Connection: Upgrade\r\n"
        response = response + "Upgrade: websocket\r\n"
        fmt.Println("response: ", response)

        websocket := NewWsSocket(conn)
        for {
            data, err := websocket.ReadIframe()
            if err != nil {
                fmt.Println(err)
                continue
            }
            fmt.Println("read: ", string(data))

            err = websocket.SendIframe(data)
            if err != nil {
                fmt.Println(err)
                continue
            }
            fmt.Println("send data successful")
        }
    } else {
        fmt.Println(string(content))
    }
}

type WebSocket struct {
    MaskingKey []byte
    Conn       net.Conn
}

func NewWsSocket(conn net.Conn) *WebSocket {
    return &WebSocket{Conn: conn}
}

func (this *WebSocket) ReadIframe() (data []byte, err error) {
    err = nil

    opcodebyte := make([]byte, 1)
    this.Conn.Read(opcodebyte)
    FIN := opcodebyte[0] >> 7
    RSV1 := opcodebyte[0] >> 6 & 1
    RSV2 := opcodebyte[0] >> 5 & 1
    RSV3 := opcodebyte[0] >> 4 & 1
    OPCODE := opcodebyte[0] & 15
    fmt.Println(FIN, RSV1, RSV2, RSV3, OPCODE)

    payloadLenByte := make([]byte, 1)
    this.Conn.Read(payloadLenByte)
    payloadLen := int(payloadLenByte[0] & 0x7f)
    fmt.Println(payloadLen)
    if payloadLen == 127 {
        extendedByte := make([]byte, 8)
        this.Conn.Read(extendedByte)
    }

    mask := payloadLenByte[0] >> 7
    maskingByte := make([]byte, 4)
    this.Conn.Read(maskingByte)

    payloadDataByte := make([]byte, payloadLen)
    this.Conn.Read(payloadDataByte)
    fmt.Println("data: ", payloadDataByte)

    dataByte := make([]byte, payloadLen)
    for i := 0; i < payloadLen; i++ {
        if mask == 1 {
            dataByte[i] = payloadDataByte[i] ^ maskingByte[i % 4]
        } else {
            dataByte[i] = payloadDataByte[i]
        }
    }

    if FIN == 1 {
        data = dataByte
        return
    }

    nextData, err := this.ReadIframe()
    if err != nil {
        return
    }

    data = append(data, nextData...)
    return
}

func (this *WebSocket) SendIframe(data []byte) error {
    if len(data) >= 125 {
        return errors.New("send data length error")
    }

    length := len(data)
    maskedData := make([]byte, length)
    for i := 0; i < length; i++ {
        if this.MaskingKey != nil {
            maskedData[i] = data[i] ^ this.MaskingKey[i % 4]
        } else {
            maskedData[i] = data[i]
        }
    }

    this.Conn.Write([]byte{0x81})

    var payLenByte byte
    if this.MaskingKey != nil && len(this.MaskingKey) != 4 {
        payLenByte = byte(0x80) | byte(length)
        this.Conn.Write([]byte{payLenByte})
        this.Conn.Write(this.MaskingKey)
    } else {
        payLenByte = byte(0x00) | byte(length)
        this.Conn.Write([]byte{payLenByte})
    }
    this.Conn.Write(data)
    return nil
}

func parseHandShake(content string) map[string]string {
    header := make(map[string]string, 10)
    lines := strings.Split(content, "\r\n")

    for _, line := range lines {
        fmt.Println(line)
        if len(line) > 0 {
            words := strings.Split(line, ":")
            if len(words) == 2 {
                header[strings.Trim(words[0], " ")] = strings.Trim(words[1], " ")
            }
        }
    }

    return header
}
