package repo

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/ozonmp/bss-office-facade/internal/database"
	"github.com/ozonmp/bss-office-facade/internal/model"
	"github.com/pkg/errors"
)

// ErrOfficeNotFound - ошибка в методах, возникающая когда запрошенная сущность office не найдена
var ErrOfficeNotFound = errors.New("office not found")

const officesTableName = "offices"

const (
	officeIDColumn          = "id"
	officeNameColumn        = "name"
	officeDescriptionColumn = "description"
	officeRemovedColumn     = "removed"
	officeCreatedAtColumn   = "created_at"
	officeUpdatedAtColumn   = "updated_at"
)

// OfficeRepo is DAO for Office
type OfficeRepo interface {
	CreateOffice(ctx context.Context, office model.OfficePayload) (uint64, error)
	RemoveOffice(ctx context.Context, officeID uint64) (bool, error)
	UpdateOffice(ctx context.Context, officeID uint64, office model.OfficePayload) (bool, error)
}

type repo struct {
	db *sqlx.DB
}

// NewOfficeRepo returns OfficeRepo interface
func NewOfficeRepo(db *sqlx.DB) OfficeRepo {
	return &repo{db: db}
}

// CreateOffice - create new office
func (r *repo) CreateOffice(ctx context.Context, office model.OfficePayload) (uint64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OfficeRepo.CreateOffice")
	defer span.Finish()

	sb := database.StatementBuilder.
		Insert(officesTableName).
		Columns(officeNameColumn, officeDescriptionColumn).
		Values(office.Name, office.Description).
		Suffix("RETURNING " + officeIDColumn)

	query, args, err := sb.ToSql()

	if err != nil {
		return 0, errors.Wrap(err, "CreateOffice:ToSql()")
	}

	row := r.db.QueryRowxContext(ctx, query, args...)

	var id uint64
	err = row.Scan(&id)

	if err != nil {
		return 0, errors.Wrap(err, "CreateOffice:Scan()")
	}

	return id, nil
}

//RemoveOffice - remove office by id
// office is not really delete, just set the removed flag to true
func (r *repo) RemoveOffice(ctx context.Context, officeID uint64) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OfficeRepo.RemoveOffice")
	defer span.Finish()

	sb := database.StatementBuilder.
		Update(officesTableName).
		Set(officeRemovedColumn, true).
		Set(officeUpdatedAtColumn, sq.Expr("NOW()")).
		Where(sq.And{
			sq.Eq{officeIDColumn: officeID},
			sq.NotEq{officeRemovedColumn: true},
		})

	query, args, err := sb.ToSql()
	if err != nil {
		return false, errors.Wrap(err, "RemoveOffice:ToSql()")
	}

	res, err := r.db.ExecContext(ctx, query, args...)

	if err != nil {
		return false, errors.Wrap(err, "db.SelectContext()")
	}

	rowsCount, err := res.RowsAffected()

	if err != nil {
		return false, errors.Wrap(err, "db.RowsAffected")
	}

	if rowsCount == 0 {
		return false, ErrOfficeNotFound
	}

	return true, nil
}

//UpdateOffice - update all editable fields in office model by id
func (r *repo) UpdateOffice(ctx context.Context, officeID uint64, officeModel model.OfficePayload) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OfficeRepo.UpdateOffice")
	defer span.Finish()

	sb := database.StatementBuilder.
		Update(officesTableName).
		Set(officeNameColumn, officeModel.Name).
		Set(officeDescriptionColumn, officeModel.Description).
		Set(officeUpdatedAtColumn, sq.Expr("NOW()")).
		Where(sq.And{
			sq.Eq{officeIDColumn: officeID},
			sq.NotEq{officeRemovedColumn: true},
		})

	query, args, err := sb.ToSql()
	if err != nil {
		return false, errors.Wrap(err, "UpdateOffice:ToSql()")
	}

	res, err := r.db.ExecContext(ctx, query, args...)

	if err != nil {
		return false, errors.Wrap(err, "db.SelectContext()")
	}

	rowsCount, err := res.RowsAffected()

	if err != nil {
		return false, errors.Wrap(err, "db.RowsAffected")
	}

	if rowsCount == 0 {
		return false, ErrOfficeNotFound
	}

	return true, nil
}
