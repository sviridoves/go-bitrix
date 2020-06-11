package crm

import (
    "time"
)

type Lead struct {
    Id int `json:"ID"`
    Title string `json:"TITLE"`
    Honorific string `json:"HONORIFIC"`
    Name string `json:"NAME"`
    SecondName string `json:"SECOND_NAME"`
    LastName string `json:"LAST_NAME"`
    Birthdate time.Time `json:"BIRTHDATE"`
    BirthdaySort int `json:"BIRTHDAY_SORT"`
    CompanyTitle string `json:"COMPANY_TITLE"`
    SourceId string `json:"SOURCE_ID"`
    SourceDescription string `json:"SOURCE_DESCRIPTION"`
    StatusId string `json:"STATUS_ID"`
    StatusDescription string `json:"STATUS_DESCRIPTION"`
    StatusSemanticId string `json:"STATUS_SEMANTIC_ID"`
    Post string `json:"POST"`
    Address string `json:"ADDRESS"`
    Address2 string `json:"ADDRESS_2"`
    AddressCity string `json:"ADDRESS_CITY"`
    AddressPostalCode string `json:"ADDRESS_POSTAL_CODE"`
    AddressRegion string `json:"ADDRESS_REGION"`
    AddressProvince string `json:"ADDRESS_PROVINCE"`
    AddressCountry string `json:"ADDRESS_COUNTRY"`
    AddressCountryCode string `json:"ADDRESS_COUNTRY_CODE"`
    CurrencyId string `json:"CURRENCY_ID"`
    Opportunity float64 `json:"OPPORTUNITY"`
    Opened string `json:"OPENED"`
    Comments string `json:"COMMENTS"`
    HasPhone string `json:"HAS_PHONE"`
    HasEmail string `json:"HAS_EMAIL"`
    HasImol string `json:"HAS_IMOL"`
    AssignedById string `json:"ASSIGNED_BY_ID"`
    CreatedById string `json:"CREATED_BY_ID"`
    ModifyById string `json:"MODIFY_BY_ID"`
    DateCreate time.Time `json:"DATE_CREATE"`
    DateModify time.Time `json:"DATE_MODIFY"`
    CompanyId string `json:"COMPANY_ID"`
    ContactId string `json:"CONTACT_ID"`
    ContactIds string `json:"CONTACT_IDS"`
    IsReturnCustomer string `json:"IS_RETURN_CUSTOMER"`
    DateClosed time.Time `json:"DATE_CLOSED"`
    OriginatorId string `json:"ORIGINATOR_ID"`
    OriginId string `json:"ORIGIN_ID"`
    UtmSource string `json:"UTM_SOURCE"`
    UtmMedium string `json:"UTM_MEDIUM"`
    UtmCampaign string `json:"UTM_CAMPAIGN"`
    UtmContent string `json:"UTM_CONTENT"`
    UtmTerm string `json:"UTM_TERM"`
    Phone string `json:"PHONE"`
    Email string `json:"EMAIL"`
    Web string `json:"WEB"`
    Im string `json:"IM"`
    UfCrm1585559766 string `json:"UF_CRM_1585559766"`
    UfCrm1585559808 string `json:"UF_CRM_1585559808"`
    UfCrm1587971702 string `json:"UF_CRM_1587971702"`
    UfCrm1386855792 int `json:"UF_CRM_1386855792"`
    UfCrm1386857211 int `json:"UF_CRM_1386857211"`
    UfCrm1472628139 string `json:"UF_CRM_1472628139"`
    UfUtmTerm string `json:"UF_UTM_TERM"`
    UfUtmContent string `json:"UF_UTM_CONTENT"`
    UfCrm1516886633 string `json:"UF_CRM_1516886633"`
    UfCrm1529520811 string `json:"UF_CRM_1529520811"`
    UfCrm1529520854 time.Time `json:"UF_CRM_1529520854"`
    UfCrm1534324060413 int `json:"UF_CRM_1534324060413"`
    UfCrm1557912745820 string `json:"UF_CRM_1557912745820"`
    UfCrm1557912807301 string `json:"UF_CRM_1557912807301"`

}