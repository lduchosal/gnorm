// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

// DataTypePrivilegeTable is the database name for the table.
const DataTypePrivilegeTable = "information_schema.data_type_privileges"

// DataTypePrivilege represents a row from 'information_schema.data_type_privileges'.
type DataTypePrivilege struct {
	ObjectCatalog SQLIdentifier `json:"object_catalog"` // object_catalog
	ObjectSchema  SQLIdentifier `json:"object_schema"`  // object_schema
	ObjectName    SQLIdentifier `json:"object_name"`    // object_name
	ObjectType    CharacterData `json:"object_type"`    // object_type
	DtdIdentifier SQLIdentifier `json:"dtd_identifier"` // dtd_identifier
}

// Constants defining each column in the table.
const (
	DataTypePrivilegeObjectCatalogField = "object_catalog"
	DataTypePrivilegeObjectSchemaField  = "object_schema"
	DataTypePrivilegeObjectNameField    = "object_name"
	DataTypePrivilegeObjectTypeField    = "object_type"
	DataTypePrivilegeDtdIdentifierField = "dtd_identifier"
)

// WhereClauses for every type in DataTypePrivilege.
var (
	DataTypePrivilegeObjectCatalogWhere SQLIdentifierField = "object_catalog"
	DataTypePrivilegeObjectSchemaWhere  SQLIdentifierField = "object_schema"
	DataTypePrivilegeObjectNameWhere    SQLIdentifierField = "object_name"
	DataTypePrivilegeObjectTypeWhere    CharacterDataField = "object_type"
	DataTypePrivilegeDtdIdentifierWhere SQLIdentifierField = "dtd_identifier"
)

// QueryOneDataTypePrivilege retrieves a row from 'information_schema.data_type_privileges' as a DataTypePrivilege.
func QueryOneDataTypePrivilege(db XODB, where WhereClause, order OrderBy) (*DataTypePrivilege, error) {
	const origsqlstr = `SELECT ` +
		`object_catalog, object_schema, object_name, object_type, dtd_identifier ` +
		`FROM information_schema.data_type_privileges WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	dtp := &DataTypePrivilege{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&dtp.ObjectCatalog, &dtp.ObjectSchema, &dtp.ObjectName, &dtp.ObjectType, &dtp.DtdIdentifier)
	if err != nil {
		return nil, err
	}
	return dtp, nil
}

// QueryDataTypePrivilege retrieves rows from 'information_schema.data_type_privileges' as a slice of DataTypePrivilege.
func QueryDataTypePrivilege(db XODB, where WhereClause, order OrderBy) ([]*DataTypePrivilege, error) {
	const origsqlstr = `SELECT ` +
		`object_catalog, object_schema, object_name, object_type, dtd_identifier ` +
		`FROM information_schema.data_type_privileges WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*DataTypePrivilege
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, err
	}
	for q.Next() {
		dtp := DataTypePrivilege{}

		err = q.Scan(&dtp.ObjectCatalog, &dtp.ObjectSchema, &dtp.ObjectName, &dtp.ObjectType, &dtp.DtdIdentifier)
		if err != nil {
			return nil, err
		}

		vals = append(vals, &dtp)
	}
	return vals, nil
}