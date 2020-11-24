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
	//-----------------------------------------------------------------> Server1
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
	var respuesta string
	go serverDN1()
	fmt.Print("El servidor de Data Node 1 esta activo\n")
	fmt.Println("Ingrese 432 y presione Enter para salir del programa")
	for{
		_, err := fmt.Scanln(&respuesta)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		if strings.Compare("432", respuesta) == 0{
			fmt.Println("Saliendo del programa. . . ")
			break
		}
		fmt.Println("que")
	}
}
