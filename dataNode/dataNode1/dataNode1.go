package main
import(
    	"os"
		"strings"
		"math/rand"
		"log"
		"io"
		"io/ioutil"
		"strconv"
        "fmt"
		"time"
		"net"
        "golang.org/x/net/context"
        "google.golang.org/grpc"
		pb "github.com/432i/T2SistemasDistribuidos/dependencias/serverclidn"
)

type Server struct {
	pb.UnimplementedChatCliDnServer
}

var mi_ip = "10.6.40.149"
var cola_chunks_de_cliente []pb.Chunk
var cola_espera []string
var estado = "liberada"
var timestamp = " "
var tipoAlgoritmo = "" //centralizado o distribuido

/*
Funcion: Find
Parametro:
    - slice: arreglo de string, contiene las peticiones en cola recibidas de la forma <tiempo>_DN<i>, donde i es el DN que hizo la peticion
    - val: string a buscar
Descripcion:
	- Busca si una peticion se encuentra en la cola
Retorno:
	- Retorna verdadero si ya se encuentra, o falso cuando se puede insertar la peticion
*/
func Find(slice []string, val string) (int, bool) {
    for i, item := range slice {
        if item == val {
            return i, true
        }
    }
    return -1, false
}

/*
Funcion: escribirLogNN
Parametro:
    - nombre, cantPartes, parte, ip: strings con la informacion para escribir en el log
Descripcion:
	- Genera la conexion con el NameNode, el mensaje que contiene la informacion de un chunk y solicita al NN que lo escriba
Retorno:
	- No hay
*/
func escribirLogNN(nombre string, cantPartes string, parte string, ip string) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.40.152:50001", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Error al establecer conexion con el NameNode")
	}
	defer conn.Close() 
	c2 := pb.NewChatCliDnClient(conn)
	
	fmt.Println("Conexion realizada correctamente con el Name Node de IP 10.6.40.152")

	msj := pb.Message{
		Body: nombre + " " + cantPartes + " " + parte + " " + ip,
	}
	response, err := c2.EscribirLog(context.Background(), &msj)
    if err != nil{
        fmt.Println("Error al enviar la informacion del chunk para escribir en el log")
	}
    log.Printf("%s", response.Body)
}

/*
Funcion: almacenarChunk
Parametro:
    - chunkcito: Contiene la info de un Chunk
Descripcion:
	- Genera el archivo del chunk y escribe la info que contiene
Retorno:
	- No hay
*/
func almacenarChunk(chunkcito pb.Chunk) {
	fileName := chunkcito.GetNombreLibro() + "#" + chunkcito.GetParte()
	_, err := os.Create(fileName)
	if err != nil {
			fmt.Println("Error al crear el archivo del chunk")
			os.Exit(1)
	}
	ioutil.WriteFile(fileName, chunkcito.GetDatos(), os.ModeAppend)
}

