package entregas // Identifica este archivo como parte del paquete entregas.

import "fmt" // Permite mostrar informacion en la terminal.

// RegistrarEntrega recibe los datos de una entrega.
// numeroPedido identifica el pedido del cliente.
// calcula cantidad plantas solicitadas.
// calcula cortesia de las plantas adicionales gratuitas.
// fecha indica cuando se realiza la entrega.
// Por ahora solo muestra los datos, no los guarda.
func RegistrarEntrega(numeroPedido int, cantidad int, cortesia int, fecha string) {

	// Muestra el encabezado de la entrega.
	fmt.Println("===== REGISTRO DE ENTREGA =====")

	// Muestra el numero del pedido.
	fmt.Println("Pedido Nro.:", numeroPedido)

	// Muestra las plantas que solicito el cliente.
	fmt.Println("Plantas solicitadas:", cantidad)

	// Muestra las plantas adicionales gratuitas.
	fmt.Println("Plantas de cortesia:", cortesia)

	// Calcula cuantas plantas se entregaran en total.
	total := CalcularTotalEntregado(cantidad, cortesia)

	// Muestra el resultado del calculo anterior.
	fmt.Println("Total de plantas entregadas:", total)

	// Muestra la fecha de entrega.
	fmt.Println("Fecha de entrega:", fecha)
}

// ValidarCortesia verifica la cantidad de plantas gratuitas.
// Devuelve true cuando la cantidad esta entre 0 y 50.
// Devuelve false cuando es negativa o supera 50.
func ValidarCortesia(cortesia int) bool {

	// && significa que ambas condiciones deben cumplirse.
	return cortesia >= 0 && cortesia <= 50
}

// CalcularTotalEntregado suma las plantas del pedido
// y las plantas adicionales de cortesia.
// Devuelve el total como un numero entero.
func CalcularTotalEntregado(cantidad int, cortesia int) int {

	// Comprueba que la cantidad solicitada sea positiva.
	if cantidad <= 0 {
		return 0 // Devuelve cero si el pedido no es valido.
	}

	// Comprueba que la cortesia este dentro del limite.
	if !ValidarCortesia(cortesia) {
		return 0 // Devuelve cero si la cortesia no es valida.
	}

	// Suma las plantas solicitadas y las gratuitas.
	total := cantidad + cortesia

	// Devuelve el numero total de plantas entregadas.
	return total
}

// ConsultarEntrega recibe el numero de un pedido.
// Devuelve un mensaje de texto.
// La busqueda real se desarrollara posteriormente.
func ConsultarEntrega(numeroPedido int) string {

	// Sprintf permite incluir el numero dentro del mensaje.
	return fmt.Sprintf(
		"Consulta pendiente de entrega del pedido Nro. %d",
		numeroPedido,
	)
}

// GestionarEntregas solicita y valida una entrega.
func GestionarEntregas() {
	var numeroPedido int // Identificador del pedido.
	var cantidad int     // Plantas solicitadas.
	var cortesia int     // Plantas adicionales gratuitas.
	var fecha string     // Fecha de entrega.

	fmt.Println("===== GESTION DE ENTREGAS =====")

	fmt.Print("Ingrese el numero del pedido: ")
	fmt.Scan(&numeroPedido) // Lee el identificador.

	fmt.Print("Ingrese las plantas solicitadas: ")
	fmt.Scan(&cantidad) // Lee la cantidad.

	fmt.Print("Ingrese las plantas de cortesia (0 a 50): ")
	fmt.Scan(&cortesia) // Lee la cortesia.

	fmt.Print("Ingrese la fecha (DD-MM-AAAA): ")
	fmt.Scan(&fecha) // Lee la fecha.

	// Comprueba que los datos sean validos.
	if numeroPedido <= 0 || cantidad <= 0 ||
		!ValidarCortesia(cortesia) {
		fmt.Println("Error: datos de entrega incorrectos.")
		return
	}

	// Muestra el registro demostrativo de entrega.
	RegistrarEntrega(numeroPedido, cantidad, cortesia, fecha)

	// Calcula la cantidad total entregada.
	total := CalcularTotalEntregado(cantidad, cortesia)

	fmt.Println("===== RESUMEN =====")
	fmt.Println("Total de plantas entregadas:", total)
	fmt.Println("Las plantas de cortesia no se cobran.")
}
