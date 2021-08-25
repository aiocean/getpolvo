package server

import "regexp"

func parseVersionOrn (orn string) (packageName string, versionName string){
	var packagesOrnRegex = regexp.MustCompile(`packages\/([^\/]+)\/versions\/([^\/]+)`)
	matches := packagesOrnRegex.FindStringSubmatch(orn)
	if len(matches) < 3 {
		return "", ""
	}

	return matches[1], matches[2]
}
