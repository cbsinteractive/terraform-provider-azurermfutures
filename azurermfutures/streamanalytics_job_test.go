package azurermfutures

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccStreamAnalyticsJob_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProviderFactories,
		CheckDestroy:             testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccExampleResource(),
				Check:  testAccCheckExampleResourceExists,
			},
		},
	})
}

func testAccPreCheck(*testing.T) {
	log.Println("[DEBUG] Inside testAccPreCheck")
}

func testAccCheckExampleResourceDestroy(*terraform.State) error {
	log.Println("[DEBUG] Inside testAccCheckExampleResourceDestroy")
	return nil
}

func testAccCheckExampleResourceExists(s *terraform.State) error {
	log.Println("[DEBUG] Inside testAccCheckExampleResourceExists")
	log.Printf("[DEBUG] state: %+v\n", s)
	job := s.RootModule().Resources["azurermfutures_streamanalytics_job.example"]
	log.Printf("[DEBUG] job: %+v\n", job)
	log.Printf("[DEBUG] job ID: %+v\n", job.Primary.ID)
	if job.Primary.ID != "some-id" {
		return fmt.Errorf("Expected ID %s; got %s", "some-id", job.Primary.ID)
	}
	return nil
}

func testAccExampleResource() string {
	log.Println("[DEBUG] Inside testAccExampleResource")
	return `
	resource "azurermfutures_streamanalytics_job" "example" {}
`
}
