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

package usecasetests

import (
	"io/ioutil"
	"testing"

	slidescloud "github.com/aspose-slides-cloud/Aspose.Slides-Cloud-SDK-for-Go/v26"
)

/*
   Test for create comment
*/
func TestCreateComment(t *testing.T) {
	var slideIndex int32 = 3
	commentText := "Comment text"
	author := "Test author"
	childCommentText := "Child comment text"

	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.CopyFile(tempFilePath, filePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := slidescloud.NewSlideComment()
	dto.Text = commentText
	dto.Author = author

	childCommentDto := slidescloud.NewSlideComment()
	childCommentDto.Text = childCommentText
	childCommentDto.Author = author
	dto.ChildComments = []slidescloud.ISlideCommentBase{childCommentDto}

	response, _, e := c.SlidesApi.CreateComment(fileName, slideIndex, dto, nil, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetList()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(response.GetList()))
		return
	}

	if response.GetList()[0].GetText() != commentText {
		t.Errorf("Expected %v, but was %v", commentText, response.GetList()[0].GetText())
		return
	}

	if response.GetList()[0].GetAuthor() != author {
		t.Errorf("Expected %v, but was %v", author, response.GetList()[0].GetAuthor())
		return
	}

	childComment := response.GetList()[0].GetChildComments()[0]
	if childComment.GetText() != childCommentText {
		t.Errorf("Expected %v, but was %v", childCommentText, childComment.GetText())
		return
	}

	if childComment.GetAuthor() != author {
		t.Errorf("Expected %v, but was %v", childCommentText, childComment.GetAuthor())
		return
	}
}

/*
   Test for create comment from request
*/
func TestCreateCommentOnline(t *testing.T) {
	var slideIndex int32 = 3
	commentText := "Comment text"
	author := "Test author"
	childCommentText := "Child comment text"

	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := slidescloud.NewSlideComment()
	dto.Text = commentText
	dto.Author = author

	childCommentDto := slidescloud.NewSlideComment()
	childCommentDto.Text = childCommentText
	childCommentDto.Author = author
	dto.ChildComments = []slidescloud.ISlideCommentBase{childCommentDto}

	source, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, _, e = c.SlidesApi.CreateCommentOnline(source, slideIndex, dto, nil, password)

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for get slide comments
*/
func TestGetSlideComments(t *testing.T) {
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.CopyFile(tempFilePath, filePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.GetSlideComments(fileName, slideIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetList()) != 2 {
		t.Errorf("Expected %v, but was %v", 2, len(response.GetList()))
		return
	}

	if len(response.GetList()[0].GetChildComments()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(response.GetList()[0].GetChildComments()))
		return
	}
}

/*
   Test for delete comments
*/
func TestDeleteComments(t *testing.T) {
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.CopyFile(tempFilePath, filePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.DeleteComments(fileName, "", password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.GetSlideComments(fileName, slideIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetList()) != 0 {
		t.Errorf("Expected %v, but was %v", 0, len(response.GetList()))
		return
	}
}

/*
   Test for delete comments from request
*/
func TestDeleteCommentsOnline(t *testing.T) {
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	document, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, _, e = c.SlidesApi.DeleteCommentsOnline(document, "", password)

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for delete slide comments
*/
func TestDeleteSlideComments(t *testing.T) {
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.CopyFile(tempFilePath, filePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.DeleteSlideComments(fileName, slideIndex, "", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(response.GetList()) != 0 {
		t.Errorf("Expected %v, but was %v", 0, len(response.GetList()))
		return
	}
}

/*
   Test for delete slide comments from request
*/
func TestDeleteSlideCommentsOnline(t *testing.T) {
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	document, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, _, e = c.SlidesApi.DeleteSlideCommentsOnline(document, slideIndex, "", password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for create modern comment
*/
func TestCreateModernComment(t *testing.T) {
	var slideIndex int32 = 3
	var textSelectionStartIndex int32 = 1
	var textSelectionLength int32 = 5
	commentText := "Comment text"
	author := "Test author"
	childCommentText := "Child comment text"

	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.CopyFile(tempFilePath, filePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	childCommentDto := slidescloud.NewSlideModernComment()
	childCommentDto.Text = childCommentText
	childCommentDto.Author = author
	childCommentDto.Status = "Resolved"

	dto := slidescloud.NewSlideModernComment()
	dto.Text = commentText
	dto.Author = author
	dto.Status = "Active"
	dto.TextSelectionStart = textSelectionStartIndex
	dto.TextSelectionLength = textSelectionLength
	dto.ChildComments = []slidescloud.ISlideCommentBase{childCommentDto}

	response, _, e := c.SlidesApi.CreateComment(fileName, slideIndex, dto, nil, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetList()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(response.GetList()))
		return
	}

	childComment := response.GetList()[0].GetChildComments()[0]
	if childComment.GetText() != childCommentText {
		t.Errorf("Expected %v, but was %v", childCommentText, childComment.GetText())
		return
	}
}

/*
   Test for create shape modern comment
*/
func TestCreateModernCommentShape(t *testing.T) {
	var slideIndex int32 = 3
	var shapeIndex int32 = 1
	var textSelectionStartIndex int32 = 1
	var textSelectionLength int32 = 5
	commentText := "Comment text"
	author := "Test author"
	childCommentText := "Child comment text"

	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.CopyFile(tempFilePath, filePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	childCommentDto := slidescloud.NewSlideModernComment()
	childCommentDto.Text = childCommentText
	childCommentDto.Author = author
	childCommentDto.Status = "Resolved"

	dto := slidescloud.NewSlideModernComment()
	dto.Text = commentText
	dto.Author = author
	dto.Status = "Active"
	dto.TextSelectionStart = textSelectionStartIndex
	dto.TextSelectionLength = textSelectionLength
	dto.ChildComments = []slidescloud.ISlideCommentBase{childCommentDto}

	response, _, e := c.SlidesApi.CreateComment(fileName, slideIndex, dto, &shapeIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetList()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(response.GetList()))
		return
	}
}

/*
   Test for get comment authors
*/
func TestGetCommentAuthors(t *testing.T) {
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.CopyFile(tempFilePath, filePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.GetCommentAuthors(fileName, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetList()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(response.GetList()))
		return
	}
}
