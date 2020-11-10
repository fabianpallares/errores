package errores

import "fmt"

// ErrorDeValidacion es una estructura que almacena errores que se han
// producido al validar los datos recibidos del cliente.
type ErrorDeValidacion struct {
	CampoEstructura string // Nombre del campo en la estructura del objeto
	CampoJSON       string // Nombre del campo en el objeto json
	CampoNombre     string // Nombre del campo que se quiere mostrar
	Errores         []string
}

// errorDeValidacion (http 400) es un error utilizado para enviar mensajes
// de error al cliente. Puede indicar que los datos recibidos contienen errores
// de validación.
type errorValidacion struct {
	errorEnvoltorio
	errores []ErrorDeValidacion
}

// crearErrorValidacionf crea un error de tipo 'errorValidacion'.
func crearErrorValidacionf(anterior error, errores []ErrorDeValidacion, formato string, args ...interface{}) *errorValidacion {
	err := &errorValidacion{
		errorEnvoltorio{
			origen:   &errorValidacion{errorEnvoltorio{mensaje: fmt.Sprintf(formato, args...)}, errores},
			anterior: anterior,
			mensaje:  fmt.Sprintf(formato, args...),
		},
		errores,
	}
	err.establecerUbicacion(2)

	return err
}

// NuevoErrorValidacion crea un nuevo error de tipo 'errorValidacion'
// con formato.
func NuevoErrorValidacion(errores []ErrorDeValidacion, formato string, args ...interface{}) error {
	return crearErrorValidacionf(nil, errores, formato, args...)
}

// EnvolverEnErrorValidacion crea un error de tipo 'errorValidacion' con
// formato y envuelve el error anterior dentro de este.
// func EnvolverEnErrorValidacion(err error, errores []ErrorDeValidacion, formato string, args ...interface{}) error {
// 	if err == nil {
// 		return nil
// 	}
// 	return crearErrorValidacionf(err, errores, formato, args...)
// }

// EsErrorValidacion devuelve el resultado de conocer si el error
// es del tipo 'errorValidacion'.
func EsErrorValidacion(err error) bool {
	origen := obtenerOrigen(err)
	_, ok := origen.(*errorValidacion)
	return ok
}

// ObtenerErroresDeValidacion devuelve una lista de objetos 'ErrorDeValidacion'.
func ObtenerErroresDeValidacion(err error) []ErrorDeValidacion {
	origen := obtenerOrigen(err)
	errVal, ok := origen.(*errorValidacion)
	if ok {
		return errVal.errores
	}

	return []ErrorDeValidacion{}
}
