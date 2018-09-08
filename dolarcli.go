package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/alexeyco/simpletable"
)

var (
	data = [][]interface{}{}
)

func main() {

	var docCompVal float64
	var docVentVal float64

	// Fetch & Parse Banco Nacion Information
	doc, err := goquery.NewDocument("https://www.bna.com.ar/Cotizador/MonedasHistorico")

	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#cotizacionesCercanas > table > tbody > tr:nth-child(1) > td").Each(func(index int, item *goquery.Selection) {

		if index == 1 {
			docCompra := item.Text()

			_docCompVal, err := strconv.ParseFloat(docCompra, 64)

			if err != nil {
				log.Fatal(err)
			} else {
				docCompVal = _docCompVal
			}
		}

		if index == 2 {
			docVenta := item.Text()
			_docVentVal, err := strconv.ParseFloat(docVenta, 64)

			if err != nil {
				log.Fatal(err)
			} else {
				docVentVal = _docVentVal
			}
		}
	})

	//Append Banco Nacion oficial value
	data = append(data, []interface{}{"Banco Nacion", docCompVal, docVentVal, docVentVal - docCompVal})

	//https://www.cronista.com/MercadosOnline/dolar.html
	//#dventa0
	docCronista, err := goquery.NewDocument("https://www.cronista.com/MercadosOnline/dolar.html")
	if err != nil {
		log.Fatal(err)
	}

	docCronista.Find("td#dcompra1").Each(func(index int, item *goquery.Selection) {
		docCompra := strings.Replace(strings.Replace(strings.Replace(item.Text(), " ", "", -1), "$", "", -1), ",", ".", -1)
		_docCompVal, err := strconv.ParseFloat(docCompra, 64)

		if err != nil {
			log.Fatal(err)
		} else {
			docCompVal = _docCompVal
		}

	})

	docCronista.Find("td#dventa1").Each(func(index int, item *goquery.Selection) {
		docVenta := strings.Replace(strings.Replace(strings.Replace(item.Text(), " ", "", -1), "$", "", -1), ",", ".", -1)
		_docVentVal, err := strconv.ParseFloat(docVenta, 64)

		if err != nil {
			log.Fatal(err)
		} else {
			docVentVal = _docVentVal
		}
	})

	data = append(data, []interface{}{"Diario El Cronista", docCompVal, docVentVal, docVentVal - docCompVal})

	// http://www.ambito.com/economia/mercados/monedas/dolar/
	// #contenido > div.row > div:nth-child(2) > div > div > div.cierreAnterior > big
	docAmbito, err := goquery.NewDocument("http://www.ambito.com/economia/mercados/monedas/dolar/")
	if err != nil {
		log.Fatal(err)
	}

	docAmbito.Find("#contenido > div.row > div:nth-child(2) > div > div > div.ultimo > big").Each(func(index int, item *goquery.Selection) {
		docCompra := strings.Replace(item.Text(), ",", ".", -1)
		_docCompVal, err := strconv.ParseFloat(docCompra, 64)

		if err != nil {
			log.Fatal(err)
		} else {
			docCompVal = _docCompVal
		}
	})

	docAmbito.Find("#contenido > div.row > div:nth-child(2) > div > div > div.cierreAnterior > big").Each(func(index int, item *goquery.Selection) {
		docVenta := strings.Replace(item.Text(), ",", ".", -1)
		_docVentVal, err := strconv.ParseFloat(docVenta, 64)

		if err != nil {
			log.Fatal(err)
		} else {
			docVentVal = _docVentVal
		}
	})

	data = append(data, []interface{}{"Diario Ambito", docCompVal, docVentVal, docVentVal - docCompVal})

	//https://www.oficialhoy.com.ar/
	// #Blog1 > article > div.post-entry > div > div:nth-child(2) > center:nth-child(3) > table:nth-child(19) > tbody > tr:nth-child(1) > td:nth-child(3)
	docOficialHoy, err := goquery.NewDocument("https://www.oficialhoy.com.ar/")
	if err != nil {
		log.Fatal(err)
	}

	docOficialHoy.Find("#Blog1 > article > div.post-entry > div > div:nth-child(2) > center:nth-child(3) > table:nth-child(22) > tbody > tr:nth-child(1) > td:nth-child(2)").Each(func(index int, item *goquery.Selection) {
		docCompra := strings.Replace(strings.Replace(strings.Replace(strings.Replace(item.Text(), " ", "", -1), "$", "", -1), ",", ".", -1), "\u00a0", "", -1)
		_docCompVal, err := strconv.ParseFloat(docCompra, 64)

		if err != nil {
			log.Fatal(err)
		} else {
			docCompVal = _docCompVal
		}
	})

	docOficialHoy.Find("#Blog1 > article > div.post-entry > div > div:nth-child(2) > center:nth-child(3) > table:nth-child(22) > tbody > tr:nth-child(1) > td:nth-child(3)").Each(func(index int, item *goquery.Selection) {
		docVenta := strings.Replace(strings.Replace(strings.Replace(strings.Replace(item.Text(), " ", "", -1), "$", "", -1), ",", ".", -1), "\u00a0", "", -1)
		_docVentVal, err := strconv.ParseFloat(docVenta, 64)

		if err != nil {
			log.Fatal(err)
		} else {
			docVentVal = _docVentVal
		}
	})

	data = append(data, []interface{}{"Banco Central (BCRA)", docCompVal, docVentVal, docVentVal - docCompVal})

	docOficialHoy.Find("#Blog1 > article > div.post-entry > div > div:nth-child(2) > center:nth-child(3) > table:nth-child(19) > tbody > tr:nth-child(1) > td:nth-child(2)").Each(func(index int, item *goquery.Selection) {
		docCompra := strings.Replace(strings.Replace(strings.Replace(strings.Replace(item.Text(), " ", "", -1), "$", "", -1), ",", ".", -1), "\u00a0", "", -1)
		_docCompVal, err := strconv.ParseFloat(docCompra, 64)

		if err != nil {
			log.Fatal(err)
		} else {
			docCompVal = _docCompVal
		}
	})

	docOficialHoy.Find("#Blog1 > article > div.post-entry > div > div:nth-child(2) > center:nth-child(3) > table:nth-child(19) > tbody > tr:nth-child(1) > td:nth-child(3)").Each(func(index int, item *goquery.Selection) {
		docVenta := strings.Replace(strings.Replace(strings.Replace(strings.Replace(item.Text(), " ", "", -1), "$", "", -1), ",", ".", -1), "\u00a0", "", -1)
		_docVentVal, err := strconv.ParseFloat(docVenta, 64)

		if err != nil {
			log.Fatal(err)
		} else {
			docVentVal = _docVentVal
		}
	})

	data = append(data, []interface{}{"Banco Santander Rio", docCompVal, docVentVal, docVentVal - docCompVal})

	docOficialHoy.Find("#Blog1 > article > div.post-entry > div > div:nth-child(2) > center:nth-child(3) > table:nth-child(14) > tbody > tr:nth-child(1) > td:nth-child(2)").Each(func(index int, item *goquery.Selection) {
		docCompra := strings.Replace(strings.Replace(strings.Replace(strings.Replace(item.Text(), " ", "", -1), "$", "", -1), ",", ".", -1), "\u00a0", "", -1)
		_docCompVal, err := strconv.ParseFloat(docCompra, 64)

		if err != nil {
			log.Fatal(err)
		} else {
			docCompVal = _docCompVal
		}
	})

	docOficialHoy.Find("#Blog1 > article > div.post-entry > div > div:nth-child(2) > center:nth-child(3) > table:nth-child(14) > tbody > tr:nth-child(1) > td:nth-child(3)").Each(func(index int, item *goquery.Selection) {
		docVenta := strings.Replace(strings.Replace(strings.Replace(strings.Replace(item.Text(), " ", "", -1), "$", "", -1), ",", ".", -1), "\u00a0", "", -1)
		_docVentVal, err := strconv.ParseFloat(docVenta, 64)

		if err != nil {
			log.Fatal(err)
		} else {
			docVentVal = _docVentVal
		}
	})

	data = append(data, []interface{}{"Banco Galicia", docCompVal, docVentVal, docVentVal - docCompVal})

	docOficialHoy.Find("#Blog1 > article > div.post-entry > div > div:nth-child(2) > center:nth-child(3) > table:nth-child(12) > tbody > tr:nth-child(1) > td:nth-child(2)").Each(func(index int, item *goquery.Selection) {
		docCompra := strings.Replace(strings.Replace(strings.Replace(strings.Replace(item.Text(), " ", "", -1), "$", "", -1), ",", ".", -1), "\u00a0", "", -1)
		_docCompVal, err := strconv.ParseFloat(docCompra, 64)

		if err != nil {
			log.Fatal(err)
		} else {
			docCompVal = _docCompVal
		}
	})

	docOficialHoy.Find("#Blog1 > article > div.post-entry > div > div:nth-child(2) > center:nth-child(3) > table:nth-child(12) > tbody > tr:nth-child(1) > td:nth-child(3)").Each(func(index int, item *goquery.Selection) {
		docVenta := strings.Replace(strings.Replace(strings.Replace(strings.Replace(item.Text(), " ", "", -1), "$", "", -1), ",", ".", -1), "\u00a0", "", -1)
		_docVentVal, err := strconv.ParseFloat(docVenta, 64)

		if err != nil {
			log.Fatal(err)
		} else {
			docVentVal = _docVentVal
		}
	})

	data = append(data, []interface{}{"Banco Hipotecario", docCompVal, docVentVal, docVentVal - docCompVal})

	docOficialHoy.Find("#Blog1 > article > div.post-entry > div > div:nth-child(2) > center:nth-child(3) > table:nth-child(9) > tbody > tr:nth-child(1) > td:nth-child(2)").Each(func(index int, item *goquery.Selection) {
		docCompra := strings.Replace(strings.Replace(strings.Replace(strings.Replace(item.Text(), " ", "", -1), "$", "", -1), ",", ".", -1), "\u00a0", "", -1)
		_docCompVal, err := strconv.ParseFloat(docCompra, 64)

		if err != nil {
			log.Fatal(err)
		} else {
			docCompVal = _docCompVal
		}
	})

	docOficialHoy.Find("#Blog1 > article > div.post-entry > div > div:nth-child(2) > center:nth-child(3) > table:nth-child(9) > tbody > tr:nth-child(1) > td:nth-child(3)").Each(func(index int, item *goquery.Selection) {
		docVenta := strings.Replace(strings.Replace(strings.Replace(strings.Replace(item.Text(), " ", "", -1), "$", "", -1), ",", ".", -1), "\u00a0", "", -1)
		_docVentVal, err := strconv.ParseFloat(docVenta, 64)

		if err != nil {
			log.Fatal(err)
		} else {
			docVentVal = _docVentVal
		}
	})

	data = append(data, []interface{}{"Banco BBVA Frances", docCompVal, docVentVal, docVentVal - docCompVal})

	docOficialHoy.Find("#Blog1 > article > div.post-entry > div > div:nth-child(2) > center:nth-child(3) > table:nth-child(17) > tbody > tr:nth-child(1) > td:nth-child(2)").Each(func(index int, item *goquery.Selection) {
		docCompra := strings.Replace(strings.Replace(strings.Replace(strings.Replace(item.Text(), " ", "", -1), "$", "", -1), ",", ".", -1), "\u00a0", "", -1)
		_docCompVal, err := strconv.ParseFloat(docCompra, 64)

		if err != nil {
			log.Fatal(err)
		} else {
			docCompVal = _docCompVal
		}
	})

	docOficialHoy.Find("#Blog1 > article > div.post-entry > div > div:nth-child(2) > center:nth-child(3) > table:nth-child(17) > tbody > tr:nth-child(1) > td:nth-child(3)").Each(func(index int, item *goquery.Selection) {
		docVenta := strings.Replace(strings.Replace(strings.Replace(strings.Replace(item.Text(), " ", "", -1), "$", "", -1), ",", ".", -1), "\u00a0", "", -1)
		_docVentVal, err := strconv.ParseFloat(docVenta, 64)

		if err != nil {
			log.Fatal(err)
		} else {
			docVentVal = _docVentVal
		}
	})

	data = append(data, []interface{}{"Banco Ciudad", docCompVal, docVentVal, docVentVal - docCompVal})

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Fuente"},
			{Align: simpletable.AlignCenter, Text: "Compra"},
			{Align: simpletable.AlignCenter, Text: "Venta"},
			{Align: simpletable.AlignCenter, Text: "Spread"},
		},
	}

	subtotal := float64(0)
	for _, row := range data {
		r := []*simpletable.Cell{
			{Text: row[0].(string)},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("$%.4f", row[1].(float64))},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("$%.4f", row[2].(float64))},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("$%.4f", row[3].(float64))},
		}

		table.Body.Cells = append(table.Body.Cells, r)
		subtotal += row[2].(float64)
	}

	table.SetStyle(simpletable.StyleRounded)
	fmt.Println(table.String())
}
