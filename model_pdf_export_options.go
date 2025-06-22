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

// Provides options that control how a presentation is saved in Pdf format.
type IPdfExportOptions interface {

	// Default regular font for rendering the presentation. 
	GetDefaultRegularFont() string
	SetDefaultRegularFont(newValue string)

	// True to delete delete all embedded binary objects.
	GetDeleteEmbeddedBinaryObjects() *bool
	SetDeleteEmbeddedBinaryObjects(newValue *bool)

	// Default regular font for rendering the presentation. 
	GetGradientStyle() string
	SetGradientStyle(newValue string)

	// Gets of sets list of font fallback rules.
	GetFontFallbackRules() []IFontFallbackRule
	SetFontFallbackRules(newValue []IFontFallbackRule)

	// Gets of sets list of font substitution rules.
	GetFontSubstRules() []IFontSubstRule
	SetFontSubstRules(newValue []IFontSubstRule)

	// True to skip hyperlinks with javascript calls when saving the presentation.
	GetSkipJavaScriptLinks() *bool
	SetSkipJavaScriptLinks(newValue *bool)

	// Export format.
	GetFormat() string
	SetFormat(newValue string)

	// Specifies compression type to be used for all textual content in the document.
	GetTextCompression() string
	SetTextCompression(newValue string)

	// Determines if all characters of font should be embedded or only used subset.
	GetEmbedFullFonts() *bool
	SetEmbedFullFonts(newValue *bool)

	// Desired conformance level for generated PDF document.
	GetCompliance() string
	SetCompliance(newValue string)

	// Returns or sets a value determining resolution of images inside PDF document.  Property affects on file size, time of export and image quality. The default value is 96.
	GetSufficientResolution() float64
	SetSufficientResolution(newValue float64)

	// Returns or sets a value determining the quality of the JPEG images inside PDF document.
	GetJpegQuality() int32
	SetJpegQuality(newValue int32)

	// True to draw black frame around each slide.
	GetDrawSlidesFrame() *bool
	SetDrawSlidesFrame(newValue *bool)

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	GetShowHiddenSlides() *bool
	SetShowHiddenSlides(newValue *bool)

	// True to convert all metafiles used in a presentation to the PNG images.
	GetSaveMetafilesAsPng() *bool
	SetSaveMetafilesAsPng(newValue *bool)

	// Setting user password to protect the PDF document. 
	GetPassword() string
	SetPassword(newValue string)

	// Determines if Aspose.Slides will embed common fonts for ASCII (33..127 code range) text. Fonts for character codes greater than 127 are always embedded. Common fonts list includes PDF's base 14 fonts and additional user specified fonts.
	GetEmbedTrueTypeFontsForASCII() *bool
	SetEmbedTrueTypeFontsForASCII(newValue *bool)

	// Returns or sets an array of user-defined names of font families which Aspose.Slides should consider common.
	GetAdditionalCommonFontFamilies() []string
	SetAdditionalCommonFontFamilies(newValue []string)

	// Slides layouting options
	GetSlidesLayoutOptions() ISlidesLayoutOptions
	SetSlidesLayoutOptions(newValue ISlidesLayoutOptions)

	// Image transparent color.
	GetImageTransparentColor() string
	SetImageTransparentColor(newValue string)

	// True to apply specified ImageTransparentColor  to an image.
	GetApplyImageTransparent() *bool
	SetApplyImageTransparent(newValue *bool)

	// Access permissions that should be granted when the document is opened with user access.  Default is AccessPermissions.None.             
	GetAccessPermissions() IAccessPermissions
	SetAccessPermissions(newValue IAccessPermissions)

	// True to hide Ink elements in exported document.
	GetHideInk() *bool
	SetHideInk(newValue *bool)

	// True to use ROP operation or Opacity for rendering brush.
	GetInterpretMaskOpAsOpacity() *bool
	SetInterpretMaskOpAsOpacity(newValue *bool)

	// True if text should be rasterized as a bitmap and saved to PDF when the font does not support bold styling. This approach can enhance the quality of text in the resulting PDF for certain fonts.
	GetRasterizeUnsupportedFontStyles() *bool
	SetRasterizeUnsupportedFontStyles(newValue *bool)

	// True to convert all OLE data from the presentation to embedded files in the resulting PDF.
	GetIncludeOleData() *bool
	SetIncludeOleData(newValue *bool)
}

