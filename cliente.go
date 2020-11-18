package main

import (
	"fmt"
	"time"
	"net"
	"encoding/gob"
)

var parar int64
var proceso Proceso

type Proceso struct {
	Id int64
	I int64
}

func (p *Proceso) HacerProceso() {
	for {
		p.I = p.I + 1
		time.Sleep(time.Millisecond * 500)
	}
}

func (p *Proceso) MostrarProceso() {
	fmt.Println("id ", p.Id,": " ,p.I)
	p.I = p.I + 1
}

func cliente()  {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := "."
	c.Write([]byte(msg))
	c.Close()
}

func clienteEscuchar()  {
	
	s, err := net.Listen("tcp", ":9998")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		handleServidor(c)
		s.Close()
		c.Close()
		return
	}
}

func handleServidor(c net.Conn)  {
	
	err := gob.NewDecoder(c).Decode(&proceso)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func mostrarProceso(proceso1 *Proceso) {
	for {

		time.Sleep(time.Millisecond * 500)
		proceso1.MostrarProceso()
		if parar != -1 {
			return
		}
	}
}

func clienteMandarProceso(){
	c, err := net.Dial("tcp", ":9997")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = gob.NewEncoder(c).Encode(proceso)
	if err != nil {
		fmt.Println(err)
	}
	
	c.Close()
}

func main()  {
	parar = -1
	go cliente()
	go clienteEscuchar()
	go mostrarProceso(&proceso)
	fmt.Scanln(&parar)
	parar = 1

	if parar != -1{
		clienteMandarProceso()
	}
} 