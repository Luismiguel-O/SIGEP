package main // Indica que este archivo contiene el programa principal.

import (
	"fmt"            // Permite mostrar mensajes y leer datos.
	"sigep/clientes" // Importa nuestro paquete clientes.
	"sigep/entregas" // Permite utilizar las funciones del paquete entregas.
	"sigep/pagos"    // Importa el paquete pagos.
	"sigep/pedidos"  // Importa el paquete pedidos.
	"sigep/plantas"
	"sigep/produccion" // Importa las funciones de produccion.
)

// Permite mostrar mensajes y leer datos.
// Importa las funciones del paquete clientes.
// Importa las funciones del paquete plantas.

// main es la funcion donde comienza la ejecucion de SIGEP.
func main() {

	// Declara una variable para guardar la opcion del menu.
	var opcion int

	// for permite repetir el menu hasta que el usuario salga.
	for {

		// Muestra las opciones disponibles.
		fmt.Println("\n===== SISTEMA SIGEP =====")
		fmt.Println("1. Gestion de clientes")
		fmt.Println("2. Gestion de plantas")
		fmt.Println("3. Gestion de pedidos")
		fmt.Println("4. Gestion de produccion")
		fmt.Println("5. Gestion de pagos")
		fmt.Println("6. Gestion de entregas")
		fmt.Println("0. Salir")

		// Solicita una opcion al usuario.
		fmt.Print("Seleccione una opcion: ")

		// Lee el numero ingresado y lo guarda en opcion.
		fmt.Scan(&opcion)

		// switch ejecuta una accion segun la opcion elegida.
		switch opcion {

		case 1:
			// Declara variables para los datos del cliente.
			var nombre string
			var cedula string
			var telefono string

			// Solicita el nombre del cliente.
			fmt.Print("Ingrese el nombre (sin espacios): ")
			fmt.Scan(&nombre)

			// Solicita la cedula del cliente.
			fmt.Print("Ingrese la cedula: ")
			fmt.Scan(&cedula)

			// Solicita el telefono del cliente.
			fmt.Print("Ingrese el telefono: ")
			fmt.Scan(&telefono)

			// Llama a la funcion del paquete clientes.
			if clientes.ValidarCliente(nombre, cedula) {

				// Ejecuta el registro demostrativo del cliente.
				clientes.RegistrarCliente(nombre, cedula, telefono)

			} else {

				// Muestra un mensaje si los datos no son validos.
				fmt.Println("Error: nombre o cedula vacios.")
			}

		case 2:
			// Declara variables para almacenar los datos de la planta.
			var nombre string   // Guarda la especie, por ejemplo: Tomate.
			var variedad string // Guarda la variedad, por ejemplo: Marimar.
			var precio float64  // Guarda el precio unitario de la planta.

			// Solicita el nombre de la especie.
			fmt.Print("Ingrese la planta (sin espacios): ")
			fmt.Scan(&nombre) // Lee el nombre ingresado.

			// Solicita la variedad de la planta.
			fmt.Print("Ingrese la variedad (sin espacios): ")
			fmt.Scan(&variedad) // Lee la variedad ingresada.

			// Solicita el precio unitario.
			fmt.Print("Ingrese el precio unitario: ")
			fmt.Scan(&precio) // Lee el precio decimal.

			// Llama a ValidarPrecio del paquete plantas.
			if plantas.ValidarPrecio(precio) {

				// Ejecuta RegistrarPlanta con los datos ingresados.
				plantas.RegistrarPlanta(nombre, variedad, precio)

			} else {

				// Muestra un error si el precio no es positivo.
				fmt.Println("Error: el precio debe ser mayor que cero.")
			}

		case 3:
			// Declara las variables necesarias para registrar el pedido.
			var cliente string      // Nombre del cliente.
			var planta string       // Especie y variedad solicitada.
			var cantidad int        // Numero de plantas solicitadas.
			var precio float64      // Precio por cada planta.
			var fechaEntrega string // Fecha prevista para entregar.

			// Solicita el nombre del cliente.
			fmt.Print("Ingrese el nombre del cliente (sin espacios): ")
			fmt.Scan(&cliente) // Guarda el nombre ingresado.

			// Solicita la planta que desea comprar.
			fmt.Print("Ingrese la planta y variedad (sin espacios): ")
			fmt.Scan(&planta) // Guarda la planta ingresada.

			// Solicita la cantidad de plantas.
			fmt.Print("Ingrese la cantidad de plantas: ")
			fmt.Scan(&cantidad) // Guarda un numero entero.

			// Solicita el precio de cada planta.
			fmt.Print("Ingrese el precio unitario: ")
			fmt.Scan(&precio) // Guarda un numero decimal.

			// Solicita la fecha prevista para entregar el pedido.
			fmt.Print("Ingrese la fecha de entrega (DD-MM-AAAA): ")
			fmt.Scan(&fechaEntrega) // Guarda la fecha como texto.

			// Comprueba que la cantidad y el precio sean positivos.
			if pedidos.ValidarPedido(cantidad, precio) {

				// Muestra los datos mediante la funcion del paquete pedidos.
				pedidos.RegistrarPedido(cliente, planta, cantidad, precio, fechaEntrega)

				// Calcula el valor total utilizando otra funcion del paquete.
				total := pedidos.CalcularTotalPedido(cantidad, precio)

				// Muestra el total con dos cifras decimales.
				fmt.Printf("TOTAL DEL PEDIDO: $%.2f\n", total)

			} else {

				// Informa si la cantidad o el precio no son validos.
				fmt.Println("Error: la cantidad y el precio deben ser mayores que cero.")
			}

		case 4:
			// Declara las variables necesarias para planificar la siembra.
			var cantidad int        // Numero de plantas solicitadas.
			var capacidad int       // Numero de cavidades por bandeja.
			var germinacion float64 // Porcentaje estimado de germinacion.

			// Muestra el encabezado del modulo.
			fmt.Println("===== PLANIFICACION DE PRODUCCION =====")

			// Solicita la cantidad de plantas del pedido.
			fmt.Print("Ingrese la cantidad de plantas solicitadas: ")
			fmt.Scan(&cantidad) // Guarda la cantidad ingresada.

			// Solicita la capacidad de cada bandeja.
			fmt.Print("Ingrese la capacidad de cada bandeja: ")
			fmt.Scan(&capacidad) // Guarda el numero de cavidades.

			// Solicita el porcentaje estimado de germinacion.
			fmt.Print("Ingrese el porcentaje de germinacion (1-100): ")
			fmt.Scan(&germinacion) // Guarda el porcentaje ingresado.

			// Comprueba que los datos sean validos.
			if cantidad > 0 && capacidad > 0 &&
				germinacion > 0 && germinacion <= 100 {

				// Calcula las bandejas sin considerar perdidas.
				bandejasBasicas := produccion.CalcularBandejas(
					cantidad, capacidad,
				)

				// Calcula las bandejas considerando la germinacion.
				bandejasNecesarias := produccion.PlanificarSiembra(
					cantidad, capacidad, germinacion,
				)

				// Calcula cuantas semillas se sembraran en total.
				semillas := bandejasNecesarias * capacidad

				// Estima cuantas plantas germinaran.
				plantasEstimadas := produccion.CalcularPlantasGerminadas(
					semillas, germinacion,
				)

				// Muestra los resultados de la planificacion.
				fmt.Println("===== RESULTADOS =====")

				// Muestra las bandejas necesarias sin perdidas.
				fmt.Println("Bandejas sin considerar perdidas:", bandejasBasicas)

				// Muestra las bandejas considerando la germinacion.
				fmt.Println("Bandejas necesarias:", bandejasNecesarias)

				// Muestra el numero total de semillas a sembrar.
				fmt.Println("Semillas a sembrar:", semillas)

				// Muestra la cantidad estimada de plantas germinadas.
				fmt.Println("Plantas germinadas estimadas:", plantasEstimadas)

			} else {

				// Informa cuando se ingresan valores incorrectos.
				fmt.Println("Error: revise la cantidad, capacidad y germinacion.")
			}

		case 5:
			// Declara las variables necesarias para calcular los pagos.
			var numeroPedido int   // Identifica el pedido del cliente.
			var total float64      // Guarda el valor total del pedido.
			var abonos float64     // Guarda la suma de los abonos anteriores.
			var nuevoAbono float64 // Guarda el nuevo pago del cliente.

			// Muestra el encabezado del modulo.
			fmt.Println("===== GESTION DE PAGOS =====")

			// Solicita el numero del pedido.
			fmt.Print("Ingrese el numero del pedido: ")
			fmt.Scan(&numeroPedido) // Lee el identificador ingresado.

			// Solicita el valor total del pedido.
			fmt.Print("Ingrese el valor total del pedido: ")
			fmt.Scan(&total) // Lee el valor total.

			// Solicita el dinero que ya ha pagado el cliente.
			fmt.Print("Ingrese el total de abonos anteriores: ")
			fmt.Scan(&abonos) // Lee los abonos anteriores.

			// Comprueba que los valores ingresados sean coherentes.
			if numeroPedido > 0 && total > 0 && abonos >= 0 && abonos <= total {

				// Calcula el saldo antes de recibir otro abono.
				saldoAnterior := pagos.CalcularSaldo(total, abonos)

				// Muestra cuanto debe actualmente el cliente.
				fmt.Printf("Saldo pendiente: $%.2f\n", saldoAnterior)

				// Solicita el nuevo abono del cliente.
				fmt.Print("Ingrese el nuevo abono (0 si no hay): ")
				fmt.Scan(&nuevoAbono) // Lee el nuevo valor ingresado.

				// Permite consultar el saldo sin registrar otro abono.
				if nuevoAbono == 0 {

					// Muestra el estado financiero actual del pedido.
					fmt.Println("Estado:", pagos.ConsultarEstadoPago(total, abonos))

				} else if pagos.ValidarAbono(nuevoAbono, saldoAnterior) {

					// Muestra los datos del nuevo abono recibido.
					pagos.RegistrarAbono(numeroPedido, nuevoAbono)

					// Suma el nuevo abono a los pagos anteriores.
					totalAbonado := abonos + nuevoAbono

					// Calcula cuanto dinero queda por pagar.
					saldoNuevo := pagos.CalcularSaldo(total, totalAbonado)

					// Muestra el saldo actualizado.
					fmt.Printf("Nuevo saldo pendiente: $%.2f\n", saldoNuevo)

					// Muestra el estado del pago actualizado.
					fmt.Println("Estado:", pagos.ConsultarEstadoPago(total, totalAbonado))

				} else {

					// Informa cuando el abono no es valido.
					fmt.Println("Error: el abono supera el saldo o no es valido.")
				}

			} else {

				// Informa cuando los datos iniciales son incorrectos.
				fmt.Println("Error: revise el pedido, total y abonos anteriores.")
			}

		case 6:
			// Declara las variables para registrar la entrega.
			var numeroPedido int // Numero que identifica el pedido.
			var cantidad int     // Cantidad solicitada por el cliente.
			var cortesia int     // Plantas adicionales gratuitas.
			var fecha string     // Fecha en que se entregan las plantas.

			// Muestra el titulo del modulo.
			fmt.Println("===== GESTION DE ENTREGAS =====")

			// Solicita el numero del pedido.
			fmt.Print("Ingrese el numero del pedido: ")
			fmt.Scan(&numeroPedido) // Guarda el numero ingresado.

			// Solicita la cantidad de plantas del pedido.
			fmt.Print("Ingrese las plantas solicitadas: ")
			fmt.Scan(&cantidad) // Guarda la cantidad ingresada.

			// Solicita la cantidad de plantas gratuitas.
			fmt.Print("Ingrese las plantas de cortesia (0 a 50): ")
			fmt.Scan(&cortesia) // Guarda la cortesia ingresada.

			// Solicita la fecha de entrega.
			fmt.Print("Ingrese la fecha (DD-MM-AAAA): ")
			fmt.Scan(&fecha) // Guarda la fecha como texto.

			// Comprueba que los datos sean validos.
			if numeroPedido > 0 && cantidad > 0 &&
				entregas.ValidarCortesia(cortesia) {

				// Llama a la funcion RegistrarEntrega del paquete entregas.
				entregas.RegistrarEntrega(numeroPedido, cantidad, cortesia, fecha)

				// Calcula el total de plantas entregadas.
				total := entregas.CalcularTotalEntregado(cantidad, cortesia)

				// Muestra el resumen de la entrega.
				fmt.Println("===== RESUMEN =====")
				fmt.Println("Total de plantas entregadas:", total)

				// Informa que las plantas gratuitas no tienen costo.
				fmt.Println("Las plantas de cortesia no se cobran.")

			} else {

				// Informa cuando existen datos incorrectos.
				fmt.Println("Error: revise los datos de la entrega.")
			}

		case 0:
			// Muestra el mensaje de despedida.
			fmt.Println("Gracias por utilizar SIGEP.")

			// return termina la ejecucion del programa.
			return

		default:
			// Se ejecuta cuando la opcion no existe.
			fmt.Println("Opcion no valida.")
		}
	}
}