type PdfExportOptions struct {

	// Default regular font for rendering the presentation. 
	DefaultRegularFont string `json:"DefaultRegularFont,omitempty"`

	// True to delete delete all embedded binary objects.
	DeleteEmbeddedBinaryObjects *bool `json:"DeleteEmbeddedBinaryObjects"`

	// Default regular font for rendering the presentation. 
	GradientStyle string `json:"GradientStyle,omitempty"`

	// Gets of sets list of font fallback rules.
	FontFallbackRules []IFontFallbackRule `json:"FontFallbackRules,omitempty"`

	// Gets of sets list of font substitution rules.
	FontSubstRules []IFontSubstRule `json:"FontSubstRules,omitempty"`

	// True to skip hyperlinks with javascript calls when saving the presentation.
	SkipJavaScriptLinks *bool `json:"SkipJavaScriptLinks"`

	// Export format.
	Format string `json:"Format,omitempty"`

	// Specifies compression type to be used for all textual content in the document.
	TextCompression string `json:"TextCompression,omitempty"`

	// Determines if all characters of font should be embedded or only used subset.
	EmbedFullFonts *bool `json:"EmbedFullFonts"`

	// Desired conformance level for generated PDF document.
	Compliance string `json:"Compliance,omitempty"`

	// Returns or sets a value determining resolution of images inside PDF document.  Property affects on file size, time of export and image quality. The default value is 96.
	SufficientResolution float64 `json:"SufficientResolution,omitempty"`

	// Returns or sets a value determining the quality of the JPEG images inside PDF document.
	JpegQuality int32 `json:"JpegQuality,omitempty"`

	// True to draw black frame around each slide.
	DrawSlidesFrame *bool `json:"DrawSlidesFrame"`

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	ShowHiddenSlides *bool `json:"ShowHiddenSlides"`

	// True to convert all metafiles used in a presentation to the PNG images.
	SaveMetafilesAsPng *bool `json:"SaveMetafilesAsPng"`

	// Setting user password to protect the PDF document. 
	Password string `json:"Password,omitempty"`

	// Determines if Aspose.Slides will embed common fonts for ASCII (33..127 code range) text. Fonts for character codes greater than 127 are always embedded. Common fonts list includes PDF's base 14 fonts and additional user specified fonts.
	EmbedTrueTypeFontsForASCII *bool `json:"EmbedTrueTypeFontsForASCII"`

	// Returns or sets an array of user-defined names of font families which Aspose.Slides should consider common.
	AdditionalCommonFontFamilies []string `json:"AdditionalCommonFontFamilies,omitempty"`

	// Slides layouting options
	SlidesLayoutOptions ISlidesLayoutOptions `json:"SlidesLayoutOptions,omitempty"`

	// Image transparent color.
	ImageTransparentColor string `json:"ImageTransparentColor,omitempty"`

	// True to apply specified ImageTransparentColor  to an image.
	ApplyImageTransparent *bool `json:"ApplyImageTransparent"`

	// Access permissions that should be granted when the document is opened with user access.  Default is AccessPermissions.None.             
	AccessPermissions IAccessPermissions `json:"AccessPermissions,omitempty"`

	// True to hide Ink elements in exported document.
	HideInk *bool `json:"HideInk"`

	// True to use ROP operation or Opacity for rendering brush.
	InterpretMaskOpAsOpacity *bool `json:"InterpretMaskOpAsOpacity"`

	// True if text should be rasterized as a bitmap and saved to PDF when the font does not support bold styling. This approach can enhance the quality of text in the resulting PDF for certain fonts.
	RasterizeUnsupportedFontStyles *bool `json:"RasterizeUnsupportedFontStyles"`

	// True to convert all OLE data from the presentation to embedded files in the resulting PDF.
	IncludeOleData *bool `json:"IncludeOleData"`
}

