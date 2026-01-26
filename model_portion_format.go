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

// Represents portion format.
type IPortionFormat interface {

	// True for bold font.
	GetFontBold() string
	SetFontBold(newValue string)

	// True for italic font.
	GetFontItalic() string
	SetFontItalic(newValue string)

	// Text underline type.
	GetFontUnderline() string
	SetFontUnderline(newValue string)

	// Text strikethrough type.
	GetStrikethroughType() string
	SetStrikethroughType(newValue string)

	// Text capitalization type.
	GetTextCapType() string
	SetTextCapType(newValue string)

	// Superscript or subscript of the text.
	GetEscapement() float64
	SetEscapement(newValue float64)

	// Intercharacter spacing increment.
	GetSpacing() float64
	SetSpacing(newValue float64)

	// Font color.
	GetFontColor() string
	SetFontColor(newValue string)

	// Highlight color.
	GetHighlightColor() string
	SetHighlightColor(newValue string)

	// Font height.
	GetFontHeight() float64
	SetFontHeight(newValue float64)

	// True to normalize the text.
	GetNormaliseHeight() string
	SetNormaliseHeight(newValue string)

	// True if the text proof should be disabled.
	GetProofDisabled() string
	SetProofDisabled(newValue string)

	// True if smart tag should be cleaned.
	GetSmartTagClean() *bool
	SetSmartTagClean(newValue *bool)

	// Minimal font size for kerning.
	GetKerningMinimalSize() float64
	SetKerningMinimalSize(newValue float64)

	// True if numbers should ignore East-Asian specific vertical text layout.
	GetKumimoji() string
	SetKumimoji(newValue string)

	// Proving language ID.
	GetLanguageId() string
	SetLanguageId(newValue string)

	// Alternative proving language ID.
	GetAlternativeLanguageId() string
	SetAlternativeLanguageId(newValue string)

	// True if underline style has own FillFormat properties.
	GetIsHardUnderlineFill() string
	SetIsHardUnderlineFill(newValue string)

	// True if underline style has own LineFormat properties.
	GetIsHardUnderlineLine() string
	SetIsHardUnderlineLine(newValue string)

	// Fill format.
	GetFillFormat() IFillFormat
	SetFillFormat(newValue IFillFormat)

	// Effect format.
	GetEffectFormat() IEffectFormat
	SetEffectFormat(newValue IEffectFormat)

	// Line format.
	GetLineFormat() ILineFormat
	SetLineFormat(newValue ILineFormat)

	// Underline fill format.
	GetUnderlineFillFormat() IFillFormat
	SetUnderlineFillFormat(newValue IFillFormat)

	// Underline line format.
	GetUnderlineLineFormat() ILineFormat
	SetUnderlineLineFormat(newValue ILineFormat)

	// Hyperlink defined for mouse click.
	GetHyperlinkClick() IHyperlink
	SetHyperlinkClick(newValue IHyperlink)

	// Hyperlink defined for mouse over.
	GetHyperlinkMouseOver() IHyperlink
	SetHyperlinkMouseOver(newValue IHyperlink)

	// Returns or sets the Latin font info.
	GetLatinFont() string
	SetLatinFont(newValue string)

	// Returns or sets the East Asian font info.
	GetEastAsianFont() string
	SetEastAsianFont(newValue string)

	// Returns or sets the complex script font info.
	GetComplexScriptFont() string
	SetComplexScriptFont(newValue string)

	// true to enable spell checking for the portion.
	GetSpellCheck() *bool
	SetSpellCheck(newValue *bool)
}

