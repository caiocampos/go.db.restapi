# go.db.restapi

[![License](https://img.shields.io/github/license/caiocampos/go.db.restapi.svg)](LICENSE)

Servidor go utilizando Fiber e Parsley, com conexão ao MongoDB, e cache com Redis ou Valkey

## Executando:

Para executar o projeto é necessário o Go instalado e configurado, siga as instruções do site a seguir para configurar:

http://www.golangbr.org/doc/instalacao

Antes de executar modifique o arquivo config.toml para apontar para o MongoDB instalado.

Após instalar o Go e configurar o arquivo config.toml compile o código, utilize o seguinte comando para isso:

> go build cmd/main.go

E depois, para executar:

> ./main

Para configurar o MongoDB localmente siga o tutorial:

https://medium.com/@NetoVieiraLeo/instalando-e-configurando-o-mongodb-no-windows-b1d4e1e58911
