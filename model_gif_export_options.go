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

// Provides options that control how a presentation is saved in Gif format.
type IGifExportOptions interface {

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

	// Gets or sets the height of slides in the output image format.
	GetHeight() int32
	SetHeight(newValue int32)

	// Gets or sets the height of slides in the output the output image format.
	GetWidth() int32
	SetWidth(newValue int32)

	// Determines whether hidden slides will be exported.
	GetExportHiddenSlides() *bool
	SetExportHiddenSlides(newValue *bool)

	// Gets or sets transition FPS [frames/sec]
	GetTransitionFps() int32
	SetTransitionFps(newValue int32)

	// Gets or sets default delay time [ms].
	GetDefaultDelay() int32
	SetDefaultDelay(newValue int32)
}

type GifExportOptions struct {

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

	// Gets or sets the height of slides in the output image format.
	Height int32 `json:"Height,omitempty"`

	// Gets or sets the height of slides in the output the output image format.
	Width int32 `json:"Width,omitempty"`

	// Determines whether hidden slides will be exported.
	ExportHiddenSlides *bool `json:"ExportHiddenSlides"`

	// Gets or sets transition FPS [frames/sec]
	TransitionFps int32 `json:"TransitionFps,omitempty"`

	// Gets or sets default delay time [ms].
	DefaultDelay int32 `json:"DefaultDelay,omitempty"`
}

func NewGifExportOptions() *GifExportOptions {
	instance := new(GifExportOptions)
	return instance
}

func (this *GifExportOptions) GetDefaultRegularFont() string {
	return this.DefaultRegularFont
}

func (this *GifExportOptions) SetDefaultRegularFont(newValue string) {
	this.DefaultRegularFont = newValue
}
func (this *GifExportOptions) GetDeleteEmbeddedBinaryObjects() *bool {
	return this.DeleteEmbeddedBinaryObjects
}

func (this *GifExportOptions) SetDeleteEmbeddedBinaryObjects(newValue *bool) {
	this.DeleteEmbeddedBinaryObjects = newValue
}
func (this *GifExportOptions) GetGradientStyle() string {
	return this.GradientStyle
}

func (this *GifExportOptions) SetGradientStyle(newValue string) {
	this.GradientStyle = newValue
}
func (this *GifExportOptions) GetFontFallbackRules() []IFontFallbackRule {
	return this.FontFallbackRules
}

func (this *GifExportOptions) SetFontFallbackRules(newValue []IFontFallbackRule) {
	this.FontFallbackRules = newValue
}
func (this *GifExportOptions) GetFontSubstRules() []IFontSubstRule {
	return this.FontSubstRules
}

func (this *GifExportOptions) SetFontSubstRules(newValue []IFontSubstRule) {
	this.FontSubstRules = newValue
}
func (this *GifExportOptions) GetSkipJavaScriptLinks() *bool {
	return this.SkipJavaScriptLinks
}

func (this *GifExportOptions) SetSkipJavaScriptLinks(newValue *bool) {
	this.SkipJavaScriptLinks = newValue
}
func (this *GifExportOptions) GetFormat() string {
	return this.Format
}

func (this *GifExportOptions) SetFormat(newValue string) {
	this.Format = newValue
}
func (this *GifExportOptions) GetHeight() int32 {
	return this.Height
}

func (this *GifExportOptions) SetHeight(newValue int32) {
	this.Height = newValue
}
func (this *GifExportOptions) GetWidth() int32 {
	return this.Width
}

func (this *GifExportOptions) SetWidth(newValue int32) {
	this.Width = newValue
}
func (this *GifExportOptions) GetExportHiddenSlides() *bool {
	return this.ExportHiddenSlides
}

func (this *GifExportOptions) SetExportHiddenSlides(newValue *bool) {
	this.ExportHiddenSlides = newValue
}
func (this *GifExportOptions) GetTransitionFps() int32 {
	return this.TransitionFps
}

func (this *GifExportOptions) SetTransitionFps(newValue int32) {
	this.TransitionFps = newValue
}
func (this *GifExportOptions) GetDefaultDelay() int32 {
	return this.DefaultDelay
}

func (this *GifExportOptions) SetDefaultDelay(newValue int32) {
	this.DefaultDelay = newValue
}

func (this *GifExportOptions) UnmarshalJSON(b []byte) error {
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
	
	if valHeight, ok := GetMapValue(objMap, "height"); ok {
		if valHeight != nil {
			var valueForHeight int32
			err = json.Unmarshal(*valHeight, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	
	if valWidth, ok := GetMapValue(objMap, "width"); ok {
		if valWidth != nil {
			var valueForWidth int32
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	
	if valExportHiddenSlides, ok := GetMapValue(objMap, "exportHiddenSlides"); ok {
		if valExportHiddenSlides != nil {
			var valueForExportHiddenSlides *bool
			err = json.Unmarshal(*valExportHiddenSlides, &valueForExportHiddenSlides)
			if err != nil {
				return err
			}
			this.ExportHiddenSlides = valueForExportHiddenSlides
		}
	}
	
	if valTransitionFps, ok := GetMapValue(objMap, "transitionFps"); ok {
		if valTransitionFps != nil {
			var valueForTransitionFps int32
			err = json.Unmarshal(*valTransitionFps, &valueForTransitionFps)
			if err != nil {
				return err
			}
			this.TransitionFps = valueForTransitionFps
		}
	}
	
	if valDefaultDelay, ok := GetMapValue(objMap, "defaultDelay"); ok {
		if valDefaultDelay != nil {
			var valueForDefaultDelay int32
			err = json.Unmarshal(*valDefaultDelay, &valueForDefaultDelay)
			if err != nil {
				return err
			}
			this.DefaultDelay = valueForDefaultDelay
		}
	}

	return nil
}
