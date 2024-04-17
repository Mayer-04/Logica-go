package main

//* Las etiquetas en las estructuras (struct), como `json:"name"`
//* Son como instrucciones para decirle al programa cómo convertir los datos en formatos como JSON.
//* Son como notas especiales que le dicen al programa cómo comportarse con los datos de la estructura.

type Person struct {
	// El campo `Name` se convertirá en la clave "name" cuando se convierta a JSON.
	Name string `json:"name"`

	// La etiqueta "-" indica que este campo debe ignorarse por completo cuando se convierte a JSON.
	Lastname string `json:"-"`

	Age int `json:"age"`

	// La etiqueta ",omitempty" indica que el campo se incluirá en el JSON solo si su valor es diferente de cero.
	Active bool `json:"active,omitempty"`
}

func main() {

}
