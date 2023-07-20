"""Command line interface for entelecheia"""

# Importing the libraries
import sys

from hyfi import about, hydra_main  # type: ignore


def main() -> None:
    """Main function for the CLI"""
    sys.argv.append(f"--config-path={about.config_path}")
    hydra_main()


if __name__ == "__main__":
    main()
