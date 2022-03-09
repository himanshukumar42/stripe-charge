package controller

import (
	"fmt"
	"net/http"

	Service "github.com/himanshuk42/stripe-charge/src/service"

	"github.com/gin-gonic/gin"
	"github.com/himanshuk42/stripe-charge/src/models"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func Payment(c *gin.Context) {
	var payment models.Charge
	c.BindJSON(&payment)

	apiKey := "sk_test_51K7GkVSDW6ziaLBqKJj1nzhalEozZH27ZiVDQsp0ALyYGatT088iERXqireJd7HVyJAgR9bJgQrLSahZLcg9bqcs006TJcc8UH"
	fmt.Println(apiKey + "asdasd")
	stripe.Key = apiKey
	_, err := charge.New(&stripe.ChargeParams{
		Amount:       stripe.Int64(payment.Amount),
		Currency:     stripe.String(string(stripe.CurrencyINR)),
		Description:  stripe.String(payment.ProductName),
		Source:       &stripe.SourceParams{Token: stripe.String("tok_visa")},
		ReceiptEmail: stripe.String(payment.ReceiptEmail),
		Shipping: &stripe.ShippingDetailsParams{
			Address: &stripe.AddressParams{
				City:       stripe.String(payment.Shipping.Address.City),
				Country:    stripe.String(payment.Shipping.Address.Country),
				Line1:      stripe.String(payment.Shipping.Address.Line1),
				Line2:      stripe.String(payment.Shipping.Address.Line2),
				PostalCode: stripe.String(payment.Shipping.Address.PostalCode),
				State:      stripe.String(payment.Shipping.Address.State),
			},
			Name:           stripe.String(payment.Shipping.Name),
			Phone:          stripe.String(payment.Shipping.Phone),
			TrackingNumber: stripe.String(payment.Shipping.TrackingNumber),
		},
	})

	if err != nil {
		c.String(http.StatusBadRequest, "Payment Unsuccessfull")
		return
	}
	err1 := Service.SavePayment(&payment)
	if err1 != nil {
		c.String(http.StatusNotFound, "error occured")

	}

}
