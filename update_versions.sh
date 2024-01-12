#!/bin/bash

# simple script to fetch last version of `replication-adapter-lib` repo
# and correctly substitute hashes in `flake.nix`

green_col=`tput setaf 2`
reset_col=`tput sgr0`

for repo in replication-adapter-lib replication-adapter
do
  VERSIONS=$(nix flake prefetch github:NilFoundation/$repo 2>&1 | sed -e 's/.*'$repo'\/\([a-z0-9]*\)'"' "'.*hash '"'"'\(.*\)'"'"'.*/\1 \2/g')
  COMMIT=$(echo $VERSIONS | awk '{ print $1; }')
  HASH=$(echo $VERSIONS | awk '{ print $2; }' 2>&1 | sed -e 's/\//\\\//g')

  echo "updating $repo version in flake.nix..."
  echo "${green_col}commit = $COMMIT, hash = $HASH${reset_col}"

  sed -e :n -e '$!N;/\n.*\n.*\n/!{$!bn
  };  s/\(.*'$repo'.git.*rev = "\)[^\n]*\(".* narHash = "\).*\(";.*\)/\1'$COMMIT'\2'$HASH'\3/;P;D' flake.nix > flake.nix.patched && mv flake.nix.patched flake.nix
done

echo "calculating vendor hash for go pkgs in nix (takes a while)..."
sed -i '' 's/^\([^a-z]*vendorHash = \)".*";/\1"";/g' flake.nix
VENDOR_HASH=$(nix build 2>&1 | grep "got:" | awk '{ print $2; }')
echo "${green_col}vendorHash = $VENDOR_HASH${reset_col}"
VENDOR_HASH=$(echo $VENDOR_HASH 2>&1 | sed -e 's/\//\\\//g')

sed -i '' 's/^\([^a-z]*vendorHash = \)".*";/\1"'$VENDOR_HASH'";/g' flake.nix
