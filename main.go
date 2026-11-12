package main // Identifica el paquete principal de SIGEP.

import (
	"fmt"              // Permite mostrar el menu y leer opciones.
	"sigep/clientes"   // Importa la gestion de clientes.
	"sigep/entregas"   // Importa la gestion de entregas.
	"sigep/pagos"      // Importa la gestion de pagos.
	"sigep/pedidos"    // Importa la gestion de pedidos.
	"sigep/plantas"    // Importa la gestion de plantas.
	"sigep/produccion" // Importa la gestion de produccion.
)

// main es la funcion donde comienza el programa.
func main() {

	// for mantiene el menu activo hasta seleccionar salir.
	for {

		// Muestra el menu principal del sistema.
		fmt.Println("\n===== SISTEMA SIGEP =====")
		fmt.Println("1. Gestion de clientes")
		fmt.Println("2. Gestion de plantas")
		fmt.Println("3. Gestion de pedidos")
		fmt.Println("4. Gestion de produccion")
		fmt.Println("5. Gestion de pagos")
		fmt.Println("6. Gestion de entregas")
		fmt.Println("0. Salir")

		// Declara la variable para guardar la opcion.
		var opcion int

		// Solicita una opcion al usuario.
		fmt.Print("Seleccione una opcion: ")

		// Lee la opcion y comprueba que sea un numero.
		_, err := fmt.Scan(&opcion)

		// Si no puede leer un numero, termina el programa.
		if err != nil {
			fmt.Println("Error al leer la opcion.")
			return
		}

		// switch selecciona el modulo que se ejecutara.
		switch opcion {

		case 1:
			// Ejecuta la gestion del paquete clientes.
			clientes.GestionarClientes()

		case 2:
			// Ejecuta la gestion del paquete plantas.
			plantas.GestionarPlantas()

		case 3:
			// Ejecuta la gestion del paquete pedidos.
			pedidos.GestionarPedidos()

		case 4:
			// Ejecuta la gestion del paquete produccion.
			produccion.GestionarProduccion()

		case 5:
			// Ejecuta la gestion del paquete pagos.
			pagos.GestionarPagos()

		case 6:
			// Ejecuta la gestion del paquete entregas.
			entregas.GestionarEntregas()

		case 0:
			// Muestra el mensaje de despedida.
			fmt.Println("Gracias por utilizar SIGEP.")

			// Finaliza la funcion main y el programa.
			return

		default:
			// Informa cuando la opcion no existe.
			fmt.Println("Opcion no valida.")
		}
	}
}
