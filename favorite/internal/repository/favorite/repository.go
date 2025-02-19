package favorite

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/escoutdoor/kotopes/common/pkg/db"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
	"github.com/escoutdoor/kotopes/favorite/internal/model"
	def "github.com/escoutdoor/kotopes/favorite/internal/repository"
	"github.com/escoutdoor/kotopes/favorite/internal/repository/favorite/converter"
	repomodel "github.com/escoutdoor/kotopes/favorite/internal/repository/favorite/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

var (
	tableName = "favorites"

	idColumn        = "id"
	userIdColumn    = "user_id"
	petIdColumn     = "pet_id"
	createdAtColumn = "created_at"

	ErrNotFound = errors.New("favorite not found")
)

type repository struct {
	db db.Client
}

var _ def.FavoriteRepository = (*repository)(nil)

func New(db db.Client) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, in *model.CreateFavorite) (string, error) {
	const op = "favorite_repository.Create"

	query, args, err := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(userIdColumn, petIdColumn).
		Values(in.UserID, in.PetID).
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

func (r *repository) IsPetFavoriteExists(ctx context.Context, petID, userID string) (bool, error) {
	const op = "favorite_repository.IsPetFavoriteExists"

	query, args, err := sq.Select(idColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.And{
			sq.Eq{userIdColumn: userID},
			sq.Eq{petIdColumn: petID},
		}).
		ToSql()
	if err != nil {
		return false, errwrap.Wrap(op, err)
	}

	q := db.Query{
		Name:     op,
		QueryRow: query,
	}

	var id string
	err = r.db.DB().QueryRow(ctx, q, args...).Scan(&id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return false, errwrap.Wrap(op, err)
	}

	return len(id) > 0, nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	const op = "favorite_repository.Delete"

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
		return errwrap.Wrap(op, fmt.Errorf("no rows deleted"))
	}
	return nil
}

func (r *repository) GetByID(ctx context.Context, id string) (*model.Favorite, error) {
	const op = "favorite_repository.GetByID"

	query, args, err := sq.Select(idColumn, userIdColumn, petIdColumn, createdAtColumn).
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

	var fav repomodel.Favorite
	err = pgxscan.ScanOne(&fav, row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model.ErrFavoriteNotFound
		}
		return nil, errwrap.Wrap(op, err)
	}

	return converter.ToFavoriteFromRepo(&fav), nil
}

func (r *repository) List(ctx context.Context, userID string) ([]*model.Favorite, error) {
	const op = "favorite_repository.List"

	query, args, err := sq.Select(idColumn, userIdColumn, petIdColumn, createdAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{userIdColumn: userID}).
		ToSql()
	if err != nil {
		return nil, err
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

	var favs []*repomodel.Favorite
	err = pgxscan.ScanAll(&favs, rows)
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	return converter.ToFavoritesFromRepo(favs), nil
}
