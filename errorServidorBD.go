package errores

import "fmt"

// errorServidorBD (http 500) es un error utilizado para registrar errores
// de bases de datos del aplicativo.
type errorServidorBD struct {
	errorEnvoltorio
}

// crearErrorServidorBDf crea un error de tipo 'errorServidorBD'.
func crearErrorServidorBDf(anterior error, formato string, args ...interface{}) *errorServidorBD {
	err := &errorServidorBD{
		errorEnvoltorio{
			origen:   &errorServidorBD{errorEnvoltorio{mensaje: fmt.Sprintf(formato, args...)}},
			anterior: anterior,
			mensaje:  fmt.Sprintf(formato, args...),
		},
	}
	err.establecerUbicacion(2)

	return err
}

// NuevoErrorServidorBD crea un nuevo error de tipo 'errorServidorBD' con formato.
func NuevoErrorServidorBD(err error, formato string, args ...interface{}) error {
	return crearErrorServidorBDf(err, formato, args...)
}

// EnvolverEnErrorServidorBD crea un error de tipo 'errorServidorDB' con
// formato y envuelve el error anterior dentro de este.
// func EnvolverEnErrorServidorBD(err error, formato string, args ...interface{}) error {
// 	if err == nil {
// 		return nil
// 	}
// 	return crearErrorServidorBDf(err, formato, args...)
// }

// EsErrorServidorBD devuelve el resultado de conocer si el error es de tipo
// 'errorServidorBD'.
func EsErrorServidorBD(err error) bool {
	origen := obtenerOrigen(err)
	_, ok := origen.(*errorServidorBD)
	return ok
}
