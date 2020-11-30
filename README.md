
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
Asumimos que no existiran dos libros que se llamen igual y que los libros no pueden tener caracteres especiales (nombre_libro por ejemplo). Se acepta nombreLibro.
Se deben ejecutar primero los data nodes, luego el name node y al final el cliente para un correcto funcionamiento


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
