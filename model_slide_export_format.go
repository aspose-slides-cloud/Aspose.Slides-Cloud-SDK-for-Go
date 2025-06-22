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
	"strings"
)

// SlideExportFormat : Slide Export Format
type SlideExportFormat string

// List of SlideExportFormat SlideExportFormat
const (
	SlideExportFormat_Jpeg SlideExportFormat = "Jpeg"
	SlideExportFormat_Png SlideExportFormat = "Png"
	SlideExportFormat_Gif SlideExportFormat = "Gif"
	SlideExportFormat_Bmp SlideExportFormat = "Bmp"
	SlideExportFormat_Tiff SlideExportFormat = "Tiff"
	SlideExportFormat_Html SlideExportFormat = "Html"
	SlideExportFormat_Pdf SlideExportFormat = "Pdf"
	SlideExportFormat_Xps SlideExportFormat = "Xps"
	SlideExportFormat_Pptx SlideExportFormat = "Pptx"
	SlideExportFormat_Odp SlideExportFormat = "Odp"
	SlideExportFormat_Otp SlideExportFormat = "Otp"
	SlideExportFormat_Ppt SlideExportFormat = "Ppt"
	SlideExportFormat_Pps SlideExportFormat = "Pps"
	SlideExportFormat_Ppsx SlideExportFormat = "Ppsx"
	SlideExportFormat_Pptm SlideExportFormat = "Pptm"
	SlideExportFormat_Ppsm SlideExportFormat = "Ppsm"
	SlideExportFormat_Potx SlideExportFormat = "Potx"
	SlideExportFormat_Pot SlideExportFormat = "Pot"
	SlideExportFormat_Potm SlideExportFormat = "Potm"
	SlideExportFormat_Svg SlideExportFormat = "Svg"
	SlideExportFormat_Fodp SlideExportFormat = "Fodp"
	SlideExportFormat_Xaml SlideExportFormat = "Xaml"
	SlideExportFormat_Html5 SlideExportFormat = "Html5"
	SlideExportFormat_Md SlideExportFormat = "Md"
	SlideExportFormat_Xml SlideExportFormat = "Xml"
	SlideExportFormat_Emf SlideExportFormat = "Emf"
)

func SlideExportFormat_Validate(value string) (bool) {
	return strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Jpeg)) || strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Png)) || strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Gif)) || strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Bmp)) || strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Tiff)) || strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Html)) || strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Pdf)) || strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Xps)) || strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Pptx)) || strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Odp)) || strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Otp)) || strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Ppt)) || strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Pps)) || strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Ppsx)) || strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Pptm)) || strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Ppsm)) || strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Potx)) || strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Pot)) || strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Potm)) || strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Svg)) || strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Fodp)) || strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Xaml)) || strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Html5)) || strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Md)) || strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Xml)) || strings.ToLower(value) == strings.ToLower(string(SlideExportFormat_Emf))
}
