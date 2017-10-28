/*客户端方法*/
package request

import (
    "golang.org/x/net/context"
    pb "test/test_grpc"
    "fmt"
    "io"
)

//简单模式
func GetUserInfo(client pb.DataClient, info *pb.URequest)  {
    req, err := client.GetUserInfo(context.Background(),info)
    if err != nil {
        fmt.Println("Could not create Customer: %v", err)
    }
    fmt.Println("userinfo is ",req.GetAge(),req.GetName(),req.GetSex())
}

//双向流模式
func ChangeUserInfo(client pb.DataClient){
    notes := []*pb.UResponse{
        {Name:"jim",Age:18,Sex:2},
        {Name:"Tom",Age:20,Sex:1},
    }
    stream, err := client.ChangeUserInfo(context.Background())
    if err != nil {
        fmt.Println("%v.RouteChat(_) = _, %v", client, err)
    }
    waitc := make(chan struct{})
    go func() {
        for {
            in, err := stream.Recv()
            if err == io.EOF {
                // read done.
                fmt.Println("read done ")
                close(waitc)
                return
            }
            if err != nil {
                fmt.Println("Failed to receive a note : %v", err)
            }
            fmt.Println("Got message %s at point(%d, %d)",in.Sex,in.Age,in.Name)
        }
    }()
    fmt.Println("notes",notes)
    for _, note := range notes {
        if err := stream.Send(note); err != nil {
            fmt.Println("Failed to send a note: %v", err)
        }
    }
    stream.CloseSend()
    <-waitc
}
