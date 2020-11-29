package main
import(
        "bufio"
        "io/ioutil"
        //"math/rand"
        "math"
        "strconv"
        "os"
        "strings"
        //"io"
        //"encoding/csv"
        "log"
        "fmt"
        "time"
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
func conexionDN(ip string) pb.ChatCliDnClient{
        var conn *grpc.ClientConn
        puerto := ":50001"
        conn, err := grpc.Dial(ip+puerto, grpc.WithInsecure())
        if err != nil {
                log.Fatalf("did not connect: %s", err)
        }else{
                fmt.Println("Conexion realizada correctamente con el Data Node de IP "+ip+"\n")
        }
        //defer conn.Close()
        c := pb.NewChatCliDnClient(conn)
        
        return c, conn
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
func conexionNN() pb.ChatCliDnClient{
        var conn *grpc.ClientConn
        conn, err := grpc.Dial("10.6.40.152:50001", grpc.WithInsecure())
        if err != nil {
                log.Fatalf("did not connect: %s", err)
        }else{
                fmt.Println("Conexion realizada correctamente con el Name Node\n")
        }
        //defer conn.Close()
        c := pb.NewChatCliDnClient(conn)
        
        return c, conn
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
func enviarChunks(tipoAlgoritmo string, nombreLibro string, c pb.ChatCliDnClient){
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
        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        stream, err := c.ChunkaDN(ctx)
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
                        TotalPartes: strconv.Itoa(int(totalPartsNum)),
                        Parte: strconv.Itoa(int(i+1)),
                        Datos: partBuffer,
                        Algoritmo: tipoAlgoritmo,
                }
                if err := stream.Send(&chunko); err != nil { //envio por stream de chunks
                        log.Fatalf("%v.Send(%v) = %v", stream, chunko, err)
                }
        }

        reply, err := stream.CloseAndRecv()
        if err != nil {
                log.Fatalf("%v.CloseAndRecv() tuvo el error %v, quiero %v", stream, err, nil)
        }
        log.Printf("Route summary: %v", reply)
        cancel()

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
func pedirDirecciones(nombreLibro string, c pb.ChatCliDnClient) []string{
        msj := pb.Message{
		Body: nombreLibro,
        }
        response, err := c.ChunksDirecciones(context.Background(), &msj)
        if err != nil{
                fmt.Println("Error al enviar la orden")
                log.Fatalf("%s", err)
                //break
        }
        partesIPS := strings.Split(response.Body, " # ") //parte 1 en la ip de la posicion 0, parte 2 en la posicion 1, etc
        return partesIPS
}
/*
Funcion: rearmarLibro
Parametro:
        - nombreLibro: string con el nombre del libro a rearmar
        - cantPartes: cantidad de chunks alojados en el directorio
Descripcion:
	- Recibe el nombre del libro y la cantidad de partes del libro para rearmarlo
Retorno:
	- Sin retorno
*/
func rearmarLibro(nombreLibro string, cantPartes int){
        newFileName := nombreLibro+"RECUPERADO"+".pdf"
        _, err := os.Create(newFileName)

        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

        //set the newFileName file to APPEND MODE!!
        // open files r and w

        file, err := os.OpenFile(newFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

        // IMPORTANT! do not defer a file.Close when opening a file for APPEND mode!
        // defer file.Close()

        // just information on which part of the new file we are appending
        var writePosition int64 = 0

        for j := uint64(0); j < uint64(cantPartes); j++ {

                //read a chunk
                currentChunkFileName := nombreLibro + "#" + strconv.FormatUint(j+1, 10)

                newFileChunk, err := os.Open(currentChunkFileName)

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

                //fmt.Println("Appending at position : [", writePosition, "] bytes")
                writePosition = writePosition + chunkSize

                // read into chunkBufferBytes
                reader := bufio.NewReader(newFileChunk)
                _, err = reader.Read(chunkBufferBytes)

                if err != nil {
                        fmt.Println(err)
                        os.Exit(1)
                }

                // DON't USE ioutil.WriteFile -- it will overwrite the previous bytes!
                // write/save buffer to disk
                //ioutil.WriteFile(newFileName, chunkBufferBytes, os.ModeAppend)

                n, err := file.Write(chunkBufferBytes)
                fmt.Println("Escritos ", n, " bytes en el disco")
                if err != nil {
                        fmt.Println(err)
                                        os.Exit(1)
                }

                file.Sync() //flush to disk

                // free up the buffer for next cycle
                // should not be a problem if the chunk size is small, but
                // can be resource hogging if the chunk size is huge.
                // also a good practice to clean up your own plate after eating

                chunkBufferBytes = nil // reset or empty our buffer
        }

        fmt.Println("Libro rearmado correctamente")
        // now, we close the newFileName
        file.Close()
}
func main(){

        cDN1, conn1 := conexionDN("10.6.40.149") //conexion Data Node 1
        cDN2, conn2 := conexionDN("10.6.40.150") //conexion Data Node 2
        cDN3, conn3 := conexionDN("10.6.40.151") //conexion Data Node 3
        cNN, conn4 := conexionNN() //conexion Name Node
        for{    
                var respuesta string
                fmt.Println("\n Quiere subir o descargar un libro? Ingrese la opcion correpsondiente y presione Enter: \n")
                fmt.Println("1 Client Uploader \n")
                fmt.Println("2 Client Downloader \n")
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
                        _, err1 := fmt.Scanln(&tipoAlgoritmo)
                        if err1 != nil {
                                fmt.Fprintln(os.Stderr, err)
                                return
                        }
                        //DN elegido aleatoriamente
                        i := 0  ////// ÑAKSDÑLASKDLASKDLASKDLÑAAAAAAAAAAAAAAAAAAAAAAAAAAKSD ARREGLAR AKI AL FINAL PORFAVOR!!!!!!!!!!!!!!!!!!!!! rand.Intn(3)rand.Intn(3)rand.Intn(3)rand.Intn(3)
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


                        for{
                                var respuesta2 string
                                fmt.Println("1 Mostrar catalogo de libros disponibles para descargar\n")
                                fmt.Println("2 Descargar un libro \n")
                                fmt.Println("432 para salir")
                                _, err := fmt.Scanln(&respuesta2)
                                if err != nil {
                                        fmt.Fprintln(os.Stderr, err)
                                        return
                                }
                                if respuesta2 == "1"{
                                        msj := pb.Message{
                                                Body: "TOY XATO",
                                        }
                                        response, err := cNN.PedirCatalogo(context.Background(), &msj)
                                        if err != nil{
                                                fmt.Println("Error al enviar la solicitud del catalogo")
                                                break
                                        }
                                        log.Printf("%s", response.Body)

                                }
                                if respuesta2 == "2"{
                                        var nombre string
                                        fmt.Println("Ingrese el nombre del libro que desea descargar sin extension y presione Enter")
                                        _, err := fmt.Scanln(&nombre)
                                        if err != nil {
                                                fmt.Fprintln(os.Stderr, err)
                                                return
                                        }
                                        direcciones := pedirDirecciones(nombre, cNN)
                                        //recorrer las direcciones ips para ir pidiendo los chunks
                                        cont := 1
                                        for _, direccion := range direcciones {
                                                if direccion == "10.6.40.149"{
                                                        msj := pb.Message{
                                                                Body: nombre+"#"+strconv.Itoa(cont), //nombreLibro#parte
                                                        }
                                                        response, err := cDN1.PedirChunk(context.Background(), &msj)
                                                        if err != nil{
                                                                fmt.Println("Error al enviar la solicitud del chunk")
                                                                break
                                                        }
                
                                                        //escribir chunk en disco
                                                        fileName := nombre+"#"+strconv.Itoa(cont)
                                                        _, err3 := os.Create(fileName)
                                                        if err3 != nil {
                                                                fmt.Println(err)
                                                                os.Exit(1)
                                                        }
                                                        // write/save buffer to disk
                                                        ioutil.WriteFile(fileName, response.GetDatos(), os.ModeAppend)
                                                        
                                                }
                                                if direccion == "10.6.40.150"{
                                                        msj := pb.Message{
                                                                Body: nombre+"#"+strconv.Itoa(cont), //nombreLibro#parte
                                                        }
                                                        response, err := cDN2.PedirChunk(context.Background(), &msj)
                                                        if err != nil{
                                                                fmt.Println("Error al enviar la solicitud del chunk")
                                                                break
                                                        }
                
                                                        //escribir chunk en disco
                                                        fileName := nombre+"#"+strconv.Itoa(cont)
                                                        _, err4 := os.Create(fileName)
                                                        if err4 != nil {
                                                                fmt.Println(err)
                                                                os.Exit(1)
                                                        }
                                                        // write/save buffer to disk
                                                        ioutil.WriteFile(fileName, response.GetDatos(), os.ModeAppend)
                                                        
                                                }else{
                                                        msj := pb.Message{
                                                                Body: nombre+"#"+strconv.Itoa(cont), //nombreLibro#parte
                                                        }
                                                        response, err := cDN3.PedirChunk(context.Background(), &msj)
                                                        if err != nil{
                                                                fmt.Println("Error al enviar la solicitud del chunk")
                                                                break
                                                        }
                
                                                        //escribir chunk en disco
                                                        fileName := nombre+"#"+strconv.Itoa(cont)
                                                        _, err5 := os.Create(fileName)
                                                        if err5 != nil {
                                                                fmt.Println(err)
                                                                os.Exit(1)
                                                        }
                                                        // write/save buffer to disk
                                                        ioutil.WriteFile(fileName, response.GetDatos(), os.ModeAppend)
                                                        
                                                }
                                                cont += 1
                
                                        }
                
                                        // just for fun, let's recombine back the chunked files in a new file
                                        rearmarLibro(nombre, cont)
                                }
                                if respuesta2 == "432"{
                                        defer conn1.Close()
                                        defer conn2.Close()
                                        defer conn3.Close()
                                        defer conn4.Close()
                                        fmt.Println("Saliendo del Client Downloader. . . ")
                                        break
                                }

                                
                                
                        }

                }

                
                if strings.Compare("432", respuesta) == 0{
                        fmt.Println("Saliendo del programa. . . ")
                        break
                }
        }
}