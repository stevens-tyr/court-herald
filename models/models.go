package models

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// WorkerResult stores the result of the test cases
type WorkerResult struct {
	ID            int    `bson:"id" json:"id" binding:"required"`
	Panicked      bool   `bson:"panicked" json:"panicked" binding:"required"`
	Passed        bool   `bson:"passed" json:"passed" binding:"required"`
	StudentFacing bool   `bson:"studentFacing" json:"studentFacing" binding:"required"`
	Output        string `bson:"output" json:"output" binding:"required"`
	HTML          string `bson:"html" json:"html" binding:"required"`
	TestCMD       string `bson:"testCMD" json:"testCMD" binding:"required"`
	Name          string `bson:"name" json:"name" binding:"required"`
}

// MongoSubmission struct the struct to represent a submission to an page.
type MongoSubmission struct {
	ID             primitive.ObjectID `bson:"_id" json:"id" binding:"required"`
	UserID         primitive.ObjectID `bson:"userID" json:"userID" binding:"required"`
	FileID         primitive.ObjectID `bson:"fileID" json:"fileID" binding:"required"`
	AssignmentID   primitive.ObjectID `bson:"assignmentID" json:"assignmentID" binding:"required"`
	AttemptNumber  int                `bson:"attemptNumber" json:"attemptNumber" binding:"required"`
	SubmissionDate primitive.DateTime `bson:"submissionDate" json:"submissionDate" binding:"required"`
	File           string             `bson:"file" json:"file" binding:"required"`
	ErrorTesting   bool               `bson:"errorTesting" json:"errorTesting" binding:"exists"`
	Results        []WorkerResult     `bson:"results" json:"results" binding:"exists"`
	InProgress     bool               `bson:"inProgress" json:"inProgress"`
}

// Test is the struct for each test case used on a student's submission
type Test struct {
	Name           string `bson:"name" json:"name" binding:"required"`
	ExpectedOutput string `bson:"expectedOutput" json:"expectedOutput" binding:"required"`
	StudentFacing  bool   `bson:"studentFacing" json:"studentFacing" binding:"exists"`
	TestCMD        string `bson:"testCMD" json:"testCMD" binding:"required"`
}

// RequestData is the struct that holds the info sent from the backend
type RequestData struct {
	Submission   MongoSubmission `json:"submission"`
	TestBuildCMD string          `json:"testBuildCMD"`
	Tests        []Test          `json:"tests"`
	Language     string          `json:"language"`
}
