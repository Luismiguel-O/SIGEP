package produccion // Identifica este archivo como parte del paquete produccion.

import (
	"fmt"  // Permite crear mensajes de texto.
	"math" // Permite realizar operaciones matematicas.
)

// CalcularBandejas recibe la cantidad de plantas solicitadas.
// capacidad indica cuantas cavidades tiene cada bandeja.
// Devuelve el numero de bandejas necesarias como un entero.
func CalcularBandejas(cantidad int, capacidad int) int {

	// Comprueba que los valores sean mayores que cero.
	if cantidad <= 0 || capacidad <= 0 {
		return 0 // Devuelve cero si los datos no son validos.
	}

	// Divide las plantas solicitadas para la capacidad.
	// math.Ceil redondea el resultado hacia arriba.
	bandejas := math.Ceil(float64(cantidad) / float64(capacidad))

	// Convierte el resultado decimal a un numero entero.
	return int(bandejas)
}

// PlanificarSiembra calcula las bandejas considerando la germinacion.
// porcentaje representa la germinacion esperada, entre 0 y 100.
func PlanificarSiembra(cantidad int, capacidad int, porcentaje float64) int {

	// Comprueba que cantidad y capacidad sean positivas.
	if cantidad <= 0 || capacidad <= 0 {
		return 0
	}

	// Verifica que el porcentaje este entre 0 y 100.
	if porcentaje <= 0 || porcentaje > 100 {
		return 0
	}

	// Convierte el porcentaje a su equivalente decimal.
	germinacion := porcentaje / 100

	// Calcula cuantas semillas debemos sembrar aproximadamente.
	semillas := float64(cantidad) / germinacion

	// Calcula las bandejas necesarias y redondea hacia arriba.
	bandejas := math.Ceil(semillas / float64(capacidad))

	// Devuelve el numero entero de bandejas.
	return int(bandejas)
}

// CalcularPlantasGerminadas estima las plantas que germinaran.
// Recibe el numero de semillas sembradas y el porcentaje esperado.
func CalcularPlantasGerminadas(semillas int, porcentaje float64) int {

	// Comprueba que los valores recibidos sean validos.
	if semillas <= 0 || porcentaje < 0 || porcentaje > 100 {
		return 0
	}

	// Multiplica las semillas por el porcentaje de germinacion.
	resultado := float64(semillas) * porcentaje / 100

	// Redondea hacia abajo para obtener plantas enteras.
	return int(math.Floor(resultado))
}

// ConsultarProduccion recibe el numero de un pedido.
// Por ahora devuelve un mensaje de demostracion.
func ConsultarProduccion(numeroPedido int) string {

	// Sprintf permite incluir el numero dentro del texto.
	return fmt.Sprintf(
		"Consulta pendiente de produccion del pedido Nro. %d",
		numeroPedido,
	)
}

// GestionarProduccion calcula las bandejas para la siembra.
func GestionarProduccion() {
	var cantidad, capacidad int // Cantidad y cavidades por bandeja.
	var germinacion float64     // Porcentaje de germinacion.

	fmt.Println("===== PLANIFICACION DE PRODUCCION =====")

	fmt.Print("Ingrese las plantas solicitadas: ")
	fmt.Scan(&cantidad) // Lee la cantidad.

	fmt.Print("Ingrese las cavidades por bandeja: ")
	fmt.Scan(&capacidad) // Lee la capacidad.

	fmt.Print("Ingrese la germinacion estimada (1-100): ")
	fmt.Scan(&germinacion) // Lee el porcentaje.

	// Verifica que los datos sean validos.
	if cantidad <= 0 || capacidad <= 0 ||
		germinacion <= 0 || germinacion > 100 {
		fmt.Println("Error: datos de produccion incorrectos.")
		return // Termina esta funcion.
	}

	// Calcula las bandejas sin considerar perdidas.
	basicas := CalcularBandejas(cantidad, capacidad)

	// Calcula las bandejas considerando la germinacion.
	necesarias := PlanificarSiembra(cantidad, capacidad, germinacion)

	// Calcula el numero de semillas a sembrar.
	semillas := necesarias * capacidad

	// Calcula la cantidad estimada de plantas germinadas.
	estimadas := CalcularPlantasGerminadas(semillas, germinacion)

	fmt.Println("===== RESULTADOS =====")
	fmt.Println("Bandejas sin perdidas:", basicas)
	fmt.Println("Bandejas necesarias:", necesarias)
	fmt.Println("Semillas a sembrar:", semillas)
	fmt.Println("Plantas germinadas estimadas:", estimadas)
}
