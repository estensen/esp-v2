// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package flags includes all API producer configurable settings.

package flags

import (
	"flag"
	"time"

	"github.com/GoogleCloudPlatform/esp-v2/src/go/commonflags"
	"github.com/GoogleCloudPlatform/esp-v2/src/go/options"
	"github.com/golang/glog"
)

var (
	// These flags are kept in sync with options.ConfigGeneratorOptions.
	defaults = options.DefaultConfigGeneratorOptions()

	// Cors related configurations.
	CorsAllowCredentials = flag.Bool("cors_allow_credentials", defaults.CorsAllowCredentials, "whether include the Access-Control-Allow-Credentials header with the value true in responses or not")
	CorsAllowHeaders     = flag.String("cors_allow_headers", defaults.CorsAllowHeaders, "set Access-Control-Allow-Headers to the specified HTTP headers")
	CorsAllowMethods     = flag.String("cors_allow_methods", defaults.CorsAllowMethods, "set Access-Control-Allow-Methods to the specified HTTP methods")
	CorsAllowOrigin      = flag.String("cors_allow_origin", defaults.CorsAllowOrigin, "set Access-Control-Allow-Origin to a specific origin")
	CorsAllowOriginRegex = flag.String("cors_allow_origin_regex", defaults.CorsAllowOriginRegex, "set Access-Control-Allow-Origin to a regular expression")
	CorsExposeHeaders    = flag.String("cors_expose_headers", defaults.CorsExposeHeaders, "set Access-Control-Expose-Headers to the specified headers")
	CorsMaxAge           = flag.Duration("cors_max_age", defaults.CorsMaxAge, "set Access-Control-Max-Age response header for CORS preflight request.")
	CorsPreset           = flag.String("cors_preset", defaults.CorsPreset, `enable CORS support, must be either "basic" or "cors_with_regex"`)

	// Backend routing configurations.
	BackendDnsLookupFamily = flag.String("backend_dns_lookup_family", defaults.BackendDnsLookupFamily, `Define the dns lookup family for all backends. The options are "auto", "v4only", "v6only", "v4preferred" and "all". The default is "v4preferred". "auto" is a legacy name, it behaves as "v6preferred".`)

	// Envoy specific configurations.
	ClusterConnectTimeout = flag.Duration("cluster_connect_timeout", defaults.ClusterConnectTimeout, "cluster connect timeout in seconds")

	// Network related configurations.
	BackendAddress               = flag.String("backend_address", defaults.BackendAddress, `The application server URI to which ESPv2 proxies requests.`)
	ListenerAddress              = flag.String("listener_address", defaults.ListenerAddress, "listener socket ip address")
	ServiceManagementURL         = flag.String("service_management_url", defaults.ServiceManagementURL, "url of service management server")
	ServiceControlURL            = flag.String("service_control_url", defaults.ServiceControlURL, "url of service control server")
	EnableBackendAddressOverride = flag.Bool("enable_backend_address_override", defaults.EnableBackendAddressOverride, "Allow the --backend flag to override the backend.rule.address for all operations.")

	ListenerPort = flag.Int("listener_port", defaults.ListenerPort, "listener port")
	Healthz      = flag.String("healthz", defaults.Healthz, "path for health check of ESPv2 proxy itself")

	// Health check grpc backend related flags.
	HealthCheckGrpcBackend        = flag.Bool("health_check_grpc_backend", defaults.HealthCheckGrpcBackend, `If true, ESPv2 periodically checks the gRPC Health service for the backend specified by the flag "--backend_address".`)
	HealthCheckGrpcBackendService = flag.String("health_check_grpc_backend_service", defaults.HealthCheckGrpcBackendService, `Specify the service name in the HealthCheckRequest when calling the backend gRPC Health service.
                       Default is empty. It only applies when the flag "--health_check_grpc_backend" is used.`)
	HealthCheckGrpcBackendInterval = flag.Duration("health_check_grpc_backend_interval", defaults.HealthCheckGrpcBackendInterval, `Specify the checking interval to call the backend gRPC Health service. Default is 1 second.
                      It only applies when the flag "--health_check_grpc_backend" is used.`)
	HealthCheckGrpcBackendNoTrafficInterval = flag.Duration("health_check_grpc_backend_no_traffic_interval", defaults.HealthCheckGrpcBackendNoTrafficInterval, `Specify the checking interval to call the backend gRPC Health service
                      when at start up or the backend did not have any traffic. Default is 60 seconds. It only applies when the flag "--health_check_grpc_backend" is used.`)

	SslServerCertPath                = flag.String("ssl_server_cert_path", defaults.SslServerCertPath, "Path to the certificate and key that ESPv2 uses to act as a HTTPS server")
	SslServerCipherSuites            = flag.String("ssl_server_cipher_suites", defaults.SslServerCipherSuites, "Cipher suites to use for downstream connections as a comma-separated list.")
	SslServerRootCertsPath           = flag.String("ssl_server_root_cert_path", defaults.SslServerRootCertPath, "The file path of root certificates that ESPv2 uses to verify downstream client certificate. If not specified, ESPv2 doesn't verify client certificates by default")
	SslSidestreamClientRootCertsPath = flag.String("ssl_sidestream_client_root_certs_path", defaults.SslSidestreamClientRootCertsPath, "Path to the root certificates to make TLS connection to all external services other than the backend.")
	SslBackendClientCertPath         = flag.String("ssl_backend_client_cert_path", defaults.SslBackendClientCertPath, "Path to the certificate and key that ESPv2 uses to enable TLS mutual authentication for HTTPS backend")
	SslBackendClientRootCertsPath    = flag.String("ssl_backend_client_root_certs_path", defaults.SslBackendClientRootCertsPath, "Path to the root certificates to make TLS connection to the HTTPS backend.")
	SslBackendClientCipherSuites     = flag.String("ssl_backend_client_cipher_suites", defaults.SslBackendClientCipherSuites, "Cipher suites to use for HTTPS backends as a comma-separated list.")
	SslMinimumProtocol               = flag.String("ssl_minimum_protocol", defaults.SslMinimumProtocol, "Minimum TLS protocol version for Downstream connections.")
	SslMaximumProtocol               = flag.String("ssl_maximum_protocol", defaults.SslMaximumProtocol, "Maximum TLS protocol version for Downstream connections.")
	EnableHSTS                       = flag.Bool("enable_strict_transport_security", defaults.EnableHSTS, "Enable HSTS (HTTP Strict Transport Security).")
	DnsResolverAddresses             = flag.String("dns_resolver_addresses", defaults.DnsResolverAddresses, `The addresses of dns resolvers. Each address should be in format of either IP_ADDR or IP_ADDR:PORT and they are separated by ';'.`)

	AddRequestHeaders = flag.String("add_request_headers", defaults.AddRequestHeaders, `Add HTTP headers to the request before sent to the upstream backend. Multiple headers are separated by ';'.
         For example --add_request_headers=key1=value1;key2=value2. If a header is already in the request, its value will be replaced with the new one.`)
	AppendRequestHeaders = flag.String("append_request_headers", defaults.AppendRequestHeaders, `Append HTTP headers to the request before sent to the upstream backend. Multiple headers are separated by ';'.
         For example --append_request_headers=key1=value1;key2=value2. If a header is already in the request, the new value will be append.`)
	AddResponseHeaders = flag.String("add_response_headers", defaults.AddResponseHeaders, `Add HTTP headers to the response before sent to the upstream backend. Multiple headers are separated by ';'.
         For example --add_response_headers=key1=value1;key2=value2. If a header is already in the response, its value will be replaced with the new one.`)
	AppendResponseHeaders = flag.String("append_response_headers", defaults.AppendResponseHeaders, `Append HTTP headers to the response before sent to the upstream backend. Multiple headers are separated by ';'.
         For example --append_response_headers=key1=value1;key2=value2. If a header is already in the response, the new value will be append.`)
	EnableOperationNameHeader = flag.Bool("enable_operation_name_header", defaults.EnableOperationNameHeader, "If enabled, the operation name for the matched route will be sent to the upstream as a request header.")

	// Flags for non_gcp deployment.
	ServiceAccountKey = flag.String("service_account_key", defaults.ServiceAccountKey, `Use the service account key JSON file to access the service control and the
	service management.  You can also set {creds_key} environment variable to the location of the service account credentials JSON file. If the option is
  omitted, the proxy contacts the metadata service to fetch an access token`)
	TokenAgentPort = flag.Uint("token_agent_port", defaults.TokenAgentPort, "Port that configmanager use to setup server to provide envoy with access token using service account credential, for accessing servicecontrol.")

	// Flags for external calls.
	DisableOidcDiscovery = flag.Bool("disable_oidc_discovery", defaults.DisableOidcDiscovery, `Disable OpenID Connect Discovery. 
  When disabled, config generator will not make external calls to determine the JWKS URI, 
	but the 'jwks_uri' field must not be empty in any authentication provider. 
	This should be disabled when the URLs configured by the API Producer cannot be trusted.`)
	DependencyErrorBehavior = flag.String("dependency_error_behavior", defaults.DependencyErrorBehavior,
		`The behavior all Envoy filter will adhere to when waiting for external dependencies during filter config.
						Value must match the enum espv2.api.envoy.v11.http.common.DependencyErrorBehavior.`)

	// Envoy configurations.
	AccessLog       = flag.String("access_log", defaults.AccessLog, "Path to a local file to which the access log entries will be written")
	AccessLogFormat = flag.String("access_log_format", defaults.AccessLogFormat, `String format to specify the format of access log.
	If unset, the following format will be used.
	https://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log#default-format-string
	For the detailed format grammar, please refer to the following document.
	https://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log#format-strings`)

	EnvoyUseRemoteAddress  = flag.Bool("envoy_use_remote_address", defaults.EnvoyUseRemoteAddress, "Envoy HttpConnectionManager configuration, please refer to envoy documentation for detailed information.")
	EnvoyXffNumTrustedHops = flag.Int("envoy_xff_num_trusted_hops", defaults.EnvoyXffNumTrustedHops, "Envoy HttpConnectionManager configuration, please refer to envoy documentation for detailed information.")

	LogJwtPayloads = flag.String("log_jwt_payloads", defaults.LogJwtPayloads, `Log corresponding JWT JSON payload primitive fields through service control, separated by comma. Example, when --log_jwt_payload=sub,project_id, log
	will have jwt_payload: sub=[SUBJECT];project_id=[PROJECT_ID] if the fields are available. The value must be a primitive field, JSON objects and arrays will not be logged.`)
	LogRequestHeaders = flag.String("log_request_headers", defaults.LogRequestHeaders, `Log corresponding request headers through service control, separated by comma. Example, when --log_request_headers=
	foo,bar, endpoint log will have request_headers: foo=foo_value;bar=bar_value if values are available;`)
	LogResponseHeaders = flag.String("log_response_headers", defaults.LogResponseHeaders, `Log corresponding response headers through service control, separated by comma. Example, when --log_response_headers=
	foo,bar,endpoint log will have response_headers: foo=foo_value;bar=bar_value if values are available.`)
	MinStreamReportIntervalMs = flag.Uint64("min_stream_report_interval_ms", defaults.MinStreamReportIntervalMs, `Minimum amount of time (milliseconds) between sending intermediate reports on a stream and the default is 10000 if not set.`)

	SuppressEnvoyHeaders = flag.Bool("suppress_envoy_headers", defaults.SuppressEnvoyHeaders, `Do not add any additional x-envoy- headers to requests or responses. This only affects the router filter
	generated *x-envoy-* headers, other Envoy filters and the HTTP connection manager may continue to set x-envoy- headers.`)
	UnderscoresInHeaders         = flag.Bool("underscores_in_headers", defaults.UnderscoresInHeaders, `When true, ESPv2 allows HTTP headers name has underscore and pass it through. Otherwise, rejects the request.`)
	NormalizePath                = flag.Bool("normalize_path", defaults.NormalizePath, `Normalizes the path according to RFC 3986 before processing requests.`)
	MergeSlashesInPath           = flag.Bool("merge_slashes_in_path", defaults.MergeSlashesInPath, `Determines if adjacent slashes in the path are merged into one before processing requests.`)
	DisallowEscapedSlashesInPath = flag.Bool("disallow_escaped_slashes_in_path", defaults.DisallowEscapedSlashesInPath, `Determines if [%2F, %2f, %2C, %2c] characters in the path are disallowed.`)

	ServiceControlNetworkFailOpen = flag.Bool("service_control_network_fail_open", defaults.ServiceControlNetworkFailOpen, ` In case of network failures when connecting to Google service control,
        the requests will be allowed if this flag is on. The default is on.`)

	EnableGrpcForHttp1 = flag.Bool("enable_grpc_for_http1", defaults.EnableGrpcForHttp1, `Enable gRPC when the downstream is HTTP/1.1. The default is on.`)

	ConnectionBufferLimitBytes = flag.Int("connection_buffer_limit_bytes", defaults.ConnectionBufferLimitBytes, `Configure the maximum amount of data that is buffered for each request/response body. 
			If not provided, Envoy will decide the default value.`)

	DisableJwksAsyncFetch      = flag.Bool("disable_jwks_async_fetch", defaults.DisableJwksAsyncFetch, `When the feature is enabled, JWKS is fetched before processing any requests. When disabled, JWKS is fetched on-demand when processing the requests.`)
	JwksAsyncFetchFastListener = flag.Bool("jwks_async_fetch_fast_listener", defaults.JwksAsyncFetchFastListener, `Only apply when --disable_jwks_async_fetch flag is not set. This flag determines if the envoy will wait for jwks_async_fetch to complete before binding the listener port. If false, it will wait. Default is false.`)
	JwksCacheDurationInS       = flag.Int("jwks_cache_duration_in_s", defaults.JwksCacheDurationInS, "Specify JWT public key cache duration in seconds. The default is 5 minutes.")

	JwksFetchNumRetries                 = flag.Int("jwks_fetch_num_retries", defaults.JwksFetchNumRetries, `Specify the remote JWKS fetch retry policy's number of retries. The default is 0, meaning no retry policy applied.`)
	JwksFetchRetryBackOffBaseIntervalMs = flag.Int("jwks_fetch_retry_back_off_base_interval_ms", int(defaults.JwksFetchRetryBackOffBaseInterval.Milliseconds()), `Specify JWKS fetch retry exponential back off base interval in milliseconds. The default is 200 milliseconds.`)
	JwksFetchRetryBackOffMaxIntervalMs  = flag.Int("jwks_fetch_retry_back_off_max_interval_ms", int(defaults.JwksFetchRetryBackOffMaxInterval.Milliseconds()), `Specify JWKS fetch retry exponential back off maximum interval in milliseconds. The default is 32 seconds.`)
	JwtPatForwardPayloadHeader          = flag.Bool("jwt_pad_forward_payload_header", defaults.JwtPadForwardPayloadHeader, `For the JWT in request, the JWT payload is forwarded to backend in the "X-Endpoint-API-UserInfo"" header by default. 
Normally JWT based64 encode doesn’t add padding. If this flag is true, the header will be padded.`)
	JwtCacheSize = flag.Uint("jwt_cache_size", defaults.JwtCacheSize, `Specify JWT cache size, the number of unique JWT tokens in the cache. The cache only stores verified good tokens. If 0, JWT cache is disabled. It limits the memory usage. The cache used memory is roughly (token size + 64 bytes) per token. If not specified, the default is 100000.`)

	ScCheckTimeoutMs  = flag.Int("service_control_check_timeout_ms", defaults.ScCheckTimeoutMs, `Set the timeout in millisecond for service control Check request. Must be > 0 and the default is 1000 if not set.`)
	ScQuotaTimeoutMs  = flag.Int("service_control_quota_timeout_ms", defaults.ScQuotaTimeoutMs, `Set the timeout in millisecond for service control Quota request. Must be > 0 and the default is 1000 if not set.`)
	ScReportTimeoutMs = flag.Int("service_control_report_timeout_ms", defaults.ScReportTimeoutMs, `Set the timeout in millisecond for service control Report request. Must be > 0 and the default is 2000 if not set.`)

	ScCheckRetries  = flag.Int("service_control_check_retries", defaults.ScCheckRetries, `Set the retry times for service control Check request. Must be >= 0 and the default is 3 if not set.`)
	ScQuotaRetries  = flag.Int("service_control_quota_retries", defaults.ScQuotaRetries, `Set the retry times for service control Quota request. Must be >= 0 and the default is 1 if not set.`)
	ScReportRetries = flag.Int("service_control_report_retries", defaults.ScReportRetries, `Set the retry times for service control Report request. Must be >= 0 and the default is 5 if not set.`)

	ComputePlatformOverride = flag.String("compute_platform_override", defaults.ComputePlatformOverride, "the overridden platform where the proxy is running at")

	// Flags for testing purpose. They are not exposed to the user via start_proxy.py
	SkipJwtAuthnFilter       = flag.Bool("skip_jwt_authn_filter", defaults.SkipJwtAuthnFilter, "skip jwt authn filter, for test purpose")
	SkipServiceControlFilter = flag.Bool("skip_service_control_filter", defaults.SkipServiceControlFilter, "skip service control filter, for test purpose")
	StreamIdleTimeout        = flag.Duration("stream_idle_timeout_test_only", defaults.StreamIdleTimeout, "The amount of time HTTP/2 streams can exist without any activity. "+
		"Set `deadline` in the service config to override this global value on a per-route basis.")

	TranscodingAlwaysPrintPrimitiveFields         = flag.Bool("transcoding_always_print_primitive_fields", defaults.TranscodingAlwaysPrintPrimitiveFields, "Whether to always print primitive fields for grpc-json transcoding")
	TranscodingAlwaysPrintEnumsAsInts             = flag.Bool("transcoding_always_print_enums_as_ints", defaults.TranscodingAlwaysPrintPrimitiveFields, "Whether to always print enums as ints for grpc-json transcoding")
	TranscodingPreserveProtoFieldNames            = flag.Bool("transcoding_preserve_proto_field_names", defaults.TranscodingPreserveProtoFieldNames, "Whether to preserve proto field names for grpc-json transcoding")
	TranscodingIgnoreQueryParameters              = flag.String("transcoding_ignore_query_parameters", defaults.TranscodingIgnoreQueryParameters, "A list of query parameters(separated by comma) to be ignored for transcoding method mapping in grpc-json transcoding.")
	TranscodingIgnoreUnknownQueryParameters       = flag.Bool("transcoding_ignore_unknown_query_parameters", defaults.TranscodingIgnoreUnknownQueryParameters, "Whether to ignore query parameters that cannot be mapped to a corresponding protobuf field in grpc-json transcoding.")
	TranscodingQueryParametersDisableUnescapePlus = flag.Bool("transcoding_query_parameters_disable_unescape_plus", defaults.TranscodingIgnoreUnknownQueryParameters, `By default, unescape "+" to space when extracting variables in
           the query parameters in grpc-json transcoding. This is to support HTML 2.0<https://tools.ietf.org/html/rfc1866#section-8.2.1>. Set this flag to true to disable this feature.`)
	TranscodingMatchUnregisteredCustomVerb = flag.Bool("transcoding_match_unregistered_custom_verb", defaults.TranscodingMatchUnregisteredCustomVerb, `If true, try to match the custom verb even if it is unregistered. By default, only match when it is registered.
  According to the http template[1], the custom verb is **":" LITERAL** at the end of http template.
	
	For a request with */foo/bar:baz* and *:baz* is not registered in any url_template, here is the behavior change
		- if the field is not set, *:baz* will not be treated as custom verb, so it will match **/foo/{x=*}**.
		- if the field is set, *:baz* is treated as custom verb,  so it will NOT match **/foo/{x=*}** since the template doesn't use any custom verb.

	[1](https://github.com/googleapis/googleapis/blob/master/google/api/http.proto#L226-L231)`)

	BackendRetryOns = flag.String("backend_retry_ons", defaults.BackendRetryOns,
		`The conditions under which ESPv2 does retry on the backends. One or more
        retryOn conditions can be specified by comma-separated list. The default
        is "reset,connect-failure,refused-stream". Disable retry by setting this flag to empty.
				This retry setting will be applied to all the backends if you have multiple ones.

        All the retryOn conditions are defined in the following
        x-envoy-retry-on(https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/router_filter#x-envoy-retry-on) and 
        x-envoy-retry-grpc-on(https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/router_filter#x-envoy-retry-on).`)
	BackendRetryNum = flag.Uint("backend_retry_num", defaults.BackendRetryNum,
		`The allowed number of retries. Must be >= 0 and defaults to 1. This retry
        setting will be applied to all the backends if you have multiple ones.`)
	BackendPerTryTimeout = flag.Duration("backend_per_try_timeout", defaults.BackendPerTryTimeout,
		`The backend timeout per retry attempt in second. Please note the 
        "deadline"" in the "x-google-backend"" extension is the total time wait
        for a full response from one request, including all retries. By default,
        backend_per_try_timeout=0 means ESPv2 will use the  "deadline"" in
        the "x-google-backend" extension. Consequently, a request that times out
        will not be retried as the total timeout budget would have been exhausted.`)
	BackendRetryOnStatusCodes = flag.String("backend_retry_on_status_codes", defaults.BackendRetryOnStatusCodes,
		`The list of backend http status codes will be retried, in
        addition to the status codes enabled for retry through other retry
        policies set in "--backend_retry_ons".
        The format is a comma-delimited String, like "501, 503`)

	EnableResponseCompression = flag.Bool("enable_response_compression", defaults.EnableResponseCompression, `Enable gzip,br compression for response data. The default is disabled.`)

	// BackendClusterMaxRequests is the maximum active requests allowed in a backend cluster.
	BackendClusterMaxRequests = flag.Int("backend_cluster_maximum_requests", defaults.BackendClusterMaxRequests,
		`The maximum allowed active requests for a backend cluster. If 0, or not set, default is 1024.
		It is the "cluster maximum requests" of Envoy circuit breaker settings(https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/circuit_breaking#circuit-breaking) that will be applied to all backend clusters.`)
)

