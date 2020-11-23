package main
import(
    	//"os"
        //"strings"
        //"io"
        //"encoding/csv"
        "log"
        "fmt"
		//"time"
		"net"
        "golang.org/x/net/context"
        "google.golang.org/grpc"
		pb "github.com/432i/T2SistemasDistribuidos/dependencias/serverclidn"
)

type Server struct {
	pb.UnimplementedChatCliDnServer
}

func (s *Server) ChunkaDN(ctx context.Context, chunkcito *pb.Chunk) (*pb.Message, error) {
	/*
	// write to disk
	fileName := "./out/" + in.GetFileName() + "_part_" + string(in.GetChunkPart())
	_, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tempBook := books{
		name:   in.GetFileName(),
		parts:  in.GetPart(),
		stored: 0,
	}
	storeInLibrary(tempBook)

	// write/save buffer to disk
	ioutil.WriteFile(fileName, []byte(in.GetChunk()), os.ModeAppend)

	//fmt.Println("Split to : ", fileName)
	*/
	msj := pb.Message{
		Body: "ok",
	}
	return &msj, nil
}

func serverDN1() {
	//--------------------------------------------------------------> Server1
	fmt.Print("Creando conexion...")
	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("failed to listen2: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChatCliDnServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve s2: %v", err)
	}
}

func main(){
	fmt.Print("Creando conexion...")
	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("failed to listen2: %v", err)
	}
	fmt.Println("hola")
	s := grpc.NewServer()
	pb.RegisterChatCliDnServer(s, &Server{})
	fmt.Println("hola2")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve s2: %v", err)
	}
	fmt.Println("hola3")
	//go serverDN1()
}
