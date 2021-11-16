package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teten-nugraha/golang-crud-fiber/model"
	"github.com/teten-nugraha/golang-crud-fiber/service"
	"github.com/teten-nugraha/golang-crud-fiber/util"
	"net/http"
	"strconv"

	//"github.com/teten-nugraha/golang-crud-fiber/util"
	//"net/http"
)

func GetProducts(c *fiber.Ctx) error {

	var products []model.Product

	products, err := service.GetProducts()
	if err != nil {
		return util.ApiResponse(c, http.StatusBadRequest, "gagal mengambil products", nil)
	}

	return util.ApiResponse(c, http.StatusOK, "sukses mengambil products", products)
}

func SaveProducts(c *fiber.Ctx) error {

	product :=new(model.Product)

	err := c.BodyParser(product)
	if err != nil {
		return util.ApiResponse(c, http.StatusBadRequest, "gagal create product", nil)
	}

	result, err := service.CreateProduct(product.KodeBarang, product.Nama, product.Supplier, product.Stok)
	if err != nil {
		return util.ApiResponse(c, http.StatusBadRequest, "gagal create product", nil)
	}

	return util.ApiResponse(c, http.StatusCreated, "sukses create product", result)
}

func FindProductById(c *fiber.Ctx) error {

	result := model.Product{}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return util.ApiResponse(c, http.StatusBadRequest, "gagal transform string ke int", nil)
	}
	result, err1 := service.FindProductById(id)
	if err1 != nil {
		return util.ApiResponse(c, http.StatusBadRequest, "ada kesalahan pada saat mengambil product", nil)
	}
	if (model.Product{}) == result {
		return util.ApiResponse(c, http.StatusNotFound, "product tidak ada", nil)
	}

	return util.ApiResponse(c, http.StatusOK, "sukses find  product", result)
}

func FindProductsByKodeBarang(c *fiber.Ctx) error {

	result := model.Product{}

	nama := c.Params("kode_barang")

	result, err1 := service.FindProductByKodeBarang(nama)
	if err1 != nil {
		return util.ApiResponse(c, http.StatusBadRequest, "ada kesalahan pada saat mengambil product", nil)
	}
	if (model.Product{}) == result {
		return util.ApiResponse(c, http.StatusNotFound, "product tidak ada", nil)
	}

	return util.ApiResponse(c, http.StatusOK, "sukses find product by kode", result)
}

func UpdateProduct(c *fiber.Ctx) error {

	product :=new(model.Product)

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return util.ApiResponse(c, http.StatusBadRequest, "gagal transform string ke int", nil)
	}

	err1 := c.BodyParser(product)
	if err1 != nil {
		return util.ApiResponse(c, http.StatusBadRequest, "gagal transform json product", nil)
	}

	result, err2 := service.UpdateProduct(id, product.KodeBarang, product.Nama, product.Supplier, product.Stok)
	if err2 != nil {
		return util.ApiResponse(c, http.StatusBadRequest, "gagal update product", nil)
	}

	return util.ApiResponse(c, http.StatusOK, "sukses update product", result)
}

func DeleteProductById(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return util.ApiResponse(c, http.StatusBadRequest, "gagal transform string ke int", nil)
	}

	err1 := service.DeleteProduct(id)
	if err1 != nil {
		return util.ApiResponse(c, http.StatusInternalServerError, "gagal delete product", nil)
	}

	return util.ApiResponse(c, http.StatusOK, "sukses delete product", nil)
}
