package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//
func Gerar() *cli.App{

	app:= cli.NewApp()
	app.Name = "app cmd"
	app.Usage = "buscar ips e servidores na internet"


	flags:= []cli.Flag{
		cli.StringFlag{

			//exemplo --host qualquerSite.com
			Name: "host",

			//caso não coloque nenhum valor
			Value: "https://davifernandes01.github.io/portfolio-daviFernandes/",
		},
	}
	//comandos 
	app.Commands = []cli.Command{

		{
			Name: "ip",
			Usage: "busca ips na internet",
			Flags: flags,
			Action: buscar,
		},
		{
			Name: "servidores",
			Usage: "buscar nomes de servidores",
			Flags: flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscar(c * cli.Context){

	host:=  c.String("host")

	ips, err := net.LookupIP(host)

	//se der algum erro, o programa para de rodar
	if err != nil {
		log.Fatal(err)
	}

	//para cada ip que o site tiver, ira printar na tela
	for _, ip := range ips{
		fmt.Println(ip)
	}
}

//mostrando os servidores
func buscarServidores(c *cli.Context){

	host:= c.String("host")

	servidores, err:= net.LookupNS(host)
	if err  != nil{
		log.Fatal(err)

	}

	//para cada servidor que o site tiver, ira printar na tela
	for _, serv := range servidores{
		fmt.Println(serv.Host)
	}

}