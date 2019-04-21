package main

import "fmt"

const helloWorld string = "Hola %s %s, bienvenido"
const testConst = "Test"

func main(){
	//name := getName()
	lastname := getLastname()
	fmt.Printf(helloWorld, getName(), lastname)
}

func getName() string {
	var name string = "Nombre prueba"
	fmt.Println("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)
	return name
}

func getLastname() string {
	lastname := "Modificar mi apellido"
	return lastname
}
