{ pkgs, ... }: {
  packages = [
    pkgs.go
    pkgs.gitmoji-cli
    pkgs.golangci-lint
  ];
 
  idx = {
    extensions = [
      "golang.go"
      "PKief.material-icon-theme"
    ];
  };
}