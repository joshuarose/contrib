// Code generated by ent, DO NOT EDIT.

package twomethodservice

const (
	// Label holds the string label denoting the twomethodservice type in the database.
	Label = "two_method_service"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the twomethodservice in the database.
	Table = "two_method_services"
)

// Columns holds all SQL columns for twomethodservice fields.
var Columns = []string{
	FieldID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
