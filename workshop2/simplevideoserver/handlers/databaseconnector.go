package handlers

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DataBaseConnector struct {
	db *sql.DB
}

func (conn *DataBaseConnector) Connect() {
	const user = "root"
	const password = "root"
	const dbname = "db1"
	const dataSource = user + ":" + password + "@/" + dbname
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	conn.db = db

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func (conn *DataBaseConnector) GetVideoList() ([]Video, error) {
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
	var video Video
	row := conn.db.QueryRow(`SELECT video_key, title, duration, url, thumbnail_url FROM video WHERE video_key = ?`, videoId)
	err := row.Scan(&video.Id, &video.Name, &video.Duration, &video.Url, &video.Thumbnail)
	return video, err
}

func (conn *DataBaseConnector) AddVideo(video Video) error {
	q := `INSERT INTO video SET video_key = ?, title = ?, duration = ?, url = ?, thumbnail_url  = ?`
	_, err := conn.db.Exec(q, video.Id, video.Name, video.Duration, video.Url, video.Thumbnail)
	return err
}

func (conn *DataBaseConnector) Close() {
	conn.db.Close()
}
