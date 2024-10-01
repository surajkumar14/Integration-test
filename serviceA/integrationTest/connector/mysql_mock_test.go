package connector

import (
	integrationtest "serviceA/integrationTest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSQLMockCRUDOperations(t *testing.T) {
	t.Run("CreateRecord", func(t *testing.T) {
		integrationtest.SQLMock.ExpectExec("INSERT INTO users").
			WithArgs(1, "John Doe").
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := CreateRecord(integrationtest.SQLMockDbClient, 1, "John Doe")
		assert.NoError(t, err)
		assert.NoError(t, integrationtest.SQLMock.ExpectationsWereMet())
	})
	t.Run("ReadRecord", func(t *testing.T) {

		rows := sqlmock.NewRows([]string{"name"}).AddRow("John Doe")
		integrationtest.SQLMock.ExpectQuery("SELECT name FROM users WHERE id = ?").
			WithArgs(1).
			WillReturnRows(rows)

		name, err := ReadRecord(integrationtest.SQLMockDbClient, 1)
		assert.NoError(t, err)
		assert.Equal(t, "John Doe", name)
		assert.NoError(t, integrationtest.SQLMock.ExpectationsWereMet())
	})
	t.Run("UpdateRecord", func(t *testing.T) {

		integrationtest.SQLMock.ExpectExec("UPDATE users SET name = \\? WHERE id = \\?").
			WithArgs("Jane Doe", 1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := UpdateRecord(integrationtest.SQLMockDbClient, 1, "Jane Doe")
		assert.NoError(t, err)
		assert.NoError(t, integrationtest.SQLMock.ExpectationsWereMet())
	})
	t.Run("DeleteRecord", func(t *testing.T) {

		integrationtest.SQLMock.ExpectExec("DELETE FROM users WHERE id = \\?").
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := DeleteRecord(integrationtest.SQLMockDbClient, 1)
		assert.NoError(t, err)
		assert.NoError(t, integrationtest.SQLMock.ExpectationsWereMet())
	})
}
