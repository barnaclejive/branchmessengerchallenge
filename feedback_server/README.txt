# run server on port 8190
go run *.go

# usage

Inserts feedback details for a given user/session.
Since this is a POST, feedback for the same user/session will be rejected and a 409 status will be returned.
Example:
curl -X "POST" "http://localhost:8190/v1/feedback/105" \
     -H 'X-UserId: mpeterson' \
     -H 'Content-Type: application/json' \
     -d $'{
  "rating": 4,
  "comment": "foo baz ispsum",
  "timestamp": "2018-04-16T14:15:16+00:00"
}'


Get a list of up to the latest 15 feedback entries. Optionally include the 'rating' parameter to limit results to feedback of the given rating.
Example:
curl "http://localhost:8190/v1/feedback?rating=4"
