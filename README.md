# Go API client for openapi

# Introduction

<div class=\"status-info\">
<p class=\"status-info-title\">Note</p>
This documents the collection of Confluent Cloud APIs. Each API documents its
<a href=\"#section/Versioning/API-Lifecycle-Policy\">lifecycle phase</a>. APIs
marked as Early Access or Preview are not ready for production usage. We're currently
working with a select group of customers to get feedback and iterate on these APIs.
</div>

Confluent Cloud APIs are a core building block of Confluent Cloud. You can use the APIs to
manage your own account or to integrate Confluent into your product.

Most of the APIs are organized around
<a href=\"http://en.wikipedia.org/wiki/Representational_State_Transfer\" target=\"_blank\">REST</a>
and the resources which make up Confluent Cloud. The APIs have predictable
resource-oriented URLs, transport data using JSON, and use standard HTTP verbs,
response codes, authentication, and design principles.

# Object Model

<div class=\"status-info\">
<p class=\"status-info-title\">Note</p>
This section describes the object model for many Confluent Cloud APIs, but not all.
The Connect v1 API group has a different object model. You can review the example
request and response bodies in <a href=\"#tag/Connectors-(v1)\">Connect v1 API</a>
to see its object model.
</div>

Confluent Cloud APIs are primarily designed to be declarative and intent-oriented. In other words, 
tell the API what you want (for example, throughput or SLOs) and it will figure out how to make it happen 
(for example, cluster sizing). A Confluent object acts as a \"record of intent\" — after you create the
object, Confluent Cloud will work tirelessly in the background to ensure that the object exists
as specified.

Confluent APIs represent objects in JSON with media-type `application/json`.

Many objects follow a model consisting of `spec` and `status`. An object's `spec` tells
Confluent the _desired state_ (specification) of the resource. The object may not be
immediately available or changes may not be immediately applied. For this reason,
many objects also have a `status` property that provides info about the
_current state_ of the resource. Confluent Cloud is continuously and actively managing
each resource's current state to match it's desired state.

All Confluent objects share a set of common properties:

* **api_version** – API objects have an `api_version` field indicating their API version.
* **kind** – API objects have a `kind` field indicating the kind of object it is.
* **id** – Each object in the API will have an identifier, indicated via its `id` field,
  and should be treated as an opaque string unless otherwise specified.

