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

// Provides options that control how a presentation is saved in Html5 format.
type IHtml5ExportOptions interface {

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

	// Gets or sets transitions animation option.
	GetAnimateTransitions() *bool
	SetAnimateTransitions(newValue *bool)

	// Gets or sets shapes animation option.
	GetAnimateShapes() *bool
	SetAnimateShapes(newValue *bool)

	// Gets or sets embed images option.
	GetEmbedImages() *bool
	SetEmbedImages(newValue *bool)

	// Slides layouting options
	GetSlidesLayoutOptions() ISlidesLayoutOptions
	SetSlidesLayoutOptions(newValue ISlidesLayoutOptions)

	// Path to custom templates
	GetTemplatesPath() string
	SetTemplatesPath(newValue string)

	// true to disable ligatures in the rendered output.
	GetDisableFontLigatures() *bool
	SetDisableFontLigatures(newValue *bool)
}

type Html5ExportOptions struct {

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

	// Gets or sets transitions animation option.
	AnimateTransitions *bool `json:"AnimateTransitions"`

	// Gets or sets shapes animation option.
	AnimateShapes *bool `json:"AnimateShapes"`

	// Gets or sets embed images option.
	EmbedImages *bool `json:"EmbedImages"`

	// Slides layouting options
	SlidesLayoutOptions ISlidesLayoutOptions `json:"SlidesLayoutOptions,omitempty"`

	// Path to custom templates
	TemplatesPath string `json:"TemplatesPath,omitempty"`

	// true to disable ligatures in the rendered output.
	DisableFontLigatures *bool `json:"DisableFontLigatures"`
}

func NewHtml5ExportOptions() *Html5ExportOptions {
	instance := new(Html5ExportOptions)
	return instance
}

func (this *Html5ExportOptions) GetDefaultRegularFont() string {
	return this.DefaultRegularFont
}

func (this *Html5ExportOptions) SetDefaultRegularFont(newValue string) {
	this.DefaultRegularFont = newValue
}
func (this *Html5ExportOptions) GetDeleteEmbeddedBinaryObjects() *bool {
	return this.DeleteEmbeddedBinaryObjects
}

func (this *Html5ExportOptions) SetDeleteEmbeddedBinaryObjects(newValue *bool) {
	this.DeleteEmbeddedBinaryObjects = newValue
}
func (this *Html5ExportOptions) GetGradientStyle() string {
	return this.GradientStyle
}

func (this *Html5ExportOptions) SetGradientStyle(newValue string) {
	this.GradientStyle = newValue
}
func (this *Html5ExportOptions) GetFontFallbackRules() []IFontFallbackRule {
	return this.FontFallbackRules
}

func (this *Html5ExportOptions) SetFontFallbackRules(newValue []IFontFallbackRule) {
	this.FontFallbackRules = newValue
}
func (this *Html5ExportOptions) GetFontSubstRules() []IFontSubstRule {
	return this.FontSubstRules
}

func (this *Html5ExportOptions) SetFontSubstRules(newValue []IFontSubstRule) {
	this.FontSubstRules = newValue
}
func (this *Html5ExportOptions) GetSkipJavaScriptLinks() *bool {
	return this.SkipJavaScriptLinks
}

func (this *Html5ExportOptions) SetSkipJavaScriptLinks(newValue *bool) {
	this.SkipJavaScriptLinks = newValue
}
func (this *Html5ExportOptions) GetFormat() string {
	return this.Format
}

func (this *Html5ExportOptions) SetFormat(newValue string) {
	this.Format = newValue
}
func (this *Html5ExportOptions) GetAnimateTransitions() *bool {
	return this.AnimateTransitions
}

func (this *Html5ExportOptions) SetAnimateTransitions(newValue *bool) {
	this.AnimateTransitions = newValue
}
func (this *Html5ExportOptions) GetAnimateShapes() *bool {
	return this.AnimateShapes
}

func (this *Html5ExportOptions) SetAnimateShapes(newValue *bool) {
	this.AnimateShapes = newValue
}
func (this *Html5ExportOptions) GetEmbedImages() *bool {
	return this.EmbedImages
}

func (this *Html5ExportOptions) SetEmbedImages(newValue *bool) {
	this.EmbedImages = newValue
}
func (this *Html5ExportOptions) GetSlidesLayoutOptions() ISlidesLayoutOptions {
	return this.SlidesLayoutOptions
}

func (this *Html5ExportOptions) SetSlidesLayoutOptions(newValue ISlidesLayoutOptions) {
	this.SlidesLayoutOptions = newValue
}
func (this *Html5ExportOptions) GetTemplatesPath() string {
	return this.TemplatesPath
}

func (this *Html5ExportOptions) SetTemplatesPath(newValue string) {
	this.TemplatesPath = newValue
}
func (this *Html5ExportOptions) GetDisableFontLigatures() *bool {
	return this.DisableFontLigatures
}

func (this *Html5ExportOptions) SetDisableFontLigatures(newValue *bool) {
	this.DisableFontLigatures = newValue
}

func (this *Html5ExportOptions) UnmarshalJSON(b []byte) error {
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
	
	if valAnimateTransitions, ok := GetMapValue(objMap, "animateTransitions"); ok {
		if valAnimateTransitions != nil {
			var valueForAnimateTransitions *bool
			err = json.Unmarshal(*valAnimateTransitions, &valueForAnimateTransitions)
			if err != nil {
				return err
			}
			this.AnimateTransitions = valueForAnimateTransitions
		}
	}
	
	if valAnimateShapes, ok := GetMapValue(objMap, "animateShapes"); ok {
		if valAnimateShapes != nil {
			var valueForAnimateShapes *bool
			err = json.Unmarshal(*valAnimateShapes, &valueForAnimateShapes)
			if err != nil {
				return err
			}
			this.AnimateShapes = valueForAnimateShapes
		}
	}
	
	if valEmbedImages, ok := GetMapValue(objMap, "embedImages"); ok {
		if valEmbedImages != nil {
			var valueForEmbedImages *bool
			err = json.Unmarshal(*valEmbedImages, &valueForEmbedImages)
			if err != nil {
				return err
			}
			this.EmbedImages = valueForEmbedImages
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
	
	if valTemplatesPath, ok := GetMapValue(objMap, "templatesPath"); ok {
		if valTemplatesPath != nil {
			var valueForTemplatesPath string
			err = json.Unmarshal(*valTemplatesPath, &valueForTemplatesPath)
			if err != nil {
				return err
			}
			this.TemplatesPath = valueForTemplatesPath
		}
	}
	
	if valDisableFontLigatures, ok := GetMapValue(objMap, "disableFontLigatures"); ok {
		if valDisableFontLigatures != nil {
			var valueForDisableFontLigatures *bool
			err = json.Unmarshal(*valDisableFontLigatures, &valueForDisableFontLigatures)
			if err != nil {
				return err
			}
			this.DisableFontLigatures = valueForDisableFontLigatures
		}
	}

	return nil
}
