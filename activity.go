package gdrivelist

import (
	s "strings"

	"github.com/DipeshTest/allstarsshared/GDrive"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

var log = logger.GetLogger("activity-gdrivelist")

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	accessToken := s.TrimSpace(context.GetInput("accessToken").(string))
	fileName := s.TrimSpace(context.GetInput("fileName").(string))
	orderBy := s.TrimSpace(context.GetInput("orderBy").(string))
	pageSize := context.GetInput("pageSize").(int)
	nextPageToken := s.TrimSpace(context.GetInput("nextPageToken").(string))
	timeout := s.TrimSpace(context.GetInput("timeout").(string))

	if len(s.TrimSpace(accessToken)) == 0 {
		//	fmt.Println(Twilio.ResponseData{501, "Failed", "Send input numbers back", "SMS body is blank"})
		context.SetOutput("statusCode", "105")

		context.SetOutput("message", "Access Token field is blank")
		//context.SetOutput("failedNumbers", to)

		//respond with this
	} else {
		if pageSize == 0 {
			pageSize = 50
		}

		if len(s.TrimSpace(timeout)) == 0 {
			timeout = "120"
		}

		code, msg, size, pageToken := GDrive.ListFile(accessToken, fileName, orderBy, int64(pageSize), nextPageToken, timeout)
		context.SetOutput("statusCode", code)

		context.SetOutput("message", msg)

		context.SetOutput("fileCount", size)

		context.SetOutput("nextPageToken", pageToken)
	}

	return true, err
}
