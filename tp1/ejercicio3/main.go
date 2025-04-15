// Implemente una estructura Alumno con nombre, una lista de notas y un método Promedio()
// que devuelve el promedio de notas. Escriba un programa que permita obtener el promedio
// de varios alumnos. Sugerencia: no es necesario que cargue los datos de los alumnos, puede
// definirlos al crearlos.
package main

import (
	"fmt"
)

type Alumno struct {
	nombre string
	notas  []float64
}

func (a Alumno) Promedio() float64 {

	promedio := 0.0
	for _, nota := range a.notas {
		promedio += nota
	}
	if promedio > 0 {
		return promedio / float64(len(a.notas))
	}
	return 0
}

func calcularNotasAlumnos(alumnos []Alumno) float64 {

	promedio := 0.0

	for _, alumno := range alumnos {
		promedio += alumno.Promedio()
	}
	return promedio / float64(len(alumnos))
}

func main() {

	alumnos := []Alumno{
		{
			nombre: "Juan Pérez",
			notas:  []float64{8, 9, 10},
		},
		{
			nombre: "María López",
			notas:  []float64{7, 8.5, 9},
		},
		{
			nombre: "Carlos García",
			notas:  []float64{6, 7, 8},
		},
		{
			nombre: "Ana Fernández",
			notas:  []float64{9, 9.5, 10},
		},
		{
			nombre: "Luis Martínez",
			notas:  []float64{5, 6, 7},
		},
	}
	fmt.Println("El promedio de las notas de todos los alumnos es de: ", calcularNotasAlumnos(alumnos))
}
