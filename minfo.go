package minfo

import "os"
import "runtime"

const (
	UNKNOWN = "unkown"
)

func OS() string {
	if runtime.GOOS == "windows" {
		return "Windows"
	}
	
	f, err := os.Open("/etc")
	if err != nil {
		// FIXME(ttacon): use other methods
		return UNKNOWN
	}

	names, err := f.Readdirnames(-1)
	if err != nil {
		// well that sucks
		return UNKNOWN
	}

	for _,entry := range names {
		if os, ok := osFileMappings[entry]; ok {
			return os
		}
	}

	return UNKNOWN
}


var (
	osFileMappings = map[string]string{
		"arch-release": "Arch Linux",
		"chakra-release": "Chakra",
		"crunchbang-lsb-release": "CrunchBang",
		"debian_version": "Debian", // TODO(ttacon): https://github.com/KittyKatt/screenFetch/blob/master/screenfetch-dev#L625
		"dragora-version": "Dragora",
		"evolveos-relase": "Evolve OS",
		"fedora-release": "Fedora", // TODO(ttacon): https://github.com/KittyKatt/screenFetch/blob/master/screenfetch-dev#L636
		"frugalware-release": "FrugalWare",
		"gentoo-release": "Gentoo",// TODO(ttacon): https://github.com/KittyKatt/screenFetch/blob/master/screenfetch-dev#L646
		"mageia-release": "Mageia",
		"mandrake-release": "Mandrake", // TODO(ttacon): https://github.com/KittyKatt/screenFetch/blob/master/screenfetch-dev#L653
		"mandriva-release": "Mandriva", // TODO(ttacon): https://github.com/KittyKatt/screenFetch/blob/master/screenfetch-dev#L659
		"NIXOS": "NixOS",
		"SuSE-release": "openSUSE",
		"pclinuxos-release": "PCLinuxOS",
		"redhat-release": "Red Hat Enterprise Linux", // TODO(ttacon): https://github.com/KittyKatt/screenFetch/blob/master/screenfetch-dev#L668
		"slackware-release": "Slackware",
		"sabayon-release": "Sabayon",
	}
)
