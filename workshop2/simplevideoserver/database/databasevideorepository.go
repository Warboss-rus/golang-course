package database

import (
	"database/sql"
	"errors"
	"github.com/Warboss-rus/golang-course/workshop2/simplevideoserver/handlers"
	"github.com/Warboss-rus/golang-course/workshop4/videoProcessingDaemon/videoprocessing"
	// we need this as is
	_ "github.com/go-sql-driver/mysql"
)

// DBVideoRepository is a MySQL implementation of handlers.VideosRepository and videoprocessing.VideoRepository
type DBVideoRepository struct {
	db *sql.DB
}

// Connect conntects to the specified database using specified user and password
func (repository *DBVideoRepository) Connect(dbname string, user string, password string) error {
	if repository.db != nil {
		return errors.New("database is already connected")
	}

	dataSource := user + ":" + password + "@/" + dbname
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return err
	}
	repository.db = db

	if err := db.Ping(); err != nil {
		return err
	}
	return repository.createTableIfNotExists()
}

func (repository *DBVideoRepository) createTableIfNotExists() error {
	_, err := repository.db.Exec(`CREATE TABLE IF NOT EXISTS video
		(
		    id            INT UNSIGNED UNIQUE NOT NULL AUTO_INCREMENT,
		    video_key     VARCHAR(255) UNIQUE,
		    title         VARCHAR(255)        NOT NULL,
		    status        TINYINT                      DEFAULT 1,
		    duration      INT UNSIGNED                 DEFAULT 0,
		    url           VARCHAR(255)        NOT NULL,
		    thumbnail_url VARCHAR(255)        NOT NULL DEFAULT '',
		    PRIMARY KEY (id)
		);`)
	return err
}

// GetVideoList returns a list of videos satisfying the criteria
func (repository *DBVideoRepository) GetVideoList(search string, start *uint, count *uint) ([]handlers.Video, error) {
	if repository.db == nil {
		return nil, errors.New("database is not connected")
	}
	videos := make([]handlers.Video, 0)
	query := `SELECT video_key, title, duration, url, thumbnail_url, status FROM video`
	var args []interface{}
	if len(search) > 0 {
		query += ` WHERE title LIKE ?`
		args = append(args, "%"+search+"%")
	}
	if count != nil {
		query += ` LIMIT ?`
		args = append(args, *count)
	}
	if start != nil {
		query += ` OFFSET ?`
		args = append(args, *start)
	}

	rows, err := repository.db.Query(query, args...)
	if err != nil {
		return videos, err
	}
	defer rows.Close()
	for rows.Next() {
		var video handlers.Video
		err := rows.Scan(&video.ID, &video.Name, &video.Duration, &video.URL, &video.Thumbnail, &video.Status)
		if err != nil {
			return videos, err
		}
		videos = append(videos, video)
	}
	return videos, nil
}

// GetVideoDetails returns the details of single video
func (repository *DBVideoRepository) GetVideoDetails(videoID string) (handlers.Video, error) {
	if repository.db == nil {
		return handlers.Video{}, errors.New("database is not connected")
	}
	var video handlers.Video
	row := repository.db.QueryRow(`SELECT video_key, title, duration, url, thumbnail_url, status FROM video WHERE video_key = ?`, videoID)
	err := row.Scan(&video.ID, &video.Name, &video.Duration, &video.URL, &video.Thumbnail, &video.Status)
	if err == sql.ErrNoRows {
		return video, &handlers.VideoNotFound{}
	}
	return video, err
}

// AddVideo adds a new video to the table
func (repository *DBVideoRepository) AddVideo(video handlers.Video) error {
	if repository.db == nil {
		return errors.New("database is not connected")
	}
	const q = `INSERT INTO video SET video_key = ?, title = ?, duration = ?, url = ?, thumbnail_url  = ?, status = ?`
	_, err := repository.db.Exec(q, video.ID, video.Name, video.Duration, video.URL, video.Thumbnail, video.Status)
	return err
}

// GetVideoStatus returns the status of the video
func (repository *DBVideoRepository) GetVideoStatus(videoID string) (handlers.Status, error) {
	if repository.db == nil {
		return handlers.Error, errors.New("database is not connected")
	}
	var status handlers.Status
	row := repository.db.QueryRow(`SELECT status FROM video WHERE video_key = ?`, videoID)
	err := row.Scan(&status)
	if err == sql.ErrNoRows {
		return handlers.Error, &handlers.VideoNotFound{}
	}
	return status, err
}

// Close closes connection to the database
func (repository *DBVideoRepository) Close() error {
	if repository.db != nil {
		return repository.db.Close()
	}
	return nil
}

func (repository *DBVideoRepository) removeTable() error {
	if repository.db == nil {
		return errors.New("database is not connected")
	}
	_, err := repository.db.Exec("DROP TABLE IF EXISTS video")
	return err
}

// GetVideosByStatus returns a list of videos that have specified status
func (repository *DBVideoRepository) GetVideosByStatus(status videoprocessing.Status) ([]videoprocessing.Video, error) {
	if repository.db == nil {
		return nil, errors.New("database is not connected")
	}
	videos := make([]videoprocessing.Video, 0)
	const query = `SELECT video_key, url FROM video WHERE status = ?`
	rows, err := repository.db.Query(query, status)
	if err != nil {
		return videos, err
	}
	defer rows.Close()
	for rows.Next() {
		var video videoprocessing.Video
		err := rows.Scan(&video.ID, &video.URL)
		if err != nil {
			return videos, err
		}
		videos = append(videos, video)
	}
	return videos, nil
}

// UpdateVideoStatus changes the status of specified video
func (repository *DBVideoRepository) UpdateVideoStatus(videoID string, status videoprocessing.Status) error {
	if repository.db == nil {
		return errors.New("database is not connected")
	}
	const q = `UPDATE video SET status = ? WHERE video_key = ?`
	_, err := repository.db.Exec(q, status, videoID)
	return err
}

// UpdateVideo inserts video details after processing
func (repository *DBVideoRepository) UpdateVideo(videoID string, duration int, thumbnail string, status videoprocessing.Status) error {
	if repository.db == nil {
		return errors.New("database is not connected")
	}
	const q = `UPDATE video SET duration = ?, thumbnail_url = ?, status = ? WHERE video_key = ?`
	_, err := repository.db.Exec(q, duration, thumbnail, status, videoID)
	return err
}
