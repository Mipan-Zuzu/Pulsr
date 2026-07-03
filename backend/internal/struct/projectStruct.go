package Struct

type Database struct {
	Host string `json:"host"`
	Version string `json:"version"`
	Postgres_engine string `json:"postgres_engine"`
	Release_channel string `json:"release_channel"`
}

type Project struct {
	Id string `json:"id"`
	Ref string `json:"ref"`
	Organization_id string `json:"organization_id"` 
	Organization_slug string `json:"organization_slug"`
	Name string `json:"name"`
	Region string `json:"region"`
	Status string `json:"status"`
	Database Database `json:"database"`
	Created_at string `json:"created_at"`
}