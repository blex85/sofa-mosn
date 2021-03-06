/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package types

//    The bunch of interfaces are used to print the access log in format designed by users.
//    Access log format consists of three parts, which are "RequestInfoFormat", "RequestHeaderFormat"
//    and "ResponseHeaderFormat", also you can get details by reading "AccessLogDetails.md".

// access log
type AccessLog interface {
	// log the access info
	// "reqHeaders" contains the request header's information, "respHeader" contains the response header's information
	// and by "requestInfo" you can get some request information
	Log(reqHeaders map[string]string, respHeaders map[string]string, requestInfo RequestInfo)
}

// filter of access log to do some filters to access log info
type AccessLogFilter interface {
	// decision about how to filter the request headers and requestInfo
	Decide(reqHeaders map[string]string, requestInfo RequestInfo) bool
}

// access log formatter
type AccessLogFormatter interface {
	// format the request headers, response headers and request info to string for printing according to log formatter
	Format(reqHeaders map[string]string, respHeaders map[string]string, requestInfo RequestInfo) string
}

//some const defined to identify how to get request info's content
const (
	// identification of request's arriving time
	LogStartTime string = "StartTime"
	// identification of duration between request arriving and request resend to upstream
	LogRequestReceivedDuration string = "RequestReceivedDuration"
	// identification of duration between request arriving and response sending
	LogResponseReceivedDuration string = "ResponseReceivedDuration"
	// identification of bytes sent
	LogBytesSent string = "BytesSent"
	// identification of bytes received
	LogBytesReceived string = "BytesReceived"
	// identification of request's protocol type
	LogProtocol string = "Protocol"
	// identification of request's response code
	LogResponseCode string = "ResponseCode"
	// identification of duration since request's starting time
	LogDuration string = "Duration"
	// identification of request's response flag
	LogResponseFlag string = "ResponseFlag"
	// identification of upstream's local address
	LogUpstreamLocalAddress string = "UpstreamLocalAddress"
	// identification of downstream's local address
	LogDownstreamLocalAddress string = "DownstreamLocalAddress"
	// identification of downstream's remote address
	LogDownstreamRemoteAddress string = "DownstreamRemoteAddress"
	// identification of host selected
	LogUpstreamHostSelectedGetter string = "UpstreamHostSelected"
)

const (
	// Prefix of request header's formatter
	ReqHeaderPrefix string = "REQ."
	// Prefix of response header's formatter
	RespHeaderPrefix string = "RESP."
)

const (
	// Default Access Log Format, for more details please read "AccessLogDetails.md"
	DefaultAccessLogFormat = "%StartTime% %RequestReceivedDuration% %ResponseReceivedDuration% %BytesSent%" + " " +
		"%BytesReceived% %Protocol% %ResponseCode% %Duration% %ResponseFlag% %ResponseCode% %UpstreamLocalAddress%" + " " +
		"%DownstreamLocalAddress% %DownstreamRemoteAddress% %UpstreamHostSelected%"
)
