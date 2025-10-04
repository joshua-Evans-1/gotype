package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Playthrough struct {
	gorm.Model
	LevelID		uint
	Wpm			float64
	Cpm			float64
	Accuracy	float64
	GameMode 	string
	ElapsedTime	float64
	Consistency	float64
}

type Level struct {
	gorm.Model
	Language 	string 
	Difficulty 	string 
	textSample	[]string 
} 

var DB *gorm.DB

func ConnDB() ( *gorm.DB, error ) {
	db, err := gorm.Open( sqlite.Open( "test.db" ), &gorm.Config{} )
    if err != nil {
        return nil, err
    }
	
	db.AutoMigrate( &Level{}, &Playthrough{} )

	DB = db

	return db, nil
}

func ( playthrough Playthrough ) Save() ( Playthrough, error ) {
	res := DB.Create( &playthrough )

	if res.Error != nil {
		return playthrough, res.Error
	}

	return playthrough, nil
}


func ( level Level ) getPlaythroughHistory() ( []Playthrough, error ) {
	var playthroughs []Playthrough
	
	res := DB.Where( "LevelID", level.ID ).Find( &playthroughs )

	if res.Error != nil {
		return nil, res.Error
	}

	return playthroughs, nil

}

func ( level Level ) Create() ( Level, error ) {
	res := DB.Create( &level )
	
	if res.Error != nil {
		return level, res.Error
	}

	return level, nil
}

func ( level Level ) GetLevels() ( []Level, error ) {
	var levels []Level
	
	res := DB.Find( &levels ) 

	if res.Error != nil {
		return nil, res.Error
	}

	return levels, nil

}


