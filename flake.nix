{
  description = "A very basic flake";

  outputs = { self, nixpkgs }:
  let
    mkDevShell = system:
    let
      pkgs = nixpkgs.legacyPackages.${system};
    in
    pkgs.mkShell {
      nativeBuildInputs = with pkgs; [
        go
        gopls
      ];
    };
  in
  {
    devShells.x86_64-linux.default = mkDevShell "x86_64-linux";
  };
}
