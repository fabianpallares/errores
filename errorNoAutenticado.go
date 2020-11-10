package errores

import "fmt"

// errorNoAutenticado (http 401) es un error utilizado para enviar mensajes
// de error al cliente. Es utilizado en servicios RESTful para indicarle al
// cliente que no puede consumir el recurso por no estar debidamente autenticado.
type errorNoAutenticado struct {
	errorEnvoltorio
}

// crearErrorValidacionf crea un error de tipo 'errorNoAutenticado'.
func crearErrorNoAutenticadof(anterior error, formato string, args ...interface{}) *errorNoAutenticado {
	err := &errorNoAutenticado{
		errorEnvoltorio{
			origen:   &errorNoAutenticado{errorEnvoltorio{mensaje: fmt.Sprintf(formato, args...)}},
			anterior: anterior,
			mensaje:  fmt.Sprintf(formato, args...),
		},
	}
	err.establecerUbicacion(2)

	return err
}

// NuevoErrorNoAutenticado crea un nuevo error de tipo 'errorNoAutenticado'
// con formato.
func NuevoErrorNoAutenticado(formato string, args ...interface{}) error {
	return crearErrorNoAutenticadof(nil, formato, args...)
}

// EnvolverEnErrorNoAutenticado crea un error de tipo 'errorNoAutenticado'
// con  formato y envuelve el error anterior dentro de este.
// func EnvolverEnErrorNoAutenticado(err error, formato string, args ...interface{}) error {
// 	if err == nil {
// 		return nil
// 	}
// 	return crearErrorNoAutenticadof(err, formato, args...)
// }

// EsErrorNoAutenticado devuelve el resultado de conocer si el error
// es del tipo 'errorNoAutenticado'.
func EsErrorNoAutenticado(err error) bool {
	origen := obtenerOrigen(err)
	_, ok := origen.(*errorNoAutenticado)
	return ok
}
