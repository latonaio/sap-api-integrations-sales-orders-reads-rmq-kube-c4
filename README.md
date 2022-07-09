# sap-api-integrations-sales-orders-reads-rmq-kube-c4 
sap-api-integrations-sales-orders-reads-rmq-kube-c4は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API キャンペーンデータを取得するマイクロサービスです。  
sap-api-integrations-sales-orders-reads-rmq-kube-c4には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-sales-orders-reads-rmq-kube-c4は、オンプレミス版である（＝クラウド版ではない）SAPC4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/salesorder/overview    

## 動作環境
sap-api-integrations-sales-orders-reads-rmq-kube-c4は、主にエッジコンピューティング環境における動作にフォーカスしています。   
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。   
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）  
・ RabbitMQ on Kubernetes  
・ RabbitMQ Client   

## クラウド環境での利用  
sap-api-integrations-sales-orders-reads-rmq-kube-c4は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## RabbitMQ からの JSON Input

sap-api-integrations-sales-orders-reads-rmq-kube-c4  は、Inputとして、RabbitMQ からのメッセージをJSON形式で受け取ります。 
Input の サンプルJSON は、Inputs フォルダ内にあります。  

## RabbitMQ からのメッセージ受信による イベントドリヴン の ランタイム実行

sap-api-integrations-sales-orders-reads-rmq-kube-c4  は、RabbitMQ からのメッセージを受け取ると、イベントドリヴンでランタイムを実行します。  
AION の仕様では、Kubernetes 上 の 当該マイクロサービスPod は 立ち上がったまま待機状態で当該メッセージを受け取り、（コンテナ起動などの段取時間をカットして）即座にランタイムを実行します。　

## RabbitMQ への JSON Output

sap-api-integrations-sales-orders-reads-rmq-kube-c4 は、Outputとして、RabbitMQ へのメッセージをJSON形式で出力します。

## RabbitMQ の マスタサーバ環境

