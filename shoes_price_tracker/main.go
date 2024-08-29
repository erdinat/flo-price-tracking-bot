package main

import (
	"github.com/erdinat/shoes_price_tracker/webScrapping"
)

func main() {
	//db := connect.Conn()

	//check := model.CheckUser(db, "ahmet", "blabla")
	//if check {
	//fmt.Println("successfully login")
	//return
	//}
	//fmt.Println("error login")
	//return

	webScrapping.Scraper()
}

//bir tane api olacak /login diye istek geldiğinde username ve password post veya get olarak alınacak ve kontrolleri CheckUser
// fonksiyonu içinde yapılacak eğer kullanıcı varsa success => true, yoksa success => false, message => kullanıcı bulunmadı olarak dönecek.

//kayıt ol => sign up diye bir sekme olucak
// adam username giricek mail adresi ve password giricek bunları users tablosuna insert edeceksin
