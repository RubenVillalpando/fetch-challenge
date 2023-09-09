package main

var POINTS_RECEIPT_1 = receipt{
	Retailer:     "Target",
	PurchaseDate: "2022-01-01",
	PurchaseTime: "13:01",
	Items: []item{
		{
			ShortDescription: "Mountain Dew 12PK",
			Price:            "6.49",
		}, {
			ShortDescription: "Emils Cheese Pizza",
			Price:            "12.25",
		}, {
			ShortDescription: "Knorr Creamy Chicken",
			Price:            "1.26",
		}, {
			ShortDescription: "Doritos Nacho Cheese",
			Price:            "3.35",
		}, {
			ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
			Price:            "12.00",
		},
	},
	Total: "35.35",
}
var POINTS_RECEIPT_2 = receipt{
	Retailer:     "M&M Corner Market",
	PurchaseDate: "2022-03-20",
	PurchaseTime: "14:33",
	Items: []item{
		{
			ShortDescription: "Gatorade",
			Price:            "2.25",
		}, {
			ShortDescription: "Gatorade",
			Price:            "2.25",
		}, {
			ShortDescription: "Gatorade",
			Price:            "2.25",
		}, {
			ShortDescription: "Gatorade",
			Price:            "2.25",
		},
	},
	Total: "9.00",
}

var LOWER_INBOUNDS_CASE_TIME = receipt{
	Retailer:     "M&M Corner Market",
	PurchaseDate: "2022-03-20",
	PurchaseTime: "14:01",
	Items: []item{
		{
			ShortDescription: "Gatorade",
			Price:            "2.25",
		},
	},
	Total: "2.25",
}
var LOWER_OUTBOUNDS_CASE_TIME = receipt{
	Retailer:     "M&M Corner Market",
	PurchaseDate: "2022-03-20",
	PurchaseTime: "14:00",
	Items: []item{
		{
			ShortDescription: "Gatorade",
			Price:            "2.25",
		},
	},
	Total: "2.25",
}
var LOWER_EDGE_CASE_TIME = receipt{
	Retailer:     "M&M Corner Market",
	PurchaseDate: "2022-03-20",
	PurchaseTime: "14:00",
	Items: []item{
		{
			ShortDescription: "Gatorade",
			Price:            "2.25",
		},
	},
	Total: "2.25",
}

var UPPER_INBOUNDS_CASE_TIME = receipt{
	Retailer:     "M&M Corner Market",
	PurchaseDate: "2022-03-20",
	PurchaseTime: "15:59",
	Items: []item{
		{
			ShortDescription: "Gatorade",
			Price:            "2.25",
		},
	},
	Total: "2.25",
}
var UPPER_EDGE_CASE_TIME = receipt{
	Retailer:     "M&M Corner Market",
	PurchaseDate: "2022-03-20",
	PurchaseTime: "16:00",
	Items: []item{
		{
			ShortDescription: "Gatorade",
			Price:            "2.25",
		},
	},
	Total: "2.25",
}
var UPPER_OUTBOUNDS_CASE_TIME = receipt{
	Retailer:     "M&M Corner Market",
	PurchaseDate: "2022-03-20",
	PurchaseTime: "16:01",
	Items: []item{
		{
			ShortDescription: "Gatorade",
			Price:            "2.25",
		},
	},
	Total: "2.25",
}
var ONE_ITEM_CASE_ITEM_QTTY = receipt{
	Retailer:     "M&M Corner Market",
	PurchaseDate: "2022-03-20",
	PurchaseTime: "16:01",
	Items: []item{
		{
			ShortDescription: "Gatorade",
			Price:            "2.25",
		},
	},
	Total: "2.25",
}
var TWO_ITEM_CASE_ITEM_QTTY = receipt{
	Retailer:     "M&M Corner Market",
	PurchaseDate: "2022-03-20",
	PurchaseTime: "16:01",
	Items: []item{
		{
			ShortDescription: "Gatorade",
			Price:            "2.25",
		},
		{
			ShortDescription: "Gatorade",
			Price:            "2.25",
		},
	},
	Total: "2.25",
}
var SEVEN_ITEM_CASE_ITEM_QTTY = receipt{
	Retailer:     "M&M Corner Market",
	PurchaseDate: "2022-03-20",
	PurchaseTime: "16:01",
	Items: []item{
		{
			ShortDescription: "Gatorade",
			Price:            "2.25",
		},
		{
			ShortDescription: "Gatorade",
			Price:            "2.25",
		},
		{
			ShortDescription: "Gatorade",
			Price:            "2.25",
		},
		{
			ShortDescription: "Gatorade",
			Price:            "2.25",
		},
		{
			ShortDescription: "Gatorade",
			Price:            "2.25",
		},
		{
			ShortDescription: "Gatorade",
			Price:            "2.25",
		},
		{
			ShortDescription: "Gatorade",
			Price:            "2.25",
		},
	},
	Total: "2.25",
}

var ALL_SPACES_CASE_SHORTDESC = receipt{
	Retailer:     "M&M Corner Market",
	PurchaseDate: "2022-03-20",
	PurchaseTime: "16:01",
	Items: []item{
		{
			ShortDescription: "	  	",
			Price:            "2.25",
		},
	},
	Total: "2.25",
}

var FREE_ITEM_CASE_SHORTDESC = receipt{
	Retailer:     "M&M Corner Market",
	PurchaseDate: "2022-03-20",
	PurchaseTime: "16:01",
	Items: []item{
		{
			ShortDescription: "AAA",
			Price:            "0.00",
		},
	},
	Total: "0.00",
}
var ONE_CENT_CASE_SHORTDESC = receipt{
	Retailer:     "M&M Corner Market",
	PurchaseDate: "2022-03-20",
	PurchaseTime: "16:01",
	Items: []item{
		{
			ShortDescription: "AAA",
			Price:            "0.01",
		},
	},
	Total: "0.00",
}
var SPACES_AROUND_CASE_SHORTDESC = receipt{
	Retailer:     "M&M Corner Market",
	PurchaseDate: "2022-03-20",
	PurchaseTime: "16:01",
	Items: []item{
		{
			ShortDescription: " 	 	 	 AAA	  	 	 	\n",
			Price:            "0.01",
		},
	},
	Total: "0.00",
}
var SYMBOLS_ON_NAME_CASE_RETAILER = receipt{
	Retailer:     "A1&-/+}´.,!\"$",
	PurchaseDate: "2022-03-20",
	PurchaseTime: "16:01",
	Items: []item{
		{
			ShortDescription: "ASDFASDFSAFD",
			Price:            "2.01",
		},
	},
	Total: "2.01",
}
