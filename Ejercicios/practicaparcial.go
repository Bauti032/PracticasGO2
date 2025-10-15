package main 

import (
	"fmt"
)

func main(){

	var(
		datos [][6]int
		res [6][7]int
		cf int
	)

datos, cf  = cargarM(datos [][6]int, cf)
res = generarRES(datos, cf)
nms := ave(res)
}


func cargarM ( datos [][6]int) [][6]int{
// Nro. de servicio (número correlativo que inicia en 1000 y NO debe ser ingresado por el usuario).
// Código de cliente (dueño del animal).
// Código clase de animal (1 = Perro, 2 = Gato, 3 = Conejo, 4 = Ave, 5 = Roedores o 6 = Otras especies).
// Código de servicio (1 = Consulta, 2 = Clínico, 3 = Preventivo, 4 = Etología o 5 = Vacunación).
// Importe cobrado (valor de tipo entero).
// Atención a domicilio (1 = SI, 2 = NO).
var(
	nrS int = 999
	cC int
	cA int 
	cS int
	iC int 
	aD int
	y int
)

for {
	fmt.Print("Ingrese numero de cliente: ")
	fmt.Scanln(&cC)
		if cC != 0 {
			datos[y+1][0] = nrS+1
			datos[y+1][1] = cC
			fmt.Print("Ingrese el codigo del animal: ")
			fmt.Scanln(&cA)
			datos[y+1][2] = cA
			fmt.Print("Ingrese el codigo de servicio: ")
			fmt.Scanln(&cS)
			datos[y+1][3] = cS
			fmt.Print("Ingrese el importe cobrado: ")
			fmt.Scanln(&iC)
			datos[y+1][4] = iC
			fmt.Print("Ingrese el codigo de atencion: ")
			fmt.Scanln(&aD)
			datos[y+1][5] = aD
			cf = cf+1
		}else {
			break
		}
	
}

return datos
}

func generarRES(datos [][6]int, cf int){
// b. Generar una matriz resumen RES[6][7] que guarde la cantidad total de servicios prestados según el código de
// servicio y el código de la clase del animal. servicio 1/5 y cod animal 1/6, 2 y 3 columnas de datos
var(
	res [6][7]int
	acF int
)

for x1:= 0; x1< 5; x1++{
	for y:=0; y<6; y++{
		for x:=0; x<cf; x++{
			if datos [x][2] == y+1  && datos [x][3] == x1+1{
				res[x1][y] += datos[x][3] 
				acF = acf + res[y][x]
			}else{}
		}
		res[x1][7] = acf
		acf = 0
	}
	
}
}

func ave(res [6][7]int) string{
var(
	mayA int
)


for x:=0; x<6; x++{
	if mayA < res [x][4]{
		res [x][4] = mayA
		mayNS := x
	}
}
// Código de servicio (1 = Consulta, 2 = Clínico, 3 = Preventivo, 4 = Etología o 5 = Vacunación).
if mayNS == 1{
	nms := "Consulta"
}else if mayNS == 2{
		nms := "Clinico"
}else if mayNS == 3{
		nms := "Preventivo"
}else if mayNS == 4{
		nms := "Etología"
}else{
		nms := "Vacunación"
}

return nms
}

func cantotalanimales(res [6][7]int){
// 	Generar un procedimiento que calcule e imprima los porcentajes que representan la cantidad total de animales
// atendidos según sus clases sobre el total general.

var(
ct int
vp [6]float
)


for x:= 0; x<6; x++{
	if x == 0  {
		c1 := float(res[5][x])
		ct += res[5][x]
	}else if x == 1  {
		c2 := float(res[5][x])
		ct += res[5][x]
	}else if x == 2  {
		c3 := float(res[5][x])
		ct += res[5][x]
	}else if x == 3  {
		c4 := float(res[5][x])
		ct += res[5][x]
	}else if x == 0  {
		c5 := float(res[5][x])
		ct += res[5][x]
	}else{
		c6 := float(res[5][x])
		ct += res[5][x]
	}
}

//porcentajes ct --> 100 
//            cn  --> x

//vector porcentajes

for x:=0; x<6;x++{
	vp[x] = c1*100/ct
}


}