package main

import (
	"fmt"
	"net/http"
	"github.com/google/uuid"
	"time"
	"log"
	"io/ioutil"
	"encoding/json"
)

type PriceReturn struct {
	BillingCurrency string
	CustomerEntityId string
	CustomerEntityType string
	Items []PriceRecord
	NextPageLink string
	Count int
}

type PriceRecord struct {
	CurrencyCode string
	TierMinimumUnits float32
	RetailPrice float32
	UnitPrice float32
	ArmRegionName string
	Location string
	EffectiveStartDate time.Time
	MeterId uuid.UUID
	MeterName string
	ProductId string
	SkuId string
	ProductName string
	SkuName string
	ServiceName string
	ServiceId string
	ServiceFamily string
	UnitOfMeasure string
	Type string
	IsPrimaryMeterRegion bool
	ArmSkuName string
}

func main() {
	var url = "https://prices.azure.com/api/retail/prices"

	fmt.Println("Connecting to Azure Retail Price API")

	var req *http.Request
	var resp *http.Response
	var err error

	for len(url) != 0 {
		req, err = http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}
	
		httpClient := http.DefaultClient
	
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
	
		var body []byte
	
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	
		var prices PriceReturn
		err = json.Unmarshal(body, &prices)
		if err != nil {
			log.Fatal(err)
		}
	
		// for _, price := range prices.Items {
			// do what you want here
		// }
	
		url = prices.NextPageLink
		fmt.Println(url)
	}

	fmt.Println("Done")
}