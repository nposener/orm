// Autogenerated by github.com/posener/orm
package allorm

func (c *TCreate) String() string {
	// Create statement has a line for each variable with it's name and it's type.
	return `CREATE TABLE 'all' ( 'int' INT PRIMARY KEY, 'string' VARCHAR(100) NOT NULL, 'bool' BOOLEAN, 'auto' INT AUTO_INCREMENT, 'time' TIMESTAMP, 'select' INT )`
}
