package repo

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-office-facade/internal/model"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

var testOffice = model.OfficePayload{
	ID:          uint64(1),
	Name:        "test",
	Description: "test",
}

type officeRepoFixture struct {
	officeRepo OfficeRepo
	db         *sqlx.DB
	dbMock     sqlmock.Sqlmock
}

func setUp(t *testing.T) officeRepoFixture {
	var fixture officeRepoFixture

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	fixture.db = sqlx.NewDb(db, "sqlmock")
	fixture.dbMock = mock
	fixture.officeRepo = NewOfficeRepo(fixture.db)

	return fixture
}

func (f *officeRepoFixture) tearDown() {
	f.db.Close()
}

func Test_repo_CreateOffice(t *testing.T) {
	f := setUp(t)
	defer f.tearDown()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	f.dbMock.ExpectQuery(`INSERT INTO offices (.+) VALUES (.+) RETURNING id`).
		WithArgs(testOffice.Name, testOffice.Description).
		WillReturnRows(rows)

	resultID, err := f.officeRepo.CreateOffice(context.Background(), testOffice)

	require.NoError(t, err)
	require.Equal(t, testOffice.ID, resultID)
}

func Test_repo_RemoveOffice(t *testing.T) {
	f := setUp(t)
	defer f.tearDown()

	expectSQL := regexp.QuoteMeta(`UPDATE offices SET removed = $1, updated_at = NOW() WHERE (id = $2 AND removed <> $3)`)

	f.dbMock.ExpectExec(expectSQL).
		WithArgs(true, testOffice.ID, true).WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := f.officeRepo.RemoveOffice(context.Background(), testOffice.ID)

	require.NoError(t, err)
	require.Equal(t, result, true)
}

func Test_repo_RemoveOffice_Not_Found(t *testing.T) {
	f := setUp(t)
	defer f.tearDown()

	expectSQL := regexp.QuoteMeta(`UPDATE offices SET removed = $1, updated_at = NOW() WHERE (id = $2 AND removed <> $3)`)

	f.dbMock.ExpectExec(expectSQL).
		WithArgs(true, testOffice.ID, true).WillReturnResult(sqlmock.NewResult(0, 0))

	result, err := f.officeRepo.RemoveOffice(context.Background(), testOffice.ID)

	require.ErrorIs(t, err, ErrOfficeNotFound)
	require.Equal(t, result, false)
}

func Test_repo_UpdateOffice(t *testing.T) {
	f := setUp(t)
	defer f.tearDown()

	expectSQL := regexp.QuoteMeta(`UPDATE offices SET name = $1, description = $2, updated_at = NOW() WHERE (id = $3 AND removed <> $4)`)

	f.dbMock.ExpectExec(expectSQL).
		WithArgs(testOffice.Name, testOffice.Description, testOffice.ID, true).WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := f.officeRepo.UpdateOffice(context.Background(), testOffice.ID, testOffice)

	require.NoError(t, err)
	require.Equal(t, result, true)
}

func Test_repo_UpdateOffice_Err_Not_Found(t *testing.T) {
	f := setUp(t)
	defer f.tearDown()

	expectSQL := regexp.QuoteMeta(`UPDATE offices SET name = $1, description = $2, updated_at = NOW() WHERE (id = $3 AND removed <> $4)`)

	f.dbMock.ExpectExec(expectSQL).
		WithArgs(testOffice.Name, testOffice.Description, testOffice.ID, true).WillReturnResult(sqlmock.NewResult(0, 0))

	result, err := f.officeRepo.UpdateOffice(context.Background(), testOffice.ID, testOffice)

	require.ErrorIs(t, err, ErrOfficeNotFound)
	require.Equal(t, result, false)
}
