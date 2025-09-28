package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	// dsn := "host=db.thwtjnsdkjkisdooxkjm.supabase.co user=postgres password=bdhqtx0fZkAIjTjg dbname=postgres port=6543 sslmode=require"
	// dsn := "postgresql://postgres:bdhqtx0fZkAIjTjg@db.thwtjnsdkjkisdooxkjm.supabase.co:5432/postgres"
	dsn := "postgresql://postgres.thwtjnsdkjkisdooxkjm:bdhqtx0fZkAIjTjg@aws-1-ap-south-1.pooler.supabase.com:6543/postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