func NewPdfExportOptions() *PdfExportOptions {
	instance := new(PdfExportOptions)
	return instance
}

func (this *PdfExportOptions) GetDefaultRegularFont() string {
	return this.DefaultRegularFont
}

func (this *PdfExportOptions) SetDefaultRegularFont(newValue string) {
	this.DefaultRegularFont = newValue
}
func (this *PdfExportOptions) GetDeleteEmbeddedBinaryObjects() *bool {
	return this.DeleteEmbeddedBinaryObjects
}

func (this *PdfExportOptions) SetDeleteEmbeddedBinaryObjects(newValue *bool) {
	this.DeleteEmbeddedBinaryObjects = newValue
}
func (this *PdfExportOptions) GetGradientStyle() string {
	return this.GradientStyle
}

func (this *PdfExportOptions) SetGradientStyle(newValue string) {
	this.GradientStyle = newValue
}
func (this *PdfExportOptions) GetFontFallbackRules() []IFontFallbackRule {
	return this.FontFallbackRules
}

func (this *PdfExportOptions) SetFontFallbackRules(newValue []IFontFallbackRule) {
	this.FontFallbackRules = newValue
}
func (this *PdfExportOptions) GetFontSubstRules() []IFontSubstRule {
	return this.FontSubstRules
}

func (this *PdfExportOptions) SetFontSubstRules(newValue []IFontSubstRule) {
	this.FontSubstRules = newValue
}
func (this *PdfExportOptions) GetSkipJavaScriptLinks() *bool {
	return this.SkipJavaScriptLinks
}

func (this *PdfExportOptions) SetSkipJavaScriptLinks(newValue *bool) {
	this.SkipJavaScriptLinks = newValue
}
func (this *PdfExportOptions) GetFormat() string {
	return this.Format
}

func (this *PdfExportOptions) SetFormat(newValue string) {
	this.Format = newValue
}
func (this *PdfExportOptions) GetTextCompression() string {
	return this.TextCompression
}

func (this *PdfExportOptions) SetTextCompression(newValue string) {
	this.TextCompression = newValue
}
func (this *PdfExportOptions) GetEmbedFullFonts() *bool {
	return this.EmbedFullFonts
}

func (this *PdfExportOptions) SetEmbedFullFonts(newValue *bool) {
	this.EmbedFullFonts = newValue
}
func (this *PdfExportOptions) GetCompliance() string {
	return this.Compliance
}

func (this *PdfExportOptions) SetCompliance(newValue string) {
	this.Compliance = newValue
}
func (this *PdfExportOptions) GetSufficientResolution() float64 {
	return this.SufficientResolution
}

func (this *PdfExportOptions) SetSufficientResolution(newValue float64) {
	this.SufficientResolution = newValue
}
func (this *PdfExportOptions) GetJpegQuality() int32 {
	return this.JpegQuality
}

func (this *PdfExportOptions) SetJpegQuality(newValue int32) {
	this.JpegQuality = newValue
}
func (this *PdfExportOptions) GetDrawSlidesFrame() *bool {
	return this.DrawSlidesFrame
}

func (this *PdfExportOptions) SetDrawSlidesFrame(newValue *bool) {
	this.DrawSlidesFrame = newValue
}
func (this *PdfExportOptions) GetShowHiddenSlides() *bool {
	return this.ShowHiddenSlides
}

