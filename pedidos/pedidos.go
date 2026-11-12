package pedidos // Identifica el paquete pedidos de SIGEP.

import "fmt" // Permite mostrar informacion en la terminal.

// RegistrarPedido recibe los datos de un pedido.
// cliente: nombre de la persona que solicita las plantas.
// planta: especie y variedad solicitada.
// cantidad: numero de plantas que necesita el cliente.
// precio: valor unitario acordado y calcula su valor total.
// fechaEntrega: fecha prevista para retirar las plantas.
// Por ahora solo muestra los datos, no los almacena.
func RegistrarPedido(cliente string, planta string, cantidad int, precio float64, fechaEntrega string) {

	// Muestra el encabezado del pedido.
	fmt.Println("===== REGISTRO DE PEDIDO =====")

	// Muestra el nombre del cliente.
	fmt.Println("Cliente:", cliente)

	// Muestra la planta y su variedad.
	fmt.Println("Planta:", planta)

	// Muestra la cantidad solicitada.
	fmt.Println("Cantidad:", cantidad)

	// Muestra el precio de cada planta.
	fmt.Printf("Precio unitario: $%.2f\n", precio)

	// Muestra la fecha programada de entrega.
	fmt.Println("Fecha de entrega:", fechaEntrega)
}

// ValidarPedido comprueba que los datos numericos sean validos.
// Devuelve true cuando cantidad y precio son positivos.
func ValidarPedido(cantidad int, precio float64) bool {

	// && significa que ambas condiciones deben cumplirse.
	return cantidad > 0 && precio > 0
}

// CalcularTotalPedido calcula el valor del pedido original.
// Solo cobra las plantas solicitadas por el cliente.
// Las plantas de cortesia no forman parte de este calculo.
func CalcularTotalPedido(cantidad int, precio float64) float64 {

	// Convierte cantidad a float64 y multiplica por el precio.
	return float64(cantidad) * precio
}

// ConsultarPedido recibe el numero identificador de un pedido.
// Devuelve un mensaje de texto de tipo string.
// La busqueda real se desarrollara en otra etapa.
func ConsultarPedido(numero int) string {

	// Sprintf construye un texto incluyendo el numero recibido.
	return fmt.Sprintf("Consulta pendiente del pedido Nro. %d", numero)
}

// GestionarPedidos solicita los datos y calcula el total.
func GestionarPedidos() {
	var cliente, planta, fechaEntrega string // Datos de texto.
	var cantidad int                         // Numero de plantas.
	var precio float64                       // Precio unitario.

	fmt.Println("===== GESTION DE PEDIDOS =====")

	fmt.Print("Ingrese el cliente (sin espacios): ")
	fmt.Scan(&cliente) // Lee el cliente.

	fmt.Print("Ingrese la planta y variedad (sin espacios): ")
	fmt.Scan(&planta) // Lee la planta.

	fmt.Print("Ingrese la cantidad: ")
	fmt.Scan(&cantidad) // Lee la cantidad.

	fmt.Print("Ingrese el precio unitario: ")
	fmt.Scan(&precio) // Lee el precio.

	fmt.Print("Ingrese la fecha de entrega (DD-MM-AAAA): ")
	fmt.Scan(&fechaEntrega) // Lee la fecha.

	if ValidarPedido(cantidad, precio) { // Comprueba el pedido.
		RegistrarPedido(cliente, planta, cantidad, precio, fechaEntrega)

		// Calcula el valor total del pedido.
		total := CalcularTotalPedido(cantidad, precio)

		// Muestra el resultado con dos decimales.
		fmt.Printf("TOTAL DEL PEDIDO: $%.2f\n", total)
	} else {
		fmt.Println("Error: cantidad o precio incorrectos.")
	}
}
