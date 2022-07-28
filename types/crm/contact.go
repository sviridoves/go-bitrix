package crm

import "time"

type CrmMultifield struct {
	ID        string `json:"ID"`
	Value     string `json:"VALUE"`
	ValueType string `json:"VALUE_TYPE"`
	TypeID    string `json:"TYPE_ID"`
}

type photo struct {
	ID          int    `json:"id"`
	ShowURL     string `json:"showUrl"`
	DownloadURL string `json:"downloadUrl"`
}

type Contact struct {
	Id                int             `json:"ID,string"`
	Honorific         string          `json:"HONORIFIC"`
	Name              string          `json:"NAME"`
	SecondName        string          `json:"SECOND_NAME"`
	LastName          string          `json:"LAST_NAME"`
	FullName          string          `json:"FULL_NAME"`
	Photo             photo           `json:"PHOTO,omitempty"`
	Birthdate         string          `json:"BIRTHDATE"`
	TypeId            string          `json:"TYPE_ID"`
	SourceId          string          `json:"SOURCE_ID"`
	SourceDescription string          `json:"SOURCE_DESCRIPTION"`
	Post              string          `json:"POST"`
	Address           string          `json:"ADDRESS"`
	Address2          string          `json:"ADDRESS_2"`
	AddressCity       string          `json:"ADDRESS_CITY"`
	AddressPostalCode string          `json:"ADDRESS_POSTAL_CODE"`
	AddressRegion     string          `json:"ADDRESS_REGION"`
	AddressProvince   string          `json:"ADDRESS_PROVINCE"`
	AddressCountry    string          `json:"ADDRESS_COUNTRY"`
	Comments          string          `json:"COMMENTS"`
	Opened            string          `json:"OPENED"`
	Export            string          `json:"EXPORT"`
	HasPhone          string          `json:"HAS_PHONE"`
	HasEmail          string          `json:"HAS_EMAIL"`
	HasImol           string          `json:"HAS_IMOL"`
	AssignedById      int             `json:"ASSIGNED_BY_ID,string"`
	CreatedById       int             `json:"CREATED_BY_ID,string"`
	ModifyById        int             `json:"MODIFY_BY_ID,string"`
	DateCreate        time.Time       `json:"DATE_CREATE"`
	DateModify        time.Time       `json:"DATE_MODIFY"`
	CompanyId         string          `json:"COMPANY_ID"`
	LeadId            int             `json:"LEAD_ID,string"`
	OriginatorId      string          `json:"ORIGINATOR_ID"`
	OriginId          string          `json:"ORIGIN_ID"`
	OriginVersion     string          `json:"ORIGIN_VERSION"`
	FaceId            int             `json:"FACE_ID,string"`
	UtmSource         string          `json:"UTM_SOURCE"`
	UtmMedium         string          `json:"UTM_MEDIUM"`
	UtmCampaign       string          `json:"UTM_CAMPAIGN"`
	UtmContent        string          `json:"UTM_CONTENT"`
	UtmTerm           string          `json:"UTM_TERM"`
	Phone             []CrmMultifield `json:"PHONE,omitempty"`
	Email             []CrmMultifield `json:"EMAIL,omitempty"`
	Web               []CrmMultifield `json:"WEB,omitempty"`

	UfCrm1543654486    interface{}   `json:"UF_CRM_1543654486"`
	UfCrm1543588165    string        `json:"UF_CRM_1543588165"`
	UfCrm5E81C9A0Bdb0E string        `json:"UF_CRM_5E81C9A0BDB0E"`
	UfCrm5E81C9A1F1125 interface{}   `json:"UF_CRM_5E81C9A1F1125"`
	UfCrm5Ea68Aa2Dec53 interface{}   `json:"UF_CRM_5EA68AA2DEC53"`
	UfCrm1409141850    []interface{} `json:"UF_CRM_1409141850"`
	UfCrm560548Ef6Fc70 interface{}   `json:"UF_CRM_560548EF6FC70"`
	UfCrm560548Ef73E94 []interface{} `json:"UF_CRM_560548EF73E94"`
	UfCrm57C93B8A19A74 string        `json:"UF_CRM_57C93B8A19A74"`
	UfCrm1497382470    interface{}   `json:"UF_CRM_1497382470"`
	UfCrm1497884835    interface{}   `json:"UF_CRM_1497884835"`
	UfCrm1498318041    interface{}   `json:"UF_CRM_1498318041"`
	UfCrm1498318077    interface{}   `json:"UF_CRM_1498318077"`
	UfCrm1499446373    interface{}   `json:"UF_CRM_1499446373"`
	UfCrm5A6845A20Bd9D interface{}   `json:"UF_CRM_5A6845A20BD9D"`
	UfCrm5A6845A28B9Ab interface{}   `json:"UF_CRM_5A6845A28B9AB"`
	UfCrm5A69Dcdec9B3C interface{}   `json:"UF_CRM_5A69DCDEC9B3C"`
	UfCrm5B2B8A93Dcd98 interface{}   `json:"UF_CRM_5B2B8A93DCD98"`
	UfCrm5B2B8A9496209 string        `json:"UF_CRM_5B2B8A9496209"`
	UfCrm5B756Db460Dc3 []interface{} `json:"UF_CRM_5B756DB460DC3"`
	UfCrm1548849779    interface{}   `json:"UF_CRM_1548849779"`
	UfCrm1548850096    interface{}   `json:"UF_CRM_1548850096"`
	UfCrm1548850308    interface{}   `json:"UF_CRM_1548850308"`
	UfCrm1548850343    interface{}   `json:"UF_CRM_1548850343"`
	UfCrm1548850386    interface{}   `json:"UF_CRM_1548850386"`
	UfCrm1548850419    interface{}   `json:"UF_CRM_1548850419"`
	UfCrm1548850472    interface{}   `json:"UF_CRM_1548850472"`
	UfCrm1548850489    interface{}   `json:"UF_CRM_1548850489"`
	UfCrm1548850518    interface{}   `json:"UF_CRM_1548850518"`
	UfCrm1548938252    interface{}   `json:"UF_CRM_1548938252"`
	UfCrm1554804867    interface{}   `json:"UF_CRM_1554804867"`
	UfCrm5Cdc0F609Baea interface{}   `json:"UF_CRM_5CDC0F609BAEA"`
	UfCrm5Cdc0F613902F interface{}   `json:"UF_CRM_5CDC0F613902F"`
	UfCrm1564650862    interface{}   `json:"UF_CRM_1564650862"`
	UfCrm1592316858906 interface{}   `json:"UF_CRM_1592316858906"`
	UfCrm1592316910662 interface{}   `json:"UF_CRM_1592316910662"`
	UfCrm1653314996    interface{}   `json:"UF_CRM_1653314996"`
	UfCrm1653315011    interface{}   `json:"UF_CRM_1653315011"`
	UfCrm1653315101    interface{}   `json:"UF_CRM_1653315101"`
	UfCrm1653388064    interface{}   `json:"UF_CRM_1653388064"`
	UfCrm1653388163    interface{}   `json:"UF_CRM_1653388163"`
	UfCrm1653388249    interface{}   `json:"UF_CRM_1653388249"`
	UfCrm1653388278    interface{}   `json:"UF_CRM_1653388278"`
	UfCrm1653388306    interface{}   `json:"UF_CRM_1653388306"`
	UfCrm5Eea226D65F50 interface{}   `json:"UF_CRM_5EEA226D65F50"`
}
