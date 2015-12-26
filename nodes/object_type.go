// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

/*
 * When a command can act on several kinds of objects with only one
 * parse structure required, use these constants to designate the
 * object type.  Note that commands typically don't support all the types.
 */
type ObjectType uint

const (
	OBJECT_AGGREGATE ObjectType = iota
	OBJECT_ATTRIBUTE            /* type's attribute, when distinct from column */
	OBJECT_CAST
	OBJECT_COLUMN
	OBJECT_CONSTRAINT
	OBJECT_COLLATION
	OBJECT_CONVERSION
	OBJECT_DATABASE
	OBJECT_DOMAIN
	OBJECT_EVENT_TRIGGER
	OBJECT_EXTENSION
	OBJECT_FDW
	OBJECT_FOREIGN_SERVER
	OBJECT_FOREIGN_TABLE
	OBJECT_FUNCTION
	OBJECT_INDEX
	OBJECT_LANGUAGE
	OBJECT_LARGEOBJECT
	OBJECT_MATVIEW
	OBJECT_OPCLASS
	OBJECT_OPERATOR
	OBJECT_OPFAMILY
	OBJECT_ROLE
	OBJECT_RULE
	OBJECT_SCHEMA
	OBJECT_SEQUENCE
	OBJECT_TABLE
	OBJECT_TABLESPACE
	OBJECT_TRIGGER
	OBJECT_TSCONFIGURATION
	OBJECT_TSDICTIONARY
	OBJECT_TSPARSER
	OBJECT_TSTEMPLATE
	OBJECT_TYPE
	OBJECT_VIEW
)