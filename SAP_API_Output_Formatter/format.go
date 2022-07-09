package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-sales-orders-reads-rmq-kube-c4/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToCustomerOrderCollection(raw []byte, l *logger.Logger) ([]CustomerOrderCollection, error) {
	pm := &responses.CustomerOrderCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to CustomerOrderCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	customerOrderCollection := make([]CustomerOrderCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		customerOrderCollection = append(customerOrderCollection, CustomerOrderCollection{
			ObjectID:                                   data.ObjectID,
			InformationLifeCycleStatusCode:             data.InformationLifeCycleStatusCode,
			InformationLifeCycleStatusCodeText:         data.InformationLifeCycleStatusCodeText,
			ClassificationCode:                         data.ClassificationCode,
			ClassificationCodeText:                     data.ClassificationCodeText,
			BuyerID:                                    data.BuyerID,
			FulfilmentBlockingReasonCode:               data.FulfilmentBlockingReasonCode,
			FulfilmentBlockingReasonCodeText:           data.FulfilmentBlockingReasonCodeText,
			ID:                                         data.ID,
			InvoicingBlockingReasonCode:                data.InvoicingBlockingReasonCode,
			InvoicingBlockingReasonCodeText:            data.InvoicingBlockingReasonCodeText,
			LastChangeDate:                             data.LastChangeDate,
			OrderExternalLifeCycleStatusCode:           data.OrderExternalLifeCycleStatusCode,
			OrderExternalLifeCycleStatusCodeText:       data.OrderExternalLifeCycleStatusCodeText,
			Name:                                       data.Name,
			LanguageCode:                               data.LanguageCode,
			LanguageCodeText:                           data.LanguageCodeText,
			ProcessingTypeCode:                         data.ProcessingTypeCode,
			ProcessingTypeCodeText:                     data.ProcessingTypeCodeText,
			ItemListCancellationStatusCode:             data.ItemListCancellationStatusCode,
			ItemListCancellationStatusCodeText:         data.ItemListCancellationStatusCodeText,
			ItemListFulfilmentProcessingStatusCode:     data.ItemListFulfilmentProcessingStatusCode,
			ItemListFulfilmentProcessingStatusCodeText: data.ItemListFulfilmentProcessingStatusCodeText,
			ApprovalStatusCode:                         data.ApprovalStatusCode,
			ApprovalStatusCodeText:                     data.ApprovalStatusCodeText,
			ConsistencyStatusCode:                      data.ConsistencyStatusCode,
			ConsistencyStatusCodeText:                  data.ConsistencyStatusCodeText,
			OverallBlockingStatusCode:                  data.OverallBlockingStatusCode,
			OverallBlockingStatusCodeText:              data.OverallBlockingStatusCodeText,
			ReplicationProcessingStatusCode:            data.ReplicationProcessingStatusCode,
			ReplicationProcessingStatusCodeText:        data.ReplicationProcessingStatusCodeText,
			DistributionChannelCode:                    data.DistributionChannelCode,
			DistributionChannelCodeText:                data.DistributionChannelCodeText,
			DivisionCode:                               data.DivisionCode,
			DivisionCodeText:                           data.DivisionCodeText,
			SalesGroupID:                               data.SalesGroupID,
			SalesOfficeID:                              data.SalesOfficeID,
			SalesOrganisationID:                        data.SalesOrganisationID,
			SalesTerritoryID:                           data.SalesTerritoryID,
			BuyerPartyID:                               data.BuyerPartyID,
			DeliveryPriorityCode:                       data.DeliveryPriorityCode,
			DeliveryPriorityCodeText:                   data.DeliveryPriorityCodeText,
			TransferLocationName:                       data.TransferLocationName,
			EmployeeResponsiblePartyID:                 data.EmployeeResponsiblePartyID,
			BuyerPartyName:                             data.BuyerPartyName,
			EmployeeResponsiblePartyName:               data.EmployeeResponsiblePartyName,
			BuyerContactPartyID:                        data.BuyerContactPartyID,
			BuyerContactPartyName:                      data.BuyerContactPartyName,
			CurrencyCode:                               data.CurrencyCode,
			CurrencyCodeText:                           data.CurrencyCodeText,
			PriceDateTime:                              data.PriceDateTime,
			ProductRecipientPartyID:                    data.ProductRecipientPartyID,
			ProductRecipientPartyName:                  data.ProductRecipientPartyName,
			RequestedFulfillmentStartDateTime:          data.RequestedFulfillmentStartDateTime,
			TimeZoneCode:                               data.TimeZoneCode,
			TimeZoneCodeText:                           data.TimeZoneCodeText,
			CancellationReasonCode:                     data.CancellationReasonCode,
			CancellationReasonCodeText:                 data.CancellationReasonCodeText,
			SalesUnitPartyID:                           data.SalesUnitPartyID,
			CalculationStatusCode:                      data.CalculationStatusCode,
			CalculationStatusCodeText:                  data.CalculationStatusCodeText,
			ExternalPriceDocumentBaseBusinessTransactionDocumentUUID: data.ExternalPriceDocumentBaseBusinessTransactionDocumentUUID,
			PricingProcedureCode:                     data.PricingProcedureCode,
			PricingProcedureCodeText:                 data.PricingProcedureCodeText,
			DateTime:                                 data.DateTime,
			OrderReasonCode:                          data.OrderReasonCode,
			OrderReasonCodeText:                      data.OrderReasonCodeText,
			MaintenanceModeInternalOnlyMainDiscount:  data.MaintenanceModeInternalOnlyMainDiscount,
			NetAmount:                                data.NetAmount,
			NetAmountCurrencyCode:                    data.NetAmountCurrencyCode,
			NetAmountCurrencyCodeText:                data.NetAmountCurrencyCodeText,
			GrossAmount:                              data.GrossAmount,
			GrossAmountCurrencyCode:                  data.GrossAmountCurrencyCode,
			GrossAmountCurrencyCodeText:              data.GrossAmountCurrencyCodeText,
			TaxAmount:                                data.TaxAmount,
			TaxAmountCurrencyCode:                    data.TaxAmountCurrencyCode,
			TaxAmountCurrencyCodeText:                data.TaxAmountCurrencyCodeText,
			InternalPricingProcedureCode:             data.InternalPricingProcedureCode,
			InternalPricingProcedureCodeText:         data.InternalPricingProcedureCodeText,
			InternalPricingCalculationStatusCode:     data.InternalPricingCalculationStatusCode,
			InternalPricingCalculationStatusCodeText: data.InternalPricingCalculationStatusCodeText,
			NetWeightMeasure:                         data.NetWeightMeasure,
			NetWeightUnitCode:                        data.NetWeightUnitCode,
			NetWeightUnitCodeText:                    data.NetWeightUnitCodeText,
			GrossWeightMeasure:                       data.GrossWeightMeasure,
			GrossWeightUnitCode:                      data.GrossWeightUnitCode,
			GrossWeightUnitCodeText:                  data.GrossWeightUnitCodeText,
			VolumeMeasure:                            data.VolumeMeasure,
			VolumeUnitCode:                           data.VolumeUnitCode,
			VolumeUnitCodeText:                       data.VolumeUnitCodeText,
			Simulate:                                 data.Simulate,
			SubmitForApproval:                        data.SubmitForApproval,
			Transfer:                                 data.Transfer,
			WithdrawFromApproval:                     data.WithdrawFromApproval,
			SetAsCompleted:                           data.SetAsCompleted,
			PlantPartyID:                             data.PlantPartyID,
			PlantPartyName:                           data.PlantPartyName,
			CreatedBy:                                data.CreatedBy,
			LastChangedBy:                            data.LastChangedBy,
			CreationIdentityUUID:                     data.CreationIdentityUUID,
			LastChangeIdentityUUID:                   data.LastChangeIdentityUUID,
			EntityLastChangedOn:                      data.EntityLastChangedOn,
			ETag:                                     data.ETag,
			DataloaderKUT:                            data.DataloaderKUT,
			ToCustomerOrderCashDiscountTerms:         data.CustomerOrderCashDiscountTerms.Deferred.URI,
			ToCustomerOrderItem:                      data.CustomerOrderItem.Deferred.URI,
			ToCustomerOrderParty:                     data.CustomerOrderParty.Deferred.URI,
		})
	}

	return customerOrderCollection, nil
}

