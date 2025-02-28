package xml

import (
	"errors"

	"github.com/beevik/etree"
)

var (
	ErrElementIsNil               = errors.New("element is nil")
	ErrInvalidElementTag          = errors.New("invalid element tag")
	ErrChildElementNotFound       = errors.New("child element not found")
	ErrMultipleChildElementsFound = errors.New("multiple child elements found")
)

func ValidateElement(el *etree.Element, tag string, namespaceUri string) error {
	if el == nil {
		return ErrElementIsNil
	}
	if el.Tag != tag || el.NamespaceURI() != namespaceUri {
		return ErrInvalidElementTag
	}
	return nil
}

func GetSingleChildElement(el *etree.Element, tag string, namespaceUri string) (*etree.Element, error) {
	elements := el.SelectElements(tag)
	if len(elements) == 0 {
		return nil, ErrChildElementNotFound
	}
	if len(elements) > 1 {
		return nil, ErrMultipleChildElementsFound
	}
	return elements[0], nil
}

func GetOptionalSingleChildElement(el *etree.Element, tag string, namespaceUri string) (*etree.Element, error) {
	elements := el.SelectElements(tag)
	if len(elements) > 1 {
		return nil, ErrMultipleChildElementsFound
	}
	if len(elements) > 0 {
		return elements[0], nil
	}
	return nil, nil
}
