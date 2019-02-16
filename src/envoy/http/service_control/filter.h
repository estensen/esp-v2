// Copyright 2018 Google Cloud Platform Proxy Authors
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

#pragma once

#include "common/common/logger.h"
#include "envoy/access_log/access_log.h"
#include "envoy/http/filter.h"
#include "src/envoy/http/service_control/filter_config.h"
#include "src/envoy/http/service_control/handler.h"

#include <string>

namespace Envoy {
namespace Extensions {
namespace HttpFilters {
namespace ServiceControl {

// The Envoy filter for Cloud ESF service control client.
class Filter : public Http::StreamDecoderFilter,
               public AccessLog::Instance,
               public Handler::CheckDoneCallback,
               public Logger::Loggable<Logger::Id::filter> {
 public:
  Filter(FilterConfigSharedPtr config) : config_(config) {}

  void onDestroy() override {}

  // Http::StreamDecoderFilter
  Http::FilterHeadersStatus decodeHeaders(Http::HeaderMap& headers,
                                          bool) override;
  Http::FilterDataStatus decodeData(Buffer::Instance&, bool) override;
  Http::FilterTrailersStatus decodeTrailers(Http::HeaderMap&) override;
  void setDecoderFilterCallbacks(
      Http::StreamDecoderFilterCallbacks& callbacks) override;

  // Called when the request is completed.
  void log(const Http::HeaderMap* request_headers,
           const Http::HeaderMap* response_headers,
           const Http::HeaderMap* response_trailers,
           const StreamInfo::StreamInfo& stream_info) override;

  // For Handler::CheckDoneCallback, called when callCheck() is done
  void onCheckDone(const ::google::protobuf::util::Status& status);

 private:
  void rejectRequest(Http::Code code, absl::string_view error_msg);

  // The callback funcion.
  Http::StreamDecoderFilterCallbacks* decoder_callbacks_;
  FilterConfigSharedPtr config_;

  // The service control request handler
  std::unique_ptr<Handler> handler_;

  // The state of the request.
  enum State { Init, Calling, Responded, Complete };
  State state_ = Init;
  // Mark if request has been stopped.
  bool stopped_ = false;
};

}  // namespace ServiceControl
}  // namespace HttpFilters
}  // namespace Extensions
}  // namespace Envoy
