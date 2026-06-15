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

// Smart art node.
type ISmartArtNode interface {

	// Node list.
	GetNodes() []ISmartArtNode
	SetNodes(newValue []ISmartArtNode)

	// Gets or sets the link to shapes.
	GetShapes() IResourceUri
	SetShapes(newValue IResourceUri)

	// True for and assistant node.
	GetIsAssistant() *bool
	SetIsAssistant(newValue *bool)

	// Node text.
	GetText() string
	SetText(newValue string)

	// Organization chart layout type associated with current node.
	GetOrgChartLayout() string
	SetOrgChartLayout(newValue string)

	// Get or sets list to paragraphs list
	GetParagraphs() IResourceUri
	SetParagraphs(newValue IResourceUri)

	// Default paragraph format for the node's text frame.
	GetDefaultParagraphFormat() IParagraphFormat
	SetDefaultParagraphFormat(newValue IParagraphFormat)
}

type SmartArtNode struct {

	// Node list.
	Nodes []ISmartArtNode `json:"Nodes,omitempty"`

	// Gets or sets the link to shapes.
	Shapes IResourceUri `json:"Shapes,omitempty"`

	// True for and assistant node.
	IsAssistant *bool `json:"IsAssistant"`

	// Node text.
	Text string `json:"Text,omitempty"`

	// Organization chart layout type associated with current node.
	OrgChartLayout string `json:"OrgChartLayout"`

	// Get or sets list to paragraphs list
	Paragraphs IResourceUri `json:"Paragraphs,omitempty"`

	// Default paragraph format for the node's text frame.
	DefaultParagraphFormat IParagraphFormat `json:"DefaultParagraphFormat,omitempty"`
}

func NewSmartArtNode() *SmartArtNode {
	instance := new(SmartArtNode)
	instance.OrgChartLayout = "Initial"
	return instance
}

func (this *SmartArtNode) GetNodes() []ISmartArtNode {
	return this.Nodes
}

func (this *SmartArtNode) SetNodes(newValue []ISmartArtNode) {
	this.Nodes = newValue
}
func (this *SmartArtNode) GetShapes() IResourceUri {
	return this.Shapes
}

func (this *SmartArtNode) SetShapes(newValue IResourceUri) {
	this.Shapes = newValue
}
func (this *SmartArtNode) GetIsAssistant() *bool {
	return this.IsAssistant
}

func (this *SmartArtNode) SetIsAssistant(newValue *bool) {
	this.IsAssistant = newValue
}
func (this *SmartArtNode) GetText() string {
	return this.Text
}

func (this *SmartArtNode) SetText(newValue string) {
	this.Text = newValue
}
func (this *SmartArtNode) GetOrgChartLayout() string {
	return this.OrgChartLayout
}

func (this *SmartArtNode) SetOrgChartLayout(newValue string) {
	this.OrgChartLayout = newValue
}
func (this *SmartArtNode) GetParagraphs() IResourceUri {
	return this.Paragraphs
}

func (this *SmartArtNode) SetParagraphs(newValue IResourceUri) {
	this.Paragraphs = newValue
}
func (this *SmartArtNode) GetDefaultParagraphFormat() IParagraphFormat {
	return this.DefaultParagraphFormat
}

func (this *SmartArtNode) SetDefaultParagraphFormat(newValue IParagraphFormat) {
	this.DefaultParagraphFormat = newValue
}

func (this *SmartArtNode) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valNodes, ok := GetMapValue(objMap, "nodes"); ok {
		if valNodes != nil {
			var valueForNodes []json.RawMessage
			err = json.Unmarshal(*valNodes, &valueForNodes)
			if err != nil {
				return err
			}
			valueForINodes := make([]ISmartArtNode, len(valueForNodes))
			for i, v := range valueForNodes {
				vObject, err := createObjectForType("SmartArtNode", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForINodes[i] = vObject.(ISmartArtNode)
				}
			}
			this.Nodes = valueForINodes
		}
	}
	
	if valShapes, ok := GetMapValue(objMap, "shapes"); ok {
		if valShapes != nil {
			var valueForShapes ResourceUri
			err = json.Unmarshal(*valShapes, &valueForShapes)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valShapes)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valShapes, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.Shapes = vInterfaceObject
			}
		}
	}
	
	if valIsAssistant, ok := GetMapValue(objMap, "isAssistant"); ok {
		if valIsAssistant != nil {
			var valueForIsAssistant *bool
			err = json.Unmarshal(*valIsAssistant, &valueForIsAssistant)
			if err != nil {
				return err
			}
			this.IsAssistant = valueForIsAssistant
		}
	}
	
	if valText, ok := GetMapValue(objMap, "text"); ok {
		if valText != nil {
			var valueForText string
			err = json.Unmarshal(*valText, &valueForText)
			if err != nil {
				return err
			}
			this.Text = valueForText
		}
	}
	this.OrgChartLayout = "Initial"
	if valOrgChartLayout, ok := GetMapValue(objMap, "orgChartLayout"); ok {
		if valOrgChartLayout != nil {
			var valueForOrgChartLayout string
			err = json.Unmarshal(*valOrgChartLayout, &valueForOrgChartLayout)
			if err != nil {
				var valueForOrgChartLayoutInt int32
				err = json.Unmarshal(*valOrgChartLayout, &valueForOrgChartLayoutInt)
				if err != nil {
					return err
				}
				this.OrgChartLayout = string(valueForOrgChartLayoutInt)
			} else {
				this.OrgChartLayout = valueForOrgChartLayout
			}
		}
	}
	
	if valParagraphs, ok := GetMapValue(objMap, "paragraphs"); ok {
		if valParagraphs != nil {
			var valueForParagraphs ResourceUri
			err = json.Unmarshal(*valParagraphs, &valueForParagraphs)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valParagraphs)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valParagraphs, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.Paragraphs = vInterfaceObject
			}
		}
	}
	
	if valDefaultParagraphFormat, ok := GetMapValue(objMap, "defaultParagraphFormat"); ok {
		if valDefaultParagraphFormat != nil {
			var valueForDefaultParagraphFormat ParagraphFormat
			err = json.Unmarshal(*valDefaultParagraphFormat, &valueForDefaultParagraphFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ParagraphFormat", *valDefaultParagraphFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valDefaultParagraphFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IParagraphFormat)
			if ok {
				this.DefaultParagraphFormat = vInterfaceObject
			}
		}
	}

	return nil
}
