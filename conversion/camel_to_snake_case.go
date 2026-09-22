package main

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"time"
)

// Reads request.json, converts its camelCase field names to snake_case (dropping
// any "CustomInfo" fields and parsing "centAmount" strings into ints), and writes response.json.
// this func is needed when we need to handle truly arbitrary/unknown JSON shapes
func camel_to_snake_case() {
	jsonData, err := os.ReadFile("request.json")
	if err != nil {
		panic(err)
	}

	var data any
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		panic(err)
	}

	convertedData := convertFieldNames(data)

	// Marshal the converted data back to JSON
	convertedJSON, err := json.MarshalIndent(convertedData, "", " ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("response.json", convertedJSON, 0600)
	if err != nil {
		panic(err)
	}
}

func convertFieldNames(data any) any {
	switch d := data.(type) {
	case map[string]any:
		newMap := make(map[string]any)
		for key, val := range d {
			if strings.Contains(key, "CustomInfo") {
				continue
			}

			newVal := convertFieldNames(val)
			if key == "centAmount" {
				centAmt, err := strconv.Atoi(newVal.(string))
				if err != nil {
					centAmt = 0
				}

				snakeCaseKey := camelToSnakeCase(key)
				newMap[snakeCaseKey] = centAmt
				continue
			}

			snakeCaseKey := camelToSnakeCase(key)
			newMap[snakeCaseKey] = newVal
		}

		return newMap
	case []any:
		newSlice := make([]any, 0, len(d))
		for _, val := range d {
			newSlice = append(newSlice, convertFieldNames(val))
		}

		return newSlice
	default:
		return data
	}
}

// Function to convert camel case to snake case
func camelToSnakeCase(s string) string {
	var result strings.Builder

	for i, ch := range s {
		if i > 0 && (ch >= 'A' && ch <= 'Z') {
			result.WriteByte('_')
		}
		result.WriteRune(ch)
	}

	return strings.ToLower(result.String())
}

type CustomTime time.Time

func (t *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	parsed, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return err
	}
	*t = CustomTime(parsed)
	return nil
}

func (t CustomTime) MarshalJSON() ([]byte, error) {
	// json.Marshal(time.Time(t).Format("2006-01-02 15:04:05")) for custom layout
	return time.Time(t).MarshalJSON()
}

func (o Order) MarshalJSON() ([]byte, error) {
	type AmountAlias struct {
		CentAmount   float32 `json:"cent_amount"`
		CurrencyCode string  `json:"currency_code"`
	}

	type AddressAlias struct {
		StreetName   string `json:"street_name"`
		StreetNumber string `json:"street_number"`
		PostalCode   string `json:"postal_code"`
		CityName     string `json:"city_name"`
		CountryCode  string `json:"country_code"`
	}

	type LineItemAlias struct {
		ProductId   string      `json:"product_id"`
		ProductName string      `json:"product_name"`
		Quantity    float32     `json:"quantity"`
		UnitPrice   AmountAlias `json:"unit_price"`
	}

	type OrderAlias struct {
		OrderID         string          `json:"order_id"`
		CustomerName    string          `json:"customer_name"`
		CustomerEmail   string          `json:"customer_email"`
		IsGiftWrapped   bool            `json:"is_gift_wrapped"`
		CreatedAt       CustomTime      `json:"created_at"`
		TotalPrice      AmountAlias     `json:"total_price"`
		ShippingAddress AddressAlias    `json:"shipping_address"`
		LineItems       []LineItemAlias `json:"line_items"`
	}

	return json.Marshal(OrderAlias{
		OrderID:       o.OrderID,
		CustomerName:  o.CustomerName,
		CustomerEmail: o.CustomerEmail,
		IsGiftWrapped: o.IsGiftWrapped,
		CreatedAt:     o.CreatedAt,
		TotalPrice: AmountAlias{
			CentAmount:   o.TotalPrice.CentAmount,
			CurrencyCode: o.TotalPrice.CurrencyCode,
		},
		ShippingAddress: AddressAlias{
			StreetName:   o.ShippingAddress.StreetName,
			StreetNumber: o.ShippingAddress.StreetNumber,
			PostalCode:   o.ShippingAddress.PostalCode,
			CityName:     o.ShippingAddress.CityName,
			CountryCode:  o.ShippingAddress.CountryCode,
		},

		LineItems: func() []LineItemAlias {
			lineItems := make([]LineItemAlias, 0, len(o.LineItems))

			for _, lineItem := range o.LineItems {
				lineItems = append(lineItems, LineItemAlias{
					ProductId:   lineItem.ProductId,
					ProductName: lineItem.ProductName,
					Quantity:    lineItem.Quantity,
					UnitPrice: AmountAlias{
						CentAmount:   lineItem.UnitPrice.CentAmount,
						CurrencyCode: lineItem.UnitPrice.CurrencyCode,
					},
				})
			}

			return lineItems
		}(),
	})
}

type Order struct {
	OrderID         string
	CustomerName    string
	CustomerEmail   string
	IsGiftWrapped   bool
	CreatedAt       CustomTime
	TotalPrice      Amount
	ShippingAddress Address
	LineItems       []LineItem
	OrderCustomInfo OrderCustomInfo
}

type Amount struct {
	CentAmount   float32 `json:"centAmount,string"`
	CurrencyCode string
}

type Address struct {
	StreetName   string
	StreetNumber string
	PostalCode   string
	CityName     string
	CountryCode  string
}

type LineItem struct {
	ProductId   string
	ProductName string
	Quantity    float32
	UnitPrice   Amount
}

type OrderCustomInfo struct {
	InternalNotes string
	RiskScore     int
}

func camelToSnakeCaseWhenJsonStructKnown() {
	data, err := os.ReadFile("request.json")
	if err != nil {
		panic(err)
	}

	var order Order

	if err := json.Unmarshal(data, &order); err != nil {
		panic(err)
	}

	respBytes, err := json.MarshalIndent(&order, "", " ")
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile("response.json", respBytes, 0600); err != nil {
		panic(err)
	}
}
