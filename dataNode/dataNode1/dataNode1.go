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
/*

type ChunkInterno struct {
	NombreLibro string
	TotalPartes string
	Parte string
	Datos bytes;
}

var cola_chunks_de_cliente []ChunkInterno
var ultima_parte int

func almacenarChunk(chunkcito *pb.Chunk) {
	// write to disk
	fileName := chunkcito.nombreLibro + "_" + chunkcito.Parte

	_, err := os.Create(fileName)
	if err != nil {
			fmt.Println("Error al crear el archivo del chunk")
			os.Exit(1)
	}

	// write/save buffer to disk
	ioutil.WriteFile(fileName, chunkcito.Datos, os.ModeAppend)
}

func (s *Server) ChunkaDN(ctx context.Context, chunkcito *pb.Chunk) (*pb.Message, error) {
	trozo := ChunkInterno {
		NombreLibro: chunkcito.GetNombreLibro,
		TotalPartes: chunkcito.GetTotalPartes,
		Parte: chunkcito.GetParte,
		Datos: chunkcito.GetDatos,
	}
	cola_chunks_de_cliente = append(cola_chunks_de_cliente, trozo)
	
	if chunkcito.Parte == chunkcito.TotalPartes {
		ultima_parte = 1
	} else {
		ultima_parte = 0
	}

	msj := pb.Message{
		Body: "Chunk recibido:\n    {nombreLibro: " + trozo.nombreLibro + ",\n    totalPartes: " + trozo.totalPartes + ",\n    parte: " + trozo.parte + ",\n    datos: " + trozo.datos + "}",
	}
	return &msj, nil
}*/
func (s *Server) ChunkaDN(stream pb.Chunk) error {

	for {
	  chunk, err := stream.Recv()
	  if err == io.EOF {
		return stream.SendAndClose(&pb.Message{
		  body: "ok",
		})
	  }
	  if err != nil {
		return err
	  }
	}
}

func propuesta() {
	var proposal;
	if chunkcito.totalPartes == 3 {
		proposal = 0
	} else if chunkcito.totalPartes < 3 {
		proposal = 1
	} else {
		proposal = 2
	}
	return proposal
}

func generarPropuesta(prop int, direccion string) {
	var connDN2 *grpc.ClientConn
	connDN2, err := grpc.Dial(direccion, grpc.WithInsecure())
	if err != nil {
			return "no"
	}
	defer conn.Close()
	c := pb.NewChatCliDnClient(connDN2)
	fmt.Println("Conexion realizada correctamente con el Data Node de IP "+ip+"\n")
	return c
}

func chunksRecibidos() {
	var prop int
	for { 
		if ultima_parte == 1 {
			prop = propuesta()
			generarPropuesta(prop, "10.6.40.150:50001")
		}
	}
}

func serverDN1() { //Comunicacion con cliente
	//-----------------------------------------------------------------> Server1
	fmt.Println("Creando conexión...")
	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("failed to listen2: %v", err)
	}
	fmt.Println("Conexión creada satisfactoriamente")
	s := grpc.NewServer()
	pb.RegisterChatCliDnServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve s2: %v", err)
	}
}

func main(){
	fmt.Println("El servidor de Data Node 1")
	var respuesta string
	go serverDN1()
	go chunksRecibidos()
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