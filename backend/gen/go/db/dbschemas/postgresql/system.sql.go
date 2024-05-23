// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: system.sql

package pg_queries

import (
	"context"
)

const getDatabaseSchema = `-- name: GetDatabaseSchema :many
SELECT
    n.nspname AS table_schema,
    c.relname AS table_name,
    a.attname AS column_name,
    pg_catalog.format_type(a.atttypid, a.atttypmod) AS data_type, -- This formats the type into something that should always be a valid postgres type. It also includes constraints if there are any
 COALESCE(
        pg_catalog.pg_get_expr(d.adbin, d.adrelid),
        ''
    )::text AS column_default,
    CASE
        WHEN a.attnotnull THEN 'NO'
        ELSE 'YES'
    END AS is_nullable,
    CASE
        WHEN pg_catalog.format_type(a.atttypid, a.atttypmod) LIKE 'character varying%' THEN
            a.atttypmod - 4 -- The -4 removes the extra bits that postgres uses for internal use
        ELSE
            -1
    END AS character_maximum_length,
    CASE
        WHEN a.atttypid = pg_catalog.regtype 'numeric'::regtype THEN
            (a.atttypmod - 4) >> 16
        -- Precision is technically only necessary for numeric values, but we are populating these here for simplicity in knowing what the type of integer is.
        -- This operates similar to the precision column in the information_schema.columns table
        WHEN a.atttypid = pg_catalog.regtype 'smallint'::regtype THEN
            16
        WHEN a.atttypid = pg_catalog.regtype 'integer'::regtype THEN
            32
        WHEN a.atttypid = pg_catalog.regtype 'bigint'::regtype THEN
            64
        ELSE
            -1
    END AS numeric_precision,
    CASE
        WHEN a.atttypid = pg_catalog.regtype 'numeric'::regtype THEN
            CASE
                -- If scale is not explicitly set, return -1 (meaning arbitrary scale)
                WHEN (a.atttypmod) = -1 THEN -1
                ELSE (a.atttypmod - 4) & 65535
            END
        -- Scale is technically only necessary for numeric values, but we are populating these here for simplicity in knowing what the type of integer is.
        -- This operates similar to the scake column in the information_schema.columns table
        WHEN a.atttypid = pg_catalog.regtype 'smallint'::regtype THEN
            0
        WHEN a.atttypid = pg_catalog.regtype 'integer'::regtype THEN
            0
        WHEN a.atttypid = pg_catalog.regtype 'bigint'::regtype THEN
            0
        ELSE
            -1
    END AS numeric_scale,
    a.attnum AS ordinal_position
    -- a.attgenerated::text as generated_type
FROM
    pg_catalog.pg_attribute a
    INNER JOIN pg_catalog.pg_class c ON a.attrelid = c.oid
    INNER JOIN pg_catalog.pg_namespace n ON c.relnamespace = n.oid
    INNER JOIN pg_catalog.pg_type pgt ON pgt.oid = a.atttypid
    LEFT JOIN pg_catalog.pg_attrdef d ON d.adrelid = a.attrelid AND d.adnum = a.attnum
WHERE
    n.nspname NOT IN('pg_catalog', 'pg_toast', 'information_schema')
    AND a.attnum > 0
    AND NOT a.attisdropped
    AND c.relkind = 'r' -- ensures only tables are present
ORDER BY
    a.attnum
`