func ConvertToCustomerOrderCashDiscountTerms(raw []byte, l *logger.Logger) (*CustomerOrderCashDiscountTerms, error) {
	pm := &responses.CustomerOrderCashDiscountTerms{}

	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to CustomerOrderCashDiscountTerms. unmarshal error: %w", err)
	}

	return &CustomerOrderCashDiscountTerms{
		ObjectID:       pm.D.Results.ObjectID,
		ParentObjectID: pm.D.Results.ParentObjectID,
		Code:           pm.D.Results.Code,
		CodeText:       pm.D.Results.CodeText,
		SalesOrderID:   pm.D.Results.SalesOrderID,
		ETag:           pm.D.Results.ETag,
	}, nil
}

func ConvertToCustomerOrderItem(raw []byte, l *logger.Logger) ([]CustomerOrderItem, error) {
	pm := &responses.CustomerOrderItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to CustomerOrderItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	customerOrderItem := make([]CustomerOrderItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		customerOrderItem = append(customerOrderItem, CustomerOrderItem{
			ObjectID:                             data.ObjectID,
			ParentObjectID:                       data.ParentObjectID,
			Description:                          data.Description,
			LanguageCode:                         data.LanguageCode,
			LanguageCodeText:                     data.LanguageCodeText,
			CreationDate:                         data.CreationDate,
			ParentItemID:                         data.ParentItemID,
			ID:                                   data.ID,
			ProcessingTypeCode:                   data.ProcessingTypeCode,
			ProcessingTypeCodeText:               data.ProcessingTypeCodeText,
			ProductID:                            data.ProductID,
			ProductStandardID:                    data.ProductStandardID,
			ProductInternalID:                    data.ProductInternalID,
			Quantity:                             data.Quantity,
			QuantityMeasureUnitCode:              data.QuantityMeasureUnitCode,
			QuantityMeasureUnitCodeText:          data.QuantityMeasureUnitCodeText,
			StartDateTime:                        data.StartDateTime,
			TimeZoneCode:                         data.TimeZoneCode,
			TimeZoneCodeText:                     data.TimeZoneCodeText,
			SalesTermsCancellationReasonCode:     data.SalesTermsCancellationReasonCode,
			SalesTermsCancellationReasonCodeText: data.SalesTermsCancellationReasonCodeText,
			MaintenanceModeInternalOnlyItemMainDiscount:   data.MaintenanceModeInternalOnlyItemMainDiscount,
			ItemProductRecipientBusinessPartnerInternalID: data.ItemProductRecipientBusinessPartnerInternalID,
			BuyerID:                            data.BuyerID,
			ProductBuyerID:                     data.ProductBuyerID,
			SalesOrderID:                       data.SalesOrderID,
			NetAmount:                          data.NetAmount,
			NetAmountCurrencyCode:              data.NetAmountCurrencyCode,
			NetAmountCurrencyCodeText:          data.NetAmountCurrencyCodeText,
			NetPriceAmount:                     data.NetPriceAmount,
			NetPriceAmountCurrencyCode:         data.NetPriceAmountCurrencyCode,
			NetPriceAmountCurrencyCodeText:     data.NetPriceAmountCurrencyCodeText,
			TaxAmount:                          data.TaxAmount,
			TaxAmountCurrencyCode:              data.TaxAmountCurrencyCode,
			TaxAmountCurrencyCodeText:          data.TaxAmountCurrencyCodeText,
			NetWeightMeasure:                   data.NetWeightMeasure,
			NetWeightUnitCode:                  data.NetWeightUnitCode,
			NetWeightUnitCodeText:              data.NetWeightUnitCodeText,
			GrossWeightMeasure:                 data.GrossWeightMeasure,
			GrossWeightUnitCode:                data.GrossWeightUnitCode,
			GrossWeightUnitCodeText:            data.GrossWeightUnitCodeText,
			VolumeMeasure:                      data.VolumeMeasure,
			VolumeUnitCode:                     data.VolumeUnitCode,
			VolumeUnitCodeText:                 data.VolumeUnitCodeText,
			NetPriceBaseQuantity:               data.NetPriceBaseQuantity,
			NetPriceBaseQuantityunitCode:       data.NetPriceBaseQuantityunitCode,
			NetPriceBaseQuantityunitCodeText:   data.NetPriceBaseQuantityunitCodeText,
			PlantPartyID:                       data.PlantPartyID,
			PlantName:                          data.PlantName,
			LifeCycleStatusCode:                data.LifeCycleStatusCode,
			LifeCycleStatusCodeText:            data.LifeCycleStatusCodeText,
			CancellationStatusCode:             data.CancellationStatusCode,
			CancellationStatusCodeText:         data.CancellationStatusCodeText,
			FulfilmentProcessingStatusCode:     data.FulfilmentProcessingStatusCode,
			FulfilmentProcessingStatusCodeText: data.FulfilmentProcessingStatusCodeText,
			InvoiceProcessingStatusCode:        data.InvoiceProcessingStatusCode,
			InvoiceProcessingStatusCodeText:    data.InvoiceProcessingStatusCodeText,
			EntityLastChangedOn:                data.EntityLastChangedOn,
			ETag:                               data.ETag,
		})
	}

	return customerOrderItem, nil
}