/*
Funcion: propuestaEntreTres
Parametro:
    - c2, c3: Conexiones del DN con los otros dos
Descripcion:
	- Le envía al DN correspondiente el chunk (o a sí mismo no sortea), dependiendo del sorteo que se hace
Retorno:
	- No hay
*/
func propuestaEntreTres(c2 pb.ChatCliDnClient, c3 pb.ChatCliDnClient) {
	/*msg2 := pb.Message {
		Body: "m",
	}
	msg3 := pb.Message {
		Body: "m",
	}*/
	i, _ := strconv.Atoi(cola_chunks_de_cliente[0].GetParte())
	chunkcito := pb.Chunk {
		NombreLibro: cola_chunks_de_cliente[0].GetNombreLibro(),
		TotalPartes: cola_chunks_de_cliente[0].GetTotalPartes(),
		Parte: cola_chunks_de_cliente[0].GetParte(),
		Datos: cola_chunks_de_cliente[0].GetDatos(),
		Algoritmo: cola_chunks_de_cliente[0].GetAlgoritmo(),
	}

	if cola_chunks_de_cliente[0].GetParte() == "1" {
		almacenarChunk(chunkcito)
		escribirLogNN(chunkcito.GetNombreLibro(), chunkcito.GetTotalPartes(), chunkcito.GetParte(), "10.6.40.149")
	}
	if cola_chunks_de_cliente[0].GetParte() == "2" {
		msg2, _ := c2.ChunkEntreDN(context.Background(), &chunkcito)
		fmt.Println(msg2.Body)
		escribirLogNN(chunkcito.GetNombreLibro(), chunkcito.GetTotalPartes(), chunkcito.GetParte(), "10.6.40.150")
	}
	if cola_chunks_de_cliente[0].GetParte() == "3" {
		msg3, _ := c3.ChunkEntreDN(context.Background(), &chunkcito)
		fmt.Println(msg3.Body)
		escribirLogNN(chunkcito.GetNombreLibro(), chunkcito.GetTotalPartes(), chunkcito.GetParte(), "10.6.40.151")
	}
	if i > 3 {
		j := rand.Intn(3)
		if j == 0 {
			almacenarChunk(chunkcito)
			escribirLogNN(chunkcito.GetNombreLibro(), chunkcito.GetTotalPartes(), chunkcito.GetParte(), "10.6.40.149")
		} else if j == 1 {
			msg2, _ := c2.ChunkEntreDN(context.Background(), &chunkcito)
			fmt.Println(msg2.Body)
			escribirLogNN(chunkcito.GetNombreLibro(), chunkcito.GetTotalPartes(), chunkcito.GetParte(), "10.6.40.150")
		} else {
			msg3, _ := c3.ChunkEntreDN(context.Background(), &chunkcito)
			fmt.Println(msg3.Body)
			escribirLogNN(chunkcito.GetNombreLibro(), chunkcito.GetTotalPartes(), chunkcito.GetParte(), "10.6.40.151")
		}
	}
}

/*
Funcion: propuestaEntreDos
Parametro:
    - c: Conexion con un DN dado
Descripcion:
	- Le envía al DN correspondiente el chunk (o a sí mismo no sortea), dependiendo del sorteo que se hace
Retorno:
	- No hay
*/
func propuestaEntreDos(c pb.ChatCliDnClient) {
	/*msg := pb.Message {
		Body: "m",
	}*/
	i, _ := strconv.Atoi(cola_chunks_de_cliente[0].GetParte())
	chunkcito := pb.Chunk {
		NombreLibro: cola_chunks_de_cliente[0].GetNombreLibro(),
		TotalPartes: cola_chunks_de_cliente[0].GetTotalPartes(),
		Parte: cola_chunks_de_cliente[0].GetParte(),
		Datos: cola_chunks_de_cliente[0].GetDatos(),
		Algoritmo: cola_chunks_de_cliente[0].GetAlgoritmo(),
	}

	if cola_chunks_de_cliente[0].GetParte() == "1" {
		almacenarChunk(chunkcito)
		escribirLogNN(chunkcito.GetNombreLibro(), chunkcito.GetTotalPartes(), chunkcito.GetParte(), "10.6.40.149")
	}
	if cola_chunks_de_cliente[0].GetParte() == "2" {
		msg, _ := c.ChunkEntreDN(context.Background(), &chunkcito)
		el_split := strings.Split(msg.Body, "#")
		fmt.Println(el_split[0])
		escribirLogNN(chunkcito.GetNombreLibro(), chunkcito.GetTotalPartes(), chunkcito.GetParte(), el_split[1])
	}
	if i > 2 {
		j := rand.Intn(2)
		if j == 0 {
			almacenarChunk(chunkcito)
		} else {
			msg, _ := c.ChunkEntreDN(context.Background(), &chunkcito)
			el_split := strings.Split(msg.Body, "#")
			fmt.Println(el_split[0])
			escribirLogNN(chunkcito.GetNombreLibro(), chunkcito.GetTotalPartes(), chunkcito.GetParte(), el_split[1])
		}
	}
}