type PortionFormat struct {

	// True for bold font.
	FontBold string `json:"FontBold,omitempty"`

	// True for italic font.
	FontItalic string `json:"FontItalic,omitempty"`

	// Text underline type.
	FontUnderline string `json:"FontUnderline,omitempty"`

	// Text strikethrough type.
	StrikethroughType string `json:"StrikethroughType,omitempty"`

	// Text capitalization type.
	TextCapType string `json:"TextCapType,omitempty"`

	// Superscript or subscript of the text.
	Escapement float64 `json:"Escapement,omitempty"`

	// Intercharacter spacing increment.
	Spacing float64 `json:"Spacing,omitempty"`

	// Font color.
	FontColor string `json:"FontColor,omitempty"`

	// Highlight color.
	HighlightColor string `json:"HighlightColor,omitempty"`

	// Font height.
	FontHeight float64 `json:"FontHeight,omitempty"`

	// True to normalize the text.
	NormaliseHeight string `json:"NormaliseHeight,omitempty"`

	// True if the text proof should be disabled.
	ProofDisabled string `json:"ProofDisabled,omitempty"`

	// True if smart tag should be cleaned.
	SmartTagClean *bool `json:"SmartTagClean"`

	// Minimal font size for kerning.
	KerningMinimalSize float64 `json:"KerningMinimalSize,omitempty"`

	// True if numbers should ignore East-Asian specific vertical text layout.
	Kumimoji string `json:"Kumimoji,omitempty"`

	// Proving language ID.
	LanguageId string `json:"LanguageId,omitempty"`

	// Alternative proving language ID.
	AlternativeLanguageId string `json:"AlternativeLanguageId,omitempty"`

	// True if underline style has own FillFormat properties.
	IsHardUnderlineFill string `json:"IsHardUnderlineFill,omitempty"`

	// True if underline style has own LineFormat properties.
	IsHardUnderlineLine string `json:"IsHardUnderlineLine,omitempty"`

	// Fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Effect format.
	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	// Line format.
	LineFormat ILineFormat `json:"LineFormat,omitempty"`

	// Underline fill format.
	UnderlineFillFormat IFillFormat `json:"UnderlineFillFormat,omitempty"`

	// Underline line format.
	UnderlineLineFormat ILineFormat `json:"UnderlineLineFormat,omitempty"`

	// Hyperlink defined for mouse click.
	HyperlinkClick IHyperlink `json:"HyperlinkClick,omitempty"`

	// Hyperlink defined for mouse over.
	HyperlinkMouseOver IHyperlink `json:"HyperlinkMouseOver,omitempty"`

	// Returns or sets the Latin font info.
	LatinFont string `json:"LatinFont,omitempty"`

	// Returns or sets the East Asian font info.
	EastAsianFont string `json:"EastAsianFont,omitempty"`

	// Returns or sets the complex script font info.
	ComplexScriptFont string `json:"ComplexScriptFont,omitempty"`

	// true to enable spell checking for the portion.
	SpellCheck *bool `json:"SpellCheck"`
}

func NewPortionFormat() *PortionFormat {
	instance := new(PortionFormat)
	return instance
}

func (this *PortionFormat) GetFontBold() string {
	return this.FontBold
}

func (this *PortionFormat) SetFontBold(newValue string) {
	this.FontBold = newValue
}
func (this *PortionFormat) GetFontItalic() string {
	return this.FontItalic
}

func (this *PortionFormat) SetFontItalic(newValue string) {
	this.FontItalic = newValue
}
func (this *PortionFormat) GetFontUnderline() string {
	return this.FontUnderline
}

func (this *PortionFormat) SetFontUnderline(newValue string) {
	this.FontUnderline = newValue
}
func (this *PortionFormat) GetStrikethroughType() string {
	return this.StrikethroughType
}

func (this *PortionFormat) SetStrikethroughType(newValue string) {
	this.StrikethroughType = newValue
}
func (this *PortionFormat) GetTextCapType() string {
	return this.TextCapType
}

func (this *PortionFormat) SetTextCapType(newValue string) {
	this.TextCapType = newValue
}
func (this *PortionFormat) GetEscapement() float64 {
	return this.Escapement
}

func (this *PortionFormat) SetEscapement(newValue float64) {
	this.Escapement = newValue
}
func (this *PortionFormat) GetSpacing() float64 {
	return this.Spacing
}

func (this *PortionFormat) SetSpacing(newValue float64) {
	this.Spacing = newValue
}
func (this *PortionFormat) GetFontColor() string {
	return this.FontColor
}

