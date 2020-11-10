# errores: manejador de errores

[![Go Report Card](https://goreportcard.com/badge/github.com/fabianpallares/errores)](https://goreportcard.com/report/github.com/fabianpallares/errores) [![GoDoc](https://godoc.org/github.com/fabianpallares/errores?status.svg)](https://godoc.org/github.com/fabianpallares/errores)

El manejo de errores en Go es muy sencillo y a la vez muy poderoso. Basta con
crear estructuras que implementen la interface 'error' para empezar a utilizarlas.

Los problemas comienzan cuando se necesita lanzar y enviar hacia arriba los errores generados por el aplicativo y encadenar los diversos errores generados por funciones lejanas.

De aquí surjen varias preguntas:<br/>
- ¿ Cómo se puede hacer para conocer la pila de errores en un momento determinado ?
- ¿ Cómo saber qué tipo de error se ha producido ?
- ¿ En qué paquete... en qué programa y en qué número de línea ?

Todas estas cuestiones, son las que se intentan solucionar con la utilización de
este paquete de errores.

Para lograr esto, me basé en varios artículos; especialmente en aquellos escritos
por Dave Cheney.

En Go, existe un provervio que dice asi:<br/>
_"No se limite a comprobar los errores, manéjelos con gracia."_

A continuación presento los enlaces de los artículos mencionados:<br/>
- https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
- https://www.youtube.com/watch?v=lsBF58Q-DnY

Otros paquetes interesantes y que utilizan la misma filosofía, son lo siguientes:<br/>
- https://github.com/pkg/errors
- https://github.com/juju/errors

## Instalación:
Para instalar el paquete utilice la siguiente sentencia:
```
go get -u github.com/fabianpallares/errores
```

## El concepto:
El manejo tradicional de errores en Go para verificar errores, es el siguiente:
```Go
err := funcion(<parámetros>)
if err != nil {
	// Tratar el error...
}

o...

if err := funcion(<parámetros>); err != nil {
	// Tratar el error...
}
```

Esta manera de manejar los errores no es muy conveniente. La forma clásica no permite persistir los errores, encadenarlos y envolverlos para poder tomar una decisión al final del camino.

## La solución:
Con este paquete de errores, es posible crear nuevos tipos, envolverlos, encadenarlos, conocer en que programa y en qué línea sucedió el error.

### Creación de errores:
En vez de crear errores al vuelo con:
```Go
return errors.New("Este es un error")

o...

return fmt.Errorf("Este es un error")
```

Podremos utilizar la creación de errores tipificados:
```Go
return errores.Nuevo("Este es un error")

o...

return errores.NuevoErrorServidor("Error en el servidor: %v", <variable>)

o...

err := errores.NuevoErrorMalFormato("Error de formato en los datos recibidos")
```

Los tipos de errores de este paquete, tienen una concordancia con las respuestas
que suelen brindarse en la utilización de servicios RESTful a traves de http(s).
Cada uno de ellos está relacionado con los códigos de estado de http:
```Go
http 400:
	ErrorExistente
	ErrorValidacion
	ErrorMalRequerimiento

http 401:
	ErrorNoAutenticado

http 403:
	ErrorNoAutorizado

http 404:
	ErrorNoEncontrado

http 413:
	ErrorRequerimientoMuyGrande

http 415:
	ErrorMalFormato

http 423:
	ErrorRecursoBloqueado

http 500:
	ErrorServidor
	ErrorServidorBD // Error en el servidor producido por el motor de la base de datos.
```

### Envolver errores:
La magia de encadenar errores y poder ser enviados hacia las capas superiores,
se logra a traves de la envoltura:

```Go
err := guardarPersona(<parámetros>)
return errores.EnvolverEnErrorServidor(err, "Error al guardar los datos de la persona")

o...

return errores.EnvolverEnErrorServidor(guardarPersona(<parámetros>), "Error al guardar los datos de la persona")
```

Cada uno de los tipos de error mencionados anteriormente, pueden ser utilizados
para crear nuevos errores o para ser envueltos dentro de ellos.

### Obtención de errores:
El paquete dispone de varios mecanismos para obtener los errores producidos en el aplicativo:

```Go
// Devuelve [][]strings.
listaDeErrores := errores.ObtenerErrores(err)

// Devuelve string.
erroresStr := errores.ObtenerErroresApilados(err)
	Ejemplo:
	[*errores.errorServidor (envoltura_test: 14)] No es posible guardar los datos
	[*errores.errorEnvoltorio (envoltura_test: 11)] Un mensaje de error producido
	por la BD.

// Devuelve string del mensaje del error de origen.
errores.ObtenerMensajeOrigen(err)
```

Además es posible preguntar si la respuesta es de un error determinado:
```Go
if errores.EsErrorNoEncontrado(err) {
	// Tratar el error...
}
```

#### Documentación:
[Documentación en godoc](https://godoc.org/github.com/fabianpallares/errores)

