package context

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/openapi/models"
)

func TestSortedPCCRuleIDsByQoSRef(t *testing.T) {
	rules := map[string]*PCCRule{
		"PccRuleId-3": NewPCCRule(&models.PccRule{
			PccRuleId:  "PccRuleId-3",
			RefQosData: []string{"3"},
		}),
		"PccRuleId-1": NewPCCRule(&models.PccRule{
			PccRuleId:  "PccRuleId-1",
			RefQosData: []string{"1"},
		}),
		"PccRuleId-2": NewPCCRule(&models.PccRule{
			PccRuleId:  "PccRuleId-2",
			RefQosData: []string{"2"},
		}),
	}

	require.Equal(t,
		[]string{"PccRuleId-1", "PccRuleId-2", "PccRuleId-3"},
		sortedPCCRuleIDsByQoSRef(rules),
	)
}
