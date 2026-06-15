/*
 * --------------------------------------------------------------------------------------------------------------------
 * <copyright company="Aspose">
 *   Copyright (c) 2018 Aspose.Slides for Cloud
 * </copyright>
 * <summary>
 *   Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 * 
 *  The above copyright notice and this permission notice shall be included in all
 *  copies or substantial portions of the Software.
 * 
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *  SOFTWARE.
 * </summary>
 * --------------------------------------------------------------------------------------------------------------------
 */

package asposeslidescloud

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"context"
	"net/http"
	"net/url"
	"time"
	"os"
	"reflect"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
	"strconv"
)

var (
	jsonCheck = regexp.MustCompile("(?i:[application|text]/json)")
	xmlCheck = regexp.MustCompile("(?i:[application|text]/xml)")
)

// APIClient manages communication with the Aspose.Slides for Cloud API Reference API v1.1
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg 	*Configuration
	common 	service 		// Reuse a single struct instead of allocating one for each service on the heap.

	 // API Services
	SlidesApi		*SlidesApiService
	SlidesAsyncApi		*SlidesAsyncApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		if (cfg.AllowInsecureRequests) {
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			http.DefaultClient.Transport = tr
		}
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.SlidesApi = (*SlidesApiService)(&c.common)
	c.SlidesAsyncApi = (*SlidesAsyncApiService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}


// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insenstive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string
	delimiter = ","

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	}

        if reflect.ValueOf(obj).Kind() == reflect.Ptr {
		return fmt.Sprintf("%v", reflect.ValueOf(obj).Elem())
	}
	return fmt.Sprintf("%v", obj)
}

// callAPI do the request. 
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	 return c.cfg.HTTPClient.Do(request)
}

// Change base path to allow switching to mocks
func (c *APIClient) ChangeBasePath (path string) {
	c.cfg.BasePath = path
}

/* SlidesApiService Get API info.
 @return ApiInfo*/
func (c *APIClient) makeRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	files [][]byte) (response *http.Response, responseBytes []byte, err error) {
	response, responseBytes, needRepeat, err := c.makeRequestWithAuthCheck(ctx, path, method, postBody, headerParams, queryParams, files)
	if needRepeat {
		c.cfg.OAuthToken = ""
		response, responseBytes, _, err = c.makeRequestWithAuthCheck(ctx, path, method, postBody, headerParams, queryParams, files)
	}
	return response, responseBytes, err
}

/* SlidesApiService Get API info.
 @return ApiInfo*/
func (c *APIClient) makeRequestWithAuthCheck(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	files [][]byte) (response *http.Response, responseBytes []byte, needRepeat bool, err error) {
	needRepeat = len(c.cfg.OAuthToken) > 0
	r, resp, err := c.prepareRequest(nil, path, method, postBody, headerParams, queryParams, files)
	if err != nil {
		return resp, nil, false, err
	}
	if c.cfg.Debug {
		c.log(fmt.Sprintf("-->: %v %v %v\n", r.Method, r.URL, r.Header))
		if !strings.Contains(r.Header.Get("Content-Type"), "multipart") && r.Body != nil {
			c.log(fmt.Sprintf("-->BODY: %v\n", r.Body))
		}
	}
	localVarHttpResponse, err := c.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, nil, false, err
	}
	if c.cfg.Debug {
		c.log(fmt.Sprintf("<--: %v\n", localVarHttpResponse))
	}
	defer localVarHttpResponse.Body.Close()
        responseLength := 0
	responseBytes, err = ioutil.ReadAll(localVarHttpResponse.Body)
	if err == nil && localVarHttpResponse != nil {
		if c.cfg.Debug && (strings.Contains(localVarHttpResponse.Header.Get("Content-Type"), "text") || strings.Contains(localVarHttpResponse.Header.Get("Content-Type"), "json")) {
			c.log(fmt.Sprintf("<--BODY: %v\n", string(responseBytes)))
		}
		responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1)
	        responseLength = len(responseBytes)
	}
	needRepeat = needRepeat &&
		(localVarHttpResponse.StatusCode == 401 ||
			(localVarHttpResponse.StatusCode == 500 && responseLength == 0))
	return localVarHttpResponse, responseBytes, needRepeat, err
}

