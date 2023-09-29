package models

type School struct {
	Id   int     `json: "id"`
	Name string  `json: "name"`
    Class [] Classes
}

type Classes struct {
	Id   int    `json: "id"`
	Name string `json: "name"`
    Pupil [] Pupils
}

type Pupils struct {
	Id    int    `json: "id"`
	Name  string `json: "name"`
}