package clientes // Identifica este archivo como parte del paquete clientes.

import "fmt" // Importa fmt para mostrar mensajes en la terminal.

// RegistrarCliente recibe los datos de un cliente.
// Los parametros son nombre, cedula y telefono, todos de tipo string.
// Por ahora solo muestra los datos, no los guarda.
func RegistrarCliente(nombre string, cedula string, telefono string) {

	// Muestra un encabezado en la terminal.
	fmt.Println("Registro de cliente:")

	// Muestra el nombre recibido por la funcion.
	fmt.Println("Nombre:", nombre)

	// Muestra la cedula del cliente.
	fmt.Println("Cedula:", cedula)

	// Muestra el telefono del cliente.
	fmt.Println("Telefono:", telefono)
}

// ConsultarCliente recibe una cedula de tipo string.
// Devuelve un texto de tipo string.
// Todavia no realiza una busqueda real.
func ConsultarCliente(cedula string) string {

	// return devuelve el resultado a quien llama la funcion.
	return "Consulta pendiente para la cedula: " + cedula
}

// ValidarCliente recibe el nombre y la cedula.
// Devuelve un valor booleano: true o false.
func ValidarCliente(nombre string, cedula string) bool {

	// != significa "diferente de".
	// && significa "Y": ambas condiciones deben cumplirse.
	// true: ambos campos tienen contenido.
	// false: al menos uno esta vacio.
	return nombre != "" && cedula != ""
}

// GestionarClientes solicita y valida los datos del cliente.
func GestionarClientes() {
	var nombre, cedula, telefono string // Variables de texto.

	fmt.Println("===== GESTION DE CLIENTES =====")

	fmt.Print("Ingrese el nombre (sin espacios): ")
	fmt.Scan(&nombre) // Lee el nombre.

	fmt.Print("Ingrese la cedula: ")
	fmt.Scan(&cedula) // Lee la cedula.

	fmt.Print("Ingrese el telefono: ")
	fmt.Scan(&telefono) // Lee el telefono.

	if ValidarCliente(nombre, cedula) { // Comprueba los datos.
		RegistrarCliente(nombre, cedula, telefono) // Muestra el registro.
	} else {
		fmt.Println("Error: nombre o cedula vacios.")
	}
}
