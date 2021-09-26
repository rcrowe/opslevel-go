// Code generated by gen.go; DO NOT EDIT.

package opslevel

// CheckStatus represents the evaluation status of the check.
type CheckStatus string

const (
	CheckStatusPassed  CheckStatus = "passed"  // The check evaluated to a truthy value based on some conditions.
	CheckStatusFailed  CheckStatus = "failed"  // The check evaluated to a falsy value based on some conditions.
	CheckStatusPending CheckStatus = "pending" // The check has not been evaluated yet..
)

// All CheckStatus as []string
func AllCheckStatus() []string {
	return []string{
		string(CheckStatusPassed),
		string(CheckStatusFailed),
		string(CheckStatusPending),
	}
}

// CheckType represents the type of check.
type CheckType string

const (
	CheckTypeHasOwner         CheckType = "has_owner"          // Verifies that the service has an owner defined.
	CheckTypeServiceProperty  CheckType = "service_property"   // Verifies that a service property is set or matches a specified format.
	CheckTypeHasServiceConfig CheckType = "has_service_config" // Verifies that the service is maintained though the use of an opslevel.yml service config.
	CheckTypeHasRepository    CheckType = "has_repository"     // Verifies that the service has a repository integrated.
	CheckTypeToolUsage        CheckType = "tool_usage"         // Verifies that the service is using a tool of a particular category or name.
	CheckTypeTagDefined       CheckType = "tag_defined"        // Verifies that the service has the specified tag defined.
	CheckTypeRepoFile         CheckType = "repo_file"          // Verifies that the service's repository contains a file with a certain path. (Optional: Can also be used to check for specific file contents).
	CheckTypeRepoSearch       CheckType = "repo_search"        // Searches the service's repository and verifies if any file matches the given contents.
	CheckTypeCustom           CheckType = "custom"             // Allows for the creation of programmatic checks that use an API to mark the status as passing or failing.
	CheckTypePayload          CheckType = "payload"            // Requires a payload integration api call to complete a check for the service.
	CheckTypeManual           CheckType = "manual"             // Requires a service owner to manually complete a check for the service.
	CheckTypeGeneric          CheckType = "generic"            // Requires a generic integration api call to complete a series of checks for multiple services.
)

// All CheckType as []string
func AllCheckType() []string {
	return []string{
		string(CheckTypeHasOwner),
		string(CheckTypeServiceProperty),
		string(CheckTypeHasServiceConfig),
		string(CheckTypeHasRepository),
		string(CheckTypeToolUsage),
		string(CheckTypeTagDefined),
		string(CheckTypeRepoFile),
		string(CheckTypeRepoSearch),
		string(CheckTypeCustom),
		string(CheckTypePayload),
		string(CheckTypeManual),
		string(CheckTypeGeneric),
	}
}

// ConnectiveEnum represents the logical operator to be used in conjunction with multiple filters (requires filters to be supplied).
type ConnectiveEnum string

const (
	ConnectiveEnumAnd ConnectiveEnum = "and" // Used to ensure **all** filters match for a given resource.
	ConnectiveEnumOr  ConnectiveEnum = "or"  // Used to ensure **any** filters match for a given resource.
)

// All ConnectiveEnum as []string
func AllConnectiveEnum() []string {
	return []string{
		string(ConnectiveEnumAnd),
		string(ConnectiveEnumOr),
	}
}

// ContactType represents the method of contact.
type ContactType string

const (
	ContactTypeSlack ContactType = "slack" // A Slack channel contact method.
	ContactTypeEmail ContactType = "email" // An email contact method.
	ContactTypeWeb   ContactType = "web"   // A website contact method.
)

// All ContactType as []string
func AllContactType() []string {
	return []string{
		string(ContactTypeSlack),
		string(ContactTypeEmail),
		string(ContactTypeWeb),
	}
}

// FrequencyTimeScale represents the time scale type for the frequency.
type FrequencyTimeScale string

const (
	FrequencyTimeScaleDay   FrequencyTimeScale = "day"   // Consider the time scale of days.
	FrequencyTimeScaleWeek  FrequencyTimeScale = "week"  // Consider the time scale of weeks.
	FrequencyTimeScaleMonth FrequencyTimeScale = "month" // Consider the time scale of months.
	FrequencyTimeScaleYear  FrequencyTimeScale = "year"  // Consider the time scale of years.
)

// All FrequencyTimeScale as []string
func AllFrequencyTimeScale() []string {
	return []string{
		string(FrequencyTimeScaleDay),
		string(FrequencyTimeScaleWeek),
		string(FrequencyTimeScaleMonth),
		string(FrequencyTimeScaleYear),
	}
}

