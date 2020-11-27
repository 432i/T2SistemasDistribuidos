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
        xd := strings.Split(ips,"-")
        return xd
}
//funcion que recibe la solicitud de escribir en el log 
//recibe un string con los datos del chunk que se guardó en tal data node
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
//recibe el nombre del libro y retorna las ips de los data nodes donde estan
func (s *Server) ChunksDirecciones(ctx context.Context, message *pb.Message) (*pb.Message, error){
        nombreLibro := message.GetBody()
        ips := recuperarLibro(nombreLibro)
        msj := pb.Message{
		Body: ips,
	}
	return &msj, nil
}

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

}