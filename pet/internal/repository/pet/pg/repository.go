package pg

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/escoutdoor/kotopes/common/pkg/db"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
	"github.com/escoutdoor/kotopes/pet/internal/model"
	"github.com/escoutdoor/kotopes/pet/internal/repository/pet/pg/converter"
	repomodel "github.com/escoutdoor/kotopes/pet/internal/repository/pet/pg/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

var (
	tableName = "pets"

	idColumn          = "id"
	nameColumn        = "name"
	descriptionColumn = "description"
	ageColumn         = "age"
	ownerIDColumn     = "owner_id"
	createdAtColumn   = "created_at"
)

type repository struct {
	db db.Client
}

func New(db db.Client) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, in *model.CreatePet) (string, error) {
	const op = "pet_repository.Create"

	query, args, err := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, descriptionColumn, ageColumn, ownerIDColumn).
		Values(in.Name, in.Description, in.Age, in.OwnerID).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return "", errwrap.Wrap(op, err)
	}

	q := db.Query{
		Name:     "pet_repository.Create",
		QueryRow: query,
	}

	var id string
	err = r.db.DB().QueryRow(ctx, q, args...).Scan(&id)
	if err != nil {
		return "", errwrap.Wrap(op, err)
	}

	return id, nil
}

func (r *repository) GetByID(ctx context.Context, id string) (*model.Pet, error) {
	const op = "pet_repository.GetByID"

	query, args, err := sq.Select(idColumn, nameColumn, descriptionColumn,
		ageColumn, ownerIDColumn, createdAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1).
		ToSql()
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	q := db.Query{
		Name:     "pet_repository.GetByID",
		QueryRow: query,
	}

	row, err := r.db.DB().Query(ctx, q, args...)
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}
	defer row.Close()

	var pet repomodel.Pet
	err = pgxscan.ScanOne(&pet, row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errwrap.Wrap(op, model.ErrPetNotFound)
		}
		return nil, errwrap.Wrap(op, err)
	}

	return converter.ToPetFromRepo(&pet), nil
}

func (r *repository) ListPets(ctx context.Context, in *model.ListPets) ([]*model.Pet, error) {
	const op = "pet_repository.ListPets"

	builder := sq.Select(idColumn, nameColumn, descriptionColumn,
		ageColumn, ownerIDColumn, createdAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName)

	if in.Limit > 0 {
		builder = builder.Limit(uint64(in.Limit))
	}
	if in.Offset > 0 {
		builder = builder.Offset(uint64(in.Offset))
	}
	if len(in.PetIDs) > 0 {
		builder = builder.Where(sq.Eq{idColumn: in.PetIDs})
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	q := db.Query{
		QueryRow: query,
		Name:     "pet_repository.ListPets",
	}

	rows, err := r.db.DB().Query(ctx, q, args...)
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}
	defer rows.Close()

	var pets []*repomodel.Pet
	err = pgxscan.ScanAll(&pets, rows)
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	return converter.ToPetsFromRepo(pets), nil
}

func (r *repository) Update(ctx context.Context, in *model.UpdatePet) error {
	const op = "pet_repository.Update"

	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar)

	if in.Name != nil {
		builder = builder.Set(nameColumn, in.Name)
	}
	if in.Description != nil {
		builder = builder.Set(descriptionColumn, in.Description)
	}
	if in.Age != nil {
		builder = builder.Set(ageColumn, in.Age)
	}

	query, args, err := builder.Where(sq.Eq{idColumn: in.ID}).ToSql()
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	q := db.Query{
		Name:     "pet_repository.Update",
		QueryRow: query,
	}

	_, err = r.db.DB().Exec(ctx, q, args...)
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	const op = "pet_repository.Delete"

	query, args, err := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id}).
		ToSql()
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	q := db.Query{
		Name:     "pet_repository.Delete",
		QueryRow: query,
	}

	_, err = r.db.DB().Exec(ctx, q, args...)
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	return nil
}
