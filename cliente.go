package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func limpiarPantalla(){
	fmt.Print("\033[H\033[2J")
}

func client() {
	c, err := rpc.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var op int64
	for {
		fmt.Println("1.- Registrar Calificacion")
		fmt.Println("2.- Mostrar Promedio De Un Alumno")
		fmt.Println("3.- Mostrar Promedio General")
		fmt.Println("4.- Mostrar Promedio De Una Materia")
		fmt.Println("5.- Salir")
		fmt.Scanln(&op)

		switch op {
		case 1:
			limpiarPantalla()
			var name string
			var calificacion string
			var materia string
			var result string
			fmt.Print("Alumno: ")
			scanner.Scan()
			name = scanner.Text()
			fmt.Print("Calificacion: ")
			scanner.Scan()
			calificacion = scanner.Text()
			fmt.Print("Materia: ")
			scanner.Scan()
			materia = scanner.Text()
			slice := []string{name, materia, calificacion}
			err = c.Call("Server.Registrar", slice, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
			limpiarPantalla()
		case 2:
			limpiarPantalla()
			var name string
			fmt.Print("Alumno: ")
			scanner.Scan()
			name = scanner.Text()

			var result float64
			err = c.Call("Server.PromedioAlumno", name, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("El promedio de ", name, " es: ", result)
			}
			scanner.Scan()
			limpiarPantalla()
		case 3:
			limpiarPantalla()
			var result float64
			err = c.Call("Server.PromedioGeneral", "Promedio General", &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("El promedio general es: ", result)
			}
			scanner.Scan()
			limpiarPantalla()
		case 4:
			limpiarPantalla()
			var name string
			fmt.Print("Nombre de la materia: ")
			scanner.Scan()
			name = scanner.Text()

			var result float64
			err = c.Call("Server.PromedioMateria", name, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("El promedio de la materia ", name, "es: ", result)
			}
			scanner.Scan()
			limpiarPantalla()
		case 5:
			return
		}
	}
}

func main() {
	client()
}