type GetDatabaseSchemaRow struct {
	TableSchema            string
	TableName              string
	ColumnName             string
	DataType               string
	ColumnDefault          string
	IsNullable             string
	CharacterMaximumLength int32
	NumericPrecision       int32
	NumericScale           int32
	OrdinalPosition        int16
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
			&i.DataType,
			&i.ColumnDefault,
			&i.IsNullable,
			&i.CharacterMaximumLength,
			&i.NumericPrecision,
			&i.NumericScale,
			&i.OrdinalPosition,
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
    n.nspname AS schema_name,
    c.relname AS table_name,
    a.attname AS column_name,
    pg_catalog.format_type(a.atttypid, a.atttypmod) AS data_type,  -- This formats the type into something that should always be a valid postgres type. It also includes constraints if there are any
 COALESCE(
        pg_catalog.pg_get_expr(d.adbin, d.adrelid),
        ''
    )::text AS column_default,
    CASE
        WHEN a.attnotnull THEN 'NO'
        ELSE 'YES'
    END AS is_nullable,
    CASE
        WHEN pg_catalog.format_type(a.atttypid, a.atttypmod) LIKE 'character varying%' THEN
            a.atttypmod - 4 -- The -4 removes the extra bits that postgres uses for internal use
        ELSE
            -1
    END AS character_maximum_length,
    CASE
        WHEN a.atttypid = pg_catalog.regtype 'numeric'::regtype THEN
            (a.atttypmod - 4) >> 16
        -- Precision is technically only necessary for numeric values, but we are populating these here for simplicity in knowing what the type of integer is.
        -- This operates similar to the precision column in the information_schema.columns table
        WHEN a.atttypid = pg_catalog.regtype 'smallint'::regtype THEN
            16
        WHEN a.atttypid = pg_catalog.regtype 'integer'::regtype THEN
            32
        WHEN a.atttypid = pg_catalog.regtype 'bigint'::regtype THEN
            64
        ELSE
            -1
    END AS numeric_precision,
    CASE
        WHEN a.atttypid = pg_catalog.regtype 'numeric'::regtype THEN
            CASE
                -- If scale is not explicitly set, return -1 (meaning arbitrary scale)
                WHEN (a.atttypmod) = -1 THEN -1
                ELSE (a.atttypmod - 4) & 65535
            END
        -- Scale is technically only necessary for numeric values, but we are populating these here for simplicity in knowing what the type of integer is.
        -- This operates similar to the scake column in the information_schema.columns table
        WHEN a.atttypid = pg_catalog.regtype 'smallint'::regtype THEN
            0
        WHEN a.atttypid = pg_catalog.regtype 'integer'::regtype THEN
            0
        WHEN a.atttypid = pg_catalog.regtype 'bigint'::regtype THEN
            0
        ELSE
            -1
    END AS numeric_scale,
    a.attnum AS ordinal_position
    -- a.attgenerated::text as generated_type
FROM
    pg_catalog.pg_attribute a
    INNER JOIN pg_catalog.pg_class c ON a.attrelid = c.oid
    INNER JOIN pg_catalog.pg_namespace n ON c.relnamespace = n.oid
    INNER JOIN pg_catalog.pg_type pgt ON pgt.oid = a.atttypid
    LEFT JOIN pg_catalog.pg_attrdef d ON d.adrelid = a.attrelid AND d.adnum = a.attnum
WHERE
    c.relname = $1
    AND n.nspname = $2
    AND a.attnum > 0
    AND NOT a.attisdropped
    AND c.relkind = 'r' -- ensures only tables are present
ORDER BY
    a.attnum
`

type GetDatabaseTableSchemaParams struct {
	Table  string
	Schema string
}

type GetDatabaseTableSchemaRow struct {
	SchemaName             string
	TableName              string
	ColumnName             string
	DataType               string
	ColumnDefault          string
	IsNullable             string
	CharacterMaximumLength int32
	NumericPrecision       int32
	NumericScale           int32
	OrdinalPosition        int16
}

func (q *Queries) GetDatabaseTableSchema(ctx context.Context, db DBTX, arg *GetDatabaseTableSchemaParams) ([]*GetDatabaseTableSchemaRow, error) {
	rows, err := db.Query(ctx, getDatabaseTableSchema, arg.Table, arg.Schema)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetDatabaseTableSchemaRow
	for rows.Next() {
		var i GetDatabaseTableSchemaRow
		if err := rows.Scan(
			&i.SchemaName,
			&i.TableName,
			&i.ColumnName,
			&i.DataType,
			&i.ColumnDefault,
			&i.IsNullable,
			&i.CharacterMaximumLength,
			&i.NumericPrecision,
			&i.NumericScale,
			&i.OrdinalPosition,
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

const getPostgresRolePermissions = `-- name: GetPostgresRolePermissions :many
SELECT
    rtg.table_schema as table_schema,
    rtg.table_name as table_name,
    rtg.privilege_type as privilege_type
FROM
    information_schema.role_table_grants as rtg
WHERE
    table_schema NOT IN ('pg_catalog', 'information_schema')
AND grantee =  $1
ORDER BY
    table_schema,
    table_name
`

type GetPostgresRolePermissionsRow struct {
	TableSchema   string
	TableName     string
	PrivilegeType string
}

func (q *Queries) GetPostgresRolePermissions(ctx context.Context, db DBTX, role interface{}) ([]*GetPostgresRolePermissionsRow, error) {
	rows, err := db.Query(ctx, getPostgresRolePermissions, role)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetPostgresRolePermissionsRow
	for rows.Next() {
		var i GetPostgresRolePermissionsRow
		if err := rows.Scan(&i.TableSchema, &i.TableName, &i.PrivilegeType); err != nil {
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
    con.conname AS constraint_name,
    con.contype::TEXT AS constraint_type,
    con.connamespace::regnamespace::TEXT AS schema_name,
    cls.relname AS table_name,
    array_agg(att.attname)::TEXT[] AS constraint_columns,
    array_agg(att.attnotnull)::BOOL[] AS notnullable,
    CASE
        WHEN con.contype = 'f' THEN fn_cl.relnamespace::regnamespace::TEXT
        ELSE ''
    END AS foreign_schema_name,
    CASE
        WHEN con.contype = 'f' THEN fn_cl.relname
        ELSE ''
    END AS foreign_table_name,
    CASE
        WHEN con.contype = 'f' THEN fk_columns.foreign_column_names
        ELSE NULL::text[]
    END as foreign_column_names,
    pg_get_constraintdef(con.oid)::TEXT AS constraint_definition
FROM
    pg_catalog.pg_constraint con
JOIN
    pg_catalog.pg_attribute att ON
    att.attrelid = con.conrelid
    AND att.attnum = ANY(con.conkey)
JOIN
    pg_catalog.pg_class cls ON
    con.conrelid = cls.oid
JOIN
    pg_catalog.pg_namespace nsp ON
    cls.relnamespace = nsp.oid
LEFT JOIN
    pg_catalog.pg_class fn_cl ON
    fn_cl.oid = con.confrelid
LEFT JOIN LATERAL (
        SELECT
            array_agg(fk_att.attname) AS foreign_column_names
        FROM
            pg_catalog.pg_attribute fk_att
        WHERE
            fk_att.attrelid = con.confrelid
            AND fk_att.attnum = ANY(con.confkey)
    ) AS fk_columns ON
    TRUE
WHERE
    con.connamespace::regnamespace::TEXT = $1
    AND con.conrelid::regclass::TEXT = $2
GROUP BY
    con.oid,
    con.connamespace,
    con.conname,
    cls.relname,
    con.contype,
    fn_cl.relnamespace,
    fn_cl.relname,
    fk_columns.foreign_column_names
`

type GetTableConstraintsParams struct {
	Schema string
	Table  string
}

type GetTableConstraintsRow struct {
	ConstraintName       string
	ConstraintType       string
	SchemaName           string
	TableName            string
	ConstraintColumns    []string
	Notnullable          []bool
	ForeignSchemaName    string
	ForeignTableName     string
	ForeignColumnNames   []string
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
			&i.ConstraintName,
			&i.ConstraintType,
			&i.SchemaName,
			&i.TableName,
			&i.ConstraintColumns,
			&i.Notnullable,
			&i.ForeignSchemaName,
			&i.ForeignTableName,
			&i.ForeignColumnNames,
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

const getTableConstraintsBySchema = `-- name: GetTableConstraintsBySchema :many
SELECT
    con.conname AS constraint_name,
    con.contype::TEXT AS constraint_type,
    con.connamespace::regnamespace::TEXT AS schema_name,
    cls.relname AS table_name,
    array_agg(att.attname)::TEXT[] AS constraint_columns,
    array_agg(att.attnotnull)::BOOL[] AS notnullable,
    CASE
        WHEN con.contype = 'f' THEN fn_cl.relnamespace::regnamespace::TEXT
        ELSE ''
    END AS foreign_schema_name,
    CASE
        WHEN con.contype = 'f' THEN fn_cl.relname
        ELSE ''
    END AS foreign_table_name,
    CASE
        WHEN con.contype = 'f' THEN fk_columns.foreign_column_names
        ELSE NULL::text[]
    END as foreign_column_names,
    pg_get_constraintdef(con.oid)::TEXT AS constraint_definition
FROM
    pg_catalog.pg_constraint con
JOIN
    pg_catalog.pg_attribute att ON
    att.attrelid = con.conrelid
    AND att.attnum = ANY(con.conkey)
JOIN
    pg_catalog.pg_class cls ON
    con.conrelid = cls.oid
JOIN
    pg_catalog.pg_namespace nsp ON
    cls.relnamespace = nsp.oid
LEFT JOIN
    pg_catalog.pg_class fn_cl ON
    fn_cl.oid = con.confrelid
LEFT JOIN LATERAL (
        SELECT
            array_agg(fk_att.attname) AS foreign_column_names
        FROM
            pg_catalog.pg_attribute fk_att
        WHERE
            fk_att.attrelid = con.confrelid
            AND fk_att.attnum = ANY(con.confkey)
    ) AS fk_columns ON
    TRUE
WHERE
    con.connamespace::regnamespace::TEXT = ANY(
        $1::TEXT[]
    )
GROUP BY
    con.oid,
    con.connamespace,
    con.conname,
    cls.relname,
    con.contype,
    fn_cl.relnamespace,
    fn_cl.relname,
    fk_columns.foreign_column_names
`

type GetTableConstraintsBySchemaRow struct {
	ConstraintName       string
	ConstraintType       string
	SchemaName           string
	TableName            string
	ConstraintColumns    []string
	Notnullable          []bool
	ForeignSchemaName    string
	ForeignTableName     string
	ForeignColumnNames   []string
	ConstraintDefinition string
}

func (q *Queries) GetTableConstraintsBySchema(ctx context.Context, db DBTX, schema []string) ([]*GetTableConstraintsBySchemaRow, error) {
	rows, err := db.Query(ctx, getTableConstraintsBySchema, schema)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetTableConstraintsBySchemaRow
	for rows.Next() {
		var i GetTableConstraintsBySchemaRow
		if err := rows.Scan(
			&i.ConstraintName,
			&i.ConstraintType,
			&i.SchemaName,
			&i.TableName,
			&i.ConstraintColumns,
			&i.Notnullable,
			&i.ForeignSchemaName,
			&i.ForeignTableName,
			&i.ForeignColumnNames,
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
