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
        pb "github.com/432i/T2SistemasDistribuidos/dependencias/serverclidn"
)

type Server struct {
	pb.UnimplementedChatCliDnServer
}
/*
Funcion: recuperarLibro
Parametro:
	- nombreLibro: nombre del libro que se quiere sacar de log.txt
Descripcion:
	- Busca en el log.txt el libro y recupera sus partes y las ip en donde se encuentran
Retorno:
	- Un slice con las ip en donde estan las partes del libro
*/
func recuperarLibro(nombreLibro string) []string{
        ips := ""
        flag := 0
        cont :=1
        cantPartes :=""
        file, err := os.Open("log.txt")
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()
    
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            //fmt.Println(scanner.Text())
            linea := scanner.Text()
            if flag == 1{
                split := strings.Split(linea, " ")
                ip := split[1]
                ips = ips+ip+" - "
                i, err := strconv.Atoi(cantPartes)
                
                if err != nil {
                        fmt.Println(err)
                }
        
                if cont ==i{
                        break
                }
                cont+=1
            }
            if true==strings.Contains(linea, nombreLibro){
                split := strings.Split(linea, " ")
                cantPartes = split[1]
                flag=1
            }
        }
    
        if err := scanner.Err(); err != nil {
                log.Fatal(err)
        }
        final := strings.Split(ips,"-")
        return final
}
/*
Funcion: pedirCatalogo
Parametro:
	- Nada
Descripcion:
	- Genera un string con todos los libros disponibles en log.txt (catalogo)
Retorno:
	- Retorna un string con el catalogo de libros disponibles
*/
func pedirCatalogo() string{
        catalogo := "Libros disponibles: \n"
        file, err := os.Open("log.txt")
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()
    
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                linea := scanner.Text()
                //fmt.Println(scanner.Text())
                if strings.Contains(linea, "parte_") && strings.Contains(linea, "."){
                        catalogo = catalogo
                }else{
                        split := strings.Split(linea, " ")
                        nombreLibro := split[0]
                        catalogo = catalogo + nombreLibro + "\n"
                }
            
        }
    
        if err := scanner.Err(); err != nil {
                log.Fatal(err)
        }

        return catalogo
}
/*
Funcion: pedirCatalogo
Parametro:
	- Nada
Descripcion:
	- Pide el catalogo a la funcion pedirCatalogo y lo devuelve al cliente
Retorno:
	- Un Message con el catalogo
*/
func (s *Server) pedirCatalogo(ctx context.Context, message *Message) (*Message, error){
        catalogo := obtenerCatalogo()
        msj := pb.Message{
		Body: catalogo,
	}
        return &msj, nil

}
/*
Funcion: propuestaCentralizado
Parametro:
	- Message: con el nodo que envia la propuesta y su propuesta
Descripcion:
	- Recibe la propuesta del data node emisor y la acepta o propone otra
Retorno:
	- Message con la respuesta
*/
func (s *Server) propuestaCentralizado(ctx context.Context, message *Message) (*Message, error){
        var se_pudo2, se_pudo3, se_pudo1 bool
        se_pudo1 = true
        se_pudo2 = true
        se_pudo3 = true

        split := strings.Split(message.GetBody(), "#")
        emisor := split[0]                              //de donde se envia la propuesta

        connDN1, err1 := grpc.Dial("10.6.40.149", grpc.WithInsecure())
        if err1 != nil {
                se_pudo1 = false
        }
        defer connDN1.Close()
        c1 := pb.NewChatCliDnClient(connDN1)
        
        connDN2, err2 := grpc.Dial("10.6.40.150", grpc.WithInsecure())
        if err2 != nil {
                se_pudo2 = false
        }
        defer connDN2.Close()
        c2 := pb.NewChatCliDnClient(connDN2)

        connDN3, err3 := grpc.Dial("10.6.40.151", grpc.WithInsecure())
        if err3 != nil {
                se_pudo3 = false
        }
        defer connDN3.Close()
        c3 := pb.NewChatCliDnClient(connDN3)

        if emisor == "DN1"{

                if se_pudo2 && se_pudo3{
                        msj := pb.Message{
                                Body: "aceptada",
                        }
                        return &msj, nil
                }
                if se_pudo2 == true && se_pudo3 == false{
                        msj := pb.Message{
                                Body: "DN2",
                        }
                        return &msj, nil
                }
                if se_pudo3 == true && se_pudo2 == false{
                        msj := pb.Message{
                                Body: "DN3",
                        }
                        return &msj, nil
                }else{
                        msj := pb.Message{
                                Body: "tu",
                        }
                        return &msj, nil
                }
        }

        if emisor == "DN2"{

                if se_pudo1 && se_pudo3{
 
                        msj := pb.Message{
                                Body: "aceptada",
                        }
                        return &msj, nil
                }
                if se_pudo1 == true && se_pudo3 == false{
                        msj := pb.Message{
                                Body: "DN1",
                        }
                        return &msj, nil
                }
                if se_pudo3 == true && se_pudo1 == false{
                        msj := pb.Message{
                                Body: "DN3",
                        }
                        return &msj, nil
                }else{
                        msj := pb.Message{
                                Body: "tu",
                        }
                        return &msj, nil
                }
        }
        
        if emisor == "DN3"{

                if se_pudo2 && se_pudo1{
 
                        msj := pb.Message{
                                Body: "aceptada",
                        }
                        return &msj, nil
                }
                if se_pudo2 == true && se_pudo1 == false{
                        msj := pb.Message{
                                Body: "DN2",
                        }
                        return &msj, nil
                }
                if se_pudo1 == true && se_pudo2 == false{
                        msj := pb.Message{
                                Body: "DN1",
                        }
                        return &msj, nil
                }else{
                        msj := pb.Message{
                                Body: "tu",
                        }
                        return &msj, nil
                }
        }
        
}

