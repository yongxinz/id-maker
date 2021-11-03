package entity

import "time"

type TimeFormat time.Time

func (t *TimeFormat) MarshalJSON() ([]byte, error) {
	if time.Time(*t).IsZero() {
		return []byte(`""`), nil
	}
	return []byte(`"` + time.Time(*t).Format("2006-01-02 15:04:05") + `"`), nil
}

type Segments struct {
	BizTag     string     `xorm:"not null pk VARCHAR(32) 'biz_tag'" json:"biz_tag" validate:"required,max=32"`
	MaxId      int64      `xorm:"BIGINT(20) 'max_id'" json:"max_id" validate:"required"`
	Step       int64      `xorm:"INT(11) 'step'" json:"step" validate:"required"`
	Remark     string     `xorm:"VARCHAR(200) 'remark'" json:"remark"`
	CreateTime TimeFormat `xorm:"created" json:"create_time"`
	UpdateTime TimeFormat `xorm:"updated" json:"update_time"`
}

func (s *Segments) TableName() string {
	return "segments"
}
