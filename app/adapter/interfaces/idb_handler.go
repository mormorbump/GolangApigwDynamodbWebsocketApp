package interfaces

type IDBHandler interface {
	DBPutItem(primaryKeyVal string, sortKeyVal string, attrKey1Val string) error
	DBDeleteItem(sortKeyVal string) error
	DBGetItem(primaryKeyVal string, sortKeyVal string, out interface{}) error
	DBQuery(primaryKeyVal string, out interface{}) error
}
