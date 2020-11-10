package errores

import "fmt"

// errorMalFormato (http 415) es un error utilizado para enviar mensajes de
// al cliente, para avisarle a este que la informaci información recibida
// contiene errores de formateo. En el caso de servicios RESTful, es utilizado
// al verificar formatos de datos JSON recibidos en el cuerpo del mensaje.
type errorMalFormato struct {
	errorEnvoltorio
}

// crearErrorValidacionf crea un error de tipo 'errorMalFormato'.
func crearErrorMalFormatof(anterior error, formato string, args ...interface{}) *errorMalFormato {
	err := &errorMalFormato{
		errorEnvoltorio{
			origen:   &errorMalFormato{errorEnvoltorio{mensaje: fmt.Sprintf(formato, args...)}},
			anterior: anterior,
			mensaje:  fmt.Sprintf(formato, args...),
		},
	}
	err.establecerUbicacion(2)

	return err
}

// NuevoErrorMalFormato crea un nuevo error de tipo 'errorMalFormato'
// con formato.
func NuevoErrorMalFormato(formato string, args ...interface{}) error {
	return crearErrorMalFormatof(nil, formato, args...)
}

// EnvolverEnErrorMalFormato crea un error de tipo 'errorMalFormato'
// con  formato y envuelve el error anterior dentro de este.
// func EnvolverEnErrorMalFormato(err error, formato string, args ...interface{}) error {
// 	if err == nil {
// 		return nil
// 	}
// 	return crearErrorMalFormatof(err, formato, args...)
// }

// EsErrorMalFormato devuelve el resultado de conocer si el error
// es del tipo 'errorMalFormato'.
func EsErrorMalFormato(err error) bool {
	origen := obtenerOrigen(err)
	_, ok := origen.(*errorMalFormato)
	return ok
}
