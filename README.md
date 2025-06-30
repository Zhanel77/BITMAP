## Bitmap

## Features

- **Header**: Read and display metadata from BMP files.

- **Mirror**: Mirror images horizontally or vertically.

- **Filter**: Apply various filters to images (e.g., blue, red, grayscale, negative).

- **Rotate**: Rotate images by specified angles.

- **Crop**: Trim images according to specified parameters.

- **Help**: Display usage information and command options.

## Usage
```bash
bitmap <command> [arguments]
```

## Commands
- **header**: Prints bitmap file header information.
```bash
./bitmap header <source_file>
```

- **apply**: Applies processing to the image and saves it to the file.
```bash
./bitmap apply [options] <source_file> <output_file>
```


## Flags
- `-h`, `--help`: Prints program usage information.

## Options

### Mirror 
- `--mirror`: flag flips a bitmap image **horizontally** (left to right) or **vertically** (top to bottom).

### Usage

```bash
./bitmap apply --mirror=<direction> [--mirror=<direction>] input.bmp output.bmp
```

###  Filters 
- `--filter` flag applies visual effects to a bitmap image. Multiple filters can be chained by specifying the `--filter` flag multiple times, and they are applied **in the order given**.

- `--filter`: Applies various filters to the image.
  - `--filter=blue`: Retains only the blue channel.
  - `--filter=red`: Retains only the red channel.
  - `--filter=green`: Retains only the green channel.
  - `--filter=grayscale`: Converts the image to grayscale.
  - `--filter=negative`: Applies a negative filter.
  - `--filter=pixelate`: Applies a pixelation effect (default block size: 20 pixels).
  - `--filter=blur`: Applies a blur effect.

### Rotate
- `--rotate` flag rotates a bitmap image by a specified angle or direction. You can use this flag multiple times to apply cumulative rotations.

### Usage

```bash
./bitmap apply --rotate=<angle|direction> [--rotate=<angle|direction>] input.bmp output.bmp
```

- `--rotate`: Rotates a bitmap image by a specified angle.
  - `--rotate=90`: Rotates the image 90 degrees clockwise.
  - `--rotate=180`: Rotates the image 180 degrees.
  - `--rotate=270`: Rotates the image 270 degrees clockwise.
  - `--rotate=left`: Rotates the image counterclockwise.
  - `--rotate=-90`: Rotates the image 90 degrees counterclockwise.
  - `--rotate=-180`: Rotates the image 180 degrees.
  - `--rotate=-270`: Rotates the image 270 degrees counterclockwise.
  - `--rotate=right`: Rotates the image clockwise.

### Crop Options
- `--crop`: Trims a bitmap image according to specified parameters.
  - Accepts values in the format: `OffsetX-OffsetY-Width-Height`
  - Example: `--crop=10-20-100-200` will crop starting from (10, 20) with width 100 and height 200.

## Edge Cases

- Only valid uncompressed 24-bit BMP files are accepted. Invalid files will result in a non-zero exit status and an error message.
- Input/output file names must come last.
- The program can overwrite existing files.
- Invalid flags will trigger an error message and a non-zero exit status.