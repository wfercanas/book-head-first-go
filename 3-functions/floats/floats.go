package main

import "fmt"

/*
	1. Para formatear floats, se puede usar %f
	2. Para controlar mejor se pueden usar los parametros min-width (x) y total-decimals (y). Quedaría %x.yf
	3. Lo primero que hace el programa es evaluar y. Colocando la cantidad de decimales solicitados.
	4. La parte entera del número nunca se afecta.
	5. Se suma la cantidad de dígitos de la parte entera, más la cantidad de decimales, más el punto decimal, para saber cuántos caracteres se usaron.
	6. Si la suma de caracteres es mayor a x, el sistema no modifica nada e imprime el número.
	7. Si la suma de caracteres es menor a x, el sistema agrega espacios a la izquierda para que la longitud total, sumando caracteres y espacios, sea igual a x.
*/

func main() {
	value := 111222.3456789
	fmt.Printf("Simple: %f\n", value)                         // 111222.345679 -> Parte entera igual. Reduce la cantidad de decimales de 7 a 6
	fmt.Printf("Short min: %4f\n", value)                     // 111222.345679 -> Parte entera igual. Igual que no haber definido min-width
	fmt.Printf("Short min and short decimal: %4.1f\n", value) // 111222.3 -> Parte entera igual. Decimales reducidos de 7 a 1
	fmt.Printf("Short min and long decimal: %4.12f\n", value) // 111222.345678900005 -> Parte entera igual. Decimales ampliados de 7 a 12
	fmt.Printf("Long min: %12f\n", value)                     // 111222.345679 -> Parte entera igual. Igual que no haber definido min-width
	fmt.Printf("Long min and short decimal: %12.1f\n", value) //      111222.3 -> Parte entera igual. Decimales reducidos de 7 a 1. Para cumplir con la longitud de 12 se agregan espacios a la izq.
	fmt.Printf("Long min and long decimal: %12.15f\n", value) // 1111222.345678900004714 -> Parte entera igual. Sin espacio a la izq y con 15 decimales.
}
