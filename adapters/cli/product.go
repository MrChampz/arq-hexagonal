package cli

import (
	"fmt"

	"github.com/MrChampz/arq-hexagonal/application"
)

func Run(
	service application.ProductServiceInterface,
	action string,
	productId string,
	productName string,
	price float64,
) (string, error) {

	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf(
			"Product (ID: %s, name: %s, price: %f, status: %s) has been created.",
			product.GetID(),
			product.GetName(),
			product.GetPrice(),
			product.GetStatus(),
		)

	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		product, err = service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf(
			"Product (ID: %s, name: %s) has been enabled.",
			product.GetID(),
			product.GetName(),
		)

	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		product, err = service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf(
			"Product (ID: %s, name: %s) has been disabled.",
			product.GetID(),
			product.GetName(),
		)

	default:
		product, err := service.Get(productId)
		if err != nil {
			return result, nil
		}
		result = fmt.Sprintf(
			"Product \nID: %s\nName: %s\nPrice: %f\nStatus: %s",
			product.GetID(),
			product.GetName(),
			product.GetPrice(),
			product.GetStatus(),
		)
	}

	return result, nil
}
