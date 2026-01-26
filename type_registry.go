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

import "encoding/json"
import "reflect"
import "unicode"

var typeRegistry = make(map[string]reflect.Type)
var derivedTypes = make(map[string]string)
var typeDeterminers = make(map[string]map[string]string)

func init() {
	typeRegistry["AccentElement"] = reflect.TypeOf(AccentElement{})
	derivedTypes["AccentElement"] = "MathElement"
	typeDeterminers["AccentElement"] = make(map[string]string)
	typeDeterminers["AccentElement"]["Type"] = "Accent"
	typeRegistry["AccessPermissions"] = reflect.TypeOf(AccessPermissions{})
	
	typeDeterminers["AccessPermissions"] = make(map[string]string)
	typeRegistry["AddLayoutSlide"] = reflect.TypeOf(AddLayoutSlide{})
	derivedTypes["AddLayoutSlide"] = "Task"
	typeDeterminers["AddLayoutSlide"] = make(map[string]string)
	typeDeterminers["AddLayoutSlide"]["Type"] = "AddLayoutSlide"
	typeRegistry["AddMasterSlide"] = reflect.TypeOf(AddMasterSlide{})
	derivedTypes["AddMasterSlide"] = "Task"
	typeDeterminers["AddMasterSlide"] = make(map[string]string)
	typeDeterminers["AddMasterSlide"]["Type"] = "AddMasterSlide"
	typeRegistry["AddShape"] = reflect.TypeOf(AddShape{})
	derivedTypes["AddShape"] = "Task"
	typeDeterminers["AddShape"] = make(map[string]string)
	typeDeterminers["AddShape"]["Type"] = "AddShape"
	typeRegistry["AddSlide"] = reflect.TypeOf(AddSlide{})
	derivedTypes["AddSlide"] = "Task"
	typeDeterminers["AddSlide"] = make(map[string]string)
	typeDeterminers["AddSlide"]["Type"] = "AddSlide"
	typeRegistry["AlphaBiLevelEffect"] = reflect.TypeOf(AlphaBiLevelEffect{})
	derivedTypes["AlphaBiLevelEffect"] = "ImageTransformEffect"
	typeDeterminers["AlphaBiLevelEffect"] = make(map[string]string)
	typeDeterminers["AlphaBiLevelEffect"]["Type"] = "AlphaBiLevel"
	typeRegistry["AlphaCeilingEffect"] = reflect.TypeOf(AlphaCeilingEffect{})
	derivedTypes["AlphaCeilingEffect"] = "ImageTransformEffect"
	typeDeterminers["AlphaCeilingEffect"] = make(map[string]string)
	typeDeterminers["AlphaCeilingEffect"]["Type"] = "AlphaCeiling"
	typeRegistry["AlphaFloorEffect"] = reflect.TypeOf(AlphaFloorEffect{})
	derivedTypes["AlphaFloorEffect"] = "ImageTransformEffect"
	typeDeterminers["AlphaFloorEffect"] = make(map[string]string)
	typeDeterminers["AlphaFloorEffect"]["Type"] = "AlphaFloor"
	typeRegistry["AlphaInverseEffect"] = reflect.TypeOf(AlphaInverseEffect{})
	derivedTypes["AlphaInverseEffect"] = "ImageTransformEffect"
	typeDeterminers["AlphaInverseEffect"] = make(map[string]string)
	typeDeterminers["AlphaInverseEffect"]["Type"] = "AlphaInverse"
	typeRegistry["AlphaModulateEffect"] = reflect.TypeOf(AlphaModulateEffect{})
	derivedTypes["AlphaModulateEffect"] = "ImageTransformEffect"
	typeDeterminers["AlphaModulateEffect"] = make(map[string]string)
	typeDeterminers["AlphaModulateEffect"]["Type"] = "AlphaModulate"
	typeRegistry["AlphaModulateFixedEffect"] = reflect.TypeOf(AlphaModulateFixedEffect{})
	derivedTypes["AlphaModulateFixedEffect"] = "ImageTransformEffect"
	typeDeterminers["AlphaModulateFixedEffect"] = make(map[string]string)
	typeDeterminers["AlphaModulateFixedEffect"]["Type"] = "AlphaModulateFixed"
	typeRegistry["AlphaReplaceEffect"] = reflect.TypeOf(AlphaReplaceEffect{})
	derivedTypes["AlphaReplaceEffect"] = "ImageTransformEffect"
	typeDeterminers["AlphaReplaceEffect"] = make(map[string]string)
	typeDeterminers["AlphaReplaceEffect"]["Type"] = "AlphaReplace"
	typeRegistry["ApiInfo"] = reflect.TypeOf(ApiInfo{})
	
	typeDeterminers["ApiInfo"] = make(map[string]string)
	typeRegistry["ArcToPathSegment"] = reflect.TypeOf(ArcToPathSegment{})
	derivedTypes["ArcToPathSegment"] = "PathSegment"
	typeDeterminers["ArcToPathSegment"] = make(map[string]string)
	typeDeterminers["ArcToPathSegment"]["Type"] = "ArcTo"
	typeRegistry["ArrayElement"] = reflect.TypeOf(ArrayElement{})
	derivedTypes["ArrayElement"] = "MathElement"
	typeDeterminers["ArrayElement"] = make(map[string]string)
	typeDeterminers["ArrayElement"]["Type"] = "Array"
	typeRegistry["ArrowHeadProperties"] = reflect.TypeOf(ArrowHeadProperties{})
	
	typeDeterminers["ArrowHeadProperties"] = make(map[string]string)
	typeRegistry["AudioFrame"] = reflect.TypeOf(AudioFrame{})
	derivedTypes["AudioFrame"] = "GeometryShape"
	typeDeterminers["AudioFrame"] = make(map[string]string)
	typeDeterminers["AudioFrame"]["Type"] = "AudioFrame"
	typeRegistry["Axes"] = reflect.TypeOf(Axes{})
	
	typeDeterminers["Axes"] = make(map[string]string)
	typeRegistry["Axis"] = reflect.TypeOf(Axis{})
	
	typeDeterminers["Axis"] = make(map[string]string)
	typeRegistry["BarElement"] = reflect.TypeOf(BarElement{})
	derivedTypes["BarElement"] = "MathElement"
	typeDeterminers["BarElement"] = make(map[string]string)
	typeDeterminers["BarElement"]["Type"] = "Bar"
	typeRegistry["Base64InputFile"] = reflect.TypeOf(Base64InputFile{})
	derivedTypes["Base64InputFile"] = "InputFile"
	typeDeterminers["Base64InputFile"] = make(map[string]string)
	typeDeterminers["Base64InputFile"]["Type"] = "Base64"
	typeRegistry["BiLevelEffect"] = reflect.TypeOf(BiLevelEffect{})
	derivedTypes["BiLevelEffect"] = "ImageTransformEffect"
	typeDeterminers["BiLevelEffect"] = make(map[string]string)
	typeDeterminers["BiLevelEffect"]["Type"] = "BiLevel"
	typeRegistry["BlockElement"] = reflect.TypeOf(BlockElement{})
	derivedTypes["BlockElement"] = "MathElement"
	typeDeterminers["BlockElement"] = make(map[string]string)
	typeDeterminers["BlockElement"]["Type"] = "Block"
	typeRegistry["BlurEffect"] = reflect.TypeOf(BlurEffect{})
	
	typeDeterminers["BlurEffect"] = make(map[string]string)
	typeRegistry["BlurImageEffect"] = reflect.TypeOf(BlurImageEffect{})
	derivedTypes["BlurImageEffect"] = "ImageTransformEffect"
	typeDeterminers["BlurImageEffect"] = make(map[string]string)
	typeDeterminers["BlurImageEffect"]["Type"] = "Blur"
	typeRegistry["BorderBoxElement"] = reflect.TypeOf(BorderBoxElement{})
	derivedTypes["BorderBoxElement"] = "MathElement"
	typeDeterminers["BorderBoxElement"] = make(map[string]string)
	typeDeterminers["BorderBoxElement"]["Type"] = "BorderBox"
	typeRegistry["BoxElement"] = reflect.TypeOf(BoxElement{})
	derivedTypes["BoxElement"] = "MathElement"
	typeDeterminers["BoxElement"] = make(map[string]string)
	typeDeterminers["BoxElement"]["Type"] = "Box"
	typeRegistry["BubbleChartDataPoint"] = reflect.TypeOf(BubbleChartDataPoint{})
	derivedTypes["BubbleChartDataPoint"] = "ScatterChartDataPoint"
	typeDeterminers["BubbleChartDataPoint"] = make(map[string]string)
	typeDeterminers["BubbleChartDataPoint"]["Type"] = "Bubble"
	typeRegistry["BubbleSeries"] = reflect.TypeOf(BubbleSeries{})
	derivedTypes["BubbleSeries"] = "XYSeries"
	typeDeterminers["BubbleSeries"] = make(map[string]string)
	typeDeterminers["BubbleSeries"]["DataPointType"] = "Bubble"
	typeRegistry["Camera"] = reflect.TypeOf(Camera{})
	
	typeDeterminers["Camera"] = make(map[string]string)
	typeRegistry["CaptionTrack"] = reflect.TypeOf(CaptionTrack{})
	derivedTypes["CaptionTrack"] = "ResourceBase"
	typeDeterminers["CaptionTrack"] = make(map[string]string)
	typeRegistry["CaptionTracks"] = reflect.TypeOf(CaptionTracks{})
	derivedTypes["CaptionTracks"] = "ResourceBase"
	typeDeterminers["CaptionTracks"] = make(map[string]string)
	typeRegistry["Chart"] = reflect.TypeOf(Chart{})
	derivedTypes["Chart"] = "ShapeBase"
	typeDeterminers["Chart"] = make(map[string]string)
	typeDeterminers["Chart"]["Type"] = "Chart"
	typeRegistry["ChartCategory"] = reflect.TypeOf(ChartCategory{})
	
	typeDeterminers["ChartCategory"] = make(map[string]string)
	typeRegistry["ChartLinesFormat"] = reflect.TypeOf(ChartLinesFormat{})
	
	typeDeterminers["ChartLinesFormat"] = make(map[string]string)
	typeRegistry["ChartSeriesGroup"] = reflect.TypeOf(ChartSeriesGroup{})
	
	typeDeterminers["ChartSeriesGroup"] = make(map[string]string)
	typeRegistry["ChartTitle"] = reflect.TypeOf(ChartTitle{})
	
	typeDeterminers["ChartTitle"] = make(map[string]string)
	typeRegistry["ChartWall"] = reflect.TypeOf(ChartWall{})
	
	typeDeterminers["ChartWall"] = make(map[string]string)
	typeRegistry["ClosePathSegment"] = reflect.TypeOf(ClosePathSegment{})
	derivedTypes["ClosePathSegment"] = "PathSegment"
	typeDeterminers["ClosePathSegment"] = make(map[string]string)
	typeDeterminers["ClosePathSegment"]["Type"] = "Close"
	typeRegistry["ColorChangeEffect"] = reflect.TypeOf(ColorChangeEffect{})
	derivedTypes["ColorChangeEffect"] = "ImageTransformEffect"
	typeDeterminers["ColorChangeEffect"] = make(map[string]string)
	typeDeterminers["ColorChangeEffect"]["Type"] = "ColorChange"
	typeRegistry["ColorReplaceEffect"] = reflect.TypeOf(ColorReplaceEffect{})
	derivedTypes["ColorReplaceEffect"] = "ImageTransformEffect"
	typeDeterminers["ColorReplaceEffect"] = make(map[string]string)
	typeDeterminers["ColorReplaceEffect"]["Type"] = "ColorReplace"
	typeRegistry["ColorScheme"] = reflect.TypeOf(ColorScheme{})
	derivedTypes["ColorScheme"] = "ResourceBase"
	typeDeterminers["ColorScheme"] = make(map[string]string)
	typeRegistry["CommentAuthor"] = reflect.TypeOf(CommentAuthor{})
	
	typeDeterminers["CommentAuthor"] = make(map[string]string)
	typeRegistry["CommentAuthors"] = reflect.TypeOf(CommentAuthors{})
	derivedTypes["CommentAuthors"] = "ResourceBase"
	typeDeterminers["CommentAuthors"] = make(map[string]string)
	typeRegistry["CommonSlideViewProperties"] = reflect.TypeOf(CommonSlideViewProperties{})
	
	typeDeterminers["CommonSlideViewProperties"] = make(map[string]string)
	typeRegistry["Connector"] = reflect.TypeOf(Connector{})
	derivedTypes["Connector"] = "GeometryShape"
	typeDeterminers["Connector"] = make(map[string]string)
	typeDeterminers["Connector"]["Type"] = "Connector"
	typeRegistry["CubicBezierToPathSegment"] = reflect.TypeOf(CubicBezierToPathSegment{})
	derivedTypes["CubicBezierToPathSegment"] = "PathSegment"
	typeDeterminers["CubicBezierToPathSegment"] = make(map[string]string)
	typeDeterminers["CubicBezierToPathSegment"]["Type"] = "CubicBezierTo"
	typeRegistry["CustomDashPattern"] = reflect.TypeOf(CustomDashPattern{})
	
	typeDeterminers["CustomDashPattern"] = make(map[string]string)
	typeRegistry["DataPoint"] = reflect.TypeOf(DataPoint{})
	
	typeDeterminers["DataPoint"] = make(map[string]string)
	typeRegistry["DataSource"] = reflect.TypeOf(DataSource{})
	
	typeDeterminers["DataSource"] = make(map[string]string)
	typeRegistry["DelimiterElement"] = reflect.TypeOf(DelimiterElement{})
	derivedTypes["DelimiterElement"] = "MathElement"
	typeDeterminers["DelimiterElement"] = make(map[string]string)
	typeDeterminers["DelimiterElement"]["Type"] = "Delimiter"
	typeRegistry["DiscUsage"] = reflect.TypeOf(DiscUsage{})
	
	typeDeterminers["DiscUsage"] = make(map[string]string)
	typeRegistry["Document"] = reflect.TypeOf(Document{})
	derivedTypes["Document"] = "ResourceBase"
	typeDeterminers["Document"] = make(map[string]string)
	typeRegistry["DocumentProperties"] = reflect.TypeOf(DocumentProperties{})
	derivedTypes["DocumentProperties"] = "ResourceBase"
	typeDeterminers["DocumentProperties"] = make(map[string]string)
	typeRegistry["DocumentProperty"] = reflect.TypeOf(DocumentProperty{})
	derivedTypes["DocumentProperty"] = "ResourceBase"
	typeDeterminers["DocumentProperty"] = make(map[string]string)
	typeRegistry["DocumentReplaceResult"] = reflect.TypeOf(DocumentReplaceResult{})
	derivedTypes["DocumentReplaceResult"] = "Document"
	typeDeterminers["DocumentReplaceResult"] = make(map[string]string)
	typeRegistry["DrawingGuide"] = reflect.TypeOf(DrawingGuide{})
	
	typeDeterminers["DrawingGuide"] = make(map[string]string)
	typeRegistry["DuotoneEffect"] = reflect.TypeOf(DuotoneEffect{})
	derivedTypes["DuotoneEffect"] = "ImageTransformEffect"
	typeDeterminers["DuotoneEffect"] = make(map[string]string)
	typeDeterminers["DuotoneEffect"]["Type"] = "Duotone"
	typeRegistry["Effect"] = reflect.TypeOf(Effect{})
	
	typeDeterminers["Effect"] = make(map[string]string)
	typeRegistry["EffectFormat"] = reflect.TypeOf(EffectFormat{})
	
	typeDeterminers["EffectFormat"] = make(map[string]string)
	typeRegistry["EntityExists"] = reflect.TypeOf(EntityExists{})
	
	typeDeterminers["EntityExists"] = make(map[string]string)
	typeRegistry["ErrorDetails"] = reflect.TypeOf(ErrorDetails{})
	
	typeDeterminers["ErrorDetails"] = make(map[string]string)
	typeRegistry["ExportOptions"] = reflect.TypeOf(ExportOptions{})
	
	typeDeterminers["ExportOptions"] = make(map[string]string)
	typeRegistry["FileVersion"] = reflect.TypeOf(FileVersion{})
	derivedTypes["FileVersion"] = "StorageFile"
	typeDeterminers["FileVersion"] = make(map[string]string)
	typeRegistry["FileVersions"] = reflect.TypeOf(FileVersions{})
	
	typeDeterminers["FileVersions"] = make(map[string]string)
	typeRegistry["FilesList"] = reflect.TypeOf(FilesList{})
	
	typeDeterminers["FilesList"] = make(map[string]string)
	typeRegistry["FilesUploadResult"] = reflect.TypeOf(FilesUploadResult{})
	
	typeDeterminers["FilesUploadResult"] = make(map[string]string)
	typeRegistry["FillFormat"] = reflect.TypeOf(FillFormat{})
	
	typeDeterminers["FillFormat"] = make(map[string]string)
	typeRegistry["FillOverlayEffect"] = reflect.TypeOf(FillOverlayEffect{})
	
	typeDeterminers["FillOverlayEffect"] = make(map[string]string)
	typeRegistry["FillOverlayImageEffect"] = reflect.TypeOf(FillOverlayImageEffect{})
	derivedTypes["FillOverlayImageEffect"] = "ImageTransformEffect"
	typeDeterminers["FillOverlayImageEffect"] = make(map[string]string)
	typeDeterminers["FillOverlayImageEffect"]["Type"] = "FillOverlay"
	typeRegistry["FontData"] = reflect.TypeOf(FontData{})
	
	typeDeterminers["FontData"] = make(map[string]string)
	typeRegistry["FontFallbackRule"] = reflect.TypeOf(FontFallbackRule{})
	
	typeDeterminers["FontFallbackRule"] = make(map[string]string)
	typeRegistry["FontScheme"] = reflect.TypeOf(FontScheme{})
	derivedTypes["FontScheme"] = "ResourceBase"
	typeDeterminers["FontScheme"] = make(map[string]string)
	typeRegistry["FontSet"] = reflect.TypeOf(FontSet{})
	
	typeDeterminers["FontSet"] = make(map[string]string)
	typeRegistry["FontSubstRule"] = reflect.TypeOf(FontSubstRule{})
	
	typeDeterminers["FontSubstRule"] = make(map[string]string)
	typeRegistry["FontsData"] = reflect.TypeOf(FontsData{})
	
	typeDeterminers["FontsData"] = make(map[string]string)
	typeRegistry["FormatScheme"] = reflect.TypeOf(FormatScheme{})
	derivedTypes["FormatScheme"] = "ResourceBase"
	typeDeterminers["FormatScheme"] = make(map[string]string)
	typeRegistry["FractionElement"] = reflect.TypeOf(FractionElement{})
	derivedTypes["FractionElement"] = "MathElement"
	typeDeterminers["FractionElement"] = make(map[string]string)
	typeDeterminers["FractionElement"]["Type"] = "Fraction"
	typeRegistry["FunctionElement"] = reflect.TypeOf(FunctionElement{})
	derivedTypes["FunctionElement"] = "MathElement"
	typeDeterminers["FunctionElement"] = make(map[string]string)
	typeDeterminers["FunctionElement"]["Type"] = "Function"
	typeRegistry["GeometryPath"] = reflect.TypeOf(GeometryPath{})
	
	typeDeterminers["GeometryPath"] = make(map[string]string)
	typeRegistry["GeometryPaths"] = reflect.TypeOf(GeometryPaths{})
	
	typeDeterminers["GeometryPaths"] = make(map[string]string)
	typeRegistry["GeometryShape"] = reflect.TypeOf(GeometryShape{})
	derivedTypes["GeometryShape"] = "ShapeBase"
	typeDeterminers["GeometryShape"] = make(map[string]string)
	typeRegistry["GifExportOptions"] = reflect.TypeOf(GifExportOptions{})
	derivedTypes["GifExportOptions"] = "ImageExportOptionsBase"
	typeDeterminers["GifExportOptions"] = make(map[string]string)
	typeRegistry["GlowEffect"] = reflect.TypeOf(GlowEffect{})
	
	typeDeterminers["GlowEffect"] = make(map[string]string)
	typeRegistry["GradientFill"] = reflect.TypeOf(GradientFill{})
	derivedTypes["GradientFill"] = "FillFormat"
	typeDeterminers["GradientFill"] = make(map[string]string)
	typeDeterminers["GradientFill"]["Type"] = "Gradient"
	typeRegistry["GradientFillStop"] = reflect.TypeOf(GradientFillStop{})
	
	typeDeterminers["GradientFillStop"] = make(map[string]string)
	typeRegistry["GraphicalObject"] = reflect.TypeOf(GraphicalObject{})
	derivedTypes["GraphicalObject"] = "ShapeBase"
	typeDeterminers["GraphicalObject"] = make(map[string]string)
	typeDeterminers["GraphicalObject"]["Type"] = "GraphicalObject"
	typeRegistry["GrayScaleEffect"] = reflect.TypeOf(GrayScaleEffect{})
	derivedTypes["GrayScaleEffect"] = "ImageTransformEffect"
	typeDeterminers["GrayScaleEffect"] = make(map[string]string)
	typeDeterminers["GrayScaleEffect"]["Type"] = "GrayScale"
	typeRegistry["GroupShape"] = reflect.TypeOf(GroupShape{})
	derivedTypes["GroupShape"] = "ShapeBase"
	typeDeterminers["GroupShape"] = make(map[string]string)
	typeDeterminers["GroupShape"]["Type"] = "GroupShape"
	typeRegistry["GroupingCharacterElement"] = reflect.TypeOf(GroupingCharacterElement{})
	derivedTypes["GroupingCharacterElement"] = "MathElement"
	typeDeterminers["GroupingCharacterElement"] = make(map[string]string)
	typeDeterminers["GroupingCharacterElement"]["Type"] = "GroupingCharacter"
	typeRegistry["HandoutLayoutingOptions"] = reflect.TypeOf(HandoutLayoutingOptions{})
	derivedTypes["HandoutLayoutingOptions"] = "SlidesLayoutOptions"
	typeDeterminers["HandoutLayoutingOptions"] = make(map[string]string)
	typeDeterminers["HandoutLayoutingOptions"]["LayoutType"] = "Handout"
	typeRegistry["HeaderFooter"] = reflect.TypeOf(HeaderFooter{})
	derivedTypes["HeaderFooter"] = "ResourceBase"
	typeDeterminers["HeaderFooter"] = make(map[string]string)
	typeRegistry["HslEffect"] = reflect.TypeOf(HslEffect{})
	derivedTypes["HslEffect"] = "ImageTransformEffect"
	typeDeterminers["HslEffect"] = make(map[string]string)
	typeDeterminers["HslEffect"]["Type"] = "Hsl"
	typeRegistry["Html5ExportOptions"] = reflect.TypeOf(Html5ExportOptions{})
	derivedTypes["Html5ExportOptions"] = "ExportOptions"
	typeDeterminers["Html5ExportOptions"] = make(map[string]string)
	typeRegistry["HtmlExportOptions"] = reflect.TypeOf(HtmlExportOptions{})
	derivedTypes["HtmlExportOptions"] = "ExportOptions"
	typeDeterminers["HtmlExportOptions"] = make(map[string]string)
	typeRegistry["Hyperlink"] = reflect.TypeOf(Hyperlink{})
	
	typeDeterminers["Hyperlink"] = make(map[string]string)
	typeRegistry["IShapeExportOptions"] = reflect.TypeOf(IShapeExportOptions{})
	
	typeDeterminers["IShapeExportOptions"] = make(map[string]string)
	typeRegistry["Image"] = reflect.TypeOf(Image{})
	derivedTypes["Image"] = "ResourceBase"
	typeDeterminers["Image"] = make(map[string]string)
	typeRegistry["ImageExportOptions"] = reflect.TypeOf(ImageExportOptions{})
	derivedTypes["ImageExportOptions"] = "ImageExportOptionsBase"
	typeDeterminers["ImageExportOptions"] = make(map[string]string)
	typeRegistry["ImageExportOptionsBase"] = reflect.TypeOf(ImageExportOptionsBase{})
	derivedTypes["ImageExportOptionsBase"] = "ExportOptions"
	typeDeterminers["ImageExportOptionsBase"] = make(map[string]string)
	typeRegistry["ImageTransformEffect"] = reflect.TypeOf(ImageTransformEffect{})
	
	typeDeterminers["ImageTransformEffect"] = make(map[string]string)
	typeRegistry["Images"] = reflect.TypeOf(Images{})
	derivedTypes["Images"] = "ResourceBase"
	typeDeterminers["Images"] = make(map[string]string)
	typeRegistry["InnerShadowEffect"] = reflect.TypeOf(InnerShadowEffect{})
	
	typeDeterminers["InnerShadowEffect"] = make(map[string]string)
	typeRegistry["Input"] = reflect.TypeOf(Input{})
	
	typeDeterminers["Input"] = make(map[string]string)
	typeRegistry["InputFile"] = reflect.TypeOf(InputFile{})
	
	typeDeterminers["InputFile"] = make(map[string]string)
	typeRegistry["InteractiveSequence"] = reflect.TypeOf(InteractiveSequence{})
	
	typeDeterminers["InteractiveSequence"] = make(map[string]string)
	typeRegistry["LayoutSlide"] = reflect.TypeOf(LayoutSlide{})
	derivedTypes["LayoutSlide"] = "ResourceBase"
	typeDeterminers["LayoutSlide"] = make(map[string]string)
	typeRegistry["LayoutSlides"] = reflect.TypeOf(LayoutSlides{})
	derivedTypes["LayoutSlides"] = "ResourceBase"
	typeDeterminers["LayoutSlides"] = make(map[string]string)
	typeRegistry["LeftSubSuperscriptElement"] = reflect.TypeOf(LeftSubSuperscriptElement{})
	derivedTypes["LeftSubSuperscriptElement"] = "MathElement"
	typeDeterminers["LeftSubSuperscriptElement"] = make(map[string]string)
	typeDeterminers["LeftSubSuperscriptElement"]["Type"] = "LeftSubSuperscriptElement"
	typeRegistry["Legend"] = reflect.TypeOf(Legend{})
	
	typeDeterminers["Legend"] = make(map[string]string)
	typeRegistry["LightRig"] = reflect.TypeOf(LightRig{})
	
	typeDeterminers["LightRig"] = make(map[string]string)
	typeRegistry["LimitElement"] = reflect.TypeOf(LimitElement{})
	derivedTypes["LimitElement"] = "MathElement"
	typeDeterminers["LimitElement"] = make(map[string]string)
	typeDeterminers["LimitElement"]["Type"] = "Limit"
	typeRegistry["LineFormat"] = reflect.TypeOf(LineFormat{})
	
	typeDeterminers["LineFormat"] = make(map[string]string)
	typeRegistry["LineToPathSegment"] = reflect.TypeOf(LineToPathSegment{})
	derivedTypes["LineToPathSegment"] = "PathSegment"
	typeDeterminers["LineToPathSegment"] = make(map[string]string)
	typeDeterminers["LineToPathSegment"]["Type"] = "LineTo"
	typeRegistry["Literals"] = reflect.TypeOf(Literals{})
	derivedTypes["Literals"] = "DataSource"
	typeDeterminers["Literals"] = make(map[string]string)
	typeDeterminers["Literals"]["Type"] = "Literals"
	typeRegistry["LuminanceEffect"] = reflect.TypeOf(LuminanceEffect{})
	derivedTypes["LuminanceEffect"] = "ImageTransformEffect"
	typeDeterminers["LuminanceEffect"] = make(map[string]string)
	typeDeterminers["LuminanceEffect"]["Type"] = "Luminance"
	typeRegistry["MarkdownExportOptions"] = reflect.TypeOf(MarkdownExportOptions{})
	derivedTypes["MarkdownExportOptions"] = "ExportOptions"
	typeDeterminers["MarkdownExportOptions"] = make(map[string]string)
	typeRegistry["MasterSlide"] = reflect.TypeOf(MasterSlide{})
	derivedTypes["MasterSlide"] = "ResourceBase"
	typeDeterminers["MasterSlide"] = make(map[string]string)
	typeRegistry["MasterSlides"] = reflect.TypeOf(MasterSlides{})
	derivedTypes["MasterSlides"] = "ResourceBase"
	typeDeterminers["MasterSlides"] = make(map[string]string)
	typeRegistry["MathElement"] = reflect.TypeOf(MathElement{})
	
	typeDeterminers["MathElement"] = make(map[string]string)
	typeRegistry["MathParagraph"] = reflect.TypeOf(MathParagraph{})
	
	typeDeterminers["MathParagraph"] = make(map[string]string)
	typeRegistry["MatrixElement"] = reflect.TypeOf(MatrixElement{})
	derivedTypes["MatrixElement"] = "MathElement"
	typeDeterminers["MatrixElement"] = make(map[string]string)
	typeDeterminers["MatrixElement"]["Type"] = "Matrix"
	typeRegistry["Merge"] = reflect.TypeOf(Merge{})
	derivedTypes["Merge"] = "Task"
	typeDeterminers["Merge"] = make(map[string]string)
	typeDeterminers["Merge"]["Type"] = "Merge"
	typeRegistry["MergingSource"] = reflect.TypeOf(MergingSource{})
	
	typeDeterminers["MergingSource"] = make(map[string]string)
	typeRegistry["ModelError"] = reflect.TypeOf(ModelError{})
	
	typeDeterminers["ModelError"] = make(map[string]string)
	typeRegistry["MoveToPathSegment"] = reflect.TypeOf(MoveToPathSegment{})
	derivedTypes["MoveToPathSegment"] = "PathSegment"
	typeDeterminers["MoveToPathSegment"] = make(map[string]string)
	typeDeterminers["MoveToPathSegment"]["Type"] = "MoveTo"
	typeRegistry["NaryOperatorElement"] = reflect.TypeOf(NaryOperatorElement{})
	derivedTypes["NaryOperatorElement"] = "MathElement"
	typeDeterminers["NaryOperatorElement"] = make(map[string]string)
	typeDeterminers["NaryOperatorElement"]["Type"] = "NaryOperator"
	typeRegistry["NoFill"] = reflect.TypeOf(NoFill{})
	derivedTypes["NoFill"] = "FillFormat"
	typeDeterminers["NoFill"] = make(map[string]string)
	typeDeterminers["NoFill"]["Type"] = "NoFill"
	typeRegistry["NormalViewRestoredProperties"] = reflect.TypeOf(NormalViewRestoredProperties{})
	
	typeDeterminers["NormalViewRestoredProperties"] = make(map[string]string)
	typeRegistry["NotesCommentsLayoutingOptions"] = reflect.TypeOf(NotesCommentsLayoutingOptions{})
	derivedTypes["NotesCommentsLayoutingOptions"] = "SlidesLayoutOptions"
	typeDeterminers["NotesCommentsLayoutingOptions"] = make(map[string]string)
	typeDeterminers["NotesCommentsLayoutingOptions"]["LayoutType"] = "NotesComments"
	typeRegistry["NotesSlide"] = reflect.TypeOf(NotesSlide{})
	derivedTypes["NotesSlide"] = "ResourceBase"
	typeDeterminers["NotesSlide"] = make(map[string]string)
	typeRegistry["NotesSlideHeaderFooter"] = reflect.TypeOf(NotesSlideHeaderFooter{})
	derivedTypes["NotesSlideHeaderFooter"] = "ResourceBase"
	typeDeterminers["NotesSlideHeaderFooter"] = make(map[string]string)
	typeRegistry["ObjectExist"] = reflect.TypeOf(ObjectExist{})
	
	typeDeterminers["ObjectExist"] = make(map[string]string)
	typeRegistry["OleObjectFrame"] = reflect.TypeOf(OleObjectFrame{})
	derivedTypes["OleObjectFrame"] = "ShapeBase"
	typeDeterminers["OleObjectFrame"] = make(map[string]string)
	typeDeterminers["OleObjectFrame"]["Type"] = "OleObjectFrame"
	typeRegistry["OneValueChartDataPoint"] = reflect.TypeOf(OneValueChartDataPoint{})
	derivedTypes["OneValueChartDataPoint"] = "DataPoint"
	typeDeterminers["OneValueChartDataPoint"] = make(map[string]string)
	typeDeterminers["OneValueChartDataPoint"]["Type"] = "OneValue"
	typeRegistry["OneValueSeries"] = reflect.TypeOf(OneValueSeries{})
	derivedTypes["OneValueSeries"] = "Series"
	typeDeterminers["OneValueSeries"] = make(map[string]string)
	typeDeterminers["OneValueSeries"]["DataPointType"] = "OneValue"
	typeRegistry["Operation"] = reflect.TypeOf(Operation{})
	
	typeDeterminers["Operation"] = make(map[string]string)
	typeRegistry["OperationError"] = reflect.TypeOf(OperationError{})
	
	typeDeterminers["OperationError"] = make(map[string]string)
	typeRegistry["OperationProgress"] = reflect.TypeOf(OperationProgress{})
	
	typeDeterminers["OperationProgress"] = make(map[string]string)
	typeRegistry["OrderedMergeRequest"] = reflect.TypeOf(OrderedMergeRequest{})
	
	typeDeterminers["OrderedMergeRequest"] = make(map[string]string)
	typeRegistry["OuterShadowEffect"] = reflect.TypeOf(OuterShadowEffect{})
	
	typeDeterminers["OuterShadowEffect"] = make(map[string]string)
	typeRegistry["OutputFile"] = reflect.TypeOf(OutputFile{})
	
	typeDeterminers["OutputFile"] = make(map[string]string)
	typeRegistry["Paragraph"] = reflect.TypeOf(Paragraph{})
	derivedTypes["Paragraph"] = "ResourceBase"
	typeDeterminers["Paragraph"] = make(map[string]string)
	typeRegistry["ParagraphFormat"] = reflect.TypeOf(ParagraphFormat{})
	
	typeDeterminers["ParagraphFormat"] = make(map[string]string)
	typeRegistry["Paragraphs"] = reflect.TypeOf(Paragraphs{})
	derivedTypes["Paragraphs"] = "ResourceBase"
	typeDeterminers["Paragraphs"] = make(map[string]string)
	typeRegistry["PathInputFile"] = reflect.TypeOf(PathInputFile{})
	derivedTypes["PathInputFile"] = "InputFile"
	typeDeterminers["PathInputFile"] = make(map[string]string)
	typeDeterminers["PathInputFile"]["Type"] = "Path"
	typeRegistry["PathOutputFile"] = reflect.TypeOf(PathOutputFile{})
	derivedTypes["PathOutputFile"] = "OutputFile"
	typeDeterminers["PathOutputFile"] = make(map[string]string)
	typeDeterminers["PathOutputFile"]["Type"] = "Path"
	typeRegistry["PathSegment"] = reflect.TypeOf(PathSegment{})
	
	typeDeterminers["PathSegment"] = make(map[string]string)
	typeRegistry["PatternFill"] = reflect.TypeOf(PatternFill{})
	derivedTypes["PatternFill"] = "FillFormat"
	typeDeterminers["PatternFill"] = make(map[string]string)
	typeDeterminers["PatternFill"]["Type"] = "Pattern"
	typeRegistry["PdfExportOptions"] = reflect.TypeOf(PdfExportOptions{})
	derivedTypes["PdfExportOptions"] = "ExportOptions"
	typeDeterminers["PdfExportOptions"] = make(map[string]string)
	typeRegistry["PdfImportOptions"] = reflect.TypeOf(PdfImportOptions{})
	
	typeDeterminers["PdfImportOptions"] = make(map[string]string)
	typeRegistry["PhantomElement"] = reflect.TypeOf(PhantomElement{})
	derivedTypes["PhantomElement"] = "MathElement"
	typeDeterminers["PhantomElement"] = make(map[string]string)
	typeDeterminers["PhantomElement"]["Type"] = "Phantom"
	typeRegistry["PictureFill"] = reflect.TypeOf(PictureFill{})
	derivedTypes["PictureFill"] = "FillFormat"
	typeDeterminers["PictureFill"] = make(map[string]string)
	typeDeterminers["PictureFill"]["Type"] = "Picture"
	typeRegistry["PictureFrame"] = reflect.TypeOf(PictureFrame{})
	derivedTypes["PictureFrame"] = "GeometryShape"
	typeDeterminers["PictureFrame"] = make(map[string]string)
	typeDeterminers["PictureFrame"]["Type"] = "PictureFrame"
	typeRegistry["Pipeline"] = reflect.TypeOf(Pipeline{})
	
	typeDeterminers["Pipeline"] = make(map[string]string)
	typeRegistry["Placeholder"] = reflect.TypeOf(Placeholder{})
	derivedTypes["Placeholder"] = "ResourceBase"
	typeDeterminers["Placeholder"] = make(map[string]string)
	typeRegistry["Placeholders"] = reflect.TypeOf(Placeholders{})
	derivedTypes["Placeholders"] = "ResourceBase"
	typeDeterminers["Placeholders"] = make(map[string]string)
	typeRegistry["PlotArea"] = reflect.TypeOf(PlotArea{})
	
	typeDeterminers["PlotArea"] = make(map[string]string)
	typeRegistry["Portion"] = reflect.TypeOf(Portion{})
	derivedTypes["Portion"] = "ResourceBase"
	typeDeterminers["Portion"] = make(map[string]string)
	typeRegistry["PortionFormat"] = reflect.TypeOf(PortionFormat{})
	
	typeDeterminers["PortionFormat"] = make(map[string]string)
	typeRegistry["Portions"] = reflect.TypeOf(Portions{})
	derivedTypes["Portions"] = "ResourceBase"
	typeDeterminers["Portions"] = make(map[string]string)
	typeRegistry["PptxExportOptions"] = reflect.TypeOf(PptxExportOptions{})
	derivedTypes["PptxExportOptions"] = "ExportOptions"
	typeDeterminers["PptxExportOptions"] = make(map[string]string)
	typeRegistry["PresentationToMerge"] = reflect.TypeOf(PresentationToMerge{})
	
	typeDeterminers["PresentationToMerge"] = make(map[string]string)
	typeRegistry["PresentationsMergeRequest"] = reflect.TypeOf(PresentationsMergeRequest{})
	
	typeDeterminers["PresentationsMergeRequest"] = make(map[string]string)
	typeRegistry["PresetShadowEffect"] = reflect.TypeOf(PresetShadowEffect{})
	
	typeDeterminers["PresetShadowEffect"] = make(map[string]string)
	typeRegistry["ProtectionProperties"] = reflect.TypeOf(ProtectionProperties{})
	derivedTypes["ProtectionProperties"] = "ResourceBase"
	typeDeterminers["ProtectionProperties"] = make(map[string]string)
	typeRegistry["QuadraticBezierToPathSegment"] = reflect.TypeOf(QuadraticBezierToPathSegment{})
	derivedTypes["QuadraticBezierToPathSegment"] = "PathSegment"
	typeDeterminers["QuadraticBezierToPathSegment"] = make(map[string]string)
	typeDeterminers["QuadraticBezierToPathSegment"]["Type"] = "QuadBezierTo"
	typeRegistry["RadicalElement"] = reflect.TypeOf(RadicalElement{})
	derivedTypes["RadicalElement"] = "MathElement"
	typeDeterminers["RadicalElement"] = make(map[string]string)
	typeDeterminers["RadicalElement"]["Type"] = "Radical"
	typeRegistry["ReflectionEffect"] = reflect.TypeOf(ReflectionEffect{})
	
	typeDeterminers["ReflectionEffect"] = make(map[string]string)
	typeRegistry["RemoveShape"] = reflect.TypeOf(RemoveShape{})
	derivedTypes["RemoveShape"] = "Task"
	typeDeterminers["RemoveShape"] = make(map[string]string)
	typeDeterminers["RemoveShape"]["Type"] = "RemoveShape"
	typeRegistry["RemoveSlide"] = reflect.TypeOf(RemoveSlide{})
	derivedTypes["RemoveSlide"] = "Task"
	typeDeterminers["RemoveSlide"] = make(map[string]string)
	typeDeterminers["RemoveSlide"]["Type"] = "RemoveSlide"
	typeRegistry["ReorderSlide"] = reflect.TypeOf(ReorderSlide{})
	derivedTypes["ReorderSlide"] = "Task"
	typeDeterminers["ReorderSlide"] = make(map[string]string)
	typeDeterminers["ReorderSlide"]["Type"] = "ReoderSlide"
	typeRegistry["ReplaceText"] = reflect.TypeOf(ReplaceText{})
	derivedTypes["ReplaceText"] = "Task"
	typeDeterminers["ReplaceText"] = make(map[string]string)
	typeDeterminers["ReplaceText"]["Type"] = "ReplaceText"
	typeRegistry["RequestInputFile"] = reflect.TypeOf(RequestInputFile{})
	derivedTypes["RequestInputFile"] = "InputFile"
	typeDeterminers["RequestInputFile"] = make(map[string]string)
	typeDeterminers["RequestInputFile"]["Type"] = "Request"
	typeRegistry["ResetSlide"] = reflect.TypeOf(ResetSlide{})
	derivedTypes["ResetSlide"] = "Task"
	typeDeterminers["ResetSlide"] = make(map[string]string)
	typeDeterminers["ResetSlide"]["Type"] = "ResetSlide"
	typeRegistry["ResourceBase"] = reflect.TypeOf(ResourceBase{})
	
	typeDeterminers["ResourceBase"] = make(map[string]string)
	typeRegistry["ResourceUri"] = reflect.TypeOf(ResourceUri{})
	
	typeDeterminers["ResourceUri"] = make(map[string]string)
	typeRegistry["ResponseOutputFile"] = reflect.TypeOf(ResponseOutputFile{})
	derivedTypes["ResponseOutputFile"] = "OutputFile"
	typeDeterminers["ResponseOutputFile"] = make(map[string]string)
	typeDeterminers["ResponseOutputFile"]["Type"] = "Response"
	typeRegistry["RightSubSuperscriptElement"] = reflect.TypeOf(RightSubSuperscriptElement{})
	derivedTypes["RightSubSuperscriptElement"] = "MathElement"
	typeDeterminers["RightSubSuperscriptElement"] = make(map[string]string)
	typeDeterminers["RightSubSuperscriptElement"]["Type"] = "RightSubSuperscriptElement"
	typeRegistry["Save"] = reflect.TypeOf(Save{})
	derivedTypes["Save"] = "Task"
	typeDeterminers["Save"] = make(map[string]string)
	typeDeterminers["Save"]["Type"] = "Save"
	typeRegistry["SaveShape"] = reflect.TypeOf(SaveShape{})
	derivedTypes["SaveShape"] = "Task"
	typeDeterminers["SaveShape"] = make(map[string]string)
	typeDeterminers["SaveShape"]["Type"] = "SaveShape"
	typeRegistry["SaveSlide"] = reflect.TypeOf(SaveSlide{})
	derivedTypes["SaveSlide"] = "Task"
	typeDeterminers["SaveSlide"] = make(map[string]string)
	typeDeterminers["SaveSlide"]["Type"] = "SaveSlide"
	typeRegistry["ScatterChartDataPoint"] = reflect.TypeOf(ScatterChartDataPoint{})
	derivedTypes["ScatterChartDataPoint"] = "DataPoint"
	typeDeterminers["ScatterChartDataPoint"] = make(map[string]string)
	typeDeterminers["ScatterChartDataPoint"]["Type"] = "Scatter"
	typeRegistry["ScatterSeries"] = reflect.TypeOf(ScatterSeries{})
	derivedTypes["ScatterSeries"] = "XYSeries"
	typeDeterminers["ScatterSeries"] = make(map[string]string)
	typeDeterminers["ScatterSeries"]["DataPointType"] = "Scatter"
	typeRegistry["Section"] = reflect.TypeOf(Section{})
	derivedTypes["Section"] = "ResourceBase"
	typeDeterminers["Section"] = make(map[string]string)
	typeRegistry["SectionZoomFrame"] = reflect.TypeOf(SectionZoomFrame{})
	derivedTypes["SectionZoomFrame"] = "ZoomObject"
	typeDeterminers["SectionZoomFrame"] = make(map[string]string)
	typeDeterminers["SectionZoomFrame"]["Type"] = "SectionZoomFrame"
	typeRegistry["Sections"] = reflect.TypeOf(Sections{})
	derivedTypes["Sections"] = "ResourceBase"
	typeDeterminers["Sections"] = make(map[string]string)
	typeRegistry["Series"] = reflect.TypeOf(Series{})
	
	typeDeterminers["Series"] = make(map[string]string)
	typeRegistry["SeriesMarker"] = reflect.TypeOf(SeriesMarker{})
	
	typeDeterminers["SeriesMarker"] = make(map[string]string)
	typeRegistry["Shape"] = reflect.TypeOf(Shape{})
	derivedTypes["Shape"] = "GeometryShape"
	typeDeterminers["Shape"] = make(map[string]string)
	typeDeterminers["Shape"]["Type"] = "Shape"
	typeRegistry["ShapeBase"] = reflect.TypeOf(ShapeBase{})
	derivedTypes["ShapeBase"] = "ResourceBase"
	typeDeterminers["ShapeBase"] = make(map[string]string)
	typeRegistry["ShapeBevel"] = reflect.TypeOf(ShapeBevel{})
	
	typeDeterminers["ShapeBevel"] = make(map[string]string)
	typeRegistry["ShapeImageExportOptions"] = reflect.TypeOf(ShapeImageExportOptions{})
	
	typeDeterminers["ShapeImageExportOptions"] = make(map[string]string)
	typeRegistry["Shapes"] = reflect.TypeOf(Shapes{})
	derivedTypes["Shapes"] = "ResourceBase"
	typeDeterminers["Shapes"] = make(map[string]string)
	typeRegistry["Slide"] = reflect.TypeOf(Slide{})
	derivedTypes["Slide"] = "ResourceBase"
	typeDeterminers["Slide"] = make(map[string]string)
	typeRegistry["SlideAnimation"] = reflect.TypeOf(SlideAnimation{})
	derivedTypes["SlideAnimation"] = "ResourceBase"
	typeDeterminers["SlideAnimation"] = make(map[string]string)
	typeRegistry["SlideBackground"] = reflect.TypeOf(SlideBackground{})
	derivedTypes["SlideBackground"] = "ResourceBase"
	typeDeterminers["SlideBackground"] = make(map[string]string)
	typeRegistry["SlideComment"] = reflect.TypeOf(SlideComment{})
	derivedTypes["SlideComment"] = "SlideCommentBase"
	typeDeterminers["SlideComment"] = make(map[string]string)
	typeDeterminers["SlideComment"]["Type"] = "Regular"
	typeRegistry["SlideCommentBase"] = reflect.TypeOf(SlideCommentBase{})
	
	typeDeterminers["SlideCommentBase"] = make(map[string]string)
	typeRegistry["SlideComments"] = reflect.TypeOf(SlideComments{})
	derivedTypes["SlideComments"] = "ResourceBase"
	typeDeterminers["SlideComments"] = make(map[string]string)
	typeRegistry["SlideModernComment"] = reflect.TypeOf(SlideModernComment{})
	derivedTypes["SlideModernComment"] = "SlideCommentBase"
	typeDeterminers["SlideModernComment"] = make(map[string]string)
	typeDeterminers["SlideModernComment"]["Type"] = "Modern"
	typeRegistry["SlideProperties"] = reflect.TypeOf(SlideProperties{})
	derivedTypes["SlideProperties"] = "ResourceBase"
	typeDeterminers["SlideProperties"] = make(map[string]string)
	typeRegistry["SlideReplaceResult"] = reflect.TypeOf(SlideReplaceResult{})
	derivedTypes["SlideReplaceResult"] = "Slide"
	typeDeterminers["SlideReplaceResult"] = make(map[string]string)
	typeRegistry["SlideShowProperties"] = reflect.TypeOf(SlideShowProperties{})
	derivedTypes["SlideShowProperties"] = "ResourceBase"
	typeDeterminers["SlideShowProperties"] = make(map[string]string)
	typeRegistry["SlideShowTransition"] = reflect.TypeOf(SlideShowTransition{})
	
	typeDeterminers["SlideShowTransition"] = make(map[string]string)
	typeRegistry["Slides"] = reflect.TypeOf(Slides{})
	derivedTypes["Slides"] = "ResourceBase"
	typeDeterminers["Slides"] = make(map[string]string)
	typeRegistry["SlidesLayoutOptions"] = reflect.TypeOf(SlidesLayoutOptions{})
	
	typeDeterminers["SlidesLayoutOptions"] = make(map[string]string)
	typeRegistry["SmartArt"] = reflect.TypeOf(SmartArt{})
	derivedTypes["SmartArt"] = "ShapeBase"
	typeDeterminers["SmartArt"] = make(map[string]string)
	typeDeterminers["SmartArt"]["Type"] = "SmartArt"
	typeRegistry["SmartArtNode"] = reflect.TypeOf(SmartArtNode{})
	
	typeDeterminers["SmartArtNode"] = make(map[string]string)
	typeRegistry["SmartArtShape"] = reflect.TypeOf(SmartArtShape{})
	derivedTypes["SmartArtShape"] = "GeometryShape"
	typeDeterminers["SmartArtShape"] = make(map[string]string)
	typeDeterminers["SmartArtShape"]["Type"] = "SmartArtShape"
	typeRegistry["SoftEdgeEffect"] = reflect.TypeOf(SoftEdgeEffect{})
	
	typeDeterminers["SoftEdgeEffect"] = make(map[string]string)
	typeRegistry["SolidFill"] = reflect.TypeOf(SolidFill{})
	derivedTypes["SolidFill"] = "FillFormat"
	typeDeterminers["SolidFill"] = make(map[string]string)
	typeDeterminers["SolidFill"]["Type"] = "Solid"
	typeRegistry["SplitDocumentResult"] = reflect.TypeOf(SplitDocumentResult{})
	derivedTypes["SplitDocumentResult"] = "ResourceBase"
	typeDeterminers["SplitDocumentResult"] = make(map[string]string)
	typeRegistry["StorageExist"] = reflect.TypeOf(StorageExist{})
	
	typeDeterminers["StorageExist"] = make(map[string]string)
	typeRegistry["StorageFile"] = reflect.TypeOf(StorageFile{})
	
	typeDeterminers["StorageFile"] = make(map[string]string)
	typeRegistry["SubscriptElement"] = reflect.TypeOf(SubscriptElement{})
	derivedTypes["SubscriptElement"] = "MathElement"
	typeDeterminers["SubscriptElement"] = make(map[string]string)
	typeDeterminers["SubscriptElement"]["Type"] = "SubscriptElement"
	typeRegistry["SummaryZoomFrame"] = reflect.TypeOf(SummaryZoomFrame{})
	derivedTypes["SummaryZoomFrame"] = "ShapeBase"
	typeDeterminers["SummaryZoomFrame"] = make(map[string]string)
	typeDeterminers["SummaryZoomFrame"]["Type"] = "SummaryZoomFrame"
	typeRegistry["SummaryZoomSection"] = reflect.TypeOf(SummaryZoomSection{})
	derivedTypes["SummaryZoomSection"] = "SectionZoomFrame"
	typeDeterminers["SummaryZoomSection"] = make(map[string]string)
	typeDeterminers["SummaryZoomSection"]["Type"] = "SummaryZoomSection"
	typeRegistry["SuperscriptElement"] = reflect.TypeOf(SuperscriptElement{})
	derivedTypes["SuperscriptElement"] = "MathElement"
	typeDeterminers["SuperscriptElement"] = make(map[string]string)
	typeDeterminers["SuperscriptElement"]["Type"] = "SuperscriptElement"
	typeRegistry["SvgExportOptions"] = reflect.TypeOf(SvgExportOptions{})
	derivedTypes["SvgExportOptions"] = "ExportOptions"
	typeDeterminers["SvgExportOptions"] = make(map[string]string)
	typeRegistry["SwfExportOptions"] = reflect.TypeOf(SwfExportOptions{})
	derivedTypes["SwfExportOptions"] = "ExportOptions"
	typeDeterminers["SwfExportOptions"] = make(map[string]string)
	typeRegistry["Table"] = reflect.TypeOf(Table{})
	derivedTypes["Table"] = "ShapeBase"
	typeDeterminers["Table"] = make(map[string]string)
	typeDeterminers["Table"]["Type"] = "Table"
	typeRegistry["TableCell"] = reflect.TypeOf(TableCell{})
	
	typeDeterminers["TableCell"] = make(map[string]string)
	typeRegistry["TableCellMergeOptions"] = reflect.TypeOf(TableCellMergeOptions{})
	
	typeDeterminers["TableCellMergeOptions"] = make(map[string]string)
	typeRegistry["TableColumn"] = reflect.TypeOf(TableColumn{})
	
	typeDeterminers["TableColumn"] = make(map[string]string)
	typeRegistry["TableRow"] = reflect.TypeOf(TableRow{})
	
	typeDeterminers["TableRow"] = make(map[string]string)
	typeRegistry["Task"] = reflect.TypeOf(Task{})
	
	typeDeterminers["Task"] = make(map[string]string)
	typeRegistry["TextBounds"] = reflect.TypeOf(TextBounds{})
	
	typeDeterminers["TextBounds"] = make(map[string]string)
	typeRegistry["TextElement"] = reflect.TypeOf(TextElement{})
	derivedTypes["TextElement"] = "MathElement"
	typeDeterminers["TextElement"] = make(map[string]string)
	typeDeterminers["TextElement"]["Type"] = "Text"
	typeRegistry["TextFrameFormat"] = reflect.TypeOf(TextFrameFormat{})
	
	typeDeterminers["TextFrameFormat"] = make(map[string]string)
	typeRegistry["TextItem"] = reflect.TypeOf(TextItem{})
	
	typeDeterminers["TextItem"] = make(map[string]string)
	typeRegistry["TextItems"] = reflect.TypeOf(TextItems{})
	derivedTypes["TextItems"] = "ResourceBase"
	typeDeterminers["TextItems"] = make(map[string]string)
	typeRegistry["Theme"] = reflect.TypeOf(Theme{})
	derivedTypes["Theme"] = "ResourceBase"
	typeDeterminers["Theme"] = make(map[string]string)
	typeRegistry["ThreeDFormat"] = reflect.TypeOf(ThreeDFormat{})
	
	typeDeterminers["ThreeDFormat"] = make(map[string]string)
	typeRegistry["TiffExportOptions"] = reflect.TypeOf(TiffExportOptions{})
	derivedTypes["TiffExportOptions"] = "ImageExportOptionsBase"
	typeDeterminers["TiffExportOptions"] = make(map[string]string)
	typeRegistry["TintEffect"] = reflect.TypeOf(TintEffect{})
	derivedTypes["TintEffect"] = "ImageTransformEffect"
	typeDeterminers["TintEffect"] = make(map[string]string)
	typeDeterminers["TintEffect"]["Type"] = "Tint"
	typeRegistry["UpdateBackground"] = reflect.TypeOf(UpdateBackground{})
	derivedTypes["UpdateBackground"] = "Task"
	typeDeterminers["UpdateBackground"] = make(map[string]string)
	typeDeterminers["UpdateBackground"]["Type"] = "UpdateBackground"
	typeRegistry["UpdateShape"] = reflect.TypeOf(UpdateShape{})
	derivedTypes["UpdateShape"] = "Task"
	typeDeterminers["UpdateShape"] = make(map[string]string)
	typeDeterminers["UpdateShape"]["Type"] = "UpdateShape"
	typeRegistry["VbaModule"] = reflect.TypeOf(VbaModule{})
	derivedTypes["VbaModule"] = "ResourceBase"
	typeDeterminers["VbaModule"] = make(map[string]string)
	typeRegistry["VbaProject"] = reflect.TypeOf(VbaProject{})
	derivedTypes["VbaProject"] = "ResourceBase"
	typeDeterminers["VbaProject"] = make(map[string]string)
	typeRegistry["VbaReference"] = reflect.TypeOf(VbaReference{})
	
	typeDeterminers["VbaReference"] = make(map[string]string)
	typeRegistry["VideoExportOptions"] = reflect.TypeOf(VideoExportOptions{})
	derivedTypes["VideoExportOptions"] = "ExportOptions"
	typeDeterminers["VideoExportOptions"] = make(map[string]string)
	typeRegistry["VideoFrame"] = reflect.TypeOf(VideoFrame{})
	derivedTypes["VideoFrame"] = "GeometryShape"
	typeDeterminers["VideoFrame"] = make(map[string]string)
	typeDeterminers["VideoFrame"]["Type"] = "VideoFrame"
	typeRegistry["ViewProperties"] = reflect.TypeOf(ViewProperties{})
	derivedTypes["ViewProperties"] = "ResourceBase"
	typeDeterminers["ViewProperties"] = make(map[string]string)
	typeRegistry["Workbook"] = reflect.TypeOf(Workbook{})
	derivedTypes["Workbook"] = "DataSource"
	typeDeterminers["Workbook"] = make(map[string]string)
	typeDeterminers["Workbook"]["Type"] = "Workbook"
	typeRegistry["XYSeries"] = reflect.TypeOf(XYSeries{})
	derivedTypes["XYSeries"] = "Series"
	typeDeterminers["XYSeries"] = make(map[string]string)
	typeRegistry["XamlExportOptions"] = reflect.TypeOf(XamlExportOptions{})
	derivedTypes["XamlExportOptions"] = "ExportOptions"
	typeDeterminers["XamlExportOptions"] = make(map[string]string)
	typeRegistry["XpsExportOptions"] = reflect.TypeOf(XpsExportOptions{})
	derivedTypes["XpsExportOptions"] = "ExportOptions"
	typeDeterminers["XpsExportOptions"] = make(map[string]string)
	typeRegistry["ZoomFrame"] = reflect.TypeOf(ZoomFrame{})
	derivedTypes["ZoomFrame"] = "ZoomObject"
	typeDeterminers["ZoomFrame"] = make(map[string]string)
	typeDeterminers["ZoomFrame"]["Type"] = "ZoomFrame"
	typeRegistry["ZoomObject"] = reflect.TypeOf(ZoomObject{})
	derivedTypes["ZoomObject"] = "ShapeBase"
	typeDeterminers["ZoomObject"] = make(map[string]string)
}

