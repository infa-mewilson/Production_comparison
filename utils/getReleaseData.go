package utils

import (
	"NEWGOLANG/config"
	"encoding/json"
	"fmt"
)

func ReleaseData(StartDate string, EndDate string, ids string, pod1 string, environment string) (map[string]config.Value, map[string]config.Value) {
	perfusw1map := make(map[string]config.Value)
	perfidsmap := make(map[string]config.Value)
	listdetails := `{
  "iics-qa-perfusw1":[
     {
      "scheduler-service":["POST /scheduler-service/api/v1/Schedules","PATCH /scheduler-service/api/v1/Schedules","GET /scheduler-service/internal/api/v1/Jobs","GET scheduler-service/api/v1/Organizations *and* /Jobs","POST /scheduler-service/api/v1/Schedules *and* /Jobs","PUT /scheduler-service/api/v1/Schedules","GET /scheduler-service/api/v1/Schedules","DELETE /scheduler-service/api/v1/Jobs","GET /scheduler-service/api/v1/Jobs","DELETE /scheduler-service/api/v1/Schedules","GET /scheduler-service/api/v1/Organizations *and* /Schedules"],
      "containername":"scheduler-service",
      "logfilepath":"*/var/log/containers/scheduler*"
    },
    {
      "kms-service":["POST /kms-service/internal/api/v1/EncryptWithDataKey","POST /kms-service/internal/api/v1/DecryptWithMasterKey"],
      "containername":"kms-service",
      "logfilepath":"*/var/log/containers/kms*"
    },
    {
        "notification-service":["GET /notification-service/api/v1/Users( *and* )/ReceivedMessages?$count=true&$filter=status%20eq%20%27NEW%27%20and%20(messageType%20eq%20%27BELL_NOTIFICATION%27)&$top=0","POST /notification-service/internal/api/v1/Topics('user *and* )/Messages","POST /notification-service/internal/api/v1/Topics('org *and* )/Messages","POST /notification-service/api/v1/MarkAllNotificationsAsRead"],
        "containername":"notification-service",
        "logfilepath":"*/var/log/containers/notification-service*"
    },
    {
        "ca-service":["POST /ca-service/api/v1/Sign"],
        "containername":"ca-service",
        "logfilepath":"*/var/log/containers/ca-service*"
    },
    {   "license-service":["POST /license-service/podlink/api/v1/OrgLicenseAssignment( *and* )/AssignEdition()","GET /license-service/api/v1/OrgLicenseAssignment( *and* )/View()","GET /license-service/api/v1/OrgLicenseAssignment( *and* )/EffectiveLicenses()","GET /license-service/internal/api/v1/OrgLicenseAssignment( *and* )/EffectiveAppExternalId()","GET /license-service/api/v1/metering/definitions","POST license-service/api/v1/metering/org/( *and* )/usage"],
        "containername":"license-service",
        "logfilepath":"*/var/log/containers/license-service*"
    },
    {
        "jls-service":["GET /jls-di/api/v1/Orgs( *and* )/JobLogEntries","GET /jls-di/api/v1/JobStatusChart","GET /jls-di/api/v1/Orgs( *and* )/FlattenedJobLogEntries(","GET /jls-di/api/v1/Orgs( *and* )/JobLogEntries?$count=true&$top=0&$filter=startedBy","GET /jls-di/api/v1/Orgs( *and* )/JobLogEntries?$skip=0&$top=25&$filter=startedBy"],
        "containername":"jls-di",
        "logfilepath":"*/var/log/containers/jls-di*"
    },
    {
        "bundle-service":["GET /bundle-service/api/v1/BundleInstallations"],
        "containername":"bundle-service",
        "logfilepath":"*/var/log/containers/bundle-service*"
    },
    {
        "session-service":["GET /session-service/api/v1/session/User","GET /session-service/api/v1/session/Agent","GET /session-service/internal/api/v1/cache/ServiceAndUserInfo","GET /session-service/api/v1/Orgs( *and* )/Parent","DELETE /session-service/api/v1/session/Logout"],
        "containername":"session-service",
        "logfilepath":"*/var/log/containers/session-service*"
    },
    {
        "frs":["GET /frs/internal/api/v1/DocumentTypes","GET /frs/api/v1/Projects","POST /frs/api/v1/Projects","POST /frs/v1/Projects *and* /Folders","GET /frs/api/v1/Projects","GET /frs/v1/Folders(","POST /frs/api/v1/UpdateEntityAccess","POST /frs/internal/api/v1/GetPermissions","POST frs/internal/api/v1/GetEffPrivilegeForDoctypeContainer","GET /frs/api/v1/Projects?$expand=sourceControlAttributes&$orderby=name","POST /frs/v1/LookupArtifactsDetailsByPath","GET /frs/api/v1/FetchProjectStatForRecentEntity()","GET /frs/api/v1/FetchDefaultLocation()","GET /frs/api/v1/BaseEntities","DELETE /frs/v1/Folders(","DELETE /frs/api/v1/Projects("],
        "containername":"frs",
        "logfilepath":"*/var/log/containers/frs*"
    },
    {
        "audit-service":["GET /auditlog-service/api/v1/Orgs( *and* )/AuditEntries?$count=true"],
        "containername":"auditlog-service",
        "logfilepath":"*/var/log/containers/auditlog-service*"
    },
    {
        "preference-service":["POST /preference-service/api/v1/Users( *and* )/Preferences","GET /preference-service/api/v1/Users *and* )/Preferences","PUT /preference-service/api/v1/Users( *and* )/Preferences","DELETE /preference-service/api/v1/Users( *and* )/Preferences","GET /preference-service/internal/api/v1/System/Preferences","POST /preference-service/api/v1/Orgs( *and* )/Preferences","GET /preference-service/api/v1/Orgs( *and* )/Preferences","DELETE /preference-service/api/v1/Orgs( *and* )/Preferences('iics.feature.onboard.wizardCompleted')"],
        "containername":"preference-service",
        "logfilepath":"*/var/log/containers/preference-service*"
    },
    {
        "migration-service":["POST /migration/api/v1/ImportJobs/$package","POST /migration/api/v1/ImportJobs( *and* /OData.migration_service.StartImportJob","GET /migration/api/v1/ImportJobs","GET /migration/api/v1/ImportJobs( *and* /ImportObjects"],
        "containername":"migration-service",
        "logfilepath":"*/var/log/containers/migration-service*"
    },
    {
        "admin-service":["POST /admin-service/api/v1/predownload/agent/ *and* /available"],
        "containername":"admin-service",
        "logfilepath":"*/var/log/containers/admin-service*"
    },
    {
        "vcs":["GET /vcs/api/v1/RepoConnection/orgId/","GET /vcs/api/v1/UserCredential/userId/"],
        "containername":"vcs",
        "logfilepath":"*/var/log/containers/vcs*"
    },
    {
        "ac-service":["GET /ac/logout","GET /ac/resources/static/images/dlg_close.gif","GET /ac/resources/static/js/lang/calendar-en.js","GET /ac/resources/static/images/Loading.gif"],
        "containername":"ac",
        "logfilepath":"*/var/log/containers/ac*"
    },
    {
        "ldm-service":["GET /ldm/api/v1/connection"],
        "containername":"ldm-service",
        "logfilepath":"*/var/log/containers/ldm-service*"
    },
    {
        "runtime":["GET /saas/api/v2/agent/details/","GET /saas/api/v2/agent/","GET /saas/api/v2/agent/name/","GET /saas/api/v2/runtimeEnvironment/","GET /saas/api/v2/runtimeEnvironment/name/","GET /saas/api/v2/connection"],
        "containername":"not using for runtime",
        "logfilepath":"/var/log/haproxy.log"
    }
  ],
  "iics-qa-perfids":[
	{
      "auth-service":["GET /authz-service/.well-known/jwks.json","POST /authz-service/oauth/token"],
      "containername":"authz-service",
      "logfilepath":"*/var/log/containers/auth*"
    },{
      "branding-service":["GET /branding-service/api/v1/OrgBranding","POST /branding-service/api/v1/OrgBranding","DELETE /branding-service/api/v1/OrgBranding","GET /branding-service/api/v1/Themes"],
      "containername":"branding-service",
      "logfilepath":"*/var/log/containers/branding-service*"
    },{
      "content-repository":["GET /content-repo/api/v1/Contents?$filter=contentType","GET /content-repo/api/v1/GetNewAssetContentsWithTags(tagNames='NEW_ASSET_MAPPINGS')?$expand=tags","GET /content-repo/api/v1/GetNewAssetContentsWithTags(tagNames='NEW_ASSET_TASKFLOWS)?$expand=tags"],
      "containername":"content-repository-service",
      "logfilepath":"*/var/log/containers/content-repository-service*"
    },{
      "ids-service":["POST /identity-service/api/v1/Login","POST /identity-service/api/v1/VerifyIdentity","GET /identity-service/agent/api/v1/Ping","GET /identity-service/api/v1/CurrentUserInfo","POST /identity-service/api/v1/Users","POST /identity-service/agent/api/v1/Login","POST /identity-service/api/v1/Orgs","GET /identity-service/home","DELETE /identity-service/api/v1/Users(","DELETE /identity-service/api/v1/Orgs(","POST /identity-service/podlink/api/v1/Token/LoginAs","POST /identity-service/podlink/api/v1/Login","POST /identity-service/podlink/agent/api/v1/Login","POST /identity-service/agent/api/v1/Logout","POST /identity-service/api/v1/Logout"],
      "containername":"identity-service",
      "logfilepath":"*/var/log/containers/identity-service*"
    },{
      "ma-service":["GET /ma/api/v3/SCIMToken","POST /ma/api/v2/user/login","POST /ma/api/v2/user/logout","POST /ma/api/v3/InternalLogin","GET /ma/api/v3/Users(","GET /ma/home","POST /ma/podlink/api/v3/Login","POST /ma/api/v3/Users","POST /ma/api/v3/DeleteUser","POST /ma/api/v3/agent/Login","POST /ma/api/v3/agent/Logout","POST /ma/api/v3/Logout"],
      "containername":"ma-service",
      "logfilepath":"*/var/log/containers/ma-service*"
    },{
      "scim-service":["GET /scim-service/Users/","GET /scim-service/v2/Users/","GET /scim-service/Groups/","POST /scim-service/Users","PATCH /scim-service/Users/","POST /scim-service/v2/Groups","PATCH /scim-service/v2/Groups/","PUT /scim-service/v2/Groups/","GET /scim-service/Groups?excludedAttributes=members&filter=displayName","DELETE /scim-service/Users/","DELETE /scim-service/Groups/"],
      "containername":"scim-service",
      "logfilepath":"*/var/log/containers/scim-service*"
    },{
      "v3api":["POST /saas/public/core/v3/login","GET /saas/public/core/v3/objects","POST /saas/public/core/v3/export","POST /saas/public/core/v3/lookup","PUT /saas/public/core/v3/Orgs/ *and* /addSamlGroupMappings","PUT /saas/public/core/v3/Orgs/ *and* /addSamlRoleMappings","PUT /saas/public/core/v3/Orgs/ *and* /removeSamlGroupMappings","PUT /saas/public/core/v3/Orgs/ *and* /removeSamlRoleMappings"],
      "containername":"v3api",
      "logfilepath":"*/var/log/containers/v3api*"
    }
]
}`
	var result config.APIDteails

	err := json.Unmarshal([]byte(listdetails), &result)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, nil
	}
	//fmt.Println(result)

	for _, perfusw1 := range result.IicsQaPerfusw1 {
		var containernames, filepaths string
		containernames = perfusw1.Containername
		filepaths = perfusw1.Logfilepath
		for _, apisofservices := range perfusw1.SchedulerService {
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames, filepaths, pod1)
			SendESRequest(apisofservices, perfusw1map, query)
		}
		for _, apis := range perfusw1.KmsService {
			query := buildmyquery(StartDate, EndDate, apis, containernames, filepaths, pod1)
			SendESRequest(apis, perfusw1map, query)
		}
		for _, apisofservices := range perfusw1.NotificationService {
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames, filepaths, pod1)
			SendESRequest(apisofservices, perfusw1map, query)
		}
		for _, apis := range perfusw1.CAService {
			query := buildmyquery(StartDate, EndDate, apis, containernames, filepaths, pod1)
			SendESRequest(apis, perfusw1map, query)
		}
		for _, apisofservices := range perfusw1.LicenseService {
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames, filepaths, pod1)
			SendESRequest(apisofservices, perfusw1map, query)
		}
		for _, apis := range perfusw1.JlsService {
			query := buildmyquery(StartDate, EndDate, apis, containernames, filepaths, pod1)
			SendESRequest(apis, perfusw1map, query)
		}
		for _, apisofservices := range perfusw1.BundleService {
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames, filepaths, pod1)
			SendESRequest(apisofservices, perfusw1map, query)
		}
		
		for _, apisofservices := range perfusw1.SessionService {
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames, filepaths, pod1)
			SendESRequest(apisofservices, perfusw1map, query)
		}
		for _, apis := range perfusw1.Frs {
			query := buildmyquery(StartDate, EndDate, apis, containernames, filepaths, pod1)
			SendESRequest(apis, perfusw1map, query)
		}
		for _, apisofservices := range perfusw1.AuditService {
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames, filepaths, pod1)
			SendESRequest(apisofservices, perfusw1map, query)
		}
		for _, apis := range perfusw1.PreferenceService {
			query := buildmyquery(StartDate, EndDate, apis, containernames, filepaths, pod1)
			SendESRequest(apis, perfusw1map, query)
		}
		for _, apis := range perfusw1.MigrationService {
			query := buildmyquery(StartDate, EndDate, apis, containernames, filepaths, pod1)
			SendESRequest(apis, perfusw1map, query)
		}
		for _, apisofservices := range perfusw1.AdminService {
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames, filepaths, pod1)
			SendESRequest(apisofservices, perfusw1map, query)
		}
		for _, apis := range perfusw1.Vcs {
			query := buildmyquery(StartDate, EndDate, apis, containernames, filepaths, pod1)
			SendESRequest(apis, perfusw1map, query)
		}
		for _, apisofservices := range perfusw1.AcService {
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames, filepaths, pod1)
			SendESRequest(apisofservices, perfusw1map, query)
		}
		for _, apis := range perfusw1.LdmService {
			query := buildmyquery(StartDate, EndDate, apis, containernames, filepaths, pod1)
			SendESRequest(apis, perfusw1map, query)
		}
		for _, apis := range perfusw1.Runtime {
			query := QueryRuntime(StartDate, EndDate, apis, filepaths, environment)
			SendESRequest(apis, perfusw1map, query)
		}
	}

	for _, perfids := range result.IicsQaPerfids {
		var containernames2, filepaths2 string
		containernames2 = perfids.Containername
		filepaths2 = perfids.Logfilepath

		//log.Println(perfusw1, podname)
		for _, apisofservices := range perfids.AuthService {
			//log.Println(apisofservices, containernames, filepaths)
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames2, filepaths2, ids)
			SendESRequest(apisofservices, perfidsmap, query)
		}
		for _, apisofservices := range perfids.BrandingService {
			//log.Println(apisofservices, containernames, filepaths)
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames2, filepaths2, ids)
			SendESRequest(apisofservices, perfidsmap, query)
		}
		for _, apisofservices := range perfids.ContentRepository {
			//log.Println(apisofservices, containernames, filepaths)
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames2, filepaths2, ids)
			SendESRequest(apisofservices, perfidsmap, query)
		}
		for _, apisofservices := range perfids.IdsService {
			//log.Println(apisofservices, containernames, filepaths)
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames2, filepaths2, ids)
			SendESRequest(apisofservices, perfidsmap, query)
		}
		for _, apisofservices := range perfids.MaService {
			//log.Println(apisofservices, containernames, filepaths)
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames2, filepaths2, ids)
			SendESRequest(apisofservices, perfidsmap, query)
		}
		for _, apisofservices := range perfids.ScimService {
			//log.Println(apisofservices, containernames, filepaths)
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames2, filepaths2, ids)
			SendESRequest(apisofservices, perfidsmap, query)
		}
		for _, apisofservice := range perfids.V3API {
			query := buildmyquery(StartDate, EndDate, apisofservice, containernames2, filepaths2, ids)
			SendESRequest(apisofservice, perfidsmap, query)
		}
	}
	//fmt.Println(perfusw1map, perfidsmap)
	return perfidsmap, perfusw1map
}
