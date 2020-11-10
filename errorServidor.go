package errores

import "fmt"

// errorServidor (http 500) es un error utilizado para registrar errores
// del aplicativo.
type errorServidor struct {
	errorEnvoltorio
}

// crearErrorServidorf crea un error de tipo 'errorServidor'.
func crearErrorServidorf(anterior error, formato string, args ...interface{}) *errorServidor {
	err := &errorServidor{
		errorEnvoltorio{
			origen:   &errorServidor{errorEnvoltorio{mensaje: fmt.Sprintf(formato, args...)}},
			anterior: anterior,
			mensaje:  fmt.Sprintf(formato, args...),
		},
	}
	err.establecerUbicacion(2)

	return err
}

// NuevoErrorServidor crea un nuevo error de tipo 'errorServidor' con formato.
func NuevoErrorServidor(formato string, args ...interface{}) error {
	return crearErrorServidorf(nil, formato, args...)
}

// EnvolverEnErrorServidor crea un error 'erroServidor' con formato y
// envuelve el error anterior dentro de este.
// func EnvolverEnErrorServidor(err error, formato string, args ...interface{}) error {
// 	if err == nil {
// 		return nil
// 	}
// 	return crearErrorServidorf(err, formato, args...)
// }

// EsErrorServidor devuelve el resultado de conocer si el error es 'errorServidor'.
func EsErrorServidor(err error) bool {
	origen := obtenerOrigen(err)
	_, ok := origen.(*errorServidor)
	return ok
}
