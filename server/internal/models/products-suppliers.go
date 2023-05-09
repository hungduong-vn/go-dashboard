package models

import (
	"gorm.io/gorm"
)

// foreign key: product_id
type ProductsSuppliers struct {
	ProductID  uint32 `gorm:"column:product_id;primaryKey" json:"product_id"`
	SupplierID uint32 `gorm:"column:supplier_id;primaryKey" json:"supplier_id"`
	Stock      uint32 `gorm:"column:stock;not null" json:"stock"`
}

func SelectProductStocks(db *gorm.DB, id int) ([]ProductsSuppliers, error) {
	var stocks []ProductsSuppliers
	err := db.Where("product_id = ?", id).Find(&stocks).Error
	if err != nil {
		return nil, err
	}

	return stocks, nil
}

func CreateStock(db *gorm.DB, Stock *ProductsSuppliers) error {
	return db.Create(Stock).Error
}

func UpdateStock(db *gorm.DB, product_id uint32, supplier_id uint32, stock uint32) error {
	return db.Model(&ProductsSuppliers{}).Where("product_id = ? AND supplier_id = ?", product_id, supplier_id).Update("stock", stock).Error
}

func DeleteStock(db *gorm.DB, product_id uint32, supplier_id uint32) error {
	return db.Where("product_id = ? AND supplier_id = ?", product_id, supplier_id).Delete(&ProductsSuppliers{}).Error
}
