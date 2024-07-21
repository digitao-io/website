db.createCollection("pages")
db.createCollection("resources")

db.createCollection("articles")
db.createCollection("tags")
db.createCollection("files")

db.articles.createIndex({
  "title": "text",
  "summary": "text",
})

db.files.createIndex({
  "title": "text"
})
