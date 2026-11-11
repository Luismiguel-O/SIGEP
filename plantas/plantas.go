package plantas // Indica que este archivo pertenece al paquete plantas.

import "fmt" // Importa fmt para mostrar informacion en pantalla.

// RegistrarPlanta recibe el nombre, la variedad y el precio.
// string representa texto y float64 representa numeros decimales.
// Esta funcion es una demostracion: todavia no guarda datos.
func RegistrarPlanta(nombre string, variedad string, precio float64) {

	// Muestra un mensaje indicando que se recibieron los datos.
	fmt.Println("Datos de la planta:")

	// Muestra el nombre de la especie, por ejemplo: Tomate.
	fmt.Println("Planta:", nombre)

	// Muestra la variedad, por ejemplo: Marimar.
	fmt.Println("Variedad:", variedad)

	// Muestra el precio de venta por cada planta.
	fmt.Printf("Precio unitario: $%.2f\n", precio)
}

// ValidarPrecio recibe el precio unitario de una planta.
// Devuelve true si el precio es positivo y false si no lo es.
func ValidarPrecio(precio float64) bool {

	// El operador > significa "mayor que".
	// El precio debe ser mayor que cero.
	return precio > 0
}

// CalcularValor recibe la cantidad de plantas y el precio unitario.
// Devuelve el valor total de la venta como un numero decimal.
func CalcularValor(cantidad int, precio float64) float64 {

	// Convierte la cantidad de int a float64 para multiplicarla.
	// El resultado representa cantidad por precio unitario.
	return float64(cantidad) * precio
}

// ConsultarPlanta recibe el nombre de una planta.
// Por ahora devuelve un mensaje, no busca en un catalogo real.
func ConsultarPlanta(nombre string) string {

	// Devuelve un texto indicando la consulta solicitada.
	return "Consulta pendiente de la planta: " + nombre
}

// GestionarPlantas solicita los datos de una planta.
func GestionarPlantas() {
	var nombre, variedad string // Variables de texto.
	var precio float64          // Variable decimal.

	fmt.Println("===== GESTION DE PLANTAS =====")

	fmt.Print("Ingrese la planta (sin espacios): ")
	fmt.Scan(&nombre) // Lee la especie.

	fmt.Print("Ingrese la variedad (sin espacios): ")
	fmt.Scan(&variedad) // Lee la variedad.

	fmt.Print("Ingrese el precio unitario: ")
	fmt.Scan(&precio) // Lee el precio.

	if ValidarPrecio(precio) { // Comprueba que sea positivo.
		RegistrarPlanta(nombre, variedad, precio) // Muestra los datos.
	} else {
		fmt.Println("Error: el precio debe ser mayor que cero.")
	}
}
