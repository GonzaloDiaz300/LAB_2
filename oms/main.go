package main

import (
	"context"
	"fmt"
	"net"
	"os"

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

func (a *oms) Notificar(ctx context.Context, in *pb.ContiReq) (*pb.Confirmacion, error) {
	fmt.Printf("Estado Recibido: %s %s %s \n", in.Estado)
	//Se envía la id a los DataNodes
	if "A" <= string(in.Apellido[0]) && string(in.Apellido[0]) <= "M" {
		go enviarMensaje(DataNodes[0], ID_datanode1, in.Nombre, in.Apellido)
		linea := fmt.Sprintf("%d %s %s\n", ID_datanode1, in.Nombre, in.Apellido)
		go escribir_archivo(linea)
		ID_datanode1 += 2
	} else if "N" <= string(in.Apellido[0]) && string(in.Apellido[0]) <= "Z" {
		go enviarMensaje(DataNodes[1], ID_datanode2, in.Nombre, in.Apellido)
		linea := fmt.Sprintf("%d %s %s\n", ID_datanode2, in.Estado)
		go escribir_archivo(linea)
		ID_datanode2 += 2
	}
	return &pb.Confirmacion{Respuesta: 1}, nil
}

func (a *oms) Nombres(ctx context.Context, in *pb.ONUReq) (*pb.ONUResp, error) {
	return &pb.ONUResp{}, nil
}

func escribir_archivo(linea string) {
	// Nombre del archivo que deseas trabajar
	filename := "DATA.txt"

	// Intentar abrir el archivo en modo lectura
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Si el archivo no existía previamente, escribir datos en él
	if _, err = file.Stat(); os.IsNotExist(err) {
		_, err = file.Write([]byte(linea))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Archivo creado exitosamente.")
	} else {
		_, err = file.Write([]byte(linea))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func enviarMensaje(datanode string, id int, nombre string, apellido string) {
	// Aquí puedes implementar la lógica para enviar el mensaje al servidor especificado
	conn, err := grpc.Dial(datanode, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error al conectar con %s: %v\n", datanode, err)
		return
	}
	defer conn.Close()
	serviceClient := pb.NewIntercambiosClient(conn)
	res, err := serviceClient.Guardar(context.Background(), &pb.OMSReq{Id: int32(id), Nombre: nombre, Apellido: apellido})
	if err != nil {
		panic("No se llego el mensaje " + err.Error())
	}
	if res != nil && !flag {
		flag = true
	}
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
