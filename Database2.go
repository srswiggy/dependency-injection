package main

type Database2 struct {
}

func (db *Database2) GetPhotos() []string {
	return []string{"photo3", "photo4"}
}
