package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Course represents a course record
type Course struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	PrimaryTopic string `json:"primary_topic"`
}

type Review struct {
	CourseID   int     `json:"course_id"`
	StudentID  int     `json:"student_id"`
	StarRating float64 `json:"star_rating"`
	ReviewText string  `json:"review"`
}

// RandomCourse returns a randomly generated course
func (d *database) RandomCourse() *Course {
	rand.Seed(time.Now().UnixNano())

	topics := []string{"Math", "Science", "History", "Art", "Music", "Literature"}

	return &Course{
		Title:        fmt.Sprintf("Course %d", rand.Intn(100)),
		Description:  fmt.Sprintf("This is a course about %s", topics[rand.Intn(len(topics))]),
		PrimaryTopic: topics[rand.Intn(len(topics))],
	}
}

// RandomReview returns a randomly generated review
func (d *database) RandomReview() *Review {
	var courseID int
	err := d.db.QueryRow("SELECT course_id FROM courses ORDER BY RANDOM() LIMIT 1").Scan(&courseID)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	rand.Seed(time.Now().UnixNano())

	// random integer between 1 and 100 (inclusive).
	studentID := rand.Intn(100) + 1
	rating := float64(rand.Intn(5) + 1) // generate a rating between 1 and 5
	return &Review{
		CourseID:   courseID,
		StudentID:  studentID,
		StarRating: rating,
		ReviewText: fmt.Sprintf("This course is %s", []string{"awful", "bad", "okay", "good", "excellent"}[int(rating-1)]),
	}
}

func (d *database) KeepInsertingCourses() {
	// Insert random data into the courses table indefinitely
	for {
		course := d.RandomCourse()

		_, err := d.db.Exec("INSERT INTO courses (title, description, primary_topic) VALUES ($1, $2, $3)", course.Title, course.Description, course.PrimaryTopic)

		if err != nil {
			panic(err)
		}
		fmt.Println("Inserted course:", course)
		time.Sleep(time.Millisecond * 500) // wait for 500ms before inserting the next record
	}
}

func (d *database) KeepInsertingReviews() {
	// Insert random data into the courses table indefinitely
	for {
		review := d.RandomReview()

		if review == nil {
			time.Sleep(time.Second)
			continue
		}

		_, err := d.db.Exec("INSERT INTO reviews (course_id, student_id, star_rating, review) VALUES ($1, $2, $3, $4)", review.CourseID, review.StudentID, review.StarRating, review.ReviewText)

		if err != nil {
			panic(err)
		}

		fmt.Println("Inserted review:", review)
		time.Sleep(time.Millisecond * 500) // wait for 500ms before inserting the next record
	}
}
