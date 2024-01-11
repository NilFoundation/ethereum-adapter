{
  description = "NIX dev env for replication-adapter erigon part";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    (flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};

        replication-adapter-lib = builtins.fetchGit {
          url = "git@github.com:NilFoundation/replication-adapter-lib.git";
          rev = "c583de36f7a539fc5fb7151773bff368c4872e1f";
          narHash = "sha256-jxd7w+vszp9wLs/p0iSazKqHpJWdxZh4z3p7IL5lFII=";
        };

        replication-adapter = builtins.fetchGit {
          url = "git@github.com:NilFoundation/replication-adapter.git";
          rev = "6db7506cb8cf0e6337a3494afcb79634962cf21f";
          narHash = "sha256-n2kRpZxCBHIKbHyaS5QvSxUaepZgySpdkzTW/H6CR28=";
        };

      in rec

      {
        packages.default = pkgs.buildGoModule {
          name = "erigon";
          src = ./.;

          # to obtain run `nix build` with vendorHash = "";
          proxyVendor = true;
          vendorHash = "sha256-4XTimF2biLXC0muAGEFe5SpblwrVWNmf3NzSzTX2NAs=";

          # skip testing
          doCheck = false;

          preBuild = ''
            ln -sf ${replication-adapter} ./adapter
            ln -sf ${replication-adapter-lib} ./adapterlib

            go mod edit -replace=github.com/NilFoundation/replication-adapter=./adapter
            go mod edit -replace=github.com/NilFoundation/replication-adapter-lib=./adapterlib
          '';


          # ======= DIRTY HACK (?) ========
          # we only need erigon binary, so skip rest
          # also full build with nix fails on ./erigon-lib
          # cause it's separate module

          buildPhase = ''
            runHook preBuild

            buildGoDir() {
              local dir="$1"

              . $TMPDIR/buildFlagsArray

              declare -a flags
              flags+=($buildFlags "''${buildFlagsArray[@]}")
              flags+=(''${tags:+-tags=''${tags// /,}})
              flags+=(''${ldflags:+-ldflags="$ldflags"})
              flags+=("-p" "$NIX_BUILD_CORES")

              local OUT
              if ! OUT="$(go install "''${flags[@]}" $dir 2>&1)"; then
                if ! echo "$OUT" | grep -qE '(no( buildable| non-test)?|build constraints exclude all) Go (source )?files'; then
                  echo "$OUT" >&2
                  return 1
                fi
              fi
              if [ -n "$OUT" ]; then
                echo "$OUT" >&2
              fi
              return 0
            }

            if (( "''${NIX_DEBUG:-0}" >= 1 )); then
              buildFlagsArray+=(-x)
            fi

            if [ ''${#buildFlagsArray[@]} -ne 0 ]; then
              declare -p buildFlagsArray > $TMPDIR/buildFlagsArray
            else
              touch $TMPDIR/buildFlagsArray
            fi
            if [ -z "$enableParallelBuilding" ]; then
                export NIX_BUILD_CORES=1
            fi

            buildTests() {
              local dir="$1" binary_name="$2"

              go test -c -o $GOPATH/bin/$binary_name $dir
            }

            echo "Building go package ./cmd/erigon"

            buildGoDir "./cmd/erigon"

            buildTests "./cmd/devnet/tests/generic/devnet_state_test.go" "devnet_state_test"
            buildTests "./cmd/devnet/tests/generic/devnet_infinite_run_test.go" "devnet_infinite_run_test"
          '';
        };

        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gotools
            go-tools
          ];

          shellHook = ''
            export GO_CFG_DIR=$HOME/.nix/go/$(basename ${self})

            mkdir -p $GO_CFG_DIR/config $GO_CFG_DIR/cache $GO_CFG_DIR/pkg/mod

            export GOENV="$GO_CFG_DIR/config/env"

            go env -w GOCACHE="$GO_CFG_DIR/cache"
            go env -w GOMODCACHE="$GO_CFG_DIR/pkg/mod"

            go env -w GOPRIVATE=github.com/NilFoundation

            go mod tidy
          '';
        };

        overlays.default = final: prev: {
          erigon = packages.default;
        };
      })
    );
}
