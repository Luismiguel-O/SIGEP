package pagos // Indica que este archivo pertenece al paquete pagos.

import "fmt" // Importa fmt para mostrar informacion en pantalla.

// RegistrarAbono recibe el numero de pedido y el valor abonado.
// numeroPedido identifica el pedido del cliente.
// monto representa la cantidad de dinero que entrega.
// Esta funcion solo muestra datos, todavia no guarda pagos.
func RegistrarAbono(numeroPedido int, monto float64) {

	// Muestra el encabezado del registro.
	fmt.Println("===== REGISTRO DE ABONO =====")

	// Muestra el numero del pedido.
	fmt.Println("Pedido Nro.:", numeroPedido)

	// Muestra el dinero abonado con dos decimales.
	fmt.Printf("Abono recibido: $%.2f\n", monto)
}

// ValidarAbono comprueba si un abono puede aceptarse.
// monto representa el dinero que desea pagar el cliente.
// saldo representa el dinero que todavia debe.
// Devuelve true si el abono es valido y false si no lo es.
func ValidarAbono(monto float64, saldo float64) bool {

	// El abono debe ser mayor que cero.
	// No puede superar el saldo pendiente.
	// && significa que ambas condiciones deben cumplirse.
	return monto > 0 && saldo > 0 && monto <= saldo
}

// CalcularSaldo determina cuanto dinero queda pendiente.
// total es el valor del pedido.
// abonos es la suma de todos los pagos recibidos.
// Devuelve el saldo pendiente como float64.
func CalcularSaldo(total float64, abonos float64) float64 {

	// Resta los abonos del valor total del pedido.
	saldo := total - abonos

	// Si el resultado es negativo, devuelve cero.
	if saldo < 0 {
		return 0
	}

	// Devuelve el dinero que el cliente sigue debiendo.
	return saldo
}

// ConsultarEstadoPago identifica el estado financiero.
// total representa el valor completo del pedido.
// abonos representa el total pagado por el cliente.
// Devuelve un mensaje de tipo string.
func ConsultarEstadoPago(total float64, abonos float64) string {

	// Comprueba que los valores recibidos sean validos.
	if total <= 0 || abonos < 0 || abonos > total {
		return "Datos de pago no validos"
	}

	// Si no existen abonos, el pedido esta pendiente.
	if abonos == 0 {
		return "Pendiente de pago"
	}

	// Si los abonos cubren todo el pedido, esta pagado.
	if abonos == total {
		return "Pagado completamente"
	}

	// En cualquier otro caso, el pago es parcial.
	return "Abonado parcialmente"
}

// GestionarPagos permite calcular abonos y saldos.
func GestionarPagos() {
	var numeroPedido int             // Identificador del pedido.
	var total, abonos, nuevo float64 // Valores monetarios.

	fmt.Println("===== GESTION DE PAGOS =====")

	fmt.Print("Ingrese el numero del pedido: ")
	fmt.Scan(&numeroPedido) // Lee el identificador.

	fmt.Print("Ingrese el valor total del pedido: ")
	fmt.Scan(&total) // Lee el total.

	fmt.Print("Ingrese los abonos anteriores: ")
	fmt.Scan(&abonos) // Lee los abonos.

	// Comprueba que los valores iniciales sean validos.
	if numeroPedido <= 0 || total <= 0 ||
		abonos < 0 || abonos > total {
		fmt.Println("Error: datos de pago incorrectos.")
		return // Termina esta funcion.
	}

	// Calcula el dinero pendiente de pago.
	saldo := CalcularSaldo(total, abonos)
	fmt.Printf("Saldo pendiente: $%.2f\n", saldo)

	fmt.Print("Ingrese el nuevo abono (0 si no hay): ")
	fmt.Scan(&nuevo) // Lee el nuevo abono.

	if nuevo == 0 { // Permite consultar sin pagar.
		fmt.Println("Estado:", ConsultarEstadoPago(total, abonos))
		return
	}

	if !ValidarAbono(nuevo, saldo) { // Comprueba el nuevo abono.
		fmt.Println("Error: el abono no es valido.")
		return
	}

	RegistrarAbono(numeroPedido, nuevo) // Muestra el abono recibido.

	abonos += nuevo // Suma el nuevo abono a los anteriores.

	saldo = CalcularSaldo(total, abonos) // Actualiza el saldo.

	fmt.Printf("Nuevo saldo pendiente: $%.2f\n", saldo)
	fmt.Println("Estado:", ConsultarEstadoPago(total, abonos))
}
