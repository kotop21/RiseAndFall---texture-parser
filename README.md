<div align="center">

# Rise And Fall: Civilization at War Asset Utilities

A collection of utilities for working with **Rise And Fall: Civilization at War** assets.

[![Go Version](https://img.shields.io/badge/Go-1.25.11-00ADD8?logo=go&logoColor=white)](https://go.dev/)

</div>

The project uses a **Taskfile** to simplify cross-platform builds for **Windows**, **Linux**, and **macOS**.

## Build

Build every tool at once:

```bash
task build-all
```

---

## Texture Converter

Convert, unpack, and repack game texture formats.

### Supported Input Formats

- `.dds`
- `.tga`
- `.sst`

### Supported Output Formats

#### Unpacking

- `.png`

#### Packing

- Same format as the original texture (`.dds`, `.tga`, or `.sst`)

### Folder Structure

```text
1_original/    Original game textures
2_editing/     Extracted PNG files for editing
3_packed/      Repacked textures
```

Running the executable **without arguments** automatically creates this folder structure.

### Commands

| Command | Description |
|---------|-------------|
| `converter unpack` | Extract textures to PNG |
| `converter pack` | Pack edited PNG files back into the original format |

---

## Matte Builder

Generate game-ready texture maps from source images.

### Supported Input Formats

- `.png`
- `.jpg`
- `.jpeg`

### Supported Output Format

- `.png`

### Commands

| Command | Description |
|---------|-------------|
| `matte-builder diffuse.png ao.png normal.png specular.png` | Generate a diffuse texture and a normal/bump texture |
| `matte-builder --diffuse diffuse.png ao.png` | Bake the ambient occlusion map into the diffuse texture |
| `matte-builder --normal normal.png specular.png` | Pack the specular map into the normal map alpha channel |
