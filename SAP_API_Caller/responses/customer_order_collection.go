package responses

type CustomerOrderCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                                                 string    `json:"ObjectID"`
			InformationLifeCycleStatusCode                           string    `json:"InformationLifeCycleStatusCode"`
			InformationLifeCycleStatusCodeText                       string    `json:"InformationLifeCycleStatusCodeText"`
			ClassificationCode                                       string    `json:"ClassificationCode"`
			ClassificationCodeText                                   string    `json:"ClassificationCodeText"`
			BuyerID                                                  string    `json:"BuyerID"`
			FulfilmentBlockingReasonCode                             string    `json:"FulfilmentBlockingReasonCode"`
			FulfilmentBlockingReasonCodeText                         string    `json:"FulfilmentBlockingReasonCodeText"`
			ID                                                       string    `json:"ID"`
			InvoicingBlockingReasonCode                              string    `json:"InvoicingBlockingReasonCode"`
			InvoicingBlockingReasonCodeText                          string    `json:"InvoicingBlockingReasonCodeText"`
			LastChangeDate                                           string    `json:"LastChangeDate"`
			OrderExternalLifeCycleStatusCode                         string    `json:"OrderExternalLifeCycleStatusCode"`
			OrderExternalLifeCycleStatusCodeText                     string    `json:"OrderExternalLifeCycleStatusCodeText"`
			Name                                                     string    `json:"Name"`
			LanguageCode                                             string    `json:"LanguageCode"`
			LanguageCodeText                                         string    `json:"LanguageCodeText"`
			ProcessingTypeCode                                       string    `json:"ProcessingTypeCode"`
			ProcessingTypeCodeText                                   string    `json:"ProcessingTypeCodeText"`
			ItemListCancellationStatusCode                           string    `json:"ItemListCancellationStatusCode"`
			ItemListCancellationStatusCodeText                       string    `json:"ItemListCancellationStatusCodeText"`
			ItemListFulfilmentProcessingStatusCode                   string    `json:"ItemListFulfilmentProcessingStatusCode"`
			ItemListFulfilmentProcessingStatusCodeText               string    `json:"ItemListFulfilmentProcessingStatusCodeText"`
			ApprovalStatusCode                                       string    `json:"ApprovalStatusCode"`
			ApprovalStatusCodeText                                   string    `json:"ApprovalStatusCodeText"`
			ConsistencyStatusCode                                    string    `json:"ConsistencyStatusCode"`
			ConsistencyStatusCodeText                                string    `json:"ConsistencyStatusCodeText"`
			OverallBlockingStatusCode                                string    `json:"OverallBlockingStatusCode"`
			OverallBlockingStatusCodeText                            string    `json:"OverallBlockingStatusCodeText"`
			ReplicationProcessingStatusCode                          string    `json:"ReplicationProcessingStatusCode"`
			ReplicationProcessingStatusCodeText                      string    `json:"ReplicationProcessingStatusCodeText"`
			DistributionChannelCode                                  string    `json:"DistributionChannelCode"`
			DistributionChannelCodeText                              string    `json:"DistributionChannelCodeText"`
			DivisionCode                                             string    `json:"DivisionCode"`
			DivisionCodeText                                         string    `json:"DivisionCodeText"`
			SalesGroupID                                             string    `json:"SalesGroupID"`
			SalesOfficeID                                            string    `json:"SalesOfficeID"`
			SalesOrganisationID                                      string    `json:"SalesOrganisationID"`
			SalesTerritoryID                                         string    `json:"SalesTerritoryID"`
			BuyerPartyID                                             string    `json:"BuyerPartyID"`
			DeliveryPriorityCode                                     string    `json:"DeliveryPriorityCode"`
			DeliveryPriorityCodeText                                 string    `json:"DeliveryPriorityCodeText"`
			TransferLocationName                                     string    `json:"TransferLocationName"`
			EmployeeResponsiblePartyID                               string    `json:"EmployeeResponsiblePartyID"`
			BuyerPartyName                                           string    `json:"BuyerPartyName"`
			EmployeeResponsiblePartyName                             string    `json:"EmployeeResponsiblePartyName"`
			BuyerContactPartyID                                      string    `json:"BuyerContactPartyID"`
			BuyerContactPartyName                                    string    `json:"BuyerContactPartyName"`
			CurrencyCode                                             string    `json:"CurrencyCode"`
			CurrencyCodeText                                         string    `json:"CurrencyCodeText"`
			PriceDateTime                                            string    `json:"PriceDateTime"`
			ProductRecipientPartyID                                  string    `json:"ProductRecipientPartyID"`
			ProductRecipientPartyName                                string    `json:"ProductRecipientPartyName"`
			RequestedFulfillmentStartDateTime                        string    `json:"RequestedFulfillmentStartDateTime"`
			TimeZoneCode                                             string    `json:"timeZoneCode"`
			TimeZoneCodeText                                         string    `json:"timeZoneCodeText"`
			CancellationReasonCode                                   string    `json:"CancellationReasonCode"`
			CancellationReasonCodeText                               string    `json:"CancellationReasonCodeText"`
			SalesUnitPartyID                                         string    `json:"SalesUnitPartyID"`
			CalculationStatusCode                                    string    `json:"CalculationStatusCode"`
			CalculationStatusCodeText                                string    `json:"CalculationStatusCodeText"`
			ExternalPriceDocumentBaseBusinessTransactionDocumentUUID string    `json:"ExternalPriceDocumentBaseBusinessTransactionDocumentUUID"`
			PricingProcedureCode                                     string    `json:"PricingProcedureCode"`
			PricingProcedureCodeText                                 string    `json:"PricingProcedureCodeText"`
			DateTime                                                 string    `json:"DateTime"`
			OrderReasonCode                                          string    `json:"OrderReasonCode"`
			OrderReasonCodeText                                      string    `json:"OrderReasonCodeText"`
			MaintenanceModeInternalOnlyMainDiscount                  string    `json:"MaintenanceModeInternalOnlyMainDiscount"`
			NetAmount                                                string    `json:"NetAmount"`
			NetAmountCurrencyCode                                    string    `json:"NetAmountCurrencyCode"`
			NetAmountCurrencyCodeText                                string    `json:"NetAmountCurrencyCodeText"`
			GrossAmount                                              string    `json:"GrossAmount"`
			GrossAmountCurrencyCode                                  string    `json:"GrossAmountCurrencyCode"`
			GrossAmountCurrencyCodeText                              string    `json:"GrossAmountCurrencyCodeText"`
			TaxAmount                                                string    `json:"TaxAmount"`
			TaxAmountCurrencyCode                                    string    `json:"TaxAmountCurrencyCode"`
			TaxAmountCurrencyCodeText                                string    `json:"TaxAmountCurrencyCodeText"`
			InternalPricingProcedureCode                             string    `json:"InternalPricingProcedureCode"`
			InternalPricingProcedureCodeText                         string    `json:"InternalPricingProcedureCodeText"`
			InternalPricingCalculationStatusCode                     string    `json:"InternalPricingCalculationStatusCode"`
			InternalPricingCalculationStatusCodeText                 string    `json:"InternalPricingCalculationStatusCodeText"`
			NetWeightMeasure                                         string    `json:"NetWeightMeasure"`
			NetWeightUnitCode                                        string    `json:"NetWeightUnitCode"`
			NetWeightUnitCodeText                                    string    `json:"NetWeightUnitCodeText"`
			GrossWeightMeasure                                       string    `json:"GrossWeightMeasure"`
			GrossWeightUnitCode                                      string    `json:"GrossWeightUnitCode"`
			GrossWeightUnitCodeText                                  string    `json:"GrossWeightUnitCodeText"`
			VolumeMeasure                                            string    `json:"VolumeMeasure"`
			VolumeUnitCode                                           string    `json:"VolumeUnitCode"`
			VolumeUnitCodeText                                       string    `json:"VolumeUnitCodeText"`
			Simulate                                                 bool      `json:"Simulate"`
			SubmitForApproval                                        bool      `json:"SubmitForApproval"`
			Transfer                                                 bool      `json:"Transfer"`
			WithdrawFromApproval                                     bool      `json:"WithdrawFromApproval"`
			SetAsCompleted                                           bool      `json:"SetAsCompleted"`
			PlantPartyID                                             string    `json:"PlantPartyID"`
			PlantPartyName                                           string    `json:"PlantPartyName"`
			CreatedBy                                                string    `json:"CreatedBy"`
			LastChangedBy                                            string    `json:"LastChangedBy"`
			CreationIdentityUUID                                     string    `json:"CreationIdentityUUID"`
			LastChangeIdentityUUID                                   string    `json:"LastChangeIdentityUUID"`
			EntityLastChangedOn                                      string    `json:"EntityLastChangedOn"`
			ETag                                                     string    `json:"ETag"`
			DataloaderKUT                                            string    `json:"dataloader_KUT"`
			CustomerOrderCashDiscountTerms struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"CustomerOrderCashDiscountTerms"`
			CustomerOrderItem struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"CustomerOrderItem"`
			CustomerOrderParty struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"CustomerOrderParty"`
		} `json:"results"`
	} `json:"d"`
}
