package model

import "time"

// BaseModel 公共 Model 字段，不适用 DeleteAt 作软删除
type BaseModel struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
