package service

import (
	"github.com/teten-nugraha/golang-crud-fiber/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetProducts()([]model.Product, error) {
	var products []model.Product

	db, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	if err != nil {
		return products, err
	}

	db.Find(&products)

	return products, nil
}

func CreateProduct(kode_barang string, nama string, supplier string, stok int) (model.Product, error) {
	
	var newProduct = model.Product{
		KodeBarang: kode_barang,
		Nama:       nama,
		Supplier:   supplier,
		Stok:       stok,
	}

	db, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	if err != nil {
		return newProduct, err
	}

	db.Create(&newProduct)

	return newProduct, nil
}

func FindProductById(id int) (model.Product, error) {

	var product model.Product

	db, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	if err != nil {
		return product, err
	}

	db.Where("id=?", id).Find(&product)

	return product, nil
}

func FindProductByKodeBarang(kode_barang string)(model.Product, error) {

	var product model.Product

	db, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	if err != nil {
		return product, err
	}

	db.Where("kode_barang=?", kode_barang).Find(&product)

	return product, nil
}

func UpdateProduct(id int, kode_barang string, nama string, supplier string, stok int) (model.Product, error) {

	product, err := FindProductById(id)
	if err != nil {
		return product, err
	}

	db, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	if err != nil {
		return product, err
	}

	db.Model(&product).Where("id=?", id).Updates(model.Product{
		KodeBarang: kode_barang,
		Nama:       nama,
		Supplier:   supplier,
		Stok:       stok,
	})

	return product, nil
}

func DeleteProduct(id int) error {

	db, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	product, err1 := FindProductById(id)
	if err1 != nil {
		return err1
	}

	db.Where("id=?", id).Delete(&product)

	return nil
}