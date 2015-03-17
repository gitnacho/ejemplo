/*
Copyright 2014 Google Inc.
Liberado bajo la Licencia Apache, versión 2.0 (la "Licencia");
no puedes utilizar este archivo salvo en conformidad con la Licencia.
Puedes obtener una copia de la Licencia en:

http://www.apache.org/licenses/LICENSE-2.0

A menos que lo requiera la ley aplicable o se acuerde por escrito, el
software distribuido bajo la licencia se distribuye "TAL CUAL", SIN
GARANTÍAS NI CONDICIONES DE NINGÚN TIPO, ya sean expresas o implícitas.
Consulta la licencia para el idioma específico que rige los permisos y
limitaciones conforme a la licencia.
*/

// yasalio es un servidor web que anuncia cuando una versión particular
// de Go ya ha sido etiquetada o no.
package main

import (
	"expvar"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
	"time"
)

// Opciones de la línea de órdenes.
var (
	dirHttp       = flag.String("http", ":8080", "Dirección de escucha")
	periodoSondeo = flag.Duration("sondeo", 5*time.Second, "Periodo del sondeo")
	versión       = flag.String("versión", "1.4", "Versión de Go")
)

const cambiaUrlBase = "https://code.google.com/p/go/source/detail?r="

func main() {
	flag.Parse()
	cambiaURL := fmt.Sprintf("%sgo%s", cambiaUrlBase, *versión)
	http.Handle("/", NuevoServidor(*versión, cambiaURL, *periodoSondeo))
	log.Fatal(http.ListenAndServe(*dirHttp, nil))
}

// Variables exportadas para monitorear el servidor.
// Estas se exportan vía HTTP como un objeto JSON en /debug/vars.
var (
	contadorVisitas       = expvar.NewInt("contadorVisitas")
	contadorSondeos       = expvar.NewInt("contadorSondeos")
	errorSondeo           = expvar.NewString("errorSondeo")
	contadorErroresSondeo = expvar.NewInt("contadorErroresSondeo")
)

// Servidor implementa el servidor yasalio.
// Este sirve la interfaz de usuario (es un http.Handler)
// y sondea los cambios en el repositorio remoto.
type Servidor struct {
	versión  string
	url      string
	periodo  time.Duration
	mu  sync.RWMutex // protege la variable sí
	sí  bool
}

// NuevoServidor regresa un servidor yasalio iniciado.
func NuevoServidor(versión, url string, periodo time.Duration) *Servidor {
	s := &Servidor{versión: versión, url: url, periodo: periodo}
	go s.sondea()
	return s
}

// sondeo sondea el cambio de URL para el periodo especificado hasta
// que la etiqueta exista.
// Luego ajusta si el campo 'sí' del Servidor es cierto y sale.
func (s *Servidor) sondea() {
	for !estáEtiquetado(s.url) {
		sondeaPausa(s.periodo)
	}
	s.mu.Lock()
	s.sí = true
	s.mu.Unlock()
	sondeaHecho()
}

// Ganchos que puedes redefinir para pruebas de integración.
var (
	sondeaPausa = time.Sleep
	sondeaHecho = func() {}
)

// estáEtiquetado hace una petición HEAD al HTTP a la URL dada e informa
// si este regresa una respuesta 200 OK.
func estáEtiquetado(url string) bool {
	contadorSondeos.Add(1)
	r, err := http.Head(url)
	if err != nil {
		log.Print(err)
		errorSondeo.Set(err.Error())
		contadorErroresSondeo.Add(1)
		return false
	}
	return r.StatusCode == http.StatusOK
}

// ServeHTTP implementa la interfaz de usuario HTTP.
func (s *Servidor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	contadorVisitas.Add(1)
	s.mu.RLock()
	datos := struct {
		URL     string
		Versión string
		Sí      bool
	}{
		s.url,
		s.versión,
		s.sí,
	}
	s.mu.RUnlock()
	err := plant.Execute(w, datos)
	if err != nil {
		log.Print(err)
	}
}

// plant es la plantilla HTML que controla la interfaz de usuario.
var plant = template.Must(template.New("plant").Parse(`<!DOCTYPE html>
<html lang="es">
  <body>
    <center>
	<h2>¿Ya salió la versión {{.Versión}} de Go?</h2>
	<h1>
	{{if .Sí}}
		<a href="{{.URL}}">¡Sí! :)</a>
	{{else}}
		¡No! :-(
	{{end}}
	</h1>
    </center>
  </body>
</html>
`))