// PredicateKeyEnum represents fields that can be used as part of filter for services.
type PredicateKeyEnum string

const (
	PredicateKeyEnumTierIndex      PredicateKeyEnum = "tier_index"      // Filter by `tier` field.
	PredicateKeyEnumLifecycleIndex PredicateKeyEnum = "lifecycle_index" // Filter by `lifecycle` field.
	PredicateKeyEnumLanguage       PredicateKeyEnum = "language"        // Filter by `language` field.
	PredicateKeyEnumFramework      PredicateKeyEnum = "framework"       // Filter by `framework` field.
	PredicateKeyEnumProduct        PredicateKeyEnum = "product"         // Filter by `product` field.
	PredicateKeyEnumName           PredicateKeyEnum = "name"            // Filter by `name` field.
	PredicateKeyEnumTags           PredicateKeyEnum = "tags"            // Filter by `tags` field.
	PredicateKeyEnumOwnerID        PredicateKeyEnum = "owner_id"        // Filter by `owner` field.
)

// All PredicateKeyEnum as []string
func AllPredicateKeyEnum() []string {
	return []string{
		string(PredicateKeyEnumTierIndex),
		string(PredicateKeyEnumLifecycleIndex),
		string(PredicateKeyEnumLanguage),
		string(PredicateKeyEnumFramework),
		string(PredicateKeyEnumProduct),
		string(PredicateKeyEnumName),
		string(PredicateKeyEnumTags),
		string(PredicateKeyEnumOwnerID),
	}
}

// PredicateTypeEnum represents operations that can be used on predicates.
type PredicateTypeEnum string

const (
	PredicateTypeEnumContains                   PredicateTypeEnum = "contains"                     // Contains a specific value.
	PredicateTypeEnumDoesNotContain             PredicateTypeEnum = "does_not_contain"             // Does not contain a specific value.
	PredicateTypeEnumDoesNotEqual               PredicateTypeEnum = "does_not_equal"               // Does not equal a specific value.
	PredicateTypeEnumDoesNotExist               PredicateTypeEnum = "does_not_exist"               // Specific attribute does not exist.
	PredicateTypeEnumEndsWith                   PredicateTypeEnum = "ends_with"                    // Ends with a specific value.
	PredicateTypeEnumEquals                     PredicateTypeEnum = "equals"                       // Equals a specific value.
	PredicateTypeEnumExists                     PredicateTypeEnum = "exists"                       // Specific attribute exists.
	PredicateTypeEnumGreaterThanOrEqualTo       PredicateTypeEnum = "greater_than_or_equal_to"     // Greater than or equal to a specific value (numeric only).
	PredicateTypeEnumLessThanOrEqualTo          PredicateTypeEnum = "less_than_or_equal_to"        // Less than or equal to a specific value (numeric only).
	PredicateTypeEnumStartsWith                 PredicateTypeEnum = "starts_with"                  // Starts with a specific value.
	PredicateTypeEnumSatisfiesVersionConstraint PredicateTypeEnum = "satisfies_version_constraint" // Satisfies version constraint (tag value only).
	PredicateTypeEnumMatchesRegex               PredicateTypeEnum = "matches_regex"                // Matches a value using a regular expression.
	PredicateTypeEnumSatisfiesJqExpression      PredicateTypeEnum = "satisfies_jq_expression"      // Satisfies an expression defined in jq.
)

// All PredicateTypeEnum as []string
func AllPredicateTypeEnum() []string {
	return []string{
		string(PredicateTypeEnumContains),
		string(PredicateTypeEnumDoesNotContain),
		string(PredicateTypeEnumDoesNotEqual),
		string(PredicateTypeEnumDoesNotExist),
		string(PredicateTypeEnumEndsWith),
		string(PredicateTypeEnumEquals),
		string(PredicateTypeEnumExists),
		string(PredicateTypeEnumGreaterThanOrEqualTo),
		string(PredicateTypeEnumLessThanOrEqualTo),
		string(PredicateTypeEnumStartsWith),
		string(PredicateTypeEnumSatisfiesVersionConstraint),
		string(PredicateTypeEnumMatchesRegex),
		string(PredicateTypeEnumSatisfiesJqExpression),
	}
}

// ServicePropertyTypeEnum represents properties of services that can be validated.
type ServicePropertyTypeEnum string

