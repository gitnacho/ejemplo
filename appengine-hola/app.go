// Copyright 2015 Google Inc.
// Liberado bajo la Licencia Apache, versión 2.0 (la "Licencia");
// no puedes utilizar este archivo salvo en conformidad con la Licencia.
// Puedes obtener una copia de la Licencia en:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// A menos que lo requiera la ley aplicable o se acuerde por escrito, el
// software distribuido bajo la licencia se distribuye "TAL CUAL", SIN
// GARANTÍAS NI CONDICIONES DE NINGÚN TIPO, ya sean expresas o implícitas.
// Consulta la licencia para el idioma específico que rige los permisos y
// limitaciones conforme a la licencia.
//
// El paquete hola es una sencilla aplicación de App Engine que responde a
// las peticiones en /hola con un mensaje de bienvenida.
package hola

import (
	"fmt"
	"net/http"
)

// init se ejecuta antes de que la aplicación inicie el servidor.
func init() {
	// Maneja todas las peticiones a la ruta /hola con la función
	// controladorHola.
	http.HandleFunc("/hola", controladorHola)
}

func controladorHola(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola desde la aplicación Go")
}
