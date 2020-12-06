package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"strconv"
)

var smaterias = make([]Materia, 0)
var salumnos = make([]Alumno, 0)

type Materia struct {
	nombre string
	m      map[string]string
}
type Alumno struct {
	nombre string
	a      map[string]string
}

type Mensaje struct {
	alumno       string
	calificacion string
	materia      string
}
type Server struct{}

func (this *Server) Registrar(slice []string, reply *string) error {
	var msg Mensaje
	msg.alumno = slice[0]
	msg.materia = slice[1]
	msg.calificacion = slice[2]

	if existeMateria(msg.materia) {

		for _, i := range smaterias {

			if i.nombre == msg.materia {
				_, prs := i.m[msg.alumno]

				if prs {
					return errors.New("Ya existe calificacion para ese alumno en la materia")
				}

			}

		}
	}

	existe := existeMateria(msg.materia)
	if existe {
		for _, i := range smaterias {
			if i.nombre == msg.materia {
				i.m[msg.alumno] = msg.calificacion

			}

		}

	} else {
		var aux Materia
		n := make(map[string]string)
		aux.nombre = msg.materia
		aux.m = n
		aux.m[msg.alumno] = msg.calificacion
		smaterias = append(smaterias, aux)

	}

	existe = existeAlumno(msg.alumno)
	if existe {
		for _, i := range salumnos {
			if i.nombre == msg.alumno {
				i.a[msg.materia] = msg.calificacion

			}

		}
	} else {
		var aux Alumno
		n := make(map[string]string)
		aux.nombre = msg.alumno
		aux.a = n
		aux.a[msg.materia] = msg.calificacion
		salumnos = append(salumnos, aux)

	}
	fmt.Println(smaterias)
	fmt.Println(salumnos)
	*reply = "Calificacion registrada"
	return nil
}

func (this *Server) PromedioAlumno(name string, reply *float64) error {
	var total float64
	existe := false
	for _, i := range salumnos {
		if i.nombre == name {
			existe = true
			for _, mapa := range i.a {
				aux, _ := strconv.ParseFloat(mapa, 64)
				total += aux
			}
			c := float64(len(i.a))
			*reply = total / c
		}

	}
	if existe {
		return nil
	} else {

		return errors.New("no existe ese alumno")
	}

}
func (this *Server) PromedioGeneral(mensaje string, reply *float64) error {
	var promedio float64
	var promedioG float64 = 0
	fmt.Println(mensaje)
	for _, i := range salumnos {
		var total float64
		for _, mapa := range i.a {
			aux, _ := strconv.ParseFloat(mapa, 64)
			total += aux
		}
		c := float64(len(i.a))
		promedio = total / c
		promedioG += promedio
	}
	c := float64(len(salumnos))
	*reply = promedioG / c

	if len(salumnos) == 0 {
		return errors.New("no existen alumnos")
		return nil
	} else {

		return nil
	}
}

func (this *Server) PromedioMateria(name string, reply *float64) error {
	var total float64
	existe := false
	for _, i := range smaterias {
		if i.nombre == name {
			existe = true
			for _, mapa := range i.m {
				aux, _ := strconv.ParseFloat(mapa, 64)
				total += aux
			}
			c := float64(len(i.m))
			*reply = total / c
		}

	}
	if existe {
		return nil
	} else {

		return errors.New("no existe esa Materia")
	}

}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}

func existeMateria(materia string) bool {
	existe := false
	for _, i := range smaterias {
		if i.nombre == materia {
			existe = true

		}

	}

	return existe
}

func existeAlumno(alumno string) bool {
	existe := false
	for _, i := range salumnos {
		if i.nombre == alumno {
			existe = true

		}

	}

	return existe
}
