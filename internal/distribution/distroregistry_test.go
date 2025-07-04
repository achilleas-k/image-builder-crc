package distribution

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/osbuild/image-builder-crc/internal/common"
)

func TestDistroRegistry_List(t *testing.T) {
	allDistros := []string{
		"rhel-8",
		"rhel-8-nightly",
		"rhel-84",
		"rhel-85",
		"rhel-86",
		"rhel-87",
		"rhel-88",
		"rhel-89",
		"rhel-8.10",
		"rhel-9",
		"rhel-9-beta",
		"rhel-9-nightly",
		"rhel-9.6-nightly",
		"rhel-9.7-nightly",
		"rhel-90",
		"rhel-91",
		"rhel-92",
		"rhel-93",
		"rhel-94",
		"rhel-95",
		"rhel-9.6",
		"centos-9",
		"rhel-10",
		"rhel-10-beta",
		"rhel-10-nightly",
		"rhel-10.0-nightly",
		"rhel-10.1-nightly",
		"rhel-10.0",
		"centos-10",
		"fedora-37",
		"fedora-38",
		"fedora-39",
		"fedora-40",
		"fedora-41",
		"fedora-42",
	}
	notEntitledDistros := []string{
		"rhel-8-nightly",
		"rhel-9-nightly",
		"rhel-9.6-nightly",
		"rhel-9.7-nightly",
		"rhel-10-nightly",
		"rhel-10.0-nightly",
		"rhel-10.1-nightly",
		"centos-9",
		"centos-10",
		"fedora-37",
		"fedora-38",
		"fedora-39",
		"fedora-40",
		"fedora-41",
		"fedora-42",
	}

	dr, err := LoadDistroRegistry("../../distributions")
	require.NoError(t, err)

	result := dr.Available(true).List()
	require.Len(t, result, len(allDistros))
	for _, distro := range result {
		require.Contains(t, allDistros, distro.Distribution.Name)
	}

	result = dr.Available(false).List()
	require.Len(t, result, len(notEntitledDistros))
	for _, distro := range result {
		require.Contains(t, notEntitledDistros, distro.Distribution.Name)
	}
}

func TestDistroRegistry_Get(t *testing.T) {
	dr, err := LoadDistroRegistry("../../distributions")
	require.NoError(t, err)

	result, err := dr.Available(true).Get("rhel-9")
	require.Equal(t, "rhel-9.6", result.Distribution.Name)
	require.Nil(t, err)

	// don't test packages, they are huge
	result.ArchX86.Packages = nil
	result.Aarch64.Packages = nil

	require.Equal(t, &DistributionFile{
		ModulePlatformID: "platform:el9",
		OscapName:        "rhel9",
		Distribution: DistributionItem{
			Description:      "Red Hat Enterprise Linux (RHEL) 9",
			Name:             "rhel-9.6",
			ComposerName:     common.ToPtr("rhel-9.6"),
			RestrictedAccess: false,
			NoPackageList:    true,
		},
		ArchX86: &Architecture{
			ImageTypes: []string{"aws", "gcp", "azure", "rhel-edge-commit", "rhel-edge-installer", "edge-commit", "edge-installer", "guest-image", "image-installer", "oci", "openshift-virt", "vsphere", "vsphere-ova", "wsl"},
			Repositories: []Repository{
				{
					Id:            "baseos",
					Baseurl:       common.ToPtr("https://cdn.redhat.com/content/dist/rhel9/9/x86_64/baseos/os"),
					Rhsm:          true,
					CheckGpg:      common.ToPtr(true),
					GpgKey:        common.ToPtr(rhelGpg),
					ImageTypeTags: nil,
				},
				{
					Id:            "appstream",
					Baseurl:       common.ToPtr("https://cdn.redhat.com/content/dist/rhel9/9/x86_64/appstream/os"),
					Rhsm:          true,
					CheckGpg:      common.ToPtr(true),
					GpgKey:        common.ToPtr(rhelGpg),
					ImageTypeTags: nil,
				},
				{
					Id:            "google-compute-engine",
					Baseurl:       common.ToPtr("https://packages.cloud.google.com/yum/repos/google-compute-engine-el9-x86_64-stable"),
					Rhsm:          false,
					CheckGpg:      common.ToPtr(true),
					GpgKey:        common.ToPtr(googleSdkGpg),
					ImageTypeTags: []string{"gcp"},
				},
				{
					Id:            "google-cloud-sdk",
					Baseurl:       common.ToPtr("https://packages.cloud.google.com/yum/repos/cloud-sdk-el9-x86_64"),
					Rhsm:          false,
					CheckGpg:      common.ToPtr(true),
					GpgKey:        common.ToPtr(googleSdkGpg),
					ImageTypeTags: []string{"gcp"},
				},
			},
		},
		Aarch64: &Architecture{
			ImageTypes: []string{"aws", "guest-image", "image-installer", "openshift-virt"},
			Repositories: []Repository{
				{
					Id:            "baseos",
					Baseurl:       common.ToPtr("https://cdn.redhat.com/content/dist/rhel9/9/aarch64/baseos/os"),
					Rhsm:          true,
					CheckGpg:      common.ToPtr(true),
					GpgKey:        common.ToPtr(rhelGpg),
					ImageTypeTags: nil,
				},
				{
					Id:            "appstream",
					Baseurl:       common.ToPtr("https://cdn.redhat.com/content/dist/rhel9/9/aarch64/appstream/os"),
					Rhsm:          true,
					CheckGpg:      common.ToPtr(true),
					GpgKey:        common.ToPtr(rhelGpg),
					ImageTypeTags: nil,
				},
			},
		},
	}, result)

	result, err = dr.Available(false).Get("toucan-42")
	require.Nil(t, result)
	require.Equal(t, ErrDistributionNotFound, err)
}
