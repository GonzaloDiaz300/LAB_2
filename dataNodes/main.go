package datanodes

import (
	"fmt"
	"net"

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
	pb.UnimplementedNotificacionServer
}

func main() {
	//Puerto oms
	listner, err := net.Listen("tcp", ":50056")

	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}

	serv := grpc.NewServer()
	fmt.Printf("ServerNode Activo\n")
	pb.RegisterNotificacionServer(serv, &serverNode{})
	if err = serv.Serve(listner); err != nil {
		panic("cannot initialize the server" + err.Error())
	}
}
