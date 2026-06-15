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

// Represents SmartArt shape resource.
type ISmartArt interface {

	// Gets or sets the link to this resource.
	GetSelfUri() IResourceUri
	SetSelfUri(newValue IResourceUri)

	// List of alternate links.
	GetAlternateLinks() []IResourceUri
	SetAlternateLinks(newValue []IResourceUri)

	// Gets or sets the name.
	GetName() string
	SetName(newValue string)

	// Gets or sets the width.
	GetWidth() float64
	SetWidth(newValue float64)

	// Gets or sets the height.
	GetHeight() float64
	SetHeight(newValue float64)

	// Gets or sets the alternative text.
	GetAlternativeText() string
	SetAlternativeText(newValue string)

	// The title of alternative text associated with the shape.
	GetAlternativeTextTitle() string
	SetAlternativeTextTitle(newValue string)

	// Gets or sets a value indicating whether this ShapeBase is hidden.
	GetHidden() *bool
	SetHidden(newValue *bool)

	// Gets or sets 'Mark as decorative' option.
	GetIsDecorative() *bool
	SetIsDecorative(newValue *bool)

	// Gets or sets the X
	GetX() float64
	SetX(newValue float64)

	// Gets or sets the Y.
	GetY() float64
	SetY(newValue float64)

	// Gets z-order position of shape
	GetZOrderPosition() int32
	SetZOrderPosition(newValue int32)

	// Gets or sets the fill format.
	GetFillFormat() IFillFormat
	SetFillFormat(newValue IFillFormat)

	// Gets or sets the effect format.
	GetEffectFormat() IEffectFormat
	SetEffectFormat(newValue IEffectFormat)

	// Gets or sets the 3D format
	GetThreeDFormat() IThreeDFormat
	SetThreeDFormat(newValue IThreeDFormat)

	// Gets or sets the line format.
	GetLineFormat() ILineFormat
	SetLineFormat(newValue ILineFormat)

	// Hyperlink defined for mouse click.
	GetHyperlinkClick() IHyperlink
	SetHyperlinkClick(newValue IHyperlink)

	// Hyperlink defined for mouse over.
	GetHyperlinkMouseOver() IHyperlink
	SetHyperlinkMouseOver(newValue IHyperlink)

	// Shape type.
	GetType() string
	SetType(newValue string)

	// Layout type.
	GetLayout() string
	SetLayout(newValue string)

	// Quick style.
	GetQuickStyle() string
	SetQuickStyle(newValue string)

	// Color style.
	GetColorStyle() string
	SetColorStyle(newValue string)

	// Collection of nodes in SmartArt object.             
	GetNodes() []ISmartArtNode
	SetNodes(newValue []ISmartArtNode)

	// The state of the SmartArt diagram with regard to (left-to-right) LTR or (right-to-left) RTL, if the diagram supports reversal.
	GetIsReversed() *bool
	SetIsReversed(newValue *bool)

	// Default paragraph format applied to all nodes. Write-only: not populated on read since there is no corresponding property at the SmartArt shape level in Aspose.Slides.
	GetDefaultParagraphFormat() IParagraphFormat
	SetDefaultParagraphFormat(newValue IParagraphFormat)
}

type SmartArt struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// Gets or sets the name.
	Name string `json:"Name,omitempty"`

	// Gets or sets the width.
	Width float64 `json:"Width,omitempty"`

	// Gets or sets the height.
	Height float64 `json:"Height,omitempty"`

	// Gets or sets the alternative text.
	AlternativeText string `json:"AlternativeText,omitempty"`

	// The title of alternative text associated with the shape.
	AlternativeTextTitle string `json:"AlternativeTextTitle,omitempty"`

	// Gets or sets a value indicating whether this ShapeBase is hidden.
	Hidden *bool `json:"Hidden"`

	// Gets or sets 'Mark as decorative' option.
	IsDecorative *bool `json:"IsDecorative"`

	// Gets or sets the X
	X float64 `json:"X,omitempty"`

	// Gets or sets the Y.
	Y float64 `json:"Y,omitempty"`

	// Gets z-order position of shape
	ZOrderPosition int32 `json:"ZOrderPosition,omitempty"`

	// Gets or sets the fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Gets or sets the effect format.
	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	// Gets or sets the 3D format
	ThreeDFormat IThreeDFormat `json:"ThreeDFormat,omitempty"`

	// Gets or sets the line format.
	LineFormat ILineFormat `json:"LineFormat,omitempty"`

	// Hyperlink defined for mouse click.
	HyperlinkClick IHyperlink `json:"HyperlinkClick,omitempty"`

	// Hyperlink defined for mouse over.
	HyperlinkMouseOver IHyperlink `json:"HyperlinkMouseOver,omitempty"`

	// Shape type.
	Type_ string `json:"Type"`

	// Layout type.
	Layout string `json:"Layout"`

	// Quick style.
	QuickStyle string `json:"QuickStyle"`

	// Color style.
	ColorStyle string `json:"ColorStyle"`

	// Collection of nodes in SmartArt object.             
	Nodes []ISmartArtNode `json:"Nodes,omitempty"`

	// The state of the SmartArt diagram with regard to (left-to-right) LTR or (right-to-left) RTL, if the diagram supports reversal.
	IsReversed *bool `json:"IsReversed"`

	// Default paragraph format applied to all nodes. Write-only: not populated on read since there is no corresponding property at the SmartArt shape level in Aspose.Slides.
	DefaultParagraphFormat IParagraphFormat `json:"DefaultParagraphFormat,omitempty"`
}

