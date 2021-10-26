package models

type DbType uint8

const (
	Postgres DbType = iota
	Mysql
	Oracle
)