/*
Funcion: generarPropuesta
Parametro:
	- cantPartes: string que indica cuantos chunks tiene un libro
Descripcion:
	- Intenta generar conexiones con los demas DNs, si alguno falla se detecta y se genera un nuevo plan de accion para poder distribuir los chunks entre los DNs que esten vivos
Retorno:
	- No hay
*/
func generarPropuesta(cantPartes string) {
	var se_pudo2, se_pudo3 bool
	se_pudo2 = true
	se_pudo3 = true
	partes, _:= strconv.Atoi(cantPartes)
	//var connDN2, connDN3 *grpc.ClientConn

	i := 0
	for i < partes {
		mensajito := pb.Message {
			Body: timestamp + "_DN1",
		}
		//entrarZona := false
		connDN2, _ := grpc.Dial("10.6.40.150:50001", grpc.WithInsecure())
		defer connDN2.Close()
		c2 := pb.NewChatCliDnClient(connDN2)
		fun2, errFunc2 := c2.MaquinaFunciona(context.Background(), &mensajito)
		if errFunc2 != nil {
			se_pudo2 = false
		} else {
			fmt.Println(fun2.Body)
			fmt.Println("Conexion realizada correctamente con el Data Node de IP 10.6.40.150")
		}

		connDN3, _ := grpc.Dial("10.6.40.151:50001", grpc.WithInsecure())
		defer connDN3.Close()
		c3 := pb.NewChatCliDnClient(connDN2)
		fun3, errFunc3 := c3.MaquinaFunciona(context.Background(), &mensajito)
		if errFunc3 != nil {
			se_pudo3 = false
		} else {
			fmt.Println(fun3.Body)
			fmt.Println("Conexion realizada correctamente con el Data Node de IP 10.6.40.151")
		}
		
		if (se_pudo2 == true && se_pudo3 == true) {
			if estado == "buscada" {
				msj2, _ := c2.EnviarPeticion(context.Background(), &mensajito)
				for msj2.Body != "ok" {
					msj2, _ = c2.EnviarPeticion(context.Background(), &mensajito)
				}
				fmt.Println(msj2.Body)
				msj3, _ := c3.EnviarPeticion(context.Background(), &mensajito)
				for msj3.Body != "ok" {
					msj3, _ = c3.EnviarPeticion(context.Background(), &mensajito)
				}
				fmt.Println(msj3.Body)
				estado = "tomada"
				propuestaEntreTres(c2, c3)
			}
			if estado == "tomada" {
				propuestaEntreTres(c2, c3)
			}
		} else if (se_pudo2 == true && se_pudo3 == false) {
			if estado == "buscada" {
				msj2, _ := c2.EnviarPeticion(context.Background(), &mensajito)
				for msj2.Body != "ok" {
					msj2, _ = c2.EnviarPeticion(context.Background(), &mensajito)
				}
				fmt.Println(msj2.Body)
				estado = "tomada"
				propuestaEntreDos(c2)
			}
			if estado == "tomada" {
				propuestaEntreDos(c2)
			}
		} else if (se_pudo2 == false && se_pudo3 == true) {
			if estado == "buscada" {
				msj3, _ := c3.EnviarPeticion(context.Background(), &mensajito)
				for msj3.Body != "ok" {
					msj3, _ = c3.EnviarPeticion(context.Background(), &mensajito)
				}
				fmt.Println(msj3.Body)
				estado = "tomada"
				propuestaEntreDos(c3)
			}
			if estado == "tomada" {
				propuestaEntreDos(c3)
			}
		} else {
			estado = "tomada"
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
	estado = "liberada"
}
/*
Funcion: generarPropuestaCentralizado
Parametro:
	- Nada
Descripcion:
	- Le envia propuesta al Name Node, este la acepta o le entrega una nueva y distribuye los chunks
Retorno:
	- No hay
*/
func generarPropuestaCentralizado(cantPartes string, nombreLibro string){
	var conn *grpc.ClientConn
	//var se_pudo2, se_pudo3 bool
	//se_pudo2 = true
	//se_pudo3 = true
	conn, err := grpc.Dial("10.6.40.152:50001", grpc.WithInsecure())
	if err != nil {
			log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := pb.NewChatCliDnClient(conn)

	partes, _:= strconv.Atoi(cantPartes)
	i := 0
	for i < partes{
		propuesta := "DN1#DN2DN3#"+nombreLibro //le enviamos la propuesta inicial donde asumimos q los demas datanodes estan activos
		//Formato: De_Donde_Envia#Demas_Nodos
		msj := pb.Message {
			Body: propuesta,
		}
		response, err := c.PropuestaCentralizado(context.Background(), &msj)
		if err != nil{
				fmt.Println("Error al enviar la propuesta")
		}

		if response.Body == "espera"{
			fmt.Println("Espere mientras otro nodo utiliza el sistema")
			for response.Body == "espera"{
				response, err = c.PropuestaCentralizado(context.Background(), &msj)
				if err != nil{
					fmt.Println("Error al enviar la propuesta")
				}
			}
			fmt.Println("Sistema liberado: ahora se continua")
		}

		if response.Body == "aceptada"{ //se reparte entre los 3 nodos

			connDN2, err2 := grpc.Dial("10.6.40.150:50001", grpc.WithInsecure())
			if err2 != nil {
				fmt.Println("Error al conectar.")
				//se_pudo2 = false
			}
			defer connDN2.Close()
			c2 := pb.NewChatCliDnClient(connDN2)
			fmt.Println("Conexion realizada correctamente con el Data Node de IP 10.6.40.150")

			connDN3, err3 := grpc.Dial("10.6.40.151:50001", grpc.WithInsecure())
			if err3 != nil {
				fmt.Println("Error al conectar.")
				//se_pudo3 = false
			}
			defer connDN3.Close()
			c3 := pb.NewChatCliDnClient(connDN2)
			fmt.Println("Conexion realizada correctamente con el Data Node de IP 10.6.40.151")

			propuestaEntreTres(c2, c3)

		}
		if response.Body == "DN2"{ //se reparte entre 2 y 1

			connDN2, err2 := grpc.Dial("10.6.40.150:50001", grpc.WithInsecure())
			if err2 != nil {
				fmt.Println("Error al conectar.")
				//se_pudo2 = false
			}
			defer connDN2.Close()
			c2 := pb.NewChatCliDnClient(connDN2)
			fmt.Println("Conexion realizada correctamente con el Data Node de IP 10.6.40.150")

			propuestaEntreDos(c2)
		}
		if response.Body == "DN3"{ //se reparte entre 3 y 1

			connDN3, err3 := grpc.Dial("10.6.40.151:50001", grpc.WithInsecure())
			if err3 != nil {
				fmt.Println("Error al conectar.")
				//se_pudo3 = false
			}
			defer connDN3.Close()
			c3 := pb.NewChatCliDnClient(connDN3)
			fmt.Println("Conexion realizada correctamente con el Data Node de IP 10.6.40.151")

			propuestaEntreDos(c3)

		}
		if response.Body == "tu"{ //te dejas todos los chunks

			almacenarChunk(cola_chunks_de_cliente[0])
			
		}
		i += 1
	}
	

}
/*
Funcion: propuestaEntreDos
Parametro:
    - No tiene
Descripcion:
	- Observa la cola de stream de chunks de los libros enviados por un cliente uploader, cuando detecta que se subió el libro comienza con la ejecución del proceso de distribucion y escritura
Retorno:
	- No hay
*/
func escucharListaChunks() {
	//var prop int
	for { 
		if len(cola_chunks_de_cliente) != 0 {
			if tipoAlgoritmo == "distribuido" {
				tiempoactual := time.Now()
				timestamp = tiempoactual.Format("02-01-2006 15:04")
				estado = "buscada"
				generarPropuesta(cola_chunks_de_cliente[0].GetTotalPartes())
			}
			if tipoAlgoritmo == "centralizado" {
				generarPropuestaCentralizado(cola_chunks_de_cliente[0].GetTotalPartes(), cola_chunks_de_cliente[0].GetNombreLibro())
			}
		}
	}
}

/*
Funcion: serverDN1
Parametro:
    - No tiene
Descripcion:
	- Crea la conexion de servidor
Retorno:
	- No hay
*/
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

func (s *Server) MaquinaFunciona(ctx context.Context, msj *pb.Message) (*pb.Message, error) {
	return &pb.Message{Body: "Exito"}, nil
}

/*
Funcion: EnviarPeticion
Parametro:
    - msj: Mensaje que contiene el instante en que el DN emisor genera la petición de escritura junto con su identificador
Descripcion:
	- Un DN receptor chequeara mediante la implementacion del algoritmo de Ricart y Agrawala si acepta o deja en cola la peticion recibida
Retorno:
	- Retorna un mensaje, cuando este es "ok", la peticion se acepta, en cualquier otro caso se manda un string vacio
*/
func (s *Server) EnviarPeticion(ctx context.Context, msj *pb.Message) (*pb.Message, error) {
	if estado == "liberada" {
		msg := pb.Message {
			Body: "ok",
		}
		return &msg, nil
	}
	if estado == "tomada" {
		_, esta := Find(cola_espera, msj.GetBody())
		if esta == false {
			cola_espera = append(cola_espera, msj.GetBody())
		}
		msg := pb.Message {
			Body: "",
		}
		return &msg, nil
	}
	if estado == "buscada" {
		mi_fecha, _ := time.Parse("02-01-2006 15:04", timestamp)
		split_msj := strings.Split(msj.GetBody(), "_") 
		fecha_emisor, _ := time.Parse("02-01-2006 15:04", split_msj[0])
		if fecha_emisor.Before(mi_fecha) {
			msg := pb.Message {
				Body: "ok",
			}
			return &msg, nil
		} else {
			_, esta := Find(cola_espera, msj.GetBody())
			if esta == false {
				cola_espera = append(cola_espera, msj.GetBody())
			}
			msg := pb.Message {
				Body: "",
			}
			return &msg, nil
		}
	}
	msg := pb.Message {
		Body: "Si te retorno esta wea xd po",
	}
	return &msg, nil
}

/*
Funcion: pedirChunk
Parametro:
    - msj: chunk que pide el cliente
Descripcion:
	- Recibe un mensaje con el chunk que quiere el cliente y retorna el chunk
Retorno:
	- Un chunk
*/
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


/*
Funcion: ChunkADN
Parametro:
    - stream: Stream que contiene el libro trozado en chunks
Descripcion:
	- Inserta cada chunk en una cola para que sean procesados
Retorno:
	- Retorna un stream de exito cuando se recibe o un error si falla la recuperacion del stream
*/
func (s *Server) ChunkaDN(stream pb.ChatCliDn_ChunkaDNServer) error {
	for {
		chunk, err := stream.Recv()
		tipoAlgoritmo = chunk.GetAlgoritmo()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Message {
			Body: "Stream recibido",
			})
		}
		if err != nil {
			return err
		}
		cola_chunks_de_cliente = append(cola_chunks_de_cliente, *chunk)
	}
	return nil
}

/*
Funcion: ChunkEntreDN
Parametro:
    - chunkcito: Recibe un chunk con su informacion respectiva
Descripcion:
	- Pide almacenar el chunk en su archivo, llamando a la funcion almacenarChunk()
Retorno:
	- Retorna un mensaje de exito junto a la ip del DN que almaceno el chunk
*/
func (s *Server) ChunkEntreDN(ctx context.Context, chunkcito *pb.Chunk) (*pb.Message, error) {
	almacenarChunk(*chunkcito)
	fmt.Println("Se ha almacenado el chunk:\n    {nombreLibro: " + chunkcito.GetNombreLibro() + ",\n    totalPartes: " + chunkcito.GetTotalPartes() + ",\n    parte: " + chunkcito.GetParte() + "}")

	msj := pb.Message{
		Body: "Chunk recibido y almacenado#" + mi_ip,
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