# data-parser

## Description

`data-parser` is a versatile command-line tool and library designed for parsing and transforming various data formats into a standardized JSON output. It aims to simplify data ingestion into systems that require structured data, providing a flexible and extensible solution for handling diverse data sources. Whether you're dealing with CSV files, log files, or custom-formatted text, `data-parser` offers a streamlined approach to data extraction and normalization.

## Features

*   **Format Support:**
    *   CSV (Comma Separated Values)
    *   TSV (Tab Separated Values)
    *   JSON (JavaScript Object Notation)
    *   TXT (Custom Delimited Text - configurable delimiters)
    *   Log Files (Configurable regex-based parsing for various log formats)
*   **Schema Definition:**
    *   Supports defining schemas for data validation and transformation.
    *   Schema definitions can specify data types, default values, and transformations.
    *   Allows renaming fields during parsing.
*   **Data Transformation:**
    *   Ability to apply transformations to data values during parsing (e.g., date formatting, numerical conversions).
    *   Supports custom transformation functions via plugins.
*   **Command-Line Interface (CLI):**
    *   Easy-to-use CLI for quick data parsing and transformation.
    *   Supports piping data from standard input.
    *   Options for specifying input file, output file, schema file, and other configuration parameters.
*   **Library Usage:**
    *   Can be used as a library in other Python projects for programmatic data parsing.
    *   Provides a well-defined API for customizing the parsing process.
*   **Error Handling:**
    *   Robust error handling with informative error messages.
    *   Option to skip invalid records or halt processing on errors.
*   **Extensible Plugin Architecture:**
    *   Allows users to create custom parsers and transformers through a plugin system.
*   **Comprehensive Documentation:**
    *   Detailed documentation with examples and usage instructions.

## Technologies Used

*   **Python:** Core implementation language.
*   **`argparse`:** For command-line argument parsing.
*   **`csv`:** For CSV parsing.
*   **`json`:** For JSON parsing and output.
*   **`re`:** For regular expression based parsing.
*   **`pluggy`:** For plugin management and extensibility.
*   **`pytest`:** For unit testing.
*   **`tox`:** For testing across multiple Python versions.

## Installation

### Prerequisites

*   Python 3.7 or higher.
*   `pip` package installer.

### Installing with pip

```bash
pip install data-parser
```

### Installing from Source (for development)

1.  Clone the repository:

    ```bash
    git clone https://github.com/your-username/data-parser.git
    cd data-parser
    ```

2.  Create a virtual environment (recommended):

    ```bash
    python3 -m venv venv
    source venv/bin/activate  # On Linux/macOS
    # venv\Scripts\activate   # On Windows
    ```

3.  Install the dependencies:

    ```bash
    pip install -e .
    pip install -r requirements-dev.txt # for development dependencies
    ```

## Usage

### CLI Usage

```bash
data-parser --input input.csv --output output.json --schema schema.json --format csv
```

**Options:**

*   `--input`: Path to the input file.
*   `--output`: Path to the output file (defaults to stdout).
*   `--schema`: Path to the schema file.
*   `--format`: Input data format (csv, tsv, json, txt, log).  Defaults to `txt`.
*   `--delimiter`: Delimiter for TXT format.
*   `--log-format`: Regex string defining the structure of log file entries.
*   `--skip-errors`:  If set, invalid records will be skipped instead of throwing an error.

For more detailed usage information, run:

```bash
data-parser --help
```

### Library Usage

```python
from data_parser.parser import DataParser
from data_parser.schema import Schema

# Example schema
schema_definition = {
    "fields": [
        {"name": "name", "type": "string"},
        {"name": "age", "type": "integer", "default": 0},
        {"name": "date", "type": "date", "format": "%Y-%m-%d"}
    ]
}

schema = Schema(schema_definition)
parser = DataParser(schema=schema, format="csv")

data = """
name,age,date
John Doe,30,2023-01-01
Jane Smith,,2023-02-01
"""

parsed_data = parser.parse(data)
print(parsed_data)

# Alternatively, parse from a file:
# with open("input.csv", "r") as f:
#     parsed_data = parser.parse(f)
#     print(parsed_data)
```

## Schema Definition

The schema file is a JSON file that defines the structure of the data being parsed. It contains a `fields` array, where each element describes a field in the data:

```json
{
  "fields": [
    {
      "name": "field1",
      "type": "string",
      "default": null
    },
    {
      "name": "field2",
      "type": "integer",
      "default": 0
    },
    {
      "name": "field3",
      "type": "date",
      "format": "%Y-%m-%d"
    },
    {
      "name": "field4",
      "type": "boolean",
      "default": false
    },
    {
      "name": "field5",
      "type": "float"
    }
  ]
}
```

**Field Properties:**

*   `name`: (Required) The name of the field.
*   `type`: (Required) The data type of the field (`string`, `integer`, `float`, `boolean`, `date`).
*   `default`: (Optional) The default value for the field if it's missing in the input data.
*   `format`: (Optional) The format string for date fields (e.g., `%Y-%m-%d`).

## Contributing

We welcome contributions to `data-parser`! Please see the `CONTRIBUTING.md` file for guidelines on how to contribute.

## License

[MIT License](LICENSE)