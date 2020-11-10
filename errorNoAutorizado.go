package errores

import "fmt"

// errorNoAutorizado (http 403) es un error utilizado para enviar mensajes
// de error al cliente. Es utilizado en servicios RESTful para indicarle al
// cliente que no puede consumir el recurso solicitado por no cumplir con
// la autorización adecuada.
type errorNoAutorizado struct {
	errorEnvoltorio
}

// crearErrorValidacionf crea un error de tipo 'errorNoAutorizado'.
func crearErrorNoAutorizadof(anterior error, formato string, args ...interface{}) *errorNoAutorizado {
	err := &errorNoAutorizado{
		errorEnvoltorio{
			origen:   &errorNoAutorizado{errorEnvoltorio{mensaje: fmt.Sprintf(formato, args...)}},
			anterior: anterior,
			mensaje:  fmt.Sprintf(formato, args...),
		},
	}
	err.establecerUbicacion(2)

	return err
}

// NuevoErrorNoAutorizado crea un nuevo error de tipo 'errorNoAutorizado'
// con formato.
func NuevoErrorNoAutorizado(formato string, args ...interface{}) error {
	return crearErrorNoAutorizadof(nil, formato, args...)
}

// EnvolverEnErrorNoAutorizado crea un error de tipo 'errorNoAutorizado'
// con  formato y envuelve el error anterior dentro de este.
// func EnvolverEnErrorNoAutorizado(err error, formato string, args ...interface{}) error {
// 	if err == nil {
// 		return nil
// 	}
// 	return crearErrorNoAutorizadof(err, formato, args...)
// }

// EsErrorNoAutorizado devuelve el resultado de conocer si el error
// es del tipo 'errorNoAutorizado'.
func EsErrorNoAutorizado(err error) bool {
	origen := obtenerOrigen(err)
	_, ok := origen.(*errorNoAutorizado)
	return ok
}
