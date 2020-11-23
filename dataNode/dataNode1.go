package main
import(
        "os"
        "strings"
        "io"
        "encoding/csv"
        "log"
        "fmt"
        "time"
        "golang.org/x/net/context"
        "google.golang.org/grpc"
		"github.com/432i/T2SistemasDistribuidos/dependencias/chatclidn"
		pb "connclidn"
)

func serverDN1() {
	//--------------------------------------------------------------> Server1
	fmt.Print("Creando conexion...")
	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("failed to listen2: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChatCliDnServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve s2: %v", err)
	}
}

func main(){
	go serverDN1()
}