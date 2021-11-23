package repo

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-office-facade/internal/model"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
	"time"
)

type eventRepoFixture struct {
	eventRepo EventRepo
	db        *sqlx.DB
	dbMock    sqlmock.Sqlmock
}

var testEventModel = model.OfficeEvent{
	OfficeID: 1,
	Type:     model.Created,
	Status:   model.Deferred,
	Created:  time.Now(),
}

func setUpEventRepo(t *testing.T) eventRepoFixture {
	var fixture eventRepoFixture

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	fixture.db = sqlx.NewDb(db, "sqlmock")
	fixture.dbMock = mock
	fixture.eventRepo = NewEventRepo(fixture.db)

	return fixture
}

func (f *eventRepoFixture) tearDown() {
	f.db.Close()
}

func Test_eventRepo_Lock(t *testing.T) {
	f := setUpEventRepo(t)
	defer f.tearDown()

	rows := sqlmock.NewRows([]string{"id", "office_id", "type", "status", "created_at", "payload"}).
		AddRow(1, 1, model.Created, model.Processed, time.Now(), "{}").
		AddRow(2, 2, model.Updated, model.Processed, time.Now(), "{}").
		AddRow(3, 3, model.Removed, model.Processed, time.Now(), "{}")

	testLimit := uint64(3)

	expectSQL := regexp.QuoteMeta(`
UPDATE offices_events
SET status = $1, updated_at = NOW()
WHERE id IN ( 
SELECT id FROM offices_events 
WHERE status = $2 
ORDER BY id 
LIMIT 3 
FOR UPDATE SKIP LOCKED ) RETURNING *`)

	f.dbMock.ExpectQuery(expectSQL).
		WithArgs(model.Processed, model.Deferred).
		WillReturnRows(rows)

	result, err := f.eventRepo.Lock(context.Background(), testLimit)

	require.NoError(t, err)
	require.Equal(t, testLimit, uint64(len(result)))
}

func Test_eventRepo_Remove(t *testing.T) {
	f := setUpEventRepo(t)
	defer f.tearDown()

	f.dbMock.ExpectExec(`DELETE FROM offices_events WHERE id IN (.+)`).
		WithArgs(1, 2).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := f.eventRepo.Remove(context.Background(), []uint64{1, 2})

	require.NoError(t, err)
}

func Test_eventRepo_Remove_Err_Not_Found(t *testing.T) {
	f := setUpEventRepo(t)
	defer f.tearDown()

	f.dbMock.ExpectExec(`DELETE FROM offices_events WHERE id IN (.+)`).
		WithArgs(1, 2).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := f.eventRepo.Remove(context.Background(), []uint64{1, 2})

	require.ErrorIs(t, err, ErrOfficeNotFound)
}

func Test_eventRepo_Unlock(t *testing.T) {
	f := setUpEventRepo(t)
	defer f.tearDown()

	f.dbMock.ExpectExec(`UPDATE offices_events SET status = \$1 WHERE id IN (.+)`).
		WithArgs(model.Deferred, 1, 2).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := f.eventRepo.Unlock(context.Background(), []uint64{1, 2})

	require.NoError(t, err)
}

func Test_eventRepo_Unlock_Err_Not_Found(t *testing.T) {
	f := setUpEventRepo(t)
	defer f.tearDown()

	f.dbMock.ExpectExec(`UPDATE offices_events SET status = \$1 WHERE id IN (.+)`).
		WithArgs(model.Deferred, 1, 2).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := f.eventRepo.Unlock(context.Background(), []uint64{1, 2})

	require.ErrorIs(t, err, ErrNoneRowsUnlock)
}

func Test_eventRepo_Add(t *testing.T) {
	f := setUpEventRepo(t)
	defer f.tearDown()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	expectPayload, err := convertBssOfficeToJsonb(&testEventModel.Payload)

	require.NoError(t, err)

	f.dbMock.ExpectQuery(`INSERT INTO offices_events  (.+) VALUES  (.+) RETURNING id`).
		WithArgs(testEventModel.OfficeID, testEventModel.Type, testEventModel.Status, expectPayload).
		WillReturnRows(rows)

	err = f.eventRepo.Add(context.Background(), &testEventModel)

	require.NoError(t, err)
}