func (this *PortionFormat) SetFontColor(newValue string) {
	this.FontColor = newValue
}
func (this *PortionFormat) GetHighlightColor() string {
	return this.HighlightColor
}

func (this *PortionFormat) SetHighlightColor(newValue string) {
	this.HighlightColor = newValue
}
func (this *PortionFormat) GetFontHeight() float64 {
	return this.FontHeight
}

func (this *PortionFormat) SetFontHeight(newValue float64) {
	this.FontHeight = newValue
}
func (this *PortionFormat) GetNormaliseHeight() string {
	return this.NormaliseHeight
}

func (this *PortionFormat) SetNormaliseHeight(newValue string) {
	this.NormaliseHeight = newValue
}
func (this *PortionFormat) GetProofDisabled() string {
	return this.ProofDisabled
}

func (this *PortionFormat) SetProofDisabled(newValue string) {
	this.ProofDisabled = newValue
}
func (this *PortionFormat) GetSmartTagClean() *bool {
	return this.SmartTagClean
}

func (this *PortionFormat) SetSmartTagClean(newValue *bool) {
	this.SmartTagClean = newValue
}
func (this *PortionFormat) GetKerningMinimalSize() float64 {
	return this.KerningMinimalSize
}

func (this *PortionFormat) SetKerningMinimalSize(newValue float64) {
	this.KerningMinimalSize = newValue
}
func (this *PortionFormat) GetKumimoji() string {
	return this.Kumimoji
}

func (this *PortionFormat) SetKumimoji(newValue string) {
	this.Kumimoji = newValue
}
func (this *PortionFormat) GetLanguageId() string {
	return this.LanguageId
}

func (this *PortionFormat) SetLanguageId(newValue string) {
	this.LanguageId = newValue
}
func (this *PortionFormat) GetAlternativeLanguageId() string {
	return this.AlternativeLanguageId
}

func (this *PortionFormat) SetAlternativeLanguageId(newValue string) {
	this.AlternativeLanguageId = newValue
}
func (this *PortionFormat) GetIsHardUnderlineFill() string {
	return this.IsHardUnderlineFill
}

func (this *PortionFormat) SetIsHardUnderlineFill(newValue string) {
	this.IsHardUnderlineFill = newValue
}
func (this *PortionFormat) GetIsHardUnderlineLine() string {
	return this.IsHardUnderlineLine
}

func (this *PortionFormat) SetIsHardUnderlineLine(newValue string) {
	this.IsHardUnderlineLine = newValue
}
func (this *PortionFormat) GetFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *PortionFormat) SetFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *PortionFormat) GetEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *PortionFormat) SetEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *PortionFormat) GetLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *PortionFormat) SetLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *PortionFormat) GetUnderlineFillFormat() IFillFormat {
	return this.UnderlineFillFormat
}

func (this *PortionFormat) SetUnderlineFillFormat(newValue IFillFormat) {
	this.UnderlineFillFormat = newValue
}
func (this *PortionFormat) GetUnderlineLineFormat() ILineFormat {
	return this.UnderlineLineFormat
}

func (this *PortionFormat) SetUnderlineLineFormat(newValue ILineFormat) {
	this.UnderlineLineFormat = newValue
}
func (this *PortionFormat) GetHyperlinkClick() IHyperlink {
	return this.HyperlinkClick
}

func (this *PortionFormat) SetHyperlinkClick(newValue IHyperlink) {
	this.HyperlinkClick = newValue
}
func (this *PortionFormat) GetHyperlinkMouseOver() IHyperlink {
	return this.HyperlinkMouseOver
}

func (this *PortionFormat) SetHyperlinkMouseOver(newValue IHyperlink) {
	this.HyperlinkMouseOver = newValue
}
func (this *PortionFormat) GetLatinFont() string {
	return this.LatinFont
}

func (this *PortionFormat) SetLatinFont(newValue string) {
	this.LatinFont = newValue
}
func (this *PortionFormat) GetEastAsianFont() string {
	return this.EastAsianFont
}

func (this *PortionFormat) SetEastAsianFont(newValue string) {
	this.EastAsianFont = newValue
}
func (this *PortionFormat) GetComplexScriptFont() string {
	return this.ComplexScriptFont
}

