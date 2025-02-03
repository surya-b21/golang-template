package model

type MyClient struct {
	Base
	MyClientApi
}

type MyClientApi struct {
	Name         *string `json:"name,omitempty" gorm:"type:char(250);not null"`
	Slug         *string `json:"slug,omitempty" gorm:"type:char(100);not null"`
	IsProject    *string `json:"is_project,omitempty" gorm:"type:varchar(30);check: is_project in ('0','1'); not null; default:'0'"`
	SelfCapture  *string `json:"self_capture,omitempty" gorm:"type:char(1); not null; default: '1"`
	ClientPrefix *string `json:"client_prefix,omitempty" gorm:"type:char(4); not null"`
	ClientLogo   *string `json:"client_logo,omitempty" gorm:"type:char(255) not null; default: 'no-image.jpg"`
	Address      *string `json:"address,omitempty" gorm:"type:text"`
	City         *string `json:"city,omitempty" gorm:"type:char(50)"`
}
