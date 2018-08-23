package db

import (
	_ "github.com/go-sql-driver/mysql"
)

// =============================
// 数据库表结构体
// =============================
type Message struct {
	Id        int64  `json:"id"`
	Index     string `json:"index"`
	Type      int64  `json:"type"`
	Uid       string `json:"uid"`
	Body      string `json:"body"`
	Status    int8   `json:"status"`
	CreatedAt int64  `json:"created_at"`
}