func isModelClass(typeName string) bool {
	_, ok := typeRegistry[typeName]
	return ok
}

func isSubclass(typeName string, baseTypeName string) bool {
	if typeName == baseTypeName {
		return true
	}
	for k, v := range derivedTypes {
		if v == baseTypeName && isSubclass(typeName, k) {
			return true
		}
	}
	return false
}

func createObjectForType(typeName string, data []byte) (interface{}, error) {
	var objmap map[string]*json.RawMessage
	err := json.Unmarshal(data, &objmap)
	if err != nil {
		return nil, err
	}
	return createObjectForTypeMap(typeName, typeName, objmap), nil
}

func createObjectForMap(typeName string, baseTypeName string, objMap map[string]interface{}) interface{} {
	for k, v := range derivedTypes {
		if v == typeName {
			subObject := createObjectForMap(k, typeName, objMap)
			if subObject != nil {
				return subObject
			}
		}
	}
	matches := 0
	if detMap, ok := typeDeterminers[typeName]; ok {
		matches = 0
		for k, v := range detMap {
			objValue, objOk := objMap[k]
			if !objOk {
				objValue, objOk = objMap[lcFirst(k)]
			}
			if !objOk {
				return nil
			}
			if objOk {
				var objString string
				var objBytes, ok = objValue.([]byte)
				if ok {
				 	err := json.Unmarshal(objBytes, &objString)
					if err == nil {
						if objString != v {
							return nil
						}
					}
				}
			}
			matches++
		}
	}
	if typeName == baseTypeName || matches > 0 {
		if reflectType, ok := typeRegistry[typeName]; ok {
			return reflect.New(reflectType).Interface()
		}
	}
	return nil
}

func createObjectForTypeMap(typeName string, baseTypeName string, objMap map[string]*json.RawMessage) interface{} {
	for k, v := range derivedTypes {
		if v == typeName {
			subObject := createObjectForTypeMap(k, typeName, objMap)
			if subObject != nil {
				return subObject
			}
		}
	}
	matches := 0
	if detMap, ok := typeDeterminers[typeName]; ok {
		matches = 0
		for k, v := range detMap {
			objValue, objOk := objMap[k]
			if !objOk {
				objValue, objOk = objMap[lcFirst(k)]
			}
			if !objOk {
				return nil
			}
			if objOk {
				var objString string
			 	err := json.Unmarshal(*objValue, &objString)
				if err == nil {
					if objString != v {
						return nil
					}
				}
			}
			matches++
		}
	}
	if typeName == baseTypeName || matches > 0 {
		if reflectType, ok := typeRegistry[typeName]; ok {
			return reflect.New(reflectType).Interface()
		}
	}
	return nil
}
 
func lcFirst(str string) string {
    for i, v := range str {                                                                                                                                           
        return string(unicode.ToLower(v)) + str[i+1:]
    }  
    return ""
}