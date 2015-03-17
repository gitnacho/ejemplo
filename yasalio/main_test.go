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

package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// StatusHandler es un http.Handler que escribe una respuesta vacía
// usándose a sí mismo como el código de estado de respuesta.
type controladorEstado int

func (h *controladorEstado) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(int(*h))
}

func TestEstáEtiquetado(t *testing.T) {
	// Configura un falso servidor web "Google Code" informando 404 No
	// encontrado.
	estado := controladorEstado(http.StatusNotFound)
	s := httptest.NuevoServidor(&estado)
	defer s.Close()
	if estáEtiquetado(s.URL) {
		t.Fatal("estáEtiquetado == true, esperaba false")
	}
	// Cambia el estado del falso servidor a 200 OK y vuelve a intentarlo.
		estado = http.StatusOK
	if !estáEtiquetado(s.URL) {
		t.Fatal("estáEtiquetado == false, esperaba true")
	}
}

func TestIntegración(t *testing.T) {
	estado := controladorEstado(http.StatusNotFound)
	te := httptest.NuevoServidor(&estado)
	defer te.Close()
	// Reemplaza el sondeaPausa con un cierre que se puede bloquear
	// y desbloquear.
	pausa := make(chan bool)
	sondeaPausa = func(time.Duration) {
		pausa <- true
		pausa <- true
	}

	// Reemplaza sondeaHecho con un cierre que nos dirá cuando
	// está saliendo del sondeo.
	hecho := make(chan bool)
	sondeaHecho = func() { hecho <- true }
	// Pone las cosas como estaban cuando la prueba finalice.
	defer func() {
		sondeaPausa = time.Sleep
		sondeaHecho = func() {}
	}()
	s := NuevoServidor("1.x", te.URL, 1*time.Millisecond)
	<-pausa // Espera al bucle sondeo para empezar la pausa.
	// Hace la primer petición al servidor.
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	s.ServeHTTP(w, r)
	if b := w.Body.String(); !strings.Contains(b, "No.") {
		t.Fatalf("body = %s, esperaba no", b)
	}
	estado = http.StatusOK
	<-pausa // Permite dejar la pausa.
	<-hecho // Espera al sondeo para ver el estado "OK" y salir.
	// Hace la segunda petición al servidor.
	w = httptest.NewRecorder()
	s.ServeHTTP(w, r)
	if b := w.Body.String(); !strings.Contains(b, "Sí!") {
		t.Fatalf("body = %q, esperaba sí", b)
	}
}
