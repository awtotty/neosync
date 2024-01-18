// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: system.sql

package pg_queries

import (
	"context"
)

const getDatabaseSchema = `-- name: GetDatabaseSchema :many
SELECT
	c.table_schema,
	c.table_name,
	c.column_name,
	c.ordinal_position,
	COALESCE(c.column_default, 'NULL') as column_default, -- must coalesce because sqlc doesn't appear to work for system structs to output a *string
	c.is_nullable,
	c.data_type
FROM
	information_schema.columns AS c
	JOIN information_schema.tables AS t ON c.table_schema = t.table_schema
		AND c.table_name = t.table_name
WHERE
	c.table_schema NOT IN('pg_catalog', 'information_schema')
	AND t.table_type = 'BASE TABLE'
`

type GetDatabaseSchemaRow struct {
	TableSchema     string
	TableName       string
	ColumnName      string
	OrdinalPosition int
	ColumnDefault   string
	IsNullable      string
	DataType        string
}

func (q *Queries) GetDatabaseSchema(ctx context.Context, db DBTX) ([]*GetDatabaseSchemaRow, error) {
	rows, err := db.Query(ctx, getDatabaseSchema)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetDatabaseSchemaRow
	for rows.Next() {
		var i GetDatabaseSchemaRow
		if err := rows.Scan(
			&i.TableSchema,
			&i.TableName,
			&i.ColumnName,
			&i.OrdinalPosition,
			&i.ColumnDefault,
			&i.IsNullable,
			&i.DataType,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getDatabaseTableSchema = `-- name: GetDatabaseTableSchema :many
SELECT
	c.table_schema,
	c.table_name,
	c.column_name,
	c.ordinal_position,
	COALESCE(c.column_default, 'NULL') as column_default, -- must coalesce because sqlc doesn't appear to work for system structs to output a *string
	c.is_nullable,
	c.data_type,
    c.character_maximum_length,
    c.numeric_precision,
    c.numeric_scale
FROM
	information_schema.columns AS c
	JOIN information_schema.tables AS t ON c.table_schema = t.table_schema
		AND c.table_name = t.table_name
WHERE
	c.table_schema = $1 AND t.table_name = $2
	AND t.table_type = 'BASE TABLE'
	ORDER BY c.ordinal_position ASC
`

type GetDatabaseTableSchemaParams struct {
	TableSchema string
	TableName   string
}

type GetDatabaseTableSchemaRow struct {
	TableSchema            string
	TableName              string
	ColumnName             string
	OrdinalPosition        int
	ColumnDefault          string
	IsNullable             string
	DataType               string
	CharacterMaximumLength interface{}
	NumericPrecision       interface{}
	NumericScale           interface{}
}

func (q *Queries) GetDatabaseTableSchema(ctx context.Context, db DBTX, arg *GetDatabaseTableSchemaParams) ([]*GetDatabaseTableSchemaRow, error) {
	rows, err := db.Query(ctx, getDatabaseTableSchema, arg.TableSchema, arg.TableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetDatabaseTableSchemaRow
	for rows.Next() {
		var i GetDatabaseTableSchemaRow
		if err := rows.Scan(
			&i.TableSchema,
			&i.TableName,
			&i.ColumnName,
			&i.OrdinalPosition,
			&i.ColumnDefault,
			&i.IsNullable,
			&i.DataType,
			&i.CharacterMaximumLength,
			&i.NumericPrecision,
			&i.NumericScale,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getForeignKeyConstraints = `-- name: GetForeignKeyConstraints :many
	SELECT
    rc.constraint_name
    ,
    kcu.table_schema AS schema_name
    ,
    kcu.table_name
    ,
    kcu.column_name
    ,
    c.is_nullable
    ,
    kcu2.table_schema AS foreign_schema_name
    ,
    kcu2.table_name AS foreign_table_name
    ,
    kcu2.column_name AS foreign_column_name
FROM
    information_schema.referential_constraints rc
JOIN information_schema.key_column_usage kcu
    ON
    kcu.constraint_name = rc.constraint_name
JOIN information_schema.key_column_usage kcu2
    ON
    kcu2.ordinal_position = kcu.position_in_unique_constraint
    AND kcu2.constraint_name = rc.unique_constraint_name
JOIN information_schema.columns as c
	ON
	c.table_schema = kcu.table_schema 
	AND c.table_name = kcu.table_name 
	AND c.column_name = kcu.column_name
WHERE
    kcu.table_schema = $1
ORDER BY
    rc.constraint_name,
    kcu.ordinal_position
`

type GetForeignKeyConstraintsRow struct {
	ConstraintName    string
	SchemaName        string
	TableName         string
	ColumnName        string
	IsNullable        string
	ForeignSchemaName string
	ForeignTableName  string
	ForeignColumnName string
}

func (q *Queries) GetForeignKeyConstraints(ctx context.Context, db DBTX, tableschema string) ([]*GetForeignKeyConstraintsRow, error) {
	rows, err := db.Query(ctx, getForeignKeyConstraints, tableschema)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetForeignKeyConstraintsRow
	for rows.Next() {
		var i GetForeignKeyConstraintsRow
		if err := rows.Scan(
			&i.ConstraintName,
			&i.SchemaName,
			&i.TableName,
			&i.ColumnName,
			&i.IsNullable,
			&i.ForeignSchemaName,
			&i.ForeignTableName,
			&i.ForeignColumnName,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPrimaryKeyConstraints = `-- name: GetPrimaryKeyConstraints :many
SELECT 
    tc.table_schema AS schema_name,
    tc.table_name as table_name,
    tc.constraint_name as constraint_name, 
    kcu.column_name as column_name 
FROM 
    information_schema.table_constraints AS tc 
JOIN information_schema.key_column_usage AS kcu
    ON tc.constraint_name = kcu.constraint_name
    AND tc.table_schema = kcu.table_schema
WHERE 
    tc.table_schema = $1
    AND tc.constraint_type = 'PRIMARY KEY'
ORDER BY 
    tc.table_name, 
    kcu.column_name
`

type GetPrimaryKeyConstraintsRow struct {
	SchemaName     string
	TableName      string
	ConstraintName string
	ColumnName     string
}

func (q *Queries) GetPrimaryKeyConstraints(ctx context.Context, db DBTX, tableschema string) ([]*GetPrimaryKeyConstraintsRow, error) {
	rows, err := db.Query(ctx, getPrimaryKeyConstraints, tableschema)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetPrimaryKeyConstraintsRow
	for rows.Next() {
		var i GetPrimaryKeyConstraintsRow
		if err := rows.Scan(
			&i.SchemaName,
			&i.TableName,
			&i.ConstraintName,
			&i.ColumnName,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTableConstraints = `-- name: GetTableConstraints :many
SELECT
    nsp.nspname AS db_schema,
    rel.relname AS table_name,
    con.conname AS constraint_name,
    pg_get_constraintdef(con.oid) AS constraint_definition
FROM
    pg_catalog.pg_constraint con
INNER JOIN pg_catalog.pg_class rel
                       ON
    rel.oid = con.conrelid
INNER JOIN pg_catalog.pg_namespace nsp
                       ON
    nsp.oid = connamespace
WHERE
    nsp.nspname = $1 AND rel.relname = $2
`

type GetTableConstraintsParams struct {
	Schema string
	Table  string
}

type GetTableConstraintsRow struct {
	DbSchema             string
	TableName            string
	ConstraintName       string
	ConstraintDefinition string
}

func (q *Queries) GetTableConstraints(ctx context.Context, db DBTX, arg *GetTableConstraintsParams) ([]*GetTableConstraintsRow, error) {
	rows, err := db.Query(ctx, getTableConstraints, arg.Schema, arg.Table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetTableConstraintsRow
	for rows.Next() {
		var i GetTableConstraintsRow
		if err := rows.Scan(
			&i.DbSchema,
			&i.TableName,
			&i.ConstraintName,
			&i.ConstraintDefinition,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
