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

// El paquete utilcadenas contiene útiles funciones para trabajar con
// cadenas de caracteres.
package utilcadenas

// Reverso invierte su argumento s dejándolo legible de izquierda
// a derecha.
func Reverso(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
