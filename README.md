# gobbeldygook: generate readable nonsense

`gobbeldygook` is a small Go command‑line tool that generates pseudo‑readable text (nonsense) using a simple Markov‑chain algorithm. It learns word transitions from a provided text file and then produces a line of a specified length.

## Features

- Generates text that mimics the style of the source material.
- Configurable number of words per generated line.
- Simple, single‑binary Go implementation.

## Installation

```sh
git clone https://github.com/yourusername/gobbeldygook.git
cd gobbeldygook
go build -o gobbeldygook .
```

## Usage

```sh
./gobbeldygook -path <source.txt> -nums <word_count>
```

- `-path` – Path to a `.txt` file containing the source paragraphs the generator will learn from. **Required**.
- `-nums` – Number of words the generated line should contain (default: 50). Must be greater than 0.

### Example

```sh
./gobbeldygook -path samples/war_and_peace.txt -nums 80
```

The program will read `war_and_peace.txt`, build a word‑pair map, and output an 80‑word line of generated text.

## How It Works

1. **Read Input** – The source file is read and split into individual words.
2. **Build Word Map** – A map of `WordPair` (two consecutive words) to possible following words is constructed.
3. **Choose Starting Pair** – A random pair whose first word starts with an uppercase letter is selected.
4. **Generate** – The generator walks the map, picking random successors until the desired word count is reached.
5. **Wrap Output** – The generated line is wrapped to 90 characters for readability.

## Contributing

Feel free to open issues or submit pull requests. When contributing:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Make your changes and ensure the code builds.
4. Open a pull request describing your changes.

## License

This project is licensed under the MIT License – see the [LICENSE](LICENSE) file for details.