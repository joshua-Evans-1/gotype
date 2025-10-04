# Database

## Database.go

implement a Gorm instance and creates fucntions to save and get metrics and levels from a sqlite database

```
┌─────────────────────┐         ┌─────────────────────┐
│ Playthrough         │         │ Level               │
│─────────────────────│         │─────────────────────│
│ ID          uint    │    ┌───►│ ID         uint     │
│ UpdatedAt   time    │    │    │ CreatedAt  time     │
│ DeletedAt   time    │    │    │ UpdatedAt  time     │
│ LevelID     uint    ├────┘    │ DeletedAt  time     │
│ Wpm         float64 │         │ Language   string   │
│ Cpm         float64 │         │ Difficulty string   │
│ Accuracy    float64 │         │ TextSample []string │
│ GameMode    string  │         └─────────────────────┘
│ ElapsedTime float64 │                                
│ Consistency float64 │                                
└─────────────────────┘                                
```
