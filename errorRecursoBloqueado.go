package errores

import "fmt"

// errorRecursoBloqueado (http 423) es un error utilizado para enviar mensajes
// de error al cliente. Utilizado en servicios RESTful para indicar que un
// recurso se encuentra bloqueado, ya sea por mantenimiento o por desuso.
type errorRecursoBloqueado struct {
	errorEnvoltorio
}

// crearErrorRecursoBloqueadof crea un error de tipo 'errorRecursoBloqueado'.
func crearErrorRecursoBloqueadof(anterior error, formato string, args ...interface{}) *errorRecursoBloqueado {
	err := &errorRecursoBloqueado{
		errorEnvoltorio{
			origen:   &errorRecursoBloqueado{errorEnvoltorio{mensaje: fmt.Sprintf(formato, args...)}},
			anterior: anterior,
			mensaje:  fmt.Sprintf(formato, args...),
		},
	}
	err.establecerUbicacion(2)

	return err
}

// NuevoErrorRecursoBloqueado crea un nuevo error de tipo 'errorRecursoBloqueado'
// con formato.
func NuevoErrorRecursoBloqueado(formato string, args ...interface{}) error {
	return crearErrorNoAutenticadof(nil, formato, args...)
}

// EnvolverEnErrorRecursoBloqueado crea un error de tipo 'errorRecursoBloqueado'
// con  formato y envuelve el error anterior dentro de este.
// func EnvolverEnErrorRecursoBloqueado(err error, formato string, args ...interface{}) error {
// 	if err == nil {
// 		return nil
// 	}
// 	return crearErrorRecursoBloqueadof(err, formato, args...)
// }

// EsErrorRecursoBloqueado devuelve el resultado de conocer si el error
// es del tipo 'errorRecursoBloqueado'.
func EsErrorRecursoBloqueado(err error) bool {
	origen := obtenerOrigen(err)
	_, ok := origen.(*errorRecursoBloqueado)
	return ok
}
