package main

import (
  "log"
  "fmt"
  "time"
  "strconv"
  "strings"
  "net/http"

  "github.com/gin-gonic/gin"
)

func startApiServer() {

  router := gin.Default()

  v1 := router.Group("/v1")
  {
    v1.POST("/feedback/:sessionId", postFeedbackRequestHandler)
    v1.GET("/feedback", getFeedbackRequestHandler)
  }

  router.Run(":8190")
}

// Inserts feedback details for a given user/session.
// Since this is a POST, feedback for the same user/session will be rejected and a 409 status will be returned.
// Example:
// curl -X "POST" "http://localhost:8190/v1/feedback/105" \
//      -H 'X-UserId: mpeterson' \
//      -H 'Content-Type: application/json' \
//      -d $'{
//   "rating": 4,
//   "comment": "foo baz ispsum",
//   "timestamp": "2018-04-16T14:15:16+00:00"
// }'
func postFeedbackRequestHandler(c *gin.Context) {

  sessionId := c.Param("sessionId")
  userId := strings.TrimSpace(c.GetHeader("X-UserId"))

  if userId == "" {
    c.String(http.StatusBadRequest, "X-UserId header is required")
    return
  }

  var feedback Feedback
	c.BindJSON(&feedback)
  feedback.SessionId = sessionId
  feedback.UserId = userId

  inserted := FeedbackDao.insertFeedback(feedback)
  if !inserted {
    c.String(http.StatusConflict, fmt.Sprintf("Feedback has already been submitted for session %s by user %s", sessionId, userId))
  } else {
    c.String(http.StatusOK, "OK")
  }
}

// Get a list of up to the latest 15 feedback entries. Optionally include the 'rating' parameter to limit results to feedback of the given rating.
// Example:
// curl "http://localhost:8190/v1/feedback?rating=4"
func getFeedbackRequestHandler(c *gin.Context) {

  ratingParam := c.DefaultQuery("rating", "0")
  rating, err := strconv.Atoi(ratingParam)
  if err != nil {
    log.Fatal(err)
  }

  feedbacks := FeedbackDao.selectFeedback(rating)

  c.JSON(http.StatusOK, feedbacks)
}

type Feedback struct {

  Timestamp time.Time `json:"timestamp"`
  SessionId string
  UserId string
  Rating int `json:"rating"`
  Comment string `json:"comment"`
}
