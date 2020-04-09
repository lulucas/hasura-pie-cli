package model

const (
	ColumnsQuery = `
	SELECT
		a.attname as col,
		case
			when left(t.typname, 1) = '_' then right(t.typname, length(t.typname)-1)
			else t.typname
		end,
		(	SELECT
				substring(pg_catalog.pg_get_expr(d.adbin, d.adrelid) for 128) as default
			FROM pg_catalog.pg_attrdef d
			WHERE d.adrelid = a.attrelid
				and d.adnum = a.attnum
				and a.atthasdef
		),
		not a.attnotnull,
		t.typcategory = 'A' as is_array,
		t.typcategory
		FROM pg_catalog.pg_attribute a
			join pg_type t on t.oid = a.atttypid
		WHERE a.attrelid = $1::regclass
			AND a.attnum > 0
			AND not a.attisdropped
		ORDER BY a.attnum;`
	TablesQuery = `
	SELECT table_name
		FROM information_schema.tables
		WHERE table_schema='public' AND table_type='BASE TABLE';`
)
