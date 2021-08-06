package main

import (
	"fmt"
)

/*definir la longitud del mapa puede liberar espacio de la memoria*/
func main() {
	nombre := make(map[string]string)
	fmt.Println(nombre)

	nombre["last_name"] = "moya"
	nombre["last_name"] = "ome"
	fmt.Println(nombre)
	/*Esta es otra forma de agregar mapas, semejante a un diccionario en python*/
	paises := map[string]int{
		"Colombia":  20,
		"Argentina": 30,
		"Venezuela": 40,
		"Ecuador":   50}

	/*Para agregar un nuevo elemento */
	paises["Bracil"] = 34
	/*para poder eliminar de un mapa*/
	// delete(paises, "Argentina")
	/*Para poder iterar en un mapa*/
	for pais, numero := range paises {
		fmt.Printf("El pais %s tiene el numero %d \n", pais, numero)
	}
	fmt.Println(paises)
	/*Para buscar un valor en el mapa */
	pais, puntaje := paises["Colombia"]
	fmt.Printf("el numero es %d, y el pais existe %t \n", pais, puntaje)
}
