package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	doc, err := goquery.NewDocument("https://www.bna.com.ar/Cotizador/MonedasHistorico")

	if err != nil {
		log.Fatal(err)
	}

	////*[@id="cotizacionesCercanas"]/table/tbody/tr[1]/td[3]

	// use CSS selector found with the browser inspector
	// for each, use index and item
	doc.Find("#cotizacionesCercanas > table > tbody > tr:nth-child(1) > td:nth-child(3)").Each(func(index int, item *goquery.Selection) {
		title := item.Text()
		fmt.Printf("Dolar Oficial: %s\n", title)
	})

	doc.Find("#cotizacionesCercanas > table > tbody > tr:nth-child(3) > td:nth-child(3)").Each(func(index int, item *goquery.Selection) {
		title := item.Text()
		fmt.Printf("Euro Oficial: %s\n", title)
	})

	//https://www.cronista.com/MercadosOnline/dolar.html
	//#dventa0
	docCronista, err := goquery.NewDocument("https://www.cronista.com/MercadosOnline/dolar.html")
	if err != nil {
		log.Fatal(err)
	}

	docCronista.Find("td#dventa1").Each(func(index int, item *goquery.Selection) {
		title := strings.Replace(strings.Replace(strings.Replace(item.Text(), " ", "", -1), "$", "", -1), ",", ".", -1)
		fmt.Printf("Dolar Blue Cronista: %s\n", title)
	})

	// http://www.ambito.com/economia/mercados/monedas/dolar/
	// #contenido > div.row > div:nth-child(2) > div > div > div.cierreAnterior > big
	docAmbito, err := goquery.NewDocument("http://www.ambito.com/economia/mercados/monedas/dolar/")
	if err != nil {
		log.Fatal(err)
	}

	docAmbito.Find("#contenido > div.row > div:nth-child(2) > div > div > div.cierreAnterior > big").Each(func(index int, item *goquery.Selection) {
		title := strings.Replace(item.Text(), ",", ".", -1)
		fmt.Printf("Dolar Blue Ambito: %s\n", title)
	})

}
