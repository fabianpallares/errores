package errores

import "fmt"

// errorMalRequerimiento (http 400) es un error utilizado para enviar mensajes
// de error al cliente. Puede ser utilizado para indicar errores de reglas
// de negocio del aplicativo.
type errorMalRequerimiento struct {
	errorEnvoltorio
}

// crearErrorValidacionf crea un error de tipo 'errorMalRequerimiento'.
func crearErrorMalRequerimientof(anterior error, formato string, args ...interface{}) *errorMalRequerimiento {
	err := &errorMalRequerimiento{
		errorEnvoltorio{
			origen:   &errorMalRequerimiento{errorEnvoltorio{mensaje: fmt.Sprintf(formato, args...)}},
			anterior: anterior,
			mensaje:  fmt.Sprintf(formato, args...),
		},
	}
	err.establecerUbicacion(2)

	return err
}

// NuevoErrorMalRequerimiento crea un nuevo error de tipo 'errorMalRequerimiento'
// con formato.
func NuevoErrorMalRequerimiento(formato string, args ...interface{}) error {
	return crearErrorMalRequerimientof(nil, formato, args...)
}

// EnvolverEnErrorMalRequerimiento crea un error de tipo 'errorMalRequerimiento'
// con  formato y envuelve el error anterior dentro de este.
// func EnvolverEnErrorMalRequerimiento(err error, formato string, args ...interface{}) error {
// 	if err == nil {
// 		return nil
// 	}
// 	return crearErrorMalRequerimientof(err, formato, args...)
// }

// EsErrorMalRequerimiento devuelve el resultado de conocer si el error
// es del tipo 'errorMalRequerimiento'.
func EsErrorMalRequerimiento(err error) bool {
	origen := obtenerOrigen(err)
	_, ok := origen.(*errorMalRequerimiento)
	return ok
}
