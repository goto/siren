# Siren APIs
Documentation of our Siren API with gRPC and
gRPC-Gateway.

## Version: 0.6

### /v1beta1/alerts/{provider_type}/{provider_id}

#### GET
##### Summary

list alerts

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| provider_type | path |  | Yes | string |
| provider_id | path |  | Yes | string (uint64) |
| resource_name | query |  | No | string |
| start_time | query |  | No | string (uint64) |
| end_time | query |  | No | string (uint64) |
| namespace_id | query |  | No | string (uint64) |
| silence_id | query |  | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [ListAlertsResponse](#listalertsresponse) |
| default | An unexpected error response. | [Status](#status) |

#### POST
##### Summary

create alerts

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| provider_type | path |  | Yes | string |
| provider_id | path |  | Yes | string (uint64) |
| body | body |  | Yes | object |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [CreateAlertsResponse](#createalertsresponse) |
| default | An unexpected error response. | [Status](#status) |

### /v1beta1/alerts/{provider_type}/{provider_id}/{namespace_id}

#### POST
##### Summary

create alerts with namespace

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| provider_type | path |  | Yes | string |
| provider_id | path |  | Yes | string (uint64) |
| namespace_id | path |  | Yes | string (uint64) |
| body | body |  | Yes | object |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [CreateAlertsWithNamespaceResponse](#createalertswithnamespaceresponse) |
| default | An unexpected error response. | [Status](#status) |

### /v1beta1/namespaces

#### GET
##### Summary

list namespaces

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [ListNamespacesResponse](#listnamespacesresponse) |
| default | An unexpected error response. | [Status](#status) |

#### POST
##### Summary

create a namespace

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | Yes | [CreateNamespaceRequest](#createnamespacerequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [CreateNamespaceResponse](#createnamespaceresponse) |
| default | An unexpected error response. | [Status](#status) |

### /v1beta1/namespaces/{id}

#### GET
##### Summary

get a namespace

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string (uint64) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [GetNamespaceResponse](#getnamespaceresponse) |
| default | An unexpected error response. | [Status](#status) |

#### DELETE
##### Summary

delete a namespace

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string (uint64) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [DeleteNamespaceResponse](#deletenamespaceresponse) |
| default | An unexpected error response. | [Status](#status) |

#### PUT
##### Summary

update a namespace

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string (uint64) |
| body | body |  | Yes | object |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [UpdateNamespaceResponse](#updatenamespaceresponse) |
| default | An unexpected error response. | [Status](#status) |

### /v1beta1/notifications

#### GET
##### Summary

List notifications

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| type | query |  | No | string |
| template | query |  | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [ListNotificationsResponse](#listnotificationsresponse) |
| default | An unexpected error response. | [Status](#status) |

#### POST
##### Summary

Post an event notification

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | Yes | [PostNotificationRequest](#postnotificationrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [PostNotificationResponse](#postnotificationresponse) |
| default | An unexpected error response. | [Status](#status) |

### /v1beta1/notifications/{notification_id}/messages

#### GET
##### Summary

List messages generated by notification

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| notification_id | path |  | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [ListNotificationMessagesResponse](#listnotificationmessagesresponse) |
| default | An unexpected error response. | [Status](#status) |

### /v1beta1/providers

#### GET
##### Summary

list providers

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| urn | query |  | No | string |
| type | query |  | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [ListProvidersResponse](#listprovidersresponse) |
| default | An unexpected error response. | [Status](#status) |

#### POST
##### Summary

create a provider

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | Yes | [CreateProviderRequest](#createproviderrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [CreateProviderResponse](#createproviderresponse) |
| default | An unexpected error response. | [Status](#status) |

### /v1beta1/providers/{id}

#### GET
##### Summary

get a provider

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string (uint64) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [GetProviderResponse](#getproviderresponse) |
| default | An unexpected error response. | [Status](#status) |

#### DELETE
##### Summary

delete a provider

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string (uint64) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [DeleteProviderResponse](#deleteproviderresponse) |
| default | An unexpected error response. | [Status](#status) |

#### PUT
##### Summary

update a provider

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string (uint64) |
| body | body |  | Yes | object |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [UpdateProviderResponse](#updateproviderresponse) |
| default | An unexpected error response. | [Status](#status) |

### /v1beta1/receivers

#### GET
##### Summary

list receivers

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [ListReceiversResponse](#listreceiversresponse) |
| default | An unexpected error response. | [Status](#status) |

#### POST
##### Summary

create a receiver

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | Yes | [CreateReceiverRequest](#createreceiverrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [CreateReceiverResponse](#createreceiverresponse) |
| default | An unexpected error response. | [Status](#status) |

### /v1beta1/receivers/{id}

#### GET
##### Summary

get a receiver

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string (uint64) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [GetReceiverResponse](#getreceiverresponse) |
| default | An unexpected error response. | [Status](#status) |

#### DELETE
##### Summary

delete a receiver

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string (uint64) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [DeleteReceiverResponse](#deletereceiverresponse) |
| default | An unexpected error response. | [Status](#status) |

#### PUT
##### Summary

update a receiver

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string (uint64) |
| body | body |  | Yes | object |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [UpdateReceiverResponse](#updatereceiverresponse) |
| default | An unexpected error response. | [Status](#status) |

### /v1beta1/receivers/{id}/send

#### POST
##### Summary

send notification to receiver

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string (uint64) |
| body | body |  | Yes | object |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [NotifyReceiverResponse](#notifyreceiverresponse) |
| default | An unexpected error response. | [Status](#status) |

### /v1beta1/rules

#### GET
##### Summary

list rules

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| name | query |  | No | string |
| namespace | query |  | No | string |
| group_name | query |  | No | string |
| template | query |  | No | string |
| provider_namespace | query |  | No | string (uint64) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [ListRulesResponse](#listrulesresponse) |
| default | An unexpected error response. | [Status](#status) |

#### PUT
##### Summary

add/update a rule

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | Yes | [UpdateRuleRequest](#updaterulerequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [UpdateRuleResponse](#updateruleresponse) |
| default | An unexpected error response. | [Status](#status) |

### /v1beta1/silences

#### GET
##### Summary

get all silences

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| subscription_id | query |  | No | string (uint64) |
| namespace_id | query |  | No | string (uint64) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [ListSilencesResponse](#listsilencesresponse) |
| default | An unexpected error response. | [Status](#status) |

#### POST
##### Summary

create a silence

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | Yes | [CreateSilenceRequest](#createsilencerequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [CreateSilenceResponse](#createsilenceresponse) |
| default | An unexpected error response. | [Status](#status) |

### /v1beta1/silences/{id}

#### GET
##### Summary

get a silence

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [GetSilenceResponse](#getsilenceresponse) |
| default | An unexpected error response. | [Status](#status) |

#### DELETE
##### Summary

expire a silence

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [ExpireSilenceResponse](#expiresilenceresponse) |
| default | An unexpected error response. | [Status](#status) |

### /v1beta1/subscriptions

#### GET
##### Summary

List subscriptions

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| namespace_id | query |  | No | string (uint64) |
| silence_id | query |  | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [ListSubscriptionsResponse](#listsubscriptionsresponse) |
| default | An unexpected error response. | [Status](#status) |

#### POST
##### Summary

Create a subscription

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | Yes | [CreateSubscriptionRequest](#createsubscriptionrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [CreateSubscriptionResponse](#createsubscriptionresponse) |
| default | An unexpected error response. | [Status](#status) |

### /v1beta1/subscriptions/{id}

#### GET
##### Summary

Get a subscription

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string (uint64) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [GetSubscriptionResponse](#getsubscriptionresponse) |
| default | An unexpected error response. | [Status](#status) |

#### DELETE
##### Summary

Delete a subscription

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string (uint64) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [DeleteSubscriptionResponse](#deletesubscriptionresponse) |
| default | An unexpected error response. | [Status](#status) |

#### PUT
##### Summary

Update a subscription

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string (uint64) |
| body | body |  | Yes | object |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [UpdateSubscriptionResponse](#updatesubscriptionresponse) |
| default | An unexpected error response. | [Status](#status) |

### /v1beta1/templates

#### GET
##### Summary

list templates

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| tag | query |  | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [ListTemplatesResponse](#listtemplatesresponse) |
| default | An unexpected error response. | [Status](#status) |

#### PUT
##### Summary

add/update a template

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | Yes | [UpsertTemplateRequest](#upserttemplaterequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [UpsertTemplateResponse](#upserttemplateresponse) |
| default | An unexpected error response. | [Status](#status) |

### /v1beta1/templates/{name}

#### GET
##### Summary

get a template

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| name | path |  | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [GetTemplateResponse](#gettemplateresponse) |
| default | An unexpected error response. | [Status](#status) |

#### DELETE
##### Summary

delete a template

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| name | path |  | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [DeleteTemplateResponse](#deletetemplateresponse) |
| default | An unexpected error response. | [Status](#status) |

### /v1beta1/templates/{name}/render

#### POST
##### Summary

render a template

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| name | path |  | Yes | string |
| body | body |  | Yes | object |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [RenderTemplateResponse](#rendertemplateresponse) |
| default | An unexpected error response. | [Status](#status) |

### Models

#### Alert

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string (uint64) |  | No |
| provider_id | string (uint64) |  | No |
| resource_name | string |  | No |
| metric_name | string |  | No |
| metric_value | string |  | No |
| severity | string |  | No |
| rule | string |  | No |
| triggered_at | dateTime |  | No |
| namespace_id | string (uint64) |  | No |
| silence_status | string |  | No |
| created_at | dateTime |  | No |
| updated_at | dateTime |  | No |
| group_key | string |  | No |
| status | string |  | No |
| annotations | object |  | No |
| labels | object |  | No |
| generator_url | string |  | No |
| fingerprint | string |  | No |

#### Any

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| @type | string |  | No |

#### CreateAlertsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| alerts | [ [Alert](#alert) ] |  | No |

#### CreateAlertsWithNamespaceResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| alerts | [ [Alert](#alert) ] |  | No |

#### CreateNamespaceRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| name | string |  | No |
| urn | string |  | No |
| provider | string (uint64) |  | No |
| credentials | object |  | No |
| labels | object |  | No |
| created_at | dateTime |  | No |
| updated_at | dateTime |  | No |

#### CreateNamespaceResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string (uint64) |  | No |

#### CreateProviderRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| host | string |  | No |
| urn | string |  | No |
| name | string |  | No |
| type | string |  | No |
| credentials | object |  | No |
| labels | object |  | No |

#### CreateProviderResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string (uint64) |  | No |

#### CreateReceiverRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| name | string |  | No |
| type | string |  | No |
| labels | object |  | No |
| configurations | object |  | No |
| parent_id | string (uint64) |  | No |

#### CreateReceiverResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string (uint64) |  | No |

#### CreateSilenceRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| namespace_id | string (uint64) |  | No |
| type | string |  | No |
| target_id | string (uint64) |  | No |
| target_expression | object |  | No |

#### CreateSilenceResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string |  | No |

#### CreateSubscriptionRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| urn | string |  | No |
| namespace | string (uint64) |  | No |
| receivers | [ [ReceiverMetadata](#receivermetadata) ] |  | No |
| match | object |  | No |
| metadata | object |  | No |
| created_by | string |  | No |

#### CreateSubscriptionResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string (uint64) |  | No |

#### DeleteNamespaceResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| DeleteNamespaceResponse | object |  |  |

#### DeleteProviderResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| DeleteProviderResponse | object |  |  |

#### DeleteReceiverResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| DeleteReceiverResponse | object |  |  |

#### DeleteSubscriptionResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| DeleteSubscriptionResponse | object |  |  |

#### DeleteTemplateResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| DeleteTemplateResponse | object |  |  |

#### ExpireSilenceResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| ExpireSilenceResponse | object |  |  |

#### GetNamespaceResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| namespace | [Namespace](#namespace) |  | No |

#### GetProviderResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| provider | [Provider](#provider) |  | No |

#### GetReceiverResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| receiver | [Receiver](#receiver) |  | No |

#### GetSilenceResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| silence | [Silence](#silence) |  | No |

#### GetSubscriptionResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| subscription | [Subscription](#subscription) |  | No |

#### GetTemplateResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| template | [Template](#template) |  | No |

#### ListAlertsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| alerts | [ [Alert](#alert) ] |  | No |

#### ListNamespacesResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| namespaces | [ [Namespace](#namespace) ] |  | No |

#### ListNotificationMessagesResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| messages | [ [NotificationMessage](#notificationmessage) ] |  | No |

#### ListNotificationsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| notifications | [ [Notification](#notification) ] |  | No |

#### ListProvidersResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| providers | [ [Provider](#provider) ] |  | No |

#### ListReceiversResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| receivers | [ [Receiver](#receiver) ] |  | No |

#### ListRulesResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| rules | [ [Rule](#rule) ] |  | No |

#### ListSilencesResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| silences | [ [Silence](#silence) ] |  | No |

#### ListSubscriptionsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| subscriptions | [ [Subscription](#subscription) ] |  | No |

#### ListTemplatesResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| templates | [ [Template](#template) ] |  | No |

#### Namespace

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string (uint64) |  | No |
| urn | string |  | No |
| name | string |  | No |
| provider | string (uint64) |  | No |
| credentials | object |  | No |
| labels | object |  | No |
| created_at | dateTime |  | No |
| updated_at | dateTime |  | No |

#### Notification

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string |  | No |
| namespace_id | string (uint64) |  | No |
| type | string |  | No |
| data | object |  | No |
| labels | object |  | No |
| valid_duration | string |  | No |
| template | string |  | No |
| create_at | dateTime |  | No |
| unique_key | string |  | No |
| receiver_selectors | [ [ReceiverSelector](#receiverselector) ] |  | No |

#### NotificationMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string |  | No |
| notification_id | string |  | No |
| status | string |  | No |
| receiver_type | string |  | No |
| details | object |  | No |
| last_error | string |  | No |
| max_tries | string (uint64) |  | No |
| try_count | string (uint64) |  | No |
| retryable | boolean |  | No |
| expired_at | dateTime |  | No |
| created_at | dateTime |  | No |
| updated_at | dateTime |  | No |
| configs | object |  | No |

#### NotifyReceiverResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| NotifyReceiverResponse | object |  |  |

#### NullValue

`NullValue` is a singleton enumeration to represent the null value for the
`Value` type union.

The JSON representation for `NullValue` is JSON `null`.

- NULL_VALUE: Null value.

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| NullValue | string | `NullValue` is a singleton enumeration to represent the null value for the `Value` type union.  The JSON representation for `NullValue` is JSON `null`.   - NULL_VALUE: Null value. |  |

#### PostNotificationRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| receivers | [ object ] |  | No |
| data | object |  | No |
| labels | object |  | No |
| template | string |  | No |

#### PostNotificationResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| notification_id | string |  | No |

#### Provider

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string (uint64) |  | No |
| host | string |  | No |
| urn | string |  | No |
| name | string |  | No |
| type | string |  | No |
| credentials | object |  | No |
| labels | object |  | No |
| created_at | dateTime |  | No |
| updated_at | dateTime |  | No |

#### Receiver

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string (uint64) |  | No |
| name | string |  | No |
| type | string |  | No |
| labels | object |  | No |
| configurations | object |  | No |
| data | object |  | No |
| created_at | dateTime |  | No |
| updated_at | dateTime |  | No |
| parent_id | string (uint64) |  | No |

#### ReceiverMetadata

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string (uint64) |  | No |
| configuration | object |  | No |

#### ReceiverSelector

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| receiver_selector | object |  | No |

#### RenderTemplateResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| body | string |  | No |

#### Rule

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string (uint64) |  | No |
| name | string |  | No |
| enabled | boolean |  | No |
| group_name | string |  | No |
| namespace | string |  | No |
| template | string |  | No |
| variables | [ [Variables](#variables) ] |  | No |
| created_at | dateTime |  | No |
| updated_at | dateTime |  | No |
| provider_namespace | string (uint64) |  | No |

#### SetConfigResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| SetConfigResponse | object |  |  |

#### Silence

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string |  | No |
| namespace_id | string (uint64) |  | No |
| type | string |  | No |
| target_id | string (uint64) |  | No |
| target_expression | object |  | No |
| created_at | dateTime |  | No |
| updated_at | dateTime |  | No |
| deleted_at | dateTime |  | No |

#### Status

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| code | integer |  | No |
| message | string |  | No |
| details | [ [Any](#any) ] |  | No |

#### Subscription

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string (uint64) |  | No |
| urn | string |  | No |
| namespace | string (uint64) |  | No |
| receivers | [ [ReceiverMetadata](#receivermetadata) ] |  | No |
| match | object |  | No |
| created_at | dateTime |  | No |
| updated_at | dateTime |  | No |
| metadata | object |  | No |
| created_by | string |  | No |
| updated_by | string |  | No |

#### SyncRuntimeConfigResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| labels | object |  | No |

#### Template

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string (uint64) |  | No |
| name | string |  | No |
| body | string |  | No |
| tags | [ string ] |  | No |
| created_at | dateTime |  | No |
| updated_at | dateTime |  | No |
| variables | [ [TemplateVariables](#templatevariables) ] |  | No |

#### TemplateVariables

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| name | string |  | No |
| type | string |  | No |
| default | string |  | No |
| description | string |  | No |

#### TransformToAlertsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| alerts | [ [Alert](#alert) ] |  | No |
| firing_num | string (uint64) |  | No |

#### UpdateNamespaceResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string (uint64) |  | No |

#### UpdateProviderResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string (uint64) |  | No |

#### UpdateReceiverResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string (uint64) |  | No |

#### UpdateRuleRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| enabled | boolean |  | No |
| group_name | string |  | No |
| namespace | string |  | No |
| template | string |  | No |
| variables | [ [Variables](#variables) ] |  | No |
| provider_namespace | string (uint64) |  | No |

#### UpdateRuleResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| rule | [Rule](#rule) |  | No |

#### UpdateSubscriptionResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string (uint64) |  | No |

#### UpsertRuleResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| UpsertRuleResponse | object |  |  |

#### UpsertTemplateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string (uint64) |  | No |
| name | string |  | No |
| body | string |  | No |
| tags | [ string ] |  | No |
| variables | [ [TemplateVariables](#templatevariables) ] |  | No |

#### UpsertTemplateResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string (uint64) |  | No |

#### Variables

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| name | string |  | No |
| value | string |  | No |
| type | string |  | No |
| description | string |  | No |
