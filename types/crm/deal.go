package crm

import (
    "time"
)

type Deal struct {
    Id int `json:"ID"`
    Title string `json:"TITLE"`
    TypeId string `json:"TYPE_ID"`
    CategoryId string `json:"CATEGORY_ID"`
    StageId string `json:"STAGE_ID"`
    StageSemanticId string `json:"STAGE_SEMANTIC_ID"`
    IsNew string `json:"IS_NEW"`
    IsRecurring string `json:"IS_RECURRING"`
    IsReturnCustomer string `json:"IS_RETURN_CUSTOMER"`
    IsRepeatedApproach string `json:"IS_REPEATED_APPROACH"`
    Probability int `json:"PROBABILITY"`
    CurrencyId string `json:"CURRENCY_ID"`
    Opportunity float64 `json:"OPPORTUNITY"`
    IsManualOpportunity string `json:"IS_MANUAL_OPPORTUNITY"`
    TaxValue float64 `json:"TAX_VALUE"`
    ExchRate float64 `json:"EXCH_RATE"`
    AccountCurrencyId string `json:"ACCOUNT_CURRENCY_ID"`
    OpportunityAccount float64 `json:"OPPORTUNITY_ACCOUNT"`
    TaxValueAccount float64 `json:"TAX_VALUE_ACCOUNT"`
    CompanyId string `json:"COMPANY_ID"`
    ContactId string `json:"CONTACT_ID"`
    ContactIds string `json:"CONTACT_IDS"`
    QuoteId string `json:"QUOTE_ID"`
    Begindate time.Time `json:"BEGINDATE"`
    Closedate time.Time `json:"CLOSEDATE"`
    Opened string `json:"OPENED"`
    Closed string `json:"CLOSED"`
    Comments string `json:"COMMENTS"`
    AssignedById string `json:"ASSIGNED_BY_ID"`
    CreatedById string `json:"CREATED_BY_ID"`
    ModifyById string `json:"MODIFY_BY_ID"`
    DateCreate time.Time `json:"DATE_CREATE"`
    DateModify time.Time `json:"DATE_MODIFY"`
    SourceId string `json:"SOURCE_ID"`
    SourceDescription string `json:"SOURCE_DESCRIPTION"`
    LeadId string `json:"LEAD_ID"`
    AdditionalInfo string `json:"ADDITIONAL_INFO"`
    LocationId string `json:"LOCATION_ID"`
    OriginatorId string `json:"ORIGINATOR_ID"`
    OriginId string `json:"ORIGIN_ID"`
    UtmSource string `json:"UTM_SOURCE"`
    UtmMedium string `json:"UTM_MEDIUM"`
    UtmCampaign string `json:"UTM_CAMPAIGN"`
    UtmContent string `json:"UTM_CONTENT"`
    UtmTerm string `json:"UTM_TERM"`
    UfCrm1436354924 string `json:"UF_CRM_1436354924"`
    UfCrm1584371569872 string `json:"UF_CRM_1584371569872"`
    UfCrm5e81c99fefe66 string `json:"UF_CRM_5E81C99FEFE66"`
    UfCrm5e81c9a06933f string `json:"UF_CRM_5E81C9A06933F"`
    UfCrm1586421147 int `json:"UF_CRM_1586421147"`
    UfCrm5ea68aa27e342 string `json:"UF_CRM_5EA68AA27E342"`
    UfCrm1573469032 string `json:"UF_CRM_1573469032"`
    UfCrm1586855542 string `json:"UF_CRM_1586855542"`
    UfCrm1438589528 int `json:"UF_CRM_1438589528"`
    UfCrm560548ef58789 int `json:"UF_CRM_560548EF58789"`
    UfCrm560548ef67bf3 int `json:"UF_CRM_560548EF67BF3"`
    UfCrm57c93b89af653 string `json:"UF_CRM_57C93B89AF653"`
    UfCrm1481552877 string `json:"UF_CRM_1481552877"`
    UfCrm1498928035 string `json:"UF_CRM_1498928035"`
    UfCrm1500371492 time.Time `json:"UF_CRM_1500371492"`
    UfCrm1500371525 time.Time `json:"UF_CRM_1500371525"`
    UfCrmB24Portal string `json:"UF_CRM_B24_PORTAL"`
    UfCrmLicenseKey string `json:"UF_CRM_LICENSE_KEY"`
    UfCrm59e71cad9783d string `json:"UF_CRM_59E71CAD9783D"`
    UfCrm59e71cae18f49 string `json:"UF_CRM_59E71CAE18F49"`
    UfCrm59e71cae4902e string `json:"UF_CRM_59E71CAE4902E"`
    UfCrmCity string `json:"UF_CRM_CITY"`
    UfCrm5b2cb0c557d86 string `json:"UF_CRM_5B2CB0C557D86"`
    UfCrm5b2cb0c5a1542 time.Time `json:"UF_CRM_5B2CB0C5A1542"`
    UfCrm5b9798d319d3b int `json:"UF_CRM_5B9798D319D3B"`
    UfCrm5cdc0f61c3ba1 string `json:"UF_CRM_5CDC0F61C3BA1"`
    UfCrm5cdc0f6206793 string `json:"UF_CRM_5CDC0F6206793"`
    UfCrm1561732218126 string `json:"UF_CRM_1561732218126"`
    UfCrm1561732303934 string `json:"UF_CRM_1561732303934"`
    UfCrm1561992408 string `json:"UF_CRM_1561992408"`
    UfCrm1563952925 string `json:"UF_CRM_1563952925"`
    UfCrm1569577102589 string `json:"UF_CRM_1569577102589"`
    UfCrm1575315051542 string `json:"UF_CRM_1575315051542"`
    UfCrm1576151727490 string `json:"UF_CRM_1576151727490"`
    UfCrm1576151920228 string `json:"UF_CRM_1576151920228"`
    UfCrm1576158789522 float64 `json:"UF_CRM_1576158789522"`
    UfCrm1581876109368 time.Time `json:"UF_CRM_1581876109368"`
    UfCrm5ed9ff182af71 string `json:"UF_CRM_5ED9FF182AF71"`
    UfCrm1584710786 string `json:"UF_CRM_1584710786"`
    UfCrm1587635693 string `json:"UF_CRM_1587635693"`
    UfCrm1587635762 string `json:"UF_CRM_1587635762"`
    UfCrm1587635996 float64 `json:"UF_CRM_1587635996"`
    UfCrm1587635897 string `json:"UF_CRM_1587635897"`

}