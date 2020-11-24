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

func recuperarLibro(nombreLibro string) string{
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
                if flag ==1{
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
                        split := strings.Split(linea, "_")
                        split = strings.Split(split[1], " ")
                        cantPartes = split[1]
                        flag=1
                }
        }
    
        if err := scanner.Err(); err != nil {
                log.Fatal(err)
        }

        
        return ips
}

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


var numLibros = 1
func escribirLog(nombre string, cantPartes string, parte string, ip string, flag int){
        archivo, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY, 0644)
        if err != nil {
                fmt.Println(err)
                return
        }
        if flag == 1{ //primera vez que se escribe el libro
                str := nombre+"_"+strconv.Itoa(numLibros)+" "+cantPartes+"_"+strconv.Itoa(numLibros)+"\n"
                _, err := archivo.WriteString(str)
                if err != nil {
                        fmt.Println(err)
                        archivo.Close()
                return
                }
        }else{ //se guardan sus partes
        str := "parte_"+strconv.Itoa(numLibros)+"_"+parte+" "+ ip+"\n"
        _, err := archivo.WriteString(str)
        if err != nil {
                fmt.Println(err)
                archivo.Close()
                return
        }
      }
}

func main(){
        crearTxt()

        nombresLibros := []string{"Foo", "Bar", "ETHIEL"}
        ips := []string{"1.1.1.1", "2.2.2.2", "3.3.3.3"}
        for _, s := range nombresLibros { 
                escribirLog(s, "3", "1", "1.66.6", 1)
                for i := 0; i < 3; i++ {
                        escribirLog(s, "3", strconv.Itoa(i+1), ips[i], 0)

	        }
        numLibros += 1
        }
}