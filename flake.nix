{
  description = "NIX dev env for replication-adapter state keeper";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

    flake-utils.url = "github:numtide/flake-utils";

    gomod2nix = {
      url = "github:tweag/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.flake-utils.follows = "flake-utils";
    };
  };

  outputs = { self, nixpkgs, flake-utils, gomod2nix }:
    (flake-utils.lib.eachDefaultSystem
      (system:
        let
          pkgs = nixpkgs.legacyPackages.${system};

          # The current default sdk for macOS fails to compile go projects, so we use a newer one for now.
          # This has no effect on other platforms.
          callPackage = pkgs.darwin.apple_sdk_11_0.callPackage or pkgs.callPackage;
        in
        {
          packages.default = callPackage ./. {
            inherit (gomod2nix.packages.${system}) buildGoApplication;
          };

          devShells.default = pkgs.mkShell {
            buildInputs = with pkgs; [
              go
              gotools
              go-tools
              gomod2nix.packages.${system}.default
            ];

            shellHook = ''
              mkdir -p $(go env GOROOT)/.config $(go env GOROOT)/.cache $(go env GOROOT)/pkg/mod
              export GOENV="$(go env GOROOT)/.config/env"

              go env -w GOCACHE="$(go env GOROOT)/.cache"
              go env -w GOMODCACHE="$(go env GOROOT)/pkg/mod"

              CUR_DIR=$(pwd); cd ${self} && go mod tidy && cd $CUR_DIR
            '';
          };
        })
    );
}