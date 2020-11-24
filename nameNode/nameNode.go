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
        for _, s := range nombresLibros { 
                escribirLog(s, "3", "1", "1.66.6", 1)
                for i := 0; i < 3; i++ {
                        escribirLog(s, "3", strconv.Itoa(i+1), "1.66.6", 0)

	        }
        numLibros += 1
        }
}