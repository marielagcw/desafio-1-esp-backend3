# Desafio 1

El presente repositorio tiene como finalidad trabajar en equipo para cumplir los requerimientos del primer examen parcial de la Especialidad Backend 3 desarrollado en Golang

# Enunciado

## Objetivo

A continuación se plantea un desafío integrador que nos permitirá evaluar los temas que
hemos visto hasta el momento.

## Planteo

Una aerolínea pequeña tiene un sistema de reservas de pasajes a diferentes países. Este retorna un archivo con la información de los pasajes sacados en las últimas 24 horas. La aerolínea necesita un programa para extraer información de las ventas del día y, así, analizar las tendencias de compra. El archivo en cuestión es del tipo valores separados por coma (CSV), donde los campos están compuestos por: id, nombre, email, país de destino, hora del vuelo y precio.

## Desafío

Realizar un programa que sirva como herramienta para calcular diferentes datos estadísticos. Para lograrlo,
deberás descargar este [repositorio](https://drive.google.com/file/d/1mPdZ_nSGJR_WCvxbh4Bpt7W7vtm3b1WB/view) que contiene un archivo CSV con datos generados y un esqueleto del proyecto.

## Requerimientos

### Requerimiento 1

- Una función que calcule cuántas personas viajan a un país determinado.

_Ejemplo:_

  `func GetTotalTickets(destination string) (int, error) {}`
  
_Tip: VS Code nos permite buscar una palabra en un archivo con Ctrl + F o ⌘ + F._

### Requerimiento 2

- Una o varias funciones que calculen cuántas personas viajan en madrugada (0 → 6), mañana (7 → 12), tarde (13 → 19), y noche (20 → 23).

_Ejemplo:_

`func GetCountByPeriod(time string) (int, error) {}`

_Tip: En Go, para manipular caracteres, tenemos el paquete strings._

### Requerimiento 3

- Calcular el porcentaje de personas que viajan a un país determinado en un día.

_Ejemplo:_

`func PercentageDestination(destination string, total int) (float64, error) {}`

_Tip: El porcentaje de x se calcula como: x̄ = ∑ x / n . 100_

### Requerimiento 4

- Ejecutar al menos una vez cada requerimiento en la función main. Las ejecuciones deben realizarse de manera concurrente (utilizando diferentes goroutines).

_Ejemplo:_

`go func(){
        totalTickets, err := GetTotalTickets(“Brazil”)
        // process all the data
}()`

_Tip: Se pueden enviar parámetros a la función anónima de la siguiente manera: go func(param string) { // function detail } (“param string”)_
