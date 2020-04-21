#!/bin/bash

NAT='0|[1-9][0-9]*'
ALPHANUM='[0-9]*[A-Za-z-][0-9A-Za-z-]*'
IDENT="$NAT|$ALPHANUM"
FIELD='[0-9A-Za-z-]+'

SEMVER_REGEX="\
^[vV]\
($NAT)\\.($NAT)\\.($NAT)\
(\\-(${IDENT})(\\.(${IDENT}))*)?\
(\\+${FIELD}(\\.${FIELD})*)?$"

semver_compare() {
  local version_a version_b pr_a pr_b
  version_a=$(echo "${1//v/}" | awk -F'-' '{print $1}')
  version_b=$(echo "${2//v/}" | awk -F'-' '{print $1}')

  if [ "$version_a" \= "$version_b" ]; then
    pr_a=$(echo "$1" | awk -F'-' '{print $2}')
    pr_b=$(echo "$2" | awk -F'-' '{print $2}')

    [ "$pr_a" \= "$pr_b" ] && echo 0 && return 0

    if [ -z "$pr_a" ]; then
      echo 1 && return 0
    fi

    number_a=$(echo ${pr_a//[!0-9]/})
    number_b=$(echo ${pr_b//[!0-9]/})
    [ -z "${number_a}" ] && number_a=0
    [ -z "${number_b}" ] && number_b=0

    [ "$pr_a" \> "$pr_b" ] && [ -n "$pr_b" ] && [ "$number_a" -gt "$number_b" ] && echo 1 && return 0

    echo -1 && return 0
  fi
  arr_version_a=(${version_a//./ })
  arr_version_b=(${version_b//./ })
  cursor=0
  # Iterate arrays from left to right and find the first difference
  while [ "$([ "${arr_version_a[$cursor]}" -eq "${arr_version_b[$cursor]}" ] && [ $cursor -lt ${#arr_version_a[@]} ] && echo true)" == true ]; do
    cursor=$((cursor + 1))
  done
  [ "${arr_version_a[$cursor]}" -gt "${arr_version_b[$cursor]}" ] && echo 1 || echo -1
}

get_latest_version() {
  git fetch --tags
  echo "$(git tag | grep 'v' | tr - \~ | sort -V | tr \~ - | tail -1)"
}

latestVersion=$(get_latest_version)

echo The latest tag: "$latestVersion"

read -p 'Enter new version (e.g: v0.0.1): ' version

if [[ ! "$version" =~ $SEMVER_REGEX ]]; then
  echo "Invalid version"
  exit
fi

compare=$(semver_compare $version "$latestVersion")

if [ $compare -le 0 ]; then
  echo "Version should be greater then $latestVersion"
  exit
fi

git tag $version

git push origin $version
