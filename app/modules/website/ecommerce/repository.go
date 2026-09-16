package ecommerce

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateCart(data *Cart) error {
	return config.DB.Create(data).Error
}

func GetAllCart() ([]Cart, error) {
	var list []Cart
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedCarts(offset, limit int, search string) ([]Cart, int64, error) {
	var list []Cart
	var total int64
	query := config.DB.Model(&Cart{}).Preload(clause.Associations)
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := query.Order("id desc").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetCartByID(id uint) (*Cart, error) {
	var data Cart
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateCart(data *Cart) error {
	return config.DB.Save(data).Error
}

func DeleteCart(id uint) error {
	return config.DB.Delete(&Cart{}, id).Error
}

func CreatePortalUser(data *PortalUser) error { return config.DB.Create(data).Error }
func GetAllPortalUser() ([]PortalUser, error) {
	var list []PortalUser
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetPortalUserByID(id uint) (*PortalUser, error) {
	var data PortalUser
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdatePortalUser(data *PortalUser) error { return config.DB.Save(data).Error }
func DeletePortalUser(id uint) error          { return config.DB.Delete(&PortalUser{}, id).Error }

func CreateShoppingCart(data *ShoppingCart) error { return config.DB.Create(data).Error }
func GetAllShoppingCart() ([]ShoppingCart, error) {
	var list []ShoppingCart
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetShoppingCartByID(id uint) (*ShoppingCart, error) {
	var data ShoppingCart
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateShoppingCart(data *ShoppingCart) error { return config.DB.Save(data).Error }
func DeleteShoppingCart(id uint) error            { return config.DB.Delete(&ShoppingCart{}, id).Error }

func CreateCartItem(data *CartItem) error { return config.DB.Create(data).Error }
func GetAllCartItem() ([]CartItem, error) {
	var list []CartItem
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetCartItemByID(id uint) (*CartItem, error) {
	var data CartItem
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateCartItem(data *CartItem) error { return config.DB.Save(data).Error }
func DeleteCartItem(id uint) error        { return config.DB.Delete(&CartItem{}, id).Error }
