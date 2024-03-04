package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	exibeIntroducao()
	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Encerrando...")
			os.Exit(0)
		default:
			fmt.Println("Comando desconhecido")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Douglas"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Esta é a versao", versao)
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("                                                                                     ")
	return comandoLido

}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do progama")

}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println('\n')
	}

	fmt.Println('\n')
}

func exibeNomes() {
	nomes := []string{"Douglas", "Daniel", "Bernardo"}
	fmt.Println(nomes)
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, "foi carregado com sucesso! ")
		registraLog(site, true)
	} else {
		fmt.Print("Site: ", site, "esta com problemas. Status code : ", resp.StatusCode)
		registraLog(site, false)
	}
}
func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " +
		strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("err")
	}

	fmt.Println(string(arquivo))

}
