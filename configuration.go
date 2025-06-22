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
	"net/http"
)

const ContextOAuth2 		int = 1
const ContextBasicAuth 		int = 2
const ContextAccessToken 	int = 3
const ContextAPIKey 		int = 4 

type BasicAuth struct {
	UserName      string            `json:"userName,omitempty"`
	Password      string            `json:"password,omitempty"`	
}

type APIKey struct {
	Key 	string
	Prefix	string
}

type Configuration struct {
	BasePath              string		`json:"BaseUrl,omitempty"`
	AsyncBasePath         string            `json:"AsyncBaseUrl,omitempty"`
	AuthBasePath          string            `json:"AuthBaseUrl,omitempty"`
	Version               string            `json:"Version,omitempty"`
	Host                  string            `json:"host,omitempty"`
	AppSid                string            `json:"AppSid,omitempty"`
	AppKey                string            `json:"AppKey,omitempty"`
	OAuthToken            string            `json:"OAuthToken,omitempty"`
	Scheme                string            `json:"scheme,omitempty"`
	ApiVersion            string            `json:"ApiVersion,omitempty"`
	Debug                 bool 	        `json:"Debug,omitempty"`
	AllowInsecureRequests bool 	        `json:"AllowInsecureRequests,omitempty"`
	Timeout               int 	        `json:"Timeout,omitempty"`
	CustomHeaders         map[string]string	`json:"CustomHeaders,omitempty"`
	HTTPClient            *http.Client
}

func NewConfiguration() *Configuration {
	cfg := &Configuration{
		BasePath:      "https://api.aspose.cloud",
		AsyncBasePath:  "https://api.aspose.cloud",
		AuthBasePath:  "https://api.aspose.cloud",
		AppSid:        "",
		AppKey:        "",
		Version:       "v3.0",
		ApiVersion:    "25.6.0",
		CustomHeaders: make(map[string]string),
	}
	return cfg
}

func (c *Configuration) GetApiUrl() string {
	return c.BasePath + "/" + c.Version
}
