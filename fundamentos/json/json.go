package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

/*
* JSON y Etiquetas en Go
- Las etiquetas en las estructuras (struct), como `json:"name"`,
son instrucciones para decirle al programa cómo convertir los datos en formatos como JSON.

* Las etiquetas son cadenas que siguen el formato `json:"clave"`:
- `"clave"`: El nombre de la clave en el JSON.
- `"-"`: Indica que el campo debe ser ignorado en la codificación y decodificación JSON.
- `"omitempty"`: El campo solo se incluye si su valor no es el valor cero del tipo.
- `"string"`: Convierte campos numéricos o booleanos a cadenas en JSON.
- Por defecto, las etiquetas están vacías, y se utilizan para personalizar la forma en que los datos se representan en JSON.

IMPORTANTE: Puedes usar el paquete `reflect` para inspeccionar las etiquetas de los campos.
*/

type Person struct {
	// El campo `Name` se convertirá en la clave "name" en JSON.
	Name string `json:"name"`

	// La etiqueta "-" indica que este campo se omitirá en JSON.
	Lastname string `json:"-"`

	// El campo `Age` se convertirá en la clave "age" en JSON.
	Age int `json:"age"`

	// La etiqueta "omitempty" indica que el campo solo se incluirá si su valor no es el valor cero.
	Active bool `json:"active,omitempty"`

	// La etiqueta "string" convierte campos numéricos o booleanos en cadenas en JSON.
	Play bool `json:"play,string"`
}

type Country struct {
	Nombre     string   `json:"nombre"`
	Habitantes int      `json:"habitantes"`
	Capital    string   `json:"capital"`
	Idiomas    []string `json:"idiomas"`
}

func main() {
	// Crear una instancia de Person con algunos campos.
	mayer := Person{
		Name:     "Mayer",
		Lastname: "Chaves",
		Age:      24,
		Active:   false,
	}

	// Convertir la estructura Person a JSON (serialización).
	jsonData, err := json.Marshal(mayer)
	if err != nil {
		panic(err)
	}

	// Convertir el slice de bytes a una cadena para imprimirlo.
	str := string(jsonData)
	fmt.Println("JSON de Person:", str) // Output: {"name":"Mayer","age":24,"play":"false"}

	// Convertir una cadena JSON a una estructura Go (deserialización).
	exampleJson := `{"nombre":"Canada","habitantes":37314442,"capital":"Ottawa","idiomas":["Inglés","Francés"]}`
	countrys := Country{}

	if err := json.Unmarshal([]byte(exampleJson), &countrys); err != nil {
		panic(err)
	}

	fmt.Println("Estructura Country:", countrys) // Output: {Canada 37314442 Ottawa [Inglés Francés]}

	// Usar reflect para inspeccionar etiquetas de los campos.
	t := reflect.TypeOf(Country{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		typeName := field.Type.Name()
		fmt.Printf("field %s (%s), tag: %q\n", field.Name, typeName, tag)
	}
}
