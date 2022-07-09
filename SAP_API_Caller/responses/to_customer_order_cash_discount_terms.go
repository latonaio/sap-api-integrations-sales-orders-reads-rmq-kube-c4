package responses

type CustomerOrderCashDiscountTerms struct {
	D struct {
		Results struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID       string `json:"ObjectID"`
			ParentObjectID string `json:"ParentObjectID"`
			Code           string `json:"Code"`
			CodeText       string `json:"CodeText"`
			SalesOrderID   string `json:"SalesOrderID"`
			ETag           string `json:"ETag"`
		} `json:"results"`
	} `json:"d"`
}