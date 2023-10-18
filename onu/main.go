package main

import (
	"context"
	"fmt"
	"net"
	"sync"

	pb "github.com/GonzaloDiaz300/LAB_2/proto"
	"google.golang.org/grpc"
)

type onu struct {
	pb.UnimplementedIntercambiosServer
}

var oms string = "localhost:50051"
var flag bool

func consultarOMS(ip_oms string, estado string, wg *sync.WaitGroup) {
	defer wg.Done()
	// Aquí puedes implementar la lógica para enviar el mensaje al servidor especificado
	conn, err := grpc.Dial(ip_oms, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error al conectar con %s: %v\n", ip_oms, err)
		return
	}
	defer conn.Close()
	serviceClient := pb.NewIntercambiosClient(conn)
	res, err := serviceClient.Nombres(context.Background(), &pb.ONUReq{Id: int32(id)})
	if err != nil {
		panic("No se llego el mensaje " + err.Error())
	}
	if res != nil && !flag {
		flag = true
	}
}

func main() {
	//Set puerto
	listener, err := net.Listen("tcp", ":50052")

	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}

	serv := grpc.NewServer()
	fmt.Printf("Servidor ONU Activo\n")
	pb.RegisterIntercambiosServer(serv, &onu{})
	if err = serv.Serve(listener); err != nil {
		panic("cannot initialize the server" + err.Error())
	}

	var wg sync.WaitGroup
	//Se envía la id a los DataNodes
	wg.Add(1)
	go consultarOMS(oms, "infectados", &wg)
	wg.Wait()
}
