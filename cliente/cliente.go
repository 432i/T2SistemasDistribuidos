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

func enviarChunks(nombreLibro string, c pb.NewChatCliDnClient){
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
        for i := uint64(0); i < totalPartsNum; i++ {
                partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
                partBuffer := make([]byte, partSize)
                file.Read(partBuffer)
                // write to disk
                fileName := "bigfile_" + strconv.FormatUint(i, 10)
                _, err := os.Create(fileName)

                if err != nil {
                        fmt.Println(err)
                        os.Exit(1)
                }
                // write/save buffer to disk
                ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)
                //fmt.Println("Split to : ", fileName)

                //enviamos el chunk correspondiente

                message := pb.Chunk{
                        NombreLibro: nombreLibro,
                        TotalPartes: strconv.Itoa(totalPartsNum),
                        Parte: strconv.Itoa(i+1),
                        Datos: partBuffer,
                }
                response, err := c.ChunkaDN(context.Background(), &message)
                if err != nil{
                        fmt.Println("Error al enviar el chunk")
                        log.Fatalf("%s", err)
                        break
                }
                log.Printf("%s", response.Body)
        }
}
func pedirDirecciones(nombre, c pb.NewChatCliDnClient){
        msj := pb.Message{
		Body: nombre,
        }
        response, err := c.ChunksDirecciones(context.Background(), &ms)
        if err != nil{
                fmt.Println("Error al enviar la orden")
                log.Fatalf("%s", err)
                break
        }
        direcciones := response.Body
        //
        //
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
                        fmt.Println("Ingrese el nombre del archivo sin extension y presione Enter")
                        _, err := fmt.Scanln(&nombre)
                        if err != nil {
                                fmt.Fprintln(os.Stderr, err)
                                return
                        }
                        //DN elegido aleatoriamente
                        i := rand.Intn(3)
                        if i == 0{
                                enviarChunks(nombre, cDN1)
                        }
                        if i == 1{
                                enviarChunks(nombre, cDN2)
                        }else{
                                enviarChunks(nombre, cDN3)
                        }

                }


                if strings.Compare("2", respuesta) == 0{
                        var nombre string
                        fmt.Println("Ingrese el nombre del archivo sin extension y presione Enter")
                        _, err := fmt.Scanln(&nombre)
                        if err != nil {
                                fmt.Fprintln(os.Stderr, err)
                                return
                        }
                        direcciones := pedirDirecciones(nombre, cNN)
                        
                        

                }

                if strings.Compare("432", respuesta) == 0{
                        fmt.Println("Saliendo del programa. . . ")
                        break
                }
        }
}