func NewSmartArt() *SmartArt {
	instance := new(SmartArt)
	instance.Type_ = "SmartArt"
	instance.Layout = "AccentProcess"
	instance.QuickStyle = "SimpleFill"
	instance.ColorStyle = "Dark1Outline"
	return instance
}

func (this *SmartArt) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *SmartArt) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *SmartArt) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *SmartArt) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *SmartArt) GetName() string {
	return this.Name
}

func (this *SmartArt) SetName(newValue string) {
	this.Name = newValue
}
func (this *SmartArt) GetWidth() float64 {
	return this.Width
}

func (this *SmartArt) SetWidth(newValue float64) {
	this.Width = newValue
}
func (this *SmartArt) GetHeight() float64 {
	return this.Height
}

func (this *SmartArt) SetHeight(newValue float64) {
	this.Height = newValue
}
func (this *SmartArt) GetAlternativeText() string {
	return this.AlternativeText
}

func (this *SmartArt) SetAlternativeText(newValue string) {
	this.AlternativeText = newValue
}
func (this *SmartArt) GetAlternativeTextTitle() string {
	return this.AlternativeTextTitle
}

func (this *SmartArt) SetAlternativeTextTitle(newValue string) {
	this.AlternativeTextTitle = newValue
}
func (this *SmartArt) GetHidden() *bool {
	return this.Hidden
}

func (this *SmartArt) SetHidden(newValue *bool) {
	this.Hidden = newValue
}
func (this *SmartArt) GetIsDecorative() *bool {
	return this.IsDecorative
}

func (this *SmartArt) SetIsDecorative(newValue *bool) {
	this.IsDecorative = newValue
}
func (this *SmartArt) GetX() float64 {
	return this.X
}

func (this *SmartArt) SetX(newValue float64) {
	this.X = newValue
}
func (this *SmartArt) GetY() float64 {
	return this.Y
}

func (this *SmartArt) SetY(newValue float64) {
	this.Y = newValue
}
func (this *SmartArt) GetZOrderPosition() int32 {
	return this.ZOrderPosition
}

func (this *SmartArt) SetZOrderPosition(newValue int32) {
	this.ZOrderPosition = newValue
}
func (this *SmartArt) GetFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *SmartArt) SetFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *SmartArt) GetEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *SmartArt) SetEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *SmartArt) GetThreeDFormat() IThreeDFormat {
	return this.ThreeDFormat
}

func (this *SmartArt) SetThreeDFormat(newValue IThreeDFormat) {
	this.ThreeDFormat = newValue
}
func (this *SmartArt) GetLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *SmartArt) SetLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *SmartArt) GetHyperlinkClick() IHyperlink {
	return this.HyperlinkClick
}

func (this *SmartArt) SetHyperlinkClick(newValue IHyperlink) {
	this.HyperlinkClick = newValue
}
func (this *SmartArt) GetHyperlinkMouseOver() IHyperlink {
	return this.HyperlinkMouseOver
}

func (this *SmartArt) SetHyperlinkMouseOver(newValue IHyperlink) {
	this.HyperlinkMouseOver = newValue
}
func (this *SmartArt) GetType() string {
	return this.Type_
}

func (this *SmartArt) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *SmartArt) GetLayout() string {
	return this.Layout
}

func (this *SmartArt) SetLayout(newValue string) {
	this.Layout = newValue
}
func (this *SmartArt) GetQuickStyle() string {
	return this.QuickStyle
}

func (this *SmartArt) SetQuickStyle(newValue string) {
	this.QuickStyle = newValue
}
func (this *SmartArt) GetColorStyle() string {
	return this.ColorStyle
}

