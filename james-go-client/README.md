# Go API client for openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.8.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/searchlight/james-go-client"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AddressAliasAPI* | [**CreateAlias**](docs/AddressAliasAPI.md#createalias) | **Put** /address/aliases/{userAddress} | Add a new alias to a user
*AddressAliasAPI* | [**DeleteAlias**](docs/AddressAliasAPI.md#deletealias) | **Delete** /address/aliases/{userAddress} | Remove an alias from a user
*AddressAliasAPI* | [**GetAlias**](docs/AddressAliasAPI.md#getalias) | **Get** /address/aliases/{userAddress} | List alias sources of a user
*AddressAliasAPI* | [**ListAliases**](docs/AddressAliasAPI.md#listaliases) | **Get** /address/aliases | List users with aliases
*AddressForwardAPI* | [**AddDestination**](docs/AddressForwardAPI.md#adddestination) | **Put** /address/forwards/{userAddress} | Add a new destination to a forward
*AddressForwardAPI* | [**DeleteDestination**](docs/AddressForwardAPI.md#deletedestination) | **Delete** /address/forwards/{userAddress} | Remove a destination from a forward
*AddressForwardAPI* | [**ListDestinations**](docs/AddressForwardAPI.md#listdestinations) | **Get** /address/forwards/{userAddress} | List destinations in a forward
*AddressForwardAPI* | [**ListForwards**](docs/AddressForwardAPI.md#listforwards) | **Get** /address/forwards | List address forwards
*AddressGroupAPI* | [**AddMember**](docs/AddressGroupAPI.md#addmember) | **Put** /address/groups/{groupAddress} | Add a group member
*AddressGroupAPI* | [**ListGroups**](docs/AddressGroupAPI.md#listgroups) | **Get** /address/groups | List address groups
*AddressGroupAPI* | [**ListMembers**](docs/AddressGroupAPI.md#listmembers) | **Get** /address/groups/{groupAddress} | List members of a group
*AddressGroupAPI* | [**RemoveMember**](docs/AddressGroupAPI.md#removemember) | **Delete** /address/groups/{groupAddress} | Remove a group member
*AddressMappingAPI* | [**AddAddressMapping**](docs/AddressMappingAPI.md#addaddressmapping) | **Post** /mappings | Add an address mapping
*AddressMappingAPI* | [**ListAddressMappings**](docs/AddressMappingAPI.md#listaddressmappings) | **Get** /mappings | List all address mappings
*AddressMappingAPI* | [**RemoveAddressMapping**](docs/AddressMappingAPI.md#removeaddressmapping) | **Delete** /mappings/address/{mappingSource}/targets/{destinationAddress} | Remove an address mapping
*CassandraExtraAPI* | [**PerformActionOnCassandraMappings**](docs/CassandraExtraAPI.md#performactiononcassandramappings) | **Post** /cassandra/mappings | Perform an action on mappings_sources table
*CassandraSchemaUpgradeAPI* | [**GetLatestAvailableSchemaVersion**](docs/CassandraSchemaUpgradeAPI.md#getlatestavailableschemaversion) | **Get** /cassandra/version/latest | Retrieve latest available Cassandra schema version
*CassandraSchemaUpgradeAPI* | [**GetSchemaVersion**](docs/CassandraSchemaUpgradeAPI.md#getschemaversion) | **Get** /cassandra/version | Retrieve current Cassandra schema version
*CassandraSchemaUpgradeAPI* | [**UpgradeSchemaVersion**](docs/CassandraSchemaUpgradeAPI.md#upgradeschemaversion) | **Post** /cassandra/version/upgrade | Upgrade to a specific Cassandra schema version
*CassandraSchemaUpgradeAPI* | [**UpgradeToLatestSchemaVersion**](docs/CassandraSchemaUpgradeAPI.md#upgradetolatestschemaversion) | **Post** /cassandra/version/upgrade/latest | Upgrade to the latest Cassandra schema version
*DeletedMessageVaultAPI* | [**ExportDeletedMessages**](docs/DeletedMessageVaultAPI.md#exportdeletedmessages) | **Post** /deletedMessages/users/{user}/actions/export | Export deleted messages for a specific user
*DeletedMessageVaultAPI* | [**PurgeMessage**](docs/DeletedMessageVaultAPI.md#purgemessage) | **Delete** /deletedMessages/users/{user}/messages/{messageId} | Permanently remove a deleted message
*DeletedMessageVaultAPI* | [**PurgeMessages**](docs/DeletedMessageVaultAPI.md#purgemessages) | **Delete** /deletedMessages | Purge all expired deleted messages
*DeletedMessageVaultAPI* | [**RestoreDeletedMessages**](docs/DeletedMessageVaultAPI.md#restoredeletedmessages) | **Post** /deletedMessages/users/{user}/actions/restore | Restore deleted messages for a specific user
*DlpAPI* | [**FetchDLPConfiguration**](docs/DlpAPI.md#fetchdlpconfiguration) | **Get** /dlp/rules/{senderDomain}/rules/{ruleId} | Fetch a DLP configuration item by sender domain and rule id
*DlpAPI* | [**ListDLPConfiguration**](docs/DlpAPI.md#listdlpconfiguration) | **Get** /dlp/rules/{senderDomain} | List DLP configuration by sender domain
*DlpAPI* | [**RemoveDLPConfiguration**](docs/DlpAPI.md#removedlpconfiguration) | **Delete** /dlp/rules/{senderDomain} | Remove DLP configuration by sender domain
*DlpAPI* | [**StoreDLPConfiguration**](docs/DlpAPI.md#storedlpconfiguration) | **Put** /dlp/rules/{senderDomain} | Store DLP configuration by sender domain
*DomainMappingAPI* | [**AddDomainMapping**](docs/DomainMappingAPI.md#adddomainmapping) | **Put** /domainMappings | Add a domain mapping
*DomainMappingAPI* | [**ListDestinationDomains**](docs/DomainMappingAPI.md#listdestinationdomains) | **Get** /domainMappings/{fromDomain} | List all destination domains for a source domain
*DomainMappingAPI* | [**ListDomainMappings**](docs/DomainMappingAPI.md#listdomainmappings) | **Get** /domainMappings | List all domain mappings
*DomainMappingAPI* | [**RemoveDomainMapping**](docs/DomainMappingAPI.md#removedomainmapping) | **Delete** /domainMappings | Remove a domain mapping
*DomainQuotaAPI* | [**DeleteDomainQuotaCount**](docs/DomainQuotaAPI.md#deletedomainquotacount) | **Delete** /quota/domains/{domainToBeUsed}/count | Delete the quota count for a domain
*DomainQuotaAPI* | [**DeleteDomainQuotaSize**](docs/DomainQuotaAPI.md#deletedomainquotasize) | **Delete** /quota/domains/{domainToBeUsed}/size | Delete the quota size for a domain
*DomainQuotaAPI* | [**GetDomainQuota**](docs/DomainQuotaAPI.md#getdomainquota) | **Get** /quota/domains/{domainToBeUsed} | Get the quota for a domain
*DomainQuotaAPI* | [**GetDomainQuotaCount**](docs/DomainQuotaAPI.md#getdomainquotacount) | **Get** /quota/domains/{domainToBeUsed}/count | Get the quota count for a domain
*DomainQuotaAPI* | [**GetDomainQuotaSize**](docs/DomainQuotaAPI.md#getdomainquotasize) | **Get** /quota/domains/{domainToBeUsed}/size | Get the quota size for a domain
*DomainQuotaAPI* | [**UpdateDomainQuota**](docs/DomainQuotaAPI.md#updatedomainquota) | **Put** /quota/domains/{domainToBeUsed} | Update the quota for a domain
*DomainQuotaAPI* | [**UpdateDomainQuotaCount**](docs/DomainQuotaAPI.md#updatedomainquotacount) | **Put** /quota/domains/{domainToBeUsed}/count | Update the quota count for a domain
*DomainQuotaAPI* | [**UpdateDomainQuotaSize**](docs/DomainQuotaAPI.md#updatedomainquotasize) | **Put** /quota/domains/{domainToBeUsed}/size | Update the quota size for a domain
*DomainsAPI* | [**CreateDomain**](docs/DomainsAPI.md#createdomain) | **Put** /domains/{domainToBeCreated} | Create a domain
*DomainsAPI* | [**CreateDomainAlias**](docs/DomainsAPI.md#createdomainalias) | **Put** /domains/{domainName}/aliases | Create an alias for a domain
*DomainsAPI* | [**DeleteDomain**](docs/DomainsAPI.md#deletedomain) | **Delete** /domains/{domainToBeDeleted} | Delete a domain
*DomainsAPI* | [**DeleteDomainAlias**](docs/DomainsAPI.md#deletedomainalias) | **Delete** /domains/{domainName}/aliases | Delete an alias for a domain
*DomainsAPI* | [**DeleteUserDataOfDomain**](docs/DomainsAPI.md#deleteuserdataofdomain) | **Post** /domains/{domainToBeUsed} | Delete all users data of a domain
*DomainsAPI* | [**ExistsDomain**](docs/DomainsAPI.md#existsdomain) | **Get** /domains/{domainName} | Test if a domain exists
*DomainsAPI* | [**ListDomainAliases**](docs/DomainsAPI.md#listdomainaliases) | **Get** /domains/{domainName}/aliases | Get the list of aliases for a domain
*DomainsAPI* | [**ListDomains**](docs/DomainsAPI.md#listdomains) | **Get** /domains | Get the list of domains
*EventDeadLetterAPI* | [**DeleteAllEvents**](docs/EventDeadLetterAPI.md#deleteallevents) | **Delete** /events/deadLetter/groups/{groupName} | Delete all events of a group
*EventDeadLetterAPI* | [**DeleteEvent**](docs/EventDeadLetterAPI.md#deleteevent) | **Delete** /events/deadLetter/groups/{groupName}/{insertionId} | Delete an event
*EventDeadLetterAPI* | [**GetEvent**](docs/EventDeadLetterAPI.md#getevent) | **Get** /events/deadLetter/groups/{groupName}/{insertionId} | Get event details
*EventDeadLetterAPI* | [**ListFailedEvents**](docs/EventDeadLetterAPI.md#listfailedevents) | **Get** /events/deadLetter/groups/{groupName} | List failed events for a given group
*EventDeadLetterAPI* | [**ListMailboxListenerGroups**](docs/EventDeadLetterAPI.md#listmailboxlistenergroups) | **Get** /events/deadLetter/groups | List Mailbox Listener Groups
*EventDeadLetterAPI* | [**RedeliverAllEvents**](docs/EventDeadLetterAPI.md#redeliverallevents) | **Post** /events/deadLetter/groups | Redeliver all events
*EventDeadLetterAPI* | [**RedeliverEvent**](docs/EventDeadLetterAPI.md#redeliverevent) | **Post** /events/deadLetter/groups/{groupName}/{insertionId}/reDeliver | Redeliver a single event
*EventDeadLetterAPI* | [**RedeliverGroupEvents**](docs/EventDeadLetterAPI.md#redelivergroupevents) | **Post** /events/deadLetter/groups/{groupName}/reDeliver | Redeliver group events
*GarbageCollectionAPI* | [**RunBlobGarbageCollector**](docs/GarbageCollectionAPI.md#runblobgarbagecollector) | **Delete** /blobs | Run blob garbage collection
*GhostMailboxAPI* | [**CorrectGhostMailbox**](docs/GhostMailboxAPI.md#correctghostmailbox) | **Post** /cassandra/mailbox/merging | Correct ghost mailbox by merging
*GlobalQuotaAPI* | [**DeleteGlobalQuotaCount**](docs/GlobalQuotaAPI.md#deleteglobalquotacount) | **Delete** /quota/count | Delete the global quota count
*GlobalQuotaAPI* | [**DeleteGlobalQuotaSize**](docs/GlobalQuotaAPI.md#deleteglobalquotasize) | **Delete** /quota/size | Delete the global quota size
*GlobalQuotaAPI* | [**GetGlobalQuota**](docs/GlobalQuotaAPI.md#getglobalquota) | **Get** /quota | Get the global quota
*GlobalQuotaAPI* | [**GetGlobalQuotaCount**](docs/GlobalQuotaAPI.md#getglobalquotacount) | **Get** /quota/count | Get the global quota count
*GlobalQuotaAPI* | [**GetGlobalQuotaSize**](docs/GlobalQuotaAPI.md#getglobalquotasize) | **Get** /quota/size | Get the global quota size
*GlobalQuotaAPI* | [**UpdateGlobalQuota**](docs/GlobalQuotaAPI.md#updateglobalquota) | **Put** /quota | Update the global quota
*GlobalQuotaAPI* | [**UpdateGlobalQuotaCount**](docs/GlobalQuotaAPI.md#updateglobalquotacount) | **Put** /quota/count | Update the global quota count
*GlobalQuotaAPI* | [**UpdateGlobalQuotaSize**](docs/GlobalQuotaAPI.md#updateglobalquotasize) | **Put** /quota/size | Update the global quota size
*HealthcheckAPI* | [**CheckAllComponents**](docs/HealthcheckAPI.md#checkallcomponents) | **Get** /healthcheck | Check all components
*HealthcheckAPI* | [**CheckComponent**](docs/HealthcheckAPI.md#checkcomponent) | **Get** /healthcheck/checks/{componentName} | Check single component
*HealthcheckAPI* | [**ListAllHealthChecks**](docs/HealthcheckAPI.md#listallhealthchecks) | **Get** /healthcheck/checks | List all health checks
*JmapUploadsAPI* | [**CleanUploadRepository**](docs/JmapUploadsAPI.md#cleanuploadrepository) | **Delete** /jmap/uploads | Clean upload repository
*MailQueueAPI* | [**DeleteMailsOfMailQueue**](docs/MailQueueAPI.md#deletemailsofmailqueue) | **Delete** /mailQueues/{mailQueueName}/mails | Delete mails from a mail queue
*MailQueueAPI* | [**FlushMailsOfMailQueue**](docs/MailQueueAPI.md#flushmailsofmailqueue) | **Patch** /mailQueues/{mailQueueName} | Flush mails from a mail queue
*MailQueueAPI* | [**ListMailQueues**](docs/MailQueueAPI.md#listmailqueues) | **Get** /mailQueues | List mail queues
*MailQueueAPI* | [**ListMailsOfMailQueue**](docs/MailQueueAPI.md#listmailsofmailqueue) | **Get** /mailQueues/{mailQueueName}/mails | List mails of a mail queue
*MailQueueAPI* | [**RepublishMailQueue**](docs/MailQueueAPI.md#republishmailqueue) | **Post** /mailQueues | RabbitMQ republishing a mail queue from Cassandra
*MailRepositoryAPI* | [**CreateMailRepository**](docs/MailRepositoryAPI.md#createmailrepository) | **Put** /mailRepositories/{encodedPathOfTheRepository} | Create a mail repository
*MailRepositoryAPI* | [**GetMailRepository**](docs/MailRepositoryAPI.md#getmailrepository) | **Get** /mailRepositories/{encodedPathOfTheRepository} | Getting additional information for a mail repository
*MailRepositoryAPI* | [**ListMailRepositories**](docs/MailRepositoryAPI.md#listmailrepositories) | **Get** /mailRepositories | Listing mail repositories
*MailRepositoryAPI* | [**ListMailsInMailRepository**](docs/MailRepositoryAPI.md#listmailsinmailrepository) | **Get** /mailRepositories/{encodedPathOfTheRepository}/mails | Listing mails contained in a mail repository
*MailboxAPI* | [**PerformActionsOnMailboxes**](docs/MailboxAPI.md#performactionsonmailboxes) | **Post** /mailboxes | Perform actions on mailboxes
*MailboxAPI* | [**ReindexMailbox**](docs/MailboxAPI.md#reindexmailbox) | **Post** /mailboxes/{mailboxId} | Reindex a mailbox
*MessagesAPI* | [**ReindexEmail**](docs/MessagesAPI.md#reindexemail) | **Post** /messages/{messageId} | Reindex a single mail by messageId
*MessagesAPI* | [**ScheduleTask**](docs/MessagesAPI.md#scheduletask) | **Post** /messages | Schedule a task for fixing message inconsistencies
*RegexMappingAPI* | [**AddRegexMapping**](docs/RegexMappingAPI.md#addregexmapping) | **Post** /mappings/regex/{mappingSource}/targets/{regex} | Add a regex mapping
*RegexMappingAPI* | [**RemoveRegexMapping**](docs/RegexMappingAPI.md#removeregexmapping) | **Delete** /mappings/regex/{mappingSource}/targets/{regex} | Remove a regex mapping
*SendMailAPI* | [**SendEmail**](docs/SendMailAPI.md#sendemail) | **Post** /mail-transfer-service | Send email
*SieveQuotaAPI* | [**GetUserSieveQuota**](docs/SieveQuotaAPI.md#getusersievequota) | **Get** /sieve/quota/users/{userEmail} | Retrieve user sieve quota
*SieveQuotaAPI* | [**RemoveUserSieveQuota**](docs/SieveQuotaAPI.md#removeusersievequota) | **Delete** /sieve/quota/users/{userEmail} | Remove user sieve quota
*SieveQuotaAPI* | [**SieveQuotaDefaultDelete**](docs/SieveQuotaAPI.md#sievequotadefaultdelete) | **Delete** /sieve/quota/default | Remove global sieve quota
*SieveQuotaAPI* | [**SieveQuotaDefaultGet**](docs/SieveQuotaAPI.md#sievequotadefaultget) | **Get** /sieve/quota/default | Retrieve global sieve quota
*SieveQuotaAPI* | [**UpdateGlobalSieveQuota**](docs/SieveQuotaAPI.md#updateglobalsievequota) | **Put** /sieve/quota/default | Update global sieve quota
*SieveQuotaAPI* | [**UpdateUserSieveQuota**](docs/SieveQuotaAPI.md#updateusersievequota) | **Put** /sieve/quota/users/{userEmail} | Update user sieve quota
*TaskAPI* | [**AwaitTaskCompletion**](docs/TaskAPI.md#awaittaskcompletion) | **Get** /tasks/{taskId}/await | Await the completion of a task
*TaskAPI* | [**CancelTask**](docs/TaskAPI.md#canceltask) | **Delete** /tasks/{taskId} | Cancel a task
*TaskAPI* | [**GetTask**](docs/TaskAPI.md#gettask) | **Get** /tasks/{taskId} | Get a task&#39;s details
*TaskAPI* | [**ListTasks**](docs/TaskAPI.md#listtasks) | **Get** /tasks | List tasks
*UserMailboxAPI* | [**ClearMailbox**](docs/UserMailboxAPI.md#clearmailbox) | **Delete** /users/{username}/mailboxes/{mailboxName}/messages | Clear mailbox content
*UserMailboxAPI* | [**CountEmails**](docs/UserMailboxAPI.md#countemails) | **Get** /users/{username}/mailboxes/{mailboxName}/messageCount | Count emails in a mailbox
*UserMailboxAPI* | [**CountUnseenEmails**](docs/UserMailboxAPI.md#countunseenemails) | **Get** /users/{username}/mailboxes/{mailboxName}/unseenMessageCount | Count unseen emails in a mailbox
*UserMailboxAPI* | [**CreateMailbox**](docs/UserMailboxAPI.md#createmailbox) | **Put** /users/{username}/mailboxes/{mailboxNameToBeCreated} | Create a mailbox
*UserMailboxAPI* | [**DeleteMailbox**](docs/UserMailboxAPI.md#deletemailbox) | **Delete** /users/{username}/mailboxes/INBOX.work | Delete a mailbox and its children
*UserMailboxAPI* | [**DeleteMailboxes**](docs/UserMailboxAPI.md#deletemailboxes) | **Delete** /users/{username}/mailboxes | Delete user mailboxes
*UserMailboxAPI* | [**ExistsMailbox**](docs/UserMailboxAPI.md#existsmailbox) | **Get** /users/{username}/mailboxes/{mailboxNameToBeTested} | Test existence of a mailbox
*UserMailboxAPI* | [**ExportMailboxes**](docs/UserMailboxAPI.md#exportmailboxes) | **Post** /users/{username}/mailboxes?action&#x3D;export | Export user mailboxes
*UserMailboxAPI* | [**ListMailboxes**](docs/UserMailboxAPI.md#listmailboxes) | **Get** /users/{username}/mailboxes | List user mailboxes
*UserMailboxAPI* | [**RecomputeCassandraFilteringProjection**](docs/UserMailboxAPI.md#recomputecassandrafilteringprojection) | **Post** /mailboxes?task&#x3D;populateFilteringProjection | Recompute Cassandra filtering projection
*UserMailboxAPI* | [**RecomputeMessageViewProjection**](docs/UserMailboxAPI.md#recomputemessageviewprojection) | **Post** /users/{username}/mailboxes?task&#x3D;recomputeFastViewProjectionItems | Recompute User JMAP fast message view projection
*UserMailboxAPI* | [**ReindexEmails**](docs/UserMailboxAPI.md#reindexemails) | **Post** /users/{username}/mailboxes?task&#x3D;reIndex | Reindex a user&#39;s mails
*UserMailboxAPI* | [**SubscribeAllMailboxes**](docs/UserMailboxAPI.md#subscribeallmailboxes) | **Post** /users/{username}/mailboxes?task&#x3D;subscribeAll | Subscribe a user to all of their mailboxes
*UserMappingAPI* | [**ListUserMappings**](docs/UserMappingAPI.md#listusermappings) | **Get** /mappings/user/{userAddress} | Listing User Mappings
*UserQuotaAPI* | [**DeleteQuotaCount**](docs/UserQuotaAPI.md#deletequotacount) | **Delete** /quota/users/{username}/count | Delete the quota count for a user
*UserQuotaAPI* | [**DeleteQuotaSize**](docs/UserQuotaAPI.md#deletequotasize) | **Delete** /quota/users/{username}/size | Delete the quota size for a user
*UserQuotaAPI* | [**GetQuota**](docs/UserQuotaAPI.md#getquota) | **Get** /quota/users/{username} | Get the quota for a user
*UserQuotaAPI* | [**GetQuotaCount**](docs/UserQuotaAPI.md#getquotacount) | **Get** /quota/users/{username}/count | Get the quota count for a user
*UserQuotaAPI* | [**GetQuotaSize**](docs/UserQuotaAPI.md#getquotasize) | **Get** /quota/users/{username}/size | Get the quota size for a user
*UserQuotaAPI* | [**RecomputeCurrentQuotas**](docs/UserQuotaAPI.md#recomputecurrentquotas) | **Post** /quota/users | Recompute current quotas for users
*UserQuotaAPI* | [**SearchByQuota**](docs/UserQuotaAPI.md#searchbyquota) | **Get** /quota/users | Search users by quota ratio
*UserQuotaAPI* | [**UpdateQuota**](docs/UserQuotaAPI.md#updatequota) | **Put** /quota/users/{username} | Update the quota for a user
*UserQuotaAPI* | [**UpdateQuotaCount**](docs/UserQuotaAPI.md#updatequotacount) | **Put** /quota/users/{username}/count | Update the quota count for a user
*UserQuotaAPI* | [**UpdateQuotaSize**](docs/UserQuotaAPI.md#updatequotasize) | **Put** /quota/users/{username}/size | Update the quota size for a user
*UsersAPI* | [**AddDelegatedUser**](docs/UsersAPI.md#adddelegateduser) | **Put** /users/{baseUser}/authorizedUsers/delegatedUser | Add a delegated user to a base user
*UsersAPI* | [**ChangeUsername**](docs/UsersAPI.md#changeusername) | **Post** /users/{oldUser}/rename/{newUser} | Change a username
*UsersAPI* | [**CreateUserIdentity**](docs/UsersAPI.md#createuseridentity) | **Post** /users/{username}/identities | Create a JMAP user identity
*UsersAPI* | [**DeleteUser**](docs/UsersAPI.md#deleteuser) | **Delete** /users/{username} | Delete a user
*UsersAPI* | [**ExistsUser**](docs/UsersAPI.md#existsuser) | **Head** /users/{username} | Test user existence
*UsersAPI* | [**ListAllowedFromHeaders**](docs/UsersAPI.md#listallowedfromheaders) | **Get** /users/{givenUser}/allowedFromHeaders | Retrieve the list of allowed From headers for a given user
*UsersAPI* | [**ListDelegatedUsers**](docs/UsersAPI.md#listdelegatedusers) | **Get** /users/{baseUser}/authorizedUsers | Retrieve the list of delegated users of a base user
*UsersAPI* | [**ListUserIdentities**](docs/UsersAPI.md#listuseridentities) | **Get** /users/{username}/identities | Retrieve the user identities
*UsersAPI* | [**ListUsers**](docs/UsersAPI.md#listusers) | **Get** /users | Retrieve the user list
*UsersAPI* | [**RemoveAllDelegatedUsers**](docs/UsersAPI.md#removealldelegatedusers) | **Delete** /users/{baseUser}/authorizedUsers | Remove all delegated users of a base user
*UsersAPI* | [**RemoveDelegatedUser**](docs/UsersAPI.md#removedelegateduser) | **Delete** /users/{baseUser}/authorizedUsers/delegatedUser | Remove a delegated user from a base user
*UsersAPI* | [**UpdateUserIdentity**](docs/UsersAPI.md#updateuseridentity) | **Put** /users/{username}/identities/{identityId} | Update a JMAP user identity
*UsersAPI* | [**UpsertUser**](docs/UsersAPI.md#upsertuser) | **Put** /users/{username} | Create or Update User


## Documentation For Models

 - [CheckAllComponents200Response](docs/CheckAllComponents200Response.md)
 - [CleanUploadRepository201Response](docs/CleanUploadRepository201Response.md)
 - [CleanUploadRepositoryRequest](docs/CleanUploadRepositoryRequest.md)
 - [CorrectGhostMailbox201Response](docs/CorrectGhostMailbox201Response.md)
 - [CorrectGhostMailbox201ResponseAdditionalInformation](docs/CorrectGhostMailbox201ResponseAdditionalInformation.md)
 - [CorrectGhostMailboxRequest](docs/CorrectGhostMailboxRequest.md)
 - [CreateUserIdentityRequest](docs/CreateUserIdentityRequest.md)
 - [Criterion](docs/Criterion.md)
 - [DomainAlias](docs/DomainAlias.md)
 - [Error](docs/Error.md)
 - [ExecutionReport](docs/ExecutionReport.md)
 - [ExportDeletedMessagesRequest](docs/ExportDeletedMessagesRequest.md)
 - [FlushMailsOfMailQueueRequest](docs/FlushMailsOfMailQueueRequest.md)
 - [GetAlias200ResponseInner](docs/GetAlias200ResponseInner.md)
 - [GetGlobalQuota200Response](docs/GetGlobalQuota200Response.md)
 - [GetLatestAvailableSchemaVersion200Response](docs/GetLatestAvailableSchemaVersion200Response.md)
 - [GetMailRepository200Response](docs/GetMailRepository200Response.md)
 - [GetSchemaVersion200Response](docs/GetSchemaVersion200Response.md)
 - [HealthCheckInfo](docs/HealthCheckInfo.md)
 - [HealthCheckResult](docs/HealthCheckResult.md)
 - [ListAddressMappings200Response](docs/ListAddressMappings200Response.md)
 - [ListAddressMappings200ResponseMappingSourceInner](docs/ListAddressMappings200ResponseMappingSourceInner.md)
 - [ListDestinations200ResponseInner](docs/ListDestinations200ResponseInner.md)
 - [ListMailRepositories200ResponseInner](docs/ListMailRepositories200ResponseInner.md)
 - [ListMailboxes200ResponseInner](docs/ListMailboxes200ResponseInner.md)
 - [ListMailsOfMailQueue200ResponseInner](docs/ListMailsOfMailQueue200ResponseInner.md)
 - [ListUserIdentities200ResponseInner](docs/ListUserIdentities200ResponseInner.md)
 - [ListUserIdentities200ResponseInnerBccInner](docs/ListUserIdentities200ResponseInnerBccInner.md)
 - [ListUserMappings200ResponseInner](docs/ListUserMappings200ResponseInner.md)
 - [ListUsers200ResponseInner](docs/ListUsers200ResponseInner.md)
 - [PerformActionOnCassandraMappings201Response](docs/PerformActionOnCassandraMappings201Response.md)
 - [PerformActionsOnMailboxes201Response](docs/PerformActionsOnMailboxes201Response.md)
 - [RecomputeCurrentQuotasRequest](docs/RecomputeCurrentQuotasRequest.md)
 - [RunBlobGarbageCollector200Response](docs/RunBlobGarbageCollector200Response.md)
 - [ScheduleTask400Response](docs/ScheduleTask400Response.md)
 - [StoreDLPConfigurationRequest](docs/StoreDLPConfigurationRequest.md)
 - [StoreDLPConfigurationRequestRulesInner](docs/StoreDLPConfigurationRequestRulesInner.md)
 - [Task](docs/Task.md)
 - [TaskErrorsInner](docs/TaskErrorsInner.md)
 - [TaskId](docs/TaskId.md)
 - [TaskRunningOptions](docs/TaskRunningOptions.md)
 - [UpdateGlobalQuotaRequest](docs/UpdateGlobalQuotaRequest.md)
 - [UpdateGlobalSieveQuotaRequest](docs/UpdateGlobalSieveQuotaRequest.md)
 - [UpdateUserIdentityRequest](docs/UpdateUserIdentityRequest.md)
 - [UpdateUserSieveQuotaRequest](docs/UpdateUserSieveQuotaRequest.md)
 - [UpgradeSchemaVersion200Response](docs/UpgradeSchemaVersion200Response.md)
 - [UpgradeSchemaVersion200ResponseAdditionalInformation](docs/UpgradeSchemaVersion200ResponseAdditionalInformation.md)
 - [UpgradeToLatestSchemaVersion200Response](docs/UpgradeToLatestSchemaVersion200Response.md)
 - [UpgradeToLatestSchemaVersion200ResponseAdditionalInformation](docs/UpgradeToLatestSchemaVersion200ResponseAdditionalInformation.md)
 - [UpsertUserRequest](docs/UpsertUserRequest.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