// prepareRequest build the request
func (c *APIClient) prepareRequest (
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	files [][]byte) (localVarRequest *http.Request, response *http.Response, err error) {

	var body *bytes.Buffer
        partCount := len(files)
	if postBody != nil {
		partCount += 1
	}
	if len(files) > 0 {
		contentType := headerParams["Content-Type"]
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)
		if postBody != nil {
			bodyBuf, _, err := setBody(postBody, contentType)
			if err != nil {
				return nil, nil, err
			}
			w.WriteField("pipeline", string(bodyBuf.Bytes()))
		}
		for i, file := range files {
			part, err := w.CreateFormFile(fmt.Sprintf("file%d", i + 1), fmt.Sprintf("file%d", i + 1))
			if err != nil {
				return nil, nil, err
			}
			_, err = part.Write(file)
			if err != nil {
				return nil, nil, err
			}
		}
		w.Close()
                contentType = w.FormDataContentType()
		headerParams["Content-Type"] = contentType
	} else if partCount == 1 {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
		}
		if postBody != nil {
			body, contentType, err = setBody(postBody, contentType)
		} else {
			body, contentType, err = setBody(files[0], contentType)
		}
		if err != nil {
			return nil, nil, err
		}
		headerParams["Content-Type"] = contentType
	}

	// Setup path and query paramters
	url, err := url.Parse(path)
	if err != nil {
		return nil, nil, err
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, nil, err
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		localVarRequest.Host = c.cfg.Host
	}

	authResponse, err := c.prepareRequestHeader(localVarRequest, headerParams)
	if err != nil {
		return localVarRequest, authResponse, err
	}
	
	return localVarRequest, nil, nil
}

// prepareRequest build the request
func (c *APIClient) prepareRequestHeader(localVarRequest *http.Request, headerParams map[string]string) (*http.Response, error) {
	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("x-aspose-client", fmt.Sprintf("go sdk v%v", c.cfg.ApiVersion))
	if c.cfg.Timeout > 0 {
		localVarRequest.Header.Add("x-aspose-timeout", strconv.FormatInt(int64(c.cfg.Timeout), 10))
	}
	for headerKey, headerValue := range c.cfg.CustomHeaders {
		localVarRequest.Header.Add(headerKey, headerValue)
	}
	if (len(c.cfg.AppSid) > 0 || len(c.cfg.OAuthToken) > 0) {
		if (len(c.cfg.OAuthToken) == 0) {
			oauthRequest, err := http.NewRequest(
				"POST",
				c.cfg.GetAuthBasePath() + "/connect/token",
				strings.NewReader("grant_type=client_credentials&client_id=" + c.cfg.AppSid + "&client_secret=" + c.cfg.AppKey))
			if (err != nil) {
				return nil, err
			}
			oauthRequest.Header.Set("Content-type", "application/x-www-form-urlencoded")
			oauthResponse, err := c.callAPI(oauthRequest)
			if (err != nil) {
        	                oauthResponse.StatusCode = 401
				return oauthResponse, err
			}
			defer oauthResponse.Body.Close()
			if oauthResponse.StatusCode >= 300 {
                        	oauthResponse.StatusCode = 401
				return oauthResponse, errors.New("Authentication error.")
 			}
			var token OAuthResponse
			if err = json.NewDecoder(oauthResponse.Body).Decode(&token); err != nil {
        	                oauthResponse.StatusCode = 401
				return oauthResponse, err
			}
			c.cfg.OAuthToken = token.AccessToken
		}
		localVarRequest.Header.Add("Authorization", "Bearer " + c.cfg.OAuthToken)
	}
	return nil, nil
}

func (c *APIClient) log(message string) {
	if c.cfg.Logger != nil {
		c.cfg.Logger(message)
	} else {
		fmt.Print(message)
	}
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) (error) {
	return fmt.Errorf(format, a...)
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, resultContentType string, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}
        if reflect.ValueOf(body).Kind() == reflect.Ptr && reflect.ValueOf(body).Elem().IsNil() {
		return bodyBuf, contentType, nil
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if b, ok := body.(*[]byte); ok {
		_, err = bodyBuf.Write(*b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if xmlCheck.MatchString(contentType) {
		xml.NewEncoder(bodyBuf).Encode(body)
	} else {
		err = json.NewEncoder(bodyBuf).Encode(body)
		if err == nil {
			contentType = "application/json"
		}
	}

	if err != nil {
		return nil, contentType, err
	}

	return bodyBuf, contentType, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()
	
	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}


// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) (time.Time) {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)
	
	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		}
		expires = now.Add(lifetime)
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) (int) {
	return utf8.RuneCountInString(s)
}

func processFileResponse(reader io.Reader) (*os.File, error) {
	out, err := ioutil.TempFile("", "*")
	if err != nil {
		return out, err
	}
	defer out.Close()
	_, err = io.Copy(out, reader)
	if err != nil {
		return out, err
	}
	return out, nil
}

func GetMapValue(objMap map[string]*json.RawMessage, key string) (*json.RawMessage, bool) {
	key = firstToLower(key)
	if value, ok := objMap[key]; ok {
		return value, true
	}
	key = firstToUpper(key)
	if value, ok := objMap[key]; ok {
		return value, true
	}
	return nil, false
}

func firstToLower(s string) string {
    r, size := utf8.DecodeRuneInString(s)
    if r == utf8.RuneError && size <= 1 {
        return s
    }
    lc := unicode.ToLower(r)
    if r == lc {
        return s
    }
    return string(lc) + s[size:]
}

func firstToUpper(s string) string {
    r, size := utf8.DecodeRuneInString(s)
    if r == utf8.RuneError && size <= 1 {
        return s
    }
    lc := unicode.ToUpper(r)
    if r == lc {
        return s
    }
    return string(lc) + s[size:]
}