/*
Funcion: escribirLog
Parametro:
	- Message: string con la informacion del chunk a guardar en el log
Descripcion:
	- Recibe la solicitud de escribir en el log con los datos del chunk a escribir
Retorno:
	- String de exito
*/
func (s *Server) escribirLog(ctx context.Context, message *Message) (*Message, error){
        split := strings.Split(message.GetBody(), " ")
        nombreLibro := split[0]
        cantPartes := split[1]
        parte := split[2]
        ip :=  split[3]

        if parte ==1{
                escribirTXT(nombreLibro, cantPartes, parte, ip, 1)
        }else{
                escribirTXT(nombreLibro, cantPartes, parte, ip, 0)
        }
        msj := Message{
                Body: "Chunk guardado en el log correctamente",
        }
        return &msj, nil
}
/*
Funcion: ChunksDirecciones
Parametro:
	- Message: string con el nombre del libro a recuperar las ip de sus partes
Descripcion:
	- Recibe el nombre del libro, busca sus partes e ips en el log.txt y las devuelve
Retorno:
	- String con las ip de las partes
*/
func (s *Server) ChunksDirecciones(ctx context.Context, message *pb.Message) (*pb.Message, error){
        nombreLibro := message.GetBody()
        ips := recuperarLibro(nombreLibro)
        msj := pb.Message{
		Body: ips,
	}
	return &msj, nil
}
/*
Funcion: crearTxt
Parametro:
	- Nada
Descripcion:
	- Crea el archivo log.txt
Retorno:
	- Nada
*/
func crearTxt(){
        archivo, err := os.Create("log.txt")
        if err != nil {
            fmt.Println(err)
            return
        }
        err = archivo.Close()
        if err != nil {
            fmt.Println(err)
            return
        }
}
/*
Funcion: escribirTXT
Parametro:
        - nombre, cantPartes, parte, ip: informacion del chunk a guardar
        - flag: 1 si es primera vez que se escribe el libro y 0 si no
Descripcion:
	- Guarda la info del libro y sus chunks en el txt
Retorno:
	- Nada
*/
func escribirTXT(nombre string, cantPartes string, parte string, ip string, flag int){
        archivo, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY, 0644)
        if err != nil {
                fmt.Println(err)
                return
        }
        if flag == 1{ //primera vez que se escribe el libro
                str := nombre+" "+cantPartes+"\n"
                _, err := archivo.WriteString(str)
                str = "parte_"+parte+" "+ ip+"\n"
                _, err = archivo.WriteString(str)
                if err != nil {
                fmt.Println(err)
                archivo.Close()
                return
                }
        }else{ //se guardan sus partes
                str := "parte_"+parte+" "+ ip+"\n"
                _, err := archivo.WriteString(str)
                if err != nil {
                        fmt.Println(err)
                        archivo.Close()
                return
                }
        }
}
/*
Funcion: serverNN
Parametro:
	- Nada
Descripcion:
	- Levanta el servidor de name node
Retorno:
	- Nada
*/
func serverNN() {
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
        crearTxt()
        var respuesta string
        go serverNN()
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