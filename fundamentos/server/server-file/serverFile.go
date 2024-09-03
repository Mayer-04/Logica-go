package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
)

/*
* Servidor web estático en Go
Este programa crea un servidor web que sirve archivos estáticos desde una carpeta específica (public).

* Paquetes utilizados en el código:
- El paquete 'path/filepath' proporciona funciones para trabajar con rutas de archivos y directorios.
- El paquete 'os' proporciona funciones para interactuar con el sistema operativo,
como obtener el directorio de trabajo actual.
- El paquete 'slog' se utiliza para registrar eventos y mensajes informativos en el sistema de registro.
*/

func main() {

	// Obtiene la ruta completa a la carpeta 'public', donde se encuentran los archivos estáticos.
	publicPath := getPublicPath()
	// Crea un servidor de archivos estáticos que servirá los archivos desde la carpeta 'public'.
	fs := http.FileServer(http.Dir(publicPath))
	// Definimos el manejador de solicitudes HTTP.
	http.Handle("/", fs)
	// Definimos el puerto en el que el servidor web escuchará las solicitudes.
	addr := ":5000"
	slog.Info("Servidor iniciado en el puerto " + addr)
	// Iniciamos el servidor web en el puerto especificado.
	http.ListenAndServe(addr, nil)
}

// getCurrentDirectory obtiene el directorio de trabajo actual.
func getCurrentDirectory() (string, error) {
	// Obtener el directorio de trabajo actual.
	dir, err := os.Getwd()
	if err != nil {
		// Devuelve un error envuelto con un mensaje explicativo si no se puede obtener el directorio.
		return "", fmt.Errorf("error obteniendo el directorio: %w", err)
	}
	// Devuelve la ruta del directorio actual si no hubo errores.
	return dir, nil
}

// getPublicPath construye la ruta completa a la carpeta 'public' donde se almacenan los archivos estáticos.
func getPublicPath() string {
	// Obtiene el directorio de trabajo actual.
	dir, err := getCurrentDirectory()
	if err != nil {
		// Registra el error si no se puede obtener el directorio de trabajo actual.
		slog.Any("getCurrentDirectory():", err)

		// Termina el programa con un estado de error genérico.
		// El código 1 indica una falla por una razón que no está claramente definida.
		os.Exit(1)
	}

	// Define una lista de directorios que componen la ruta relativa a la carpeta 'public'.
	directorios := []string{"fundamentos", "server", "server-file", "public"}
	for _, directorio := range directorios {
		// Acumula la ruta completa añadiendo cada directorio de la lista.
		dir = filepath.Join(dir, directorio)
	}

	// Devuelve la ruta completa a la carpeta 'public'.
	return dir
}
