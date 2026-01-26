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

// Represents portion resource
type IPortion interface {

	// Gets or sets the link to this resource.
	GetSelfUri() IResourceUri
	SetSelfUri(newValue IResourceUri)

	// List of alternate links.
	GetAlternateLinks() []IResourceUri
	SetAlternateLinks(newValue []IResourceUri)

	// Text.
	GetText() string
	SetText(newValue string)

	// Math paragraph.
	GetMathParagraph() IMathParagraph
	SetMathParagraph(newValue IMathParagraph)

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

	// true to enable spell checking for the portion.
	GetSpellCheck() *bool
	SetSpellCheck(newValue *bool)

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
}

type Portion struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// Text.
	Text string `json:"Text,omitempty"`

	// Math paragraph.
	MathParagraph IMathParagraph `json:"MathParagraph,omitempty"`

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

	// true to enable spell checking for the portion.
	SpellCheck *bool `json:"SpellCheck"`

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
}

func NewPortion() *Portion {
	instance := new(Portion)
	return instance
}

func (this *Portion) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *Portion) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *Portion) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *Portion) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *Portion) GetText() string {
	return this.Text
}

func (this *Portion) SetText(newValue string) {
	this.Text = newValue
}
func (this *Portion) GetMathParagraph() IMathParagraph {
	return this.MathParagraph
}

func (this *Portion) SetMathParagraph(newValue IMathParagraph) {
	this.MathParagraph = newValue
}
func (this *Portion) GetFontBold() string {
	return this.FontBold
}

func (this *Portion) SetFontBold(newValue string) {
	this.FontBold = newValue
}
func (this *Portion) GetFontItalic() string {
	return this.FontItalic
}

func (this *Portion) SetFontItalic(newValue string) {
	this.FontItalic = newValue
}
func (this *Portion) GetFontUnderline() string {
	return this.FontUnderline
}

func (this *Portion) SetFontUnderline(newValue string) {
	this.FontUnderline = newValue
}
func (this *Portion) GetStrikethroughType() string {
	return this.StrikethroughType
}

func (this *Portion) SetStrikethroughType(newValue string) {
	this.StrikethroughType = newValue
}
func (this *Portion) GetTextCapType() string {
	return this.TextCapType
}

func (this *Portion) SetTextCapType(newValue string) {
	this.TextCapType = newValue
}
func (this *Portion) GetEscapement() float64 {
	return this.Escapement
}

func (this *Portion) SetEscapement(newValue float64) {
	this.Escapement = newValue
}
func (this *Portion) GetSpacing() float64 {
	return this.Spacing
}

func (this *Portion) SetSpacing(newValue float64) {
	this.Spacing = newValue
}
func (this *Portion) GetFontColor() string {
	return this.FontColor
}

func (this *Portion) SetFontColor(newValue string) {
	this.FontColor = newValue
}
func (this *Portion) GetHighlightColor() string {
	return this.HighlightColor
}

func (this *Portion) SetHighlightColor(newValue string) {
	this.HighlightColor = newValue
}
func (this *Portion) GetFontHeight() float64 {
	return this.FontHeight
}

func (this *Portion) SetFontHeight(newValue float64) {
	this.FontHeight = newValue
}
func (this *Portion) GetNormaliseHeight() string {
	return this.NormaliseHeight
}

func (this *Portion) SetNormaliseHeight(newValue string) {
	this.NormaliseHeight = newValue
}
func (this *Portion) GetProofDisabled() string {
	return this.ProofDisabled
}

func (this *Portion) SetProofDisabled(newValue string) {
	this.ProofDisabled = newValue
}
func (this *Portion) GetSmartTagClean() *bool {
	return this.SmartTagClean
}

func (this *Portion) SetSmartTagClean(newValue *bool) {
	this.SmartTagClean = newValue
}
func (this *Portion) GetKerningMinimalSize() float64 {
	return this.KerningMinimalSize
}

