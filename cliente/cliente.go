package main
import(
        "bufio"
        "io/ioutil"
        "math"
        "strconv"
        "os"
        "strings"
        "io"
        "encoding/csv"
        "log"
        "fmt"
        "time"
        "golang.org/x/net/context"
        "google.golang.org/grpc"
        "github.com/432i/T2SistemasDistribuidos/dependencias/serverclidn"
)


func main(){
        var conn *grpc.ClientConn
        conn, err := grpc.Dial("10.6.40.149:9000", grpc.WithInsecure())
        if err != nil {
                log.Fatalf("did not connect: %s", err)
        }
        defer conn.Close()

        c := serverclidn.NewChatServiceClient(conn)

        //response, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello From Client!"})
        //if err != nil {
        //        log.Fatalf("Error when calling SayHello: %s", err)
        //}
        //log.Printf("Response from server: %s", response.Body)

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
                        fmt.Println("Ingrese el nombre del archivo que desea subir en formato 'archivo.pdf'")
                        _, err := fmt.Scanln(&nombre)
                        if err != nil {
                                fmt.Fprintln(os.Stderr, err)
                                return
                        }
                        fileToBeChunked := nombre // change here!

                        file, err := os.Open(fileToBeChunked)

                        if err != nil {
                                fmt.Println(err)
                                os.Exit(1)
                        }

                        defer file.Close()

                        fileInfo, _ := file.Stat()

                        var fileSize int64 = fileInfo.Size()

                        const fileChunk = 250000*8 // 1 MB, change this to your requirement

                        // calculate total number of parts the file will be chunked into

                        totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

                        fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

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

                                fmt.Println("Split to : ", fileName)

                                //enviamos el chunk correspondiente

                                message := serverclidn.Chunk{
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


                if strings.Compare("2", respuesta) == 0{
                        fmt.Println("xd")
                        

                }

                if strings.Compare("432", respuesta) == 0{
                        fmt.Println("Saliendo del programa. . . ")
                        break
                }
        }
}