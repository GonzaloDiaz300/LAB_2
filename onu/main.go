package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/GonzaloDiaz300/LAB_2/proto"
	"google.golang.org/grpc"
)

type onu struct {
	pb.UnimplementedIntercambiosServer
}

var oms string = "dist077:50051"
var flag bool

func consultarOMS(ip_oms string, estado string) {
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
	conn, err := grpc.Dial(oms, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error al conectar con %s: %v\n", oms, err)
		return
	}
	defer conn.Close()
	serviceClient := pb.NewIntercambiosClient(conn)
	for {
		fmt.Println("\nSeleccione una opción:")
		fmt.Println("1) Preguntar por infectados")
		fmt.Println("2) Preguntar por muertos")
		fmt.Println("Opción: ")

		var opcion int
		_, err := fmt.Scanf("%d\n", &opcion)

		if err != nil {
			fmt.Println("Opción no válida")
			continue
		}
		switch opcion {
		case 1:
			respuesta, err := serviceClient.Nombres(context.Background(), &pb.ONUReq{Estado: "INFECTADA"})
			if err != nil {
				log.Fatalf("Error al solicitar infectados: %v", err)
			}
			fmt.Printf("Esperando respuesta...\n")
			for _, nombre := range respuesta.Persona {
				fmt.Println(nombre)
			}
		case 2:
			respuesta, err := serviceClient.Nombres(context.Background(), &pb.ONUReq{Estado: "MUERTA"})
			if err != nil {
				log.Fatalf("Error al solicitar muertos: %v", err)
			}
			fmt.Printf("Esperando respuesta...\n")
			// Itera sobre la lista de personas e imprime los nombres.
			for _, nombre := range respuesta.Persona {
				fmt.Println(nombre)
			}
		default:
			fmt.Println("Opción no válida")
		}
	}
}
