package main

import(
	"net"
	"fmt"
	"context"
	pb "github.com/GonzaloDiaz300/LAB_2/proto"
	"google.golang.org/grpc"
)

type oms struct{
	pb.UnimplementedIntercambiosServer
}


func (a *oms) Notificar(ctx context.Context, in *pb.ContiReq) (*pb.ContiResp, error){
	fmt.Printf("Estado Recibido: %s %s %s \n",in.Nombre,in.Apellido,in.Estado)
	return &pb.ContiResp{Respuesta: 1}, nil
}

func main() {
	listner, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}

	serv := grpc.NewServer()
	fmt.Printf("Servidor OMS Activo\n")
	pb.RegisterIntercambiosServer(serv, &oms{})
	if err = serv.Serve(listner); err != nil {
		panic("cannot initialize the server" + err.Error())
	}
}