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

// ShapeExportFormat : Represents a format for individual shape export.
type ShapeExportFormat string

// List of ShapeExportFormat ShapeExportFormat
const (
	ShapeExportFormat_Jpeg ShapeExportFormat = "Jpeg"
	ShapeExportFormat_Png ShapeExportFormat = "Png"
	ShapeExportFormat_Gif ShapeExportFormat = "Gif"
	ShapeExportFormat_Bmp ShapeExportFormat = "Bmp"
	ShapeExportFormat_Tiff ShapeExportFormat = "Tiff"
	ShapeExportFormat_Svg ShapeExportFormat = "Svg"
	ShapeExportFormat_Emf ShapeExportFormat = "Emf"
)

func ShapeExportFormat_Validate(value string) (bool) {
	return strings.ToLower(value) == strings.ToLower(string(ShapeExportFormat_Jpeg)) || strings.ToLower(value) == strings.ToLower(string(ShapeExportFormat_Png)) || strings.ToLower(value) == strings.ToLower(string(ShapeExportFormat_Gif)) || strings.ToLower(value) == strings.ToLower(string(ShapeExportFormat_Bmp)) || strings.ToLower(value) == strings.ToLower(string(ShapeExportFormat_Tiff)) || strings.ToLower(value) == strings.ToLower(string(ShapeExportFormat_Svg)) || strings.ToLower(value) == strings.ToLower(string(ShapeExportFormat_Emf))
}
