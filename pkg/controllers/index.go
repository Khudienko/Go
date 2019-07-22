package controllers

import (
	"Pet/models"
	"Pet/pkg/driver"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := driver.Database.Query("SELECT * FROM productdb.Products")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	products := []models.Product{}

	for rows.Next() {
		p := models.Product{}
		err := rows.Scan(&p.Id, &p.Code, &p.Vendor,&p.Country, &p.Warranty)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}


	rows2,err2:= driver.Database.Query("select name,mobile,email,code,vendor,country,warranty FROM userproduct join users,products where userproduct.id_user=users.id and userproduct.id_product=products.id")
	if err2 != nil {
		log.Println(err)
	}
	defer rows2.Close()
	orders:=[]models.Orders{}

	for rows2.Next(){
		o:=models.Orders{}
		err:=rows2.Scan(&o.Name,&o.Mobile,&o.Email,&o.Code,&o.Vendor,&o.Country,&o.Warranty)
		switch o.Country {
		case "Hungary":
			o.Warranty += 24
		case "Germany":
			o.Warranty += 36
		case "USA":
			o.Warranty += 24
		case "Japanese":
			o.Warranty += 36
		default:
			o.Warranty+=12
		}


		if err != nil {
			fmt.Println(err)
			continue
		}
		orders = append(orders, o)
	}

	tmpl, _ := template.ParseFiles("web/index.html")
	tmpl.Execute(w, products)
	tmpl2, _ := template.ParseFiles("web/orders.html")
	tmpl2.Execute(w, orders)

}
