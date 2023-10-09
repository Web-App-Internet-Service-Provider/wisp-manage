package sms

import (
	"database/sql"
	"os"

	"github.com/AbdulrahmanDaud10/wisp-manage/internal/repository"
)

var (
	TililApiKeySetting   = os.Getenv("TILIL_API_KEY_BULK_SMS_SETTING_LABEL")
	SenderIDSetting      = os.Getenv("TILIL_SENDER_ID_SETTING_LABEL")
	TililBulkSMSEndpoint = os.Getenv("TILIL_SMS_ENDPOINT_SETTING_LABEL")
)

type TililSettings struct {
	APIkey   string `json:"api_key"`
	SenderID string `json:"sender_id"`
}

// Handler for GET /settings/tilil, which retrieves settings required
// to communicate with TILIL API
func GetTililSettings() (TililSettings, error) {
	var tililSettings TililSettings
	var err error

	ts := repository.NewOrganizationSettingRepository()

	if tililSettings.APIkey, err = ts.GetSettingValue(TililApiKeySetting); err != nil {
		if err != sql.ErrNoRows {
			return tililSettings, err
		}
	}

	if tililSettings.SenderID, err = ts.GetSettingValue(SenderIDSetting); err != nil {
		if err != sql.ErrNoRows {
			return tililSettings, err
		}
	}

	return tililSettings, nil
}