func (this *Portion) SetKerningMinimalSize(newValue float64) {
	this.KerningMinimalSize = newValue
}
func (this *Portion) GetKumimoji() string {
	return this.Kumimoji
}

func (this *Portion) SetKumimoji(newValue string) {
	this.Kumimoji = newValue
}
func (this *Portion) GetSpellCheck() *bool {
	return this.SpellCheck
}

func (this *Portion) SetSpellCheck(newValue *bool) {
	this.SpellCheck = newValue
}
func (this *Portion) GetLanguageId() string {
	return this.LanguageId
}

func (this *Portion) SetLanguageId(newValue string) {
	this.LanguageId = newValue
}
func (this *Portion) GetAlternativeLanguageId() string {
	return this.AlternativeLanguageId
}

func (this *Portion) SetAlternativeLanguageId(newValue string) {
	this.AlternativeLanguageId = newValue
}
func (this *Portion) GetIsHardUnderlineFill() string {
	return this.IsHardUnderlineFill
}

func (this *Portion) SetIsHardUnderlineFill(newValue string) {
	this.IsHardUnderlineFill = newValue
}
func (this *Portion) GetIsHardUnderlineLine() string {
	return this.IsHardUnderlineLine
}

func (this *Portion) SetIsHardUnderlineLine(newValue string) {
	this.IsHardUnderlineLine = newValue
}
func (this *Portion) GetFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *Portion) SetFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *Portion) GetEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *Portion) SetEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *Portion) GetLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *Portion) SetLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *Portion) GetUnderlineFillFormat() IFillFormat {
	return this.UnderlineFillFormat
}

func (this *Portion) SetUnderlineFillFormat(newValue IFillFormat) {
	this.UnderlineFillFormat = newValue
}
func (this *Portion) GetUnderlineLineFormat() ILineFormat {
	return this.UnderlineLineFormat
}

func (this *Portion) SetUnderlineLineFormat(newValue ILineFormat) {
	this.UnderlineLineFormat = newValue
}
func (this *Portion) GetHyperlinkClick() IHyperlink {
	return this.HyperlinkClick
}

func (this *Portion) SetHyperlinkClick(newValue IHyperlink) {
	this.HyperlinkClick = newValue
}
func (this *Portion) GetHyperlinkMouseOver() IHyperlink {
	return this.HyperlinkMouseOver
}

func (this *Portion) SetHyperlinkMouseOver(newValue IHyperlink) {
	this.HyperlinkMouseOver = newValue
}
func (this *Portion) GetLatinFont() string {
	return this.LatinFont
}

func (this *Portion) SetLatinFont(newValue string) {
	this.LatinFont = newValue
}
func (this *Portion) GetEastAsianFont() string {
	return this.EastAsianFont
}

func (this *Portion) SetEastAsianFont(newValue string) {
	this.EastAsianFont = newValue
}
func (this *Portion) GetComplexScriptFont() string {
	return this.ComplexScriptFont
}

func (this *Portion) SetComplexScriptFont(newValue string) {
	this.ComplexScriptFont = newValue
}

func (this *Portion) UnmarshalJSON(b []byte) error {
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
	
	if valText, ok := GetMapValue(objMap, "text"); ok {
		if valText != nil {
			var valueForText string
			err = json.Unmarshal(*valText, &valueForText)
			if err != nil {
				return err
			}
			this.Text = valueForText
		}
	}
	
	if valMathParagraph, ok := GetMapValue(objMap, "mathParagraph"); ok {
		if valMathParagraph != nil {
			var valueForMathParagraph MathParagraph
			err = json.Unmarshal(*valMathParagraph, &valueForMathParagraph)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("MathParagraph", *valMathParagraph)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valMathParagraph, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IMathParagraph)
			if ok {
				this.MathParagraph = vInterfaceObject
			}
		}
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

	return nil
}
