package config

type Body struct {
	ResponseCode int
	Message      string
}

type APIDteails struct {
	IicsQaPerfusw1 []struct {
		SchedulerService    []string `json:"scheduler-service,omitempty"`
		Containername       string   `json:"containername"`
		Logfilepath         string   `json:"logfilepath"`
		KmsService          []string `json:"kms-service,omitempty"`
		NotificationService []string `json:"notification-service,omitempty"`
		CAService           []string `json:"ca-service,omitempty"`
		LicenseService      []string `json:"license-service,omitempty"`
		JlsService          []string `json:"jls-service,omitempty"`
		BundleService       []string `json:"bundle-service,omitempty"`
		SessionService      []string `json:"session-service,omitempty"`
		Frs                 []string `json:"frs,omitempty"`
		AuditService        []string `json:"audit-service,omitempty"`
		PreferenceService   []string `json:"preference-service,omitempty"`
		MigrationService    []string `json:"migration-service,omitempty"`
		AdminService        []string `json:"admin-service,omitempty"`
		Vcs                 []string `json:"vcs,omitempty"`
		AcService           []string `json:"ac-service,omitempty"`
		LdmService          []string `json:"ldm-service,omitempty"`
		Runtime             []string `json:"runtime,omitempty"`
	} `json:"iics-qa-perfusw1"`
	IicsQaPerfids []struct {
		Containername     string   `json:"containername"`
		Logfilepath       string   `json:"logfilepath"`
		AuthService       []string `json:"auth-service,omitempty"`
		BrandingService   []string `json:"branding-service,omitempty"`
		ContentRepository []string `json:"content-repository,omitempty"`
		IdsService        []string `json:"ids-service,omitempty"`
		MaService         []string `json:"ma-service,omitempty"`
		ScimService       []string `json:"scim-service,omitempty"`
		V3API             []string `json:"v3api,omitempty"`
	} `json:"iics-qa-perfids"`
}

type ESResponse struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore interface{}   `json:"max_score"`
		Hits     []interface{} `json:"hits"`
	} `json:"hits"`
	Aggregations struct {
		Num0 struct {
			Values struct {
				Nine00 float64 `json:"90.0"`
				Nine51 float64 `json:"95.0"`
			} `json:"values"`
		} `json:"0"`
		Num1 struct {
			Value float64 `json:"value"`
		} `json:"1"`
	} `json:"aggregations"`
}

type Value struct {
	AverageResponseTime float64
	Responsetime95      float64
	TotalHits           int
}
