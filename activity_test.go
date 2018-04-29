package gdrivelist

import (
	"io/ioutil"
	"testing"

	"fmt"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

/*
func TestGDriveListFile_InvalidToken(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accessToken", "")
	//tc.SetInput("fileName", "Murder.pdf")
	//	tc.SetInput("orderBy", "createdTime")

	act.Eval(tc)

	//check result attr

	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "105")

	tc.SetInput("accessToken", "ya29.GlurBW7n5A2Fk_grpX9KMVeXLEOT4k0OhmSnF_w7626K9kgKmempF_xTDJ6uQVMkdWWWIMiNcb-ht6Rv0gcnuUb2VhtF9h7nltFw0iniwp10dmDQsFT49giOqSA8")

	act.Eval(tc)

	code = tc.GetOutput("statusCode")
	msg = tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "401")
}

func TestGDriveListFile_Success(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs



	tc.SetInput("accessToken", "ya29.GlurBT7VcVOXqc9DJjdUMPslohoMe0abAyj2dMvwgdY2-xN_jQmhCnKb4I9kyJa4hARr2tmsFngHMKJ0R-WOnhJ5Jbhi_Mc4CdVtQVk9-JTEsSUbs64rcDHwJRaO")

	act.Eval(tc)

	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "200")
}


func TestGDriveListFile_fileName(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs



	tc.SetInput("accessToken", "ya29.GlurBT7VcVOXqc9DJjdUMPslohoMe0abAyj2dMvwgdY2-xN_jQmhCnKb4I9kyJa4hARr2tmsFngHMKJ0R-WOnhJ5Jbhi_Mc4CdVtQVk9-JTEsSUbs64rcDHwJRaO")
	tc.SetInput("fileName", "CodeOfConduct.pdf")
	act.Eval(tc)

	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "200")
}


func TestGDriveListFile_InvalidfileName(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs



	tc.SetInput("accessToken", "ya29.GlurBT7VcVOXqc9DJjdUMPslohoMe0abAyj2dMvwgdY2-xN_jQmhCnKb4I9kyJa4hARr2tmsFngHMKJ0R-WOnhJ5Jbhi_Mc4CdVtQVk9-JTEsSUbs64rcDHwJRaO")
	tc.SetInput("fileName", "CodeOfConduct1.pdf")
	act.Eval(tc)

	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "200")
}
*/

func TestGDriveListFile_orderByfileName(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accessToken", "ya29.GlurBT7VcVOXqc9DJjdUMPslohoMe0abAyj2dMvwgdY2-xN_jQmhCnKb4I9kyJa4hARr2tmsFngHMKJ0R-WOnhJ5Jbhi_Mc4CdVtQVk9-JTEsSUbs64rcDHwJRaO")
	tc.SetInput("fileName", "")
	tc.SetInput("orderBy", "createdTime")

	act.Eval(tc)

	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "200")
}
