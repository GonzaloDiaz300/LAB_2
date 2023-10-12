package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"strings"

	pb "github.com/GonzaloDiaz300/LAB_2/proto"
	"google.golang.org/grpc"
)

/*
import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	pb "github.com/GonzaloDiaz300/LAB_2/proto"
)
*/

type serverNode struct {
	pb.UnimplementedIntercambiosServer
}

func (a *serverNode) Buscar(ctx context.Context, in *pb.OMSReq) (*pb.DTNResp, error) {
	var id string
	var nombre string
	var apellido string
	fmt.Printf("Se recibió request de OMS para id: %d", in.GetId())
	//Abrir DATA.txt
	archivo, err := os.Open("DATA.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return nil, nil
	}
	defer archivo.Close()

	// Crea un lector para el archivo
	lector := bufio.NewScanner(archivo)

	if lector.Scan() {
		// Lee la línea actual
		linea := lector.Text()

		// Divide la línea en dos partes utilizando " "
		partes := strings.Split(linea, " ")
		// Convierte las partes en enteros
		id = partes[0]
		nombre = partes[1]
		apellido = partes[2]

	}

	fmt.Printf("Respuesta: %s %s (id: %s)", nombre, apellido, id)
	return &pb.DTNResp{Nombre: nombre, Apellido: apellido}, nil
}

func main() {
	//Puerto oms
	listener, err := net.Listen("tcp", ":50057")

	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}

	serv := grpc.NewServer()
	fmt.Printf("ServerNode Activo\n")
	pb.RegisterIntercambiosServer(serv, &serverNode{})
	if err = serv.Serve(listener); err != nil {
		panic("cannot initialize the server" + err.Error())
	}
}