func (this *PdfExportOptions) SetShowHiddenSlides(newValue *bool) {
	this.ShowHiddenSlides = newValue
}
func (this *PdfExportOptions) GetSaveMetafilesAsPng() *bool {
	return this.SaveMetafilesAsPng
}

func (this *PdfExportOptions) SetSaveMetafilesAsPng(newValue *bool) {
	this.SaveMetafilesAsPng = newValue
}
func (this *PdfExportOptions) GetPassword() string {
	return this.Password
}

func (this *PdfExportOptions) SetPassword(newValue string) {
	this.Password = newValue
}
func (this *PdfExportOptions) GetEmbedTrueTypeFontsForASCII() *bool {
	return this.EmbedTrueTypeFontsForASCII
}

func (this *PdfExportOptions) SetEmbedTrueTypeFontsForASCII(newValue *bool) {
	this.EmbedTrueTypeFontsForASCII = newValue
}
func (this *PdfExportOptions) GetAdditionalCommonFontFamilies() []string {
	return this.AdditionalCommonFontFamilies
}

func (this *PdfExportOptions) SetAdditionalCommonFontFamilies(newValue []string) {
	this.AdditionalCommonFontFamilies = newValue
}
func (this *PdfExportOptions) GetSlidesLayoutOptions() ISlidesLayoutOptions {
	return this.SlidesLayoutOptions
}

func (this *PdfExportOptions) SetSlidesLayoutOptions(newValue ISlidesLayoutOptions) {
	this.SlidesLayoutOptions = newValue
}
func (this *PdfExportOptions) GetImageTransparentColor() string {
	return this.ImageTransparentColor
}

func (this *PdfExportOptions) SetImageTransparentColor(newValue string) {
	this.ImageTransparentColor = newValue
}
func (this *PdfExportOptions) GetApplyImageTransparent() *bool {
	return this.ApplyImageTransparent
}

func (this *PdfExportOptions) SetApplyImageTransparent(newValue *bool) {
	this.ApplyImageTransparent = newValue
}
func (this *PdfExportOptions) GetAccessPermissions() IAccessPermissions {
	return this.AccessPermissions
}

func (this *PdfExportOptions) SetAccessPermissions(newValue IAccessPermissions) {
	this.AccessPermissions = newValue
}
func (this *PdfExportOptions) GetHideInk() *bool {
	return this.HideInk
}

func (this *PdfExportOptions) SetHideInk(newValue *bool) {
	this.HideInk = newValue
}
func (this *PdfExportOptions) GetInterpretMaskOpAsOpacity() *bool {
	return this.InterpretMaskOpAsOpacity
}

func (this *PdfExportOptions) SetInterpretMaskOpAsOpacity(newValue *bool) {
	this.InterpretMaskOpAsOpacity = newValue
}
func (this *PdfExportOptions) GetRasterizeUnsupportedFontStyles() *bool {
	return this.RasterizeUnsupportedFontStyles
}

func (this *PdfExportOptions) SetRasterizeUnsupportedFontStyles(newValue *bool) {
	this.RasterizeUnsupportedFontStyles = newValue
}
func (this *PdfExportOptions) GetIncludeOleData() *bool {
	return this.IncludeOleData
}

func (this *PdfExportOptions) SetIncludeOleData(newValue *bool) {
	this.IncludeOleData = newValue
}

