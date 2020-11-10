package errores

import "fmt"

// errorNoEncontrado (http 404) es un error utilizado para enviar mensajes
// de error al cliente. Utilizado para indicar que un recurso solicitado es
// inexistente.
type errorNoEncontrado struct {
	errorEnvoltorio
}

// crearErrorNoEncontradof crea un error de tipo 'errorNoEncontrado'.
func crearErrorNoEncontradof(anterior error, formato string, args ...interface{}) *errorNoEncontrado {
	err := &errorNoEncontrado{
		errorEnvoltorio{
			origen:   &errorNoEncontrado{errorEnvoltorio{mensaje: fmt.Sprintf(formato, args...)}},
			anterior: anterior,
			mensaje:  fmt.Sprintf(formato, args...),
		},
	}
	err.establecerUbicacion(2)

	return err
}

// NuevoErrorNoEncontrado crea un nuevo error de tipo 'errorNoEncontrado'
// con formato.
func NuevoErrorNoEncontrado(formato string, args ...interface{}) error {
	return crearErrorNoEncontradof(nil, formato, args...)
}

// EnvolverEnErrorNoEncontrado crea un error de tipo 'errorNoEncontrado'
// con  formato y envuelve el error anterior dentro de este.
// func EnvolverEnErrorNoEncontrado(err error, formato string, args ...interface{}) error {
// 	if err == nil {
// 		return nil
// 	}
// 	return crearErrorNoEncontradof(err, formato, args...)
// }

// EsErrorNoEncontrado devuelve el resultado de conocer si el error
// es del tipo 'errorNoEncontrado'.
func EsErrorNoEncontrado(err error) bool {
	origen := obtenerOrigen(err)
	_, ok := origen.(*errorNoEncontrado)
	return ok
}
