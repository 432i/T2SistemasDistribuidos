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
	
	msj := pb.Message{
		Body: "ok",
	}
	return &msj, nil
}
//sirve para obtener la ip de la maquina
func obtenerIP() string{
	ipDN="10.6.40.150"
	return ipDN
}
//envia los datos del chunk para que se escriban en el log
func escribirLogNN(nombre string, cantPartes string, parte string, ip string, c pb.NewChatCliDnClient){
	msj = Message{
		Body: nombre+" "+cantPartes+" "+parte+" "+ip,
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