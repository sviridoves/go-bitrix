package crm

import (
    "time"
)

type Activity struct {
    Id int `json:"ID"`
    OwnerId int `json:"OWNER_ID"`
    OwnerTypeId string `json:"OWNER_TYPE_ID"`
    TypeId string `json:"TYPE_ID"`
    ProviderId string `json:"PROVIDER_ID"`
    ProviderTypeId string `json:"PROVIDER_TYPE_ID"`
    ProviderGroupId string `json:"PROVIDER_GROUP_ID"`
    AssociatedEntityId int `json:"ASSOCIATED_ENTITY_ID"`
    Subject string `json:"SUBJECT"`
    StartTime time.Time `json:"START_TIME"`
    EndTime time.Time `json:"END_TIME"`
    Deadline time.Time `json:"DEADLINE"`
    Completed string `json:"COMPLETED"`
    Status string `json:"STATUS"`
    ResponsibleId string `json:"RESPONSIBLE_ID"`
    Priority string `json:"PRIORITY"`
    NotifyType string `json:"NOTIFY_TYPE"`
    NotifyValue int `json:"NOTIFY_VALUE"`
    Description string `json:"DESCRIPTION"`
    DescriptionType string `json:"DESCRIPTION_TYPE"`
    Direction string `json:"DIRECTION"`
    Location string `json:"LOCATION"`
    Created time.Time `json:"CREATED"`
    AuthorId string `json:"AUTHOR_ID"`
    LastUpdated time.Time `json:"LAST_UPDATED"`
    EditorId string `json:"EDITOR_ID"`
    Settings string `json:"SETTINGS"`
    OriginId string `json:"ORIGIN_ID"`
    OriginatorId string `json:"ORIGINATOR_ID"`
    ResultStatus int `json:"RESULT_STATUS"`
    ResultStream int `json:"RESULT_STREAM"`
    ResultSourceId string `json:"RESULT_SOURCE_ID"`
    ProviderParams string `json:"PROVIDER_PARAMS"`
    ProviderData string `json:"PROVIDER_DATA"`
    ResultMark int `json:"RESULT_MARK"`
    ResultValue float64 `json:"RESULT_VALUE"`
    ResultSum float64 `json:"RESULT_SUM"`
    ResultCurrencyId string `json:"RESULT_CURRENCY_ID"`
    AutocompleteRule int `json:"AUTOCOMPLETE_RULE"`
    Bindings string `json:"BINDINGS"`
    Communications string `json:"COMMUNICATIONS"`
    Files string `json:"FILES"`
    WebdavElements string `json:"WEBDAV_ELEMENTS"`

}