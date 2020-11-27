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

var cola_chunks_de_cliente []pb.Chunk

func almacenarChunk(chunkcito pb.Chunk) {
	fileName := chunkcito.nombreLibro + "#" + chunkcito.Parte
	_, err := os.Create(fileName)
	if err != nil {
			fmt.Println("Error al crear el archivo del chunk")
			os.Exit(1)
	}
	ioutil.WriteFile(fileName, chunkcito.Datos, os.ModeAppend)
}

func propuestaEntreTres(c2 pb.NewChatCliDnClient, c3 pb.NewChatCliDnClient) {
	msg2 := pb.Message {
		body: "m",
	}
	msg3 := pb.Message {
		body: "m",
	}
	i, _ := strconv.Atoi(cola_chunks_de_cliente[0].parte)
	chunkcito := pb.Chunk {
		nombreLibro: cola_chunks_de_cliente[0].nombreLibro,
		totalPartes: cola_chunks_de_cliente[0].totalPartes,
		parte: cola_chunks_de_cliente[0].parte,
		datos: cola_chunks_de_cliente[0].datos,
		algoritmo: cola_chunks_de_cliente[0].algoritmo,
	}

	if cola_chunks_de_cliente[0].parte == "1" {
		almacenarChunk(chunkcito)
	}
	if cola_chunks_de_cliente[0].parte == "2" {
		msg2, _ = c2.ChunkEntreDN(context.Background(), &chunkcito)
		fmt.Println(msg2.body)
	}
	if cola_chunks_de_cliente[0].parte == "3" {
		msg3, _ = c3.ChunkEntreDN(context.Background(), &chunkcito)
		fmt.Println(msg3.body)
	}
	if i > 3 {
		j := rand.Intn(3)
		if j == 0 {
			almacenarChunk(chunkcito)
		} else if j == 1 {
			msg2, _ = c2.ChunkEntreDN(context.Background(), chunkcito)
			fmt.Println(msg2.body)
		} else {
			msg3, _ = c3.ChunkEntreDN(context.Background(), chunkcito)
			fmt.Println(msg3.body)
		}
	}
}

func propuestaEntreDos(c pb.NewChatCliDnClient) {
	msg := pb.Message {
		body: "m",
	}
	i, _ := strconv.Atoi(cola_chunks_de_cliente[0].parte)
	chunkcito := pb.Chunk {
		nombreLibro: cola_chunks_de_cliente[0].nombreLibro,
		totalPartes: cola_chunks_de_cliente[0].totalPartes,
		parte: cola_chunks_de_cliente[0].parte,
		datos: cola_chunks_de_cliente[0].datos,
		algoritmo: cola_chunks_de_cliente[0].algoritmo,
	}

	if cola_chunks_de_cliente[0].parte == "1" {
		almacenarChunk(chunkcito)
	}
	if cola_chunks_de_cliente[0].parte == "2" {
		msg, _ = c.ChunkEntreDN(context.Background(), &chunkcito)
		fmt.Println(msg.body)
	}
	if i > 2 {
		j := rand.Intn(2)
		if j == 0 {
			almacenarChunk(chunkcito)
		} else {
			msg, _ = c.ChunkEntreDN(context.Background(), chunkcito)
			fmt.Println(msg.body)
		}
	}
}

func generarPropuesta(cantPartes string) {
	var se_pudo2, se_pudo3 bool
	se_pudo2 = true
	se_pudo3 = true
	partes, _:= strconv.Atoi(cantPartes)
	var connDN2, connDN3 *grpc.ClientConn
	i := 0
	for i < partes {
		connDN2, err2 := grpc.Dial(direccion, grpc.WithInsecure())
		if err2 != nil {
			se_pudo2 = false
		}
		defer connDN2.Close()
		c2 := pb.NewChatCliDnClient(connDN2)
		fmt.Println("Conexion realizada correctamente con el Data Node de IP 10.6.40.150")

		connDN3, err3 := grpc.Dial(direccion, grpc.WithInsecure())
		if err3 != nil {
			se_pudo3 = false
		}
		defer connDN3.Close()
		c3 := pb.NewChatCliDnClient(connDN2)
		fmt.Println("Conexion realizada correctamente con el Data Node de IP 10.6.40.151")
		
		if (se_pudo2 == true && se_pudo3 == true) {
			propuestaEntreTres(c2, c3, prop)
		} else if (se_pudo2 == true && se_pudo3 == false) {
			propuestaEntreDos(c2, prop)
		} else if (se_pudo2 == false && se_pudo3 == true) {
			propuestaEntreDos(c3, prop)
		} else {
			almacenarChunk(cola_chunks_de_cliente[0])
		}

		if len(cola_chunks_de_cliente) == 1 {
			cola_chunks_de_cliente = make([]pb.Chunk, 0)
		} else {
			cola_chunks_de_cliente = cola_chunks_de_cliente[1:]
		}
		se_pudo2 = true
	se_pudo3 = true
		i++
	}
}

func escucharListaChunks() {
	var prop int
	for { 
		if len(cola_chunks_de_cliente) != 0 {
			generarPropuesta(cola_chunks_de_cliente[0].totalPartes)
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

func (s *Server) ChunkaDN(stream pb.Chunk) error {
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Message {
			body: "Stream recibido",
			})
		}
		if err != nil {
			return err
		}
		cola_chunks_de_cliente = append(cola_chunks_de_cliente, chunk)
	}
}

func (s *Server) ChunkEntreDN(ctx context.Context, chunkcito *pb.Chunk) (*pb.Message, error) {
	almacenarChunk(chunkcito)
	fnt.Println("Se ha almacenado el chunk:\n    {nombreLibro: " + chunkcito.nombreLibro + ",\n    totalPartes: " + chunkcito.totalPartes + ",\n    parte: " + chunkcito.parte + ",\n    datos: " + chunkcito.datos + "}",

	msj := pb.Message{
		Body: "Chunk recibido y almacenado:\n    {nombreLibro: " + chunkcito.nombreLibro + ",\n    totalPartes: " + chunkcito.totalPartes + ",\n    parte: " + chunkcito.parte + ",\n    datos: " + chunkcito.datos + "}",
	}
	return &msj, nil
}

func main(){
	fmt.Println("El servidor de Data Node 1")
	var respuesta string
	go serverDN1()
	go escucharListaChunks()
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