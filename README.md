# ColorUtils

[![Go Reference](https://pkg.go.dev/badge/github.com/shadiestgoat/colorutils.svg)](https://pkg.go.dev/github.com/shadiestgoat/colorutils)

This package is intended to provide a few utility functions when it comes to colors in go.

## Installation

To install, run the following command:

```sh
go get -u github.com/shadiestgoat/colorutils
```

## Color Space Conversion

This package gives free conversion between RGB, HSV and HSL:

| from | to | function |
|:-:|:-:|:-:|
| RGB | HSL | RGBToHSL |
| RGB | HSV | RGBToHSV |
| HSL | RGB | HSLToRGB |
| HSL | HSV | HSLToHSV |
| HSV | RGB | HSVToRGB |
| HSV | HSL | HSVToHSL |

## Luminosity

This package provides a few utils for luminosity calculations. This is useful for creating colors that truly contrast each other, and meet WCAG standards. For this, the package provides `RelativeLuminosity` to calculate relative luminosity (Y) of a color, and `ContrastRatio` to calculate the ratio between 2 colors.

The package also provides `NewContrastColor`, for when you want to generate a WCAG compliant contrasting color, for both a lighter and a dark background.

## Hexadecimal

The package provides a util method for creating hex color string - `Hexadecimal` and `HexadecimalOpacity`. 

## Color Schemes

This package can also create color schemes, given a base hue:
- Monochromatic
- Analogous
- Complementary
- Triadic
- Compound

These can be accessed with `ColorScheme{Name}`. 
There are also util functions for `ColorScheme{Name}RGB`, which are wrappers than change the values in RGB. These are usually not recommended, since they are more restricting in terms of hue/saturation/lightness
