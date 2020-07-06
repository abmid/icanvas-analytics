package entity

type Account struct {
	ID                         uint32 `json:"id"`
	Name                       string `json:"name"`
	UUID                       string `json:"uuid"`
	ParentAccountID            uint32 `json:"parrent_account_id"`
	RootAccountID              uint32 `json:"root_account_id"`
	DefaultStorageQuotaMB      uint32 `json:"default_storage_quota_mb"`
	DefaultUserStorageQuotaMB  uint32 `json:"default_user_storage_quota_mb"`
	DefaultGroupStorageQuotaMB uint32 `json:"default_group_storgae_quota_mb"`
	DefaultTimeZone            string `json:"default_time_zone"`
	SISAccountID               string `json:"sis_account_id"`
	IntegrationID              string `json:"integration_id"`
	SISImportID                uint32 `json:"sis_import_id"`
	LtiGuid                    string `json:"lti_guid"`
	WorkflowState              string `json:"workflow_state"`
}
