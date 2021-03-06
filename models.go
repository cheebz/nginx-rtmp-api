package main

import (
	"database/sql"
	"time"
)

// Configuration struct
type Configuration struct {
	Debug      bool       `json:"debug"`
	Port       int        `json:"port"`
	Db         DataSource `json:"db"`
	Filesystem bool       `json:"filesystem"`
}

// DataSource struct
type DataSource struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

// Stream struct
type Stream struct {
	ID        int          `json:"id"`
	Title     string       `json:"title"`
	Live      bool         `json:"live"`
	StartTime time.Time    `json:"start_time"`
	EndTime   sql.NullTime `json:"end_time"`
	UserID    int          `json:"user_id"`
}
