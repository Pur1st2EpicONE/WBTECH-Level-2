## L2.8

This Go snippet demonstrates a utility that fetches the current time from an NTP server and formats it according to a YAML configuration, with configurable time output, optional display of nanoseconds, timezone info, and safe handling of external input.

The program loads config.yaml using Viper, which defines the NTP server and formatting options such as showing the date, nanoseconds, UTC offset, and timezone name. If show_all is true, the full time.Time value is printed; otherwise, the output is customized. 

Errors in loading the config or fetching NTP time cause the program to exit with distinct codes (1 for config errors, 2 for server errors). The format function builds the output string efficiently with strings.Builder, appending each component according to the selected flags.

The utility can be run using **make**, which builds and runs the program, and the code can be checked with **make lint**, which runs golangci-lint to detect style issues and potential problems.
