// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/google/go-github/v58/github"
)

const allowlistHeader = `# Code generated by githubgen. DO NOT EDIT.
#####################################################
#
# List of components in OpenTelemetry Collector Contrib
# waiting on owners to be assigned
#
#####################################################
#
# Learn about membership in OpenTelemetry community:
#  https://github.com/open-telemetry/community/blob/main/community-membership.md
#
#
# Learn about CODEOWNERS file format:
#  https://help.github.com/en/articles/about-code-owners
#

## 
# NOTE: New components MUST have one or more codeowners. Add codeowners to the component metadata.yaml and run make gengithub
##

## COMMON & SHARED components
internal/common

`

const unmaintainedHeader = `

## UNMAINTAINED components

`

const codeownersHeader = `# Code generated by githubgen. DO NOT EDIT.
#####################################################
#
# List of codeowners for OpenTelemetry Collector Contrib
#
#####################################################
#
# Learn about membership in OpenTelemetry community:
# https://github.com/open-telemetry/community/blob/main/community-membership.md
#
#
# Learn about CODEOWNERS file format:
# https://help.github.com/en/articles/about-code-owners
#

* @open-telemetry/collector-contrib-approvers
`

const distributionCodeownersHeader = `
#####################################################
#
# List of distribution maintainers for OpenTelemetry Collector Contrib
#
#####################################################
`

type codeownersGenerator struct {
}

func (cg codeownersGenerator) generate(data *githubData) error {
	allowlistData, err := os.ReadFile(data.allowlistFilePath)
	if err != nil {
		return err
	}
	allowlistLines := strings.Split(string(allowlistData), "\n")

	allowlist := make(map[string]struct{}, len(allowlistLines))
	unusedAllowlist := make(map[string]struct{}, len(allowlistLines))
	for _, line := range allowlistLines {
		if line == "" {
			continue
		}
		allowlist[line] = struct{}{}
		unusedAllowlist[line] = struct{}{}
	}
	var missingCodeowners []string
	var duplicateCodeowners []string
	members, err := getGithubMembers()
	if err != nil {
		return err
	}
	for _, codeowner := range data.codeowners {
		_, present := members[codeowner]

		if !present {
			_, allowed := allowlist[codeowner]
			delete(unusedAllowlist, codeowner)
			allowed = allowed || strings.HasPrefix(codeowner, "open-telemetry/")
			if !allowed {
				missingCodeowners = append(missingCodeowners, codeowner)
			}
		} else if _, ok := allowlist[codeowner]; ok {
			duplicateCodeowners = append(duplicateCodeowners, codeowner)
		}
	}
	if len(missingCodeowners) > 0 {
		sort.Strings(missingCodeowners)
		return fmt.Errorf("codeowners are not members: %s", strings.Join(missingCodeowners, ", "))
	}
	if len(duplicateCodeowners) > 0 {
		sort.Strings(duplicateCodeowners)
		return fmt.Errorf("codeowners members duplicate in allowlist: %s", strings.Join(duplicateCodeowners, ", "))
	}
	if len(unusedAllowlist) > 0 {
		var unused []string
		for k := range unusedAllowlist {
			unused = append(unused, k)
		}
		sort.Strings(unused)
		return fmt.Errorf("unused members in allowlist: %s", strings.Join(unused, ", "))
	}

	codeowners := codeownersHeader
	deprecatedList := "## DEPRECATED components\n"
	unmaintainedList := "\n## UNMAINTAINED components\n"

	unmaintainedCodeowners := unmaintainedHeader
	currentFirstSegment := ""
LOOP:
	for _, key := range data.folders {
		m := data.components[key]
		for stability := range m.Status.Stability {
			if stability == unmaintainedStatus {
				unmaintainedList += key + "/\n"
				unmaintainedCodeowners += fmt.Sprintf("%s/%s @open-telemetry/collector-contrib-approvers \n", key, strings.Repeat(" ", data.maxLength-len(key)))
				continue LOOP
			}
			if stability == "deprecated" && (m.Status.Codeowners == nil || len(m.Status.Codeowners.Active) == 0) {
				deprecatedList += key + "/\n"
			}
		}

		if m.Status.Codeowners != nil {
			parts := strings.Split(key, string(os.PathSeparator))
			firstSegment := parts[0]
			if firstSegment != currentFirstSegment {
				currentFirstSegment = firstSegment
				codeowners += "\n"
			}
			owners := ""
			for _, owner := range m.Status.Codeowners.Active {
				owners += " "
				owners += "@" + owner
			}
			codeowners += fmt.Sprintf("%s/%s @open-telemetry/collector-contrib-approvers%s\n", key, strings.Repeat(" ", data.maxLength-len(key)), owners)
		}
	}

	codeowners += distributionCodeownersHeader
	longestName := 0
	for _, dist := range data.distributions {
		if longestName < len(dist.Name) {
			longestName = len(dist.Name)
		}
	}

	for _, dist := range data.distributions {
		var maintainers []string
		for _, m := range dist.Maintainers {
			maintainers = append(maintainers, fmt.Sprintf("@%s", m))
		}
		codeowners += fmt.Sprintf("reports/distributions/%s.yaml%s @open-telemetry/collector-contrib-approvers %s\n", dist.Name, strings.Repeat(" ", longestName-len(dist.Name)), strings.Join(maintainers, " "))
	}

	err = os.WriteFile(filepath.Join(".github", "CODEOWNERS"), []byte(codeowners+unmaintainedCodeowners), 0600)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(".github", "ALLOWLIST"), []byte(allowlistHeader+deprecatedList+unmaintainedList), 0600)
	if err != nil {
		return err
	}
	return nil
}

func getGithubMembers() (map[string]struct{}, error) {
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		return nil, fmt.Errorf("Set the environment variable `GITHUB_TOKEN` to a PAT token to authenticate")
	}
	client := github.NewTokenClient(context.Background(), githubToken)
	var allUsers []*github.User
	pageIndex := 0
	for {
		users, resp, err := client.Organizations.ListMembers(context.Background(), "open-telemetry",
			&github.ListMembersOptions{
				PublicOnly: false,
				ListOptions: github.ListOptions{
					PerPage: 50,
					Page:    pageIndex,
				},
			},
		)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if len(users) == 0 {
			break
		}
		allUsers = append(allUsers, users...)
		pageIndex++
	}

	usernames := make(map[string]struct{}, len(allUsers))
	for _, u := range allUsers {
		usernames[*u.Login] = struct{}{}
	}
	return usernames, nil
}
