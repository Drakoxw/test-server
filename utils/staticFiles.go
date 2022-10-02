package utils

import (
	"fmt"
	"os"
)

func OpenFile(path string) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Ups!, al parecer el programa no finalizo de forma correcta")
		}
	}()

	// if file, err := os.Open("static/hola.json"); err != nil {
	if file, err := os.Open(path); err != nil {
		panic(err)
		// return "", err
	} else {
		defer func() {
			fmt.Println("Cerrando el archivo!")
			file.Close()
		}()

		contenido := make([]byte, 254)
		long, _ := file.Read(contenido)
		contenidoArchivo := string(contenido[:long])
		return contenidoArchivo, err
	}
}
