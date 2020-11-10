package errores

import "fmt"

// errorRequerimientoMuyGrande (http 413) es un error utilizado para enviar
// mensajes de error al cliente. Utilizado en servicios RESTful para indicar
// que los valores enviados por el cliente son muy grandes (exceden el límite
// permitido).
type errorRequerimientoMuyGrande struct {
	errorEnvoltorio
}

// crearErrorValidacionf crea un error de tipo 'errorRequerimientoMuyGrande'.
func crearErrorRequerimientoMuyGrandef(anterior error, formato string, args ...interface{}) *errorRequerimientoMuyGrande {
	err := &errorRequerimientoMuyGrande{
		errorEnvoltorio{
			origen:   &errorRequerimientoMuyGrande{errorEnvoltorio{mensaje: fmt.Sprintf(formato, args...)}},
			anterior: anterior,
			mensaje:  fmt.Sprintf(formato, args...),
		},
	}
	err.establecerUbicacion(2)

	return err
}

// NuevoErrorRequerimientoMuyGrande crea un nuevo error de tipo 'errorRequerimientoMuyGrande'
// con formato.
func NuevoErrorRequerimientoMuyGrande(formato string, args ...interface{}) error {
	return crearErrorRequerimientoMuyGrandef(nil, formato, args...)
}

// EnvolverEnErrorRequerimientoMuyGrande crea un error de tipo 'errorRequerimientoMuyGrande'
// con  formato y envuelve el error anterior dentro de este.
// func EnvolverEnErrorRequerimientoMuyGrande(err error, formato string, args ...interface{}) error {
// 	if err == nil {
// 		return nil
// 	}
// 	return crearErrorRequerimientoMuyGrandef(err, formato, args...)
// }

// EsErrorRequerimientoMuyGrande devuelve el resultado de conocer si el error
// es del tipo 'errorRequerimientoMuyGrande'.
func EsErrorRequerimientoMuyGrande(err error) bool {
	origen := obtenerOrigen(err)
	_, ok := origen.(*errorRequerimientoMuyGrande)
	return ok
}
