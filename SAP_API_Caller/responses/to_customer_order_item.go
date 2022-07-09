package responses

type CustomerOrderItem struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                                      string `json:"ObjectID"`
			ParentObjectID                                string `json:"ParentObjectID"`
			Description                                   string `json:"Description"`
			LanguageCode                                  string `json:"languageCode"`
			LanguageCodeText                              string `json:"languageCodeText"`
			CreationDate                                  string `json:"CreationDate"`
			ParentItemID                                  string `json:"ParentItemID"`
			ID                                            string `json:"ID"`
			ProcessingTypeCode                            string `json:"ProcessingTypeCode"`
			ProcessingTypeCodeText                        string `json:"ProcessingTypeCodeText"`
			ProductID                                     string `json:"ProductID"`
			ProductStandardID                             string `json:"ProductStandardID"`
			ProductInternalID                             string `json:"ProductInternalID"`
			Quantity                                      string `json:"Quantity"`
			QuantityMeasureUnitCode                       string `json:"QuantityMeasureUnitCode"`
			QuantityMeasureUnitCodeText                   string `json:"QuantityMeasureUnitCodeText"`
			StartDateTime                                 string `json:"StartDateTime"`
			TimeZoneCode                                  string `json:"TimeZoneCode"`
			TimeZoneCodeText                              string `json:"TimeZoneCodeText"`
			SalesTermsCancellationReasonCode              string `json:"SalesTermsCancellationReasonCode"`
			SalesTermsCancellationReasonCodeText          string `json:"SalesTermsCancellationReasonCodeText"`
			MaintenanceModeInternalOnlyItemMainDiscount   string `json:"MaintenanceModeInternalOnlyItemMainDiscount"`
			ItemProductRecipientBusinessPartnerInternalID string `json:"ItemProductRecipientBusinessPartnerInternalID"`
			BuyerID                                       string `json:"BuyerID"`
			ProductBuyerID                                string `json:"ProductBuyerID"`
			SalesOrderID                                  string `json:"SalesOrderID"`
			NetAmount                                     string `json:"NetAmount"`
			NetAmountCurrencyCode                         string `json:"NetAmountCurrencyCode"`
			NetAmountCurrencyCodeText                     string `json:"NetAmountCurrencyCodeText"`
			NetPriceAmount                                string `json:"NetPriceAmount"`
			NetPriceAmountCurrencyCode                    string `json:"NetPriceAmountCurrencyCode"`
			NetPriceAmountCurrencyCodeText                string `json:"NetPriceAmountCurrencyCodeText"`
			TaxAmount                                     string `json:"TaxAmount"`
			TaxAmountCurrencyCode                         string `json:"TaxAmountCurrencyCode"`
			TaxAmountCurrencyCodeText                     string `json:"TaxAmountCurrencyCodeText"`
			NetWeightMeasure                              string `json:"NetWeightMeasure"`
			NetWeightUnitCode                             string `json:"NetWeightUnitCode"`
			NetWeightUnitCodeText                         string `json:"NetWeightUnitCodeText"`
			GrossWeightMeasure                            string `json:"GrossWeightMeasure"`
			GrossWeightUnitCode                           string `json:"GrossWeightUnitCode"`
			GrossWeightUnitCodeText                       string `json:"GrossWeightUnitCodeText"`
			VolumeMeasure                                 string `json:"VolumeMeasure"`
			VolumeUnitCode                                string `json:"VolumeUnitCode"`
			VolumeUnitCodeText                            string `json:"VolumeUnitCodeText"`
			NetPriceBaseQuantity                          string `json:"NetPriceBaseQuantity"`
			NetPriceBaseQuantityunitCode                  string `json:"NetPriceBaseQuantityunitCode"`
			NetPriceBaseQuantityunitCodeText              string `json:"NetPriceBaseQuantityunitCodeText"`
			PlantPartyID                                  string `json:"PlantPartyID"`
			PlantName                                     string `json:"PlantName"`
			LifeCycleStatusCode                           string `json:"LifeCycleStatusCode"`
			LifeCycleStatusCodeText                       string `json:"LifeCycleStatusCodeText"`
			CancellationStatusCode                        string `json:"CancellationStatusCode"`
			CancellationStatusCodeText                    string `json:"CancellationStatusCodeText"`
			FulfilmentProcessingStatusCode                string `json:"FulfilmentProcessingStatusCode"`
			FulfilmentProcessingStatusCodeText            string `json:"FulfilmentProcessingStatusCodeText"`
			InvoiceProcessingStatusCode                   string `json:"InvoiceProcessingStatusCode"`
			InvoiceProcessingStatusCodeText               string `json:"InvoiceProcessingStatusCodeText"`
			EntityLastChangedOn                           string `json:"EntityLastChangedOn"`
			ETag                                          string `json:"ETag"`
		} `json:"results"`
	} `json:"d"`
}