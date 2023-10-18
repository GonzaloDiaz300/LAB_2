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

func (a *serverNode) Guardar(ctx context.Context, in *pb.OMSReq) (*pb.Confirmacion, error) {
	//Abrir DATA.txt
	archivo, err := os.Open("DATA.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return nil, nil
	}
	defer archivo.Close()
	linea := fmt.Sprintf("%d %s %s", in.Id, in.Nombre, in.Apellido)
	go escribir_archivo(linea)
	return &pb.Confirmacion{Respuesta: 1}, nil
}

func escribir_archivo(linea string) {
    rutaCompleta := "dataNodes/DATA.txt" // Utilizando "/" como separador de ruta

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


func (a *serverNode) Buscar(ctx context.Context, in *pb.OMSONUReq) (*pb.DTNResp, error) {
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
	//Set puerto
	listener, err := net.Listen("tcp", ":50053")

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
