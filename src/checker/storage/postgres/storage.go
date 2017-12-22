package postgres

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	_ "github.com/lib/pq"

	"checker/config"
	"checker/storage"
)

const (
	fetchLimit = 1000
	queueSize = 100
	queueTimeout = 10 * time.Second
)

type store struct {
	db       *sql.DB
	logsChan chan *storage.Log
}

// NewStore init store and apply migrations
func NewStore(dbConfig *config.DSNConfig) (storage.Store, error) {
	db, err := sql.Open("postgres", dbConfig.String())
	if err != nil {
		log.Fatal(err)
	}

	if err := initMigrations(dbConfig.DbName(), db); err != nil {
		return nil, fmt.Errorf("couldn't apply migrations for current database, err: %s", err.Error())
	}

	store := &store{
		db:       db,
		logsChan: make(chan *storage.Log, queueSize),
	}

	go store.logsStoreWorker()

	return store, nil
}

// AddLog adds log entity to database storage
func (s *store) AddLog(sLog *storage.Log) error {
	select {
		case s.logsChan <- sLog:
		case <- time.After(queueTimeout):
			return fmt.Errorf("timeout: queue is full, log for url: '%s' is skipped", sLog.Url)
	}

	return nil
}

// FetchLogs retrieves ordered list of resources status
func (s *store) FetchLogs(start, end time.Time, url string) ([]*storage.Log, error) {
	var arguments []interface{}
	var whereArgs []string

	var sqlstr = `SELECT created_at, url, successful FROM checker_log`

	arguments = append(arguments, start, end)
	whereArgs = append(whereArgs, fmt.Sprintf(
		"created_at > $%d AND created_at < $%d", len(arguments)-1, len(arguments)),
	)

	if url != "" {
		arguments = append(arguments, url)
		whereArgs = append(whereArgs, fmt.Sprintf("url = $%d", len(arguments)))
	}

	sqlstr += " WHERE " + strings.Join(whereArgs, " AND ")

	arguments = append(arguments, fetchLimit)
	sqlstr += fmt.Sprintf(" ORDER BY created_at LIMIT $%d", len(arguments))

	rows, err := s.db.Query(sqlstr, arguments...)
	if err == sql.ErrNoRows {
		return []*storage.Log{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer Close(rows)

	var sLogs []*storage.Log
	for rows.Next() {
		sLog := &storage.Log{}

		err = rows.Scan(&sLog.CreatedAt, &sLog.Url, &sLog.IsHealthy)
		if err != nil {
			return nil, err
		}

		sLogs = append(sLogs, sLog)
	}

	return sLogs, nil
}

func (s *store) logsStoreWorker() {
	for {
		sLog := <- s.logsChan

		_, err := s.db.Exec(
			"INSERT INTO checker_log(created_at, url, successful) VALUES ($1, $2, $3)",
			sLog.CreatedAt, sLog.Url, sLog.IsHealthy,
		)

		if err != nil {
			log.Println(err.Error())
		}
	}
}

// Close is helper func to use with 'defer' and handling error
func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Println("can not close resource: ", err)
	}
}
