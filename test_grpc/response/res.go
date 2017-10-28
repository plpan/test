/*服务端的方法*/
package response

import (
    "golang.org/x/net/context"
    pb "test/test_grpc"
    "fmt"
    "io"
)

type Server struct{
    routeNotes    []*pb.UResponse
}

//简单模式
func (this *Server)GetUserInfo(ctx context.Context, in *pb.URequest)(*pb.UResponse,error){
    uid := in.GetUid()
    fmt.Println("The uid is ",uid)
    return &pb.UResponse{
        Name : "Jim",
        Age  : 18,
        Sex : 0,
    },nil
}

//双向流模式
func (this *Server) ChangeUserInfo(stream pb.Data_ChangeUserInfoServer)(error){
    for {
        in, err := stream.Recv()
        if err == io.EOF {
            fmt.Println("read done")
            return nil
        }
        if err != nil {
            fmt.Println("ERR",err)
            return err
        }
        fmt.Println("userinfo ",in)
        for _, note := range this.routeNotes{
            if err := stream.Send(note); err != nil {
                return err
            }
        }
    }
}
