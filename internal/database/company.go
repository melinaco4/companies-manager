package database

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
	AmountofEmployees int
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

func (d *Database) GetCompany(
	ctx context.Context,
	uuid string,
) (company.Company, error) {
	var cmpnRow CompanyRow
	row := d.Client.QueryRowContext(
		ctx,
		`SELECT id, name, description, amountofemployees, registered, type
		FROM companies
		WHERE id = $1`,
		uuid,
	)
	err := row.Scan(&cmpnRow.ID, &cmpnRow.Name, &cmpnRow.Description, &cmpnRow.AmountofEmployees, &cmpnRow.Registered, &cmpnRow.Type)
	if err != nil {
		return company.Company{}, fmt.Errorf("error fetching the company by uuid: %w", err)
	}

	return convertCompanyRowToCompany(cmpnRow), nil
}

func (d *Database) PostCompany(ctx context.Context, cmpn company.Company) (company.Company, error) {
	cmpn.ID = uuid.NewV4().String()
	postRow := CompanyRow{
		ID:                cmpn.ID,
		Name:              sql.NullString{String: cmpn.Name, Valid: true},
		Description:       sql.NullString{String: cmpn.Description, Valid: true},
		AmountofEmployees: cmpn.AmountofEmployees,
		Registered:        cmpn.Registered,
		Type:              sql.NullString{String: cmpn.Type, Valid: true},
	}
	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO comments
		(id, name, description, amountofemployees, registered, type)
		VALUES
		(:id, :name, :description, :amountofemployees, :registered, :type)`,
		postRow,
	)
	if err != nil {
		return company.Company{}, fmt.Errorf("failed to insert company: %w", err)
	}
	if err := rows.Close(); err != nil {
		return company.Company{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return cmpn, nil
}

func (d *Database) DeleteCompany(ctx context.Context, id string) error {
	_, err := d.Client.ExecContext(
		ctx,
		`DELETE FROM companies where id = $1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("failed to delete company from database: %w", err)
	}
	return nil
}

func (d *Database) UpdateCompany(
	ctx context.Context,
	id string,
	cmpn company.Company,
) (company.Company, error) {
	cmpnRow := CompanyRow{
		ID:                id,
		Name:              sql.NullString{String: cmpn.Name, Valid: true},
		Description:       sql.NullString{String: cmpn.Description, Valid: true},
		AmountofEmployees: cmpn.AmountofEmployees,
		Registered:        cmpn.Registered,
		Type:              sql.NullString{String: cmpn.Type, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`UPDATE comments SET
		name = :slug,
		description = :author,
		amountofemployees = :body,
		registered = :registered,
		type = :type
		WHERE id = :id`,
		cmpnRow,
	)
	if err != nil {
		return company.Company{}, fmt.Errorf("failed to update company: %w", err)
	}
	if err := rows.Close(); err != nil {
		return company.Company{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return convertCompanyRowToCompany(cmpnRow), nil
}
