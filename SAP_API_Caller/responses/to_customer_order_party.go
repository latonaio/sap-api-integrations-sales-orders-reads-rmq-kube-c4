package responses

type CustomerOrderParty struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                                 string `json:"ObjectID"`
			ParentObjectID                           string `json:"ParentObjectID"`
			TimeZoneCode                             string `json:"TimeZoneCode"`
			TimeZoneCodeText                         string `json:"TimeZoneCodeText"`
			PartyName                                string `json:"PartyName"`
			PartyID                                  string `json:"PartyID"`
			RoleCode                                 string `json:"RoleCode"`
			RoleCodeText                             string `json:"RoleCodeText"`
			SalesOrderID                             string `json:"SalesOrderID"`
			PreferredCommunicationMediumTypeCode     string `json:"PreferredCommunicationMediumTypeCode"`
			PreferredCommunicationMediumTypeCodeText string `json:"PreferredCommunicationMediumTypeCodeText"`
			MainIndicator                            bool   `json:"MainIndicator"`
			ConventionalPhone                        string `json:"ConventionalPhone"`
			Email                                    string `json:"Email"`
			Web                                      string `json:"Web"`
			Fax                                      string `json:"Fax"`
			MobilePhone                              string `json:"MobilePhone"`
			FirstLineName                            string `json:"FirstLineName"`
			SecondLineName                           string `json:"SecondLineName"`
			ThirdLineName                            string `json:"ThirdLineName"`
			FourthLineName                           string `json:"FourthLineName"`
			CareOfName                               string `json:"CareOfName"`
			CityName                                 string `json:"CityName"`
			CountryCode                              string `json:"CountryCode"`
			CountryCodeText                          string `json:"CountryCodeText"`
			CountyName                               string `json:"CountyName"`
			DistrictName                             string `json:"DistrictName"`
			StreetPrefixName                         string `json:"StreetPrefixName"`
			AdditionalStreetPrefixName               string `json:"AdditionalStreetPrefixName"`
			HouseID                                  string `json:"HouseID"`
			StreetName                               string `json:"StreetName"`
			StreetSuffixName                         string `json:"StreetSuffixName"`
			AdditionalStreetSuffixName               string `json:"AdditionalStreetSuffixName"`
			RegionCode                               string `json:"RegionCode"`
			RegionCodeText                           string `json:"RegionCodeText"`
			StreetPostalCode                         string `json:"StreetPostalCode"`
			POBoxIndicator                           bool   `json:"POBoxIndicator"`
			POBoxID                                  string `json:"POBoxID"`
			POBoxPostalCode                          string `json:"POBoxPostalCode"`
			POBoxDeviatingCityName                   string `json:"POBoxDeviatingCityName"`
			POBoxDeviatingCountryCode                string `json:"POBoxDeviatingCountryCode"`
			POBoxDeviatingCountryCodeText            string `json:"POBoxDeviatingCountryCodeText"`
			POBoxDeviatingRegionCode                 string `json:"POBoxDeviatingRegionCode"`
			POBoxDeviatingRegionCodeText             string `json:"POBoxDeviatingRegionCodeText"`
			ETag                                     string `json:"ETag"`
		} `json:"results"`
	} `json:"d"`
}
