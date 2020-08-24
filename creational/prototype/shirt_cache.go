package prototype

import (
	"errors"
	"fmt"
)

type ShirtCache struct{}
func (s *ShirtCache) GetClone(c int) (ItemInfoGetter, error){
	switch c{
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New(fmt.Sprintf("Prototype %d not implemented yet", c))
	}
}