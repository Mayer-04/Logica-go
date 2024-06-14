package main

import (
	"encoding/json"
	"fmt"
)

//* Las etiquetas en las estructuras (struct), como `json:"name"`
//* Son instrucciones para decirle al programa cómo convertir los datos en formatos como JSON.
//* Son como notas especiales que le dicen al programa cómo comportarse con los datos de la estructura.
//* El valor predeterminado de cada etiqueta en un campo es una cadena en blanco.
// TODO: Podemos usar el paquete reflect para inspeccionar la información de la etiqueta de cada campo.

type Person struct {
	// El campo `Name` se convertirá en la clave "name" cuando se convierta a JSON.
	Name string `json:"name"`

	// La etiqueta "-" indica que este campo debe ignorarse por completo cuando se convierte a JSON.
	Lastname string `json:"-"`

	Age int `json:"age"`

	// La etiqueta "omitempty" indica que el campo se incluirá en el JSON solo si su valor es diferente de cero.
	Active bool `json:"active,omitempty"`

	// La etiqueta "string" convierta campos numéricos o booleanos en cadenas
	Play bool `json:"play,string"`
}

type Pais struct {
	Nombre     string
	Habitantes int
	Capital    string
	Idiomas    []string
}

func main() {

	mayer := Person{
		Name:     "Mayer",
		Lastname: "Chaves",
		Age:      24,
		Active:   false,
	}

	//* json.Marshal: Convertir estructuras de datos Go en una cadena JSON - Serializar (codificación)
	// Devuelve un slice de bytes []byte y un error
	jsonData, err := json.Marshal(mayer)

	if err != nil {
		panic(err)
	}

	str := string(jsonData)

	fmt.Println(str) // Output: {"name":"Mayer","age":24,"play":"false"}

	//* json.Unmarshal: Convertir una cadena JSON a código Go - Decodificación
	//* Importante pasar un "puntero" como segundo argumento

	//TODO: 1. Preparamos la cadena JSON a convertir
	exampleJson := `{"nombre":"Canada","habitantes":37314442,"capital":"Ottawa","idiomas":["Inglés","Frances"]}`
	//TODO: 2. Preparamos la estructura que recibirá el JSON decodificado.
	countrys := Pais{} // var countrys Pais

	if err := json.Unmarshal([]byte(exampleJson), &countrys); err != nil {
		panic(err)
	}

	fmt.Println(countrys) // Output: {Canada 37314442 Ottawa [Inglés Frances]}

}
