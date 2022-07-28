package crm

import "time"

type Lead struct {
	Id                 int             `json:"ID,string"`
	Title              string          `json:"TITLE"`
	Honorific          string          `json:"HONORIFIC"`
	Name               string          `json:"NAME"`
	SecondName         string          `json:"SECOND_NAME"`
	LastName           string          `json:"LAST_NAME"`
	Birthdate          string          `json:"BIRTHDATE"`
	CompanyTitle       string          `json:"COMPANY_TITLE"`
	SourceId           string          `json:"SOURCE_ID"`
	SourceDescription  string          `json:"SOURCE_DESCRIPTION"`
	StatusId           string          `json:"STATUS_ID"`
	StatusDescription  string          `json:"STATUS_DESCRIPTION"`
	StatusSemanticId   string          `json:"STATUS_SEMANTIC_ID"`
	Post               string          `json:"POST"`
	Address            string          `json:"ADDRESS"`
	Address2           string          `json:"ADDRESS_2"`
	AddressCity        string          `json:"ADDRESS_CITY"`
	AddressPostalCode  string          `json:"ADDRESS_POSTAL_CODE"`
	AddressRegion      string          `json:"ADDRESS_REGION"`
	AddressProvince    string          `json:"ADDRESS_PROVINCE"`
	AddressCountry     string          `json:"ADDRESS_COUNTRY"`
	AddressCountryCode string          `json:"ADDRESS_COUNTRY_CODE"`
	CurrencyId         string          `json:"CURRENCY_ID"`
	Opportunity        float64         `json:"OPPORTUNITY,string"`
	Opened             string          `json:"OPENED"`
	Comments           string          `json:"COMMENTS"`
	HasPhone           string          `json:"HAS_PHONE"`
	HasEmail           string          `json:"HAS_EMAIL"`
	HasImol            string          `json:"HAS_IMOL"`
	AssignedById       int             `json:"ASSIGNED_BY_ID,string"`
	CreatedById        int             `json:"CREATED_BY_ID,string"`
	ModifyById         int             `json:"MODIFY_BY_ID,string"`
	DateCreate         time.Time       `json:"DATE_CREATE"`
	DateModify         time.Time       `json:"DATE_MODIFY"`
	CompanyId          int             `json:"COMPANY_ID,string"`
	ContactId          int             `json:"CONTACT_ID,string"`
	IsReturnCustomer   string          `json:"IS_RETURN_CUSTOMER"`
	DateClosed         string          `json:"DATE_CLOSED"`
	OriginatorId       string          `json:"ORIGINATOR_ID"`
	OriginId           string          `json:"ORIGIN_ID"`
	UtmSource          string          `json:"UTM_SOURCE"`
	UtmMedium          string          `json:"UTM_MEDIUM"`
	UtmCampaign        string          `json:"UTM_CAMPAIGN"`
	UtmContent         string          `json:"UTM_CONTENT"`
	UtmTerm            string          `json:"UTM_TERM"`
	Phone              []CrmMultifield `json:"PHONE"`
	Email              []CrmMultifield `json:"EMAIL"`

	UfCrm1585559766    string        `json:"UF_CRM_1585559766"`
	UfCrm1585559808    string        `json:"UF_CRM_1585559808"`
	UfCrm1587971702    interface{}   `json:"UF_CRM_1587971702"`
	UfCrm1386855792    interface{}   `json:"UF_CRM_1386855792"`
	UfCrm1386857211    []interface{} `json:"UF_CRM_1386857211"`
	UfCrm1472628139    string        `json:"UF_CRM_1472628139"`
	UfUtmTerm          interface{}   `json:"UF_UTM_TERM"`
	UfUtmContent       interface{}   `json:"UF_UTM_CONTENT"`
	UfCrm1516886633    interface{}   `json:"UF_CRM_1516886633"`
	UfCrm1529520811    interface{}   `json:"UF_CRM_1529520811"`
	UfCrm1529520854    string        `json:"UF_CRM_1529520854"`
	UfCrm1534324060413 []interface{} `json:"UF_CRM_1534324060413"`
	UfCrm1557912745820 interface{}   `json:"UF_CRM_1557912745820"`
	UfCrm1557912807301 interface{}   `json:"UF_CRM_1557912807301"`
	UfCrm1543588165    interface{}   `json:"UF_CRM_1543588165"`
	UfCrm1592399114    interface{}   `json:"UF_CRM_1592399114"`
}
