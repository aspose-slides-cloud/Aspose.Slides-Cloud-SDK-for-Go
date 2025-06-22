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

// Provides options that control how a presentation is saved in SWF format.
type ISwfExportOptions interface {

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

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	GetShowHiddenSlides() *bool
	SetShowHiddenSlides(newValue *bool)

	// Specifies whether the generated SWF document should be compressed or not. Default is true. 
	GetCompressed() *bool
	SetCompressed(newValue *bool)

	// Specifies whether the generated SWF document should include the integrated document viewer or not. Default is true. 
	GetViewerIncluded() *bool
	SetViewerIncluded(newValue *bool)

	// Specifies whether border around pages should be shown. Default is true. 
	GetShowPageBorder() *bool
	SetShowPageBorder(newValue *bool)

	// Show/hide fullscreen button. Can be overridden in flashvars. Default is true. 
	GetShowFullScreen() *bool
	SetShowFullScreen(newValue *bool)

	// Show/hide page stepper. Can be overridden in flashvars. Default is true. 
	GetShowPageStepper() *bool
	SetShowPageStepper(newValue *bool)

	// Show/hide search section. Can be overridden in flashvars. Default is true. 
	GetShowSearch() *bool
	SetShowSearch(newValue *bool)

	// Show/hide whole top pane. Can be overridden in flashvars. Default is true. 
	GetShowTopPane() *bool
	SetShowTopPane(newValue *bool)

	// Show/hide bottom pane. Can be overridden in flashvars. Default is true. 
	GetShowBottomPane() *bool
	SetShowBottomPane(newValue *bool)

	// Show/hide left pane. Can be overridden in flashvars. Default is true. 
	GetShowLeftPane() *bool
	SetShowLeftPane(newValue *bool)

	// Start with opened left pane. Can be overridden in flashvars. Default is false. 
	GetStartOpenLeftPane() *bool
	SetStartOpenLeftPane(newValue *bool)

	// Enable/disable context menu. Default is true. 
	GetEnableContextMenu() *bool
	SetEnableContextMenu(newValue *bool)

	// Image that will be displayed as logo in the top right corner of the viewer. The image data is a base 64 string. Image should be 32x64 pixels PNG image, otherwise logo can be displayed improperly. 
	GetLogoImage() string
	SetLogoImage(newValue string)

	// Gets or sets the full hyperlink address for a logo. Has an effect only if a LogoImage is specified. 
	GetLogoLink() string
	SetLogoLink(newValue string)

	// Specifies the quality of JPEG images. Default is 95.
	GetJpegQuality() int32
	SetJpegQuality(newValue int32)

	// Slides layouting options
	GetSlidesLayoutOptions() ISlidesLayoutOptions
	SetSlidesLayoutOptions(newValue ISlidesLayoutOptions)
}

