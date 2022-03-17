package api

type Xdlb struct {
	Type      string `json:"type"`
	VideoPath struct {
		PptVideo     string `json:"pptVideo"`
		Mobile       string `json:"mobile"`
		TeacherTrack string `json:"teacherTrack"`
		StudentFull  string `json:"studentFull"`
	} `json:"videoPath"`
}
