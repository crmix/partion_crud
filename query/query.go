package query

var SCHOOL_SQL = ` 
  SELECT * FROM schools OFFSET $1 LIMIT $2

`
var CLASS_SQL = `
  SELECT 
    id,
    name
  FROM classes
  WHERE schools_id=$1

`

var PUPIL_SQL = `
  SELECT
   id,
   name
  FROM pupils
  WHERE classes_id=$1

`
