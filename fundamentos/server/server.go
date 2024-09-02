package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

/*
* Servidor HTTP en Go
- El paquete 'net/http' permite crear y gestionar servidores HTTP de manera nativa en Go.
- El paquete 'encoding/json' facilita el trabajo con datos en formato JSON,
proporcionando funciones para su codificación y decodificación.

* Multiplexor (Mux):
- Se utiliza para organizar y gestionar cómo se manejan diferentes rutas en un servidor HTTP.
- Permite definir rutas específicas que corresponden a diferentes manejadores (handlers).
- Cuando llega una solicitud HTTP, el multiplexor verifica la ruta solicitada y la asigna al manejador
correspondiente para procesar la petición y responder al cliente.

* Algunos conceptos a tener en cuenta:
- Escribir datos: Consiste en enviar o transferir información desde un programa (como un servidor)
a un cliente HTTP, una conexión de red, un archivo o un buffer.
- fmt.Fprintf(): Permite escribir datos formateados a un destino, que puede ser un archivo, una conexión de red o un buffer.
El primer argumento es un 'io.Writer', que determina el destino de la escritura.
- http.Error(): Envía un mensaje de error al cliente como parte de la respuesta HTTP,
junto con el código de estado correspondiente.
- w.Write(): Escribe datos en una respuesta HTTP que se enviará al cliente. El argumento que toma es un []byte,
que generalmente representa datos binarios o texto. Estos datos se enviarán en el cuerpo de la respuesta HTTP.
*/

// Definimos una estructura 'User' para representar un usuario.
// Las etiquetas `json:""` especifican cómo los campos serán representados en JSON.
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// * Simulación de una base de datos en memoria.
// La variable 'users' almacena los usuarios en un mapa.
// La clave es un entero que representa el ID del usuario, y el valor es un struct `User`.
var users = make(map[int]User)

func main() {

	// Crea un servidor HTTP.
	// Nuevo multiplexor de solicitudes HTTP.
	mux := http.NewServeMux()

	// Define la ruta y el manejo de la petición.
	mux.HandleFunc("/", handleRoot)

	// La solicitud debe ser de tipo 'POST' y la ruta debe ser '/users'.
	// Esta sintaxis es solo válida a partir de la versión 1.22 de Go.
	mux.HandleFunc("POST /users", createUser)

	// La solicitud debe ser de tipo 'GET' y la ruta debe ser '/users/{id}'.
	// Recupera el parámetro '{id}' de la solicitud.
	mux.HandleFunc("GET /users/{id}", getUsers)

	// La solicitud debe ser de tipo 'DELETE' y la ruta debe ser '/users/{id}'.
	// Agregamos una exprsión regular que consiste en un número entre 0 y 9.
	mux.HandleFunc("DELETE /users/{id:[0-9]+}", deleteUser)

	// Imprimimos un mensaje cuando el servidor se inicia.
	fmt.Println("Servidor iniciado en el puerto :8080")

	//* Iniciamos el servidor.
	// El primer argumento es la dirección o puerto donde escuchará el servidor.
	// El segundo argumento es el manejador de solicitudes (en este caso, `mux`).
	http.ListenAndServe(":8080", mux)
}

// Controlador para manejar la ruta raíz.
// Escribe un mensaje de "Hola Mundo" como respuesta.
func handleRoot(w http.ResponseWriter, r *http.Request) {

	// Nos permite escribir cualquier contenido en la respuesta.
	fmt.Fprintf(w, "Hola Mundo!")
	// w.Write([]byte("Hello, World!"))
}

func createUser(w http.ResponseWriter, r *http.Request) {

	// Creamos una instancia vacía de `User` para almacenar los datos del cuerpo de la solicitud.
	var user User

	// Crea un decodificador basado en el cuerpo (body) de la petición.
	// Decodifica los datos de la petición y los almacena en la variable 'user'.
	// La variable 'user' contendra todos los datos de que se especifique en el body de la petición.
	// Evidentemente el cuerpo de la solicitud debe coincidir con nuestra estructura 'User'.
	if error := json.NewDecoder(r.Body).Decode(&user); error != nil {
		// Le pasamos la respuesta, el error y el estado de la petición en este caso 400.
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}

	//* Validaciones básicas.
	// Verificamos que los campos requeridos no estén vacíos.
	if user.Name == "" {
		http.Error(w, "el nombre es requerido", http.StatusBadRequest)
		return
	}

	if user.Email == "" {
		http.Error(w, "el email es requerido", http.StatusBadRequest)
		return
	}

	// Guardamos el usuario en la base de datos.
	users[len(users)+1] = user

	// Indicamos que la solicitud fue exitosa con un estado 201.
	w.WriteHeader(http.StatusCreated)

	// Respondemos con un mensaje de confirmación.
	fmt.Fprintf(w, "el usuario %s fue creado exitosamente", user.Name)
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	// Recuperamos el valor del parámetro '{id}' de la ruta de la solicitud.
	// Este método `PathValue` está disponible desde la versión 1.22 de Go.
	userId := r.PathValue("id")

	// Convertimos el ID de usuario de string a entero.
	id, error := strconv.Atoi(userId)

	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}

	// Buscamos el usuario en el mapa `users`.
	user, ok := users[id]

	if !ok {
		// Si el usuario no existe, devolvemos un error 404 (Not Found).
		http.Error(w, "el usuario no fue encontrado", http.StatusNotFound)
		return
	}

	// Configuramos la cabecera de la respuesta para indicar que es de tipo JSON.
	w.Header().Set("Content-Type", "application/json")

	// Transformamos el usuario a JSON para mandarlo como respuesta al cliente.
	data, error := json.Marshal(user)

	if error != nil {
		// Si ocurre un error al codificar el JSON, devolvemos un error 500 (Internal Server Error).
		// Es un error interno del servidor porque hasta este punto el usuario ingreso las datos correctamente.
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	// Devuelve la respuesta al cliente con el estado 200.
	w.WriteHeader(http.StatusOK)
	// Devuelve los datos del usuario en formato JSON.
	w.Write(data)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {

	// Recuperamos el valor del parámetro '{id}' de la ruta de la solicitud.
	userId := r.PathValue("id")

	// Convertimos el ID de usuario de string a entero.
	id, error := strconv.Atoi(userId)

	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}

	// Verificamos si el usuario existe en el mapa.
	if _, ok := users[id]; !ok {
		http.Error(w, "el usuario no fue encontrado", http.StatusNotFound)
		return
	}

	// Eliminamos el par clave y valor del mapa por su id.
	delete(users, id)

	w.WriteHeader(http.StatusNoContent)
	fmt.Fprintf(w, "el usuario %s fue eliminado exitosamente", users[id].Name)
}
