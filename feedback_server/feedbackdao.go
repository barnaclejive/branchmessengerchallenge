package main

import (
  "log"
  "os"
  "time"

  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)

// The database will be exposed as a singleton of the FeedbackDaoManager type.
// Database operations are exposed as methods on the FeedbackDao object.
type FeedbackDaoManager struct {
  db *sql.DB
}
var FeedbackDao FeedbackDaoManager

// Initalize the database.
// To keep things simple for this exercise I chose to use a on disk sqlite database.
// You can re-create it on startup my manually setting 'initialDatabase' to true.
const initialDatabase = true

func init() {

  if initialDatabase {
    os.Remove("./feedback.db")
  }

  db, err := sql.Open("sqlite3", "./feedback.db")
  if err != nil {
    log.Fatal(err)
  }

  if initialDatabase {

    schemaSql := `
    CREATE TABLE feedback (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      timestamp INTEGER NOT NULL,
      session_id TEXT NOT NULL,
      user_id TEXT NOT NULL,
      rating INTEGER NOT NULL,
      comment TEXT,
      UNIQUE(session_id, user_id) ON CONFLICT FAIL
    );
    `
    _, err = db.Exec(schemaSql)
    if err != nil {
      log.Fatal(err)
    }
  }

  FeedbackDao = FeedbackDaoManager{db}
}

// Inserts a new Feedback item into the base if one does not already exist for the user+session.
func (m FeedbackDaoManager) insertFeedback(feedback Feedback) bool {

  if m.userHasFeedbackForSession(feedback.UserId, feedback.SessionId) {
    return false
  }

  tx, err := m.db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO feedback(timestamp, session_id, user_id, rating, comment) values(?, ?, ?, ?, ?);")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

  _, err = stmt.Exec(int32(feedback.Timestamp.Unix()), feedback.SessionId, feedback.UserId, feedback.Rating, feedback.Comment)
  if err != nil {
    log.Fatal(err)
  }
  tx.Commit()

  return true
}

// Checks if feedback already exists for a given user+session combination.
func (m FeedbackDaoManager) userHasFeedbackForSession(userId string, sessionId string) bool {

  stmt, err := m.db.Prepare("SELECT id FROM feedback WHERE session_id = ? AND user_id = ?;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

  rows, err := stmt.Query(sessionId, userId)
	if err != nil {
		log.Fatal(err)
  }
  defer rows.Close()

  return rows.Next()
}

// Select up to 15 feedback records, sorted by timestamp descending, optionally filtered by rating.
// Passing 0 for rating will not cause any filtering by rating.
func (m FeedbackDaoManager) selectFeedback(rating int) []Feedback {

  selectSql := `
  SELECT timestamp, session_id, user_id, rating, comment
  FROM feedback
  `
  if rating > 0 {
    selectSql += `
    WHERE rating = ?
    `
  }
  selectSql += `
  ORDER BY timestamp DESC
  LIMIT 15;
  `
  stmt, err := m.db.Prepare(selectSql)
	if err != nil {
		log.Fatal(err)
	}
  defer stmt.Close()

  var rows *sql.Rows
  if rating > 0 {
    rows, err = stmt.Query(rating)
  } else {
    rows, err = stmt.Query()
  }
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  results := []Feedback{}

  for rows.Next() {

  	var timestamp int64
    var session_id string
    var user_id string
    var rating int
    var comment string

  	err = rows.Scan(&timestamp, &session_id, &user_id, &rating, &comment)
  	if err != nil {
  		log.Fatal(err)
  	}

    feedback := Feedback{time.Unix(timestamp, 0), session_id, user_id, rating, comment}
    results = append(results, feedback)
  }

  return results
}
