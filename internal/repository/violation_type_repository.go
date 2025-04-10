package repository

import (
	"database/sql"
	"errors"
	"violation-type-service/internal/model"

	sq "github.com/Masterminds/squirrel"
)

type violationTypeRepository struct {
	db         *sql.DB
	sqlBuilder sq.StatementBuilderType
}

func NewViolationTypeRepository(db *sql.DB) ViolationTypeRepository {
	return &violationTypeRepository{
		db:         db,
		sqlBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *violationTypeRepository) FindAll() ([]model.ViolationType, error) {
	query := r.sqlBuilder.Select("id", "name", "other_info").From("violation_types")
	rows, err := query.RunWith(r.db).Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.ViolationType
	for rows.Next() {
		var v model.ViolationType
		if err := rows.Scan(&v.ID, &v.Name, &v.OtherInfo); err != nil {
			return nil, err
		}
		list = append(list, v)
	}
	return list, nil
}

func (r *violationTypeRepository) FindByID(id int64) (model.ViolationType, error) {
	query := r.sqlBuilder.Select("id", "name", "other_info").From("violation_types").Where(sq.Eq{"id": id})
	row := query.RunWith(r.db).QueryRow()

	var v model.ViolationType
	if err := row.Scan(&v.ID, &v.Name, &v.OtherInfo); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return v, errors.New("not found")
		}
		return v, err
	}
	return v, nil
}

func (r *violationTypeRepository) Create(v model.ViolationType) (int64, error) {
	query := r.sqlBuilder.Insert("violation_types").Columns("name", "other_info").
		Values(v.Name, v.OtherInfo).
		Suffix("RETURNING id")

	var id int64
	err := query.RunWith(r.db).QueryRow().Scan(&id)
	return id, err
}

func (r *violationTypeRepository) Update(id int64, v model.ViolationType) error {
	query := r.sqlBuilder.Update("violation_types").
		Set("name", v.Name).
		Set("other_info", v.OtherInfo).
		Where(sq.Eq{"id": id})

	_, err := query.RunWith(r.db).Exec()
	return err
}

func (r *violationTypeRepository) Delete(id int64) error {
	query := r.sqlBuilder.Delete("violation_types").Where(sq.Eq{"id": id})
	_, err := query.RunWith(r.db).Exec()
	return err
}

func (r *violationTypeRepository) BulkInsert(list []model.ViolationType) error {
	q := r.sqlBuilder.Insert("violation_types").Columns("name", "other_info")
	for _, v := range list {
		q = q.Values(v.Name, v.OtherInfo)
	}
	_, err := q.RunWith(r.db).Exec()
	return err
}
