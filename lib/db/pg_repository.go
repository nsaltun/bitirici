package db

import "gorm.io/gorm"

type GormResponse struct {
	*gorm.DB
}

// This is just an example. We can customize repository methods this way in future.
func (p *PgDB) Create(value interface{}) *gorm.DB {
	return p.DB.Create(value)
}
