#!/bin/bash

print_usage() {
    echo -e "Usage:\n\t $0 <major | minor | patch>\n" >&2
}

if [ -z "$1" ]; then
    echo "ERROR: missing parameter" >&2
    print_usage
    exit 1
fi

if [ "$(git branch --show-current)" != "main" ]; then
    echo "ERROR: Must be on main branch." >&2
    exit 1
fi

LAST_TAG=$(git tag -l 'v[0-9]*.[0-9]*.[0-9]*' | sort -V | tail -n 1)

if [ -z "$LAST_TAG" ]; then
    LAST_TAG="v0.0.0"
fi

NEXT_TAG=`echo "$LAST_TAG" | awk -F'[v.]' -v type="$1" '
{
  major = $2; minor = $3; patch = $4;
  
  if (type == "major") {
    major++; minor = 0; patch = 0;
  } else if (type == "minor") {
    minor++; patch = 0;
  } else if (type == "patch") {
    patch++;
  } else {
    print "ERROR: Invalid type";
    exit 1;
  }
  
  printf "v%d.%d.%d\n", major, minor, patch;
}'`

if [[ "$NEXT_TAG" == ERROR* ]]; then
    echo $NEXT_TAG >&2
    print_usage
    exit 1
fi
echo $NEXT_TAG

git checkout -b release/$NEXT_TAG
./script/generate_readme.md $NEXT_TAG
git add README.md
git add version.txt
git commit -a -m \"💾 Prepare release $NEXT_TAG\"
git push --set-upstream origin release/$NEXT_TAG
git checkout main
echo The release/$NEXT_TAG is ready. You can open a pull request on GitHub.

