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
	"encoding/json"
)

// Represents list of Links to Paragraphs resources
type ICaptionTrack interface {

	// Gets or sets the link to this resource.
	GetSelfUri() IResourceUri
	SetSelfUri(newValue IResourceUri)

	// List of alternate links.
	GetAlternateLinks() []IResourceUri
	SetAlternateLinks(newValue []IResourceUri)

	// Caption ID.
	GetCaptionId() string
	SetCaptionId(newValue string)

	// Label.
	GetLabel() string
	SetLabel(newValue string)

	// Caption track data as string.
	GetDataAsString() string
	SetDataAsString(newValue string)
}

type CaptionTrack struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// Caption ID.
	CaptionId string `json:"CaptionId"`

	// Label.
	Label string `json:"Label,omitempty"`

	// Caption track data as string.
	DataAsString string `json:"DataAsString,omitempty"`
}

func NewCaptionTrack() *CaptionTrack {
	instance := new(CaptionTrack)
	return instance
}

func (this *CaptionTrack) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *CaptionTrack) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *CaptionTrack) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *CaptionTrack) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *CaptionTrack) GetCaptionId() string {
	return this.CaptionId
}

func (this *CaptionTrack) SetCaptionId(newValue string) {
	this.CaptionId = newValue
}
func (this *CaptionTrack) GetLabel() string {
	return this.Label
}

func (this *CaptionTrack) SetLabel(newValue string) {
	this.Label = newValue
}
func (this *CaptionTrack) GetDataAsString() string {
	return this.DataAsString
}

func (this *CaptionTrack) SetDataAsString(newValue string) {
	this.DataAsString = newValue
}

func (this *CaptionTrack) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valSelfUri, ok := GetMapValue(objMap, "selfUri"); ok {
		if valSelfUri != nil {
			var valueForSelfUri ResourceUri
			err = json.Unmarshal(*valSelfUri, &valueForSelfUri)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valSelfUri)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSelfUri, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.SelfUri = vInterfaceObject
			}
		}
	}
	
	if valAlternateLinks, ok := GetMapValue(objMap, "alternateLinks"); ok {
		if valAlternateLinks != nil {
			var valueForAlternateLinks []json.RawMessage
			err = json.Unmarshal(*valAlternateLinks, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			valueForIAlternateLinks := make([]IResourceUri, len(valueForAlternateLinks))
			for i, v := range valueForAlternateLinks {
				vObject, err := createObjectForType("ResourceUri", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIAlternateLinks[i] = vObject.(IResourceUri)
				}
			}
			this.AlternateLinks = valueForIAlternateLinks
		}
	}
	
	if valCaptionId, ok := GetMapValue(objMap, "captionId"); ok {
		if valCaptionId != nil {
			var valueForCaptionId string
			err = json.Unmarshal(*valCaptionId, &valueForCaptionId)
			if err != nil {
				return err
			}
			this.CaptionId = valueForCaptionId
		}
	}
	
	if valLabel, ok := GetMapValue(objMap, "label"); ok {
		if valLabel != nil {
			var valueForLabel string
			err = json.Unmarshal(*valLabel, &valueForLabel)
			if err != nil {
				return err
			}
			this.Label = valueForLabel
		}
	}
	
	if valDataAsString, ok := GetMapValue(objMap, "dataAsString"); ok {
		if valDataAsString != nil {
			var valueForDataAsString string
			err = json.Unmarshal(*valDataAsString, &valueForDataAsString)
			if err != nil {
				return err
			}
			this.DataAsString = valueForDataAsString
		}
	}

	return nil
}
