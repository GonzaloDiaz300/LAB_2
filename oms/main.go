package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"sync"

	pb "github.com/GonzaloDiaz300/LAB_2/proto"
	"google.golang.org/grpc"
)

type oms struct {
	pb.UnimplementedIntercambiosServer
}

var onu = "localhost:50052"

var DataNodes = []string{
	"localhost:50053",
	"localhost:50054",
}

// IP´s de cada servidor:
var servidores = []string{
	"localhost:50055",
	"localhost:50056",
	"localhost:50057",
	"localhost:50058",
}

var ID_datanode1 = 0 //IDs que se le asignaran al datanode1, que seran siempre numeroos pares
var ID_datanode2 = 1 //IDs que se le asignaran al datanode2, que seran siempre numeroos impares

var flag bool

func (a *oms) Notificar(ctx context.Context, in *pb.ContiReq) (*pb.ContiResp, error) {
	fmt.Printf("Estado Recibido: %s %s %s \n", in.Nombre, in.Apellido, in.Estado)
	return &pb.ContiResp{Respuesta: 1}, nil
}

func enviarMensaje(datanode string, id int, nombre string, apellido string, wg *sync.WaitGroup) {
	defer wg.Done()
	// Aquí puedes implementar la lógica para enviar el mensaje al servidor especificado
	conn, err := grpc.Dial(datanode, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error al conectar con %s: %v\n", datanode, err)
		return
	}
	defer conn.Close()
	serviceClient := pb.NewIntercambiosClient(conn)
	res, err := serviceClient.Buscar(context.Background(), &pb.OMSReq{Id: int32(id), Nombre: nombre, Apellido: apellido})
	if err != nil {
		panic("No se llego el mensaje " + err.Error())
	}
	if res != nil && !flag {
		flag = true
	}
}

func main() {

	//se crea el archivo DATA.txt
	archivo, err := os.Create("DATA.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer archivo.Close()

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

	//Tiene que recibir los nombres de los continentes primero
	var apellido string
	var nombre string

	var wg sync.WaitGroup
	//Se envía la id a los DataNodes
	if "A" <= string(apellido[0]) && string(apellido[0]) <= "M" {
		wg.Add(1)
		go enviarMensaje(DataNodes[0], ID_datanode1, nombre, apellido, &wg)
		linea := fmt.Sprintf("%d %s %s\n", ID_datanode1, nombre, apellido) //Se crea la línea que se escribirá en el txt
		archivo.WriteString(linea)                                         //Se escribe en el txt
		ID_datanode1 += 2
	} else if "N" <= string(apellido[0]) && string(apellido[0]) <= "Z" {
		wg.Add(1)
		go enviarMensaje(DataNodes[1], ID_datanode2, nombre, apellido, &wg)
		linea := fmt.Sprintf("%d %s %s\n", ID_datanode2, nombre, apellido) //Se crea la línea que se escribirá en el txt
		archivo.WriteString(linea)                                         //Se escribe en el txt
		ID_datanode2 += 2
	}
	wg.Wait()
}
