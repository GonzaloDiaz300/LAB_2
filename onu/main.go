package main

import (
	"context"
	"fmt"
	"net"
	"sync"
	"log"
	"os"
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
	res, err := serviceClient.Nombres(context.Background(), &pb.ONUReq{Estado: estado})
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
	fmt.Printf("Servidor ONU Activo en localhost:50052\n")
	pb.RegisterIntercambiosServer(serv, &onu{})
	if err = serv.Serve(listener); err != nil {
		panic("cannot initialize the server" + err.Error())
	}
	conn, err := grpc.Dial(oms, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error al conectar con %s: %v\n", oms, err)
		return
	}
	defer conn.Close()
	serviceClient := pb.NewIntercambiosClient(conn)

	for {
		fmt.Println("Seleccione una opción:")
		fmt.Println("1) Preguntar por infectados")
		fmt.Println("2) Preguntar por muertos")
		fmt.Println("3) Cerrar los servidores")
		fmt.Print("Opción: ")

		var opcion int
		_, err := fmt.Scanf("%d", &opcion)
		if err != nil {
			fmt.Println("Opción no válida")
			continue
		}
		switch opcion {
		case 1:
			respuesta, err := serviceClient.Nombres(context.Background(), &pb.ONUReq{Estado: "INFECTADOS"})
			if err != nil {
				log.Fatalf("Error al solicitar infectados: %v", err)
			}
			fmt.Printf("Respuesta: %s\n", respuesta)
		case 2:
			respuesta, err := serviceClient.Nombres(context.Background(), &pb.ONUReq{Estado: "MUERTOS"})
			if err != nil {
				log.Fatalf("Error al solicitar muertos: %v", err)
			}
			fmt.Printf("Respuesta: %s\n", respuesta)
		case 3:
			fmt.Println("Servidores cerrados. Saliendo...")
			os.Exit(0)
		default:
			fmt.Println("Opción no válida")
		}
	}

}