type SwfExportOptions struct {

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

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	ShowHiddenSlides *bool `json:"ShowHiddenSlides"`

	// Specifies whether the generated SWF document should be compressed or not. Default is true. 
	Compressed *bool `json:"Compressed"`

	// Specifies whether the generated SWF document should include the integrated document viewer or not. Default is true. 
	ViewerIncluded *bool `json:"ViewerIncluded"`

	// Specifies whether border around pages should be shown. Default is true. 
	ShowPageBorder *bool `json:"ShowPageBorder"`

	// Show/hide fullscreen button. Can be overridden in flashvars. Default is true. 
	ShowFullScreen *bool `json:"ShowFullScreen"`

	// Show/hide page stepper. Can be overridden in flashvars. Default is true. 
	ShowPageStepper *bool `json:"ShowPageStepper"`

	// Show/hide search section. Can be overridden in flashvars. Default is true. 
	ShowSearch *bool `json:"ShowSearch"`

	// Show/hide whole top pane. Can be overridden in flashvars. Default is true. 
	ShowTopPane *bool `json:"ShowTopPane"`

	// Show/hide bottom pane. Can be overridden in flashvars. Default is true. 
	ShowBottomPane *bool `json:"ShowBottomPane"`

	// Show/hide left pane. Can be overridden in flashvars. Default is true. 
	ShowLeftPane *bool `json:"ShowLeftPane"`

	// Start with opened left pane. Can be overridden in flashvars. Default is false. 
	StartOpenLeftPane *bool `json:"StartOpenLeftPane"`

	// Enable/disable context menu. Default is true. 
	EnableContextMenu *bool `json:"EnableContextMenu"`

	// Image that will be displayed as logo in the top right corner of the viewer. The image data is a base 64 string. Image should be 32x64 pixels PNG image, otherwise logo can be displayed improperly. 
	LogoImage string `json:"LogoImage,omitempty"`

	// Gets or sets the full hyperlink address for a logo. Has an effect only if a LogoImage is specified. 
	LogoLink string `json:"LogoLink,omitempty"`

	// Specifies the quality of JPEG images. Default is 95.
	JpegQuality int32 `json:"JpegQuality,omitempty"`

	// Slides layouting options
	SlidesLayoutOptions ISlidesLayoutOptions `json:"SlidesLayoutOptions,omitempty"`
}

func NewSwfExportOptions() *SwfExportOptions {
	instance := new(SwfExportOptions)
	return instance
}

func (this *SwfExportOptions) GetDefaultRegularFont() string {
	return this.DefaultRegularFont
}

func (this *SwfExportOptions) SetDefaultRegularFont(newValue string) {
	this.DefaultRegularFont = newValue
}
func (this *SwfExportOptions) GetDeleteEmbeddedBinaryObjects() *bool {
	return this.DeleteEmbeddedBinaryObjects
}

func (this *SwfExportOptions) SetDeleteEmbeddedBinaryObjects(newValue *bool) {
	this.DeleteEmbeddedBinaryObjects = newValue
}
func (this *SwfExportOptions) GetGradientStyle() string {
	return this.GradientStyle
}

func (this *SwfExportOptions) SetGradientStyle(newValue string) {
	this.GradientStyle = newValue
}
func (this *SwfExportOptions) GetFontFallbackRules() []IFontFallbackRule {
	return this.FontFallbackRules
}

func (this *SwfExportOptions) SetFontFallbackRules(newValue []IFontFallbackRule) {
	this.FontFallbackRules = newValue
}
func (this *SwfExportOptions) GetFontSubstRules() []IFontSubstRule {
	return this.FontSubstRules
}

func (this *SwfExportOptions) SetFontSubstRules(newValue []IFontSubstRule) {
	this.FontSubstRules = newValue
}
func (this *SwfExportOptions) GetSkipJavaScriptLinks() *bool {
	return this.SkipJavaScriptLinks
}

func (this *SwfExportOptions) SetSkipJavaScriptLinks(newValue *bool) {
	this.SkipJavaScriptLinks = newValue
}
func (this *SwfExportOptions) GetFormat() string {
	return this.Format
}

func (this *SwfExportOptions) SetFormat(newValue string) {
	this.Format = newValue
}
func (this *SwfExportOptions) GetShowHiddenSlides() *bool {
	return this.ShowHiddenSlides
}

func (this *SwfExportOptions) SetShowHiddenSlides(newValue *bool) {
	this.ShowHiddenSlides = newValue
}
func (this *SwfExportOptions) GetCompressed() *bool {
	return this.Compressed
}

func (this *SwfExportOptions) SetCompressed(newValue *bool) {
	this.Compressed = newValue
}
func (this *SwfExportOptions) GetViewerIncluded() *bool {
	return this.ViewerIncluded
}

func (this *SwfExportOptions) SetViewerIncluded(newValue *bool) {
	this.ViewerIncluded = newValue
}
func (this *SwfExportOptions) GetShowPageBorder() *bool {
	return this.ShowPageBorder
}

