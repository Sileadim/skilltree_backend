package mysql

import (
	"database/sql"
	"errors"

	"github.com/Sileadim/skilltree_backend/pkg/models"
)

// Define a SnippetModel type which wraps a sql.DB connection pool.
type TreeModel struct {
	DB *sql.DB
}

// Insert tree
func (m *TreeModel) Insert(title, uuid, content string) (int, error) {
	// Write the SQL statement we want to execute. I've split it over two lines
	// for readability (which is why it's surrounded with backquotes instead
	// of normal double quotes).
	stmt := `INSERT INTO trees (title, uuid, content, created)
    VALUES(?, ?, ?, UTC_TIMESTAMP())`

	// Use the Exec() method on the embedded connection pool to execute the
	// statement. The first parameter is the SQL statement, followed by the
	// title, content and expiry values for the placeholder parameters. This
	// method returns a sql.Result object, which contains some basic
	// information about what happened when the statement was executed.
	result, err := m.DB.Exec(stmt, title, uuid, content)
	if err != nil {
		return 0, err
	}

	// Use the LastInsertId() method on the result object to get the ID of our
	// newly inserted record in the snippets table.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// The ID returned has the type int64, so we convert it to an int type
	// before returning.
	return int(id), nil
}

// Get ids
func (m *TreeModel) Get(id int) (*models.Tree, error) {
	// Write the SQL statement we want to execute. Again, I've split it over two
	// lines for readability.
	stmt := `SELECT id, uuid, title, content, created FROM trees
	WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	// Initialize a pointer to a new zeroed Tree struct.
	s := &models.Tree{}

	// Use row.Scan() to copy the values from each field in sql.Row to the
	// corresponding field in the Snippet struct. Notice that the arguments
	// to row.Scan are *pointers* to the place you want to copy the data into,
	// and the number of arguments must be exactly the same as the number of
	// columns returned by your statement.

	err := row.Scan(&s.ID, &s.Uuid, &s.Title, &s.Content, &s.Created)
	if err != nil {
		// If the query returns no rows, then row.Scan() will return a
		// sql.ErrNoRows error. We use the errors.Is() function check for that
		// error specifically, and return our own models.ErrNoRecord error
		// instead.
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	// If everything went OK then return the Snippet object.
	return s, nil
}

// This will return the 10 most recently created snippets.
func (m *TreeModel) List() ([]*models.Tree, error) {
	// todo: make this
	stmt := `SELECT id, title, uuid, content, created FROM trees
    ORDER BY created DESC LIMIT 100`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	trees := []*models.Tree{}
	for rows.Next() {
		// Create a pointer to a new zeroed Snippet struct.
		t := &models.Tree{}
		err = rows.Scan(&t.ID, &t.Title, &t.Uuid, &t.Content, &t.Created)
		if err != nil {
			return nil, err
		}
		trees = append(trees, t)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return trees, nil
}
