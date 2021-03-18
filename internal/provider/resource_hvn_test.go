package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-hcp/internal/clients"
)

var (
	hvn = fmt.Sprintf(`
resource "hcp_hvn" "test" {
	hvn_id         = "test-hvn"
	cloud_provider = "aws"
	region         = "us-west-2"
}`)
)

func TestAccHvn(t *testing.T) {
	resourceName := "hcp_hvn.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckHvnDestroy,
		Steps: []resource.TestStep{
			{
				Config: testConfig(hvn),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHvnExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "hvn_id", "test-hvn"),
					resource.TestCheckResourceAttr(resourceName, "cloud_provider", "aws"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west-2"),
					resource.TestCheckResourceAttrSet(resourceName, "cidr_block"),
					resource.TestCheckResourceAttrSet(resourceName, "organization_id"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "provider_account_id"),
				),
			},
			{
				ResourceName: resourceName,
				ImportState:  true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources[resourceName]
					if !ok {
						return "", fmt.Errorf("not found: %s", resourceName)
					}

					return rs.Primary.Attributes["hvn_id"], nil
				},
				ImportStateVerify: true,
			},
			{
				Config: testConfig(hvn),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHvnExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "hvn_id", "test-hvn"),
					resource.TestCheckResourceAttr(resourceName, "cloud_provider", "aws"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west-2"),
					resource.TestCheckResourceAttrSet(resourceName, "cidr_block"),
					resource.TestCheckResourceAttrSet(resourceName, "organization_id"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "provider_account_id"),
				),
			},
		},
	})
}

func testAccCheckHvnExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}

		id := rs.Primary.ID
		if id == "" {
			return fmt.Errorf("no ID is set")
		}

		client := testAccProvider.Meta().(*clients.Client)

		link, err := buildLinkFromURL(id, HvnResourceType, client.Config.OrganizationID)
		if err != nil {
			return fmt.Errorf("unable to build link for %q: %v", id, err)
		}

		hvnID := link.ID
		loc := link.Location

		if _, err := clients.GetHvnByID(context.Background(), client, loc, hvnID); err != nil {
			return fmt.Errorf("unable to read HVN %q: %v", id, err)
		}

		return nil
	}
}

func testAccCheckHvnDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*clients.Client)

	for _, rs := range s.RootModule().Resources {
		switch rs.Type {
		case "hcp_hvn":
			id := rs.Primary.ID

			link, err := buildLinkFromURL(id, HvnResourceType, client.Config.OrganizationID)
			if err != nil {
				return fmt.Errorf("unable to build link for %q: %v", id, err)
			}

			hvnID := link.ID
			loc := link.Location

			_, err = clients.GetHvnByID(context.Background(), client, loc, hvnID)
			if err == nil || !clients.IsResponseCodeNotFound(err) {
				return fmt.Errorf("didn't get a 404 when reading destroyed HVN %q: %v", id, err)
			}

		default:
			continue
		}
	}
	return nil
}