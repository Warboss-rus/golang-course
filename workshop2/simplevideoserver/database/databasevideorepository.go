package database

import (
	"database/sql"
	"errors"
	. "github.com/Warboss-rus/golang-course/workshop2/simplevideoserver/handlers"
	_ "github.com/go-sql-driver/mysql"
)

type DataBaseVideoRepository struct {
	db *sql.DB
}

func (repository *DataBaseVideoRepository) Connect() error {
	if repository.db != nil {
		return errors.New("database is already connected")
	}
	const user = "root"
	const password = "root"
	const dbname = "db1"
	const dataSource = user + ":" + password + "@/" + dbname
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return err
	}
	repository.db = db

	if err := db.Ping(); err != nil {
		return err
	}
	return nil
}

func (repository *DataBaseVideoRepository) ConnectTestDatabase() error {
	if repository.db != nil {
		return errors.New("database is already connected")
	}
	const user = "root"
	const password = "root"
	const dbname = "dbtest"
	const dataSource = user + ":" + password + "@/" + dbname
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return err
	}
	repository.db = db

	if err := db.Ping(); err != nil {
		return err
	}

	_, err = db.Exec("DROP TABLE IF EXISTS video;")
	if err != nil {
		return err
	}
	_, err = db.Exec(`CREATE TABLE video
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

func (repository *DataBaseVideoRepository) GetVideoList(search string, start *uint, count *uint) ([]Video, error) {
	if repository.db == nil {
		return nil, errors.New("database is not connected")
	}
	var videos []Video
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
		var video Video
		err := rows.Scan(&video.Id, &video.Name, &video.Duration, &video.Url, &video.Thumbnail, &video.Status)
		if err != nil {
			return videos, err
		}
		videos = append(videos, video)
	}
	return videos, nil
}

func (repository *DataBaseVideoRepository) GetVideoDetails(videoId string) (Video, error) {
	if repository.db == nil {
		return Video{}, errors.New("database is not connected")
	}
	var video Video
	row := repository.db.QueryRow(`SELECT video_key, title, duration, url, thumbnail_url, status FROM video WHERE video_key = ?`, videoId)
	err := row.Scan(&video.Id, &video.Name, &video.Duration, &video.Url, &video.Thumbnail, &video.Status)
	if err == sql.ErrNoRows {
		return video, &VideoNotFound{}
	}
	return video, err
}

func (repository *DataBaseVideoRepository) AddVideo(video Video) error {
	if repository.db == nil {
		return errors.New("database is not connected")
	}
	q := `INSERT INTO video SET video_key = ?, title = ?, duration = ?, url = ?, thumbnail_url  = ?, status = ?`
	_, err := repository.db.Exec(q, video.Id, video.Name, video.Duration, video.Url, video.Thumbnail, video.Status)
	return err
}

func (repository *DataBaseVideoRepository) GetVideoStatus(videoId string) (Status, error) {
	if repository.db == nil {
		return Error, errors.New("database is not connected")
	}
	var status Status
	row := repository.db.QueryRow(`SELECT status FROM video WHERE video_key = ?`, videoId)
	err := row.Scan(&status)
	if err == sql.ErrNoRows {
		return Error, &VideoNotFound{}
	}
	return status, err
}

func (repository *DataBaseVideoRepository) Close() error {
	if repository.db != nil {
		return repository.db.Close()
	}
	return nil
}

func (repository *DataBaseVideoRepository) ClearVideos() error {
	if repository.db == nil {
		return errors.New("database is not connected")
	}
	_, err := repository.db.Exec("DROP TABLE IF EXISTS video")
	return err
}
