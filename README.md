# DolarCli

Visto y considerando que en los tiempos que corren en Argentina el dolar es algo que tenemos que estar revisando cada 15 minutos es dificil tener que andar revisando todas las web para comparar los precios, es por eso que DolarCli es una pequeña y simple utilidad escrita en golang para obtener los valores de referencia en los principales bancos y diarios de la Argentina. 

# Compilacion y/o Instalación
Para compilar el programa es necesario tener instalado los siguientes paquetes:

 - golang 1.11 https://golang.org/dl/
 - dep https://golang.github.io/dep/docs/installation.html

Una vez instalados, clonar el repositorio y ejecutar el siguiente comando para compilar y ejecutar:

    > env GOOS=[OS] GOARCH=[ARCH] go build -v -o dolarcli dolarcli.go
    > ./dolarcli
    .----------------------.----------.----------.---------.
	|        Fuente        |  Compra  |  Venta   | Spread  |
	+----------------------+----------+----------+---------+
	| Banco Nacion         | $36.7900 | $36.9900 | $0.2000 |
	| Diario El Cronista   | $36.0000 | $36.5000 | $0.5000 |
	| Diario Ambito        | $36.5000 | $37.5000 | $1.0000 |
	| Banco Central (BCRA) | $36.0660 | $37.8110 | $1.7450 |
	| Banco Santander Rio  | $35.9900 | $38.0400 | $2.0500 |
	| Banco Galicia        | $36.0000 | $38.0000 | $2.0000 |
	| Banco Hipotecario    | $35.9000 | $37.9000 | $2.0000 |
	| Banco BBVA Frances   | $35.7600 | $37.7400 | $1.9800 |
	| Banco Ciudad         | $36.2500 | $37.7500 | $1.5000 |
	'----------------------'----------'----------'---------'

Donde [OS] y [ARCH] corresponde al sistema operativo y la arquitectura donde se va a correr el mismo, para referencia revisar el siguiente link https://golang.org/doc/install/source#environment

Para instalar el mismo es necesario ejecutar el siguiente comando:

    go install
    
