package domains

type ProgramClass struct {
	ProgramObj Program `json:"program"`
	Classes    []Class `json:"classes"`
}

type ProgramClassesBySemester struct {
	Semester       Semester       `json:"semester"`
	ProgramClasses []ProgramClass `json:"programClasses"`
}