func (this *PdfExportOptions) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valDefaultRegularFont, ok := GetMapValue(objMap, "defaultRegularFont"); ok {
		if valDefaultRegularFont != nil {
			var valueForDefaultRegularFont string
			err = json.Unmarshal(*valDefaultRegularFont, &valueForDefaultRegularFont)
			if err != nil {
				return err
			}
			this.DefaultRegularFont = valueForDefaultRegularFont
		}
	}
	
	if valDeleteEmbeddedBinaryObjects, ok := GetMapValue(objMap, "deleteEmbeddedBinaryObjects"); ok {
		if valDeleteEmbeddedBinaryObjects != nil {
			var valueForDeleteEmbeddedBinaryObjects *bool
			err = json.Unmarshal(*valDeleteEmbeddedBinaryObjects, &valueForDeleteEmbeddedBinaryObjects)
			if err != nil {
				return err
			}
			this.DeleteEmbeddedBinaryObjects = valueForDeleteEmbeddedBinaryObjects
		}
	}
	
	if valGradientStyle, ok := GetMapValue(objMap, "gradientStyle"); ok {
		if valGradientStyle != nil {
			var valueForGradientStyle string
			err = json.Unmarshal(*valGradientStyle, &valueForGradientStyle)
			if err != nil {
				var valueForGradientStyleInt int32
				err = json.Unmarshal(*valGradientStyle, &valueForGradientStyleInt)
				if err != nil {
					return err
				}
				this.GradientStyle = string(valueForGradientStyleInt)
			} else {
				this.GradientStyle = valueForGradientStyle
			}
		}
	}
	
	if valFontFallbackRules, ok := GetMapValue(objMap, "fontFallbackRules"); ok {
		if valFontFallbackRules != nil {
			var valueForFontFallbackRules []json.RawMessage
			err = json.Unmarshal(*valFontFallbackRules, &valueForFontFallbackRules)
			if err != nil {
				return err
			}
			valueForIFontFallbackRules := make([]IFontFallbackRule, len(valueForFontFallbackRules))
			for i, v := range valueForFontFallbackRules {
				vObject, err := createObjectForType("FontFallbackRule", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIFontFallbackRules[i] = vObject.(IFontFallbackRule)
				}
			}
			this.FontFallbackRules = valueForIFontFallbackRules
		}
	}
	
	if valFontSubstRules, ok := GetMapValue(objMap, "fontSubstRules"); ok {
		if valFontSubstRules != nil {
			var valueForFontSubstRules []json.RawMessage
			err = json.Unmarshal(*valFontSubstRules, &valueForFontSubstRules)
			if err != nil {
				return err
			}
			valueForIFontSubstRules := make([]IFontSubstRule, len(valueForFontSubstRules))
			for i, v := range valueForFontSubstRules {
				vObject, err := createObjectForType("FontSubstRule", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIFontSubstRules[i] = vObject.(IFontSubstRule)
				}
			}
			this.FontSubstRules = valueForIFontSubstRules
		}
	}
	
	if valSkipJavaScriptLinks, ok := GetMapValue(objMap, "skipJavaScriptLinks"); ok {
		if valSkipJavaScriptLinks != nil {
			var valueForSkipJavaScriptLinks *bool
			err = json.Unmarshal(*valSkipJavaScriptLinks, &valueForSkipJavaScriptLinks)
			if err != nil {
				return err
			}
			this.SkipJavaScriptLinks = valueForSkipJavaScriptLinks
		}
	}
	
	if valFormat, ok := GetMapValue(objMap, "format"); ok {
		if valFormat != nil {
			var valueForFormat string
			err = json.Unmarshal(*valFormat, &valueForFormat)
			if err != nil {
				return err
			}
			this.Format = valueForFormat
		}
	}
	
	if valTextCompression, ok := GetMapValue(objMap, "textCompression"); ok {
		if valTextCompression != nil {
			var valueForTextCompression string
			err = json.Unmarshal(*valTextCompression, &valueForTextCompression)
			if err != nil {
				var valueForTextCompressionInt int32
				err = json.Unmarshal(*valTextCompression, &valueForTextCompressionInt)
				if err != nil {
					return err
				}
				this.TextCompression = string(valueForTextCompressionInt)
			} else {
				this.TextCompression = valueForTextCompression
			}
		}
	}
	
	if valEmbedFullFonts, ok := GetMapValue(objMap, "embedFullFonts"); ok {
		if valEmbedFullFonts != nil {
			var valueForEmbedFullFonts *bool
			err = json.Unmarshal(*valEmbedFullFonts, &valueForEmbedFullFonts)
			if err != nil {
				return err
			}
			this.EmbedFullFonts = valueForEmbedFullFonts
		}
	}
	
	if valCompliance, ok := GetMapValue(objMap, "compliance"); ok {
		if valCompliance != nil {
			var valueForCompliance string
			err = json.Unmarshal(*valCompliance, &valueForCompliance)
			if err != nil {
				var valueForComplianceInt int32
				err = json.Unmarshal(*valCompliance, &valueForComplianceInt)
				if err != nil {
					return err
				}
				this.Compliance = string(valueForComplianceInt)
			} else {
				this.Compliance = valueForCompliance
			}
		}
	}
	
	if valSufficientResolution, ok := GetMapValue(objMap, "sufficientResolution"); ok {
		if valSufficientResolution != nil {
			var valueForSufficientResolution float64
			err = json.Unmarshal(*valSufficientResolution, &valueForSufficientResolution)
			if err != nil {
				return err
			}
			this.SufficientResolution = valueForSufficientResolution
		}
	}
	
	if valJpegQuality, ok := GetMapValue(objMap, "jpegQuality"); ok {
		if valJpegQuality != nil {
			var valueForJpegQuality int32
			err = json.Unmarshal(*valJpegQuality, &valueForJpegQuality)
			if err != nil {
				return err
			}
			this.JpegQuality = valueForJpegQuality
		}
	}
	
	if valDrawSlidesFrame, ok := GetMapValue(objMap, "drawSlidesFrame"); ok {
		if valDrawSlidesFrame != nil {
			var valueForDrawSlidesFrame *bool
			err = json.Unmarshal(*valDrawSlidesFrame, &valueForDrawSlidesFrame)
			if err != nil {
				return err
			}
			this.DrawSlidesFrame = valueForDrawSlidesFrame
		}
	}
	
	if valShowHiddenSlides, ok := GetMapValue(objMap, "showHiddenSlides"); ok {
		if valShowHiddenSlides != nil {
			var valueForShowHiddenSlides *bool
			err = json.Unmarshal(*valShowHiddenSlides, &valueForShowHiddenSlides)
			if err != nil {
				return err
			}
			this.ShowHiddenSlides = valueForShowHiddenSlides
		}
	}
	
	if valSaveMetafilesAsPng, ok := GetMapValue(objMap, "saveMetafilesAsPng"); ok {
		if valSaveMetafilesAsPng != nil {
			var valueForSaveMetafilesAsPng *bool
			err = json.Unmarshal(*valSaveMetafilesAsPng, &valueForSaveMetafilesAsPng)
			if err != nil {
				return err
			}
			this.SaveMetafilesAsPng = valueForSaveMetafilesAsPng
		}
	}
	
	if valPassword, ok := GetMapValue(objMap, "password"); ok {
		if valPassword != nil {
			var valueForPassword string
			err = json.Unmarshal(*valPassword, &valueForPassword)
			if err != nil {
				return err
			}
			this.Password = valueForPassword
		}
	}
	
	if valEmbedTrueTypeFontsForASCII, ok := GetMapValue(objMap, "embedTrueTypeFontsForASCII"); ok {
		if valEmbedTrueTypeFontsForASCII != nil {
			var valueForEmbedTrueTypeFontsForASCII *bool
			err = json.Unmarshal(*valEmbedTrueTypeFontsForASCII, &valueForEmbedTrueTypeFontsForASCII)
			if err != nil {
				return err
			}
			this.EmbedTrueTypeFontsForASCII = valueForEmbedTrueTypeFontsForASCII
		}
	}
	
	if valAdditionalCommonFontFamilies, ok := GetMapValue(objMap, "additionalCommonFontFamilies"); ok {
		if valAdditionalCommonFontFamilies != nil {
			var valueForAdditionalCommonFontFamilies []string
			err = json.Unmarshal(*valAdditionalCommonFontFamilies, &valueForAdditionalCommonFontFamilies)
			if err != nil {
				return err
			}
			this.AdditionalCommonFontFamilies = valueForAdditionalCommonFontFamilies
		}
	}
	
	if valSlidesLayoutOptions, ok := GetMapValue(objMap, "slidesLayoutOptions"); ok {
		if valSlidesLayoutOptions != nil {
			var valueForSlidesLayoutOptions SlidesLayoutOptions
			err = json.Unmarshal(*valSlidesLayoutOptions, &valueForSlidesLayoutOptions)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("SlidesLayoutOptions", *valSlidesLayoutOptions)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSlidesLayoutOptions, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ISlidesLayoutOptions)
			if ok {
				this.SlidesLayoutOptions = vInterfaceObject
			}
		}
	}
	
	if valImageTransparentColor, ok := GetMapValue(objMap, "imageTransparentColor"); ok {
		if valImageTransparentColor != nil {
			var valueForImageTransparentColor string
			err = json.Unmarshal(*valImageTransparentColor, &valueForImageTransparentColor)
			if err != nil {
				return err
			}
			this.ImageTransparentColor = valueForImageTransparentColor
		}
	}
	
	if valApplyImageTransparent, ok := GetMapValue(objMap, "applyImageTransparent"); ok {
		if valApplyImageTransparent != nil {
			var valueForApplyImageTransparent *bool
			err = json.Unmarshal(*valApplyImageTransparent, &valueForApplyImageTransparent)
			if err != nil {
				return err
			}
			this.ApplyImageTransparent = valueForApplyImageTransparent
		}
	}
	
	if valAccessPermissions, ok := GetMapValue(objMap, "accessPermissions"); ok {
		if valAccessPermissions != nil {
			var valueForAccessPermissions AccessPermissions
			err = json.Unmarshal(*valAccessPermissions, &valueForAccessPermissions)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("AccessPermissions", *valAccessPermissions)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valAccessPermissions, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IAccessPermissions)
			if ok {
				this.AccessPermissions = vInterfaceObject
			}
		}
	}
	
	if valHideInk, ok := GetMapValue(objMap, "hideInk"); ok {
		if valHideInk != nil {
			var valueForHideInk *bool
			err = json.Unmarshal(*valHideInk, &valueForHideInk)
			if err != nil {
				return err
			}
			this.HideInk = valueForHideInk
		}
	}
	
	if valInterpretMaskOpAsOpacity, ok := GetMapValue(objMap, "interpretMaskOpAsOpacity"); ok {
		if valInterpretMaskOpAsOpacity != nil {
			var valueForInterpretMaskOpAsOpacity *bool
			err = json.Unmarshal(*valInterpretMaskOpAsOpacity, &valueForInterpretMaskOpAsOpacity)
			if err != nil {
				return err
			}
			this.InterpretMaskOpAsOpacity = valueForInterpretMaskOpAsOpacity
		}
	}
	
	if valRasterizeUnsupportedFontStyles, ok := GetMapValue(objMap, "rasterizeUnsupportedFontStyles"); ok {
		if valRasterizeUnsupportedFontStyles != nil {
			var valueForRasterizeUnsupportedFontStyles *bool
			err = json.Unmarshal(*valRasterizeUnsupportedFontStyles, &valueForRasterizeUnsupportedFontStyles)
			if err != nil {
				return err
			}
			this.RasterizeUnsupportedFontStyles = valueForRasterizeUnsupportedFontStyles
		}
	}
	
	if valIncludeOleData, ok := GetMapValue(objMap, "includeOleData"); ok {
		if valIncludeOleData != nil {
			var valueForIncludeOleData *bool
			err = json.Unmarshal(*valIncludeOleData, &valueForIncludeOleData)
			if err != nil {
				return err
			}
			this.IncludeOleData = valueForIncludeOleData
		}
	}

	return nil
}
