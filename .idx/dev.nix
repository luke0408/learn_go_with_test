{ pkgs, ... }: {
  packages = [
    pkgs.go
    pkgs.gitmoji-cli
  ];
 
  idx = {
    extensions = [
      "golang.go"
      "PKief.material-icon-theme"
    ];
  };
}