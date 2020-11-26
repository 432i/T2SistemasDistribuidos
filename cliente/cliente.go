package main
import(
        //"bufio"
        "io/ioutil"
        "math/rand"
        "math"
        "strconv"
        "os"
        "strings"
        //"io"
        //"encoding/csv"
        "log"
        "fmt"
        //"time"
        "golang.org/x/net/context"
        "google.golang.org/grpc"
        pb "github.com/432i/T2SistemasDistribuidos/dependencias/serverclidn"
)
/*
Funcion: conexionDN
Parametro:
	- ip: ip en formato string para hacer conexion
Descripcion:
	- Recibe una ip en formato string y genera una conexion gRPC con ese Data Node
Retorno:
	- Retorna la conexion con el Data Node de la ip recibida
*/
func conexionDN(ip){
        var conn *grpc.ClientConn
        puerto := ":50001"
        conn, err := grpc.Dial(ip+puerto, grpc.WithInsecure())
        if err != nil {
                log.Fatalf("did not connect: %s", err)
        }
        defer conn.Close()
        c := pb.NewChatCliDnClient(conn)
        fmt.Println("Conexion realizada correctamente con el Data Node de IP "+ip+"\n")
        return c
}
/*
Funcion: conexionNN
Parametro:
	- Nada
Descripcion:
	- Realiza la conexion gRPC con el Name Node
Retorno:
	- Retorna la conexion con el Name Node
*/
func conexionNN(){
        var conn *grpc.ClientConn
        conn, err := grpc.Dial("10.6.40.152:50001", grpc.WithInsecure())
        if err != nil {
                log.Fatalf("did not connect: %s", err)
        }
        defer conn.Close()
        c := pb.NewChatCliDnClient(conn)
        fmt.Println("Conexion realizada correctamente con el Name Node\n")
        return c
}
/*
Funcion: enviarChunks
Parametro:
        - tipoAlgoritmo: string que indica el algoritmo a utilizar (centralizado o distribuido)
        - nombreLibro: el nombre del libro a subir
        - c: conexion gRPC
Descripcion:
	- Recibe la información del libro a subir para mandarselo a un Data Node
Retorno:
	- Nada
*/
func enviarChunks(tipoAlgoritmo string, nombreLibro string, c pb.NewChatCliDnClient){
        fileToBeChunked := nombreLibro + ".pdf"
        file, err := os.Open(fileToBeChunked)
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
        defer file.Close()
        fileInfo, _ := file.Stat()
        var fileSize int64 = fileInfo.Size()
        const fileChunk = 250000
        // calculate total number of parts the file will be chunked into
        totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))
        fmt.Printf("El libro se dividio en %d piezas, subiendo al Data Node. . .\n", totalPartsNum)

        //creamos el stream
        stream, err := c.ChunkaDN(context.Background()) 
        if err != nil {
                log.Fatalf("%v.RecordRoute(_) = _, %v", c, err)
        }
        //creamos los chunks y los mandamos por stream
        for i := uint64(0); i < totalPartsNum; i++ {
                partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
                partBuffer := make([]byte, partSize)
                file.Read(partBuffer)
                
                //enviamos el chunk correspondiente

                chunko := pb.Chunk{
                        NombreLibro: nombreLibro,
                        TotalPartes: strconv.Itoa(totalPartsNum),
                        Parte: strconv.Itoa(i+1),
                        Datos: partBuffer,
                        Algoritmo: tipoAlgoritmo,
                }
                if err := stream.Send(chunko); err != nil { //envio por stream de chunks
                        log.Fatalf("%v.Send(%v) = %v", stream, chunko, err)
                }
        }

        reply, err := stream.CloseAndRecv()
        if err != nil {
                log.Fatalf("%v.CloseAndRecv() tuvo el error %v, quiero %v", stream, err, nil)
        }
        log.Printf("Route summary: %v", reply)

}

/*
Funcion: pedirDirecciones
Parametro:
        - nombreLibro: string con el nombre del libro a descargar
        - c: conexion gRPC
Descripcion:
	- Recibe el nombre del libro y la conexion del name node para buscar las ips en donde estan repartidas las partes del libro
Retorno:
	- Retorna las ips de los data nodes en donde están las partes
*/
func pedirDirecciones(nombreLibro string, c pb.NewChatCliDnClient) []string{
        msj := pb.Message{
		Body: nombreLibro,
        }
        response, err := c.ChunksDirecciones(context.Background(), &msj)
        if err != nil{
                fmt.Println("Error al enviar la orden")
                log.Fatalf("%s", err)
                break
        }
        partesIPS := strings.Split(response.Body, "-")
        partesIPS = partesIPS[:len(partesIPS)-1] //parte 1 en la ip de la posicion 0, parte 2 en la posicion 1, etc
        return partesIPS
}

func main(){
        cDN1 := conexionDN("10.6.40.149") //conexion Data Node 1
        cDN2 := conexionDN("10.6.40.150") //conexion Data Node 2
        cDN3 := conexionDN("10.6.40.151") //conexion Data Node 3
        cNN := conexionNN() //conexion Name Node
        for{    
                var respuesta string
                fmt.Println("\n Quiere subir o descargar un libro? Ingrese la opcion correpsondiente y presione Enter: \n")
                fmt.Println("1 Subir un libro a la red \n")
                fmt.Println("2 Descargar un libro de la red \n")
                fmt.Println("432 para salir")
                _, err := fmt.Scanln(&respuesta)
                if err != nil {
                        fmt.Fprintln(os.Stderr, err)
                        return
                }


                if strings.Compare("1", respuesta) == 0{
                        var nombre string
                        var tipoAlgoritmo string
                        fmt.Println("Ingrese el nombre del libro sin extension y presione Enter")
                        _, err := fmt.Scanln(&nombre)
                        if err != nil {
                                fmt.Fprintln(os.Stderr, err)
                                return
                        }
                        fmt.Println("Ingrese el algoritmo que desea usar: 'centralizado' o 'distribuido' y presione Enter")
                        _, err := fmt.Scanln(&tipoAlgoritmo)
                        if err != nil {
                                fmt.Fprintln(os.Stderr, err)
                                return
                        }
                        //DN elegido aleatoriamente
                        i := rand.Intn(3)
                        if i == 0{
                                enviarChunks(tipoAlgoritmo, nombre, cDN1)
                        }
                        if i == 1{
                                enviarChunks(tipoAlgoritmo, nombre, cDN2)
                        }else{
                                enviarChunks(tipoAlgoritmo, nombre, cDN3)
                        }

                }


                if strings.Compare("2", respuesta) == 0{
                        var nombre string
                        fmt.Println("Ingrese el nombre del libro sin extension y presione Enter")
                        _, err := fmt.Scanln(&nombre)
                        if err != nil {
                                fmt.Fprintln(os.Stderr, err)
                                return
                        }
                        direcciones := pedirDirecciones(nombre, cNN)
                        //recorrer las direcciones ips para ir pidiendo los chunks
                        for _, direccion := range direcciones {
                                fmt.Printf(direccion)
                        }
                        

                }

                if strings.Compare("432", respuesta) == 0{
                        fmt.Println("Saliendo del programa. . . ")
                        break
                }
        }
}