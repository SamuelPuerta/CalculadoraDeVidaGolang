package main

import (
	"fmt"
)

type Persona struct {
	Edad int
}

func (p *Persona) ObtenerIngresos() float64 {
	switch {
	case p.Edad <= 13:
		fmt.Println("Tus ingresos mensuales son: 0")
		return 0.0
	case p.Edad <= 18:
		fmt.Println("Ingresa tu mesada: ")
		return Scanner()
	case p.Edad <= 25:
		fmt.Println("Ingresa tu salario o mesada: ")
		return Scanner()
	case p.Edad <= 65:
		fmt.Print("Ingresa tu salario: ")
		return Scanner()
	case p.Edad <= 100:
		fmt.Println("Ingresa tu pensión: ")
		return Scanner()
	default:
		fmt.Println("Error: Ingresa una edad válida")
		return 0.0
	}
}

func (p *Persona) ObtenerGastos() float64 {
	switch {
	case p.Edad <= 13:
		return 0.0
	case p.Edad <= 18:
		fmt.Println("Ingresa tus gastos en diversión: ")
		diversion := Scanner()
		fmt.Println("Ingresa tus gastos en comida: ")
		comida := Scanner()
		fmt.Println("Ingresa tus gastos en pareja: ")
		pareja := Scanner()
		gastos := diversion + comida + pareja
		return gastos
	case p.Edad <= 25:
		fmt.Println("Ingresa tus gastos en diversión: ")
		diversion := Scanner()
		fmt.Println("Ingresa tus gastos en comida: ")
		comida := Scanner()
		fmt.Println("Ingresa tus gastos en pareja: ")
		pareja := Scanner()
		fmt.Println("Ingresa tus gastos en transporte: ")
		transporte := Scanner()
		fmt.Println("Ingresa tus gastos en educación: ")
		educacion := Scanner()
		gastos := pareja + comida + diversion + transporte + educacion
		return gastos
	case p.Edad <= 65:
		fmt.Println("Ingresa tus gastos en diversión: ")
		diversion := Scanner()
		fmt.Println("Ingresa tus gastos en comida: ")
		comida := Scanner()
		fmt.Println("Ingresa tus gastos en pareja: ")
		pareja := Scanner()
		fmt.Println("Ingresa tus gastos en transporte: ")
		transporte := Scanner()
		fmt.Println("Ingresa tus gastos en educación: ")
		educacion := Scanner()
		fmt.Println("Ingresa tus gastos en salud: ")
		salud := Scanner()
		fmt.Println("Ingresa tus gastos en hijos: ")
		hijos := Scanner()
		fmt.Println("Ingresa tus gastos en arriendo: ")
		arriendo := Scanner()
		fmt.Println("Ingresa tus gastos en servicios: ")
		servicios := Scanner()
		gastos := pareja + comida + diversion + transporte + educacion + salud + hijos + arriendo + servicios
		return gastos
	case p.Edad <= 100:
		fmt.Println("Ingresa tus gastos en diversión: ")
		diversion := Scanner()
		fmt.Println("Ingresa tus gastos en comida: ")
		comida := Scanner()
		fmt.Println("Ingresa tus gastos en salud: ")
		salud := Scanner()
		gastos := salud + comida + diversion
		return gastos
	default:
		fmt.Println("Error: Ingresa una edad válida")
		return 0.0
	}
}

func (p Persona) CalcularEstimado() {
	ingresos := p.ObtenerIngresos()
	gastos := p.ObtenerGastos()
	total := (ingresos - gastos) * 12
	fmt.Printf("Tus ingresos mensuales son: %.2f $\n", ingresos)
	fmt.Printf("Tus gastos mensuales son: %.2f $\n", gastos)
	fmt.Printf("Tu estimado de dinero al año es: %.2f $\n", total)
}

func Scanner() float64 {
	var scanner float64
	_, err := fmt.Scanln(&scanner)
	if err != nil {
		fmt.Println("Error: Ingresa un número válido")
		fmt.Scanln()
		return Scanner()
	}
	return scanner
}

func main() {
	fmt.Println("Calculadora de estimado de dinero según la edad al año")
	fmt.Print("Ingresa tu edad: ")
	edad := Scanner()
	persona := Persona{Edad: int(edad)}
	persona.CalcularEstimado()
}
