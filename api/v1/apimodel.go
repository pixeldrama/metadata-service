// Code generated by OpenAPI Generator (https://openapi-generator.tech); then postprocessed; DO NOT EDIT.

package openapi

import "time"

type DeletionDto struct {
	// The jira issue to use for committing the deletion.
	JiraIssue string `json:"jiraIssue"`
}

type ErrorDto struct {
	Details   *string    `json:"details,omitempty"`
	Message   *string    `json:"message,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

type HealthComponent struct {
	Description *string `json:"description,omitempty"`
	Status      *string `json:"status,omitempty"`
}

type OwnerCreateDto struct {
	// The contact information of the owner
	Contact string `json:"contact"`
	// The product owner of this owner space
	ProductOwner *string `json:"productOwner,omitempty"`
	// The default jira project that is used by this owner space
	DefaultJiraProject *string `json:"defaultJiraProject,omitempty"`
	// The jira issue to use for committing a change, or the last jira issue used.
	JiraIssue string `json:"jiraIssue"`
}

type OwnerDto struct {
	// The contact information of the owner
	Contact string `json:"contact" yaml:"contact"`
	// The product owner of this owner space
	ProductOwner *string `json:"productOwner,omitempty" yaml:"productOwner,omitempty"`
	// The default jira project that is used by this owner space
	DefaultJiraProject *string `json:"defaultJiraProject,omitempty" yaml:"defaultJiraProject,omitempty"`
	// ISO-8601 UTC date time at which this information was originally committed. When sending an update, include the original timestamp you got so we can detect concurrent updates.
	TimeStamp string `json:"timeStamp" yaml:"-"`
	// The git commit hash this information was originally committed under. When sending an update, include the original commitHash you got so we can detect concurrent updates.
	CommitHash string `json:"commitHash" yaml:"-"`
	// The jira issue to use for committing a change, or the last jira issue used.
	JiraIssue string `json:"jiraIssue" yaml:"-"`
}

type OwnerListDto struct {
	Owners map[string]OwnerDto `json:"owners"`
	// ISO-8601 UTC date time at which the list of owners was obtained from service-metadata
	TimeStamp string `json:"timeStamp"`
}

type OwnerPatchDto struct {
	// The contact information of the owner
	Contact *string `json:"contact,omitempty"`
	// The product owner of this owner space
	ProductOwner *string `json:"productOwner,omitempty"`
	// The default jira project that is used by this owner space
	DefaultJiraProject *string `json:"defaultJiraProject,omitempty"`
	// ISO-8601 UTC date time at which this information was originally committed. When sending an update, include the original timestamp you got so we can detect concurrent updates.
	TimeStamp string `json:"timeStamp"`
	// The git commit hash this information was originally committed under. When sending an update, include the original commitHash you got so we can detect concurrent updates.
	CommitHash string `json:"commitHash"`
	// The jira issue to use for committing a change, or the last jira issue used.
	JiraIssue string `json:"jiraIssue"`
}

type Quicklink struct {
	Url         *string `json:"url,omitempty" yaml:"url,omitempty"`
	Title       *string `json:"title,omitempty" yaml:"title,omitempty"`
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
}

type RepositoryConfigurationAccessKeyDto struct {
	Key        string  `json:"key" yaml:"key"`
	Permission *string `json:"permission,omitempty" yaml:"permission,omitempty"`
}

type RepositoryConfigurationDto struct {
	// Ssh-Keys configured on the repository.
	AccessKeys []RepositoryConfigurationAccessKeyDto `json:"accessKeys,omitempty" yaml:"accessKeys,omitempty"`
	// Adds a corresponding commit message regex.
	CommitMessageType *string `json:"commitMessageType,omitempty" yaml:"commitMessageType,omitempty"`
	// Configures JQL matcher with query: issuetype in (Story, Bug) AND 'Risk Level' is not EMPTY
	RequireIssue *bool `json:"requireIssue,omitempty" yaml:"requireIssue,omitempty"`
	// Set the required successful builds counter.
	RequireSuccessfulBuilds *int32                              `json:"requireSuccessfulBuilds,omitempty" yaml:"requireSuccessfulBuilds,omitempty"`
	Webhooks                *RepositoryConfigurationWebhooksDto `json:"webhooks,omitempty" yaml:"webhooks,omitempty"`
	// Map of string (group name e.g. some-owner) of strings (list of approvers), one approval for each group is required.
	Approvers        *map[string][]string `json:"approvers,omitempty" yaml:"approvers,omitempty"`
	DefaultReviewers []string             `json:"defaultReviewers,omitempty" yaml:"defaultReviewers,omitempty"`
	// List of users, who can sign a pull request.
	SignedApprovers []string `json:"signedApprovers,omitempty" yaml:"signedApprovers,omitempty"`
}

type RepositoryConfigurationWebhookDto struct {
	Name string `json:"name" yaml:"name"`
	Url  string `json:"url" yaml:"url"`
	// Events the webhook should be triggered with.
	Events        []string           `json:"events,omitempty" yaml:"events,omitempty"`
	Configuration *map[string]string `json:"configuration,omitempty" yaml:"configuration,omitempty"`
}

type RepositoryConfigurationWebhooksDto struct {
	// Default pipeline trigger webhook.
	PipelineTrigger *bool `json:"pipelineTrigger,omitempty" yaml:"pipelineTrigger,omitempty"`
	// Additional webhooks to be configured.
	Additional []RepositoryConfigurationWebhookDto `json:"additional,omitempty" yaml:"additional,omitempty"`
}

type RepositoryCreateDto struct {
	// The alias of the repository owner
	Owner    string `json:"owner"`
	Url      string `json:"url"`
	Mainline string `json:"mainline"`
	// the generator used for the initial contents of this repository
	Generator *string `json:"generator,omitempty"`
	// this repository contains unit tests (currently ignored except for helm charts)
	Unittest      *bool                       `json:"unittest,omitempty"`
	Configuration *RepositoryConfigurationDto `json:"configuration,omitempty"`
	// The jira issue to use for committing a change, or the last jira issue used.
	JiraIssue string `json:"jiraIssue"`
}

type RepositoryDto struct {
	// The alias of the repository owner
	Owner    string `json:"owner" yaml:"-"`
	Url      string `json:"url" yaml:"url"`
	Mainline string `json:"mainline" yaml:"mainline"`
	// the generator used for the initial contents of this repository
	Generator *string `json:"generator,omitempty" yaml:"generator,omitempty"`
	// this repository contains unit tests (currently ignored except for helm charts)
	Unittest      *bool                       `json:"unittest,omitempty" yaml:"unittest,omitempty"`
	Configuration *RepositoryConfigurationDto `json:"configuration,omitempty" yaml:"configuration,omitempty"`
	// ISO-8601 UTC date time at which this information was originally committed. When sending an update, include the original timestamp you got so we can detect concurrent updates.
	TimeStamp string `json:"timeStamp" yaml:"-"`
	// The git commit hash this information was originally committed under. When sending an update, include the original commitHash you got so we can detect concurrent updates.
	CommitHash string `json:"commitHash" yaml:"-"`
	// The jira issue to use for committing a change, or the last jira issue used.
	JiraIssue string `json:"jiraIssue" yaml:"-"`
}

type RepositoryListDto struct {
	Repositories map[string]RepositoryDto `json:"repositories"`
	// ISO-8601 UTC date time at which the list of repositories was obtained from service-metadata
	TimeStamp string `json:"timeStamp"`
}

type RepositoryPatchDto struct {
	// The alias of the repository owner
	Owner    *string `json:"owner,omitempty"`
	Url      *string `json:"url,omitempty"`
	Mainline *string `json:"mainline,omitempty"`
	// the generator used for the initial contents of this repository
	Generator *string `json:"generator,omitempty"`
	// this repository contains unit tests (currently ignored except for helm charts)
	Unittest      *bool                       `json:"unittest,omitempty"`
	Configuration *RepositoryConfigurationDto `json:"configuration,omitempty"`
	// ISO-8601 UTC date time at which this information was originally committed. When sending an update, include the original timestamp you got so we can detect concurrent updates.
	TimeStamp string `json:"timeStamp"`
	// The git commit hash this information was originally committed under. When sending an update, include the original commitHash you got so we can detect concurrent updates.
	CommitHash string `json:"commitHash"`
	// The jira issue to use for committing a change, or the last jira issue used.
	JiraIssue string `json:"jiraIssue"`
}

type ServiceCreateDto struct {
	// The alias of the service owner. Note, an update with changed owner will move the service and any associated repositories to the new owner, but of course this will not move e.g. Jenkins jobs. That's your job.
	Owner string `json:"owner"`
	// A list of quicklinks related to the service
	Quicklinks []Quicklink `json:"quicklinks"`
	// The keys of repositories associated with the service. When sending an update, they must refer to repositories that belong to this service, or the update will fail
	Repositories []string `json:"repositories"`
	// The default channel used to send any alerts of the service to. Can be an email address or a Teams webhook URL
	AlertTarget string `json:"alertTarget"`
	// True for services that will be permanently deployed to the Development environment only.
	DevelopmentOnly *bool `json:"developmentOnly,omitempty"`
	// The operation type of the service. 'WORKLOAD' follows the default deployment strategy of one instance per environment, 'PLATFORM' one instance per cluster or node.
	OperationType *string `json:"operationType,omitempty"`
	// The security scans that are required for this service. Optional, SAST and/or SCA.
	RequiredScans []string `json:"requiredScans,omitempty"`
	// The jira issue to use for committing a change, or the last jira issue used.
	JiraIssue string `json:"jiraIssue"`
}

type ServiceDto struct {
	// The alias of the service owner. Note, an update with changed owner will move the service and any associated repositories to the new owner, but of course this will not move e.g. Jenkins jobs. That's your job.
	Owner string `json:"owner" yaml:"-"`
	// A list of quicklinks related to the service
	Quicklinks []Quicklink `json:"quicklinks" yaml:"quicklinks"`
	// The keys of repositories associated with the service. When sending an update, they must refer to repositories that belong to this service, or the update will fail
	Repositories []string `json:"repositories" yaml:"repositories"`
	// The default channel used to send any alerts of the service to. Can be an email address or a Teams webhook URL
	AlertTarget string `json:"alertTarget" yaml:"alertTarget"`
	// True for services that will be permanently deployed to the Development environment only.
	DevelopmentOnly *bool `json:"developmentOnly,omitempty" yaml:"developmentOnly,omitempty"`
	// The operation type of the service. 'WORKLOAD' follows the default deployment strategy of one instance per environment, 'PLATFORM' one instance per cluster or node.
	OperationType *string `json:"operationType,omitempty" yaml:"operationType,omitempty"`
	// The security scans that are required for this service. Optional, SAST and/or SCA.
	RequiredScans []string `json:"requiredScans,omitempty" yaml:"requiredScans,omitempty"`
	// ISO-8601 UTC date time at which this information was originally committed. When sending an update, include the original timestamp you got so we can detect concurrent updates.
	TimeStamp string `json:"timeStamp" yaml:"-"`
	// The git commit hash this information was originally committed under. When sending an update, include the original commitHash you got so we can detect concurrent updates.
	CommitHash string `json:"commitHash" yaml:"-"`
	// The jira issue to use for committing a change, or the last jira issue used.
	JiraIssue string `json:"jiraIssue" yaml:"-"`
}

type ServiceListDto struct {
	Services map[string]ServiceDto `json:"services"`
	// ISO-8601 UTC date time at which the list of services was obtained from service-metadata
	TimeStamp string `json:"timeStamp"`
}

type ServicePatchDto struct {
	// The alias of the service owner. Note, a patch with changed owner will move the service and any associated repositories to the new owner, but of course this will not move e.g. Jenkins jobs. That's your job.
	Owner *string `json:"owner,omitempty"`
	// A list of quicklinks related to the service
	Quicklinks []Quicklink `json:"quicklinks,omitempty"`
	// The keys of repositories associated with the service. When sending an update, they must refer to repositories that belong to this service, or the update will fail
	Repositories []string `json:"repositories,omitempty"`
	// The default channel used to send any alerts of the service to. Can be an email address or a Teams webhook URL
	AlertTarget *string `json:"alertTarget,omitempty"`
	// True for services that will be permanently deployed to the Development environment only.
	DevelopmentOnly *bool `json:"developmentOnly,omitempty"`
	// The operation type of the service. 'WORKLOAD' follows the default deployment strategy of one instance per environment, 'PLATFORM' one instance per cluster or node.
	OperationType *string `json:"operationType,omitempty"`
	// The security scans that are required for this service. Optional, SAST and/or SCA.
	RequiredScans []string `json:"requiredScans,omitempty"`
	// ISO-8601 UTC date time at which this information was originally committed. When sending an update, include the original timestamp you got so we can detect concurrent updates.
	TimeStamp string `json:"timeStamp"`
	// The git commit hash this information was originally committed under. When sending an update, include the original commitHash you got so we can detect concurrent updates.
	CommitHash string `json:"commitHash"`
	// The jira issue to use for committing a change, or the last jira issue used.
	JiraIssue string `json:"jiraIssue"`
}
