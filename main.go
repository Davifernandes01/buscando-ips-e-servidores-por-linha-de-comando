package main

import (
	"cmd/app"
	"log"
	"os"
)


func main(){
	

	appli:= app.Gerar()

	
	if err:= appli.Run(os.Args); err != nil{
		log.Fatal(err)
	}

	
}