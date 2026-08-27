{
  description = "Pokedex development environment";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      supportedSystems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];
      forEachSupportedSystem = f: nixpkgs.lib.genAttrs supportedSystems (system: f {
        pkgs = nixpkgs.legacyPackages.${system};
      });
    in
    {
      devShells = forEachSupportedSystem ({ pkgs }:
        {
          default = pkgs.mkShell {
            packages = with pkgs; [
              go
              golangci-lint
              gitlab-ci-ls
              nil
              fish
            ];
            shellHook = ''
            # Isolate Go binaries so bootdev is local to the project
            export GOBIN="$PWD/.nix/bin"
            export PATH="$GOBIN:$PATH"

            # Automatically install boot.dev CLI (bootdev) if missing
            if ! command -v bootdev &> /dev/null; then
            echo "Installing boot.dev CLI (bootdev) locally to $GOBIN..."
            mkdir -p "$GOBIN"
            go install github.com/bootdotdev/bootdev@latest
            fi

            echo "=== Pokedex Nix Dev Environment ==="
            echo "Available tools:"
            echo "  - go: $(go version)"
            echo "  - golangci-lint: $(golangci-lint --version)"
            echo "  - bootdev: $(bootdev --version 2>/dev/null || echo "Installed at $GOBIN/bootdev")"
            echo ""
            echo "============================="

            if [ -z "$IN_POKEDEX_DEVSHELL" ]; then
            export IN_POKEDEX_DEVSHELL=1
            exec fish
            fi

            '';
          };
        });
    };
}
