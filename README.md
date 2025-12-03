# Captcha Generator

## Usage

### Flags

| name     | flag         | type   | default  | desc.                         |
| -------- | ------------ | ------ | -------- | ----------------------------- |
| val      | `--val`      | string | 0        | value / text captcha          |
| w        | `--w`        | int    | 0        | render width                  |
| h        | `--h`        | int    | 0        | render height                 |
| out      | `--out`      | string | `base64` | output type `base64` or `png` |
| fontpath | `--fontpath` | string | ""       | path to font `ttf`            |

### Output to Base64

#### Base64 flag

| name | flag    | type   | value    |
| ---- | ------- | ------ | -------- |
| out  | `--out` | string | `base64` |

see more [flags](#flags)

#### Example

```shell
captcha_generator.exe --val=foobar --out=base64
```

will output like this

<pre style="white-space: pre-wrap; word-wrap: break-word;">
iVBORw0KGgoAAAANSUhEUgAAAJYAAAA8CAIAAAAL5NQ9AAAdlUlEQVR4nOR9WWxk15neWe5a99ZeRRZ3skn2ql7UUmu1WrE948gay5sMLxMZCWYQ+CVAkAV5CJAA8zDZHgZBECDIQ4I8eOzYUbzFtsay3RqP7JZsqd0re+W+FKtYrPXu954lqCqyuHSxSLYot2x9Igjx1q1T557v/P///f85p1p47c1Pg4PD6mLNd+l7bwdj2D0SO4ge/eEDHWxzxNuRv2z0sWz0LEVi/...
</pre>

you can preview using [Base64 viewer](https://base64-viewer.onrender.com/)

![Base64 viewer](./docs/images/base64_viewer.jpg)

### Output to PNG

#### PNG flag

| name | flag    | type   | value |
| ---- | ------- | ------ | ----- |
| out  | `--out` | string | `png` |

see more [flags](#flags)

#### Example

```shell
captcha_generator.exe --val=foobar --out=png
```

will output like this

```txt
successfully generate : captcha_foobar.png
```

or preview like this

![math_captcha_2_3.png](./docs/images/captcha_foobar.png)

## Task

Using [go-task/task](https://github.com/go-task/task)

### Build

Build to executable `captcha_generator.exe` at folder `build`

```shell
task b
```

or

```shell
task build
```

check [Taskfile.yml](Taskfile.yml)

### Run output base64

Run executable `captcha_generator.exe` in folder `build`

```shell
task base64
```

or

```shell
task test_base64
```

check [Taskfile.yml](Taskfile.yml)

### Run output PNG

Run executable `captcha_generator.exe` in folder `build`

```shell
task png
```

or

```shell
task test_png
```

check [Taskfile.yml](Taskfile.yml)

## Build from source

```
go build -o ./build/captcha_generator.exe ./cmd/math/
```

## Todo

- [ ] Make magic numbers more customizable
- [ ] Add more outputs (JPG, ...)
