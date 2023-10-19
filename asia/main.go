package main

import (
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	pb "github.com/GonzaloDiaz300/LAB_2/proto"
	"google.golang.org/grpc"
)

var servidor = "localhost:50051"

func main() {
	// Aquí puedes implementar la lógica para enviar el mensaje al servidor especificado
	conn, err := grpc.Dial(servidor, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error al conectar con %s: %v\n", servidor, err)
		return
	}
	defer conn.Close()
	// Abre el archivo para lectura
	archivo, err := os.Open("names.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer archivo.Close()
	// Crea un escáner (scanner) para leer el archivo línea por línea.
	scanner := bufio.NewScanner(archivo)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
	}

	// Crear un canal para enviar líneas algorítmicamente.
	lineChannel := make(chan string)

	// Use una goroutine para enviar las primeras 5 líneas.
	go func() {
		for i := 0; i < 5; i++ {
			if scanner.Scan() {
				lineChannel <- scanner.Text()
			}
		}
		for scanner.Scan() {
			time.Sleep(3 * time.Second) // Espera 3 segundos antes de enviar las líneas restantes.
			lineChannel <- scanner.Text()
		}
		close(lineChannel) // Cerramos el canal cuando terminamos de enviar todas las líneas.
	}()

	// Recibe y procesa las líneas desde el canal.
	for line := range lineChannel {
		// Inicializa el generador de números aleatorios con una semilla.
		rand.Seed(time.Now().UnixNano())

		// Genera un número aleatorio entre 0 y 1.
		randomValue := rand.Float64()

		var estado string

		// Asigna el estado en función de los porcentajes.
		if randomValue < 0.55 {
			estado = "INFECTADA"
		} else {
			estado = "MUERTA"
		}
		// Divide la línea en dos partes utilizando "-"
		partes := strings.Split(line, " ")
		serviceClient := pb.NewIntercambiosClient(conn)
		res, err := serviceClient.Notificar(context.Background(), &pb.ContiReq{Nombre: partes[0], Apellido: partes[1], Estado: estado})
		if err != nil {
			panic("No se llego el mensaje " + err.Error())
		}
		if res.Respuesta == 1 {
			fmt.Println("Persona se envio correctamente")
		}
	}
}
