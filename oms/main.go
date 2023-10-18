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
	fmt.Printf("Estado Recibido: %s %s %s \n", in.Nombre, in.Apellido, in.Estado)
	//Se envía la id a los DataNodes
	if "A" <= string(in.Apellido[0]) && string(in.Apellido[0]) <= "M" {
		go enviarMensaje(DataNodes[0], ID_datanode1, in.Nombre, in.Apellido)
		linea := fmt.Sprintf("%d 1 %s", ID_datanode1, in.Estado)
		go escribir_archivo(linea)
		ID_datanode1 += 2
	} else if "N" <= string(in.Apellido[0]) && string(in.Apellido[0]) <= "Z" {
		go enviarMensaje(DataNodes[1], ID_datanode2, in.Nombre, in.Apellido)
		linea := fmt.Sprintf("%d 2 %s", ID_datanode1, in.Estado)
		go escribir_archivo(linea)
		ID_datanode2 += 2
	}
	return &pb.Confirmacion{Respuesta: 1}, nil
}

func (a *oms) Nombres(ctx context.Context, in *pb.ONUReq) (*pb.ONUResp, error) {
	return &pb.ONUResp{Nombre: "Gonzalo", Apellido: "Diaz"}, nil
}

func escribir_archivo(linea string) {
    rutaCompleta := "oms/DATA.txt" // Utilizando "/" como separador de ruta

    archivo, err := os.OpenFile(rutaCompleta, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        fmt.Println("Error al abrir el archivo:", err)
        return
    }
    defer archivo.Close()

    _, err = fmt.Fprintln(archivo, linea)
    if err != nil {
        fmt.Println("Error al escribir en el archivo:", err)
        return
    }
    fmt.Println("Datos escritos en el archivo exitosamente.")
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
	if res.Respuesta == 1{
		fmt.Println("Persona se envio correctamente")
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
