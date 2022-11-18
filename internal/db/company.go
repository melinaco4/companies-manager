package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/melinaco4/companies-manager/internal/company"
	uuid "github.com/satori/go.uuid"
)

type CompanyRow struct {
	ID                string
	Name              sql.NullString
	Description       sql.NullString
	AmountofEmployees string
	Registered        bool
	Type              sql.NullString
}

func convertCompanyRowToCompany(c CompanyRow) company.Company {
	return company.Company{
		ID:                c.ID,
		Name:              c.Name.String,
		Description:       c.Description.String,
		AmountofEmployees: c.AmountofEmployees,
		Registered:        c.Registered,
		Type:              c.Type.String,
	}
}

func (d *Database) GetCompany(ctx context.Context, uuid string) (company.Company, error) {
	var cmpnyRow CompanyRow
	row := d.Client.QueryRowContext(
		ctx,
		`SELECT id, name, description, amountofemployees, registered, type 
		FROM companies
		WHERE id = $1`,
		uuid,
	)

	err := row.Scan(&cmpnyRow.ID,
		&cmpnyRow.Name,
		&cmpnyRow.Description,
		&cmpnyRow.AmountofEmployees,
		&cmpnyRow.Registered,
		&cmpnyRow.Type)

	if err != nil {
		return company.Company{}, fmt.Errorf("an error occurred fetching a comment by uuid: %w", err)
	}
	// sqlx with context to ensure context cancelation is honoured
	return convertCompanyRowToCompany(cmpnyRow), nil
}

func (d *Database) PostCompany(ctx context.Context, cmpny company.Company) (company.Company, error) {
	cmpny.ID = uuid.NewV4().String()
	postRow := CompanyRow{
		ID:                cmpny.ID,
		Name:              sql.NullString{String: cmpny.Name, Valid: true},
		Description:       sql.NullString{String: cmpny.Description, Valid: true},
		AmountofEmployees: cmpny.AmountofEmployees,
		Registered:        cmpny.Registered,
		Type:              sql.NullString{String: cmpny.Type, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO comments 
		(id, name, description, amountofemployees, registered, type) VALUES
		(:id, :name, :description, :amountofemployees, :registered, :type)`,
		postRow,
	).
	if err != nil {
		return company.Company{}, fmt.Errorf("failed to insert company: %w", err)
	}
	if err := rows.Close(); err != nil {
		return company.Company{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return cmt, nil
}


func (d *Database) UpdateComment(ctx context.Context, id string, cmt company.Company) (company.Company, error) {
	cmpnyRow := CompanyRow{
		ID:                cmpny.ID,
		Name:              sql.NullString{String: cmpny.Name, Valid: true},
		Description:       sql.NullString{String: cmpny.Description, Valid: true},
		AmountofEmployees: cmpny.AmountofEmployees,
		Registered:        cmpny.Registered,
		Type:              sql.NullString{String: cmpny.Type, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`UPDATE companies SET
		name = :name,
		description = :description,
		amountofemployees = :amountofemployees,
		registered = registered,
		type = type 
		WHERE id = :id`,
		cmpnyRow,
	)
	if err != nil {
		return company.Company{}, fmt.Errorf("failed to insert comment: %w", err)
	}
	if err := rows.Close(); err != nil {
		return company.Company{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return convertCompanyRowToCompany(cmpnyRow), nil
}

func (d *Database) DeleteCompany(ctx context.Context, id string) error {
	_, err := d.Client.ExecContext(
		ctx,
		`DELETE FROM companies where id = $1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("failed to delete company from the database: %w", err)
	}
	return nil
}
