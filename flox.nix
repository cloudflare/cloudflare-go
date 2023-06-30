{
  # flox environment
  #
  # To learn basics about flox commands see:
  #   https://floxdev.com/docs/basics
  #
  # Check other options you can configure in this file:
  #   https://floxdev.com/docs/reference/flox-nix-config
  #
  # Get help:
  #   https://discourse.floxdev.com
  #
  # Happy hacking!

  # Look for more packages with `flox search` command.
  #packages.nixpkgs-flox.go = {};
  #packages.nixpkgs-flox.nodejs = {}

  # Set environment variables
  environmentVariables.YOLOCOPTER = "jacob";

  # Run shell hook when you enter flox environment
  shell.hook = ''
    echo "Welcome to flox environment"
  '';
}
