package notification

import (
	"github.com/goto/siren/pkg/errors"
)

type ReceiverSelectors []map[string]any

func (rs ReceiverSelectors) parseAndValidate() ([]map[string]string, map[string]any, error) {
	// Check if any selector contains a config
	var selectorConfig map[string]any
	for i := 0; i < len(rs); i++ {
		selector := rs[i]
		if v, cok := selector["config"]; cok {
			if m, ok := v.(map[string]any); ok {
				selectorConfig = m
				delete(rs[i], "config")
			} else {
				return nil, nil, errors.ErrInvalid.WithMsgf("config should be in map and follow notification config")
			}
			break
		}
	}

	if selectorConfig != nil && len(rs) > 1 {
		return nil, nil, errors.ErrInvalid.WithMsgf("config override could only be used with one selector")
	}

	castedSelectors := make([]map[string]string, len(rs))
	for i, selector := range rs {
		castedSelectors[i] = make(map[string]string)
		for k, v := range selector {
			if str, ok := v.(string); ok {
				castedSelectors[i][k] = str
			} else {
				return nil, nil, errors.ErrInvalid.WithMsgf("receiver selector value of '%s' should be a string", k)
			}
		}
	}

	return castedSelectors, selectorConfig, nil
}
