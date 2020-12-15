package nutanix

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccNutanixKarbonClustersDataSource_basic(t *testing.T) {
	r := acctest.RandInt()
	//resourceName := "nutanix_karbon_cluster.cluster"
	subnetName := "Rx-Automation-Network"
	defaultContainter := "default-container-85827904983728"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccKarbonClustersDataSourceConfig(subnetName, r, defaultContainter, 1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.nutanix_karbon_clusters.kclusters", "clusters.#"),
				),
			},
		},
	})
}

func testAccKarbonClustersDataSourceConfig(subnetName string, r int, containter string, workers int) string {
	return testAccNutanixKarbonClusterConfig(subnetName, r, containter, workers) + `
	data "nutanix_karbon_clusters" "kclusters" {}

	`
}
