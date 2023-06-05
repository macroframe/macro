# macroframe

## Requirements

<details>
<summary><h3>macOS</h3></summary>

```bash
xcode-select --install
```

If you want to use screen capture functions, you might have to give your programm the right permissions:  
`System Preferences > Security and Privacy > Privacy > Accessibility and Screen Recording`

</details>

<details>
<summary><h3>Debian / Ubuntu</h3></summary>

```bash
sudo apt install gcc libc6-dev libx11-dev xorg-dev libxtst-dev xsel xclip libpng++-dev xcb libxcb-xkb-dev x11-xkb-utils libx11-xcb-dev libxkbcommon-x11-dev libxkbcommon-dev
```

</details>

<details>
<summary><h3>Windows</h3></summary>

CGO support is needed.

1. Install [TDM-GCC](https://jmeubank.github.io/tdm-gcc/download/).
   1. Download the `tdm-gcc-webdl.exe` file.
   2. Run the installer.
   3. Select 64 and 32 bit version.
2. Download [zlib](http://sourceforge.net/projects/mingw-w64/files/External%20binary%20packages%20(Win64%20hosted)/Binaries%20(64-bit))
   1. You only need the `zlib-1.2.5-bin-x64.zip` file.
3. Extract the zlib archive to your TDM-GCC installation directory.
   1. `zlib/bin` → `TDM-GCC-64/bin`
   2. `zlib/include` → `TDM-GCC-64/include`
   3. `zlib/lib` → `TDM-GCC-64/lib`

</details>

<details>
<summary><h3>Fedora</h3></summary>

```bash
sudo dnf install libXtst-devel xsel xclip libpng-devel libxkbcommon-devel libxkbcommon-x11-devel xorg-x11-xkb-utils-devel
```

</details>
