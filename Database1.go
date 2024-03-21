package main

type Database1 struct{}

func (db *Database1) GetPhotos() []string {
	return []string{"photo1", "photo2"}
}
