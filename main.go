package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)


func main() {
	if len(os.Args) != 2{
		fmt.Println("\nlisten.exe 8080\n")
		fmt.Println("listen.exe 8080-9000\n")
		fmt.Println("listen.exe 8080,9000,9090\n")
		fmt.Println("listen.exe 8080-9000,7000,7001\n")
		fmt.Println("listen.exe 8080-9000,7000-7070\n")
		os.Exit(1)
	}

	var ports []string

	if strings.Contains(os.Args[1],","){
		tmp_port0 := strings.Split(os.Args[1],",")
		for _,v := range tmp_port0 {
			if strings.Contains(v, "-") {
				v0, _ := strconv.Atoi(strings.Split(v, "-")[0])
				v1, _ := strconv.Atoi(strings.Split(v, "-")[1])
				for i := v0; i <= v1; i++ {
					ports = append(ports, ":"+strconv.Itoa(i))
				}
			} else {
				ports = append(ports, ":"+v)
			}
		}
	}else if strings.Contains(os.Args[1],"-"){
		v0,_ := strconv.Atoi(strings.Split(os.Args[1],"-")[0])
		v1,_ := strconv.Atoi(strings.Split(os.Args[1],"-")[1])
		for i := v0 ;i <=v1;i++ {
			ports = append(ports,":"+strconv.Itoa(i))
		}
	}else{
		ports = append(ports,":"+os.Args[1])
	}

	fmt.Println(ports)

	for _,v := range ports{
		go func(port string){
			//建立socket，监听端口
			netListen, err := net.Listen("tcp", "localhost"+port)
			if err != nil{
				CheckError(err)
				return
			}

			defer netListen.Close()

			Log("Waiting for clients "+port)
			for {
				conn, err := netListen.Accept()
				if err != nil {
					continue
				}

				Log(conn.RemoteAddr().String(), " tcp connect success"+port)
				handleConnection(conn)
			}
		}(v)
	}
	select{}
}
//处理连接
func handleConnection(conn net.Conn) {

	buffer := make([]byte, 2048)

	for {

		_, err := conn.Read(buffer)

		if err != nil {
			//Log(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		//Log(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))
	}

}
func Log(v ...interface{}) {
	log.Println(v...)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		//os.Exit(1)
	}
}