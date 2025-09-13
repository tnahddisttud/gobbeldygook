# gobbeldygook

**gobbeldygook** is a tiny Go command‑line tool that generates readable nonsense text based on a source paragraph. It builds a second‑order Markov chain from the input file and then produces a line of a configurable length.

## Features

- Reads any plain‑text file and learns word transitions.
- Generates a line with a specified number of words (`-nums`).
- Starts sentences with capitalised words for a more natural feel.
- Wraps output to 90 characters for easy reading.

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/gobbeldygook.git
cd gobbeldygook

# Build the binary
go build -o gobbeldygook
```

Alternatively, you can install it directly with `go install`:

```bash
go install github.com/yourusername/gobbeldygook@latest
```

## Usage

```bash
./gobbeldygook -path <input.txt> -nums <word_count>
```

- `-path` – Path to a `.txt` file containing the source paragraphs.
- `-nums` – Number of words to generate per line (default: 50).

### Example

Assuming you have a file `sample.txt` with some text:

```bash
./gobbeldygook -path sample.txt -nums 30
```

The program will output a line of 30 words that resembles the style of the source text.

## How It Works

1. **Reading Input** – The program reads the entire file and splits it into words.
2. **Building the Markov Map** – It creates a map where each key is a pair of consecutive words (`WordPair`). The value is a slice of possible next words that followed this pair in the source.
3. **Choosing a Start** – It selects a random key whose first word starts with an uppercase letter, ensuring the generated line begins like a proper sentence.
4. **Generating Text** – Using the map, it repeatedly picks a random next word based on the current word pair until the requested number of words is reached or no continuation exists.
5. **Wrapping** – The final string is wrapped at 90 characters for readability.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License – see the [LICENSE](LICENSE) file for details.