with (import <nixpkgs> {});

with pkgs; stdenv.mkDerivation {
  name = "server-side-rendering";
  buildInputs = [ nodejs yarn ];
}