func (this *SwfExportOptions) SetShowPageBorder(newValue *bool) {
	this.ShowPageBorder = newValue
}
func (this *SwfExportOptions) GetShowFullScreen() *bool {
	return this.ShowFullScreen
}

func (this *SwfExportOptions) SetShowFullScreen(newValue *bool) {
	this.ShowFullScreen = newValue
}
func (this *SwfExportOptions) GetShowPageStepper() *bool {
	return this.ShowPageStepper
}

func (this *SwfExportOptions) SetShowPageStepper(newValue *bool) {
	this.ShowPageStepper = newValue
}
func (this *SwfExportOptions) GetShowSearch() *bool {
	return this.ShowSearch
}

func (this *SwfExportOptions) SetShowSearch(newValue *bool) {
	this.ShowSearch = newValue
}
func (this *SwfExportOptions) GetShowTopPane() *bool {
	return this.ShowTopPane
}

func (this *SwfExportOptions) SetShowTopPane(newValue *bool) {
	this.ShowTopPane = newValue
}
func (this *SwfExportOptions) GetShowBottomPane() *bool {
	return this.ShowBottomPane
}

func (this *SwfExportOptions) SetShowBottomPane(newValue *bool) {
	this.ShowBottomPane = newValue
}
func (this *SwfExportOptions) GetShowLeftPane() *bool {
	return this.ShowLeftPane
}

func (this *SwfExportOptions) SetShowLeftPane(newValue *bool) {
	this.ShowLeftPane = newValue
}
func (this *SwfExportOptions) GetStartOpenLeftPane() *bool {
	return this.StartOpenLeftPane
}

func (this *SwfExportOptions) SetStartOpenLeftPane(newValue *bool) {
	this.StartOpenLeftPane = newValue
}
func (this *SwfExportOptions) GetEnableContextMenu() *bool {
	return this.EnableContextMenu
}

func (this *SwfExportOptions) SetEnableContextMenu(newValue *bool) {
	this.EnableContextMenu = newValue
}
func (this *SwfExportOptions) GetLogoImage() string {
	return this.LogoImage
}

func (this *SwfExportOptions) SetLogoImage(newValue string) {
	this.LogoImage = newValue
}
func (this *SwfExportOptions) GetLogoLink() string {
	return this.LogoLink
}

func (this *SwfExportOptions) SetLogoLink(newValue string) {
	this.LogoLink = newValue
}
func (this *SwfExportOptions) GetJpegQuality() int32 {
	return this.JpegQuality
}

func (this *SwfExportOptions) SetJpegQuality(newValue int32) {
	this.JpegQuality = newValue
}
func (this *SwfExportOptions) GetSlidesLayoutOptions() ISlidesLayoutOptions {
	return this.SlidesLayoutOptions
}

func (this *SwfExportOptions) SetSlidesLayoutOptions(newValue ISlidesLayoutOptions) {
	this.SlidesLayoutOptions = newValue
}

