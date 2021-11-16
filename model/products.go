package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	KodeBarang string `json:"kode_barang"`
	Nama string `json:"nama"`
	Supplier string `json:"supplier"`
	Stok int `json:"stok"`
}
