package errores

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// envoltura es la interface que deberán implementar los errores.
type envoltura interface {
	// obtenerOrigen devuelve el origen del error. Este será el Error original,
	// o un tipo forzado por la acción de invocar una función Envolver<enTipo>,
	// ya que al envolver <enTipo> reasigna el tipo del error.
	obtenerOrigen() error

	// obtenerAnterior devuelve el error anterior o nil si no hay ninguno.
	obtenerAnterior() error

	// obtenerMensaje devuelve el mensaje de error.
	obtenerMensaje() string

	// obtenerUbicacion devuelve el nombre del paquete, el nombre del archivo y
	// el número de línea donde se registró el error.
	obtenerUbicacion() (string, string, int)

	// establecerUbicacion establece la ubicación donde se ha generado el error.
	// Se guardan los valores del nombre de paquete, el nombre del archivo y el
	// número de línea donde se genera la llamada a esta función.
	// Se debe especificar la 'profundidad' de la llamada.
	establecerUbicacion(int)
}

// errorEnvoltorio contiene una descripción de un error junto con información
// sobre donde se creó el error.
// Puede estar incrustado en tipos de envoltorio personalizados para agregar
// información adicional que este paquete de errores puede comprender.
type errorEnvoltorio struct {
	origen   error
	anterior error
	mensaje  string
	paquete  string
	archivo  string
	linea    int
}

// Error implementa la interface error.
func (e *errorEnvoltorio) Error() string {
	if e.mensaje != "" {
		return e.mensaje
	}
	return e.origen.Error()
}

// obtenerOrigen implementa la interface envoltura.
func (e *errorEnvoltorio) obtenerOrigen() error {
	return e.origen
}

// obtenerAnterior implementa la interface envoltura.
func (e *errorEnvoltorio) obtenerAnterior() error {
	return e.anterior
}

// obtenerMensaje implementa la interface envoltura.
func (e *errorEnvoltorio) obtenerMensaje() string {
	return e.mensaje
}

// obtenerUbicacion implementa la interface envoltura.
func (e *errorEnvoltorio) obtenerUbicacion() (string, string, int) {
	return e.paquete, e.archivo, e.linea
}

// establecerUbicacion implementa la interface envoltura.
func (e *errorEnvoltorio) establecerUbicacion(profundidad int) {
	_, paquete, _, ok := runtime.Caller(0)
	if ok {
		e.paquete = paquete
	}

	_, archivo, linea, ok := runtime.Caller(profundidad + 1)
	if ok {
		pos := strings.LastIndex(archivo, "/")
		e.archivo = archivo[pos+1 : len(archivo)-3]
		e.linea = linea
	}
}

// EsEnvoltura devuelve verdadero si el error recibido es de tipo
// 'envoltura' (significa que cumple con la interface 'envoltura').
func EsEnvoltura(err error) bool {
	_, ok := err.(envoltura)
	return ok
}

// Nuevo crea un error de tipo envoltura, sólo con el mensaje formateado
// y la ubicación.
//
// Por ejemplo:
// 	return errores.Nuevo("Un error")
//
func Nuevo(formato string, args ...interface{}) error {
	nuevoEnvoltorio := &errorEnvoltorio{mensaje: fmt.Sprintf(formato, args...)}
	nuevoEnvoltorio.establecerUbicacion(1)

	return nuevoEnvoltorio
}

// Envolver se utiliza para agregar contexto (mensaje con formato) adicional a
// un error existente. La ubicación se guarda.
//
// Por ejemplo:
// 	if err := AlgunaFunc(); err != nil {
// 		return errores.Envolver(err, "El error es: %v", "Fallo adicional")
// 	}
//
func Envolver(err error, formato string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	nuevoEnvoltorio := &errorEnvoltorio{
		origen:   obtenerOrigen(err),
		anterior: err,
		mensaje:  fmt.Sprintf(formato, args...),
	}
	nuevoEnvoltorio.establecerUbicacion(1)

	return nuevoEnvoltorio
}

// ObtenerOrigen devuelve el error de origen.
func ObtenerOrigen(err error) error {
	return obtenerOrigen(err)
}

// ObtenerAnterior devuelve el error anterior al origen.
func ObtenerAnterior(err error) error {
	errEnvoltorio, ok := err.(*errorEnvoltorio)
	if !ok {
		return err
	}

	return errEnvoltorio.anterior
}

// obtenerOrigen devuelve el origen del error.
// Si el origen del error no es una envoltura, devuelve el mismo error recibido.
// Si el error es una envoltura, pueden pasar dos cosas:
// 	Si el origen del error (err.origen) es nulo, devuelve el mismo error recibido.
// 	Si el origen del error (err.origen) no es nulo, devuelve ese error.
func obtenerOrigen(err error) error {
	if errEnvoltura, ok := err.(envoltura); ok {
		origen := errEnvoltura.obtenerOrigen()
		if origen != nil {
			return origen
		}
	}

	return err
}

// ObtenerPrimerError devuelve el primer error.
func ObtenerPrimerError(err error) error {
	return obtenerPrimerError(err)
}

// ObtenerErrores devuelve un string con la pila de errores concatenados y
// listo para ser mostrados con el siguiente formato:
// 	[tipo de error (nombre del programa: número de línea)] mensaje del error
//
func ObtenerErrores(err error) [][]string {
	return obtenerErrores(err)
}

// ObtenerErroresApilados devuelve un slice con la pila de errores
// del error recibido. Los valores son los siguientes:
// 	Tipo de error
// 	Mensaje
// 	Nombre del paquete
// 	Nombre del programa
// 	Número de línea
func ObtenerErroresApilados(err error) string {
	var errs string
	for _, v := range obtenerErrores(err) {
		if v[3] != "" {
			errs += fmt.Sprintf("[%v (%v: %v)] %v\n", v[0], v[3], v[4], v[1])
		} else {
			errs += fmt.Sprintf("[%v] %v\n", v[0], v[1])
		}
	}

	return errs
}

// obtenerErrores devuelve un slice con la pila de errores
// del error recibido. Los valores son los siguientes:
// 	Tipo de error
// 	Mensaje
// 	Nombre del paquete
// 	Nombre del programa
// 	Número de línea
func obtenerErrores(err error) [][]string {
	crearElemento := func(err error) []string {
		var elemento []string
		env, ok := err.(envoltura)
		if !ok {
			if err != nil {
				elemento = append(elemento, fmt.Sprintf("%T", err), err.Error(), "", "", "")
			}
		} else {
			paq, arch, lin := env.obtenerUbicacion()
			elemento = append(elemento, fmt.Sprintf("%T", env), env.obtenerMensaje(), paq, arch, strconv.Itoa(lin))
		}

		return elemento
	}

	if err == nil {
		return nil
	}

	var retorno [][]string
	for {
		elemento := crearElemento(err)
		retorno = append(retorno, elemento)

		env, ok := err.(envoltura)
		if !ok {
			break
		}
		err = env.obtenerAnterior()
		if err == nil {
			break
		}
	}

	return retorno
}

func obtenerPrimerError(err error) error {
	if err == nil {
		return nil
	}

	for {
		env, ok := err.(envoltura)
		if !ok {
			break
		}
		err = env.obtenerAnterior()
		if err == nil {
			break
		}
	}

	return err
}
