
# SIGEP - Sistema de Gestión de Pedidos de Plantas

## 1. Descripción del proyecto

SIGEP es un prototipo desarrollado en el lenguaje de
programación Go para apoyar la gestión de una pilonera dedicada a la producción y comercialización de plantas hortícolas.

El sistema está diseñado para organizar los procesos
relacionados con clientes, plantas, pedidos, producción, pagos y entregas.

Este proyecto corresponde a la primera etapa de
planificación y desarrollo modular de un sistema
de gestión empresarial.

## 2. Objetivo general

Desarrollar un prototipo de sistema de gestión de pedidos de plantas mediante el lenguaje Go, utilizando funciones y paquetes que permitan organizar los principales procesos administrativos y productivos de una pilonera.

## 3. Módulos del sistema

### Clientes
Permite demostrar el registro y la validación
de datos de clientes.

### Plantas
Permite mostrar información sobre especies,
variedades y precios unitarios de plantas.

### Pedidos
Permite registrar datos de demostración de pedidos
y calcular su valor total.

### Producción
Permite estimar la cantidad de bandejas necesarias
para cumplir un pedido, considerando el porcentaje
de germinación de las semillas.

### Pagos
Permite calcular abonos, saldos pendientes
y estados de pago de los pedidos.

### Entregas
Permite calcular las plantas entregadas,
incluyendo cantidades adicionales de cortesía
que no generan cobros.

## 4. Estructura del proyecto

SIGEP/
- main.go
- go.mod
- README.md
- clientes/clientes.go
- plantas/plantas.go
- pedidos/pedidos.go
- produccion/produccion.go
- pagos/pagos.go
- entregas/entregas.go

## 5. Tecnologías utilizadas

- Lenguaje de programación: Go.
- Editor de código: Visual Studio Code.
- Organización: paquetes y funciones.
- Control de versiones: Git y GitHub.

## 6. Ejecución del programa

Abrir una terminal dentro de la carpeta SIGEP.

Para ejecutar el programa:

    go run main.go

Para comprobar la compilación de los paquetes:

    go build ./...

## 7. Estado actual del proyecto

El proyecto cuenta con seis paquetes y un menú
principal que permite ejecutar funciones básicas.

Actualmente, el sistema permite realizar cálculos
y mostrar información de demostración.

En esta etapa no dispone de almacenamiento
permanente de clientes, pedidos, pagos o entregas.

## 8. Contexto académico

Proyecto individual desarrollado para la asignatura
Programación Orientada a Objetos, utilizando
el lenguaje de programación Go.
