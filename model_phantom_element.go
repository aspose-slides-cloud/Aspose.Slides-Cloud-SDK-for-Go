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

// Specifies an instance of mathematical text that contained within a MathParagraph and starts on its own line.
type IPhantomElement interface {

	// Element type
	GetType() string
	SetType(newValue string)

	// Base element
	GetBase() IMathElement
	SetBase(newValue IMathElement)

	// true if the base element is displayed.
	GetShow() *bool
	SetShow(newValue *bool)

	// true if the the width of the base element should be treated as zero.
	GetZeroWidth() *bool
	SetZeroWidth(newValue *bool)

	// true if the the ascent (height above baseline) of the base element should be treated as zero.
	GetZeroAsc() *bool
	SetZeroAsc(newValue *bool)

	// true if the the descent (depth below baseline) of the base element should be treated as zero.
	GetZeroDesc() *bool
	SetZeroDesc(newValue *bool)

	// true if operators and symbols inside the phantom still affect mathematical spacing around the phantom (as if visible).
	GetTransp() *bool
	SetTransp(newValue *bool)
}

type PhantomElement struct {

	// Element type
	Type_ string `json:"Type"`

	// Base element
	Base IMathElement `json:"Base,omitempty"`

	// true if the base element is displayed.
	Show *bool `json:"Show"`

	// true if the the width of the base element should be treated as zero.
	ZeroWidth *bool `json:"ZeroWidth"`

	// true if the the ascent (height above baseline) of the base element should be treated as zero.
	ZeroAsc *bool `json:"ZeroAsc"`

	// true if the the descent (depth below baseline) of the base element should be treated as zero.
	ZeroDesc *bool `json:"ZeroDesc"`

	// true if operators and symbols inside the phantom still affect mathematical spacing around the phantom (as if visible).
	Transp *bool `json:"Transp"`
}

func NewPhantomElement() *PhantomElement {
	instance := new(PhantomElement)
	instance.Type_ = "Phantom"
	return instance
}

func (this *PhantomElement) GetType() string {
	return this.Type_
}

func (this *PhantomElement) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *PhantomElement) GetBase() IMathElement {
	return this.Base
}

func (this *PhantomElement) SetBase(newValue IMathElement) {
	this.Base = newValue
}
func (this *PhantomElement) GetShow() *bool {
	return this.Show
}

func (this *PhantomElement) SetShow(newValue *bool) {
	this.Show = newValue
}
func (this *PhantomElement) GetZeroWidth() *bool {
	return this.ZeroWidth
}

func (this *PhantomElement) SetZeroWidth(newValue *bool) {
	this.ZeroWidth = newValue
}
func (this *PhantomElement) GetZeroAsc() *bool {
	return this.ZeroAsc
}

func (this *PhantomElement) SetZeroAsc(newValue *bool) {
	this.ZeroAsc = newValue
}
func (this *PhantomElement) GetZeroDesc() *bool {
	return this.ZeroDesc
}

func (this *PhantomElement) SetZeroDesc(newValue *bool) {
	this.ZeroDesc = newValue
}
func (this *PhantomElement) GetTransp() *bool {
	return this.Transp
}

func (this *PhantomElement) SetTransp(newValue *bool) {
	this.Transp = newValue
}

func (this *PhantomElement) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Phantom"
	if valType, ok := GetMapValue(objMap, "type"); ok {
		if valType != nil {
			var valueForType string
			err = json.Unmarshal(*valType, &valueForType)
			if err != nil {
				var valueForTypeInt int32
				err = json.Unmarshal(*valType, &valueForTypeInt)
				if err != nil {
					return err
				}
				this.Type_ = string(valueForTypeInt)
			} else {
				this.Type_ = valueForType
			}
		}
	}
	
	if valBase, ok := GetMapValue(objMap, "base"); ok {
		if valBase != nil {
			var valueForBase MathElement
			err = json.Unmarshal(*valBase, &valueForBase)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("MathElement", *valBase)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valBase, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IMathElement)
			if ok {
				this.Base = vInterfaceObject
			}
		}
	}
	
	if valShow, ok := GetMapValue(objMap, "show"); ok {
		if valShow != nil {
			var valueForShow *bool
			err = json.Unmarshal(*valShow, &valueForShow)
			if err != nil {
				return err
			}
			this.Show = valueForShow
		}
	}
	
	if valZeroWidth, ok := GetMapValue(objMap, "zeroWidth"); ok {
		if valZeroWidth != nil {
			var valueForZeroWidth *bool
			err = json.Unmarshal(*valZeroWidth, &valueForZeroWidth)
			if err != nil {
				return err
			}
			this.ZeroWidth = valueForZeroWidth
		}
	}
	
	if valZeroAsc, ok := GetMapValue(objMap, "zeroAsc"); ok {
		if valZeroAsc != nil {
			var valueForZeroAsc *bool
			err = json.Unmarshal(*valZeroAsc, &valueForZeroAsc)
			if err != nil {
				return err
			}
			this.ZeroAsc = valueForZeroAsc
		}
	}
	
	if valZeroDesc, ok := GetMapValue(objMap, "zeroDesc"); ok {
		if valZeroDesc != nil {
			var valueForZeroDesc *bool
			err = json.Unmarshal(*valZeroDesc, &valueForZeroDesc)
			if err != nil {
				return err
			}
			this.ZeroDesc = valueForZeroDesc
		}
	}
	
	if valTransp, ok := GetMapValue(objMap, "transp"); ok {
		if valTransp != nil {
			var valueForTransp *bool
			err = json.Unmarshal(*valTransp, &valueForTransp)
			if err != nil {
				return err
			}
			this.Transp = valueForTransp
		}
	}

	return nil
}