func (this *SmartArt) SetColorStyle(newValue string) {
	this.ColorStyle = newValue
}
func (this *SmartArt) GetNodes() []ISmartArtNode {
	return this.Nodes
}

func (this *SmartArt) SetNodes(newValue []ISmartArtNode) {
	this.Nodes = newValue
}
func (this *SmartArt) GetIsReversed() *bool {
	return this.IsReversed
}

func (this *SmartArt) SetIsReversed(newValue *bool) {
	this.IsReversed = newValue
}
func (this *SmartArt) GetDefaultParagraphFormat() IParagraphFormat {
	return this.DefaultParagraphFormat
}

func (this *SmartArt) SetDefaultParagraphFormat(newValue IParagraphFormat) {
	this.DefaultParagraphFormat = newValue
}

func (this *SmartArt) UnmarshalJSON(b []byte) error {
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
	
	if valName, ok := GetMapValue(objMap, "name"); ok {
		if valName != nil {
			var valueForName string
			err = json.Unmarshal(*valName, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	
	if valWidth, ok := GetMapValue(objMap, "width"); ok {
		if valWidth != nil {
			var valueForWidth float64
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	
	if valHeight, ok := GetMapValue(objMap, "height"); ok {
		if valHeight != nil {
			var valueForHeight float64
			err = json.Unmarshal(*valHeight, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	
	if valAlternativeText, ok := GetMapValue(objMap, "alternativeText"); ok {
		if valAlternativeText != nil {
			var valueForAlternativeText string
			err = json.Unmarshal(*valAlternativeText, &valueForAlternativeText)
			if err != nil {
				return err
			}
			this.AlternativeText = valueForAlternativeText
		}
	}
	
	if valAlternativeTextTitle, ok := GetMapValue(objMap, "alternativeTextTitle"); ok {
		if valAlternativeTextTitle != nil {
			var valueForAlternativeTextTitle string
			err = json.Unmarshal(*valAlternativeTextTitle, &valueForAlternativeTextTitle)
			if err != nil {
				return err
			}
			this.AlternativeTextTitle = valueForAlternativeTextTitle
		}
	}
	
	if valHidden, ok := GetMapValue(objMap, "hidden"); ok {
		if valHidden != nil {
			var valueForHidden *bool
			err = json.Unmarshal(*valHidden, &valueForHidden)
			if err != nil {
				return err
			}
			this.Hidden = valueForHidden
		}
	}
	
	if valIsDecorative, ok := GetMapValue(objMap, "isDecorative"); ok {
		if valIsDecorative != nil {
			var valueForIsDecorative *bool
			err = json.Unmarshal(*valIsDecorative, &valueForIsDecorative)
			if err != nil {
				return err
			}
			this.IsDecorative = valueForIsDecorative
		}
	}
	
	if valX, ok := GetMapValue(objMap, "x"); ok {
		if valX != nil {
			var valueForX float64
			err = json.Unmarshal(*valX, &valueForX)
			if err != nil {
				return err
			}
			this.X = valueForX
		}
	}
	
	if valY, ok := GetMapValue(objMap, "y"); ok {
		if valY != nil {
			var valueForY float64
			err = json.Unmarshal(*valY, &valueForY)
			if err != nil {
				return err
			}
			this.Y = valueForY
		}
	}
	
	if valZOrderPosition, ok := GetMapValue(objMap, "zOrderPosition"); ok {
		if valZOrderPosition != nil {
			var valueForZOrderPosition int32
			err = json.Unmarshal(*valZOrderPosition, &valueForZOrderPosition)
			if err != nil {
				return err
			}
			this.ZOrderPosition = valueForZOrderPosition
		}
	}
	
	if valFillFormat, ok := GetMapValue(objMap, "fillFormat"); ok {
		if valFillFormat != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormat, &valueForFillFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("FillFormat", *valFillFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valFillFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IFillFormat)
			if ok {
				this.FillFormat = vInterfaceObject
			}
		}
	}
	
	if valEffectFormat, ok := GetMapValue(objMap, "effectFormat"); ok {
		if valEffectFormat != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormat, &valueForEffectFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("EffectFormat", *valEffectFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valEffectFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IEffectFormat)
			if ok {
				this.EffectFormat = vInterfaceObject
			}
		}
	}
	
	if valThreeDFormat, ok := GetMapValue(objMap, "threeDFormat"); ok {
		if valThreeDFormat != nil {
			var valueForThreeDFormat ThreeDFormat
			err = json.Unmarshal(*valThreeDFormat, &valueForThreeDFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ThreeDFormat", *valThreeDFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valThreeDFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IThreeDFormat)
			if ok {
				this.ThreeDFormat = vInterfaceObject
			}
		}
	}
	
	if valLineFormat, ok := GetMapValue(objMap, "lineFormat"); ok {
		if valLineFormat != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormat, &valueForLineFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("LineFormat", *valLineFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valLineFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILineFormat)
			if ok {
				this.LineFormat = vInterfaceObject
			}
		}
	}
	
	if valHyperlinkClick, ok := GetMapValue(objMap, "hyperlinkClick"); ok {
		if valHyperlinkClick != nil {
			var valueForHyperlinkClick Hyperlink
			err = json.Unmarshal(*valHyperlinkClick, &valueForHyperlinkClick)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Hyperlink", *valHyperlinkClick)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valHyperlinkClick, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IHyperlink)
			if ok {
				this.HyperlinkClick = vInterfaceObject
			}
		}
	}
	
	if valHyperlinkMouseOver, ok := GetMapValue(objMap, "hyperlinkMouseOver"); ok {
		if valHyperlinkMouseOver != nil {
			var valueForHyperlinkMouseOver Hyperlink
			err = json.Unmarshal(*valHyperlinkMouseOver, &valueForHyperlinkMouseOver)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Hyperlink", *valHyperlinkMouseOver)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valHyperlinkMouseOver, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IHyperlink)
			if ok {
				this.HyperlinkMouseOver = vInterfaceObject
			}
		}
	}
	this.Type_ = "SmartArt"
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
	this.Layout = "AccentProcess"
	if valLayout, ok := GetMapValue(objMap, "layout"); ok {
		if valLayout != nil {
			var valueForLayout string
			err = json.Unmarshal(*valLayout, &valueForLayout)
			if err != nil {
				var valueForLayoutInt int32
				err = json.Unmarshal(*valLayout, &valueForLayoutInt)
				if err != nil {
					return err
				}
				this.Layout = string(valueForLayoutInt)
			} else {
				this.Layout = valueForLayout
			}
		}
	}
	this.QuickStyle = "SimpleFill"
	if valQuickStyle, ok := GetMapValue(objMap, "quickStyle"); ok {
		if valQuickStyle != nil {
			var valueForQuickStyle string
			err = json.Unmarshal(*valQuickStyle, &valueForQuickStyle)
			if err != nil {
				var valueForQuickStyleInt int32
				err = json.Unmarshal(*valQuickStyle, &valueForQuickStyleInt)
				if err != nil {
					return err
				}
				this.QuickStyle = string(valueForQuickStyleInt)
			} else {
				this.QuickStyle = valueForQuickStyle
			}
		}
	}
	this.ColorStyle = "Dark1Outline"
	if valColorStyle, ok := GetMapValue(objMap, "colorStyle"); ok {
		if valColorStyle != nil {
			var valueForColorStyle string
			err = json.Unmarshal(*valColorStyle, &valueForColorStyle)
			if err != nil {
				var valueForColorStyleInt int32
				err = json.Unmarshal(*valColorStyle, &valueForColorStyleInt)
				if err != nil {
					return err
				}
				this.ColorStyle = string(valueForColorStyleInt)
			} else {
				this.ColorStyle = valueForColorStyle
			}
		}
	}
	
	if valNodes, ok := GetMapValue(objMap, "nodes"); ok {
		if valNodes != nil {
			var valueForNodes []json.RawMessage
			err = json.Unmarshal(*valNodes, &valueForNodes)
			if err != nil {
				return err
			}
			valueForINodes := make([]ISmartArtNode, len(valueForNodes))
			for i, v := range valueForNodes {
				vObject, err := createObjectForType("SmartArtNode", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForINodes[i] = vObject.(ISmartArtNode)
				}
			}
			this.Nodes = valueForINodes
		}
	}
	
	if valIsReversed, ok := GetMapValue(objMap, "isReversed"); ok {
		if valIsReversed != nil {
			var valueForIsReversed *bool
			err = json.Unmarshal(*valIsReversed, &valueForIsReversed)
			if err != nil {
				return err
			}
			this.IsReversed = valueForIsReversed
		}
	}
	
	if valDefaultParagraphFormat, ok := GetMapValue(objMap, "defaultParagraphFormat"); ok {
		if valDefaultParagraphFormat != nil {
			var valueForDefaultParagraphFormat ParagraphFormat
			err = json.Unmarshal(*valDefaultParagraphFormat, &valueForDefaultParagraphFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ParagraphFormat", *valDefaultParagraphFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valDefaultParagraphFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IParagraphFormat)
			if ok {
				this.DefaultParagraphFormat = vInterfaceObject
			}
		}
	}

	return nil
}
