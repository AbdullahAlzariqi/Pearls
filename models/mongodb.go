// models/mongodb.go
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Content represents the Content collection
type Content struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	JobID       string             `bson:"job_id" json:"job_id"`
	URL         string             `bson:"url" json:"url"`
	HTMLContent string             `bson:"html_content" json:"html_content"`
	TextContent string             `bson:"text_content" json:"text_content"`
	Media       []Media            `bson:"media" json:"media"`
	Metadata    Metadata           `bson:"metadata" json:"metadata"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

type Media struct {
	MediaType string `bson:"media_type" json:"media_type"`
	MediaURL  string `bson:"media_url" json:"media_url"`
}

type Metadata struct {
	Title       string   `bson:"title" json:"title"`
	Description string   `bson:"description" json:"description"`
	Keywords    []string `bson:"keywords" json:"keywords"`
}

// Chunk represents the Chunks collection
type Chunk struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ContentID      primitive.ObjectID `bson:"content_id" json:"content_id"`
	ChunkText      string             `bson:"chunk_text" json:"chunk_text"`
	SequenceNumber int                `bson:"sequence_number" json:"sequence_number"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
}

// CleanedData represents the CleanedData collection
type CleanedData struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ContentID      primitive.ObjectID `bson:"content_id" json:"content_id"`
	CleanedText    string             `bson:"cleaned_text" json:"cleaned_text"`
	CleaningConfig CleaningConfig     `bson:"cleaning_config" json:"cleaning_config"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
}

type CleaningConfig struct {
	RemoveStopwords bool     `bson:"remove_stopwords" json:"remove_stopwords"`
	Lowercase       bool     `bson:"lowercase" json:"lowercase"`
	CustomRules     []string `bson:"custom_rules" json:"custom_rules"`
}

// Summaries represents the Summaries collection
type Summary struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ContentID     primitive.ObjectID `bson:"content_id" json:"content_id"`
	SummaryText   string             `bson:"summary_text" json:"summary_text"`
	SummaryConfig SummaryConfig      `bson:"summary_config" json:"summary_config"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
}

type SummaryConfig struct {
	Length string `bson:"length" json:"length"`
	Style  string `bson:"style" json:"style"`
}

// Classification represents the Classifications collection
type Classification struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ContentID primitive.ObjectID `bson:"content_id" json:"content_id"`
	Tags      []string           `bson:"tags" json:"tags"`
	Entities  []string           `bson:"entities" json:"entities"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

// Metric represents the Metrics collection
type Metric struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	JobID       string             `bson:"job_id" json:"job_id"`
	MetricsData MetricsData        `bson:"metrics_data" json:"metrics_data"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}

type MetricsData struct {
	CrawlDuration float64 `bson:"crawl_duration" json:"crawl_duration"`
	PagesCrawled  int     `bson:"pages_crawled" json:"pages_crawled"`
	ChunkCount    int     `bson:"chunk_count" json:"chunk_count"`
}

// Log represents the Logs collection
type Log struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	JobID     string             `bson:"job_id" json:"job_id"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`
	Message   string             `bson:"message" json:"message"`
	Level     string             `bson:"level" json:"level"`
}
