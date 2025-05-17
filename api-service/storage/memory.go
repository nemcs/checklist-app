package storage

import "github.com/nemcs/checklist-app/api-service/model"

//├── storage/                # работа с базой (пока map, потом будет PostgreSQL)
//│   └── memory.go           # временная реализация на map

var Tasks = make(map[string]model.Task)
