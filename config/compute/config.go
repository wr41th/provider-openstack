package compute

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("openstack_compute_instance_v2", func(r *config.Resource) {
		// r.Name = "openstack_compute_instance_v2"
		// r.ExternalName = config.ExternalName{}
		r.ShortGroup = "compute"
		r.References["key_pair"] = config.Reference{
			Type: "KeypairV2",
		}
		r.References["security_groups"] = config.Reference{
			Type: "SecgroupV2",
		}
		r.References["flavor_name"] = config.Reference{
			Type: "FlavorV2",
		}
		r.References["image_name"] = config.Reference{
			Type: "ImageV2",
		}
	})
	p.AddResourceConfigurator("openstack_compute_quotaset_v2", func(r *config.Resource) {
		r.References["project_id"] = config.Reference{
			Type: "github.com/wr41th/provider-openstack/apis/identity/v1alpha1.ProjectV3",
		}
	})
	p.AddResourceConfigurator("openstack_compute_keypair_v2", func(r *config.Resource) {

		// example for connection secrets
		// r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
		// 	a := map[string][]byte{
		// 		"example secret": []byte("test"),
		// 	}
		// 	return a, nil
		// }
	})
}
