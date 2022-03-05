package oop

import (
	"fmt"
	"time"
)

type Product struct {
	barcodeNumber string
	price         int64
	name          string
}

func (p *Product) information() string {
	return fmt.Sprintf("Informasi Produk, nomor barcode: %s, harga: %d, nama: %s.", p.barcodeNumber, p.price, p.name)
}

type Ramen struct {
	Product
	variant     string
	isFried     bool
	color       string
	ExpiredDate time.Time
}

func (r *Ramen) PrintInformation() {
	fmt.Printf("%s\n", r.information())
	fmt.Printf("Varian: %s.\n", r.variant)
	fmt.Printf("Goreng: %t.\n", r.isFried)
	fmt.Printf("Warna: %s.\n", r.color)
	expired := r.ExpiredDate
	expiredString := fmt.Sprintf("%v.%v.%v", expired.Day(), int(expired.Month()), expired.Year())
	fmt.Printf("Tanggal Kedaluwarsa: %s.\n", expiredString)
}

func inheritance() {
	katsuRamen := Ramen{
		Product: Product{
			barcodeNumber: "123321",
			price:         25000,
			name:          "Katsu Ramen",
		},
		variant:     "cheese",
		isFried:     false,
		color:       "red",
		ExpiredDate: time.Now().Add(12 * 30 * 24 * time.Hour),
	}
	katsuRamen.PrintInformation()
}
