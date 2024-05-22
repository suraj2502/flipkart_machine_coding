package constants

// {
// 	“Category”: “Mobiles”,
// 	“Percentage”: “10”,
// 	“MaxCap”: “50” 		//we pay min(10% of orderPrice, maxCap)
// 	}

type Category struct {
	CategoryName string
	CategoryType string
	MaxCap       int
}

var CategoryMap = map[string]Category{
	"Mobiles": Category{
		CategoryName: "Mobiles",
		CategoryType: "Percentage",
		MaxCap:       50,
	},
	"MobileCovers": Category{
		CategoryName: "MobileCovers",
		CategoryType: "FlatRate",
		MaxCap:       5,
	},
}

var cancelWindow = 10
var cancelWindowInHours = 24 * cancelWindow
