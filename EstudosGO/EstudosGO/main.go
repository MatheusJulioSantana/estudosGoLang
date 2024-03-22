package main

import (
	"EstudosGO/calcmc"
	"EstudosGO/calcmcub"
	"EstudosGO/n1n2perc"
	"fmt"
	"os"
)

func main() {
	mc, err := calcmc.NewCalcMQ(calcmc.WithAltura(10), calcmc.WithBase(10))
	if err != nil {
		fmt.Println("Erro ao criar objeto:", err)
		os.Exit(1)
	}
	fmt.Println("Area do retângulo:", mc.CalculaMQ())

	mcub, err := calcmcub.NewCalcMC(calcmcub.WithAltura(10), calcmcub.WithBase(10), calcmcub.WithLargura(10))
	if err != nil {
		fmt.Println("Erro ao criar o objeto:", err)
		os.Exit(1)
	}
	fmt.Println("A area em m3 do retangulo é:", mcub.CalculaMC())

	perc, err := n1n2perc.NewCalcPerc(n1n2perc.WithN1(10), n1n2perc.WithN2(4), n1n2perc.WithMulti(0.12))
	if err != nil {
		fmt.Println("Erro ao criar estrutura:", err)
		os.Exit(1)
	}
	fmt.Println("12% de 400 é:", perc.CalculaPerc())
}
