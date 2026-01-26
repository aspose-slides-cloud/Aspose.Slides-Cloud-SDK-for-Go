![](https://img.shields.io/badge/api-v3.0-lightgrey)  [![GitHub license](https://img.shields.io/github/license/aspose-slides-cloud/aspose-slides-cloud-go)](https://github.com/aspose-slides-cloud/aspose-slides-cloud-go)

# Go REST API to Process Presentation in Cloud
This repository contains Aspose.Slides Cloud SDK for Go source code. This SDK allows you to [process & manipulate PPT, PPTX, ODP, OTP](https://products.aspose.cloud/slides/go) using Aspose.slides Cloud REST APIs in your applications.

You may want to check out Aspose free [Powerpoint to PDF](https://products.aspose.app/slides/conversion), [Powerpoint to Word](https://products.aspose.app/slides/conversion/ppt-to-word), [Powerpoint to JPG](https://products.aspose.app/slides/conversion/ppt-to-jpg), [Powerpoint to PNG](https://products.aspose.app/slides/conversion/ppt-to-png), [PDF to Powerpoint](https://products.aspose.app/slides/import/pdf-to-powerpoint), [JPG to Powerpoint](https://products.aspose.app/slides/import/jpg-to-ppt), and [PNG to Powerpoint](https://products.aspose.app/slides/import/png-to-ppt) converters because they are live implementations of popular conversion processes.

## Presentation Processing Features

- Fetch presentation images in any of the supported file formats.
- Copy the layout side or clone the master slide from the source presentation.
- Process slides shapes, slides notes, placeholders, colors & font theme info.
- Programmatically create a presentation from HTML & export it to various formats.
- Merge multiple presentations or split the single presentation into multiple ones.
- Extract and replace text from a specific slide or an entire presentation.

## Read & Write Presentation Formats

**Microsoft PowerPoint:** PPT, PPTX, PPS, PPSX, PPTM, PPSM, POTX, POTM
**OpenOffice:** ODP, OTP, FODP
**Other**: PDF, PDF/A

## Save Presentation As

**Fixed Layout:** XPS
**Images:** JPEG, PNG, BMP, TIFF, GIF, EMF, SVG
**Web:** HTML/HTML5
**Other:** MPEG4, SWF (export whole presentations)

## Enhancements in Version 26.1

* Added **SpellCheck** boolean property to **PortionFormat**.
* Extended **SlideShowTransition** with a **Duration** field.
* Introduced the **Phantom** math element. It represents an invisible math element whose spacing behavior can be configured.

## Enhancements in Version 25.10

* Added **SetEmbeddedFonts** and **SetEmbeddedFontsOnline** methods to embed many fonts in one call.
* Removed **SetEmbeddedFontFromRequest** and **SetEmbeddedFontFromRequestOnline** methods. The above methods can be used instead.
* Added optional **shapeIndex** and **shapeIndex** parameters to **GetSlideImages** method.
* Renamed **GetVideoCaptionTracks**, **CreateVideoCaptionTrack**, **DeleteVideoCaptionTracks** and **DeleteVideoCaptionTracks** methods into **GetCaptionTracks**, **CreateCaptionTrack**, **DeleteCaptionTracks** and **DeleteCaptionTracks**. The methods now work with video frames and audio frames.

## Enhancements in Version 25.6

* Added **GetVideoCaptionTracks**, **CreateVideoCaptionTrack**, **DeleteVideoCaptionTracks**, **DeleteVideoCaptionTracks** method.
* Removed deprecated **DeletePictureCroppedAreas** method.
* Added an optional **quality** parameter to **DownloadImage**, **DownloadImageOnline**, **DownloadImages**, **DownloadImagesOnline** method. It applies to jpeg fromat.
* Added **RemoveEmptyLines**, **SlideNumberFormat** and **HandleRepeatedSpaces** properties to **MarkdownExportOptions** class.
* Added **DeletePictureCroppedAreas** and **Resolution** parameters to **PictureFill** class that allow to compress and/or delete cropped aread of picture fill formats.
* Added **Emf** to the list of allowed values for **ExportFormat**, **SlideExportFormat** and **ShapeExportFormat** enumerations. You can now export presentations, slides and shapes to EMF format.
* Added **VolumeValue**, **FadeInDuration**, **FadeOutDuration**, **TrimFromStart** and **TrimFromEnd** properties to **AudioFrame** class.
* Added **SkipJavaScriptLinks** boolean property to **ExportOptions** class.
* Added **DisableFontLigatures** boolean property to **HtmlExportOptions**, **Html5ExportOptions** and **SvgExportOptions** class.
* Added **SlidesLayoutOptions** property and removed **NotesCommentsLayouting** property from **Html5ExportOptions** class.
* Added **RemoveEmptyLines**, **SlideNumberFormat** and **HandleRepeatedSpaces** properties to **MarkdownExportOptions** class.
* Added **SlidesLayoutOptions** property and removed **CommentsPosition**, **CommentsAreaColor**, **CommentsAreaWidth**, **NotesPosition** and **ShowCommentsByNoAuthor** properties from **SwfExportOptions** class.
* Added **RefreshThumbnail** boolean property to **PptxExportOptions** class.

## Licensing
All Aspose.Slides Cloud SDKs are licensed under MIT License.

## How to use the SDK?
The complete source code is available in this repository folder.

## Prerequisites
To use Aspose Slides Cloud SDK for Go you need to register an account with [Aspose Cloud](https://www.aspose.cloud/) and lookup/create App Key and SID at [Cloud Dashboard](https://dashboard.aspose.cloud/#/apps). There is free quota available. For more details, see [Aspose Cloud Pricing](https://purchase.aspose.cloud/pricing).

### Installation

```sh
go get github.com/aspose-slides-cloud/aspose-slides-cloud-go
```

### Sample usage

The example code below converts a PowerPoint document to PDF format using asposeslidescloud library:
```slides
    cfg := asposeslidescloud.NewConfiguration()
    cfg.AppSid = "MyClientId"
    cfg.AppKey = "MyClientSecret"
    apiClient := asposeslidescloud.NewAPIClient(cfg)
    file, _ := ioutil.ReadFile("MyPresentation.pptx")
    r, _, _ := apiClient.SlidesApi.Convert(file, "pdf", "", "", "", nil, nil)
    fmt.Printf("My PDF was saved to %s", r.Name())
```

## Aspose.Slides Cloud SDKs in Popular Languages

| .NET | Java | PHP | Python | Ruby | Node.js | Android | Swift|Perl|Go|
|---|---|---|---|---|---|---|--|--|--|
| [GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-dotnet) | [GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-java) | [GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-php) | [GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-python) | [GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-ruby)  | [GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-nodejs) | [GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-android) | [GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-swift)|[GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-perl) |[GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-go) |
| [NuGet](https://www.nuget.org/packages/Aspose.slides-Cloud/) | [Maven](https://repository.aspose.cloud/webapp/#/artifacts/browse/tree/General/repo/com/aspose/aspose-slides-cloud) | [Composer](https://packagist.org/packages/aspose/slides-sdk-php) | [PIP](https://pypi.org/project/asposeslidescloud/) | [GEM](https://rubygems.org/gems/aspose_slides_cloud)  | [NPM](https://www.npmjs.com/package/asposeslidescloud) | [Maven](https://repository.aspose.cloud/webapp/#/artifacts/browse/tree/General/repo/com/aspose/aspose-slides-cloud) | [Cocoapods](https://cocoapods.org/pods/AsposeslidesCloud)|[Meta Cpan](https://metacpan.org/release/AsposeSlidesCloud-SlidesApi) | [Go.Dev](https://pkg.go.dev/github.com/aspose-slides-cloud/aspose-slides-cloud-go/) |

[Product Page](https://products.aspose.cloud/slides/slides) | [Documentation](https://docs.aspose.cloud/display/slidescloud/Home) | [API Reference](https://apireference.aspose.cloud/slides/) | [Code Samples](https://github.com/aspose-slides-cloud/aspose-slides-cloud-slides) | [Blog](https://blog.aspose.cloud/cateslidesry/slides/) | [Free Support](https://forum.aspose.cloud/c/slides) | [Free Trial](https://dashboard.aspose.cloud/#/apps)