func EnvoyConfigOptionsFromFlags() options.ConfigGeneratorOptions {
	opts := options.ConfigGeneratorOptions{
		CommonOptions:                                 commonflags.DefaultCommonOptionsFromFlags(),
		BackendAddress:                                *BackendAddress,
		EnableBackendAddressOverride:                  *EnableBackendAddressOverride,
		AccessLog:                                     *AccessLog,
		AccessLogFormat:                               *AccessLogFormat,
		ComputePlatformOverride:                       *ComputePlatformOverride,
		CorsAllowCredentials:                          *CorsAllowCredentials,
		CorsAllowHeaders:                              *CorsAllowHeaders,
		CorsAllowMethods:                              *CorsAllowMethods,
		CorsAllowOrigin:                               *CorsAllowOrigin,
		CorsAllowOriginRegex:                          *CorsAllowOriginRegex,
		CorsExposeHeaders:                             *CorsExposeHeaders,
		CorsMaxAge:                                    *CorsMaxAge,
		CorsPreset:                                    *CorsPreset,
		BackendDnsLookupFamily:                        *BackendDnsLookupFamily,
		ClusterConnectTimeout:                         *ClusterConnectTimeout,
		StreamIdleTimeout:                             *StreamIdleTimeout,
		ListenerAddress:                               *ListenerAddress,
		ServiceManagementURL:                          *ServiceManagementURL,
		ServiceControlURL:                             *ServiceControlURL,
		ListenerPort:                                  *ListenerPort,
		Healthz:                                       *Healthz,
		HealthCheckGrpcBackend:                        *HealthCheckGrpcBackend,
		HealthCheckGrpcBackendService:                 *HealthCheckGrpcBackendService,
		HealthCheckGrpcBackendInterval:                *HealthCheckGrpcBackendInterval,
		HealthCheckGrpcBackendNoTrafficInterval:       *HealthCheckGrpcBackendNoTrafficInterval,
		SslSidestreamClientRootCertsPath:              *SslSidestreamClientRootCertsPath,
		SslBackendClientCertPath:                      *SslBackendClientCertPath,
		SslBackendClientRootCertsPath:                 *SslBackendClientRootCertsPath,
		SslBackendClientCipherSuites:                  *SslBackendClientCipherSuites,
		SslServerCertPath:                             *SslServerCertPath,
		SslServerCipherSuites:                         *SslServerCipherSuites,
		SslServerRootCertPath:                         *SslServerRootCertsPath,
		SslMinimumProtocol:                            *SslMinimumProtocol,
		SslMaximumProtocol:                            *SslMaximumProtocol,
		EnableHSTS:                                    *EnableHSTS,
		DnsResolverAddresses:                          *DnsResolverAddresses,
		AddRequestHeaders:                             *AddRequestHeaders,
		AppendRequestHeaders:                          *AppendRequestHeaders,
		AddResponseHeaders:                            *AddResponseHeaders,
		AppendResponseHeaders:                         *AppendResponseHeaders,
		EnableOperationNameHeader:                     *EnableOperationNameHeader,
		ServiceAccountKey:                             *ServiceAccountKey,
		TokenAgentPort:                                *TokenAgentPort,
		DisableOidcDiscovery:                          *DisableOidcDiscovery,
		DependencyErrorBehavior:                       *DependencyErrorBehavior,
		SkipJwtAuthnFilter:                            *SkipJwtAuthnFilter,
		SkipServiceControlFilter:                      *SkipServiceControlFilter,
		EnvoyUseRemoteAddress:                         *EnvoyUseRemoteAddress,
		EnvoyXffNumTrustedHops:                        *EnvoyXffNumTrustedHops,
		LogJwtPayloads:                                *LogJwtPayloads,
		LogRequestHeaders:                             *LogRequestHeaders,
		LogResponseHeaders:                            *LogResponseHeaders,
		MinStreamReportIntervalMs:                     *MinStreamReportIntervalMs,
		SuppressEnvoyHeaders:                          *SuppressEnvoyHeaders,
		UnderscoresInHeaders:                          *UnderscoresInHeaders,
		NormalizePath:                                 *NormalizePath,
		MergeSlashesInPath:                            *MergeSlashesInPath,
		DisallowEscapedSlashesInPath:                  *DisallowEscapedSlashesInPath,
		ServiceControlNetworkFailOpen:                 *ServiceControlNetworkFailOpen,
		EnableGrpcForHttp1:                            *EnableGrpcForHttp1,
		ConnectionBufferLimitBytes:                    *ConnectionBufferLimitBytes,
		DisableJwksAsyncFetch:                         *DisableJwksAsyncFetch,
		JwksAsyncFetchFastListener:                    *JwksAsyncFetchFastListener,
		JwksCacheDurationInS:                          *JwksCacheDurationInS,
		JwksFetchNumRetries:                           *JwksFetchNumRetries,
		JwksFetchRetryBackOffBaseInterval:             time.Duration(*JwksFetchRetryBackOffBaseIntervalMs) * time.Millisecond,
		JwksFetchRetryBackOffMaxInterval:              time.Duration(*JwksFetchRetryBackOffMaxIntervalMs) * time.Millisecond,
		JwtPadForwardPayloadHeader:                    *JwtPatForwardPayloadHeader,
		JwtCacheSize:                                  *JwtCacheSize,
		BackendRetryOns:                               *BackendRetryOns,
		BackendRetryNum:                               *BackendRetryNum,
		BackendPerTryTimeout:                          *BackendPerTryTimeout,
		BackendRetryOnStatusCodes:                     *BackendRetryOnStatusCodes,
		ScCheckTimeoutMs:                              *ScCheckTimeoutMs,
		ScQuotaTimeoutMs:                              *ScQuotaTimeoutMs,
		ScReportTimeoutMs:                             *ScReportTimeoutMs,
		ScCheckRetries:                                *ScCheckRetries,
		ScQuotaRetries:                                *ScQuotaRetries,
		ScReportRetries:                               *ScReportRetries,
		BackendClusterMaxRequests:                     *BackendClusterMaxRequests,
		TranscodingAlwaysPrintPrimitiveFields:         *TranscodingAlwaysPrintPrimitiveFields,
		TranscodingAlwaysPrintEnumsAsInts:             *TranscodingAlwaysPrintEnumsAsInts,
		TranscodingPreserveProtoFieldNames:            *TranscodingPreserveProtoFieldNames,
		TranscodingIgnoreQueryParameters:              *TranscodingIgnoreQueryParameters,
		TranscodingIgnoreUnknownQueryParameters:       *TranscodingIgnoreUnknownQueryParameters,
		TranscodingQueryParametersDisableUnescapePlus: *TranscodingQueryParametersDisableUnescapePlus,
		TranscodingMatchUnregisteredCustomVerb:        *TranscodingMatchUnregisteredCustomVerb,
		EnableResponseCompression:                     *EnableResponseCompression,

		// These options are not for ESPv2 users. They are overridden internally.
		APIAllowList:       []string{},
		AllowDiscoveryAPIs: false,
	}

	glog.Infof("Config Generator options: %+v", opts)
	return opts
}
