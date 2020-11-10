package errores

import (
	"errors"
	"fmt"
	"testing"
)

func TestNuevo(t *testing.T) {
	// t.Skip()
	err := Nuevo("Un mensaje de error producido por la BD")
	fmt.Printf("%#v\n", err)
	fmt.Println()
	err = EnvolverEnErrorServidor(err, "No es posible guardar los datos")
	fmt.Println(ObtenerErroresApilados(err))
	fmt.Println()
	fmt.Println(ObtenerErrores(err))
}

func TestEnvolver(t *testing.T) {
	t.Skip()
	err := errors.New("un error")
	fmt.Println("es envoltura", EsTipoEnvoltura(err))

	err = Envolver(err, "adentro hay un error común")
	fmt.Println("es envoltura", EsTipoEnvoltura(err))
	fmt.Println()
	fmt.Println(ObtenerErrores(err))
}

// func ExampleNuevo() {
// 	err := errores.Nuevo("Este es un mensaje de error")
// 	fmt.Println(err)
// }
