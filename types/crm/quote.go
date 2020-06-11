package crm

import (
    "time"
)

type Quote struct {
    Id int `json:"ID"`
    QuoteNumber string `json:"QUOTE_NUMBER"`
    Title string `json:"TITLE"`
    StatusId string `json:"STATUS_ID"`
    CurrencyId string `json:"CURRENCY_ID"`
    Opportunity float64 `json:"OPPORTUNITY"`
    TaxValue float64 `json:"TAX_VALUE"`
    ExchRate float64 `json:"EXCH_RATE"`
    AccountCurrencyId string `json:"ACCOUNT_CURRENCY_ID"`
    OpportunityAccount float64 `json:"OPPORTUNITY_ACCOUNT"`
    TaxValueAccount float64 `json:"TAX_VALUE_ACCOUNT"`
    CompanyId string `json:"COMPANY_ID"`
    MycompanyId string `json:"MYCOMPANY_ID"`
    ContactId string `json:"CONTACT_ID"`
    ContactIds string `json:"CONTACT_IDS"`
    Begindate time.Time `json:"BEGINDATE"`
    Closedate time.Time `json:"CLOSEDATE"`
    Opened string `json:"OPENED"`
    Closed string `json:"CLOSED"`
    Comments string `json:"COMMENTS"`
    Content string `json:"CONTENT"`
    Terms string `json:"TERMS"`
    ClientTitle string `json:"CLIENT_TITLE"`
    ClientAddr string `json:"CLIENT_ADDR"`
    ClientContact string `json:"CLIENT_CONTACT"`
    ClientEmail string `json:"CLIENT_EMAIL"`
    ClientPhone string `json:"CLIENT_PHONE"`
    ClientTpId string `json:"CLIENT_TP_ID"`
    ClientTpaId string `json:"CLIENT_TPA_ID"`
    AssignedById string `json:"ASSIGNED_BY_ID"`
    CreatedById string `json:"CREATED_BY_ID"`
    ModifyById string `json:"MODIFY_BY_ID"`
    DateCreate time.Time `json:"DATE_CREATE"`
    DateModify time.Time `json:"DATE_MODIFY"`
    LeadId string `json:"LEAD_ID"`
    DealId string `json:"DEAL_ID"`
    PersonTypeId int `json:"PERSON_TYPE_ID"`
    LocationId string `json:"LOCATION_ID"`
    UtmSource string `json:"UTM_SOURCE"`
    UtmMedium string `json:"UTM_MEDIUM"`
    UtmCampaign string `json:"UTM_CAMPAIGN"`
    UtmContent string `json:"UTM_CONTENT"`
    UtmTerm string `json:"UTM_TERM"`
    UfCrm598073318401f string `json:"UF_CRM_598073318401F"`
    UfCrm5ddfd49c6407b string `json:"UF_CRM_5DDFD49C6407B"`
    UfCrm59807332c8d4e int `json:"UF_CRM_59807332C8D4E"`
    UfCrm5980733306522 int `json:"UF_CRM_5980733306522"`
    UfCrm598073332feba int `json:"UF_CRM_598073332FEBA"`
    UfCrm5980733360f7a string `json:"UF_CRM_5980733360F7A"`
    UfCrm598073339eb0c string `json:"UF_CRM_598073339EB0C"`
    UfCrm59807333d0584 string `json:"UF_CRM_59807333D0584"`
    UfCrm5980733407e35 time.Time `json:"UF_CRM_5980733407E35"`
    UfCrm5980733434d59 time.Time `json:"UF_CRM_5980733434D59"`
    UfCrm5b30fb884b331 string `json:"UF_CRM_5B30FB884B331"`
    UfCrm5b30fb887c221 string `json:"UF_CRM_5B30FB887C221"`
    UfCrm5b30fb88a4bd2 string `json:"UF_CRM_5B30FB88A4BD2"`
    UfCrm5b30fb88cf842 string `json:"UF_CRM_5B30FB88CF842"`
    UfCrm5b30fb8903a93 string `json:"UF_CRM_5B30FB8903A93"`
    UfCrm5b30fb892fb5b string `json:"UF_CRM_5B30FB892FB5B"`
    UfCrm5b30fb8955190 string `json:"UF_CRM_5B30FB8955190"`
    UfCrm5b30fb89795de time.Time `json:"UF_CRM_5B30FB89795DE"`
    UfCrm5ce3edfef1081 int `json:"UF_CRM_5CE3EDFEF1081"`
    UfCrm5ce3edff3a378 string `json:"UF_CRM_5CE3EDFF3A378"`
    UfCrm5ce3edff60f6a string `json:"UF_CRM_5CE3EDFF60F6A"`
    UfCrm5d41c62d53c1a string `json:"UF_CRM_5D41C62D53C1A"`
    UfCrm5d41c62daca93 string `json:"UF_CRM_5D41C62DACA93"`
    UfCrm5d41c62df1db5 string `json:"UF_CRM_5D41C62DF1DB5"`
    UfCrm5d41c62e23274 string `json:"UF_CRM_5D41C62E23274"`
    UfCrm5ddfd49ccb1a4 string `json:"UF_CRM_5DDFD49CCB1A4"`

}