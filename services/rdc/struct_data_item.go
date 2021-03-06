package rdc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// DataItem is a nested struct in rdc response
type DataItem struct {
	User                  string      `json:"User" xml:"User"`
	VerifierId            int         `json:"VerifierId" xml:"VerifierId"`
	PossibleValues        string      `json:"PossibleValues" xml:"PossibleValues"`
	UserId                int         `json:"UserId" xml:"UserId"`
	PriorityId            int         `json:"PriorityId" xml:"PriorityId"`
	AttachmentIds         string      `json:"AttachmentIds" xml:"AttachmentIds"`
	ProjectIds            string      `json:"ProjectIds" xml:"ProjectIds"`
	FieldFormat           string      `json:"FieldFormat" xml:"FieldFormat"`
	LogicalStatus         string      `json:"LogicalStatus" xml:"LogicalStatus"`
	Source                string      `json:"Source" xml:"Source"`
	SkipCollab            bool        `json:"SkipCollab" xml:"SkipCollab"`
	Attachmented          bool        `json:"Attachmented" xml:"Attachmented"`
	IgnoreCheck           bool        `json:"IgnoreCheck" xml:"IgnoreCheck"`
	AkPaths               string      `json:"AkPaths" xml:"AkPaths"`
	LinePath              string      `json:"LinePath" xml:"LinePath"`
	Watched               bool        `json:"Watched" xml:"Watched"`
	VersionId             int         `json:"VersionId" xml:"VersionId"`
	StatusId              int         `json:"StatusId" xml:"StatusId"`
	PropertyType          string      `json:"PropertyType" xml:"PropertyType"`
	Type                  string      `json:"Type" xml:"Type"`
	CommentList           string      `json:"CommentList" xml:"CommentList"`
	CreatedAt             int64       `json:"CreatedAt" xml:"CreatedAt"`
	DefaultValue          string      `json:"DefaultValue" xml:"DefaultValue"`
	RecordChangeLog       bool        `json:"RecordChangeLog" xml:"RecordChangeLog"`
	AkVersionIds          string      `json:"AkVersionIds" xml:"AkVersionIds"`
	MinLength             int         `json:"MinLength" xml:"MinLength"`
	Other                 string      `json:"Other" xml:"Other"`
	TrackerIds            string      `json:"TrackerIds" xml:"TrackerIds"`
	AoneType              string      `json:"AoneType" xml:"AoneType"`
	RelatedAKProjectGuids string      `json:"RelatedAKProjectGuids" xml:"RelatedAKProjectGuids"`
	Scope                 int         `json:"Scope" xml:"Scope"`
	ModuleList            string      `json:"ModuleList" xml:"ModuleList"`
	AssignedTo            string      `json:"AssignedTo" xml:"AssignedTo"`
	OldValue              string      `json:"OldValue" xml:"OldValue"`
	BlackListNotice       string      `json:"BlackListNotice" xml:"BlackListNotice"`
	TargetId              int         `json:"TargetId" xml:"TargetId"`
	Status                string      `json:"Status" xml:"Status"`
	AttachmentList        string      `json:"AttachmentList" xml:"AttachmentList"`
	IssueTypeId           int         `json:"IssueTypeId" xml:"IssueTypeId"`
	Name                  string      `json:"Name" xml:"Name"`
	Dynamic               bool        `json:"Dynamic" xml:"Dynamic"`
	TemplateId            int         `json:"TemplateId" xml:"TemplateId"`
	VerifierStaffId       string      `json:"VerifierStaffId" xml:"VerifierStaffId"`
	RelatedUserIds        string      `json:"RelatedUserIds" xml:"RelatedUserIds"`
	DevDuration           int         `json:"DevDuration" xml:"DevDuration"`
	FullName              string      `json:"FullName" xml:"FullName"`
	CommitDate            int64       `json:"CommitDate" xml:"CommitDate"`
	WatcherIdList         string      `json:"WatcherIdList" xml:"WatcherIdList"`
	PropertyKey           string      `json:"PropertyKey" xml:"PropertyKey"`
	SendWangwang          bool        `json:"SendWangwang" xml:"SendWangwang"`
	Id                    int         `json:"Id" xml:"Id"`
	SeriousLevel          string      `json:"SeriousLevel" xml:"SeriousLevel"`
	SprintId              int         `json:"SprintId" xml:"SprintId"`
	AssignedToIdList      string      `json:"AssignedToIdList" xml:"AssignedToIdList"`
	SeriousLevelId        int         `json:"SeriousLevelId" xml:"SeriousLevelId"`
	ParentId              int         `json:"ParentId" xml:"ParentId"`
	IdPath                string      `json:"IdPath" xml:"IdPath"`
	ModuleUpdated         bool        `json:"ModuleUpdated" xml:"ModuleUpdated"`
	Identifier            string      `json:"Identifier" xml:"Identifier"`
	UserStaffId           string      `json:"UserStaffId" xml:"UserStaffId"`
	FixedUserId           int         `json:"FixedUserId" xml:"FixedUserId"`
	VersionList           string      `json:"VersionList" xml:"VersionList"`
	AssignedToId          int         `json:"AssignedToId" xml:"AssignedToId"`
	RespondDuration       int         `json:"RespondDuration" xml:"RespondDuration"`
	TargetType            string      `json:"TargetType" xml:"TargetType"`
	IsDelete              bool        `json:"IsDelete" xml:"IsDelete"`
	StatusStage           int         `json:"StatusStage" xml:"StatusStage"`
	Editable              bool        `json:"Editable" xml:"Editable"`
	ModuleIds             string      `json:"ModuleIds" xml:"ModuleIds"`
	AkProjectId           int         `json:"AkProjectId" xml:"AkProjectId"`
	SourceId              int         `json:"SourceId" xml:"SourceId"`
	ChangeLogList         string      `json:"ChangeLogList" xml:"ChangeLogList"`
	MaxLength             int         `json:"MaxLength" xml:"MaxLength"`
	IsRemember            bool        `json:"IsRemember" xml:"IsRemember"`
	RelatedAKProjectIds   string      `json:"RelatedAKProjectIds" xml:"RelatedAKProjectIds"`
	FixedDuration         int         `json:"FixedDuration" xml:"FixedDuration"`
	VersionIds            string      `json:"VersionIds" xml:"VersionIds"`
	IssueRelations        string      `json:"IssueRelations" xml:"IssueRelations"`
	Stamp                 string      `json:"Stamp" xml:"Stamp"`
	ProjectId             int         `json:"ProjectId" xml:"ProjectId"`
	NameI18N              string      `json:"NameI18N" xml:"NameI18N"`
	Region                string      `json:"Region" xml:"Region"`
	IgnoreIntegrate       bool        `json:"IgnoreIntegrate" xml:"IgnoreIntegrate"`
	Verifier              string      `json:"Verifier" xml:"Verifier"`
	AssignedToIds         string      `json:"AssignedToIds" xml:"AssignedToIds"`
	Mode                  string      `json:"Mode" xml:"Mode"`
	Subject               string      `json:"Subject" xml:"Subject"`
	IsRequired            bool        `json:"IsRequired" xml:"IsRequired"`
	TestsuiteId           int         `json:"TestsuiteId" xml:"TestsuiteId"`
	UpdatedAt             int64       `json:"UpdatedAt" xml:"UpdatedAt"`
	AssignedToStaffId     string      `json:"AssignedToStaffId" xml:"AssignedToStaffId"`
	Description           string      `json:"Description" xml:"Description"`
	ClosedDuration        int         `json:"ClosedDuration" xml:"ClosedDuration"`
	ScopeUserIds          string      `json:"ScopeUserIds" xml:"ScopeUserIds"`
	Priority              string      `json:"Priority" xml:"Priority"`
	AoneId                int         `json:"AoneId" xml:"AoneId"`
	Trackers              string      `json:"Trackers" xml:"Trackers"`
	Solution              int         `json:"Solution" xml:"Solution"`
	NewValue              string      `json:"NewValue" xml:"NewValue"`
	TagIdList             string      `json:"TagIdList" xml:"TagIdList"`
	Icons                 []string    `json:"Icons" xml:"Icons"`
	CustomFieldMap        []string    `json:"CustomFieldMap" xml:"CustomFieldMap"`
	Users                 []UsersItem `json:"Users" xml:"Users"`
}
