package store

import "../../store/database/sql"

func (s *PostgresStore) AuthenticateRequest(phone, token string) (bool, string, error) {
	query := `SELECT token, role FROM customer WHERE phone = $1`

	var dbToken sql.NullString
	var role string

	err := s.db.QueryRow(query, phone).Scan(&dbToken, &role)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, "", nil
		}
		return false, "", err
	}

	if dbToken.Valid && dbToken.String == token && len(role) > 0 {
		return true, role, nil
	}

	return false, "", nil
}

func (s *PostgresStore) AuthenticateRequestPacker(phone, token string) (bool, string, error) {
	query := `SELECT token, role FROM packer WHERE phone = $1`

	var dbToken sql.NullString
	var role string

	err := s.db.QueryRow(query, phone).Scan(&dbToken, &role)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, "", nil
		}
		return false, "", err
	}

	if dbToken.Valid && dbToken.String == token && len(role) > 0 {
		return true, role, nil
	}

	return false, "", nil
}

func (s *PostgresStore) AuthenticateRequestDeliveryPartner(phone, token string) (bool, string, error) {
	query := `SELECT token, role FROM delivery_partner WHERE phone = $1`

	var dbToken sql.NullString
	var role string

	err := s.db.QueryRow(query, phone).Scan(&dbToken, &role)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, "", nil
		}
		return false, "", err
	}

	if dbToken.Valid && dbToken.String == token && len(role) > 0 {
		return true, role, nil
	}

	return false, "", nil
}
