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
//le piden un chunk y lo devuelve
func (s *Server) pedirChunk(ctx context.Context, msj *pb.Message) (*pb.Chunk, error) {
	split := strings.Split(msj.GetBody(), "#")
	nombreLibro := split[0]
	parte := split[1]

	//seleccionamos el archivo del directorio
	 //read a chunk
	 newFileChunk, err := os.Open(msj.GetBody())
	 if err != nil {
			 fmt.Println(err)
			 os.Exit(1)
	 }
	 defer newFileChunk.Close()
	 chunkInfo, err := newFileChunk.Stat()
	 if err != nil {
			 fmt.Println(err)
			 os.Exit(1)
	 }
	 // calculate the bytes size of each chunk
	 // we are not going to rely on previous data and constant
	 var chunkSize int64 = chunkInfo.Size()
	 chunkBufferBytes := make([]byte, chunkSize)

	chunko := pb.Chunk{
		NombreLibro: nombreLibro,
		TotalPartes: "0",
		Parte: parte,
		Datos: chunkBufferBytes,
		Algoritmo: "0",
	}

	chunkBufferBytes = nil

	return &chunko, nil
}

//envia los datos del chunk para que se escriban en el log
func escribirLogNN(nombre string, cantPartes string, parte string, ip string, c pb.NewChatCliDnClient){
	msj = Message{
		Body: nombre + " " + cantPartes + " " + parte + " " + ip,
	}
	response, err := c.escribirLog(context.Background(), &msj)
    if err != nil{
        fmt.Println("Error al enviar la informacion del chunk para escribir en el log")
        break
    }
    log.Printf("%s", response.Body)
}

func serverDN1() {
	//-----------------------------------------------------------------> Server2
	fmt.Fprintln("Creando conexión...")
	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("failed to listen2: %v", err)
	}
	fmt.Fprintln("Conexión creada satisfactoriamente")
	s := grpc.NewServer()
	pb.RegisterChatCliDnServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve s2: %v", err)
	}
}

func main(){
	var respuesta string
	go serverDN1()
	fmt.Print("El servidor de Data Node 2 esta activo\n")
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
