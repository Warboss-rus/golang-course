package handlers

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

type DataBaseConnector struct {
	db *sql.DB
}

func (conn *DataBaseConnector) Connect() error {
	if conn.db != nil {
		return errors.New("Database is already connected")
	}
	const user = "root"
	const password = "root"
	const dbname = "db1"
	const dataSource = user + ":" + password + "@/" + dbname
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return err
	}
	conn.db = db

	if err := db.Ping(); err != nil {
		return err
	}
	return nil
}

func (conn *DataBaseConnector) ConnectTestDatabase() error {
	if conn.db != nil {
		return errors.New("Database is already connected")
	}
	const user = "root"
	const password = "root"
	const dbname = "dbtest"
	const dataSource = user + ":" + password + "@/" + dbname
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return err
	}
	conn.db = db

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

func (conn *DataBaseConnector) GetVideoList() ([]Video, error) {
	if conn.db == nil {
		return nil, errors.New("Database is not connected")
	}
	var videos []Video
	rows, err := conn.db.Query(`SELECT video_key, title, duration, url, thumbnail_url FROM video`)
	if err != nil {
		return videos, err
	}
	defer rows.Close()
	for rows.Next() {
		var video Video
		err := rows.Scan(&video.Id, &video.Name, &video.Duration, &video.Url, &video.Thumbnail)
		if err != nil {
			return videos, err
		}
		videos = append(videos, video)
	}
	return videos, nil
}

func (conn *DataBaseConnector) GetVideoDetails(videoId string) (Video, error) {
	if conn.db == nil {
		return Video{}, errors.New("Database is not connected")
	}
	var video Video
	row := conn.db.QueryRow(`SELECT video_key, title, duration, url, thumbnail_url FROM video WHERE video_key = ?`, videoId)
	err := row.Scan(&video.Id, &video.Name, &video.Duration, &video.Url, &video.Thumbnail)
	return video, err
}

func (conn *DataBaseConnector) AddVideo(video Video) error {
	if conn.db == nil {
		return errors.New("Database is not connected")
	}
	q := `INSERT INTO video SET video_key = ?, title = ?, duration = ?, url = ?, thumbnail_url  = ?`
	_, err := conn.db.Exec(q, video.Id, video.Name, video.Duration, video.Url, video.Thumbnail)
	return err
}

func (conn *DataBaseConnector) Close() error {
	if conn.db != nil {
		return conn.db.Close()
	}
	return nil
}

func (conn *DataBaseConnector) ClearVideos() error {
	if conn.db == nil {
		return errors.New("Database is not connected")
	}
	_, err := conn.db.Exec("DROP TABLE IF EXISTS video")
	return err
}
