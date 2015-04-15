# Proyectos Go de ejemplo

Este repositorio contiene una colección de programas y bibliotecas que
demuestran el lenguaje Go, bibliotecas y herramientas estándar.

## Los ejemplos

### [Hola](hola/) [![GoDoc](https://godoc.org/github.com/gitnacho/ejemplo/hola?status.svg)](https://godoc.org/github.com/gitnacho/ejemplo/hola) y [utilcadenas](utilcadenas/) [![GoDoc](https://godoc.org/github.com/gitnacho/ejemplo/utilcadenas?status.svg)](https://godoc.org/github.com/gitnacho/ejemplo/utilcadenas)

	go get github.com/gitnacho/ejemplo/hola

Un programa "Hola, mundo" trivial que utiliza el paquete utilcadenas.

La orden [hola](hola/) cubre:

* La forma básica de una orden ejecutable
* Importación de paquetes (de la biblioteca estándar y el repositorio
  local)
* Impresión de cadenas ([fmt](//golang.org/pkg/fmt/))

La biblioteca [utilcadenas](utilcadenas/) cubre:

* La forma básica de una biblioteca
* La conversión entre cadena y []rune
* Pruebas unitarias basadas en tablas ([testing](//golang.org/pkg/testing/))

### [yasalio](yasalio/)([godoc](//godoc.org/github.com/gitnacho/ejemplo/yasalio))

	go get github.com/gitnacho/ejemplo/yasalio

Un servidor web que responde a la pregunta: "¿ya salió Go 1.x?"

Temas tratados:

* Opciones de la línea de órdenes ([flag](//golang.org/pkg/flag/))
* Servidores web ([net/http](//golang.org/pkg/net/http/))
* Plantillas HTML ([html/template](//golang.org/pkg/html/template/))
* Registro cronológico de eventos ([log](//golang.org/pkg/log/))
* Procesos de larga duración en segundo plano
* Sincronización de acceso a datos entre rutinasgo ([sync](//golang.org/pkg/sync/))
* Exportación de estado del servidor para monitoreo ([expvar](//golang.org/pkg/expvar/))
* Pruebas unitarias y de integración ([testing](//golang.org/pkg/testing/))
* Inyección de dependencias
* Tiempo ([time](//golang.org/pkg/time/))

### [appengine-hola](appengine-hola/)([godoc](//godoc.org/github.com/gitnacho/ejemplo/appengine-hola))

	goapp get github.com/gitnacho/ejemplo/appengine-hola

Una aplicación App Engine "Hola, mundo" trivial destinada a utilizarse
como punto de partida para tu propio código.

_Nota_: La herramienta `goapp` es parte del
[Google App Engine SDK para Go](https://cloud.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go).
