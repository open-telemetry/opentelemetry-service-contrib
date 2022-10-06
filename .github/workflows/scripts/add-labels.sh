#!/usr/bin/env bash
#
#   Copyright The OpenTelemetry Authors.
#
#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#       http://www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.
#
#

if [ -z "${ISSUE}" ] || [ -z "${COMMENT}" ]; then
    exit 0
fi

declare -A labels
labels["/good-first-issue"]="good first issue"
labels["/help-wanted"]="help wanted"
labels["/never-stale"]="never stale"

if [ -n "${labels["${COMMENT}"]}" ] ; then
    gh issue edit "${ISSUE}" --add-label "${labels["${COMMENT}"]}"
fi
