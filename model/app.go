package model

type App struct {
	BaseModel
	tableName   struct{} `sql:"app"`
	Name        string
	Description string
}

type CreateAppReq struct {
	Name        string `json: nmae`
	Description string `json: description`
}

type CreateAppResp struct {
	ID          int64  `json: id`
	Name        string `json: name`
	Description string `json: description`
}
