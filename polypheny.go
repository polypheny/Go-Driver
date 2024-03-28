package polypheny

import (
	sql "database/sql"
	driver "database/sql/driver"
	"strings"

	"golang.org/x/net/context"
)

type Driver struct{}

type Statement struct {
	conn     *Conn
	id       int32
	numInput int
}

func (stmt *Statement) Close() error {
	stmt.conn.conn.handleCloseStatementRequest(stmt.id)
	return nil
}

func (stmt *Statement) NumInput() int {
	return stmt.numInput
}

func (stmt *Statement) Exec(args []driver.Value) (driver.Result, error) {
	parameters := make([]interface{}, len(args))
	for i, v := range args {
		parameters[i] = v
	}
	request := IndexedParametersRequest{
		parameters: parameters,
	}
	stmt.conn.conn.handleExecuteIndexedStatementRequest(stmt.id, request, nil)
	return nil, nil
}

func (stmt *Statement) Query(args []driver.Value) (driver.Rows, error) {
	parameters := make([]interface{}, len(args))
	for i, v := range args {
		parameters[i] = v
	}
	request := IndexedParametersRequest{
		parameters: parameters,
	}
	stmt.conn.conn.handleExecuteIndexedStatementRequest(stmt.id, request, nil)
	return nil, nil
}

type Connector struct {
	address  string
	username string
	password string // TODO: Consider to use C string which can be earise to erase for security
}

type Conn struct {
	conn prismClient
}

type Rows struct {
	columns   []string
	result    [][]interface{}
	readIndex int
}

func (rows *Rows) Columns() []string {
	return rows.columns
}

func (rows *Rows) Close() error {
	return nil
}

func (rows *Rows) Next(dest []driver.Value) error {
	for i, _ := range dest {
		dest[i] = rows.result[rows.readIndex][i]
	}
	rows.readIndex++
	return nil
}

func (c *Conn) Prepare(query string) (driver.Stmt, error) {
	response := c.conn.handlePrepareIndexedStatementRequest(strings.Split(query, ":")[0], strings.Split(query, ":")[1], nil)
	var key int32
	var value []ParameterMetaResponse
	for k, v := range response {
		key = k
		value = v
	}
	stmt := Statement{
		conn:     c,
		id:       key,
		numInput: len(value),
	}

	return &stmt, nil
}

func (c *Conn) Close() error {
	c.conn.handleDisconnectRequest()
	return nil
}

func (c *Conn) Begin() (driver.Tx, error) {
	return nil, nil
}

func (c *Conn) QueryContext(ctx context.Context, query string, args []interface{}) (driver.Rows, error) {
	// TODO: support args
	request := UnparameterizedStatementRequest{
		language:      strings.Split(query, ":")[0],
		statement:     strings.Split(query, ":")[1],
		fetchSize:     nil,
		namespaceName: nil,
	}
	columns, result := c.conn.handleExecuteUnparameterizedStatementRequest(request)
	rows := Rows{
		columns:   columns,
		result:    result,
		readIndex: 0,
	}
	return &rows, nil
}

func init() {
	sql.Register("polypheny", Driver{})
}

func (c *Connector) Driver() driver.Driver {
	return Driver{}
}

func (c *Connector) Connect(context.Context) (driver.Conn, error) {
	client := handleConnectRequest(c.address, c.username, c.password)
	conn := Conn{
		conn: *client,
	}
	return &conn, nil
}

func (d Driver) Open(dsn string) (driver.Conn, error) {
	// dsn format: addr:port,user:password
	connectionStrings := strings.Split(dsn, ",")
	addr := connectionStrings[0]
	username := strings.Split(connectionStrings[1], ":")[0]
	password := strings.Split(connectionStrings[1], ":")[1]
	connector := Connector{
		address:  addr,
		username: username,
		password: password,
	}
	return connector.Connect(context.Background())
}
