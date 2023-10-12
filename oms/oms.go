package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	pb "github.com/GonzaloDiaz300/LAB_2/proto"
	"google.golang.org/grpc"
)

// IP´s de cada servidor:
var servidores = []string{
	"dist077:50051",
	"dist078:50052",
	"dist079:50054",
	"dist080:50056",
}

var DataNodes = []string{
	"dist077:50057",
	"dist078:50058",
}

var onu = "dist079:50059"

// ID´s que se le asignara a cada persona que llegue, comienza en 0 y se le va sumando 1
var ID = 0

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

var flag bool

func enviarMensaje(databnode string, mensaje int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Aquí puedes implementar la lógica para enviar el mensaje al servidor especificado
	conn, err := grpc.Dial(databnode, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error al conectar con %s: %v\n", databnode, err)
		return
	}
	defer conn.Close()
	serviceClient := pb.NewNotificacionClient(conn)
	res, err := serviceClient.OMSReq(context.Background(), &pb.NotiReq{Solicitud: 100})
	if err != nil {
		panic("No se llego el mensaje " + err.Error())
	}
	if res != nil && !flag {
		flag = true
	}
}

func main() {

	var wg sync.WaitGroup
	for _, servidor := range servidores {
		wg.Add(1)
		go enviarMensaje(servidor, numero_llaves, &wg)
	}

	wg.Wait()
}