func (this *SwfExportOptions) UnmarshalJSON(b []byte) error {
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
	
	if valCompressed, ok := GetMapValue(objMap, "compressed"); ok {
		if valCompressed != nil {
			var valueForCompressed *bool
			err = json.Unmarshal(*valCompressed, &valueForCompressed)
			if err != nil {
				return err
			}
			this.Compressed = valueForCompressed
		}
	}
	
	if valViewerIncluded, ok := GetMapValue(objMap, "viewerIncluded"); ok {
		if valViewerIncluded != nil {
			var valueForViewerIncluded *bool
			err = json.Unmarshal(*valViewerIncluded, &valueForViewerIncluded)
			if err != nil {
				return err
			}
			this.ViewerIncluded = valueForViewerIncluded
		}
	}
	
	if valShowPageBorder, ok := GetMapValue(objMap, "showPageBorder"); ok {
		if valShowPageBorder != nil {
			var valueForShowPageBorder *bool
			err = json.Unmarshal(*valShowPageBorder, &valueForShowPageBorder)
			if err != nil {
				return err
			}
			this.ShowPageBorder = valueForShowPageBorder
		}
	}
	
	if valShowFullScreen, ok := GetMapValue(objMap, "showFullScreen"); ok {
		if valShowFullScreen != nil {
			var valueForShowFullScreen *bool
			err = json.Unmarshal(*valShowFullScreen, &valueForShowFullScreen)
			if err != nil {
				return err
			}
			this.ShowFullScreen = valueForShowFullScreen
		}
	}
	
	if valShowPageStepper, ok := GetMapValue(objMap, "showPageStepper"); ok {
		if valShowPageStepper != nil {
			var valueForShowPageStepper *bool
			err = json.Unmarshal(*valShowPageStepper, &valueForShowPageStepper)
			if err != nil {
				return err
			}
			this.ShowPageStepper = valueForShowPageStepper
		}
	}
	
	if valShowSearch, ok := GetMapValue(objMap, "showSearch"); ok {
		if valShowSearch != nil {
			var valueForShowSearch *bool
			err = json.Unmarshal(*valShowSearch, &valueForShowSearch)
			if err != nil {
				return err
			}
			this.ShowSearch = valueForShowSearch
		}
	}
	
	if valShowTopPane, ok := GetMapValue(objMap, "showTopPane"); ok {
		if valShowTopPane != nil {
			var valueForShowTopPane *bool
			err = json.Unmarshal(*valShowTopPane, &valueForShowTopPane)
			if err != nil {
				return err
			}
			this.ShowTopPane = valueForShowTopPane
		}
	}
	
	if valShowBottomPane, ok := GetMapValue(objMap, "showBottomPane"); ok {
		if valShowBottomPane != nil {
			var valueForShowBottomPane *bool
			err = json.Unmarshal(*valShowBottomPane, &valueForShowBottomPane)
			if err != nil {
				return err
			}
			this.ShowBottomPane = valueForShowBottomPane
		}
	}
	
	if valShowLeftPane, ok := GetMapValue(objMap, "showLeftPane"); ok {
		if valShowLeftPane != nil {
			var valueForShowLeftPane *bool
			err = json.Unmarshal(*valShowLeftPane, &valueForShowLeftPane)
			if err != nil {
				return err
			}
			this.ShowLeftPane = valueForShowLeftPane
		}
	}
	
	if valStartOpenLeftPane, ok := GetMapValue(objMap, "startOpenLeftPane"); ok {
		if valStartOpenLeftPane != nil {
			var valueForStartOpenLeftPane *bool
			err = json.Unmarshal(*valStartOpenLeftPane, &valueForStartOpenLeftPane)
			if err != nil {
				return err
			}
			this.StartOpenLeftPane = valueForStartOpenLeftPane
		}
	}
	
	if valEnableContextMenu, ok := GetMapValue(objMap, "enableContextMenu"); ok {
		if valEnableContextMenu != nil {
			var valueForEnableContextMenu *bool
			err = json.Unmarshal(*valEnableContextMenu, &valueForEnableContextMenu)
			if err != nil {
				return err
			}
			this.EnableContextMenu = valueForEnableContextMenu
		}
	}
	
	if valLogoImage, ok := GetMapValue(objMap, "logoImage"); ok {
		if valLogoImage != nil {
			var valueForLogoImage string
			err = json.Unmarshal(*valLogoImage, &valueForLogoImage)
			if err != nil {
				return err
			}
			this.LogoImage = valueForLogoImage
		}
	}
	
	if valLogoLink, ok := GetMapValue(objMap, "logoLink"); ok {
		if valLogoLink != nil {
			var valueForLogoLink string
			err = json.Unmarshal(*valLogoLink, &valueForLogoLink)
			if err != nil {
				return err
			}
			this.LogoLink = valueForLogoLink
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

	return nil
}
