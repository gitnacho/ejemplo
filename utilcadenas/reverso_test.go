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

package utilcadenas

import "testing"
func TestReverso(t *testing.T) {
	for _, c := range []struct {
		ingresado, esperaba string
	}{
		{"Hola, mundo", "odnum ,aloH"},
		{"Hola, 世界", "界世 ,aloH"},
		{"", ""},
	} {
		obtuve := Reverso(c.ingresado)
		if obtuve != c.esperaba {
			t.Errorf("Reverso(%q) == %q, esperaba %q", c.ingresado, obtuve, c.esperaba)
		}
	}
}
