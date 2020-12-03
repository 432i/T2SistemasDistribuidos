
# Tarea 2 del ramo Sistema Distribuidos
## https://github.com/432i/T2SistemasDistribuidos
## Comenzando 🚀

## 1. Entrar a la máquina correspondiente:
- Máquina 1 (dist09: ip 10.6.40.149): Data Node 1
- Máquina 2 (dist10: ip 10.6.40.150): Data Node 2
- Máquina 3 (dist11: ip 10.6.40.151): Data Node 3
- Máquina 4 (dist12: ip 10.6.40.152): Name Node

## 2. Entrar a la carpeta T2SistemasDistribuidos y entrar a la carpeta correspondiente según la entidad de la máquina (el cliente puede estar alojado en cualquiera de estas máquinas)

- Para cliente: carpeta cliente
- Para data node: carpeta dataNode y luego DataNode1, DataNode2 o DataNode3
- Para name node: carpeta nameNode

## 3. Escribir make y presionar enter en la consola para ejecutar el código
## Consideraciones (leer antes):
La maquina que es seleccionada para recibir el stream de chunks no puede ser apagada en ningun momento ni antes de que se envien los chunks.
Asumimos que no existiran dos libros que se llamen igual y que los libros no pueden tener caracteres especiales (nombre_libro por ejemplo). Se acepta nombreLibro.
Se deben ejecutar primero los data nodes, luego el name node y al final el cliente para un correcto funcionamiento

Por otra parte, y dado lo mencionado anteriormente, a continuacion se presenta como nosotros interpretamos la tarea:
En primer lugar, la propuesta realizada, ya sea para Distribuido o Centralizado, la realizamos por cada chunk (esto lo pueden notar en las funciones generarPropuesta y generarPropuestaCentralizado, donde hay un for que itera en las partes de un stream), ya que pensamos que en el proceso de envío de un libro, puede ocurrir que algún data node se caiga entre medio del envío de los chunks mismos, lo que suponemos que es algo que puede pasar en la realidad, servidores tipo data nodes distribuidos no necesariamente en el mismo lugar fisico y que por causas externas se apaguen repentinamente. Ante esto, puede suceder que algunos chunks se hayan almacenado en data nodes que en algún momento estuvieron encendidos, y que por lo tanto, para poder recuperar el libro, se deban restablecer (es decir, que vuelvan a funcionar, como un proceso de recuperacion real (?)

Por esta misma razón, no pudimos implementar directamente la medición de tiempo solo entre la escritura del log (ya que como nuestro proceso se hace por chunk, golang no permite hacer sumas de tiempo entre variables de tipo time.Time -la funcion encargada de sumar recibe un tipo time.Duration y no se puede convertir de time.Time a time.Duration para lograrlo-), así que, a modo de compensación, optamos por medir el tiempo que se demora cada opción de reparticion (distribuida vs centralizada) desde que se recibía el stream de un libro hasta que este se lograba almacenar en los data nodes.

A modo de aclaración, nos percatamos de que nuestra interpretación difería con la de otros compañeros demasiado tarde como para poder cambiar toda la logica del código


Antes de hacerlo se deben exportar variables, para esto escribir los siguientes comandos en consola:
- export GOROOT=/usr/local/go
- export GOPATH=$HOME/go
- export GOBIN=$GOPATH/bin
- export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN

Presionar enter y ejecutar los programas haciendo make

Asegurarse de que el firewall está desactivado o los métodos gRPC no funcionarán:

- service firewalld stop

## Autores ✒️

* **Ignacio Aedo, rol 201773556-2** - *Desarrollo* - [432i](https://github.com/432i)
* **Ethiel Carmona, rol 201773533-3** - *Desarrollo* - [ethielc](https://github.com/ethielc)

## Construido con 🛠️
* Lenguaje Go
* gRPC
* Protocol Buffers