sap-api-integrations-sales-orders-reads-rmq-kube-c4 が利用する RabbitMQ のマスタサーバ環境は、[rabbitmq-on-kubernetes](https://github.com/latonaio/rabbitmq-on-kubernetes) です。  
当該マスタサーバ環境は、同じエッジコンピューティングデバイスに配置されても、別の物理(仮想)サーバ内に配置されても、どちらでも構いません。

## RabbitMQ の Golang Runtime ライブラリ
sap-api-integrations-sales-orders-reads-rmq-kube-c4 は、RabbitMQ の Golang Runtime ライブラリ として、[rabbitmq-golang-client](https://github.com/latonaio/rabbitmq-golang-client)を利用しています。

## デプロイ・稼働
sap-api-integrations-sales-orders-reads-rmq-kube-c4 の デプロイ・稼働 を行うためには、aion-service-definitions の services.yml に、本レポジトリの services.yml を設定する必要があります。

kubectl apply - f 等で Deployment作成後、以下のコマンドで Pod が正しく生成されていることを確認してください。
```
$ kubectl get pods
```

## 本レポジトリ が 対応する API サービス
sap-api-integrations-sales-orders-reads-rmq-kube-c4が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/salesorder/overview     
* APIサービス名(=baseURL): c4codataapi

## 本レポジトリ に 含まれる API名
sap-api-integrations-sales-orders-reads-rmq-kube-c4には、次の API をコールするためのリソースが含まれています。  

* CustomerOrderCollection（受注 - 顧客受注）※顧客受注の詳細データを取得するために、ToCustomerOrderCashDiscountTerms、ToCustomerOrderParty、ToCustomerOrderItem、と合わせて利用されます。
* ToCustomerOrderCashDiscountTerms（受注 - 顧客受注現金割引条件 ※To）
* ToCustomerOrderParty（受注 - 顧客受注情報 ※To）
* ToCustomerOrderItem（受注 - 顧客受注明細 ※To）

## API への 値入力条件 の 初期値
sap-api-integrations-sales-orders-reads-rmq-kube-c4において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.CustomerOrderCollection.ID（ID）  


## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"CustomerOrderCollection" が指定されています。    
  
```
    "api_schema": "SalesOrderCollection",
	"accepter": ["CustomerOrderCollection"],
	"customerorder_Collection_code": "9000000152",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "SalesOrderCollection",
	"accepter": ["All"],
	"customerorder_Collection_code": "9000000152",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
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
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 受注  の 顧客受注データ が取得された結果の JSON の例です。  
以下の項目のうち、"ObjectID" ～ "CustomerOrderParty" は、/SAP_API_Output_Formatter/type.go 内 の Type CustomerOrderCollection {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-sales-orders-reads-rmq-kube-c4/SAP_API_Caller/caller.go#L54",
	"function": "sap-api-integrations-sales-orders-reads-rmq-kube-c4/SAP_API_Caller.(*SAPAPICaller).CustomerOrderCollection",
	"level": "INFO",
	"message": [
		{
			"ObjectID": "00163E0C6CDB1ED694C4D37B9E324AAE",
			"InformationLifeCycleStatusCode": "AC",
			"InformationLifeCycleStatusCodeText": "Active",
			"ClassificationCode": "",
			"ClassificationCodeText": "",
			"BuyerID": "",
			"FulfilmentBlockingReasonCode": "",
			"FulfilmentBlockingReasonCodeText": "",
			"ID": "9000000152",
			"InvoicingBlockingReasonCode": "",
			"InvoicingBlockingReasonCodeText": "",
			"LastChangeDate": "2016-07-25T09:00:00+09:00",
			"OrderExternalLifeCycleStatusCode": "1",
			"OrderExternalLifeCycleStatusCodeText": "Open",
			"Name": "",
			"LanguageCode": "EN",
			"LanguageCodeText": "English",
			"ProcessingTypeCode": "OR",
			"ProcessingTypeCodeText": "Standard Order",
			"ItemListCancellationStatusCode": "1",
			"ItemListCancellationStatusCodeText": "Not Canceled",
			"ItemListFulfilmentProcessingStatusCode": "4",
			"ItemListFulfilmentProcessingStatusCodeText": "Not Relevant",
			"ApprovalStatusCode": "",
			"ApprovalStatusCodeText": "",
			"ConsistencyStatusCode": "2",
			"ConsistencyStatusCodeText": "Inconsistent",
			"OverallBlockingStatusCode": "1",
			"OverallBlockingStatusCodeText": "Not Blocked",
			"ReplicationProcessingStatusCode": "1",
			"ReplicationProcessingStatusCodeText": "Not Started",
			"DistributionChannelCode": "01",
			"DistributionChannelCodeText": "Direct sales",
			"DivisionCode": "01",
			"DivisionCodeText": "Division 1",
			"SalesGroupID": "",
			"SalesOfficeID": "",
			"SalesOrganisationID": "US1100",
			"SalesTerritoryID": "7",
			"BuyerPartyID": "10332",
			"DeliveryPriorityCode": "3",
			"DeliveryPriorityCodeText": "Normal",
			"TransferLocationName": "",
			"EmployeeResponsiblePartyID": "8000000001",
			"BuyerPartyName": "Alpha Solutions",
			"EmployeeResponsiblePartyName": "Mike Summers",
			"BuyerContactPartyID": "1000394",
			"BuyerContactPartyName": "Catherine Lesacher",
			"CurrencyCode": "USD",
			"CurrencyCodeText": "US Dollar",
			"PriceDateTime": "2016-07-25T14:14:06+09:00",
			"ProductRecipientPartyID": "10332",
			"ProductRecipientPartyName": "Alpha Solutions",
			"RequestedFulfillmentStartDateTime": "2016-07-27T00:00:00Z",
			"timeZoneCode": "UTC",
			"timeZoneCodeText": "UTC+0",
			"CancellationReasonCode": "",
			"CancellationReasonCodeText": "",
			"SalesUnitPartyID": "US1100",
			"CalculationStatusCode": "1",
			"CalculationStatusCodeText": "Not Calculated",
			"ExternalPriceDocumentBaseBusinessTransactionDocumentUUID": "00163E0C-6CDB-1ED6-94C4-D37B9E324AAE",
			"PricingProcedureCode": "",
			"PricingProcedureCodeText": "",
			"DateTime": "2016-07-25T14:14:06+09:00",
			"OrderReasonCode": "",
			"OrderReasonCodeText": "",
			"MaintenanceModeInternalOnlyMainDiscount": "0.00000000000000",
			"NetAmount": "0.000000",
			"NetAmountCurrencyCode": "",
			"NetAmountCurrencyCodeText": "",
			"GrossAmount": "0.000000",
			"GrossAmountCurrencyCode": "",
			"GrossAmountCurrencyCodeText": "",
			"TaxAmount": "0.000000",
			"TaxAmountCurrencyCode": "",
			"TaxAmountCurrencyCodeText": "",
			"InternalPricingProcedureCode": "",
			"InternalPricingProcedureCodeText": "",
			"InternalPricingCalculationStatusCode": "",
			"InternalPricingCalculationStatusCodeText": "",
			"NetWeightMeasure": "0.00000000000000",
			"NetWeightUnitCode": "",
			"NetWeightUnitCodeText": "",
			"GrossWeightMeasure": "0.00000000000000",
			"GrossWeightUnitCode": "",
			"GrossWeightUnitCodeText": "",
			"VolumeMeasure": "0.00000000000000",
			"VolumeUnitCode": "",
			"VolumeUnitCodeText": "",
			"Simulate": false,
			"SubmitForApproval": false,
			"Transfer": false,
			"WithdrawFromApproval": false,
			"SetAsCompleted": false,
			"PlantPartyID": "",
			"PlantPartyName": "",
			"CreatedBy": "Mike Summers",
			"LastChangedBy": "Mike Summers",
			"CreationIdentityUUID": "00163E03-A070-1EE2-88BA-37FB8F7790A9",
			"LastChangeIdentityUUID": "00163E03-A070-1EE2-88BA-37FB8F7790A9",
			"EntityLastChangedOn": "2016-07-25T18:48:44+09:00",
			"ETag": "2019-11-12T21:30:56+09:00",
			"dataloader_KUT": "",
			"CustomerOrderCashDiscountTerms": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/CustomerOrderCollection('00163E0C6CDB1ED694C4D37B9E324AAE')/CustomerOrderCashDiscountTerms",
			"CustomerOrderItem": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/CustomerOrderCollection('00163E0C6CDB1ED694C4D37B9E324AAE')/CustomerOrderItem",
			"CustomerOrderParty": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/CustomerOrderCollection('00163E0C6CDB1ED694C4D37B9E324AAE')/CustomerOrderParty"
		}
	],
	"time": "2022-06-01T22:33:49+09:00"
}

```