There are a number of other [standard properties](#standard-properties) and that you'll encounter
used by many API objects. And of course, objects have plenty of non-standard fields that are
specific to each object _kind_... this is what makes them interesting!

# Authentication

Confluent uses API keys and Java Web Tokens (JWTs) to integrate your applications
and workflows to your Confluent Cloud resources using the Confluent Cloud REST APIs.
Your applications and workflows must be authenticated and authorized in order to
access and manage Confluent Cloud resources.

## API keys

You can create and manage your API keys using the Confluent Cloud Console or
Confluent CLI. For more information, see [Use API Keys to Control Access in Confluent Cloud](https://docs.confluent.io/cloud/current/access-management/authenticate/api-keys/api-keys.html).

Confluent Cloud uses the following two categories of API keys:

* A **Cloud API key** grants access to the Confluent Cloud Management APIs,
  such as for Provisioning and Metrics integrations.
* A **resource-specific API key** grants access to a Confluent Kafka cluster
  (Kafka API key), a Confluent Cloud Schema Registry (Schema Registry API key),
  or a ksqlDB application.

Each Confluent Cloud API key is associated with a principal (specific user or
service account) and inherits the permissions granted to the owner.

* For example, if service account `Armageddon` is granted ACLs on Kafka cluster
  `neptune`, then a Kafka API Key for `neptune` owned by `Armageddon` will have
  these ACLs enforced.
* **Note:** API keys are automatically deleted when the associated user or service
  account is deleted (for example, when an employee leaves the company or moves to
  a new department and an SSO integration removes the Confluent Cloud user as they
  no longer require access).
* Confluent **strongly recommends** that you use service accounts for all
  production-critical access.

Confluent Cloud API keys grant access to Confluent Cloud resources, so **keep them secure**!
Do not share your API keys and secrets in publicly-accessible locations, such as 
GitHub or client-side code.

All API requests must be made over HTTPS. Calls made over plain HTTP will fail.
API requests without authentication will also fail.

To use an API key, you must send it in an `Authorization: Basic {credentials}` header.
Remember that HTTP Basic authentication requires you to provide your credentials as
the API key ID and associated API secret separated by a colon and encoded using Base64
format. For example, if your API key ID is `ABCDEFGH123456789` and the API key Secret 
is `XNCIW93I2L1SQPJSJ823K1LS902KLDFMCZPWEO`, then the authorization header is:

```text​
Authorization: Basic QUJDREVGR0gxMjM0NTY3ODk6WE5DSVc5M0kyTDFTUVBKU0o4MjNLMUxTOTAyS0xERk1DWlBXRU8=
```

You can generate this header example from the API key:

macOS:

```shell
$ echo -n \"ABCDEFGH123456789:XNCIW93I2L1SQPJSJ823K1LS902KLDFMCZPWEO\" | base64

```

Linux:

```shell
$ echo -n \"ABCDEFGH123456789:XNCIW93I2L1SQPJSJ823K1LS902KLDFMCZPWEO\" | base64 -w 0
```

## External OAuth

You can use [OAuth/OIDC support for Confluent Cloud](https://docs.confluent.io/cloud/current/access-management/authenticate/oauth/overview.html)
to authenticate and authorize access to applications and workloads for the
following Confluent Cloud REST APIs:

* **Kafka REST API**: [Kafka REST API for Clusters(V3)](https://docs.confluent.io/cloud/current/api.html#tag/Cluster-(v3)).
  For an API overview and examples, see [Cluster Management with Kafka REST API](https://docs.confluent.io/cloud/current/kafka-rest/kafka-rest-cc.html).
* **Schema Registry REST API**: [Schema Registry REST API for Schemas(V1)](https://docs.confluent.io/cloud/current/api.html#tag/Schemas-(v1))
  and [Subjects](https://docs.confluent.io/cloud/current/api.html#tag/Subjects-(v1)).
  For an API overview and examples, see [Schema Registry REST API for Confluent Cloud](https://docs.confluent.io/cloud/current/sr/sr-rest-apis.html).

## Confluent STS tokens

Confluent Security Token Service (STS) issues access tokens (`confluent-sts-access-token`)
by exchanging an external token for a `confluent-sts-access-token`. You can use
Confluent STS tokens to authenticate to Confluent Cloud APIs that support the
`confluent-sts-access-token` notation.

To find out an API operation supports Confluent STS tokens, look in the **AUTHORIZATIONS**
listing for `confluent-sts-access-token`.

## Partner OAuth

Approved partners can fetch Partner tokens (`oauth`) that validate their identity
and grant access to the Partner API (`partner/v2`), which lets them sign up
an organization on behalf of a customer, manage entitlements (create, read, and list),
and read or list organizations they have signed up.

<!-- TODO: port this back to the Confluent API Design Guide -->

<SecurityDefinitions />

# Errors

<div class=\"status-info\">
<p class=\"status-info-title\">Note</p>
This section describes the structure of error responses for many Confluent Cloud APIs, but not all.
The Connect v1 API group has a different set of structures for error responses. Please review the example
request and response bodies in the Connect v1 API documentation <a href=\"#tag/Connectors-(v1)\">below</a>
to see its error behaviour.
</div>

Confluent uses conventional [HTTP status codes](#section/HTTP-Guidelines/Status-Codes) to
indicate the success or failure of an API request.

Failures follow a standard model to tell you about what went wrong. They may include
one or more error objects with the following fields:

| Field      | Type    | Description
|------------|---------|--------------------------------------
| id*        | UUID    | A unique identifier for this particular occurrence of the problem.
| status     | String  | The HTTP status code applicable to this problem.
| code       | String  | An application-specific error code.
| title      | String  | A short, human-readable summary of the problem that **should not** change from occurrence to occurrence of the problem, except for purposes of localization.
| detail*    | String  | A human-readable explanation specific to this occurrence of the problem. Like title, this field’s value can be localized.
| source     | Object  | An object that references the source of the error, and optionally includes any of the following members:
| &nbsp;&nbsp;pointer   | String  | A <a href=\"https://tools.ietf.org/html/rfc6901\" target=\"_blank\">JSON Pointer</a> to the associated entity in the request document (e.g. `\"/spec/title\"` for a specific attribute).
| &nbsp;&nbsp;parameter | String  | A string indicating which URI query parameter caused the error.
| meta       | Object  | A meta object that contains non-standard meta-information about the error.
| resolution | String  | Instructions for the end-user for correcting the error.

\\* indicates a required field

All errors include an `id` and some `detail` message. The `id` is a unique identifier — use it
when you're working with Confluent support to debug a problem with a specific API call. The
`detail` describes what went wrong.

Some errors that could be handled programmatically (e.g., a Kafka cluster config is invalid)
may include an error `code` that briefly explains the error reported.

Validation issues and similar errors include a `source` which tells you exactly
what in the request was responsible for the error.

For example, a failure may look like

    {
      \"errors\": [{
        \"status\": \"422\",
        \"code\": \"invalid_configuration\",
        \"id\": \"30ce6058-87da-11e4-b116-123b93f75cba\",
        \"title\": \"The Kafka cluster configuration is invalid\",
        \"detail\": \"The property '/cluster/storage_size' of type string did not match the following type: integer\",
        \"source\": {
          \"pointer\": \"/cluster/storage_size\"
        }
      }]
    }

If a request fails validation, it will return an HTTP `422 Unprocessable Entity`
with a list of fields that failed validation.

# Pagination

<div class=\"status-info\">
<p class=\"status-info-title\">Note</p>
This section describes the pagination behavior of “list” operations for many Confluent Cloud APIs, but not all.
The Connect v1 API list operations do not support pagination.
</div>

All API resources have support for bulk reads via \"list\" API operations. For example,
you can \"list Kafka clusters\", \"list api keys\", and \"list environments\". These \"list\"
operations require pagination; by requesting smaller subsets of data, API clients
receive a response much faster than requesting the entire, potentially large, data set.

All \"list\" operations follow the same pattern with the following parameters:

* `page_size` –  client-provided max number of items per page, only valid on the first request.
* `page_token` –  server-generated token used for traversing through the result set.

A paginated response may include any of the following pagination links. API clients may
follow the respective link to page forward or backward through the result set as desired.

| [Link Relation](https://www.iana.org/assignments/link-relations/link-relations.xml) | Description
|---------|---------------------------------------
| `next`  | A link to the next page of results. A response that does not contain a next link does not have further data to fetch.
| `prev`  | A link to the previous page of results. A response that does not contain a prev link has no previous data. This link is **optional** for collections that cannot be traversed backward.
| `first` | A link to the first page of results. This link is **optional** for collections that cannot be indexed directly to a given page.
| `last`  | A link to the last page of results. This link is **optional** for collections that cannot be indexed directly to a given page.

API clients must treat pagination links and the `page_token` parameter in particular as an opaque string. 

An example paginated list response may look like
```
{
    \"api_version\": \"v2\",
    \"kind\": \"KafkaClusterList\",
    \"metadata\": {
        \"next\": \"https://api.confluent.cloud/kafka-clusters?page_token=ABCDEFGHIJKLMNOP1234567890\"
    }
    \"data\": [
        {
            \"metadata\": {
                \"id\": \"lkc-abc123\",
                \"self\": \"https://api.confluent.cloud/kafka-clusters/lkc-abc123\",
                \"resource_name\": \"crn://confluent.cloud/kafka=lkc-abc123\",
            }
            \"spec\": {
                \"display_name\": \"My Kafka Cluster\",
                <snip>
            },
            \"status\": {
                \"phase\": \"RUNNING\",
                <snip>
            }
        },
        <snip>
    ]
}
```

# Rate Limiting

To protect the stability of the API and keep it available to all users, Confluent employs
multiple safeguards. If you send too many requests in quick succession or perform too many
concurrent operations, you may be throttled or have your request rejected with an error.

When a rate limit is breached, an HTTP `429 Too Many Requests` error is returned.
<!-- with
the following header:

| Header                  | Description
|-------------------------|----------------------------------------
| `Retry-After`           | The number of seconds to wait until the rate limit window resets. Only sent when the rate limit is reached.

-->
<!-- TODO make this true

Confluent enforces multiple kinds of limits, including request rate and concurrency limits, both
per user and organization wide. Unauthenticated requests are associated with the originating
IP address, and not the user making requests.

-->

Integrations should gracefully handle these limits by watching for `429` error responses and
building in a retry mechanism. This mechanism should follow a capped exponential backoff policy to
prevent [retry amplification](https://landing.google.com/sre/sre-book/chapters/addressing-cascading-failures/)
(\"retry storms\") and also introduce some randomness (\"jitter\") to avoid the
[thundering herd effect](https://en.wikipedia.org/wiki/Thundering_herd_problem).


If you’re running into this error and think you need a higher rate limit, contact Confluent at
[support@confluent.io](mailto:support@confluent.io).

# Identifiers and URLs

Most resources have multiple identifiers:
* `id` is the \"natural identifier\" for an object. It is only unique within its parent resource.
  The `id` is unique across time: the ID will not be reclaimed and reused after an object is deleted.
* `resource_name` is a Uniform Resource Identifier (URI) that is globally unique across all resources.
  This encompasses all parent resource `kind`s and `id`s necessary to uniquely identify a particular
  instance of this object `kind`. Because it uses object `id`s, the CRN will not be reclaimed and
  reused after an object is deleted. It is represented as a Confluent Resource Name (see below). 
* `self` is a Uniform Resource Locator (URL) at which an object can be addressed.
  This URL encodes the service location, API version, and other particulars necessary to
  locate the resource at a point in time.

To see how these relate to each other, consider `KafkaBroker` with `broker.id=2` in a `KafkaCluster`
in Confluent Cloud identified as `lkc-xsi8201`. In such an example, the `KafkaBroker` has `id=2`,
the `resource_name` is `crn://confluent.cloud/kafka=lkc-xsi8201/broker=2` and the `self` URL may be
something like `https://pkc-8wlk2n.us-west-2.aws.confluent.cloud`. Note that different identifiers
carry different information for different purposes, but the `resource_name` is the most complete
and canonical identifier.

## Confluent Resource Names (CRNs)

*Confluent Resource Names* (CRNs) are used to uniquely identify all Confluent resources.

A CRN is a valid URI having an \"authority\" of `confluent.cloud` or a self-managed
<a href=\"https://docs.confluent.io/current/security/rbac/configure-mds/index.html\" target=\"_blank\">
metadata service URL</a>, followed by the minimal hierarchical set of key-value
pairs necessary to uniquely identify a resource.

Here are some examples for basic resources in Confluent Cloud:

| Resource                   | Example CRN                                                                                                                                                              |
|----------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Organization               | crn://confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a                                                                                                  |
| Environment                | crn://confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/environment=env-456xy                                                                            |
| User                       | crn://confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/user=u-rst9876                                                                                   |
| API Key                    | crn://confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/user=u-zyx98/api-key=ABCDEFG9876543210                                                           |
| Service Account            | crn://confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/service-account=sa-abc1234                                                                       |
| Kafka Cluster              | crn://confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/environment=env-456xy/cloud-cluster=lkc-123abc/kafka=lkc-123abc                                  |
| Kafka Topic                | crn://confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/environment=env-456xy/cloud-cluster=lkc-123abc/kafka=lkc-123abc/topic=my_kafka_topic             |
| Consumer Group             | crn://confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/environment=env-456xy/cloud-cluster=lkc-123abc/kafka=lkc-123abc/group=confluent_cli_consumer_123 |
| Network                    | crn://confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/environment=env-456xy/network=n-123abc                                                           |
| Peering                    | crn://confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/environment=env-456xy/network=n-123abc/peering=p-123abc                                          |
| Private Link Access        | crn://confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/environment=env-456xy/network=n-123abc/private-link-access=pla-123abc                            |
| Transit Gateway Attachment | crn://confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/environment=env-456xy/network=n-123abc/transit-gateway-attachment=tgwa-123abc                    |
| Schema Registry Cluster    | crn://confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/environment=env-456xy/schema-registry=lsrc-789qw                                                 |
| Schema Subject             | crn://confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/environment=env-456xy/schema-registry=lsrc-789qw/subject=test                                    |

# Data Types

## Primitive Types

| Data Type  | Representation
|------------|---------------------
| Integers   | Each API may specify the type as `int32` or `int64`. Note that many languages, including JavaScript, are limited to a max size of approx `2**53` and don't correctly handle large `int64` values with their default JSON parser.
| Dates      | <a href=\"https://www.ietf.org/rfc/rfc3339.txt\" target=\"_blank\">RFC 3339</a> formatted string. UTC timezones are assumed, unless otherwise given.
| Times      | <a href=\"https://www.ietf.org/rfc/rfc3339.txt\" target=\"_blank\">RFC 3339</a> formatted string. UTC timezones are assumed, unless otherwise given.
| Durations  | <a href=\"https://www.ietf.org/rfc/rfc3339.txt\" target=\"_blank\">RFC 3339</a> formatted string.
| Periods    | <a href=\"https://www.ietf.org/rfc/rfc3339.txt\" target=\"_blank\">RFC 3339</a> formatted string. UTC timezones are assumed, unless otherwise given.
| Ranges     | All ranges are represented using half-open intervals with naming conventions like `[start_XXX, end_XXX)` such as `[start_time, end_time)`.
| Enums      | Most APIs use <a href=\"https://opensource.zalando.com/restful-api-guidelines/#112\" target=\"_blank\">`x-extensible-enum`</a> as an open-ended list of values. This improves compatibility compared with a standard `enum` which by definition represents a closed set. All enums have a `0`-valued entry which either serves as the default for common cases, or represents `UNSPECIFIED` when no default exists and results in an error.

<!-- TODO
### Standard Objects

| Money Object | https://schema.org/MonetaryAmount or https://opensource.zalando.com/restful-api-guidelines/#173
| Price Specification | https://schema.org/PriceSpecification -> https://schema.org/UnitPriceSpecification and https://schema.org/PaymentChargeSpecification
-->

### Standard Properties

Confluent uses this set of standard properties to ensure common concepts use
the same name and semantics across different APIs.

| Name             | Description
|------------------|------------------------------------------
| **api_version**  | Many API objects have an `api_version` field indicating their API version. See the [Object Model](#section/Object-Model).
| **kind**         | Many API objects have a `kind` field indicating the kind of object it is. See the [Object Model](#section/Object-Model).
| **id**           | Many objects in the API will have an identifier, indicated via its `id` field, and should be treated as an opaque string unless otherwise specified. See the [Object Model](#section/Object-Model).
| **name**         | Objects which support a client-provided unique identifier instead of a generated `id` will indicate this identifier via its `name` field.
| **display_name** | The human-readable display name of an API object.
| **title**        | The official name of an API object, such as a company name. It should be treated as the formal version of `display_name`.
| **description**  | One or more paragraphs of text description of an entity.
| **created_at**   | The date and time the object was created, represented as a string in <a href=\"https://www.ietf.org/rfc/rfc3339.txt\" target=\"_blank\">RFC 3339</a> format.
| **updated_at**   | The date and time the object was last modified, represented as a string in <a href=\"https://www.ietf.org/rfc/rfc3339.txt\" target=\"_blank\">RFC 3339</a> format.
| **deleted_at**   | If present, the date and time after which the object was/will be deleted, represented as a string in <a href=\"https://www.ietf.org/rfc/rfc3339.txt\" target=\"_blank\">RFC 3339</a> format.
| **page_token**   | The pagination token in the List request. See [Pagination](#section/Pagination).
| **page_size**    | The pagination size in the List request. See [Pagination](#section/Pagination).
| **total_size**   | The total count of items in the list irrespective of pagination. See [Pagination](#section/Pagination).
| **spec**         | The _desired state_ specification of the resource, as observed by Confluent Cloud.
| **status**       | The _current state_ of the resource, as observed by Confluent Cloud.

# Versioning

Confluent APIs ensure stability for your integrations by avoiding the introduction
of breaking changes to customers unexpectedly. Confluent will make non-breaking
API changes without advance notice. Thus, API clients **must**  follow the
[Compatibility Policy](#section/Versioning/Compatibility-Policy) below to ensure your
ingtegration remains stable. All APIs follow the API Lifecycle Policy described below,
which describes the guarantees API clients can rely on.

Breaking changes will be [widely communicated](#communication) in advance in accordance
with the Confluent [Deprecation Policy](#section/Versioning/Deprecation-Policy). Confluent will provide 
timelines and a migration path for all API changes, where available. Be sure to subscribe
to one or more [communication channels](#communication) so you don't miss any updates!

One exception to these guidelines is for critical security issues. Confluent will take any necessary
actions to mitigate any critical security issue as soon as possible, which may include disabling
the vulnerable functionality until a proper solution is available.

Do not consume any Confluent API unless it is documented in the API Reference. All undocumented
endpoints should be considered private, subject to change without notice, and not covered by any
agreements.

> Note: The version in the URL (e.g. \"v1\" or \"v2\") is not a \"major version\" in the
[Semantic Versioning](https://semver.org/) sense. It is a \"generational version\" or \"meta version\", as seen in
APIs like <a href=\"https://developer.github.com/v3/versions/\" target=\"_blank\">Github API</a> or the
<a href=\"https://stripe.com/docs/api/versioning\" target=\"_blank\">Stripe API</a>.

## API Groups

Confluent APIs are divided into API Groups, such as the Cluster Management for Apache Kafka (CMK) API group,
the Connect API group, and the Data Catalog API group. Each group has its own set of endpoints and resources,
as well as its own API group version.

Because different API groups have different versions, there is no single version for the \"Confluent Cloud API\".
The latest version of the Connect API group may be `connect/v1`, while the latest version of the CMK API group
may be `cmk/v2`.

When a breaking change is introduced into one API group, Confluent will increase the API version for that API group
only, leaving the other API groups' versions unchanged. This makes it easier for you to understand whether a given
breaking change impacts your usage of the APIs.

## Known Issues

During the Early Access and Preview periods, we have a few known issues.

| Issue          | Description                                                                   | Proposed Resolution
|----------------|-------------------------------------------------------------------------------|-----------------------------------------------------
| Quota Exceeded | Some \"Quota Exceeded\" errors will be returned as HTTP 400 instead of HTTP 402 | Return 402 consistently for \"Quota Exceeded\" errors 

## API Lifecycle Policy

The following status labels are applicable to APIs, features, and SDK versions, based on
the current support status of each:

* **Early Access** – May change at any time. Not recommended for production usage. Not officially supported by
  Confluent. Intended for user feedback only. Users must be granted explicit access to the API by Confluent.
* **Preview** – Unlikely to change between Preview and General Availability. Not recommended for production usage.
  Officially supported by Confluent for non-production usage. Accessible to all users.
* **Limited Availability (LA)** - Available to key select customers in a subset of regions/providers/networks and recommended for production usage.  
* **Generally Available (GA)** – Will not change at short notice. Recommended for production usage.
  Officially supported by Confluent for non-production and production usage.
* **Deprecated** – Still supported, but no longer under active development. Existing usage will continue to function
  but migration following the upgrade guide is strongly recommended. New use cases should be built against the new
  version. Deprecated feature or version will be removed in the future at the announced date.
* **Sunset** – Removed, and no longer supported or available.

An API is \"Generally Available\" unless explicitly marked otherwise.

## Compatibility Policy

Confluent Cloud APIs are governed by
<a href=\"https://docs.confluent.io/cloud/current/clusters/upgrade-policy.html\" target=\"_blank\">
Confluent Cloud Upgrade Policy</a>, which means that backward incompatible changes and
deprecations will be made approximately once per year, and 180 days notice will be provided via email to all
registered Confluent Cloud users.

### Backward Compatibility

> _An API version is backward compatible if a program written against the previous version of the API will continue to work the same way, without modification, against this version of the API._

Confluent considers the following changes to be backward compatible:

* Adding new API resources.
* Adding new optional parameters to existing API requests (e.g., query string).
* Adding new properties to existing API resources (e.g., request body).
* Changing the order of properties in existing API responses.
* Changing the length or format of object IDs or other opaque strings.
  * Unless otherwise documented, you can safely assume object IDs generated by Confluent will never exceed 255
    characters, but you should be able to handle IDs of up to that length. If you're using MySQL,
    for example, you should store IDs in a `VARCHAR(255) COLLATE utf8_bin` column.
  * This includes adding or removing fixed prefixes (such as `lkc-` on Kafka cluster IDs).
  * This includes API keys, API tokens, and similar authentication mechanisms.
  * This includes all strings described as \"opaque\" in the docs, such as pagination cursors.
* Adding new API event types.
* Adding new properties to existing API event types.
* Omitting properties with null values from existing API responses.

### Forward Compatibility

> _An API version is forward compatible if a program written against the next version of the API
> will continue to work the same way, without modification, against this version of the API._

In other words, a forward compatible API will accept input intended for a later version of itself.

Confluent does not guarantee the forward compatibility of the APIs, but Confluent does generally follow the guidelines
given by the [Robustness principle](https://en.wikipedia.org/wiki/Robustness_principle).
This means that the API determines what to do with a request based only on the parts that it recognizes.

This is often referred to as the MUST IGNORE rule.

* Request parameters that are not recognized will be ignored (e.g., query string).
* Request properties that are not recognized will be ignored (e.g., request body).
* Request metadata that are not recognized will be ignored (e.g., request headers).

API clients must also follow the MUST IGNORE rule.

* Response properties that are not recognized must be ignored (e.g., response body).
* Response metadata that are not recognized must be ignored (e.g., response headers).

Additionally, there is a more subtle related rule called the MUST FORWARD rule. Any parts of
a request that an API doesn't recognize must be forwarded unchanged.

* Response properties that are not recognized must be included in any input subsequent updates (e.g., request body)
  * This includes future `PUT` requests in a read/modify/write operation.
    (This isn't required for `PATCH` partial updates, which is why Confluent APIs use `PATCH`.)
* Event processors must not strip unknown properties before forwarding messages.

### Client Responsibilities

* Resource and rate limits, and the default and maximum sizes of paginated data **are not**
  considered part of the API contract and may change (possibly dynamically). It is the client's
  responsibility to read the road signs and obey the speed limit.
* If a property has a primitive type and the API documentation does not explicitly limit its
  possible values, clients **must not** assume the values are constrained to a particular set
  of possible responses.
* If a property of an object is not explicitly declared as mandatory in the API, clients
  **must not** assume it will be present.
* A resource **may** be modified to return a \"redirection\" response (e.g. `301`, `307`) instead of
  directly returning the resource. Clients **must** handle HTTP-level redirects, and respect HTTP
  headers (e.g. `Location`).

## Deprecation Policy

Confluent will announce deprecations at least 180 days in advance of a breaking change
and will continue to maintain the deprecated APIs in their original form during this time.

Exceptions to this policy apply in case of critical security vulnerabilities or functional defects.

### Communication

When a deprecation is announced, the details and any relevant migration
information will be available on one or more of the following channels:

* Announcements on the <a href=\"https://www.confluent.io/blog/\" target=\"_blank\">Developer Blog</a>,
  <a href=\"https://confluentcommunity.slack.com\" target=\"_blank\">Community Slack</a>
  (<a href=\"https://slackpass.io/confluentcommunity\" target=\"_blank\">join!</a>),
  <a href=\"https://groups.google.com/forum/#!forum/confluent-platform\" target=\"_blank\">Google Group</a>,
  the <a href=\"https://twitter.com/ConfluentInc\" target=\"_blank\">@ConfluentInc twitter</a>
  account, and similar channels
* Enterprise customers may receive information by email to their specified Confluent contact, if applicable.

<!-- TODO:
### Discoverability
-->

# HTTP Guidelines

## Status Codes

Confluent respects the meanings and behavior of HTTP status codes as defined
in <a href=\"https://tools.ietf.org/html/rfc2616\">RFC2616</a> and elsewhere.

* Codes in the `2xx` range indicate success
* Codes in the `3xx` range indicate redirection
* Codes in the `4xx` range indicate an error caused by the client request
  (e.g., a required parameter was omitted, an invalid cluster configuration was provided, etc.)
* Codes in the `5xx` range indicate an error with Confluent's servers (these are rare)

The various HTTP status codes that might be returned are listed below.

| Code | Title                  | Description
|------|------------------------|--------------------------------
| 200  | OK                     | Everything worked as expected.
| 201  | Created                | The resource was created. Follow the `Location` header.
| 204  | No Content             | Everything worked and there is no content to return.
| 400  | Bad Request         | The request was unacceptable, often due to malformed syntax, or a missing or malformed parameter.
| 401  | Unauthorized           | No valid credentials provided. or the credentials are unsuitable, invalid, or unauthorized.
| 402  | Over Quota             | The request was valid, but you've exceeded your plan quota or limits.
| 404  | Not Found              | The requested resource doesn't exist or you're unauthorized to know it exists.
| 409  | Conflict               | The request conflicts with another request (perhaps it already exists or was based on a stale version of data).
| 422  | Validation Failed      | The request was parsed correctly but failed some sort of validation.
| 429  | Too Many Requests      | Too many requests hit the API too quickly. Confluent recommends an exponential backoff of your requests.
| 500, 502, 503, 504 | Server Errors | Something went wrong on Confluent's end. (These are rare.)

This list is not exhaustive; other standard HTTP error codes may be used,
including `304`, `307`, `308`, `405`, `406`, `408`, `410`, and `415`.

For more details, see https://httpstatuses.com.

<!--

## Method Overriding

Some firewalls and HTTP clients restrict the use of verbs other than `GET` and `POST`. In those
environments, Confluent API operations that require `PUT`, `PATCH`, and `DELETE` verbs will be inaccessible.

To avoid this issue, Confluent APIs support the `X-HTTP-Method-Override` header, allowing clients to
\"tunnel\" `PUT`, `PATCH`, and `DELETE` requests via a `POST` request.

For example, to call a Confluent `PATCH` resource via a `POST` request, you can
include `X-HTTP-Method-Override: PATCH` as a header.

## User Agent Required

Confluent API requests **should** include a valid `User-Agent` header. Requests with no `User-Agent`
header may be rejected. You should use the name of your integration for the `User-Agent`
header value and include contact information so that Confluent can contact you if there are problems.

> If your integration is acting as a proxy or gateway, you **should** forward the User-Agent
> of the originating client with your API requests.

Here's a complete example:

    User-Agent: CoolToolName/1.2.3 (https://example.org/CoolTool/; CoolTool@example.org) UsedBaseLibrary/2.1.0

The minimum user agent string is the integration name and version: `name/version`.
You can string together multiple values in a space-separated list. The full syntax is:

    name/version [(comments)] [name/version [(comments)]] [...]​

For the integration name, use a string (without whitespace) that clearly and meaningfully
identifies your integration.

* Avoid ambiguous names: `Confluent-Integration`, `Kafka-Sink`
* Use clear and meaningful names: `ABCTools-ToolName`, `StackStorm-Confluent-Plugin`

For the version, use a semantic version, build ID, commit hash, or other identifier
that is updated automatically when you release a new version.

Wrap comments in parentheses `()` as a semi-colon separated list. Helpful comments to include:

* A public URL for your integration, such as a GitHub link or a page in your
  docs site that describes the integration.
* Contact information so that Confluent can easily reach the integration publisher. This
  information from the user agent string will never be shared nor used by Confluent for
  any purpose other than discussing the integration with its publisher.

If you provide an invalid `User-Agent` header, you may receive a `403 Forbidden` response.

-->


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.confluent.io/cloud-contact-us/](https://www.confluent.io/cloud-contact-us/)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

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

All URIs are relative to *https://api.confluent.cloud*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ACLV3Api* | [**BatchCreateKafkaAcls**](docs/ACLV3Api.md#batchcreatekafkaacls) | **Post** /kafka/v3/clusters/{cluster_id}/acls:batch | Batch Create ACLs
*ACLV3Api* | [**CreateKafkaAcls**](docs/ACLV3Api.md#createkafkaacls) | **Post** /kafka/v3/clusters/{cluster_id}/acls | Create an ACL
*ACLV3Api* | [**DeleteKafkaAcls**](docs/ACLV3Api.md#deletekafkaacls) | **Delete** /kafka/v3/clusters/{cluster_id}/acls | Delete ACLs
*ACLV3Api* | [**GetKafkaAcls**](docs/ACLV3Api.md#getkafkaacls) | **Get** /kafka/v3/clusters/{cluster_id}/acls | Search ACLs
*APIKeysIamV2Api* | [**CreateIamV2ApiKey**](docs/APIKeysIamV2Api.md#createiamv2apikey) | **Post** /iam/v2/api-keys | Create an API Key
*APIKeysIamV2Api* | [**DeleteIamV2ApiKey**](docs/APIKeysIamV2Api.md#deleteiamv2apikey) | **Delete** /iam/v2/api-keys/{id} | Delete an API Key
*APIKeysIamV2Api* | [**GetIamV2ApiKey**](docs/APIKeysIamV2Api.md#getiamv2apikey) | **Get** /iam/v2/api-keys/{id} | Read an API Key
*APIKeysIamV2Api* | [**ListIamV2ApiKeys**](docs/APIKeysIamV2Api.md#listiamv2apikeys) | **Get** /iam/v2/api-keys | List of API Keys
*APIKeysIamV2Api* | [**UpdateIamV2ApiKey**](docs/APIKeysIamV2Api.md#updateiamv2apikey) | **Patch** /iam/v2/api-keys/{id} | Update an API Key
*AppliedQuotasServiceQuotaV1Api* | [**GetServiceQuotaV1AppliedQuota**](docs/AppliedQuotasServiceQuotaV1Api.md#getservicequotav1appliedquota) | **Get** /service-quota/v1/applied-quotas/{id} | Read an Applied Quota
*AppliedQuotasServiceQuotaV1Api* | [**ListServiceQuotaV1AppliedQuotas**](docs/AppliedQuotasServiceQuotaV1Api.md#listservicequotav1appliedquotas) | **Get** /service-quota/v1/applied-quotas | List of Applied Quotas
*ClientQuotasKafkaQuotasV1Api* | [**CreateKafkaQuotasV1ClientQuota**](docs/ClientQuotasKafkaQuotasV1Api.md#createkafkaquotasv1clientquota) | **Post** /kafka-quotas/v1/client-quotas | Create a Client Quota
*ClientQuotasKafkaQuotasV1Api* | [**DeleteKafkaQuotasV1ClientQuota**](docs/ClientQuotasKafkaQuotasV1Api.md#deletekafkaquotasv1clientquota) | **Delete** /kafka-quotas/v1/client-quotas/{id} | Delete a Client Quota
*ClientQuotasKafkaQuotasV1Api* | [**GetKafkaQuotasV1ClientQuota**](docs/ClientQuotasKafkaQuotasV1Api.md#getkafkaquotasv1clientquota) | **Get** /kafka-quotas/v1/client-quotas/{id} | Read a Client Quota
*ClientQuotasKafkaQuotasV1Api* | [**ListKafkaQuotasV1ClientQuotas**](docs/ClientQuotasKafkaQuotasV1Api.md#listkafkaquotasv1clientquotas) | **Get** /kafka-quotas/v1/client-quotas | List of Client Quotas
*ClientQuotasKafkaQuotasV1Api* | [**UpdateKafkaQuotasV1ClientQuota**](docs/ClientQuotasKafkaQuotasV1Api.md#updatekafkaquotasv1clientquota) | **Patch** /kafka-quotas/v1/client-quotas/{id} | Update a Client Quota
*ClusterLinkingV3Api* | [**CreateKafkaLink**](docs/ClusterLinkingV3Api.md#createkafkalink) | **Post** /kafka/v3/clusters/{cluster_id}/links | Create a cluster link
*ClusterLinkingV3Api* | [**CreateKafkaMirrorTopic**](docs/ClusterLinkingV3Api.md#createkafkamirrortopic) | **Post** /kafka/v3/clusters/{cluster_id}/links/{link_name}/mirrors | Create a mirror topic
*ClusterLinkingV3Api* | [**DeleteKafkaLink**](docs/ClusterLinkingV3Api.md#deletekafkalink) | **Delete** /kafka/v3/clusters/{cluster_id}/links/{link_name} | Delete the cluster link
*ClusterLinkingV3Api* | [**DeleteKafkaLinkConfig**](docs/ClusterLinkingV3Api.md#deletekafkalinkconfig) | **Delete** /kafka/v3/clusters/{cluster_id}/links/{link_name}/configs/{config_name} | Reset the given config to default value
*ClusterLinkingV3Api* | [**GetKafkaLink**](docs/ClusterLinkingV3Api.md#getkafkalink) | **Get** /kafka/v3/clusters/{cluster_id}/links/{link_name} | Describe the cluster link
*ClusterLinkingV3Api* | [**GetKafkaLinkConfigs**](docs/ClusterLinkingV3Api.md#getkafkalinkconfigs) | **Get** /kafka/v3/clusters/{cluster_id}/links/{link_name}/configs/{config_name} | Describe the config under the cluster link
*ClusterLinkingV3Api* | [**ListKafkaLinkConfigs**](docs/ClusterLinkingV3Api.md#listkafkalinkconfigs) | **Get** /kafka/v3/clusters/{cluster_id}/links/{link_name}/configs | List all configs of the cluster link
*ClusterLinkingV3Api* | [**ListKafkaLinks**](docs/ClusterLinkingV3Api.md#listkafkalinks) | **Get** /kafka/v3/clusters/{cluster_id}/links | List all cluster links in the dest cluster
*ClusterLinkingV3Api* | [**ListKafkaMirrorTopics**](docs/ClusterLinkingV3Api.md#listkafkamirrortopics) | **Get** /kafka/v3/clusters/{cluster_id}/links/-/mirrors | List mirror topics
*ClusterLinkingV3Api* | [**ListKafkaMirrorTopicsUnderLink**](docs/ClusterLinkingV3Api.md#listkafkamirrortopicsunderlink) | **Get** /kafka/v3/clusters/{cluster_id}/links/{link_name}/mirrors | List mirror topics
*ClusterLinkingV3Api* | [**ReadKafkaMirrorTopic**](docs/ClusterLinkingV3Api.md#readkafkamirrortopic) | **Get** /kafka/v3/clusters/{cluster_id}/links/{link_name}/mirrors/{mirror_topic_name} | Describe the mirror topic
*ClusterLinkingV3Api* | [**UpdateKafkaLinkConfig**](docs/ClusterLinkingV3Api.md#updatekafkalinkconfig) | **Put** /kafka/v3/clusters/{cluster_id}/links/{link_name}/configs/{config_name} | Alter the config under the cluster link
*ClusterLinkingV3Api* | [**UpdateKafkaLinkConfigBatch**](docs/ClusterLinkingV3Api.md#updatekafkalinkconfigbatch) | **Put** /kafka/v3/clusters/{cluster_id}/links/{link_name}/configs:alter | Batch Alter Cluster Link Configs
*ClusterLinkingV3Api* | [**UpdateKafkaMirrorTopicsFailover**](docs/ClusterLinkingV3Api.md#updatekafkamirrortopicsfailover) | **Post** /kafka/v3/clusters/{cluster_id}/links/{link_name}/mirrors:failover | Failover the mirror topics
*ClusterLinkingV3Api* | [**UpdateKafkaMirrorTopicsPause**](docs/ClusterLinkingV3Api.md#updatekafkamirrortopicspause) | **Post** /kafka/v3/clusters/{cluster_id}/links/{link_name}/mirrors:pause | Pause the mirror topics
*ClusterLinkingV3Api* | [**UpdateKafkaMirrorTopicsPromote**](docs/ClusterLinkingV3Api.md#updatekafkamirrortopicspromote) | **Post** /kafka/v3/clusters/{cluster_id}/links/{link_name}/mirrors:promote | Promote the mirror topics
*ClusterLinkingV3Api* | [**UpdateKafkaMirrorTopicsResume**](docs/ClusterLinkingV3Api.md#updatekafkamirrortopicsresume) | **Post** /kafka/v3/clusters/{cluster_id}/links/{link_name}/mirrors:resume | Resume the mirror topics
*ClusterV3Api* | [**GetKafkaCluster**](docs/ClusterV3Api.md#getkafkacluster) | **Get** /kafka/v3/clusters/{cluster_id} | Get Cluster
*ClustersCmkV2Api* | [**CreateCmkV2Cluster**](docs/ClustersCmkV2Api.md#createcmkv2cluster) | **Post** /cmk/v2/clusters | Create a Cluster
*ClustersCmkV2Api* | [**DeleteCmkV2Cluster**](docs/ClustersCmkV2Api.md#deletecmkv2cluster) | **Delete** /cmk/v2/clusters/{id} | Delete a Cluster
*ClustersCmkV2Api* | [**GetCmkV2Cluster**](docs/ClustersCmkV2Api.md#getcmkv2cluster) | **Get** /cmk/v2/clusters/{id} | Read a Cluster
*ClustersCmkV2Api* | [**ListCmkV2Clusters**](docs/ClustersCmkV2Api.md#listcmkv2clusters) | **Get** /cmk/v2/clusters | List of Clusters
*ClustersCmkV2Api* | [**UpdateCmkV2Cluster**](docs/ClustersCmkV2Api.md#updatecmkv2cluster) | **Patch** /cmk/v2/clusters/{id} | Update a Cluster
*ClustersKsqldbcmV2Api* | [**CreateKsqldbcmV2Cluster**](docs/ClustersKsqldbcmV2Api.md#createksqldbcmv2cluster) | **Post** /ksqldbcm/v2/clusters | Create a Cluster
*ClustersKsqldbcmV2Api* | [**DeleteKsqldbcmV2Cluster**](docs/ClustersKsqldbcmV2Api.md#deleteksqldbcmv2cluster) | **Delete** /ksqldbcm/v2/clusters/{id} | Delete a Cluster
*ClustersKsqldbcmV2Api* | [**GetKsqldbcmV2Cluster**](docs/ClustersKsqldbcmV2Api.md#getksqldbcmv2cluster) | **Get** /ksqldbcm/v2/clusters/{id} | Read a Cluster
*ClustersKsqldbcmV2Api* | [**ListKsqldbcmV2Clusters**](docs/ClustersKsqldbcmV2Api.md#listksqldbcmv2clusters) | **Get** /ksqldbcm/v2/clusters | List of Clusters
*ClustersSrcmV2Api* | [**CreateSrcmV2Cluster**](docs/ClustersSrcmV2Api.md#createsrcmv2cluster) | **Post** /srcm/v2/clusters | Create a Cluster
*ClustersSrcmV2Api* | [**DeleteSrcmV2Cluster**](docs/ClustersSrcmV2Api.md#deletesrcmv2cluster) | **Delete** /srcm/v2/clusters/{id} | Delete a Cluster
*ClustersSrcmV2Api* | [**GetSrcmV2Cluster**](docs/ClustersSrcmV2Api.md#getsrcmv2cluster) | **Get** /srcm/v2/clusters/{id} | Read a Cluster
*ClustersSrcmV2Api* | [**ListSrcmV2Clusters**](docs/ClustersSrcmV2Api.md#listsrcmv2clusters) | **Get** /srcm/v2/clusters | List of Clusters
*ClustersSrcmV2Api* | [**UpdateSrcmV2Cluster**](docs/ClustersSrcmV2Api.md#updatesrcmv2cluster) | **Patch** /srcm/v2/clusters/{id} | Update a Cluster
*CompatibilityV1Api* | [**TestCompatibilityBySubjectName**](docs/CompatibilityV1Api.md#testcompatibilitybysubjectname) | **Post** /compatibility/subjects/{subject}/versions/{version} | Test schema compatibility against a particular schema subject-version
*CompatibilityV1Api* | [**TestCompatibilityForSubject**](docs/CompatibilityV1Api.md#testcompatibilityforsubject) | **Post** /compatibility/subjects/{subject}/versions | Test schema compatibility against all schemas under a subject
*ConfigV1Api* | [**DeleteSubjectConfig**](docs/ConfigV1Api.md#deletesubjectconfig) | **Delete** /config/{subject} | Delete subject compatibility level
*ConfigV1Api* | [**DeleteTopLevelConfig**](docs/ConfigV1Api.md#deletetoplevelconfig) | **Delete** /config | Delete global compatibility level
*ConfigV1Api* | [**GetClusterConfig**](docs/ConfigV1Api.md#getclusterconfig) | **Get** /clusterconfig | Get cluster config
*ConfigV1Api* | [**GetSubjectLevelConfig**](docs/ConfigV1Api.md#getsubjectlevelconfig) | **Get** /config/{subject} | Get subject compatibility level
*ConfigV1Api* | [**GetTopLevelConfig**](docs/ConfigV1Api.md#gettoplevelconfig) | **Get** /config | Get global compatibility level
*ConfigV1Api* | [**UpdateSubjectLevelConfig**](docs/ConfigV1Api.md#updatesubjectlevelconfig) | **Put** /config/{subject} | Update subject compatibility level
*ConfigV1Api* | [**UpdateTopLevelConfig**](docs/ConfigV1Api.md#updatetoplevelconfig) | **Put** /config | Update global compatibility level
*ConfigsV3Api* | [**DeleteKafkaClusterConfig**](docs/ConfigsV3Api.md#deletekafkaclusterconfig) | **Delete** /kafka/v3/clusters/{cluster_id}/broker-configs/{name} | Reset Cluster Config
*ConfigsV3Api* | [**DeleteKafkaTopicConfig**](docs/ConfigsV3Api.md#deletekafkatopicconfig) | **Delete** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/configs/{name} | Reset Topic Config
*ConfigsV3Api* | [**GetKafkaClusterConfig**](docs/ConfigsV3Api.md#getkafkaclusterconfig) | **Get** /kafka/v3/clusters/{cluster_id}/broker-configs/{name} | Get Cluster Config
*ConfigsV3Api* | [**GetKafkaTopicConfig**](docs/ConfigsV3Api.md#getkafkatopicconfig) | **Get** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/configs/{name} | Get Topic Config
*ConfigsV3Api* | [**ListKafkaAllTopicConfigs**](docs/ConfigsV3Api.md#listkafkaalltopicconfigs) | **Get** /kafka/v3/clusters/{cluster_id}/topics/-/configs | Get All Topic Configs
*ConfigsV3Api* | [**ListKafkaClusterConfigs**](docs/ConfigsV3Api.md#listkafkaclusterconfigs) | **Get** /kafka/v3/clusters/{cluster_id}/broker-configs | List Cluster Configs
*ConfigsV3Api* | [**ListKafkaDefaultTopicConfigs**](docs/ConfigsV3Api.md#listkafkadefaulttopicconfigs) | **Get** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/default-configs | List Default Topic Configs
*ConfigsV3Api* | [**ListKafkaTopicConfigs**](docs/ConfigsV3Api.md#listkafkatopicconfigs) | **Get** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/configs | List Topic Configs
*ConfigsV3Api* | [**UpdateKafkaClusterConfig**](docs/ConfigsV3Api.md#updatekafkaclusterconfig) | **Put** /kafka/v3/clusters/{cluster_id}/broker-configs/{name} | Update Cluster Config
*ConfigsV3Api* | [**UpdateKafkaClusterConfigs**](docs/ConfigsV3Api.md#updatekafkaclusterconfigs) | **Post** /kafka/v3/clusters/{cluster_id}/broker-configs:alter | Batch Alter Cluster Configs
*ConfigsV3Api* | [**UpdateKafkaTopicConfig**](docs/ConfigsV3Api.md#updatekafkatopicconfig) | **Put** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/configs/{name} | Update Topic Config
*ConfigsV3Api* | [**UpdateKafkaTopicConfigBatch**](docs/ConfigsV3Api.md#updatekafkatopicconfigbatch) | **Post** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/configs:alter | Batch Alter Topic Configs
*ConnectorsV1Api* | [**CreateConnectv1Connector**](docs/ConnectorsV1Api.md#createconnectv1connector) | **Post** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors | Create a Connector
*ConnectorsV1Api* | [**CreateOrUpdateConnectv1ConnectorConfig**](docs/ConnectorsV1Api.md#createorupdateconnectv1connectorconfig) | **Put** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors/{connector_name}/config | Create or Update a Connector Configuration
*ConnectorsV1Api* | [**DeleteConnectv1Connector**](docs/ConnectorsV1Api.md#deleteconnectv1connector) | **Delete** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors/{connector_name} | Delete a Connector
*ConnectorsV1Api* | [**GetConnectv1ConnectorConfig**](docs/ConnectorsV1Api.md#getconnectv1connectorconfig) | **Get** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors/{connector_name}/config | Read a Connector Configuration
*ConnectorsV1Api* | [**ListConnectv1Connectors**](docs/ConnectorsV1Api.md#listconnectv1connectors) | **Get** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors | List of Connectors
*ConnectorsV1Api* | [**ListConnectv1ConnectorsWithExpansions**](docs/ConnectorsV1Api.md#listconnectv1connectorswithexpansions) | **Get** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors?expand&#x3D;info,status,id | List of Connectors with Expansions
*ConnectorsV1Api* | [**ReadConnectv1Connector**](docs/ConnectorsV1Api.md#readconnectv1connector) | **Get** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors/{connector_name} | Read a Connector
*ConsumerGroupV3Api* | [**GetKafkaConsumer**](docs/ConsumerGroupV3Api.md#getkafkaconsumer) | **Get** /kafka/v3/clusters/{cluster_id}/consumer-groups/{consumer_group_id}/consumers/{consumer_id} | Get Consumer
*ConsumerGroupV3Api* | [**GetKafkaConsumerGroup**](docs/ConsumerGroupV3Api.md#getkafkaconsumergroup) | **Get** /kafka/v3/clusters/{cluster_id}/consumer-groups/{consumer_group_id} | Get Consumer Group
*ConsumerGroupV3Api* | [**GetKafkaConsumerGroupLagSummary**](docs/ConsumerGroupV3Api.md#getkafkaconsumergrouplagsummary) | **Get** /kafka/v3/clusters/{cluster_id}/consumer-groups/{consumer_group_id}/lag-summary | Get Consumer Group Lag Summary
*ConsumerGroupV3Api* | [**GetKafkaConsumerLag**](docs/ConsumerGroupV3Api.md#getkafkaconsumerlag) | **Get** /kafka/v3/clusters/{cluster_id}/consumer-groups/{consumer_group_id}/lags/{topic_name}/partitions/{partition_id} | Get Consumer Lag
*ConsumerGroupV3Api* | [**ListKafkaConsumerGroups**](docs/ConsumerGroupV3Api.md#listkafkaconsumergroups) | **Get** /kafka/v3/clusters/{cluster_id}/consumer-groups | List Consumer Groups
*ConsumerGroupV3Api* | [**ListKafkaConsumerLags**](docs/ConsumerGroupV3Api.md#listkafkaconsumerlags) | **Get** /kafka/v3/clusters/{cluster_id}/consumer-groups/{consumer_group_id}/lags | List Consumer Lags
*ConsumerGroupV3Api* | [**ListKafkaConsumers**](docs/ConsumerGroupV3Api.md#listkafkaconsumers) | **Get** /kafka/v3/clusters/{cluster_id}/consumer-groups/{consumer_group_id}/consumers | List Consumers
*ContextsV1Api* | [**ListContexts**](docs/ContextsV1Api.md#listcontexts) | **Get** /contexts | List contexts
*CostsBillingV1Api* | [**ListBillingV1Costs**](docs/CostsBillingV1Api.md#listbillingv1costs) | **Get** /billing/v1/costs | List of Costs
*EntitlementsPartnerV2Api* | [**CreatePartnerV2Entitlement**](docs/EntitlementsPartnerV2Api.md#createpartnerv2entitlement) | **Post** /partner/v2/entitlements | Create an Entitlement
*EntitlementsPartnerV2Api* | [**GetPartnerV2Entitlement**](docs/EntitlementsPartnerV2Api.md#getpartnerv2entitlement) | **Get** /partner/v2/entitlements/{id} | Read an Entitlement
*EntitlementsPartnerV2Api* | [**ListPartnerV2Entitlements**](docs/EntitlementsPartnerV2Api.md#listpartnerv2entitlements) | **Get** /partner/v2/entitlements | List of Entitlements
*EntityV1Api* | [**CreateTags**](docs/EntityV1Api.md#createtags) | **Post** /catalog/v1/entity/tags | Bulk Create Tags
*EntityV1Api* | [**DeleteTag**](docs/EntityV1Api.md#deletetag) | **Delete** /catalog/v1/entity/type/{typeName}/name/{qualifiedName}/tags/{tagName} | Delete a Tag for an Entity
*EntityV1Api* | [**GetByUniqueAttributes**](docs/EntityV1Api.md#getbyuniqueattributes) | **Get** /catalog/v1/entity/type/{typeName}/name/{qualifiedName} | Read an Entity
*EntityV1Api* | [**GetTags**](docs/EntityV1Api.md#gettags) | **Get** /catalog/v1/entity/type/{typeName}/name/{qualifiedName}/tags | Read Tags for an Entity
*EntityV1Api* | [**UpdateTags**](docs/EntityV1Api.md#updatetags) | **Put** /catalog/v1/entity/tags | Bulk Update Tags
*EnvironmentsOrgV2Api* | [**CreateOrgV2Environment**](docs/EnvironmentsOrgV2Api.md#createorgv2environment) | **Post** /org/v2/environments | Create an Environment
*EnvironmentsOrgV2Api* | [**DeleteOrgV2Environment**](docs/EnvironmentsOrgV2Api.md#deleteorgv2environment) | **Delete** /org/v2/environments/{id} | Delete an Environment
*EnvironmentsOrgV2Api* | [**GetOrgV2Environment**](docs/EnvironmentsOrgV2Api.md#getorgv2environment) | **Get** /org/v2/environments/{id} | Read an Environment
*EnvironmentsOrgV2Api* | [**ListOrgV2Environments**](docs/EnvironmentsOrgV2Api.md#listorgv2environments) | **Get** /org/v2/environments | List of Environments
*EnvironmentsOrgV2Api* | [**UpdateOrgV2Environment**](docs/EnvironmentsOrgV2Api.md#updateorgv2environment) | **Patch** /org/v2/environments/{id} | Update an Environment
*IdentityPoolsIamV2Api* | [**CreateIamV2IdentityPool**](docs/IdentityPoolsIamV2Api.md#createiamv2identitypool) | **Post** /iam/v2/identity-providers/{provider_id}/identity-pools | Create an Identity Pool
*IdentityPoolsIamV2Api* | [**DeleteIamV2IdentityPool**](docs/IdentityPoolsIamV2Api.md#deleteiamv2identitypool) | **Delete** /iam/v2/identity-providers/{provider_id}/identity-pools/{id} | Delete an Identity Pool
*IdentityPoolsIamV2Api* | [**GetIamV2IdentityPool**](docs/IdentityPoolsIamV2Api.md#getiamv2identitypool) | **Get** /iam/v2/identity-providers/{provider_id}/identity-pools/{id} | Read an Identity Pool
*IdentityPoolsIamV2Api* | [**ListIamV2IdentityPools**](docs/IdentityPoolsIamV2Api.md#listiamv2identitypools) | **Get** /iam/v2/identity-providers/{provider_id}/identity-pools | List of Identity Pools
*IdentityPoolsIamV2Api* | [**UpdateIamV2IdentityPool**](docs/IdentityPoolsIamV2Api.md#updateiamv2identitypool) | **Patch** /iam/v2/identity-providers/{provider_id}/identity-pools/{id} | Update an Identity Pool
*IdentityProvidersIamV2Api* | [**CreateIamV2IdentityProvider**](docs/IdentityProvidersIamV2Api.md#createiamv2identityprovider) | **Post** /iam/v2/identity-providers | Create an Identity Provider
*IdentityProvidersIamV2Api* | [**DeleteIamV2IdentityProvider**](docs/IdentityProvidersIamV2Api.md#deleteiamv2identityprovider) | **Delete** /iam/v2/identity-providers/{id} | Delete an Identity Provider
*IdentityProvidersIamV2Api* | [**GetIamV2IdentityProvider**](docs/IdentityProvidersIamV2Api.md#getiamv2identityprovider) | **Get** /iam/v2/identity-providers/{id} | Read an Identity Provider
*IdentityProvidersIamV2Api* | [**ListIamV2IdentityProviders**](docs/IdentityProvidersIamV2Api.md#listiamv2identityproviders) | **Get** /iam/v2/identity-providers | List of Identity Providers
*IdentityProvidersIamV2Api* | [**UpdateIamV2IdentityProvider**](docs/IdentityProvidersIamV2Api.md#updateiamv2identityprovider) | **Patch** /iam/v2/identity-providers/{id} | Update an Identity Provider
*IntegrationsNotificationsV1Api* | [**CreateNotificationsV1Integration**](docs/IntegrationsNotificationsV1Api.md#createnotificationsv1integration) | **Post** /notifications/v1/integrations | Create an Integration
*IntegrationsNotificationsV1Api* | [**DeleteNotificationsV1Integration**](docs/IntegrationsNotificationsV1Api.md#deletenotificationsv1integration) | **Delete** /notifications/v1/integrations/{id} | Delete an Integration
*IntegrationsNotificationsV1Api* | [**GetNotificationsV1Integration**](docs/IntegrationsNotificationsV1Api.md#getnotificationsv1integration) | **Get** /notifications/v1/integrations/{id} | Read an Integration
*IntegrationsNotificationsV1Api* | [**ListNotificationsV1Integrations**](docs/IntegrationsNotificationsV1Api.md#listnotificationsv1integrations) | **Get** /notifications/v1/integrations | List of Integrations
*IntegrationsNotificationsV1Api* | [**TestNotificationsV1Integration**](docs/IntegrationsNotificationsV1Api.md#testnotificationsv1integration) | **Post** /notifications/v1/integrations:test | Test a Webhook, Slack or Microsoft Teams integration
*IntegrationsNotificationsV1Api* | [**UpdateNotificationsV1Integration**](docs/IntegrationsNotificationsV1Api.md#updatenotificationsv1integration) | **Patch** /notifications/v1/integrations/{id} | Update an Integration
*InvitationsIamV2Api* | [**CreateIamV2Invitation**](docs/InvitationsIamV2Api.md#createiamv2invitation) | **Post** /iam/v2/invitations | Create an Invitation
*InvitationsIamV2Api* | [**DeleteIamV2Invitation**](docs/InvitationsIamV2Api.md#deleteiamv2invitation) | **Delete** /iam/v2/invitations/{id} | Delete an Invitation
*InvitationsIamV2Api* | [**GetIamV2Invitation**](docs/InvitationsIamV2Api.md#getiamv2invitation) | **Get** /iam/v2/invitations/{id} | Read an Invitation
*InvitationsIamV2Api* | [**ListIamV2Invitations**](docs/InvitationsIamV2Api.md#listiamv2invitations) | **Get** /iam/v2/invitations | List of Invitations
*JwksIamV2Api* | [**RefreshIamV2JsonWebKeySet**](docs/JwksIamV2Api.md#refreshiamv2jsonwebkeyset) | **Patch** /iam/v2/identity-providers/{provider_id}/jwks | Refresh a provider&#39;s JWKS
*KeysByokV1Api* | [**CreateByokV1Key**](docs/KeysByokV1Api.md#createbyokv1key) | **Post** /byok/v1/keys | Create a Key
*KeysByokV1Api* | [**DeleteByokV1Key**](docs/KeysByokV1Api.md#deletebyokv1key) | **Delete** /byok/v1/keys/{id} | Delete a Key
*KeysByokV1Api* | [**GetByokV1Key**](docs/KeysByokV1Api.md#getbyokv1key) | **Get** /byok/v1/keys/{id} | Read a Key
*KeysByokV1Api* | [**ListByokV1Keys**](docs/KeysByokV1Api.md#listbyokv1keys) | **Get** /byok/v1/keys | List of Keys
*LifecycleV1Api* | [**PauseConnectv1Connector**](docs/LifecycleV1Api.md#pauseconnectv1connector) | **Put** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors/{connector_name}/pause | Pause a Connector
*LifecycleV1Api* | [**ResumeConnectv1Connector**](docs/LifecycleV1Api.md#resumeconnectv1connector) | **Put** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors/{connector_name}/resume | Resume a Connector
*ModesV1Api* | [**DeleteSubjectMode**](docs/ModesV1Api.md#deletesubjectmode) | **Delete** /mode/{subject} | Delete subject mode
*ModesV1Api* | [**GetMode**](docs/ModesV1Api.md#getmode) | **Get** /mode/{subject} | Get subject mode
*ModesV1Api* | [**GetTopLevelMode**](docs/ModesV1Api.md#gettoplevelmode) | **Get** /mode | Get global mode
*ModesV1Api* | [**UpdateMode**](docs/ModesV1Api.md#updatemode) | **Put** /mode/{subject} | Update subject mode
*ModesV1Api* | [**UpdateTopLevelMode**](docs/ModesV1Api.md#updatetoplevelmode) | **Put** /mode | Update global mode
*NetworkLinkEndpointsNetworkingV1Api* | [**CreateNetworkingV1NetworkLinkEndpoint**](docs/NetworkLinkEndpointsNetworkingV1Api.md#createnetworkingv1networklinkendpoint) | **Post** /networking/v1/network-link-endpoints | Create a Network Link Endpoint
*NetworkLinkEndpointsNetworkingV1Api* | [**DeleteNetworkingV1NetworkLinkEndpoint**](docs/NetworkLinkEndpointsNetworkingV1Api.md#deletenetworkingv1networklinkendpoint) | **Delete** /networking/v1/network-link-endpoints/{id} | Delete a Network Link Endpoint
*NetworkLinkEndpointsNetworkingV1Api* | [**GetNetworkingV1NetworkLinkEndpoint**](docs/NetworkLinkEndpointsNetworkingV1Api.md#getnetworkingv1networklinkendpoint) | **Get** /networking/v1/network-link-endpoints/{id} | Read a Network Link Endpoint
*NetworkLinkEndpointsNetworkingV1Api* | [**ListNetworkingV1NetworkLinkEndpoints**](docs/NetworkLinkEndpointsNetworkingV1Api.md#listnetworkingv1networklinkendpoints) | **Get** /networking/v1/network-link-endpoints | List of Network Link Endpoints
*NetworkLinkEndpointsNetworkingV1Api* | [**UpdateNetworkingV1NetworkLinkEndpoint**](docs/NetworkLinkEndpointsNetworkingV1Api.md#updatenetworkingv1networklinkendpoint) | **Patch** /networking/v1/network-link-endpoints/{id} | Update a Network Link Endpoint
*NetworkLinkServiceAssociationsNetworkingV1Api* | [**GetNetworkingV1NetworkLinkServiceAssociation**](docs/NetworkLinkServiceAssociationsNetworkingV1Api.md#getnetworkingv1networklinkserviceassociation) | **Get** /networking/v1/network-link-service-associations/{id} | Read a Network Link Service Association
*NetworkLinkServiceAssociationsNetworkingV1Api* | [**ListNetworkingV1NetworkLinkServiceAssociations**](docs/NetworkLinkServiceAssociationsNetworkingV1Api.md#listnetworkingv1networklinkserviceassociations) | **Get** /networking/v1/network-link-service-associations | List of Network Link Service Associations
*NetworkLinkServicesNetworkingV1Api* | [**CreateNetworkingV1NetworkLinkService**](docs/NetworkLinkServicesNetworkingV1Api.md#createnetworkingv1networklinkservice) | **Post** /networking/v1/network-link-services | Create a Network Link Service
*NetworkLinkServicesNetworkingV1Api* | [**DeleteNetworkingV1NetworkLinkService**](docs/NetworkLinkServicesNetworkingV1Api.md#deletenetworkingv1networklinkservice) | **Delete** /networking/v1/network-link-services/{id} | Delete a Network Link Service
*NetworkLinkServicesNetworkingV1Api* | [**GetNetworkingV1NetworkLinkService**](docs/NetworkLinkServicesNetworkingV1Api.md#getnetworkingv1networklinkservice) | **Get** /networking/v1/network-link-services/{id} | Read a Network Link Service
*NetworkLinkServicesNetworkingV1Api* | [**ListNetworkingV1NetworkLinkServices**](docs/NetworkLinkServicesNetworkingV1Api.md#listnetworkingv1networklinkservices) | **Get** /networking/v1/network-link-services | List of Network Link Services
*NetworkLinkServicesNetworkingV1Api* | [**UpdateNetworkingV1NetworkLinkService**](docs/NetworkLinkServicesNetworkingV1Api.md#updatenetworkingv1networklinkservice) | **Patch** /networking/v1/network-link-services/{id} | Update a Network Link Service
*NetworksNetworkingV1Api* | [**CreateNetworkingV1Network**](docs/NetworksNetworkingV1Api.md#createnetworkingv1network) | **Post** /networking/v1/networks | Create a Network
*NetworksNetworkingV1Api* | [**DeleteNetworkingV1Network**](docs/NetworksNetworkingV1Api.md#deletenetworkingv1network) | **Delete** /networking/v1/networks/{id} | Delete a Network
*NetworksNetworkingV1Api* | [**GetNetworkingV1Network**](docs/NetworksNetworkingV1Api.md#getnetworkingv1network) | **Get** /networking/v1/networks/{id} | Read a Network
*NetworksNetworkingV1Api* | [**ListNetworkingV1Networks**](docs/NetworksNetworkingV1Api.md#listnetworkingv1networks) | **Get** /networking/v1/networks | List of Networks
*NetworksNetworkingV1Api* | [**UpdateNetworkingV1Network**](docs/NetworksNetworkingV1Api.md#updatenetworkingv1network) | **Patch** /networking/v1/networks/{id} | Update a Network
*NotificationTypesNotificationsV1Api* | [**GetNotificationsV1NotificationType**](docs/NotificationTypesNotificationsV1Api.md#getnotificationsv1notificationtype) | **Get** /notifications/v1/notification-types/{id} | Read a Notification Type
*NotificationTypesNotificationsV1Api* | [**ListNotificationsV1NotificationTypes**](docs/NotificationTypesNotificationsV1Api.md#listnotificationsv1notificationtypes) | **Get** /notifications/v1/notification-types | List of Notification Types
*OAuthTokensStsV1Api* | [**ExchangeStsV1OauthToken**](docs/OAuthTokensStsV1Api.md#exchangestsv1oauthtoken) | **Post** /sts/v1/oauth2/token | Exchange an OAuth Token
*OrganizationsOrgV2Api* | [**GetOrgV2Organization**](docs/OrganizationsOrgV2Api.md#getorgv2organization) | **Get** /org/v2/organizations/{id} | Read an Organization
*OrganizationsOrgV2Api* | [**ListOrgV2Organizations**](docs/OrganizationsOrgV2Api.md#listorgv2organizations) | **Get** /org/v2/organizations | List of Organizations
*OrganizationsOrgV2Api* | [**UpdateOrgV2Organization**](docs/OrganizationsOrgV2Api.md#updateorgv2organization) | **Patch** /org/v2/organizations/{id} | Update an Organization
*OrganizationsPartnerV2Api* | [**GetPartnerV2Organization**](docs/OrganizationsPartnerV2Api.md#getpartnerv2organization) | **Get** /partner/v2/organizations/{id} | Read an Organization
*OrganizationsPartnerV2Api* | [**ListPartnerV2Organizations**](docs/OrganizationsPartnerV2Api.md#listpartnerv2organizations) | **Get** /partner/v2/organizations | List of Organizations
*PartitionV3Api* | [**GetKafkaPartition**](docs/PartitionV3Api.md#getkafkapartition) | **Get** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/partitions/{partition_id} | Get Partition
*PartitionV3Api* | [**ListKafkaPartitions**](docs/PartitionV3Api.md#listkafkapartitions) | **Get** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/partitions | List Partitions
*PeeringsNetworkingV1Api* | [**CreateNetworkingV1Peering**](docs/PeeringsNetworkingV1Api.md#createnetworkingv1peering) | **Post** /networking/v1/peerings | Create a Peering
*PeeringsNetworkingV1Api* | [**DeleteNetworkingV1Peering**](docs/PeeringsNetworkingV1Api.md#deletenetworkingv1peering) | **Delete** /networking/v1/peerings/{id} | Delete a Peering
*PeeringsNetworkingV1Api* | [**GetNetworkingV1Peering**](docs/PeeringsNetworkingV1Api.md#getnetworkingv1peering) | **Get** /networking/v1/peerings/{id} | Read a Peering
*PeeringsNetworkingV1Api* | [**ListNetworkingV1Peerings**](docs/PeeringsNetworkingV1Api.md#listnetworkingv1peerings) | **Get** /networking/v1/peerings | List of Peerings
*PeeringsNetworkingV1Api* | [**UpdateNetworkingV1Peering**](docs/PeeringsNetworkingV1Api.md#updatenetworkingv1peering) | **Patch** /networking/v1/peerings/{id} | Update a Peering
*PipelinesSdV1Api* | [**CreateSdV1Pipeline**](docs/PipelinesSdV1Api.md#createsdv1pipeline) | **Post** /sd/v1/pipelines | Create a Pipeline
*PipelinesSdV1Api* | [**DeleteSdV1Pipeline**](docs/PipelinesSdV1Api.md#deletesdv1pipeline) | **Delete** /sd/v1/pipelines/{id} | Delete a Pipeline
*PipelinesSdV1Api* | [**GetSdV1Pipeline**](docs/PipelinesSdV1Api.md#getsdv1pipeline) | **Get** /sd/v1/pipelines/{id} | Read a Pipeline
*PipelinesSdV1Api* | [**ListSdV1Pipelines**](docs/PipelinesSdV1Api.md#listsdv1pipelines) | **Get** /sd/v1/pipelines | List of Pipelines
*PipelinesSdV1Api* | [**UpdateSdV1Pipeline**](docs/PipelinesSdV1Api.md#updatesdv1pipeline) | **Patch** /sd/v1/pipelines/{id} | Update a Pipeline
*PluginsV1Api* | [**ListConnectv1ConnectorPlugins**](docs/PluginsV1Api.md#listconnectv1connectorplugins) | **Get** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connector-plugins | List of Connector Plugins
*PluginsV1Api* | [**ValidateConnectv1ConnectorPlugin**](docs/PluginsV1Api.md#validateconnectv1connectorplugin) | **Put** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connector-plugins/{plugin_name}/config/validate | Validate a Connector Plugin
*PrivateLinkAccessesNetworkingV1Api* | [**CreateNetworkingV1PrivateLinkAccess**](docs/PrivateLinkAccessesNetworkingV1Api.md#createnetworkingv1privatelinkaccess) | **Post** /networking/v1/private-link-accesses | Create a Private Link Access
*PrivateLinkAccessesNetworkingV1Api* | [**DeleteNetworkingV1PrivateLinkAccess**](docs/PrivateLinkAccessesNetworkingV1Api.md#deletenetworkingv1privatelinkaccess) | **Delete** /networking/v1/private-link-accesses/{id} | Delete a Private Link Access
*PrivateLinkAccessesNetworkingV1Api* | [**GetNetworkingV1PrivateLinkAccess**](docs/PrivateLinkAccessesNetworkingV1Api.md#getnetworkingv1privatelinkaccess) | **Get** /networking/v1/private-link-accesses/{id} | Read a Private Link Access
*PrivateLinkAccessesNetworkingV1Api* | [**ListNetworkingV1PrivateLinkAccesses**](docs/PrivateLinkAccessesNetworkingV1Api.md#listnetworkingv1privatelinkaccesses) | **Get** /networking/v1/private-link-accesses | List of Private Link Accesses
*PrivateLinkAccessesNetworkingV1Api* | [**UpdateNetworkingV1PrivateLinkAccess**](docs/PrivateLinkAccessesNetworkingV1Api.md#updatenetworkingv1privatelinkaccess) | **Patch** /networking/v1/private-link-accesses/{id} | Update a Private Link Access
*RecordsV3Api* | [**ProduceRecord**](docs/RecordsV3Api.md#producerecord) | **Post** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/records | Produce Records
*RegionsSrcmV2Api* | [**GetSrcmV2Region**](docs/RegionsSrcmV2Api.md#getsrcmv2region) | **Get** /srcm/v2/regions/{id} | Read a Region
*RegionsSrcmV2Api* | [**ListSrcmV2Regions**](docs/RegionsSrcmV2Api.md#listsrcmv2regions) | **Get** /srcm/v2/regions | List of Regions
*RoleBindingsIamV2Api* | [**CreateIamV2RoleBinding**](docs/RoleBindingsIamV2Api.md#createiamv2rolebinding) | **Post** /iam/v2/role-bindings | Create a Role Binding
*RoleBindingsIamV2Api* | [**DeleteIamV2RoleBinding**](docs/RoleBindingsIamV2Api.md#deleteiamv2rolebinding) | **Delete** /iam/v2/role-bindings/{id} | Delete a Role Binding
*RoleBindingsIamV2Api* | [**GetIamV2RoleBinding**](docs/RoleBindingsIamV2Api.md#getiamv2rolebinding) | **Get** /iam/v2/role-bindings/{id} | Read a Role Binding
*RoleBindingsIamV2Api* | [**ListIamV2RoleBindings**](docs/RoleBindingsIamV2Api.md#listiamv2rolebindings) | **Get** /iam/v2/role-bindings | List of Role Bindings
*SchemasV1Api* | [**GetSchema**](docs/SchemasV1Api.md#getschema) | **Get** /schemas/ids/{id} | Get schema string by ID
*SchemasV1Api* | [**GetSchemaOnly**](docs/SchemasV1Api.md#getschemaonly) | **Get** /schemas/ids/{id}/schema | Get schema by ID
*SchemasV1Api* | [**GetSchemaTypes**](docs/SchemasV1Api.md#getschematypes) | **Get** /schemas/types | List supported schema types
*SchemasV1Api* | [**GetSchemas**](docs/SchemasV1Api.md#getschemas) | **Get** /schemas | List schemas
*SchemasV1Api* | [**GetSubjects**](docs/SchemasV1Api.md#getsubjects) | **Get** /schemas/ids/{id}/subjects | List subjects associated to schema ID
*SchemasV1Api* | [**GetVersions**](docs/SchemasV1Api.md#getversions) | **Get** /schemas/ids/{id}/versions | List subject-versions associated to schema ID
*ScopesServiceQuotaV1Api* | [**GetServiceQuotaV1Scope**](docs/ScopesServiceQuotaV1Api.md#getservicequotav1scope) | **Get** /service-quota/v1/scopes/{id} | Read a Scope
*ScopesServiceQuotaV1Api* | [**ListServiceQuotaV1Scopes**](docs/ScopesServiceQuotaV1Api.md#listservicequotav1scopes) | **Get** /service-quota/v1/scopes | List of Scopes
*SearchV1Api* | [**SearchUsingAttribute**](docs/SearchV1Api.md#searchusingattribute) | **Get** /catalog/v1/search/attribute | Search by Attribute
*SearchV1Api* | [**SearchUsingBasic**](docs/SearchV1Api.md#searchusingbasic) | **Get** /catalog/v1/search/basic | Search by Fulltext Query
*ServiceAccountsIamV2Api* | [**CreateIamV2ServiceAccount**](docs/ServiceAccountsIamV2Api.md#createiamv2serviceaccount) | **Post** /iam/v2/service-accounts | Create a Service Account
*ServiceAccountsIamV2Api* | [**DeleteIamV2ServiceAccount**](docs/ServiceAccountsIamV2Api.md#deleteiamv2serviceaccount) | **Delete** /iam/v2/service-accounts/{id} | Delete a Service Account
*ServiceAccountsIamV2Api* | [**GetIamV2ServiceAccount**](docs/ServiceAccountsIamV2Api.md#getiamv2serviceaccount) | **Get** /iam/v2/service-accounts/{id} | Read a Service Account
*ServiceAccountsIamV2Api* | [**ListIamV2ServiceAccounts**](docs/ServiceAccountsIamV2Api.md#listiamv2serviceaccounts) | **Get** /iam/v2/service-accounts | List of Service Accounts
*ServiceAccountsIamV2Api* | [**UpdateIamV2ServiceAccount**](docs/ServiceAccountsIamV2Api.md#updateiamv2serviceaccount) | **Patch** /iam/v2/service-accounts/{id} | Update a Service Account
*SignupPartnerV2Api* | [**ActivateSignup**](docs/SignupPartnerV2Api.md#activatesignup) | **Post** /partner/v2/signup/activate | Activate an Incomplete Signup
*SignupPartnerV2Api* | [**Signup**](docs/SignupPartnerV2Api.md#signup) | **Post** /partner/v2/signup | Signup an Organization on behalf of a Customer
*SignupPartnerV2Api* | [**SignupPartnerV2Link**](docs/SignupPartnerV2Api.md#signuppartnerv2link) | **Post** /partner/v2/signup/link | Signup a Customer by Linking to an Existing Organization
*StatusV1Api* | [**ListConnectv1ConnectorTasks**](docs/StatusV1Api.md#listconnectv1connectortasks) | **Get** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors/{connector_name}/tasks | List of Connector Tasks
*StatusV1Api* | [**ReadConnectv1ConnectorStatus**](docs/StatusV1Api.md#readconnectv1connectorstatus) | **Get** /connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors/{connector_name}/status | Read a Connector Status
*SubjectsV1Api* | [**DeleteSchemaVersion**](docs/SubjectsV1Api.md#deleteschemaversion) | **Delete** /subjects/{subject}/versions/{version} | Delete schema version
*SubjectsV1Api* | [**DeleteSubject**](docs/SubjectsV1Api.md#deletesubject) | **Delete** /subjects/{subject} | Delete subject
*SubjectsV1Api* | [**GetReferencedBy**](docs/SubjectsV1Api.md#getreferencedby) | **Get** /subjects/{subject}/versions/{version}/referencedby | List schemas referencing a schema
*SubjectsV1Api* | [**GetSchemaByVersion**](docs/SubjectsV1Api.md#getschemabyversion) | **Get** /subjects/{subject}/versions/{version} | Get schema by version
*SubjectsV1Api* | [**GetSchemaOnly1**](docs/SubjectsV1Api.md#getschemaonly1) | **Get** /subjects/{subject}/versions/{version}/schema | Get schema string by version
*SubjectsV1Api* | [**List**](docs/SubjectsV1Api.md#list) | **Get** /subjects | List subjects
*SubjectsV1Api* | [**ListVersions**](docs/SubjectsV1Api.md#listversions) | **Get** /subjects/{subject}/versions | List versions under subject
*SubjectsV1Api* | [**LookUpSchemaUnderSubject**](docs/SubjectsV1Api.md#lookupschemaundersubject) | **Post** /subjects/{subject} | Lookup schema under subject
*SubjectsV1Api* | [**Register**](docs/SubjectsV1Api.md#register) | **Post** /subjects/{subject}/versions | Register schema under a subject
*SubscriptionsNotificationsV1Api* | [**CreateNotificationsV1Subscription**](docs/SubscriptionsNotificationsV1Api.md#createnotificationsv1subscription) | **Post** /notifications/v1/subscriptions | Create a Subscription
*SubscriptionsNotificationsV1Api* | [**DeleteNotificationsV1Subscription**](docs/SubscriptionsNotificationsV1Api.md#deletenotificationsv1subscription) | **Delete** /notifications/v1/subscriptions/{id} | Delete a Subscription
*SubscriptionsNotificationsV1Api* | [**GetNotificationsV1Subscription**](docs/SubscriptionsNotificationsV1Api.md#getnotificationsv1subscription) | **Get** /notifications/v1/subscriptions/{id} | Read a Subscription
*SubscriptionsNotificationsV1Api* | [**ListNotificationsV1Subscriptions**](docs/SubscriptionsNotificationsV1Api.md#listnotificationsv1subscriptions) | **Get** /notifications/v1/subscriptions | List of Subscriptions
*SubscriptionsNotificationsV1Api* | [**UpdateNotificationsV1Subscription**](docs/SubscriptionsNotificationsV1Api.md#updatenotificationsv1subscription) | **Patch** /notifications/v1/subscriptions/{id} | Update a Subscription
*TopicV3Api* | [**CreateKafkaTopic**](docs/TopicV3Api.md#createkafkatopic) | **Post** /kafka/v3/clusters/{cluster_id}/topics | Create Topic
*TopicV3Api* | [**DeleteKafkaTopic**](docs/TopicV3Api.md#deletekafkatopic) | **Delete** /kafka/v3/clusters/{cluster_id}/topics/{topic_name} | Delete Topic
*TopicV3Api* | [**GetKafkaTopic**](docs/TopicV3Api.md#getkafkatopic) | **Get** /kafka/v3/clusters/{cluster_id}/topics/{topic_name} | Get Topic
*TopicV3Api* | [**ListKafkaTopics**](docs/TopicV3Api.md#listkafkatopics) | **Get** /kafka/v3/clusters/{cluster_id}/topics | List Topics
*TopicV3Api* | [**UpdatePartitionCountKafkaTopic**](docs/TopicV3Api.md#updatepartitioncountkafkatopic) | **Patch** /kafka/v3/clusters/{cluster_id}/topics/{topic_name} | Update partition count
*TransitGatewayAttachmentsNetworkingV1Api* | [**CreateNetworkingV1TransitGatewayAttachment**](docs/TransitGatewayAttachmentsNetworkingV1Api.md#createnetworkingv1transitgatewayattachment) | **Post** /networking/v1/transit-gateway-attachments | Create a Transit Gateway Attachment
*TransitGatewayAttachmentsNetworkingV1Api* | [**DeleteNetworkingV1TransitGatewayAttachment**](docs/TransitGatewayAttachmentsNetworkingV1Api.md#deletenetworkingv1transitgatewayattachment) | **Delete** /networking/v1/transit-gateway-attachments/{id} | Delete a Transit Gateway Attachment
*TransitGatewayAttachmentsNetworkingV1Api* | [**GetNetworkingV1TransitGatewayAttachment**](docs/TransitGatewayAttachmentsNetworkingV1Api.md#getnetworkingv1transitgatewayattachment) | **Get** /networking/v1/transit-gateway-attachments/{id} | Read a Transit Gateway Attachment
*TransitGatewayAttachmentsNetworkingV1Api* | [**ListNetworkingV1TransitGatewayAttachments**](docs/TransitGatewayAttachmentsNetworkingV1Api.md#listnetworkingv1transitgatewayattachments) | **Get** /networking/v1/transit-gateway-attachments | List of Transit Gateway Attachments
*TransitGatewayAttachmentsNetworkingV1Api* | [**UpdateNetworkingV1TransitGatewayAttachment**](docs/TransitGatewayAttachmentsNetworkingV1Api.md#updatenetworkingv1transitgatewayattachment) | **Patch** /networking/v1/transit-gateway-attachments/{id} | Update a Transit Gateway Attachment
*TypesV1Api* | [**CreateTagDefs**](docs/TypesV1Api.md#createtagdefs) | **Post** /catalog/v1/types/tagdefs | Bulk Create Tag Definitions
*TypesV1Api* | [**DeleteTagDef**](docs/TypesV1Api.md#deletetagdef) | **Delete** /catalog/v1/types/tagdefs/{tagName} | Delete Tag Definition
*TypesV1Api* | [**GetAllTagDefs**](docs/TypesV1Api.md#getalltagdefs) | **Get** /catalog/v1/types/tagdefs | Bulk Read Tag Definitions
*TypesV1Api* | [**GetTagDefByName**](docs/TypesV1Api.md#gettagdefbyname) | **Get** /catalog/v1/types/tagdefs/{tagName} | Read Tag Definition
*TypesV1Api* | [**UpdateTagDefs**](docs/TypesV1Api.md#updatetagdefs) | **Put** /catalog/v1/types/tagdefs | Bulk Update Tag Definitions
*UsersIamV2Api* | [**DeleteIamV2User**](docs/UsersIamV2Api.md#deleteiamv2user) | **Delete** /iam/v2/users/{id} | Delete a User
*UsersIamV2Api* | [**GetIamV2User**](docs/UsersIamV2Api.md#getiamv2user) | **Get** /iam/v2/users/{id} | Read a User
*UsersIamV2Api* | [**ListIamV2Users**](docs/UsersIamV2Api.md#listiamv2users) | **Get** /iam/v2/users | List of Users
*UsersIamV2Api* | [**UpdateIamV2User**](docs/UsersIamV2Api.md#updateiamv2user) | **Patch** /iam/v2/users/{id} | Update a User


## Documentation For Models

 - [AbstractConfigData](docs/AbstractConfigData.md)
 - [AbstractConfigDataAllOf](docs/AbstractConfigDataAllOf.md)
 - [AclData](docs/AclData.md)
 - [AclDataAllOf](docs/AclDataAllOf.md)
 - [AclDataList](docs/AclDataList.md)
 - [AclDataListAllOf](docs/AclDataListAllOf.md)
 - [AclResourceType](docs/AclResourceType.md)
 - [ActivatePartnerSignupRequest](docs/ActivatePartnerSignupRequest.md)
 - [AlterBrokerReplicaExclusionData](docs/AlterBrokerReplicaExclusionData.md)
 - [AlterBrokerReplicaExclusionDataAllOf](docs/AlterBrokerReplicaExclusionDataAllOf.md)
 - [AlterBrokerReplicaExclusionDataList](docs/AlterBrokerReplicaExclusionDataList.md)
 - [AlterBrokerReplicaExclusionDataListAllOf](docs/AlterBrokerReplicaExclusionDataListAllOf.md)
 - [AlterConfigBatchRequestData](docs/AlterConfigBatchRequestData.md)
 - [AlterConfigBatchRequestDataDataInner](docs/AlterConfigBatchRequestDataDataInner.md)
 - [AlterMirrorStatusResponseData](docs/AlterMirrorStatusResponseData.md)
 - [AlterMirrorStatusResponseDataAllOf](docs/AlterMirrorStatusResponseDataAllOf.md)
 - [AlterMirrorStatusResponseDataList](docs/AlterMirrorStatusResponseDataList.md)
 - [AlterMirrorStatusResponseDataListAllOf](docs/AlterMirrorStatusResponseDataListAllOf.md)
 - [AlterMirrorsRequestData](docs/AlterMirrorsRequestData.md)
 - [AnyUnevenLoadData](docs/AnyUnevenLoadData.md)
 - [AnyUnevenLoadDataAllOf](docs/AnyUnevenLoadDataAllOf.md)
 - [AttributeDef](docs/AttributeDef.md)
 - [AzureSSOConfig](docs/AzureSSOConfig.md)
 - [BalancerStatusData](docs/BalancerStatusData.md)
 - [BalancerStatusDataAllOf](docs/BalancerStatusDataAllOf.md)
 - [BillingV1Cost](docs/BillingV1Cost.md)
 - [BillingV1CostList](docs/BillingV1CostList.md)
 - [BillingV1CostListDataInner](docs/BillingV1CostListDataInner.md)
 - [BillingV1CostListMetadata](docs/BillingV1CostListMetadata.md)
 - [BillingV1CostListMetadataAllOf](docs/BillingV1CostListMetadataAllOf.md)
 - [BillingV1CostResource](docs/BillingV1CostResource.md)
 - [BillingV1Environment](docs/BillingV1Environment.md)
 - [BillingV1Resource](docs/BillingV1Resource.md)
 - [BillingV1ResourceEnvironment](docs/BillingV1ResourceEnvironment.md)
 - [BrokerConfigData](docs/BrokerConfigData.md)
 - [BrokerConfigDataAllOf](docs/BrokerConfigDataAllOf.md)
 - [BrokerConfigDataList](docs/BrokerConfigDataList.md)
 - [BrokerConfigDataListAllOf](docs/BrokerConfigDataListAllOf.md)
 - [BrokerData](docs/BrokerData.md)
 - [BrokerDataAllOf](docs/BrokerDataAllOf.md)
 - [BrokerDataList](docs/BrokerDataList.md)
 - [BrokerDataListAllOf](docs/BrokerDataListAllOf.md)
 - [BrokerRemovalData](docs/BrokerRemovalData.md)
 - [BrokerRemovalDataAllOf](docs/BrokerRemovalDataAllOf.md)
 - [BrokerRemovalDataList](docs/BrokerRemovalDataList.md)
 - [BrokerRemovalDataListAllOf](docs/BrokerRemovalDataListAllOf.md)
 - [BrokerReplicaExclusionBatchRequestData](docs/BrokerReplicaExclusionBatchRequestData.md)
 - [BrokerReplicaExclusionData](docs/BrokerReplicaExclusionData.md)
 - [BrokerReplicaExclusionDataAllOf](docs/BrokerReplicaExclusionDataAllOf.md)
 - [BrokerReplicaExclusionDataList](docs/BrokerReplicaExclusionDataList.md)
 - [BrokerReplicaExclusionDataListAllOf](docs/BrokerReplicaExclusionDataListAllOf.md)
 - [BrokerReplicaExclusionRequestData](docs/BrokerReplicaExclusionRequestData.md)
 - [BrokerTaskData](docs/BrokerTaskData.md)
 - [BrokerTaskDataAllOf](docs/BrokerTaskDataAllOf.md)
 - [BrokerTaskDataList](docs/BrokerTaskDataList.md)
 - [BrokerTaskDataListAllOf](docs/BrokerTaskDataListAllOf.md)
 - [BrokerTaskType](docs/BrokerTaskType.md)
 - [ByokV1AwsKey](docs/ByokV1AwsKey.md)
 - [ByokV1AzureKey](docs/ByokV1AzureKey.md)
 - [ByokV1Key](docs/ByokV1Key.md)
 - [ByokV1KeyKey](docs/ByokV1KeyKey.md)
 - [ByokV1KeyList](docs/ByokV1KeyList.md)
 - [ByokV1KeyListDataInner](docs/ByokV1KeyListDataInner.md)
 - [ByokV1KeyListMetadata](docs/ByokV1KeyListMetadata.md)
 - [ByokV1KeyListMetadataAllOf](docs/ByokV1KeyListMetadataAllOf.md)
 - [ByokV1KeyMetadata](docs/ByokV1KeyMetadata.md)
 - [ByokV1KeyMetadataAllOf](docs/ByokV1KeyMetadataAllOf.md)
 - [Classification](docs/Classification.md)
 - [ClusterConfig](docs/ClusterConfig.md)
 - [ClusterConfigData](docs/ClusterConfigData.md)
 - [ClusterConfigDataAllOf](docs/ClusterConfigDataAllOf.md)
 - [ClusterConfigDataList](docs/ClusterConfigDataList.md)
 - [ClusterConfigDataListAllOf](docs/ClusterConfigDataListAllOf.md)
 - [ClusterData](docs/ClusterData.md)
 - [ClusterDataAllOf](docs/ClusterDataAllOf.md)
 - [ClusterDataList](docs/ClusterDataList.md)
 - [ClusterDataListAllOf](docs/ClusterDataListAllOf.md)
 - [CmkV2Basic](docs/CmkV2Basic.md)
 - [CmkV2Cluster](docs/CmkV2Cluster.md)
 - [CmkV2ClusterList](docs/CmkV2ClusterList.md)
 - [CmkV2ClusterListDataInner](docs/CmkV2ClusterListDataInner.md)
 - [CmkV2ClusterListDataInnerAllOf](docs/CmkV2ClusterListDataInnerAllOf.md)
 - [CmkV2ClusterListMetadata](docs/CmkV2ClusterListMetadata.md)
 - [CmkV2ClusterListMetadataAllOf](docs/CmkV2ClusterListMetadataAllOf.md)
 - [CmkV2ClusterMetadata](docs/CmkV2ClusterMetadata.md)
 - [CmkV2ClusterMetadataAllOf](docs/CmkV2ClusterMetadataAllOf.md)
 - [CmkV2ClusterSpec](docs/CmkV2ClusterSpec.md)
 - [CmkV2ClusterSpecByok](docs/CmkV2ClusterSpecByok.md)
 - [CmkV2ClusterSpecConfig](docs/CmkV2ClusterSpecConfig.md)
 - [CmkV2ClusterSpecEnvironment](docs/CmkV2ClusterSpecEnvironment.md)
 - [CmkV2ClusterSpecNetwork](docs/CmkV2ClusterSpecNetwork.md)
 - [CmkV2ClusterSpecUpdate](docs/CmkV2ClusterSpecUpdate.md)
 - [CmkV2ClusterStatus](docs/CmkV2ClusterStatus.md)
 - [CmkV2ClusterUpdate](docs/CmkV2ClusterUpdate.md)
 - [CmkV2Dedicated](docs/CmkV2Dedicated.md)
 - [CmkV2Standard](docs/CmkV2Standard.md)
 - [CompatibilityCheckResponse](docs/CompatibilityCheckResponse.md)
 - [Config](docs/Config.md)
 - [ConfigData](docs/ConfigData.md)
 - [ConfigSynonymData](docs/ConfigSynonymData.md)
 - [ConfigUpdateRequest](docs/ConfigUpdateRequest.md)
 - [ConnectV1Connector](docs/ConnectV1Connector.md)
 - [ConnectV1ConnectorConfig](docs/ConnectV1ConnectorConfig.md)
 - [ConnectV1ConnectorError](docs/ConnectV1ConnectorError.md)
 - [ConnectV1ConnectorErrorError](docs/ConnectV1ConnectorErrorError.md)
 - [ConnectV1ConnectorExpansion](docs/ConnectV1ConnectorExpansion.md)
 - [ConnectV1ConnectorExpansionId](docs/ConnectV1ConnectorExpansionId.md)
 - [ConnectV1ConnectorExpansionInfo](docs/ConnectV1ConnectorExpansionInfo.md)
 - [ConnectV1ConnectorExpansionInfoConfig](docs/ConnectV1ConnectorExpansionInfoConfig.md)
 - [ConnectV1ConnectorExpansionStatus](docs/ConnectV1ConnectorExpansionStatus.md)
 - [ConnectV1ConnectorExpansionStatusConnector](docs/ConnectV1ConnectorExpansionStatusConnector.md)
 - [ConnectV1ConnectorTasksInner](docs/ConnectV1ConnectorTasksInner.md)
 - [ConnectV1ConnectorsInner](docs/ConnectV1ConnectorsInner.md)
 - [ConnectV1ConnectorsInnerConfig](docs/ConnectV1ConnectorsInnerConfig.md)
 - [ConnectV1ConnectorsInnerId](docs/ConnectV1ConnectorsInnerId.md)
 - [ConstraintDef](docs/ConstraintDef.md)
 - [ConsumerAssignmentData](docs/ConsumerAssignmentData.md)
 - [ConsumerAssignmentDataAllOf](docs/ConsumerAssignmentDataAllOf.md)
 - [ConsumerAssignmentDataList](docs/ConsumerAssignmentDataList.md)
 - [ConsumerAssignmentDataListAllOf](docs/ConsumerAssignmentDataListAllOf.md)
 - [ConsumerData](docs/ConsumerData.md)
 - [ConsumerDataAllOf](docs/ConsumerDataAllOf.md)
 - [ConsumerDataList](docs/ConsumerDataList.md)
 - [ConsumerDataListAllOf](docs/ConsumerDataListAllOf.md)
 - [ConsumerGroupData](docs/ConsumerGroupData.md)
 - [ConsumerGroupDataAllOf](docs/ConsumerGroupDataAllOf.md)
 - [ConsumerGroupDataList](docs/ConsumerGroupDataList.md)
 - [ConsumerGroupDataListAllOf](docs/ConsumerGroupDataListAllOf.md)
 - [ConsumerGroupLagSummaryData](docs/ConsumerGroupLagSummaryData.md)
 - [ConsumerGroupLagSummaryDataAllOf](docs/ConsumerGroupLagSummaryDataAllOf.md)
 - [ConsumerLagData](docs/ConsumerLagData.md)
 - [ConsumerLagDataAllOf](docs/ConsumerLagDataAllOf.md)
 - [ConsumerLagDataList](docs/ConsumerLagDataList.md)
 - [ConsumerLagDataListAllOf](docs/ConsumerLagDataListAllOf.md)
 - [CreateAclRequestData](docs/CreateAclRequestData.md)
 - [CreateAclRequestDataList](docs/CreateAclRequestDataList.md)
 - [CreateAclRequestDataListAllOf](docs/CreateAclRequestDataListAllOf.md)
 - [CreateByokV1KeyRequest](docs/CreateByokV1KeyRequest.md)
 - [CreateCmkV2Cluster202Response](docs/CreateCmkV2Cluster202Response.md)
 - [CreateCmkV2Cluster202ResponseAllOf](docs/CreateCmkV2Cluster202ResponseAllOf.md)
 - [CreateCmkV2Cluster202ResponseAllOf1](docs/CreateCmkV2Cluster202ResponseAllOf1.md)
 - [CreateCmkV2ClusterRequest](docs/CreateCmkV2ClusterRequest.md)
 - [CreateCmkV2ClusterRequestAllOf](docs/CreateCmkV2ClusterRequestAllOf.md)
 - [CreateCmkV2ClusterRequestAllOf1](docs/CreateCmkV2ClusterRequestAllOf1.md)
 - [CreateCmkV2ClusterRequestAllOf1Spec](docs/CreateCmkV2ClusterRequestAllOf1Spec.md)
 - [CreateConnectv1Connector400Response](docs/CreateConnectv1Connector400Response.md)
 - [CreateConnectv1Connector500Response](docs/CreateConnectv1Connector500Response.md)
 - [CreateConnectv1ConnectorRequest](docs/CreateConnectv1ConnectorRequest.md)
 - [CreateConnectv1ConnectorRequestConfig](docs/CreateConnectv1ConnectorRequestConfig.md)
 - [CreateIamV2ApiKey202Response](docs/CreateIamV2ApiKey202Response.md)
 - [CreateIamV2ApiKey202ResponseAllOf](docs/CreateIamV2ApiKey202ResponseAllOf.md)
 - [CreateIamV2ApiKeyRequest](docs/CreateIamV2ApiKeyRequest.md)
 - [CreateIamV2ApiKeyRequestAllOf](docs/CreateIamV2ApiKeyRequestAllOf.md)
 - [CreateIamV2ApiKeyRequestAllOf1](docs/CreateIamV2ApiKeyRequestAllOf1.md)
 - [CreateIamV2ApiKeyRequestAllOf1Spec](docs/CreateIamV2ApiKeyRequestAllOf1Spec.md)
 - [CreateIamV2IdentityPoolRequest](docs/CreateIamV2IdentityPoolRequest.md)
 - [CreateIamV2IdentityProviderRequest](docs/CreateIamV2IdentityProviderRequest.md)
 - [CreateIamV2InvitationRequest](docs/CreateIamV2InvitationRequest.md)
 - [CreateIamV2RoleBindingRequest](docs/CreateIamV2RoleBindingRequest.md)
 - [CreateIamV2ServiceAccountRequest](docs/CreateIamV2ServiceAccountRequest.md)
 - [CreateKafkaQuotasV1ClientQuota202Response](docs/CreateKafkaQuotasV1ClientQuota202Response.md)
 - [CreateKafkaQuotasV1ClientQuota202ResponseAllOf](docs/CreateKafkaQuotasV1ClientQuota202ResponseAllOf.md)
 - [CreateKafkaQuotasV1ClientQuotaRequest](docs/CreateKafkaQuotasV1ClientQuotaRequest.md)
 - [CreateKafkaQuotasV1ClientQuotaRequestAllOf](docs/CreateKafkaQuotasV1ClientQuotaRequestAllOf.md)
 - [CreateKafkaQuotasV1ClientQuotaRequestAllOf1](docs/CreateKafkaQuotasV1ClientQuotaRequestAllOf1.md)
 - [CreateKafkaQuotasV1ClientQuotaRequestAllOf1Spec](docs/CreateKafkaQuotasV1ClientQuotaRequestAllOf1Spec.md)
 - [CreateKsqldbcmV2Cluster202Response](docs/CreateKsqldbcmV2Cluster202Response.md)
 - [CreateKsqldbcmV2Cluster202ResponseAllOf](docs/CreateKsqldbcmV2Cluster202ResponseAllOf.md)
 - [CreateKsqldbcmV2Cluster202ResponseAllOf1](docs/CreateKsqldbcmV2Cluster202ResponseAllOf1.md)
 - [CreateKsqldbcmV2ClusterRequest](docs/CreateKsqldbcmV2ClusterRequest.md)
 - [CreateKsqldbcmV2ClusterRequestAllOf](docs/CreateKsqldbcmV2ClusterRequestAllOf.md)
 - [CreateKsqldbcmV2ClusterRequestAllOf1](docs/CreateKsqldbcmV2ClusterRequestAllOf1.md)
 - [CreateKsqldbcmV2ClusterRequestAllOf1Spec](docs/CreateKsqldbcmV2ClusterRequestAllOf1Spec.md)
 - [CreateLinkRequestData](docs/CreateLinkRequestData.md)
 - [CreateMirrorTopicRequestData](docs/CreateMirrorTopicRequestData.md)
 - [CreateNetworkingV1Network202Response](docs/CreateNetworkingV1Network202Response.md)
 - [CreateNetworkingV1Network202ResponseAllOf](docs/CreateNetworkingV1Network202ResponseAllOf.md)
 - [CreateNetworkingV1Network202ResponseAllOf1](docs/CreateNetworkingV1Network202ResponseAllOf1.md)
 - [CreateNetworkingV1NetworkLinkEndpoint202Response](docs/CreateNetworkingV1NetworkLinkEndpoint202Response.md)
 - [CreateNetworkingV1NetworkLinkEndpoint202ResponseAllOf](docs/CreateNetworkingV1NetworkLinkEndpoint202ResponseAllOf.md)
 - [CreateNetworkingV1NetworkLinkEndpoint202ResponseAllOf1](docs/CreateNetworkingV1NetworkLinkEndpoint202ResponseAllOf1.md)
 - [CreateNetworkingV1NetworkLinkEndpointRequest](docs/CreateNetworkingV1NetworkLinkEndpointRequest.md)
 - [CreateNetworkingV1NetworkLinkEndpointRequestAllOf](docs/CreateNetworkingV1NetworkLinkEndpointRequestAllOf.md)
 - [CreateNetworkingV1NetworkLinkEndpointRequestAllOf1](docs/CreateNetworkingV1NetworkLinkEndpointRequestAllOf1.md)
 - [CreateNetworkingV1NetworkLinkEndpointRequestAllOf1Spec](docs/CreateNetworkingV1NetworkLinkEndpointRequestAllOf1Spec.md)
 - [CreateNetworkingV1NetworkLinkService202Response](docs/CreateNetworkingV1NetworkLinkService202Response.md)
 - [CreateNetworkingV1NetworkLinkService202ResponseAllOf](docs/CreateNetworkingV1NetworkLinkService202ResponseAllOf.md)
 - [CreateNetworkingV1NetworkLinkServiceRequest](docs/CreateNetworkingV1NetworkLinkServiceRequest.md)
 - [CreateNetworkingV1NetworkLinkServiceRequestAllOf](docs/CreateNetworkingV1NetworkLinkServiceRequestAllOf.md)
 - [CreateNetworkingV1NetworkRequest](docs/CreateNetworkingV1NetworkRequest.md)
 - [CreateNetworkingV1NetworkRequestAllOf](docs/CreateNetworkingV1NetworkRequestAllOf.md)
 - [CreateNetworkingV1NetworkRequestAllOf1](docs/CreateNetworkingV1NetworkRequestAllOf1.md)
 - [CreateNetworkingV1NetworkRequestAllOf1Spec](docs/CreateNetworkingV1NetworkRequestAllOf1Spec.md)
 - [CreateNetworkingV1Peering202Response](docs/CreateNetworkingV1Peering202Response.md)
 - [CreateNetworkingV1Peering202ResponseAllOf](docs/CreateNetworkingV1Peering202ResponseAllOf.md)
 - [CreateNetworkingV1Peering202ResponseAllOf1](docs/CreateNetworkingV1Peering202ResponseAllOf1.md)
 - [CreateNetworkingV1PeeringRequest](docs/CreateNetworkingV1PeeringRequest.md)
 - [CreateNetworkingV1PeeringRequestAllOf](docs/CreateNetworkingV1PeeringRequestAllOf.md)
 - [CreateNetworkingV1PeeringRequestAllOf1](docs/CreateNetworkingV1PeeringRequestAllOf1.md)
 - [CreateNetworkingV1PeeringRequestAllOf1Spec](docs/CreateNetworkingV1PeeringRequestAllOf1Spec.md)
 - [CreateNetworkingV1PrivateLinkAccess202Response](docs/CreateNetworkingV1PrivateLinkAccess202Response.md)
 - [CreateNetworkingV1PrivateLinkAccessRequest](docs/CreateNetworkingV1PrivateLinkAccessRequest.md)
 - [CreateNetworkingV1TransitGatewayAttachment202Response](docs/CreateNetworkingV1TransitGatewayAttachment202Response.md)
 - [CreateNetworkingV1TransitGatewayAttachmentRequest](docs/CreateNetworkingV1TransitGatewayAttachmentRequest.md)
 - [CreateNotificationsV1IntegrationRequest](docs/CreateNotificationsV1IntegrationRequest.md)
 - [CreateNotificationsV1SubscriptionRequest](docs/CreateNotificationsV1SubscriptionRequest.md)
 - [CreateOrUpdateConnectv1ConnectorConfigRequest](docs/CreateOrUpdateConnectv1ConnectorConfigRequest.md)
 - [CreateOrgV2EnvironmentRequest](docs/CreateOrgV2EnvironmentRequest.md)
 - [CreatePartnerV2EntitlementRequest](docs/CreatePartnerV2EntitlementRequest.md)
 - [CreateSdV1Pipeline202Response](docs/CreateSdV1Pipeline202Response.md)
 - [CreateSdV1Pipeline202ResponseAllOf](docs/CreateSdV1Pipeline202ResponseAllOf.md)
 - [CreateSdV1PipelineRequest](docs/CreateSdV1PipelineRequest.md)
 - [CreateSdV1PipelineRequestAllOf](docs/CreateSdV1PipelineRequestAllOf.md)
 - [CreateSdV1PipelineRequestAllOf1](docs/CreateSdV1PipelineRequestAllOf1.md)
 - [CreateSdV1PipelineRequestAllOf1Spec](docs/CreateSdV1PipelineRequestAllOf1Spec.md)
 - [CreateSrcmV2Cluster202Response](docs/CreateSrcmV2Cluster202Response.md)
 - [CreateSrcmV2Cluster202ResponseAllOf](docs/CreateSrcmV2Cluster202ResponseAllOf.md)
 - [CreateSrcmV2Cluster202ResponseAllOf1](docs/CreateSrcmV2Cluster202ResponseAllOf1.md)
 - [CreateSrcmV2ClusterRequest](docs/CreateSrcmV2ClusterRequest.md)
 - [CreateSrcmV2ClusterRequestAllOf](docs/CreateSrcmV2ClusterRequestAllOf.md)
 - [CreateSrcmV2ClusterRequestAllOf1](docs/CreateSrcmV2ClusterRequestAllOf1.md)
 - [CreateSrcmV2ClusterRequestAllOf1Spec](docs/CreateSrcmV2ClusterRequestAllOf1Spec.md)
 - [CreateTopicRequestData](docs/CreateTopicRequestData.md)
 - [CreateTopicRequestDataConfigsInner](docs/CreateTopicRequestDataConfigsInner.md)
 - [DeleteConnectv1Connector200Response](docs/DeleteConnectv1Connector200Response.md)
 - [DeleteKafkaAcls200Response](docs/DeleteKafkaAcls200Response.md)
 - [Entity](docs/Entity.md)
 - [EntityHeader](docs/EntityHeader.md)
 - [EntityWithExtInfo](docs/EntityWithExtInfo.md)
 - [EnvScopedObjectReference](docs/EnvScopedObjectReference.md)
 - [Error](docs/Error.md)
 - [ErrorMessage](docs/ErrorMessage.md)
 - [ErrorSource](docs/ErrorSource.md)
 - [Failure](docs/Failure.md)
 - [GetByokV1Key200Response](docs/GetByokV1Key200Response.md)
 - [GetCmkV2Cluster200Response](docs/GetCmkV2Cluster200Response.md)
 - [GetCmkV2Cluster200ResponseAllOf](docs/GetCmkV2Cluster200ResponseAllOf.md)
 - [GetConnectv1ConnectorConfig200Response](docs/GetConnectv1ConnectorConfig200Response.md)
 - [GetIamV2ApiKey200Response](docs/GetIamV2ApiKey200Response.md)
 - [GetIamV2ApiKey200ResponseAllOf](docs/GetIamV2ApiKey200ResponseAllOf.md)
 - [GetIamV2IdentityPool200Response](docs/GetIamV2IdentityPool200Response.md)
 - [GetIamV2IdentityProvider200Response](docs/GetIamV2IdentityProvider200Response.md)
 - [GetIamV2Invitation200Response](docs/GetIamV2Invitation200Response.md)
 - [GetIamV2RoleBinding200Response](docs/GetIamV2RoleBinding200Response.md)
 - [GetIamV2ServiceAccount200Response](docs/GetIamV2ServiceAccount200Response.md)
 - [GetIamV2User200Response](docs/GetIamV2User200Response.md)
 - [GetKafkaQuotasV1ClientQuota200Response](docs/GetKafkaQuotasV1ClientQuota200Response.md)
 - [GetKafkaQuotasV1ClientQuota200ResponseAllOf](docs/GetKafkaQuotasV1ClientQuota200ResponseAllOf.md)
 - [GetKsqldbcmV2Cluster200Response](docs/GetKsqldbcmV2Cluster200Response.md)
 - [GetKsqldbcmV2Cluster200ResponseAllOf](docs/GetKsqldbcmV2Cluster200ResponseAllOf.md)
 - [GetNetworkingV1Network200Response](docs/GetNetworkingV1Network200Response.md)
 - [GetNetworkingV1Network200ResponseAllOf](docs/GetNetworkingV1Network200ResponseAllOf.md)
 - [GetNetworkingV1NetworkLinkEndpoint200Response](docs/GetNetworkingV1NetworkLinkEndpoint200Response.md)
 - [GetNetworkingV1NetworkLinkEndpoint200ResponseAllOf](docs/GetNetworkingV1NetworkLinkEndpoint200ResponseAllOf.md)
 - [GetNetworkingV1NetworkLinkService200Response](docs/GetNetworkingV1NetworkLinkService200Response.md)
 - [GetNetworkingV1NetworkLinkService200ResponseAllOf](docs/GetNetworkingV1NetworkLinkService200ResponseAllOf.md)
 - [GetNetworkingV1NetworkLinkServiceAssociation200Response](docs/GetNetworkingV1NetworkLinkServiceAssociation200Response.md)
 - [GetNetworkingV1NetworkLinkServiceAssociation200ResponseAllOf](docs/GetNetworkingV1NetworkLinkServiceAssociation200ResponseAllOf.md)
 - [GetNetworkingV1NetworkLinkServiceAssociation200ResponseAllOf1](docs/GetNetworkingV1NetworkLinkServiceAssociation200ResponseAllOf1.md)
 - [GetNetworkingV1Peering200Response](docs/GetNetworkingV1Peering200Response.md)
 - [GetNetworkingV1Peering200ResponseAllOf](docs/GetNetworkingV1Peering200ResponseAllOf.md)
 - [GetNetworkingV1PrivateLinkAccess200Response](docs/GetNetworkingV1PrivateLinkAccess200Response.md)
 - [GetNetworkingV1TransitGatewayAttachment200Response](docs/GetNetworkingV1TransitGatewayAttachment200Response.md)
 - [GetNotificationsV1Integration200Response](docs/GetNotificationsV1Integration200Response.md)
 - [GetNotificationsV1NotificationType200Response](docs/GetNotificationsV1NotificationType200Response.md)
 - [GetNotificationsV1Subscription200Response](docs/GetNotificationsV1Subscription200Response.md)
 - [GetOrgV2Environment200Response](docs/GetOrgV2Environment200Response.md)
 - [GetOrgV2Organization200Response](docs/GetOrgV2Organization200Response.md)
 - [GetPartnerV2Organization200Response](docs/GetPartnerV2Organization200Response.md)
 - [GetSdV1Pipeline200Response](docs/GetSdV1Pipeline200Response.md)
 - [GetSdV1Pipeline200ResponseAllOf](docs/GetSdV1Pipeline200ResponseAllOf.md)
 - [GetServiceQuotaV1AppliedQuota200Response](docs/GetServiceQuotaV1AppliedQuota200Response.md)
 - [GetServiceQuotaV1AppliedQuota200ResponseAllOf](docs/GetServiceQuotaV1AppliedQuota200ResponseAllOf.md)
 - [GetServiceQuotaV1Scope200Response](docs/GetServiceQuotaV1Scope200Response.md)
 - [GetSrcmV2Cluster200Response](docs/GetSrcmV2Cluster200Response.md)
 - [GetSrcmV2Cluster200ResponseAllOf](docs/GetSrcmV2Cluster200ResponseAllOf.md)
 - [GetSrcmV2Region200Response](docs/GetSrcmV2Region200Response.md)
 - [GetSrcmV2Region200ResponseAllOf](docs/GetSrcmV2Region200ResponseAllOf.md)
 - [GlobalObjectReference](docs/GlobalObjectReference.md)
 - [IamV2ApiKey](docs/IamV2ApiKey.md)
 - [IamV2ApiKeyList](docs/IamV2ApiKeyList.md)
 - [IamV2ApiKeyListDataInner](docs/IamV2ApiKeyListDataInner.md)
 - [IamV2ApiKeyListDataInnerAllOf](docs/IamV2ApiKeyListDataInnerAllOf.md)
 - [IamV2ApiKeyListMetadata](docs/IamV2ApiKeyListMetadata.md)
 - [IamV2ApiKeyListMetadataAllOf](docs/IamV2ApiKeyListMetadataAllOf.md)
 - [IamV2ApiKeyMetadata](docs/IamV2ApiKeyMetadata.md)
 - [IamV2ApiKeyMetadataAllOf](docs/IamV2ApiKeyMetadataAllOf.md)
 - [IamV2ApiKeySpec](docs/IamV2ApiKeySpec.md)
 - [IamV2ApiKeySpecOwner](docs/IamV2ApiKeySpecOwner.md)
 - [IamV2ApiKeySpecResource](docs/IamV2ApiKeySpecResource.md)
 - [IamV2ApiKeySpecUpdate](docs/IamV2ApiKeySpecUpdate.md)
 - [IamV2ApiKeyUpdate](docs/IamV2ApiKeyUpdate.md)
 - [IamV2IdentityPool](docs/IamV2IdentityPool.md)
 - [IamV2IdentityPoolList](docs/IamV2IdentityPoolList.md)
 - [IamV2IdentityPoolListDataInner](docs/IamV2IdentityPoolListDataInner.md)
 - [IamV2IdentityPoolListMetadata](docs/IamV2IdentityPoolListMetadata.md)
 - [IamV2IdentityPoolListMetadataAllOf](docs/IamV2IdentityPoolListMetadataAllOf.md)
 - [IamV2IdentityPoolMetadata](docs/IamV2IdentityPoolMetadata.md)
 - [IamV2IdentityPoolMetadataAllOf](docs/IamV2IdentityPoolMetadataAllOf.md)
 - [IamV2IdentityProvider](docs/IamV2IdentityProvider.md)
 - [IamV2IdentityProviderList](docs/IamV2IdentityProviderList.md)
 - [IamV2IdentityProviderListDataInner](docs/IamV2IdentityProviderListDataInner.md)
 - [IamV2IdentityProviderListMetadata](docs/IamV2IdentityProviderListMetadata.md)
 - [IamV2IdentityProviderListMetadataAllOf](docs/IamV2IdentityProviderListMetadataAllOf.md)
 - [IamV2IdentityProviderMetadata](docs/IamV2IdentityProviderMetadata.md)
 - [IamV2IdentityProviderMetadataAllOf](docs/IamV2IdentityProviderMetadataAllOf.md)
 - [IamV2IdentityProviderUpdate](docs/IamV2IdentityProviderUpdate.md)
 - [IamV2Invitation](docs/IamV2Invitation.md)
 - [IamV2InvitationCreator](docs/IamV2InvitationCreator.md)
 - [IamV2InvitationList](docs/IamV2InvitationList.md)
 - [IamV2InvitationListDataInner](docs/IamV2InvitationListDataInner.md)
 - [IamV2InvitationListMetadata](docs/IamV2InvitationListMetadata.md)
 - [IamV2InvitationListMetadataAllOf](docs/IamV2InvitationListMetadataAllOf.md)
 - [IamV2InvitationMetadata](docs/IamV2InvitationMetadata.md)
 - [IamV2InvitationMetadataAllOf](docs/IamV2InvitationMetadataAllOf.md)
 - [IamV2InvitationUser](docs/IamV2InvitationUser.md)
 - [IamV2Jwks](docs/IamV2Jwks.md)
 - [IamV2JwksObject](docs/IamV2JwksObject.md)
 - [IamV2JwksSpec](docs/IamV2JwksSpec.md)
 - [IamV2JwksStatus](docs/IamV2JwksStatus.md)
 - [IamV2RoleBinding](docs/IamV2RoleBinding.md)
 - [IamV2RoleBindingList](docs/IamV2RoleBindingList.md)
 - [IamV2RoleBindingListDataInner](docs/IamV2RoleBindingListDataInner.md)
 - [IamV2RoleBindingListMetadata](docs/IamV2RoleBindingListMetadata.md)
 - [IamV2RoleBindingListMetadataAllOf](docs/IamV2RoleBindingListMetadataAllOf.md)
 - [IamV2RoleBindingMetadata](docs/IamV2RoleBindingMetadata.md)
 - [IamV2RoleBindingMetadataAllOf](docs/IamV2RoleBindingMetadataAllOf.md)
 - [IamV2ServiceAccount](docs/IamV2ServiceAccount.md)
 - [IamV2ServiceAccountList](docs/IamV2ServiceAccountList.md)
 - [IamV2ServiceAccountListDataInner](docs/IamV2ServiceAccountListDataInner.md)
 - [IamV2ServiceAccountListMetadata](docs/IamV2ServiceAccountListMetadata.md)
 - [IamV2ServiceAccountListMetadataAllOf](docs/IamV2ServiceAccountListMetadataAllOf.md)
 - [IamV2ServiceAccountMetadata](docs/IamV2ServiceAccountMetadata.md)
 - [IamV2ServiceAccountMetadataAllOf](docs/IamV2ServiceAccountMetadataAllOf.md)
 - [IamV2ServiceAccountUpdate](docs/IamV2ServiceAccountUpdate.md)
 - [IamV2User](docs/IamV2User.md)
 - [IamV2UserList](docs/IamV2UserList.md)
 - [IamV2UserListDataInner](docs/IamV2UserListDataInner.md)
 - [IamV2UserListMetadata](docs/IamV2UserListMetadata.md)
 - [IamV2UserListMetadataAllOf](docs/IamV2UserListMetadataAllOf.md)
 - [IamV2UserMetadata](docs/IamV2UserMetadata.md)
 - [IamV2UserMetadataAllOf](docs/IamV2UserMetadataAllOf.md)
 - [IamV2UserUpdate](docs/IamV2UserUpdate.md)
 - [KafkaQuotasV1ClientQuota](docs/KafkaQuotasV1ClientQuota.md)
 - [KafkaQuotasV1ClientQuotaList](docs/KafkaQuotasV1ClientQuotaList.md)
 - [KafkaQuotasV1ClientQuotaListDataInner](docs/KafkaQuotasV1ClientQuotaListDataInner.md)
 - [KafkaQuotasV1ClientQuotaListDataInnerAllOf](docs/KafkaQuotasV1ClientQuotaListDataInnerAllOf.md)
 - [KafkaQuotasV1ClientQuotaListMetadata](docs/KafkaQuotasV1ClientQuotaListMetadata.md)
 - [KafkaQuotasV1ClientQuotaListMetadataAllOf](docs/KafkaQuotasV1ClientQuotaListMetadataAllOf.md)
 - [KafkaQuotasV1ClientQuotaMetadata](docs/KafkaQuotasV1ClientQuotaMetadata.md)
 - [KafkaQuotasV1ClientQuotaMetadataAllOf](docs/KafkaQuotasV1ClientQuotaMetadataAllOf.md)
 - [KafkaQuotasV1ClientQuotaSpec](docs/KafkaQuotasV1ClientQuotaSpec.md)
 - [KafkaQuotasV1ClientQuotaSpecCluster](docs/KafkaQuotasV1ClientQuotaSpecCluster.md)
 - [KafkaQuotasV1ClientQuotaSpecThroughput](docs/KafkaQuotasV1ClientQuotaSpecThroughput.md)
 - [KafkaQuotasV1ClientQuotaSpecUpdate](docs/KafkaQuotasV1ClientQuotaSpecUpdate.md)
 - [KafkaQuotasV1ClientQuotaUpdate](docs/KafkaQuotasV1ClientQuotaUpdate.md)
 - [KafkaQuotasV1Throughput](docs/KafkaQuotasV1Throughput.md)
 - [KsqldbcmV2Cluster](docs/KsqldbcmV2Cluster.md)
 - [KsqldbcmV2ClusterList](docs/KsqldbcmV2ClusterList.md)
 - [KsqldbcmV2ClusterListDataInner](docs/KsqldbcmV2ClusterListDataInner.md)
 - [KsqldbcmV2ClusterListDataInnerAllOf](docs/KsqldbcmV2ClusterListDataInnerAllOf.md)
 - [KsqldbcmV2ClusterListMetadata](docs/KsqldbcmV2ClusterListMetadata.md)
 - [KsqldbcmV2ClusterListMetadataAllOf](docs/KsqldbcmV2ClusterListMetadataAllOf.md)
 - [KsqldbcmV2ClusterMetadata](docs/KsqldbcmV2ClusterMetadata.md)
 - [KsqldbcmV2ClusterMetadataAllOf](docs/KsqldbcmV2ClusterMetadataAllOf.md)
 - [KsqldbcmV2ClusterSpec](docs/KsqldbcmV2ClusterSpec.md)
 - [KsqldbcmV2ClusterSpecCredentialIdentity](docs/KsqldbcmV2ClusterSpecCredentialIdentity.md)
 - [KsqldbcmV2ClusterSpecEnvironment](docs/KsqldbcmV2ClusterSpecEnvironment.md)
 - [KsqldbcmV2ClusterSpecKafkaCluster](docs/KsqldbcmV2ClusterSpecKafkaCluster.md)
 - [KsqldbcmV2ClusterStatus](docs/KsqldbcmV2ClusterStatus.md)
 - [ListBillingV1Costs200Response](docs/ListBillingV1Costs200Response.md)
 - [ListByokV1Keys200Response](docs/ListByokV1Keys200Response.md)
 - [ListCmkV2Clusters200Response](docs/ListCmkV2Clusters200Response.md)
 - [ListCmkV2Clusters200ResponseAllOf](docs/ListCmkV2Clusters200ResponseAllOf.md)
 - [ListCmkV2Clusters200ResponseAllOfDataInner](docs/ListCmkV2Clusters200ResponseAllOfDataInner.md)
 - [ListCmkV2Clusters200ResponseAllOfDataInnerSpec](docs/ListCmkV2Clusters200ResponseAllOfDataInnerSpec.md)
 - [ListConnectv1ConnectorPlugins200ResponseInner](docs/ListConnectv1ConnectorPlugins200ResponseInner.md)
 - [ListIamV2ApiKeys200Response](docs/ListIamV2ApiKeys200Response.md)
 - [ListIamV2ApiKeys200ResponseAllOf](docs/ListIamV2ApiKeys200ResponseAllOf.md)
 - [ListIamV2ApiKeys200ResponseAllOfDataInner](docs/ListIamV2ApiKeys200ResponseAllOfDataInner.md)
 - [ListIamV2ApiKeys200ResponseAllOfDataInnerSpec](docs/ListIamV2ApiKeys200ResponseAllOfDataInnerSpec.md)
 - [ListIamV2IdentityPools200Response](docs/ListIamV2IdentityPools200Response.md)
 - [ListIamV2IdentityProviders200Response](docs/ListIamV2IdentityProviders200Response.md)
 - [ListIamV2Invitations200Response](docs/ListIamV2Invitations200Response.md)
 - [ListIamV2RoleBindings200Response](docs/ListIamV2RoleBindings200Response.md)
 - [ListIamV2ServiceAccounts200Response](docs/ListIamV2ServiceAccounts200Response.md)
 - [ListIamV2Users200Response](docs/ListIamV2Users200Response.md)
 - [ListKafkaQuotasV1ClientQuotas200Response](docs/ListKafkaQuotasV1ClientQuotas200Response.md)
 - [ListKafkaQuotasV1ClientQuotas200ResponseAllOf](docs/ListKafkaQuotasV1ClientQuotas200ResponseAllOf.md)
 - [ListKafkaQuotasV1ClientQuotas200ResponseAllOfDataInner](docs/ListKafkaQuotasV1ClientQuotas200ResponseAllOfDataInner.md)
 - [ListKafkaQuotasV1ClientQuotas200ResponseAllOfDataInnerSpec](docs/ListKafkaQuotasV1ClientQuotas200ResponseAllOfDataInnerSpec.md)
 - [ListKsqldbcmV2Clusters200Response](docs/ListKsqldbcmV2Clusters200Response.md)
 - [ListKsqldbcmV2Clusters200ResponseAllOf](docs/ListKsqldbcmV2Clusters200ResponseAllOf.md)
 - [ListKsqldbcmV2Clusters200ResponseAllOfDataInner](docs/ListKsqldbcmV2Clusters200ResponseAllOfDataInner.md)
 - [ListKsqldbcmV2Clusters200ResponseAllOfDataInnerSpec](docs/ListKsqldbcmV2Clusters200ResponseAllOfDataInnerSpec.md)
 - [ListLinkConfigsResponseData](docs/ListLinkConfigsResponseData.md)
 - [ListLinkConfigsResponseDataAllOf](docs/ListLinkConfigsResponseDataAllOf.md)
 - [ListLinkConfigsResponseDataList](docs/ListLinkConfigsResponseDataList.md)
 - [ListLinkConfigsResponseDataListAllOf](docs/ListLinkConfigsResponseDataListAllOf.md)
 - [ListLinksResponseData](docs/ListLinksResponseData.md)
 - [ListLinksResponseDataAllOf](docs/ListLinksResponseDataAllOf.md)
 - [ListLinksResponseDataList](docs/ListLinksResponseDataList.md)
 - [ListLinksResponseDataListAllOf](docs/ListLinksResponseDataListAllOf.md)
 - [ListMeta](docs/ListMeta.md)
 - [ListMirrorTopicsResponseData](docs/ListMirrorTopicsResponseData.md)
 - [ListMirrorTopicsResponseDataAllOf](docs/ListMirrorTopicsResponseDataAllOf.md)
 - [ListMirrorTopicsResponseDataList](docs/ListMirrorTopicsResponseDataList.md)
 - [ListMirrorTopicsResponseDataListAllOf](docs/ListMirrorTopicsResponseDataListAllOf.md)
 - [ListNetworkingV1NetworkLinkEndpoints200Response](docs/ListNetworkingV1NetworkLinkEndpoints200Response.md)
 - [ListNetworkingV1NetworkLinkEndpoints200ResponseAllOf](docs/ListNetworkingV1NetworkLinkEndpoints200ResponseAllOf.md)
 - [ListNetworkingV1NetworkLinkEndpoints200ResponseAllOfDataInner](docs/ListNetworkingV1NetworkLinkEndpoints200ResponseAllOfDataInner.md)
 - [ListNetworkingV1NetworkLinkEndpoints200ResponseAllOfDataInnerSpec](docs/ListNetworkingV1NetworkLinkEndpoints200ResponseAllOfDataInnerSpec.md)
 - [ListNetworkingV1NetworkLinkServiceAssociations200Response](docs/ListNetworkingV1NetworkLinkServiceAssociations200Response.md)
 - [ListNetworkingV1NetworkLinkServiceAssociations200ResponseAllOf](docs/ListNetworkingV1NetworkLinkServiceAssociations200ResponseAllOf.md)
 - [ListNetworkingV1NetworkLinkServiceAssociations200ResponseAllOfDataInner](docs/ListNetworkingV1NetworkLinkServiceAssociations200ResponseAllOfDataInner.md)
 - [ListNetworkingV1NetworkLinkServiceAssociations200ResponseAllOfDataInnerSpec](docs/ListNetworkingV1NetworkLinkServiceAssociations200ResponseAllOfDataInnerSpec.md)
 - [ListNetworkingV1NetworkLinkServices200Response](docs/ListNetworkingV1NetworkLinkServices200Response.md)
 - [ListNetworkingV1Networks200Response](docs/ListNetworkingV1Networks200Response.md)
 - [ListNetworkingV1Networks200ResponseAllOf](docs/ListNetworkingV1Networks200ResponseAllOf.md)
 - [ListNetworkingV1Networks200ResponseAllOfDataInner](docs/ListNetworkingV1Networks200ResponseAllOfDataInner.md)
 - [ListNetworkingV1Networks200ResponseAllOfDataInnerSpec](docs/ListNetworkingV1Networks200ResponseAllOfDataInnerSpec.md)
 - [ListNetworkingV1Peerings200Response](docs/ListNetworkingV1Peerings200Response.md)
 - [ListNetworkingV1Peerings200ResponseAllOf](docs/ListNetworkingV1Peerings200ResponseAllOf.md)
 - [ListNetworkingV1Peerings200ResponseAllOfDataInner](docs/ListNetworkingV1Peerings200ResponseAllOfDataInner.md)
 - [ListNetworkingV1Peerings200ResponseAllOfDataInnerSpec](docs/ListNetworkingV1Peerings200ResponseAllOfDataInnerSpec.md)
 - [ListNetworkingV1PrivateLinkAccesses200Response](docs/ListNetworkingV1PrivateLinkAccesses200Response.md)
 - [ListNetworkingV1TransitGatewayAttachments200Response](docs/ListNetworkingV1TransitGatewayAttachments200Response.md)
 - [ListNotificationsV1Integrations200Response](docs/ListNotificationsV1Integrations200Response.md)
 - [ListNotificationsV1NotificationTypes200Response](docs/ListNotificationsV1NotificationTypes200Response.md)
 - [ListNotificationsV1Subscriptions200Response](docs/ListNotificationsV1Subscriptions200Response.md)
 - [ListOrgV2Environments200Response](docs/ListOrgV2Environments200Response.md)
 - [ListOrgV2Organizations200Response](docs/ListOrgV2Organizations200Response.md)
 - [ListSdV1Pipelines200Response](docs/ListSdV1Pipelines200Response.md)
 - [ListSdV1Pipelines200ResponseAllOf](docs/ListSdV1Pipelines200ResponseAllOf.md)
 - [ListSdV1Pipelines200ResponseAllOfDataInner](docs/ListSdV1Pipelines200ResponseAllOfDataInner.md)
 - [ListSdV1Pipelines200ResponseAllOfDataInnerSpec](docs/ListSdV1Pipelines200ResponseAllOfDataInnerSpec.md)
 - [ListServiceQuotaV1AppliedQuotas200Response](docs/ListServiceQuotaV1AppliedQuotas200Response.md)
 - [ListServiceQuotaV1AppliedQuotas200ResponseAllOf](docs/ListServiceQuotaV1AppliedQuotas200ResponseAllOf.md)
 - [ListServiceQuotaV1AppliedQuotas200ResponseAllOfDataInner](docs/ListServiceQuotaV1AppliedQuotas200ResponseAllOfDataInner.md)
 - [ListServiceQuotaV1Scopes200Response](docs/ListServiceQuotaV1Scopes200Response.md)
 - [ListSrcmV2Clusters200Response](docs/ListSrcmV2Clusters200Response.md)
 - [ListSrcmV2Clusters200ResponseAllOf](docs/ListSrcmV2Clusters200ResponseAllOf.md)
 - [ListSrcmV2Clusters200ResponseAllOfDataInner](docs/ListSrcmV2Clusters200ResponseAllOfDataInner.md)
 - [ListSrcmV2Clusters200ResponseAllOfDataInnerSpec](docs/ListSrcmV2Clusters200ResponseAllOfDataInnerSpec.md)
 - [ListSrcmV2Regions200Response](docs/ListSrcmV2Regions200Response.md)
 - [MirrorLag](docs/MirrorLag.md)
 - [MirrorTopicStatus](docs/MirrorTopicStatus.md)
 - [Mode](docs/Mode.md)
 - [ModeUpdateRequest](docs/ModeUpdateRequest.md)
 - [NetworkingV1AwsNetwork](docs/NetworkingV1AwsNetwork.md)
 - [NetworkingV1AwsPeering](docs/NetworkingV1AwsPeering.md)
 - [NetworkingV1AwsPrivateLinkAccess](docs/NetworkingV1AwsPrivateLinkAccess.md)
 - [NetworkingV1AwsTransitGatewayAttachment](docs/NetworkingV1AwsTransitGatewayAttachment.md)
 - [NetworkingV1AwsTransitGatewayAttachmentStatus](docs/NetworkingV1AwsTransitGatewayAttachmentStatus.md)
 - [NetworkingV1AzureNetwork](docs/NetworkingV1AzureNetwork.md)
 - [NetworkingV1AzurePeering](docs/NetworkingV1AzurePeering.md)
 - [NetworkingV1AzurePrivateLinkAccess](docs/NetworkingV1AzurePrivateLinkAccess.md)
 - [NetworkingV1DnsConfig](docs/NetworkingV1DnsConfig.md)
 - [NetworkingV1GcpNetwork](docs/NetworkingV1GcpNetwork.md)
 - [NetworkingV1GcpPeering](docs/NetworkingV1GcpPeering.md)
 - [NetworkingV1GcpPrivateServiceConnectAccess](docs/NetworkingV1GcpPrivateServiceConnectAccess.md)
 - [NetworkingV1Network](docs/NetworkingV1Network.md)
 - [NetworkingV1NetworkLinkEndpoint](docs/NetworkingV1NetworkLinkEndpoint.md)
 - [NetworkingV1NetworkLinkEndpointList](docs/NetworkingV1NetworkLinkEndpointList.md)
 - [NetworkingV1NetworkLinkEndpointListDataInner](docs/NetworkingV1NetworkLinkEndpointListDataInner.md)
 - [NetworkingV1NetworkLinkEndpointListDataInnerAllOf](docs/NetworkingV1NetworkLinkEndpointListDataInnerAllOf.md)
 - [NetworkingV1NetworkLinkEndpointListMetadata](docs/NetworkingV1NetworkLinkEndpointListMetadata.md)
 - [NetworkingV1NetworkLinkEndpointListMetadataAllOf](docs/NetworkingV1NetworkLinkEndpointListMetadataAllOf.md)
 - [NetworkingV1NetworkLinkEndpointMetadata](docs/NetworkingV1NetworkLinkEndpointMetadata.md)
 - [NetworkingV1NetworkLinkEndpointMetadataAllOf](docs/NetworkingV1NetworkLinkEndpointMetadataAllOf.md)
 - [NetworkingV1NetworkLinkEndpointSpec](docs/NetworkingV1NetworkLinkEndpointSpec.md)
 - [NetworkingV1NetworkLinkEndpointSpecNetworkLinkService](docs/NetworkingV1NetworkLinkEndpointSpecNetworkLinkService.md)
 - [NetworkingV1NetworkLinkEndpointSpecUpdate](docs/NetworkingV1NetworkLinkEndpointSpecUpdate.md)
 - [NetworkingV1NetworkLinkEndpointStatus](docs/NetworkingV1NetworkLinkEndpointStatus.md)
 - [NetworkingV1NetworkLinkEndpointUpdate](docs/NetworkingV1NetworkLinkEndpointUpdate.md)
 - [NetworkingV1NetworkLinkService](docs/NetworkingV1NetworkLinkService.md)
 - [NetworkingV1NetworkLinkServiceAcceptPolicy](docs/NetworkingV1NetworkLinkServiceAcceptPolicy.md)
 - [NetworkingV1NetworkLinkServiceAssociation](docs/NetworkingV1NetworkLinkServiceAssociation.md)
 - [NetworkingV1NetworkLinkServiceAssociationList](docs/NetworkingV1NetworkLinkServiceAssociationList.md)
 - [NetworkingV1NetworkLinkServiceAssociationListDataInner](docs/NetworkingV1NetworkLinkServiceAssociationListDataInner.md)
 - [NetworkingV1NetworkLinkServiceAssociationListDataInnerAllOf](docs/NetworkingV1NetworkLinkServiceAssociationListDataInnerAllOf.md)
 - [NetworkingV1NetworkLinkServiceAssociationListMetadata](docs/NetworkingV1NetworkLinkServiceAssociationListMetadata.md)
 - [NetworkingV1NetworkLinkServiceAssociationListMetadataAllOf](docs/NetworkingV1NetworkLinkServiceAssociationListMetadataAllOf.md)
 - [NetworkingV1NetworkLinkServiceAssociationMetadata](docs/NetworkingV1NetworkLinkServiceAssociationMetadata.md)
 - [NetworkingV1NetworkLinkServiceAssociationMetadataAllOf](docs/NetworkingV1NetworkLinkServiceAssociationMetadataAllOf.md)
 - [NetworkingV1NetworkLinkServiceAssociationSpec](docs/NetworkingV1NetworkLinkServiceAssociationSpec.md)
 - [NetworkingV1NetworkLinkServiceAssociationStatus](docs/NetworkingV1NetworkLinkServiceAssociationStatus.md)
 - [NetworkingV1NetworkLinkServiceList](docs/NetworkingV1NetworkLinkServiceList.md)
 - [NetworkingV1NetworkLinkServiceListDataInner](docs/NetworkingV1NetworkLinkServiceListDataInner.md)
 - [NetworkingV1NetworkLinkServiceListDataInnerAllOf](docs/NetworkingV1NetworkLinkServiceListDataInnerAllOf.md)
 - [NetworkingV1NetworkLinkServiceListMetadata](docs/NetworkingV1NetworkLinkServiceListMetadata.md)
 - [NetworkingV1NetworkLinkServiceListMetadataAllOf](docs/NetworkingV1NetworkLinkServiceListMetadataAllOf.md)
 - [NetworkingV1NetworkLinkServiceMetadata](docs/NetworkingV1NetworkLinkServiceMetadata.md)
 - [NetworkingV1NetworkLinkServiceMetadataAllOf](docs/NetworkingV1NetworkLinkServiceMetadataAllOf.md)
 - [NetworkingV1NetworkLinkServiceSpec](docs/NetworkingV1NetworkLinkServiceSpec.md)
 - [NetworkingV1NetworkLinkServiceSpecAccept](docs/NetworkingV1NetworkLinkServiceSpecAccept.md)
 - [NetworkingV1NetworkLinkServiceSpecNetwork](docs/NetworkingV1NetworkLinkServiceSpecNetwork.md)
 - [NetworkingV1NetworkLinkServiceSpecUpdate](docs/NetworkingV1NetworkLinkServiceSpecUpdate.md)
 - [NetworkingV1NetworkLinkServiceStatus](docs/NetworkingV1NetworkLinkServiceStatus.md)
 - [NetworkingV1NetworkLinkServiceUpdate](docs/NetworkingV1NetworkLinkServiceUpdate.md)
 - [NetworkingV1NetworkList](docs/NetworkingV1NetworkList.md)
 - [NetworkingV1NetworkListDataInner](docs/NetworkingV1NetworkListDataInner.md)
 - [NetworkingV1NetworkListDataInnerAllOf](docs/NetworkingV1NetworkListDataInnerAllOf.md)
 - [NetworkingV1NetworkListMetadata](docs/NetworkingV1NetworkListMetadata.md)
 - [NetworkingV1NetworkListMetadataAllOf](docs/NetworkingV1NetworkListMetadataAllOf.md)
 - [NetworkingV1NetworkMetadata](docs/NetworkingV1NetworkMetadata.md)
 - [NetworkingV1NetworkMetadataAllOf](docs/NetworkingV1NetworkMetadataAllOf.md)
 - [NetworkingV1NetworkSpec](docs/NetworkingV1NetworkSpec.md)
 - [NetworkingV1NetworkSpecDnsConfig](docs/NetworkingV1NetworkSpecDnsConfig.md)
 - [NetworkingV1NetworkSpecEnvironment](docs/NetworkingV1NetworkSpecEnvironment.md)
 - [NetworkingV1NetworkSpecUpdate](docs/NetworkingV1NetworkSpecUpdate.md)
 - [NetworkingV1NetworkStatus](docs/NetworkingV1NetworkStatus.md)
 - [NetworkingV1NetworkStatusCloud](docs/NetworkingV1NetworkStatusCloud.md)
 - [NetworkingV1NetworkUpdate](docs/NetworkingV1NetworkUpdate.md)
 - [NetworkingV1Peering](docs/NetworkingV1Peering.md)
 - [NetworkingV1PeeringList](docs/NetworkingV1PeeringList.md)
 - [NetworkingV1PeeringListDataInner](docs/NetworkingV1PeeringListDataInner.md)
 - [NetworkingV1PeeringListDataInnerAllOf](docs/NetworkingV1PeeringListDataInnerAllOf.md)
 - [NetworkingV1PeeringListMetadata](docs/NetworkingV1PeeringListMetadata.md)
 - [NetworkingV1PeeringListMetadataAllOf](docs/NetworkingV1PeeringListMetadataAllOf.md)
 - [NetworkingV1PeeringMetadata](docs/NetworkingV1PeeringMetadata.md)
 - [NetworkingV1PeeringMetadataAllOf](docs/NetworkingV1PeeringMetadataAllOf.md)
 - [NetworkingV1PeeringSpec](docs/NetworkingV1PeeringSpec.md)
 - [NetworkingV1PeeringSpecCloud](docs/NetworkingV1PeeringSpecCloud.md)
 - [NetworkingV1PeeringSpecNetwork](docs/NetworkingV1PeeringSpecNetwork.md)
 - [NetworkingV1PeeringSpecUpdate](docs/NetworkingV1PeeringSpecUpdate.md)
 - [NetworkingV1PeeringStatus](docs/NetworkingV1PeeringStatus.md)
 - [NetworkingV1PeeringUpdate](docs/NetworkingV1PeeringUpdate.md)
 - [NetworkingV1PrivateLinkAccess](docs/NetworkingV1PrivateLinkAccess.md)
 - [NetworkingV1PrivateLinkAccessList](docs/NetworkingV1PrivateLinkAccessList.md)
 - [NetworkingV1PrivateLinkAccessListDataInner](docs/NetworkingV1PrivateLinkAccessListDataInner.md)
 - [NetworkingV1PrivateLinkAccessListMetadata](docs/NetworkingV1PrivateLinkAccessListMetadata.md)
 - [NetworkingV1PrivateLinkAccessListMetadataAllOf](docs/NetworkingV1PrivateLinkAccessListMetadataAllOf.md)
 - [NetworkingV1PrivateLinkAccessMetadata](docs/NetworkingV1PrivateLinkAccessMetadata.md)
 - [NetworkingV1PrivateLinkAccessMetadataAllOf](docs/NetworkingV1PrivateLinkAccessMetadataAllOf.md)
 - [NetworkingV1PrivateLinkAccessSpec](docs/NetworkingV1PrivateLinkAccessSpec.md)
 - [NetworkingV1PrivateLinkAccessSpecCloud](docs/NetworkingV1PrivateLinkAccessSpecCloud.md)
 - [NetworkingV1PrivateLinkAccessSpecUpdate](docs/NetworkingV1PrivateLinkAccessSpecUpdate.md)
 - [NetworkingV1PrivateLinkAccessStatus](docs/NetworkingV1PrivateLinkAccessStatus.md)
 - [NetworkingV1PrivateLinkAccessUpdate](docs/NetworkingV1PrivateLinkAccessUpdate.md)
 - [NetworkingV1TransitGatewayAttachment](docs/NetworkingV1TransitGatewayAttachment.md)
 - [NetworkingV1TransitGatewayAttachmentList](docs/NetworkingV1TransitGatewayAttachmentList.md)
 - [NetworkingV1TransitGatewayAttachmentListDataInner](docs/NetworkingV1TransitGatewayAttachmentListDataInner.md)
 - [NetworkingV1TransitGatewayAttachmentListMetadata](docs/NetworkingV1TransitGatewayAttachmentListMetadata.md)
 - [NetworkingV1TransitGatewayAttachmentListMetadataAllOf](docs/NetworkingV1TransitGatewayAttachmentListMetadataAllOf.md)
 - [NetworkingV1TransitGatewayAttachmentMetadata](docs/NetworkingV1TransitGatewayAttachmentMetadata.md)
 - [NetworkingV1TransitGatewayAttachmentMetadataAllOf](docs/NetworkingV1TransitGatewayAttachmentMetadataAllOf.md)
 - [NetworkingV1TransitGatewayAttachmentSpec](docs/NetworkingV1TransitGatewayAttachmentSpec.md)
 - [NetworkingV1TransitGatewayAttachmentSpecCloud](docs/NetworkingV1TransitGatewayAttachmentSpecCloud.md)
 - [NetworkingV1TransitGatewayAttachmentSpecUpdate](docs/NetworkingV1TransitGatewayAttachmentSpecUpdate.md)
 - [NetworkingV1TransitGatewayAttachmentStatus](docs/NetworkingV1TransitGatewayAttachmentStatus.md)
 - [NetworkingV1TransitGatewayAttachmentStatusCloud](docs/NetworkingV1TransitGatewayAttachmentStatusCloud.md)
 - [NetworkingV1TransitGatewayAttachmentUpdate](docs/NetworkingV1TransitGatewayAttachmentUpdate.md)
 - [NetworkingV1ZoneInfo](docs/NetworkingV1ZoneInfo.md)
 - [NotificationsV1Integration](docs/NotificationsV1Integration.md)
 - [NotificationsV1IntegrationList](docs/NotificationsV1IntegrationList.md)
 - [NotificationsV1IntegrationListDataInner](docs/NotificationsV1IntegrationListDataInner.md)
 - [NotificationsV1IntegrationListMetadata](docs/NotificationsV1IntegrationListMetadata.md)
 - [NotificationsV1IntegrationListMetadataAllOf](docs/NotificationsV1IntegrationListMetadataAllOf.md)
 - [NotificationsV1IntegrationMetadata](docs/NotificationsV1IntegrationMetadata.md)
 - [NotificationsV1IntegrationMetadataAllOf](docs/NotificationsV1IntegrationMetadataAllOf.md)
 - [NotificationsV1IntegrationTarget](docs/NotificationsV1IntegrationTarget.md)
 - [NotificationsV1MsTeamsTarget](docs/NotificationsV1MsTeamsTarget.md)
 - [NotificationsV1NotificationType](docs/NotificationsV1NotificationType.md)
 - [NotificationsV1NotificationTypeList](docs/NotificationsV1NotificationTypeList.md)
 - [NotificationsV1NotificationTypeListDataInner](docs/NotificationsV1NotificationTypeListDataInner.md)
 - [NotificationsV1NotificationTypeListMetadata](docs/NotificationsV1NotificationTypeListMetadata.md)
 - [NotificationsV1NotificationTypeListMetadataAllOf](docs/NotificationsV1NotificationTypeListMetadataAllOf.md)
 - [NotificationsV1NotificationTypeMetadata](docs/NotificationsV1NotificationTypeMetadata.md)
 - [NotificationsV1NotificationTypeMetadataAllOf](docs/NotificationsV1NotificationTypeMetadataAllOf.md)
 - [NotificationsV1RoleEmailTarget](docs/NotificationsV1RoleEmailTarget.md)
 - [NotificationsV1SlackTarget](docs/NotificationsV1SlackTarget.md)
 - [NotificationsV1Subscription](docs/NotificationsV1Subscription.md)
 - [NotificationsV1SubscriptionList](docs/NotificationsV1SubscriptionList.md)
 - [NotificationsV1SubscriptionListDataInner](docs/NotificationsV1SubscriptionListDataInner.md)
 - [NotificationsV1SubscriptionListMetadata](docs/NotificationsV1SubscriptionListMetadata.md)
 - [NotificationsV1SubscriptionListMetadataAllOf](docs/NotificationsV1SubscriptionListMetadataAllOf.md)
 - [NotificationsV1SubscriptionMetadata](docs/NotificationsV1SubscriptionMetadata.md)
 - [NotificationsV1SubscriptionMetadataAllOf](docs/NotificationsV1SubscriptionMetadataAllOf.md)
 - [NotificationsV1SubscriptionNotificationType](docs/NotificationsV1SubscriptionNotificationType.md)
 - [NotificationsV1SubscriptionUpdate](docs/NotificationsV1SubscriptionUpdate.md)
 - [NotificationsV1Target](docs/NotificationsV1Target.md)
 - [NotificationsV1UserEmailTarget](docs/NotificationsV1UserEmailTarget.md)
 - [NotificationsV1WebhookTarget](docs/NotificationsV1WebhookTarget.md)
 - [ObjectMeta](docs/ObjectMeta.md)
 - [ObjectReference](docs/ObjectReference.md)
 - [OrgV2Environment](docs/OrgV2Environment.md)
 - [OrgV2EnvironmentList](docs/OrgV2EnvironmentList.md)
 - [OrgV2EnvironmentListDataInner](docs/OrgV2EnvironmentListDataInner.md)
 - [OrgV2EnvironmentListMetadata](docs/OrgV2EnvironmentListMetadata.md)
 - [OrgV2EnvironmentListMetadataAllOf](docs/OrgV2EnvironmentListMetadataAllOf.md)
 - [OrgV2EnvironmentMetadata](docs/OrgV2EnvironmentMetadata.md)
 - [OrgV2EnvironmentMetadataAllOf](docs/OrgV2EnvironmentMetadataAllOf.md)
 - [OrgV2Organization](docs/OrgV2Organization.md)
 - [OrgV2OrganizationList](docs/OrgV2OrganizationList.md)
 - [OrgV2OrganizationListDataInner](docs/OrgV2OrganizationListDataInner.md)
 - [OrgV2OrganizationListMetadata](docs/OrgV2OrganizationListMetadata.md)
 - [OrgV2OrganizationListMetadataAllOf](docs/OrgV2OrganizationListMetadataAllOf.md)
 - [OrgV2OrganizationMetadata](docs/OrgV2OrganizationMetadata.md)
 - [OrgV2OrganizationMetadataAllOf](docs/OrgV2OrganizationMetadataAllOf.md)
 - [PartitionData](docs/PartitionData.md)
 - [PartitionDataAllOf](docs/PartitionDataAllOf.md)
 - [PartitionDataList](docs/PartitionDataList.md)
 - [PartitionDataListAllOf](docs/PartitionDataListAllOf.md)
 - [PartnerLinkRequest](docs/PartnerLinkRequest.md)
 - [PartnerLinkRequestOrganization](docs/PartnerLinkRequestOrganization.md)
 - [PartnerSignupRequest](docs/PartnerSignupRequest.md)
 - [PartnerSignupRequestEntitlement](docs/PartnerSignupRequestEntitlement.md)
 - [PartnerSignupRequestEntitlementOneOf](docs/PartnerSignupRequestEntitlementOneOf.md)
 - [PartnerSignupRequestEntitlementOneOf1](docs/PartnerSignupRequestEntitlementOneOf1.md)
 - [PartnerSignupRequestOrganization](docs/PartnerSignupRequestOrganization.md)
 - [PartnerSignupResponse](docs/PartnerSignupResponse.md)
 - [PartnerV2Entitlement](docs/PartnerV2Entitlement.md)
 - [PartnerV2EntitlementList](docs/PartnerV2EntitlementList.md)
 - [PartnerV2EntitlementListDataInner](docs/PartnerV2EntitlementListDataInner.md)
 - [PartnerV2EntitlementOrganization](docs/PartnerV2EntitlementOrganization.md)
 - [PartnerV2Organization](docs/PartnerV2Organization.md)
 - [PartnerV2OrganizationList](docs/PartnerV2OrganizationList.md)
 - [PartnerV2OrganizationListDataInner](docs/PartnerV2OrganizationListDataInner.md)
 - [PartnerV2OrganizationSsoConfig](docs/PartnerV2OrganizationSsoConfig.md)
 - [ProduceRequest](docs/ProduceRequest.md)
 - [ProduceRequestData](docs/ProduceRequestData.md)
 - [ProduceRequestHeader](docs/ProduceRequestHeader.md)
 - [ProduceResponse](docs/ProduceResponse.md)
 - [ProduceResponseData](docs/ProduceResponseData.md)
 - [ReadConnectv1ConnectorStatus200Response](docs/ReadConnectv1ConnectorStatus200Response.md)
 - [ReadConnectv1ConnectorStatus200ResponseConnector](docs/ReadConnectv1ConnectorStatus200ResponseConnector.md)
 - [ReadConnectv1ConnectorStatus200ResponseTasksInner](docs/ReadConnectv1ConnectorStatus200ResponseTasksInner.md)
 - [ReassignmentData](docs/ReassignmentData.md)
 - [ReassignmentDataAllOf](docs/ReassignmentDataAllOf.md)
 - [ReassignmentDataList](docs/ReassignmentDataList.md)
 - [ReassignmentDataListAllOf](docs/ReassignmentDataListAllOf.md)
 - [RefreshIamV2JsonWebKeySet200Response](docs/RefreshIamV2JsonWebKeySet200Response.md)
 - [RefreshIamV2JsonWebKeySet200ResponseAllOf](docs/RefreshIamV2JsonWebKeySet200ResponseAllOf.md)
 - [RegisterSchemaRequest](docs/RegisterSchemaRequest.md)
 - [RegisterSchemaResponse](docs/RegisterSchemaResponse.md)
 - [Relationship](docs/Relationship.md)
 - [RemoveBrokerTaskData](docs/RemoveBrokerTaskData.md)
 - [RemoveBrokerTaskDataAllOf](docs/RemoveBrokerTaskDataAllOf.md)
 - [RemoveBrokerTaskDataList](docs/RemoveBrokerTaskDataList.md)
 - [RemoveBrokerTaskDataListAllOf](docs/RemoveBrokerTaskDataListAllOf.md)
 - [RemoveBrokersRequestData](docs/RemoveBrokersRequestData.md)
 - [ReplicaData](docs/ReplicaData.md)
 - [ReplicaDataAllOf](docs/ReplicaDataAllOf.md)
 - [ReplicaDataList](docs/ReplicaDataList.md)
 - [ReplicaDataListAllOf](docs/ReplicaDataListAllOf.md)
 - [ReplicaStatusData](docs/ReplicaStatusData.md)
 - [ReplicaStatusDataAllOf](docs/ReplicaStatusDataAllOf.md)
 - [ReplicaStatusDataList](docs/ReplicaStatusDataList.md)
 - [ReplicaStatusDataListAllOf](docs/ReplicaStatusDataListAllOf.md)
 - [Resource](docs/Resource.md)
 - [ResourceCollection](docs/ResourceCollection.md)
 - [ResourceCollectionMetadata](docs/ResourceCollectionMetadata.md)
 - [ResourceMetadata](docs/ResourceMetadata.md)
 - [Schema](docs/Schema.md)
 - [SchemaReference](docs/SchemaReference.md)
 - [SchemaString](docs/SchemaString.md)
 - [SdV1Pipeline](docs/SdV1Pipeline.md)
 - [SdV1PipelineList](docs/SdV1PipelineList.md)
 - [SdV1PipelineListDataInner](docs/SdV1PipelineListDataInner.md)
 - [SdV1PipelineListDataInnerAllOf](docs/SdV1PipelineListDataInnerAllOf.md)
 - [SdV1PipelineListMetadata](docs/SdV1PipelineListMetadata.md)
 - [SdV1PipelineListMetadataAllOf](docs/SdV1PipelineListMetadataAllOf.md)
 - [SdV1PipelineMetadata](docs/SdV1PipelineMetadata.md)
 - [SdV1PipelineMetadataAllOf](docs/SdV1PipelineMetadataAllOf.md)
 - [SdV1PipelineSpec](docs/SdV1PipelineSpec.md)
 - [SdV1PipelineSpecKafkaCluster](docs/SdV1PipelineSpecKafkaCluster.md)
 - [SdV1PipelineSpecKsqlCluster](docs/SdV1PipelineSpecKsqlCluster.md)
 - [SdV1PipelineSpecSourceCode](docs/SdV1PipelineSpecSourceCode.md)
 - [SdV1PipelineSpecStreamGovernanceCluster](docs/SdV1PipelineSpecStreamGovernanceCluster.md)
 - [SdV1PipelineSpecUpdate](docs/SdV1PipelineSpecUpdate.md)
 - [SdV1PipelineStatus](docs/SdV1PipelineStatus.md)
 - [SdV1PipelineUpdate](docs/SdV1PipelineUpdate.md)
 - [SdV1SourceCodeObject](docs/SdV1SourceCodeObject.md)
 - [SearchParams](docs/SearchParams.md)
 - [SearchResult](docs/SearchResult.md)
 - [ServiceQuotaV1AppliedQuota](docs/ServiceQuotaV1AppliedQuota.md)
 - [ServiceQuotaV1AppliedQuotaEnvironment](docs/ServiceQuotaV1AppliedQuotaEnvironment.md)
 - [ServiceQuotaV1AppliedQuotaKafkaCluster](docs/ServiceQuotaV1AppliedQuotaKafkaCluster.md)
 - [ServiceQuotaV1AppliedQuotaList](docs/ServiceQuotaV1AppliedQuotaList.md)
 - [ServiceQuotaV1AppliedQuotaListDataInner](docs/ServiceQuotaV1AppliedQuotaListDataInner.md)
 - [ServiceQuotaV1AppliedQuotaListMetadata](docs/ServiceQuotaV1AppliedQuotaListMetadata.md)
 - [ServiceQuotaV1AppliedQuotaListMetadataAllOf](docs/ServiceQuotaV1AppliedQuotaListMetadataAllOf.md)
 - [ServiceQuotaV1AppliedQuotaMetadata](docs/ServiceQuotaV1AppliedQuotaMetadata.md)
 - [ServiceQuotaV1AppliedQuotaMetadataAllOf](docs/ServiceQuotaV1AppliedQuotaMetadataAllOf.md)
 - [ServiceQuotaV1AppliedQuotaNetwork](docs/ServiceQuotaV1AppliedQuotaNetwork.md)
 - [ServiceQuotaV1AppliedQuotaOrganization](docs/ServiceQuotaV1AppliedQuotaOrganization.md)
 - [ServiceQuotaV1AppliedQuotaUser](docs/ServiceQuotaV1AppliedQuotaUser.md)
 - [ServiceQuotaV1Scope](docs/ServiceQuotaV1Scope.md)
 - [ServiceQuotaV1ScopeList](docs/ServiceQuotaV1ScopeList.md)
 - [ServiceQuotaV1ScopeListDataInner](docs/ServiceQuotaV1ScopeListDataInner.md)
 - [ServiceQuotaV1ScopeListMetadata](docs/ServiceQuotaV1ScopeListMetadata.md)
 - [ServiceQuotaV1ScopeListMetadataAllOf](docs/ServiceQuotaV1ScopeListMetadataAllOf.md)
 - [ServiceQuotaV1ScopeMetadata](docs/ServiceQuotaV1ScopeMetadata.md)
 - [ServiceQuotaV1ScopeMetadataAllOf](docs/ServiceQuotaV1ScopeMetadataAllOf.md)
 - [SrcmV2Cluster](docs/SrcmV2Cluster.md)
 - [SrcmV2ClusterList](docs/SrcmV2ClusterList.md)
 - [SrcmV2ClusterListDataInner](docs/SrcmV2ClusterListDataInner.md)
 - [SrcmV2ClusterListDataInnerAllOf](docs/SrcmV2ClusterListDataInnerAllOf.md)
 - [SrcmV2ClusterListMetadata](docs/SrcmV2ClusterListMetadata.md)
 - [SrcmV2ClusterListMetadataAllOf](docs/SrcmV2ClusterListMetadataAllOf.md)
 - [SrcmV2ClusterMetadata](docs/SrcmV2ClusterMetadata.md)
 - [SrcmV2ClusterMetadataAllOf](docs/SrcmV2ClusterMetadataAllOf.md)
 - [SrcmV2ClusterSpec](docs/SrcmV2ClusterSpec.md)
 - [SrcmV2ClusterSpecEnvironment](docs/SrcmV2ClusterSpecEnvironment.md)
 - [SrcmV2ClusterSpecRegion](docs/SrcmV2ClusterSpecRegion.md)
 - [SrcmV2ClusterSpecUpdate](docs/SrcmV2ClusterSpecUpdate.md)
 - [SrcmV2ClusterStatus](docs/SrcmV2ClusterStatus.md)
 - [SrcmV2ClusterUpdate](docs/SrcmV2ClusterUpdate.md)
 - [SrcmV2Region](docs/SrcmV2Region.md)
 - [SrcmV2RegionList](docs/SrcmV2RegionList.md)
 - [SrcmV2RegionListDataInner](docs/SrcmV2RegionListDataInner.md)
 - [SrcmV2RegionListDataInnerAllOf](docs/SrcmV2RegionListDataInnerAllOf.md)
 - [SrcmV2RegionListMetadata](docs/SrcmV2RegionListMetadata.md)
 - [SrcmV2RegionListMetadataAllOf](docs/SrcmV2RegionListMetadataAllOf.md)
 - [SrcmV2RegionMetadata](docs/SrcmV2RegionMetadata.md)
 - [SrcmV2RegionMetadataAllOf](docs/SrcmV2RegionMetadataAllOf.md)
 - [SrcmV2RegionSpec](docs/SrcmV2RegionSpec.md)
 - [StsV1TokenExchangeReply](docs/StsV1TokenExchangeReply.md)
 - [SubjectVersion](docs/SubjectVersion.md)
 - [Tag](docs/Tag.md)
 - [TagDef](docs/TagDef.md)
 - [TagDefResponse](docs/TagDefResponse.md)
 - [TagResponse](docs/TagResponse.md)
 - [TermAssignmentHeader](docs/TermAssignmentHeader.md)
 - [TimeBoundary](docs/TimeBoundary.md)
 - [TopicConfigData](docs/TopicConfigData.md)
 - [TopicConfigDataAllOf](docs/TopicConfigDataAllOf.md)
 - [TopicConfigDataList](docs/TopicConfigDataList.md)
 - [TopicConfigDataListAllOf](docs/TopicConfigDataListAllOf.md)
 - [TopicData](docs/TopicData.md)
 - [TopicDataAllOf](docs/TopicDataAllOf.md)
 - [TopicDataList](docs/TopicDataList.md)
 - [TopicDataListAllOf](docs/TopicDataListAllOf.md)
 - [TypedEnvScopedObjectReference](docs/TypedEnvScopedObjectReference.md)
 - [TypedGlobalObjectReference](docs/TypedGlobalObjectReference.md)
 - [UpdateCmkV2ClusterRequest](docs/UpdateCmkV2ClusterRequest.md)
 - [UpdateCmkV2ClusterRequestAllOf](docs/UpdateCmkV2ClusterRequestAllOf.md)
 - [UpdateCmkV2ClusterRequestAllOfSpec](docs/UpdateCmkV2ClusterRequestAllOfSpec.md)
 - [UpdateConfigRequestData](docs/UpdateConfigRequestData.md)
 - [UpdateKafkaQuotasV1ClientQuotaRequest](docs/UpdateKafkaQuotasV1ClientQuotaRequest.md)
 - [UpdateLinkConfigRequestData](docs/UpdateLinkConfigRequestData.md)
 - [UpdateNetworkingV1NetworkLinkEndpointRequest](docs/UpdateNetworkingV1NetworkLinkEndpointRequest.md)
 - [UpdateNetworkingV1NetworkLinkServiceRequest](docs/UpdateNetworkingV1NetworkLinkServiceRequest.md)
 - [UpdateNetworkingV1NetworkRequest](docs/UpdateNetworkingV1NetworkRequest.md)
 - [UpdateNetworkingV1PeeringRequest](docs/UpdateNetworkingV1PeeringRequest.md)
 - [UpdateNetworkingV1PrivateLinkAccessRequest](docs/UpdateNetworkingV1PrivateLinkAccessRequest.md)
 - [UpdateNetworkingV1TransitGatewayAttachmentRequest](docs/UpdateNetworkingV1TransitGatewayAttachmentRequest.md)
 - [UpdatePartitionCountRequestData](docs/UpdatePartitionCountRequestData.md)
 - [UpdateSdV1PipelineRequest](docs/UpdateSdV1PipelineRequest.md)
 - [UpdateSrcmV2ClusterRequest](docs/UpdateSrcmV2ClusterRequest.md)
 - [ValidateConnectv1ConnectorPlugin200Response](docs/ValidateConnectv1ConnectorPlugin200Response.md)
 - [ValidateConnectv1ConnectorPlugin200ResponseConfigsInner](docs/ValidateConnectv1ConnectorPlugin200ResponseConfigsInner.md)
 - [ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition](docs/ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition.md)
 - [ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerValue](docs/ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerValue.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### api-key

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

### confluent-sts-access-token


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: N/A

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```

### oauth


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **partner:alter**: enables partners to alter entitlements
 - **partner:create**: enables partners to create entitlements and signup on behalf of customers
 - **partner:delete**: enables partners to delete entitlements and organizations
 - **partner:describe**: enables partners to read and list entitlements and organizations

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


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

support@confluent.io

