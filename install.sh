#!/usr/bin/env bash

# cores
normal=$'\e[0m'  
C=$(printf '\033')                                                 
green="${C}[1;32m"
yellow="${C}[1;33m"
RED="${C}[1;31m"
# fim das cores

instalacao(){
	echo "${green}[OK] ${normal}iniciando a instalação."
	echo "${yellow}[!!] ${normal}procurando path do go."

	if [ -x "$(command -v go)" ]; then
		echo "${green}[OK] ${normal}sucesso ao encontrar."
		go mod init pink.go
		go mod tidy		

	else
		echo "${RED} favor instalar golang em seu sistema."
}

instalacao