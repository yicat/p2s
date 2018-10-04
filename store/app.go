package store

import (
	"github.com/yicat/p2s/model"
	"time"
)

func (s *Store) GetApp(appID int64) (*model.App, error) {
	var app model.App
	err := s.DB.Model(&app).Where("id = ?", appID).Select()
	return &app, err
}

func (ds *Store) GetAppList(offset, limit int) ([]model.App, error) {
	appList := make([]model.App, 0)
	err := ds.DB.Model(&appList).Offset(offset).Limit(limit).Select()
	return appList, err
}

func (ds *Store) CreateApp(app *model.App) error {
	app.CreatedAt = time.Now()
	app.UpdatedAt = time.Now()
	err := ds.DB.Insert(app)
	return err
}

func (ds *Store) DeleteApp(app *model.App) error {
	err := ds.DB.Delete(app)
	return err
}

func (ds *Store) UpdateApp(app *model.App) error {
	err := ds.DB.Update(app)
	return err
}
