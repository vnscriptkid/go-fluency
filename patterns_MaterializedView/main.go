package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type database struct {
	db *sql.DB
}

func InitDbConn() *database {
	const (
		host     = "localhost"
		port     = 5434
		user     = "postgres"
		password = "123456"
		dbname   = "udemy"
	)

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic("failed to open db")
	}

	err = db.Ping()

	if err != nil {
		panic("ping err: " + err.Error())
	}

	return &database{
		db: db,
	}
}

func (d *database) CreateTables(ctx context.Context) error {
	fmt.Println("Create tables")
	txn, err := d.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	_, err = txn.Exec(`
		DROP MATERIALIZED VIEW IF EXISTS mv_course_reviews;
		DROP TRIGGER IF EXISTS update_mv_course_reviews ON reviews;
		DROP FUNCTION IF EXISTS refresh_mv_course_reviews();
	
		CREATE TABLE IF NOT EXISTS courses (
			course_id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			description TEXT NOT NULL,
			primary_topic TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS reviews (
			review_id SERIAL PRIMARY KEY,
			course_id INTEGER NOT NULL,
			student_id INTEGER NOT NULL,
			star_rating FLOAT NOT NULL,
			review TEXT
		);

		CREATE MATERIALIZED VIEW mv_course_reviews AS
		SELECT 
			c.course_id, 
			c.title, 
			c.description, 
			AVG(r.star_rating) as avg_rating,
			SUM(r.star_rating) as total_rating,
			COUNT(*) as count
		FROM courses c
		INNER JOIN reviews r ON c.course_id = r.course_id
		GROUP BY c.course_id
		ORDER BY avg_rating DESC, total_rating DESC;

		CREATE OR REPLACE FUNCTION refresh_mv_course_reviews()
		RETURNS TRIGGER AS $$
		BEGIN
			REFRESH MATERIALIZED VIEW mv_course_reviews;
			RETURN NULL;
		END
		$$ LANGUAGE plpgsql;
		
		CREATE TRIGGER update_mv_course_reviews
		AFTER INSERT OR UPDATE OR DELETE ON reviews
		FOR EACH ROW
		EXECUTE FUNCTION refresh_mv_course_reviews();
	`)

	if err != nil {
		err2 := txn.Rollback()

		if err2 != nil {
			fmt.Println("failed to rollback " + err2.Error())
			return err2
		}

		return err
	}

	return txn.Commit()
}

func main() {

	fmt.Println("Start")
	dbConn := InitDbConn()

	defer dbConn.db.Close()

	err := dbConn.CreateTables(context.Background())

	if err != nil {
		panic("failed to create tables " + err.Error())
	}

	done := make(chan bool)

	go dbConn.KeepInsertingCourses()

	go dbConn.KeepInsertingReviews()

	fmt.Println("End", <-done)
}
