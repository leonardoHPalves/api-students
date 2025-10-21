# api-students
API to manage  `Golang` course stundents 
 
  Routes:
  - GET /students - List all students 
  - POST /students - Create students
  - GET /students/:id - Get infos from a specific student 
  - PUT /students/:id - Update student 
  - DELETE /students/:id - Delete student

  Struct Student:
  - Name (string)
  - CPF  (int)
  - Email (string)
  - Age (int)
  - Active (bool)