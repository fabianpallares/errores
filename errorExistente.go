package errores

import "fmt"

// errorExistente (http 400) es un error utilizado para enviar mensajes de
// error al cliente.
// Principalmente pensado para informar que ya existe un elemento en el repositorio
// de datos, con el mismo valor clave que se está recibiendo.
type errorExistente struct {
	errorEnvoltorio
}

// crearErrorExistentef crea un error de tipo 'errorExistente'.
func crearErrorExistentef(anterior error, formato string, args ...interface{}) *errorExistente {
	err := &errorExistente{
		errorEnvoltorio{
			origen:   &errorExistente{errorEnvoltorio{mensaje: fmt.Sprintf(formato, args...)}},
			anterior: anterior,
			mensaje:  fmt.Sprintf(formato, args...),
		},
	}
	err.establecerUbicacion(2)

	return err
}

// NuevoErrorExistente crea un nuevo error de tipo 'errorExistente'
// con formato.
func NuevoErrorExistente(formato string, args ...interface{}) error {
	return crearErrorExistentef(nil, formato, args...)
}

// EnvolverEnErrorExistente crea un error de tipo 'errorExistente'
// con  formato y envuelve el error anterior dentro de este.
// func EnvolverEnErrorExistente(err error, formato string, args ...interface{}) error {
// 	if err == nil {
// 		return nil
// 	}
// 	return crearErrorExistentef(err, formato, args...)
// }

// EsErrorExistente devuelve el resultado de conocer si el error
// es del tipo 'errorExistente'.
func EsErrorExistente(err error) bool {
	origen := obtenerOrigen(err)
	_, ok := origen.(*errorExistente)
	return ok
}
