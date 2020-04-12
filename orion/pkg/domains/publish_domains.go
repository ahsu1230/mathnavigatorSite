package domains

type UnpublishedDomains struct {
	Programs  []Program  `json:"programs"`
	Classes   []Class    `json:"classes"`
	Locations []Location `json:"locations"`
	Achieves  []Achieve  `json:"achieves"`
	Semesters []Semester `json:"semesters"`
	Sessions  []Session  `json:"sessions"`
}

type PublishErrorBody struct {
	RowId    uint   `json:"rowId,omitempty"`
	StringId string `json:"stringId,omitempty"`
	Error    error  `json:"error"`
}