const (
	ServicePropertyTypeEnumDescription    ServicePropertyTypeEnum = "description"     // The description of a service.
	ServicePropertyTypeEnumName           ServicePropertyTypeEnum = "name"            // The name of a service.
	ServicePropertyTypeEnumLanguage       ServicePropertyTypeEnum = "language"        // The primary programming language of a service.
	ServicePropertyTypeEnumFramework      ServicePropertyTypeEnum = "framework"       // The primary software development framework of a service.
	ServicePropertyTypeEnumProduct        ServicePropertyTypeEnum = "product"         // The product that is associated with a service.
	ServicePropertyTypeEnumLifecycleIndex ServicePropertyTypeEnum = "lifecycle_index" // The index of the lifecycle a service belongs to.
	ServicePropertyTypeEnumTierIndex      ServicePropertyTypeEnum = "tier_index"      // The index of the tier a service belongs to.
)

// All ServicePropertyTypeEnum as []string
func AllServicePropertyTypeEnum() []string {
	return []string{
		string(ServicePropertyTypeEnumDescription),
		string(ServicePropertyTypeEnumName),
		string(ServicePropertyTypeEnumLanguage),
		string(ServicePropertyTypeEnumFramework),
		string(ServicePropertyTypeEnumProduct),
		string(ServicePropertyTypeEnumLifecycleIndex),
		string(ServicePropertyTypeEnumTierIndex),
	}
}

// TaggableResource represents possible types to apply tags to.
type TaggableResource string

const (
	TaggableResourceService    TaggableResource = "Service"    // Used to identify a Service.
	TaggableResourceRepository TaggableResource = "Repository" // Used to identify a Repository.
)

// All TaggableResource as []string
func AllTaggableResource() []string {
	return []string{
		string(TaggableResourceService),
		string(TaggableResourceRepository),
	}
}

// ToolCategory represents the specific categories that a tool can belong to.
type ToolCategory string

const (
	ToolCategoryAdmin                 ToolCategory = "admin"                  // Tools used for administrative purposes.
	ToolCategoryAPIDocumentation      ToolCategory = "api_documentation"      // Tools used as API documentation for this service.
	ToolCategoryCode                  ToolCategory = "code"                   // Tools used for source code.
	ToolCategoryContinuousIntegration ToolCategory = "continuous_integration" // Tools used for building/unit testing a service.
	ToolCategoryDeployment            ToolCategory = "deployment"             // Tools used for deploying changes to a service.
	ToolCategoryErrors                ToolCategory = "errors"                 // Tools used for tracking/reporting errors.
	ToolCategoryFeatureFlag           ToolCategory = "feature_flag"           // Tools used for managing feature flags.
	ToolCategoryHealthChecks          ToolCategory = "health_checks"          // Tools used for tracking/reporting the health of a service.
	ToolCategoryIncidents             ToolCategory = "incidents"              // Tools used to surface incidents on a service.
	ToolCategoryLogs                  ToolCategory = "logs"                   // Tools used for displaying logs from services.
	ToolCategoryMetrics               ToolCategory = "metrics"                // Tools used for tracking/reporting service metrics.
	ToolCategoryOrchestrator          ToolCategory = "orchestrator"           // Tools used for orchestrating a service.
	ToolCategoryRunbooks              ToolCategory = "runbooks"               // Tools used for managing runbooks for a service.
	ToolCategoryStatusPage            ToolCategory = "status_page"            // Tools used for reporting the status of a service.
	ToolCategoryWiki                  ToolCategory = "wiki"                   // Tools used as a wiki for this service.
	ToolCategoryOther                 ToolCategory = "other"                  // Tools that do not fit into the available categories.
)

// All ToolCategory as []string
func AllToolCategory() []string {
	return []string{
		string(ToolCategoryAdmin),
		string(ToolCategoryAPIDocumentation),
		string(ToolCategoryCode),
		string(ToolCategoryContinuousIntegration),
		string(ToolCategoryDeployment),
		string(ToolCategoryErrors),
		string(ToolCategoryFeatureFlag),
		string(ToolCategoryHealthChecks),
		string(ToolCategoryIncidents),
		string(ToolCategoryLogs),
		string(ToolCategoryMetrics),
		string(ToolCategoryOrchestrator),
		string(ToolCategoryRunbooks),
		string(ToolCategoryStatusPage),
		string(ToolCategoryWiki),
		string(ToolCategoryOther),
	}
}

// UserRole represents a role that can be assigned to a user.
type UserRole string

const (
	UserRoleUser  UserRole = "user"  // A regular user on the account.
	UserRoleAdmin UserRole = "admin" // An administrator on the account.
)

// All UserRole as []string
func AllUserRole() []string {
	return []string{
		string(UserRoleUser),
		string(UserRoleAdmin),
	}
}
