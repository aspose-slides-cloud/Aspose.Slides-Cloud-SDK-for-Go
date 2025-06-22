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

// ExportFormat : Available export formats
type ExportFormat string

// List of ExportFormat ExportFormat
const (
	ExportFormat_Pdf ExportFormat = "Pdf"
	ExportFormat_Xps ExportFormat = "Xps"
	ExportFormat_Tiff ExportFormat = "Tiff"
	ExportFormat_Pptx ExportFormat = "Pptx"
	ExportFormat_Odp ExportFormat = "Odp"
	ExportFormat_Otp ExportFormat = "Otp"
	ExportFormat_Ppt ExportFormat = "Ppt"
	ExportFormat_Pps ExportFormat = "Pps"
	ExportFormat_Ppsx ExportFormat = "Ppsx"
	ExportFormat_Pptm ExportFormat = "Pptm"
	ExportFormat_Ppsm ExportFormat = "Ppsm"
	ExportFormat_Pot ExportFormat = "Pot"
	ExportFormat_Potx ExportFormat = "Potx"
	ExportFormat_Potm ExportFormat = "Potm"
	ExportFormat_Html ExportFormat = "Html"
	ExportFormat_Html5 ExportFormat = "Html5"
	ExportFormat_Swf ExportFormat = "Swf"
	ExportFormat_Svg ExportFormat = "Svg"
	ExportFormat_Jpeg ExportFormat = "Jpeg"
	ExportFormat_Png ExportFormat = "Png"
	ExportFormat_Gif ExportFormat = "Gif"
	ExportFormat_Bmp ExportFormat = "Bmp"
	ExportFormat_Fodp ExportFormat = "Fodp"
	ExportFormat_Xaml ExportFormat = "Xaml"
	ExportFormat_Mpeg4 ExportFormat = "Mpeg4"
	ExportFormat_Md ExportFormat = "Md"
	ExportFormat_Xml ExportFormat = "Xml"
	ExportFormat_Emf ExportFormat = "Emf"
)

func ExportFormat_Validate(value string) (bool) {
	return strings.ToLower(value) == strings.ToLower(string(ExportFormat_Pdf)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Xps)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Tiff)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Pptx)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Odp)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Otp)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Ppt)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Pps)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Ppsx)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Pptm)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Ppsm)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Pot)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Potx)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Potm)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Html)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Html5)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Swf)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Svg)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Jpeg)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Png)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Gif)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Bmp)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Fodp)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Xaml)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Mpeg4)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Md)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Xml)) || strings.ToLower(value) == strings.ToLower(string(ExportFormat_Emf))
}
