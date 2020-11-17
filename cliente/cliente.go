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
        "github.com/432i/T1SisDistribuidos/logistica/chat"
)

func main(){
        var conn *grpc.ClientConn
        conn, err := grpc.Dial("10.6.40.149:9000", grpc.WithInsecure())
        if err != nil {
                log.Fatalf("did not connect: %s", err)
        }
        defer conn.Close()

        c := chat.NewChatServiceClient(conn)

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
                        fmt.Println("xd")

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