func (this *PortionFormat) SetComplexScriptFont(newValue string) {
	this.ComplexScriptFont = newValue
}
func (this *PortionFormat) GetSpellCheck() *bool {
	return this.SpellCheck
}

func (this *PortionFormat) SetSpellCheck(newValue *bool) {
	this.SpellCheck = newValue
}

func (this *PortionFormat) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valFontBold, ok := GetMapValue(objMap, "fontBold"); ok {
		if valFontBold != nil {
			var valueForFontBold string
			err = json.Unmarshal(*valFontBold, &valueForFontBold)
			if err != nil {
				var valueForFontBoldInt int32
				err = json.Unmarshal(*valFontBold, &valueForFontBoldInt)
				if err != nil {
					return err
				}
				this.FontBold = string(valueForFontBoldInt)
			} else {
				this.FontBold = valueForFontBold
			}
		}
	}
	
	if valFontItalic, ok := GetMapValue(objMap, "fontItalic"); ok {
		if valFontItalic != nil {
			var valueForFontItalic string
			err = json.Unmarshal(*valFontItalic, &valueForFontItalic)
			if err != nil {
				var valueForFontItalicInt int32
				err = json.Unmarshal(*valFontItalic, &valueForFontItalicInt)
				if err != nil {
					return err
				}
				this.FontItalic = string(valueForFontItalicInt)
			} else {
				this.FontItalic = valueForFontItalic
			}
		}
	}
	
	if valFontUnderline, ok := GetMapValue(objMap, "fontUnderline"); ok {
		if valFontUnderline != nil {
			var valueForFontUnderline string
			err = json.Unmarshal(*valFontUnderline, &valueForFontUnderline)
			if err != nil {
				var valueForFontUnderlineInt int32
				err = json.Unmarshal(*valFontUnderline, &valueForFontUnderlineInt)
				if err != nil {
					return err
				}
				this.FontUnderline = string(valueForFontUnderlineInt)
			} else {
				this.FontUnderline = valueForFontUnderline
			}
		}
	}
	
	if valStrikethroughType, ok := GetMapValue(objMap, "strikethroughType"); ok {
		if valStrikethroughType != nil {
			var valueForStrikethroughType string
			err = json.Unmarshal(*valStrikethroughType, &valueForStrikethroughType)
			if err != nil {
				var valueForStrikethroughTypeInt int32
				err = json.Unmarshal(*valStrikethroughType, &valueForStrikethroughTypeInt)
				if err != nil {
					return err
				}
				this.StrikethroughType = string(valueForStrikethroughTypeInt)
			} else {
				this.StrikethroughType = valueForStrikethroughType
			}
		}
	}
	
	if valTextCapType, ok := GetMapValue(objMap, "textCapType"); ok {
		if valTextCapType != nil {
			var valueForTextCapType string
			err = json.Unmarshal(*valTextCapType, &valueForTextCapType)
			if err != nil {
				var valueForTextCapTypeInt int32
				err = json.Unmarshal(*valTextCapType, &valueForTextCapTypeInt)
				if err != nil {
					return err
				}
				this.TextCapType = string(valueForTextCapTypeInt)
			} else {
				this.TextCapType = valueForTextCapType
			}
		}
	}
	
	if valEscapement, ok := GetMapValue(objMap, "escapement"); ok {
		if valEscapement != nil {
			var valueForEscapement float64
			err = json.Unmarshal(*valEscapement, &valueForEscapement)
			if err != nil {
				return err
			}
			this.Escapement = valueForEscapement
		}
	}
	
	if valSpacing, ok := GetMapValue(objMap, "spacing"); ok {
		if valSpacing != nil {
			var valueForSpacing float64
			err = json.Unmarshal(*valSpacing, &valueForSpacing)
			if err != nil {
				return err
			}
			this.Spacing = valueForSpacing
		}
	}
	
	if valFontColor, ok := GetMapValue(objMap, "fontColor"); ok {
		if valFontColor != nil {
			var valueForFontColor string
			err = json.Unmarshal(*valFontColor, &valueForFontColor)
			if err != nil {
				return err
			}
			this.FontColor = valueForFontColor
		}
	}
	
	if valHighlightColor, ok := GetMapValue(objMap, "highlightColor"); ok {
		if valHighlightColor != nil {
			var valueForHighlightColor string
			err = json.Unmarshal(*valHighlightColor, &valueForHighlightColor)
			if err != nil {
				return err
			}
			this.HighlightColor = valueForHighlightColor
		}
	}
	
	if valFontHeight, ok := GetMapValue(objMap, "fontHeight"); ok {
		if valFontHeight != nil {
			var valueForFontHeight float64
			err = json.Unmarshal(*valFontHeight, &valueForFontHeight)
			if err != nil {
				return err
			}
			this.FontHeight = valueForFontHeight
		}
	}
	
	if valNormaliseHeight, ok := GetMapValue(objMap, "normaliseHeight"); ok {
		if valNormaliseHeight != nil {
			var valueForNormaliseHeight string
			err = json.Unmarshal(*valNormaliseHeight, &valueForNormaliseHeight)
			if err != nil {
				var valueForNormaliseHeightInt int32
				err = json.Unmarshal(*valNormaliseHeight, &valueForNormaliseHeightInt)
				if err != nil {
					return err
				}
				this.NormaliseHeight = string(valueForNormaliseHeightInt)
			} else {
				this.NormaliseHeight = valueForNormaliseHeight
			}
		}
	}
	
	if valProofDisabled, ok := GetMapValue(objMap, "proofDisabled"); ok {
		if valProofDisabled != nil {
			var valueForProofDisabled string
			err = json.Unmarshal(*valProofDisabled, &valueForProofDisabled)
			if err != nil {
				var valueForProofDisabledInt int32
				err = json.Unmarshal(*valProofDisabled, &valueForProofDisabledInt)
				if err != nil {
					return err
				}
				this.ProofDisabled = string(valueForProofDisabledInt)
			} else {
				this.ProofDisabled = valueForProofDisabled
			}
		}
	}
	
	if valSmartTagClean, ok := GetMapValue(objMap, "smartTagClean"); ok {
		if valSmartTagClean != nil {
			var valueForSmartTagClean *bool
			err = json.Unmarshal(*valSmartTagClean, &valueForSmartTagClean)
			if err != nil {
				return err
			}
			this.SmartTagClean = valueForSmartTagClean
		}
	}
	
	if valKerningMinimalSize, ok := GetMapValue(objMap, "kerningMinimalSize"); ok {
		if valKerningMinimalSize != nil {
			var valueForKerningMinimalSize float64
			err = json.Unmarshal(*valKerningMinimalSize, &valueForKerningMinimalSize)
			if err != nil {
				return err
			}
			this.KerningMinimalSize = valueForKerningMinimalSize
		}
	}
	
	if valKumimoji, ok := GetMapValue(objMap, "kumimoji"); ok {
		if valKumimoji != nil {
			var valueForKumimoji string
			err = json.Unmarshal(*valKumimoji, &valueForKumimoji)
			if err != nil {
				var valueForKumimojiInt int32
				err = json.Unmarshal(*valKumimoji, &valueForKumimojiInt)
				if err != nil {
					return err
				}
				this.Kumimoji = string(valueForKumimojiInt)
			} else {
				this.Kumimoji = valueForKumimoji
			}
		}
	}
	
	if valLanguageId, ok := GetMapValue(objMap, "languageId"); ok {
		if valLanguageId != nil {
			var valueForLanguageId string
			err = json.Unmarshal(*valLanguageId, &valueForLanguageId)
			if err != nil {
				return err
			}
			this.LanguageId = valueForLanguageId
		}
	}
	
	if valAlternativeLanguageId, ok := GetMapValue(objMap, "alternativeLanguageId"); ok {
		if valAlternativeLanguageId != nil {
			var valueForAlternativeLanguageId string
			err = json.Unmarshal(*valAlternativeLanguageId, &valueForAlternativeLanguageId)
			if err != nil {
				return err
			}
			this.AlternativeLanguageId = valueForAlternativeLanguageId
		}
	}
	
	if valIsHardUnderlineFill, ok := GetMapValue(objMap, "isHardUnderlineFill"); ok {
		if valIsHardUnderlineFill != nil {
			var valueForIsHardUnderlineFill string
			err = json.Unmarshal(*valIsHardUnderlineFill, &valueForIsHardUnderlineFill)
			if err != nil {
				var valueForIsHardUnderlineFillInt int32
				err = json.Unmarshal(*valIsHardUnderlineFill, &valueForIsHardUnderlineFillInt)
				if err != nil {
					return err
				}
				this.IsHardUnderlineFill = string(valueForIsHardUnderlineFillInt)
			} else {
				this.IsHardUnderlineFill = valueForIsHardUnderlineFill
			}
		}
	}
	
	if valIsHardUnderlineLine, ok := GetMapValue(objMap, "isHardUnderlineLine"); ok {
		if valIsHardUnderlineLine != nil {
			var valueForIsHardUnderlineLine string
			err = json.Unmarshal(*valIsHardUnderlineLine, &valueForIsHardUnderlineLine)
			if err != nil {
				var valueForIsHardUnderlineLineInt int32
				err = json.Unmarshal(*valIsHardUnderlineLine, &valueForIsHardUnderlineLineInt)
				if err != nil {
					return err
				}
				this.IsHardUnderlineLine = string(valueForIsHardUnderlineLineInt)
			} else {
				this.IsHardUnderlineLine = valueForIsHardUnderlineLine
			}
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
	
	if valUnderlineFillFormat, ok := GetMapValue(objMap, "underlineFillFormat"); ok {
		if valUnderlineFillFormat != nil {
			var valueForUnderlineFillFormat FillFormat
			err = json.Unmarshal(*valUnderlineFillFormat, &valueForUnderlineFillFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("FillFormat", *valUnderlineFillFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valUnderlineFillFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IFillFormat)
			if ok {
				this.UnderlineFillFormat = vInterfaceObject
			}
		}
	}
	
	if valUnderlineLineFormat, ok := GetMapValue(objMap, "underlineLineFormat"); ok {
		if valUnderlineLineFormat != nil {
			var valueForUnderlineLineFormat LineFormat
			err = json.Unmarshal(*valUnderlineLineFormat, &valueForUnderlineLineFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("LineFormat", *valUnderlineLineFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valUnderlineLineFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILineFormat)
			if ok {
				this.UnderlineLineFormat = vInterfaceObject
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
	
	if valLatinFont, ok := GetMapValue(objMap, "latinFont"); ok {
		if valLatinFont != nil {
			var valueForLatinFont string
			err = json.Unmarshal(*valLatinFont, &valueForLatinFont)
			if err != nil {
				return err
			}
			this.LatinFont = valueForLatinFont
		}
	}
	
	if valEastAsianFont, ok := GetMapValue(objMap, "eastAsianFont"); ok {
		if valEastAsianFont != nil {
			var valueForEastAsianFont string
			err = json.Unmarshal(*valEastAsianFont, &valueForEastAsianFont)
			if err != nil {
				return err
			}
			this.EastAsianFont = valueForEastAsianFont
		}
	}
	
	if valComplexScriptFont, ok := GetMapValue(objMap, "complexScriptFont"); ok {
		if valComplexScriptFont != nil {
			var valueForComplexScriptFont string
			err = json.Unmarshal(*valComplexScriptFont, &valueForComplexScriptFont)
			if err != nil {
				return err
			}
			this.ComplexScriptFont = valueForComplexScriptFont
		}
	}
	
	if valSpellCheck, ok := GetMapValue(objMap, "spellCheck"); ok {
		if valSpellCheck != nil {
			var valueForSpellCheck *bool
			err = json.Unmarshal(*valSpellCheck, &valueForSpellCheck)
			if err != nil {
				return err
			}
			this.SpellCheck = valueForSpellCheck
		}
	}

	return nil
}
