# REST API

POST /create -- create task -- 201, JSON (ID, Message)
GET /list -- get list tasks -- 200, JSON ([]Task)
DELETE /delete -- delete task -- 400
PUT /done -- change status -- 400