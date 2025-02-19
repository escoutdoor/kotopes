package user

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/escoutdoor/kotopes/auth/internal/model"
	"github.com/escoutdoor/kotopes/auth/internal/repository/user/converter"
	repomodel "github.com/escoutdoor/kotopes/auth/internal/repository/user/model"
	"github.com/escoutdoor/kotopes/common/pkg/db"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

var (
	tableName = "users"

	idColumn        = "id"
	firstNameColumn = "first_name"
	lastNameColumn  = "last_name"
	emailColumn     = "email"
	passwordColumn  = "password"
	roleColumn      = "role"
	phoneColumn     = "phone"
	cityColumn      = "city"
	countryColumn   = "country"
	createdAtColumn = "created_at"
)

type repository struct {
	db db.Client
}

func New(db db.Client) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, in *model.CreateUser) (string, error) {
	const op = "user_repository.Create"

	query, args, err := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(
			firstNameColumn,
			lastNameColumn,
			emailColumn,
			passwordColumn,
		).Values(in.FirstName, in.LastName, in.Email, in.Password).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return "", errwrap.Wrap(op, err)
	}

	q := db.Query{
		Name:     op,
		QueryRow: query,
	}

	var id string
	err = r.db.DB().QueryRow(ctx, q, args...).Scan(&id)
	if err != nil {
		return "", errwrap.Wrap(op, err)
	}

	return id, nil
}

func (r *repository) List(ctx context.Context, userIDs []string) ([]*model.User, error) {
	const op = "user_repository.List"

	builder := sq.Select(
		idColumn,
		firstNameColumn,
		lastNameColumn,
		emailColumn,
		passwordColumn,
		roleColumn,
		phoneColumn,
		cityColumn,
		countryColumn,
		createdAtColumn,
	).PlaceholderFormat(sq.Dollar).
		From(tableName)

	if len(userIDs) > 0 {
		builder = builder.Where(sq.Eq{idColumn: userIDs})
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	q := db.Query{
		Name:     op,
		QueryRow: query,
	}

	rows, err := r.db.DB().Query(ctx, q, args...)
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}
	defer rows.Close()

	var users []*repomodel.User
	err = pgxscan.ScanAll(&users, rows)
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	return converter.ToUsersFromRepo(users), nil
}

func (r *repository) GetByID(ctx context.Context, id string) (*model.User, error) {
	const op = "user_repository.GetByID"

	query, args, err := sq.Select(
		idColumn,
		firstNameColumn,
		lastNameColumn,
		emailColumn,
		passwordColumn,
		roleColumn,
		phoneColumn,
		cityColumn,
		countryColumn,
		createdAtColumn,
	).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1).
		ToSql()
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	q := db.Query{
		Name:     op,
		QueryRow: query,
	}

	row, err := r.db.DB().Query(ctx, q, args...)
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}
	defer row.Close()

	var user repomodel.User
	err = pgxscan.ScanOne(&user, row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errwrap.Wrap(op, model.ErrUserNotFound)
		}
		return nil, errwrap.Wrap(op, err)
	}

	return converter.ToUserFromRepo(&user), nil
}

func (r *repository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	const op = "user_repository.GetByEmail"

	query, args, err := sq.Select(
		idColumn,
		firstNameColumn,
		lastNameColumn,
		emailColumn,
		passwordColumn,
		roleColumn,
		phoneColumn,
		cityColumn,
		countryColumn,
		createdAtColumn,
	).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{emailColumn: email}).
		Limit(1).
		ToSql()
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	q := db.Query{
		Name:     op,
		QueryRow: query,
	}

	row, err := r.db.DB().Query(ctx, q, args...)
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}
	defer row.Close()

	var user repomodel.User
	err = pgxscan.ScanOne(&user, row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errwrap.Wrap(op, model.ErrUserNotFound)
		}
		return nil, errwrap.Wrap(op, err)
	}

	return converter.ToUserFromRepo(&user), nil
}

func (r *repository) Update(ctx context.Context, in *model.UpdateUser) error {
	const op = "user_repository.Update"

	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar)

	if in.FirstName != nil {
		builder = builder.Set(firstNameColumn, in.FirstName)
	}
	if in.LastName != nil {
		builder = builder.Set(lastNameColumn, in.LastName)
	}
	if in.Email != nil {
		builder = builder.Set(emailColumn, in.Email)
	}
	if in.Password != nil {
		builder = builder.Set(passwordColumn, in.Password)
	}
	if in.Phone != nil {
		builder = builder.Set(phoneColumn, in.Phone)
	}
	if in.City != nil {
		builder = builder.Set(cityColumn, in.City)
	}
	if in.Country != nil {
		builder = builder.Set(countryColumn, in.Country)
	}

	query, args, err := builder.Where(sq.Eq{idColumn: in.ID}).ToSql()
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	q := db.Query{
		Name:     op,
		QueryRow: query,
	}

	cmd, err := r.db.DB().Exec(ctx, q, args...)
	if err != nil {
		return errwrap.Wrap(op, err)
	}
	if cmd.RowsAffected() == 0 {
		return errwrap.Wrap(op, fmt.Errorf("no rows affected"))
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	const op = "user_repository.Delete"

	query, args, err := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id}).
		ToSql()
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	q := db.Query{
		Name:     op,
		QueryRow: query,
	}

	cmd, err := r.db.DB().Exec(ctx, q, args...)
	if err != nil {
		return errwrap.Wrap(op, err)
	}
	if cmd.RowsAffected() == 0 {
		return errwrap.Wrap(op, fmt.Errorf("no rows affected"))
	}

	return nil
}
