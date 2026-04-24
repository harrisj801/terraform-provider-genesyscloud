package external_user

import (
	"fmt"
	"net/url"
	"strings"
)

//	func createCompoundKey(userId, authorityName, externalKey string) string {
//		return fmt.Sprintf("%s/%s/%s", userId, authorityName, externalKey)
//	}
func createCompoundKey(userId, authorityName, externalKey string) string {
	return fmt.Sprintf("%s/%s/%s",
		url.QueryEscape(userId),
		url.QueryEscape(authorityName),
		url.QueryEscape(externalKey),
	)
}

//	func splitCompoundKey(compoundKey string) (string, string, string, error) {
//		split := strings.Split(compoundKey, "/")
//		if len(split) != 3 {
//			return "", "", "", fmt.Errorf("invalid compound key: %s", compoundKey)
//		}
//		return split[0], split[1], split[2], nil
//	}
func splitCompoundKey(compoundKey string) (string, string, string, error) {
	split := strings.Split(compoundKey, "/")
	if len(split) != 3 {
		return "", "", "", fmt.Errorf("invalid compound key: %s", compoundKey)
	}

	u, err := url.QueryUnescape(split[0])
	if err != nil {
		return "", "", "", err
	}
	a, err := url.QueryUnescape(split[1])
	if err != nil {
		return "", "", "", err
	}
	e, err := url.QueryUnescape(split[2])
	if err != nil {
		return "", "", "", err
	}

	return u, a, e, nil
}
func generateExternalUserIdentity(resourceLabel, userId, authorityName, externalKey string) string {
	return fmt.Sprintf(`resource "genesyscloud_externalusers_identity" "%s" {
        user_id = %s
        authority_name = "%s"
        external_key = "%s"
	}
	`, resourceLabel, userId, authorityName, externalKey)
}
