/*
Copyright 2015 Google Inc.
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

"use strict";

function recuperaMensaje() {
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.open("GET", "/hola", false);
    xmlHttp.send(null);
    document.getElementById("mensaje").innerHTML = xmlHttp.responseText;
}