func ConvertToCustomerOrderParty(raw []byte, l *logger.Logger) ([]CustomerOrderParty, error) {
	pm := &responses.CustomerOrderParty{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to CustomerOrderParty. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	customerOrderParty := make([]CustomerOrderParty, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		customerOrderParty = append(customerOrderParty, CustomerOrderParty{
			ObjectID:                                 data.ObjectID,
			ParentObjectID:                           data.ParentObjectID,
			TimeZoneCode:                             data.TimeZoneCode,
			TimeZoneCodeText:                         data.TimeZoneCodeText,
			PartyName:                                data.PartyName,
			PartyID:                                  data.PartyID,
			RoleCode:                                 data.RoleCode,
			RoleCodeText:                             data.RoleCodeText,
			SalesOrderID:                             data.SalesOrderID,
			PreferredCommunicationMediumTypeCode:     data.PreferredCommunicationMediumTypeCode,
			PreferredCommunicationMediumTypeCodeText: data.PreferredCommunicationMediumTypeCodeText,
			MainIndicator:                            data.MainIndicator,
			ConventionalPhone:                        data.ConventionalPhone,
			Email:                                    data.Email,
			Web:                                      data.Web,
			Fax:                                      data.Fax,
			MobilePhone:                              data.MobilePhone,
			FirstLineName:                            data.FirstLineName,
			SecondLineName:                           data.SecondLineName,
			ThirdLineName:                            data.ThirdLineName,
			FourthLineName:                           data.FourthLineName,
			CareOfName:                               data.CareOfName,
			CityName:                                 data.CityName,
			CountryCode:                              data.CountryCode,
			CountryCodeText:                          data.CountryCodeText,
			CountyName:                               data.CountyName,
			DistrictName:                             data.DistrictName,
			StreetPrefixName:                         data.StreetPrefixName,
			AdditionalStreetPrefixName:               data.AdditionalStreetPrefixName,
			HouseID:                                  data.HouseID,
			StreetName:                               data.StreetName,
			StreetSuffixName:                         data.StreetSuffixName,
			AdditionalStreetSuffixName:               data.AdditionalStreetSuffixName,
			RegionCode:                               data.RegionCode,
			RegionCodeText:                           data.RegionCodeText,
			StreetPostalCode:                         data.StreetPostalCode,
			POBoxIndicator:                           data.POBoxIndicator,
			POBoxID:                                  data.POBoxID,
			POBoxPostalCode:                          data.POBoxPostalCode,
			POBoxDeviatingCityName:                   data.POBoxDeviatingCityName,
			POBoxDeviatingCountryCode:                data.POBoxDeviatingCountryCode,
			POBoxDeviatingCountryCodeText:            data.POBoxDeviatingCountryCodeText,
			POBoxDeviatingRegionCode:                 data.POBoxDeviatingRegionCode,
			POBoxDeviatingRegionCodeText:             data.POBoxDeviatingRegionCodeText,
			ETag:                                     data.ETag,
		})
	}

	return customerOrderParty, nil
}

