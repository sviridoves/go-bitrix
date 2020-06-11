package crm

import (
    "time"
)

type Contact struct {
    Id int `json:"ID"`
    Honorific string `json:"HONORIFIC"`
    Name string `json:"NAME"`
    SecondName string `json:"SECOND_NAME"`
    LastName string `json:"LAST_NAME"`
    FullName string `json:"FULL_NAME"`
    Photo int `json:"PHOTO"`
    Birthdate time.Time `json:"BIRTHDATE"`
    BirthdaySort int `json:"BIRTHDAY_SORT"`
    TypeId string `json:"TYPE_ID"`
    SourceId string `json:"SOURCE_ID"`
    SourceDescription string `json:"SOURCE_DESCRIPTION"`
    Post string `json:"POST"`
    Address string `json:"ADDRESS"`
    Address2 string `json:"ADDRESS_2"`
    AddressCity string `json:"ADDRESS_CITY"`
    AddressPostalCode string `json:"ADDRESS_POSTAL_CODE"`
    AddressRegion string `json:"ADDRESS_REGION"`
    AddressProvince string `json:"ADDRESS_PROVINCE"`
    AddressCountry string `json:"ADDRESS_COUNTRY"`
    AddressCountryCode string `json:"ADDRESS_COUNTRY_CODE"`
    Comments string `json:"COMMENTS"`
    Opened string `json:"OPENED"`
    Export string `json:"EXPORT"`
    HasPhone string `json:"HAS_PHONE"`
    HasEmail string `json:"HAS_EMAIL"`
    HasImol string `json:"HAS_IMOL"`
    AssignedById string `json:"ASSIGNED_BY_ID"`
    CreatedById string `json:"CREATED_BY_ID"`
    ModifyById string `json:"MODIFY_BY_ID"`
    DateCreate time.Time `json:"DATE_CREATE"`
    DateModify time.Time `json:"DATE_MODIFY"`
    CompanyId string `json:"COMPANY_ID"`
    CompanyIds string `json:"COMPANY_IDS"`
    LeadId string `json:"LEAD_ID"`
    OriginatorId string `json:"ORIGINATOR_ID"`
    OriginId string `json:"ORIGIN_ID"`
    OriginVersion string `json:"ORIGIN_VERSION"`
    FaceId int `json:"FACE_ID"`
    UtmSource string `json:"UTM_SOURCE"`
    UtmMedium string `json:"UTM_MEDIUM"`
    UtmCampaign string `json:"UTM_CAMPAIGN"`
    UtmContent string `json:"UTM_CONTENT"`
    UtmTerm string `json:"UTM_TERM"`
    Phone string `json:"PHONE"`
    Email string `json:"EMAIL"`
    Web string `json:"WEB"`
    Im string `json:"IM"`
    UfCrm1543654486 string `json:"UF_CRM_1543654486"`
    UfCrm1543588165 string `json:"UF_CRM_1543588165"`
    UfCrm5e81c9a0bdb0e string `json:"UF_CRM_5E81C9A0BDB0E"`
    UfCrm5e81c9a1f1125 string `json:"UF_CRM_5E81C9A1F1125"`
    UfCrm5ea68aa2dec53 string `json:"UF_CRM_5EA68AA2DEC53"`
    UfCrm1409141850 int `json:"UF_CRM_1409141850"`
    UfCrm560548ef6fc70 int `json:"UF_CRM_560548EF6FC70"`
    UfCrm560548ef73e94 int `json:"UF_CRM_560548EF73E94"`
    UfCrm57c93b8a19a74 string `json:"UF_CRM_57C93B8A19A74"`
    UfCrm1497382470 string `json:"UF_CRM_1497382470"`
    UfCrm1497884835 string `json:"UF_CRM_1497884835"`
    UfCrm1498318041 string `json:"UF_CRM_1498318041"`
    UfCrm1498318077 string `json:"UF_CRM_1498318077"`
    UfCrm1499446373 string `json:"UF_CRM_1499446373"`
    UfCrm5a6845a20bd9d string `json:"UF_CRM_5A6845A20BD9D"`
    UfCrm5a6845a28b9ab string `json:"UF_CRM_5A6845A28B9AB"`
    UfCrm5a69dcdec9b3c string `json:"UF_CRM_5A69DCDEC9B3C"`
    UfCrm5b2b8a93dcd98 string `json:"UF_CRM_5B2B8A93DCD98"`
    UfCrm5b2b8a9496209 time.Time `json:"UF_CRM_5B2B8A9496209"`
    UfCrm5b756db460dc3 int `json:"UF_CRM_5B756DB460DC3"`
    UfCrm1548849779 string `json:"UF_CRM_1548849779"`
    UfCrm1548850096 string `json:"UF_CRM_1548850096"`
    UfCrm1548850308 string `json:"UF_CRM_1548850308"`
    UfCrm1548850343 string `json:"UF_CRM_1548850343"`
    UfCrm1548850386 string `json:"UF_CRM_1548850386"`
    UfCrm1548850419 string `json:"UF_CRM_1548850419"`
    UfCrm1548850472 string `json:"UF_CRM_1548850472"`
    UfCrm1548850489 string `json:"UF_CRM_1548850489"`
    UfCrm1548850518 string `json:"UF_CRM_1548850518"`
    UfCrm1548938252 string `json:"UF_CRM_1548938252"`
    UfCrm1554804867 string `json:"UF_CRM_1554804867"`
    UfCrm5cdc0f609baea string `json:"UF_CRM_5CDC0F609BAEA"`
    UfCrm5cdc0f613902f string `json:"UF_CRM_5CDC0F613902F"`
    UfCrm1564650862 string `json:"UF_CRM_1564650862"`

}