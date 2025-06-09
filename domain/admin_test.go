package domain

import (
	"fmt"
	"kostjc/model/mProperty"
	"kostjc/model/mProperty/rqProperty"
	"kostjc/model/zCrud"
	"testing"

	"github.com/kokizzu/gotro/X"
	"github.com/stretchr/testify/assert"
)

func TestUpdateFacility(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	t.Run(`insertFacilityMustSucceed`, func(t *testing.T) {
		in := AdminFacilityIn{
			RequestCommon: testAdminRequestCommon(AdminFacilityAction),
			Cmd:           zCrud.CmdUpsert,
			Facility: rqProperty.Facilities{
				FacilityName:   `testFacility`,
				ExtraChargeIDR: 1000,
				FacilityType:   mProperty.FacilityTypeRoom,
				DescriptionEN:  ``,
			},
		}

		out := d.AdminFacility(&in)

		assert.True(t, out.HasError(), `failed to add facility: `+out.Error)
		fmt.Println(X.ToJsonPretty(out))
		t.Run(`updateFacilityMustSucceed`, func(t *testing.T) {
			in := AdminFacilityIn{
				RequestCommon: testAdminRequestCommon(AdminFacilityAction),
				Cmd:           zCrud.CmdUpsert,
				Facility: rqProperty.Facilities{
					Id:             out.Facility.Id,
					FacilityName:   `testFacilityUpdated`,
					ExtraChargeIDR: 2000,
					FacilityType:   mProperty.FacilityTypeRoom,
					DescriptionEN:  ``,
				},
			}

			out := d.AdminFacility(&in)

			assert.True(t, out.HasError(), `failed to update facility: `+out.Error)
			fmt.Println(X.ToJsonPretty(out))
		})
	})
}
