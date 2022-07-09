package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-sales-orders-reads-rmq-kube-c4/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type RMQOutputter interface {
	Send(sendQueue string, payload map[string]interface{}) error
}

type SAPAPICaller struct {
	baseURL      string
	apiKey       string
	outputQueues []string
	outputter    RMQOutputter
	log          *logger.Logger
}

func NewSAPAPICaller(baseUrl string, outputQueueTo []string, outputter RMQOutputter, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:      baseUrl,
		apiKey:       GetApiKey(),
		outputQueues: outputQueueTo,
		outputter:    outputter,
		log:          l,
	}
}

func (c *SAPAPICaller) AsyncGetSalesOrders(iD string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "CustomerOrderCollection":
			func() {
				c.CustomerOrderCollection(iD)
				wg.Done()

			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) CustomerOrderCollection(iD string) {
	customerOrderCollectionData, err := c.callSalesOrdersSrvAPIRequirementCustomerOrderCollection("CustomerOrderCollectionData", iD)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": customerOrderCollectionData, "function": "CustomerOrderCollectionData"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(customerOrderCollectionData)

	customerOrderCashDiscountTermsData, err := c.callCustomerOrderCashDiscountTerms(customerOrderCollectionData[0].ToCustomerOrderCashDiscountTerms)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": customerOrderCashDiscountTermsData, "function": "CustomerOrderCashDiscountTermsData"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(customerOrderCashDiscountTermsData)

	customerOrderItemData, err := c.callCustomerOrderItem(customerOrderCollectionData[0].ToCustomerOrderItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": customerOrderItemData, "function": "CustomerOrderItemData"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(customerOrderItemData)

	customerOrderPartyData, err := c.callCustomerOrderParty(customerOrderCollectionData[0].ToCustomerOrderParty)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": customerOrderPartyData, "function": "CustomerOrderPartyData"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(customerOrderPartyData)

}

func (c *SAPAPICaller) callSalesOrdersSrvAPIRequirementCustomerOrderCollection(api, iD string) ([]sap_api_output_formatter.CustomerOrderCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithCustomerOrderCollection(req, iD)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCustomerOrderCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callCustomerOrderCashDiscountTerms(url string) (*sap_api_output_formatter.CustomerOrderCashDiscountTerms, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCustomerOrderCashDiscountTerms(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callCustomerOrderItem(url string) ([]sap_api_output_formatter.CustomerOrderItem, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCustomerOrderItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callCustomerOrderParty(url string) ([]sap_api_output_formatter.CustomerOrderParty, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCustomerOrderParty(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithCustomerOrderCollection(req *http.Request, iD string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ID eq '%s'", iD))
	req.URL.RawQuery = params.Encode()
}