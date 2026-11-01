<p align="center">
  <img src="build/appicon.png" alt="croc-desktop 图标" width="112" height="112" />
</p>

<h1 align="center">croc-desktop</h1>

<p align="center">
  <b>在设备之间发送文件和文字：简单、私密、不经过云盘。</b>
</p>

<p align="center">
  为
  <a href="https://github.com/schollz/croc">croc</a>
  做的桌面应用：端到端加密，点对点传输。
  <br />
  <b>桌面：</b> Linux · macOS · Windows &nbsp;·&nbsp;
  <b>移动：</b> iOS · Android
  <br />
  <sub>基于 <a href="https://v3.wails.io">Wails v3</a>（旧文档中的 v2 已过时）</sub>
</p>

<p align="center">
  <a href="https://github.com/SihanTeng/croc-desktop/releases/latest"><img src="https://img.shields.io/github/v/release/SihanTeng/croc-desktop?style=for-the-badge&label=Download&color=1F6FEB" alt="下载" /></a>
  &nbsp;
  <a href="https://github.com/SihanTeng/croc-desktop/actions/workflows/ci.yml"><img src="https://img.shields.io/github/actions/workflow/status/SihanTeng/croc-desktop/ci.yml?branch=main&style=for-the-badge&label=CI" alt="CI" /></a>
  &nbsp;
  <img src="https://img.shields.io/badge/Desktop-Linux%20%7C%20macOS%20%7C%20Windows-1F6FEB?style=for-the-badge" alt="桌面平台" />
  &nbsp;
  <img src="https://img.shields.io/badge/Mobile-iOS%20%7C%20Android-0E8A16?style=for-the-badge" alt="移动平台" />
  &nbsp;
  <a href="https://v3.wails.io"><img src="https://img.shields.io/badge/Wails-v3-DF3A2C?style=for-the-badge" alt="Wails v3" /></a>
</p>

<p align="center">
  <a href="#install-desktop">安装桌面端</a> ·
  <a href="#mobile">移动端</a> ·
  <a href="#how-to-use">怎么用</a> ·
  <a href="#features">功能</a> ·
  <a href="#faq">常见问题</a> ·
  <a href="#contributing">参与贡献</a> ·
  <a href="#development">开发</a>
</p>

---

<p align="center">
  <img src="docs/images/send.png" alt="发送页：取件码与二维码" width="48%" />
  &nbsp;
  <img src="docs/images/receive.png" alt="接收页：文件预览" width="48%" />
</p>

<p align="center">
  <img src="docs/images/history.png" alt="历史传输记录" width="48%" />
  &nbsp;
  <img src="docs/images/logs.png" alt="实时日志" width="48%" />
</p>

## 为什么选择 croc-desktop？ {#why}

把文件送到另一台电脑，不必先上传到别人的服务器，也不必先学会命令行。

**croc-desktop** 是 [croc](https://github.com/schollz/croc) 的完整图形界面：

| | |
| --- | --- |
| **默认私密** | 端到端加密，点对点传输。文件不会存到云端。 |
| **一个取件码（或二维码）** | 把短码发给对方，或让对方扫码，即可开始。 |
| **每次都由你决定** | 可以接受或拒绝。覆盖和续传也由你选择。 |
| **能和命令行配合** | 取件码与命令行 `croc` 相同，桌面端和命令行可以混用。 |
| **不用开终端** | 点选、拖放、粘贴即可。 |

它在进程内运行 croc，因此有真正的对话框、实时进度、历史记录和中继。薄封装只会在窗口里启动 `croc` 命令。

> **技术栈：** 本项目使用 **[Wails v3](https://v3.wails.io)**（`github.com/wailsapp/wails/v3`）。旧文档或旧分支里的 Wails v2 已过时。

## 安装（桌面端） {#install-desktop}

### 一行安装

**Linux / macOS：**

```sh
curl -fsSL https://raw.githubusercontent.com/SihanTeng/croc-desktop/main/install.sh | bash
```

**Windows（PowerShell）：**

```powershell
irm https://raw.githubusercontent.com/SihanTeng/croc-desktop/main/install.ps1 | iex
```

| 平台 | 脚本安装的内容 | 说明 |
| --- | --- | --- |
| **Linux**（amd64） | AppImage → `~/.local/bin/croc-desktop` | 会添加桌面菜单；请确认 `~/.local/bin` 在 `PATH` 中 |
| **macOS 13+**（Apple Silicon） | DMG → `/Applications` | 临时签名。首次打开可能需要在**隐私与安全性**里选择**仍要打开** |
| **Windows**（x64） | 通过 `msiexec` 安装 MSI | 可能会请求管理员权限 |

可选：在一行命令之前指定版本。Unix 用 `VERSION=v0.2.0`，PowerShell 用 `$env:VERSION = "v0.2.0"`。

### 手动下载

**[前往 Releases →](https://github.com/SihanTeng/croc-desktop/releases/latest)**

| 平台 | 文件 | 提示 |
| --- | --- | --- |
| **Linux** | `croc-desktop_*_linux-amd64.AppImage` | 先 `chmod +x`，再双击或直接运行 |
| **macOS 13+**（Apple Silicon） | `croc-desktop_*_darwin-arm64.dmg` | 打开 DMG，把应用拖到「应用程序」 |
| **Windows** | `croc-desktop_*_windows-amd64.msi` | 运行安装包 |

### 包管理器

**Homebrew（macOS，Apple Silicon）：**

```sh
brew install --cask sihanteng/tap/croc-desktop
```

下列渠道仍在计划中，目前还没有发布包：

| 渠道 | 平台 | 状态 |
| --- | --- | --- |
| **AUR** | Arch Linux | 计划中，尚未发布 |
| **Fedora**（Copr / 官方仓库） | Fedora | 计划中，尚未发布 |

在此之前，请用上面的一行安装脚本，或从 Releases 手动下载。

> 传输双方各自安装本应用或 [croc 命令行](https://github.com/schollz/croc) 即可。

## 移动端（iOS · Android） {#mobile}

同一套应用、同一套界面（手机宽度下使用底部标签）。由 **Wails v3** 的原生移动目标驱动。

| 客户端 | 状态 | 位置 | 如何运行 |
| --- | --- | --- | --- |
| **iOS** | 实验性，需从源码构建 | 脚手架在 [`build/ios/`](build/ios/) | 模拟器 / 真机（macOS 上需要完整 **Xcode**） |
| **Android** | 实验性，需从源码构建 | 脚手架在 [`build/android/`](build/android/) | 模拟器 / 真机（需要 **Android SDK + NDK**） |

**App Store 和 Play 商店还没有安装包。** 移动端面向自行从本仓库构建的人（或你自己配的 CI）。可下载的正式途径仍是 [Releases](https://github.com/SihanTeng/croc-desktop/releases) 上的桌面版。

### 构建并运行移动端

前提与[开发](#development)相同（Go、Node、Wails v3 CLI），并安装对应平台的 SDK。

```sh
# 安装 Wails v3 CLI（版本与 go.mod 一致）
go install -tags gtk3 github.com/wailsapp/wails/v3/cmd/wails3@v3.0.0-alpha2.119

# 前端资源（会嵌入应用）
npm --prefix frontend ci && npm --prefix frontend run build

# ── iOS（仅 macOS + Xcode）─────────────────────────────────
wails3 task ios:run                          # 模拟器
# 真机：
# wails3 task ios:package IOS_PLATFORM=device \
#   CODESIGN_IDENTITY="Apple Development: You (TEAMID)"

# ── Android ──────────────────────────────────────────────────
wails3 task android:run                      # 模拟器或 adb 设备
```

| | 值 |
| --- | --- |
| iOS bundle ID | `io.github.sihanteng.croc-desktop` |
| Android applicationId | `io.github.sihanteng.crocdesktop`（不含连字符） |
| 配置 | [`build/config.yml`](build/config.yml) 中的 `ios:` 段 |
| 平台钩子 | [`app_options_ios.go`](app_options_ios.go)、[`app_options_android.go`](app_options_android.go) |

**界面预览：** 把桌面窗口拉到大约 360 像素宽，即可看到手机上的底部标签，不必真机。

## 怎么用 {#how-to-use}

### 发送

1. 打开**发送**。
2. 拖入文件或文件夹，或浏览选择；也可以切换到**文字**发送一段文本。
3. 点击**发送**。界面会给出短**取件码**和**二维码**。
4. 把取件码或二维码图片发给接收方。
5. 等待对方接受。进度会实时显示。可随时用按钮或 `Esc` 取消。

### 接收

1. 打开**接收**。
2. 粘贴取件码（也可以是整段 `croc …` 命令或分享链接），或者载入一张**二维码截图**。
3. 如需要，选择下载目录，然后点**接收**。
4. 出现提示时选择**接受**或拒绝。
5. 完成后，可在应用内预览图片、视频、音频或文字。

### 日常功能

- **历史**：可搜索以往的发送和接收。
- **收藏**：保存经常使用的取件码。
- **再发一次**：不必重新挑选，再次发送同一批文件或文字。
- **设置**：主题（浅色 / 深色 / 跟随系统）、语言、自定义中继、代理等。
- **中继**：可在应用内临时开一台中继，用于受限网络。
- **自建中继**：发送密钥只留在发送端。见[自建中继操作说明](docs/自建中继.md)。

## croc 引擎兼容性 {#engine}

本构建使用 **croc v11.6.0**。双方都需要 croc v11；仍在使用 v10 的旧桌面端必须升级。取件码现为三个单词，文字消息上限为 1 MiB。迁移行为和校验见[迁移说明](docs/croc-v11-migration.md)，其中也包括保留的中继与代理路由。

## 功能 {#features}

### 传输

- 文件和文件夹（拖放或选择）
- 文字片段
- 二维码，以及复制取件码 / 复制命令行
- 粘贴常见格式的取件码（`code`、`croc code`、`CROC_SECRET=…`、分享链接）
- 从二维码截图识别取件码
- 接受 / 拒绝，以及覆盖 / 续传提示
- 实时进度、停滞提示，随时取消
- 中断后的接收可以续传
- 接收完成后在应用内预览
- 完成或失败时的桌面通知

### 应用

- 重启后仍在的传输**历史**
- 可按级别筛选的实时**日志**
- 可选的应用内**中继**
- 自定义中继地址和密码、加密曲线、哈希算法
- 仅局域网 / 关闭局域网发现、压缩、代理（SOCKS5 / HTTP）
- 发送前压缩文件夹、排除规则、上传限速
- 浅色 / 深色 / 跟随系统主题
- 语言：English、简体中文、繁體中文、Español、Français、Deutsch、日本語
- 自适应布局：桌面用侧栏，手机或窄窗口用**底部标签**

## 常见问题 {#faq}

<details>
<summary><b>文件会上传到服务器吗？</b></summary>

<br>

不会。传输是端到端加密的点对点传输。公共**中继**只帮助双方找到彼此（与 croc 命令行相同），文件内容不会留在中继上。你可以在应用里自己开中继，或在设置里填写私人中继。

</details>

<details>
<summary><b>双方都必须安装这个应用吗？</b></summary>

<br>

不必。一方用 **croc-desktop**，另一方用官方 [`croc` 命令行](https://github.com/schollz/croc) 即可，反过来也行。取件码互相兼容。

</details>

<details>
<summary><b>可以离线，或只在局域网里用吗？</b></summary>

<br>

可以。两台设备在同一网络时，在设置里打开**仅局域网**（或使用局域网发现）。也可以在**中继**页自己运行一台中继。

</details>

<details>
<summary><b>设置和历史存在哪里？</b></summary>

<br>

与 croc 命令行相同的配置目录（可用环境变量 `CROC_CONFIG_DIR` 覆盖）：

- `croc-desktop.json`：偏好设置
- `croc-desktop-history.json`：传输历史
- `croc-desktop.log`：持久日志（JSON 行，大约 2 MiB × 3 个文件轮转）

旧版使用 `croc-gui.json` 的安装，会在首次运行时自动升级。

</details>

<details>
<summary><b>macOS 提示应用来自未识别的开发者</b></summary>

<br>

当前发布包是**临时签名**，尚未公证。首次被拦截后，打开**系统设置 → 隐私与安全性**并允许打开，或右键点击后选择**打开**。

</details>

<details>
<summary><b>这还是 Wails v2 吗？</b></summary>

<br>

本分支使用 **Wails v3**（`github.com/wailsapp/wails/v3`，见 `go.mod`）。文档：[v3.wails.io](https://v3.wails.io)。若在 GitHub 上仍看到 “Wails v2”，说明看的是合并前的旧分支。

</details>

<details>
<summary><b>手机应用在哪里下载？</b></summary>

<br>

见上面的[移动端（iOS · Android）](#mobile)。它们用 `wails3 task ios:run` / `android:run` 从本仓库构建，还没有上架商店。

</details>

<details>
<summary><b>使用 Clash 等代理时连不上中继？请将中继设为直连（DIRECT）</b></summary>

<br>

若本机开启了 **Clash / Clash Meta（mihomo）/ Clash Verge** 等代理（尤其是
**fake-IP** 模式），公共中继域名可能被解析成 `198.18.x.x` 等虚假地址，导致
发送/接收失败（超时、EOF、「无法连接中继」等）。

**请在代理工具中把 croc 公共中继设为直连（DIRECT）**，并建议加入 fake-IP
过滤。配置后**重载配置**，并**重启 croc-desktop**。应用内设置请填写
**域名**（如 `croc.schollz.com:9009`），不要填写 `198.18.…` 地址。
自建中继若填写的是公网 IP，请为该 IP 添加 `IP-CIDR,...,DIRECT` 规则，详见[自建中继操作说明](docs/自建中继.md)。

**规则示例**（写在 `rules:` **靠前**位置，保证优先生效）：

```yaml
rules:
  - DOMAIN,croc.schollz.com,DIRECT
  - DOMAIN,croc6.schollz.com,DIRECT
  # 可选：同域名后缀一并直连
  - DOMAIN-SUFFIX,schollz.com,DIRECT
  # …其余规则…
```

**DNS fake-IP 过滤（推荐，Clash Meta / mihomo）：**

```yaml
dns:
  enable: true
  enhanced-mode: fake-ip
  fake-ip-filter:
    - croc.schollz.com
    - croc6.schollz.com
    - +.schollz.com   # 可选
```

图形界面用户可在「规则 / Rules」中手动添加上述 **DOMAIN → DIRECT** 项，
保存后重载配置。

> 直连只在你的真实网络能够访问该中继时有效。若运营商屏蔽了公共中继，
> 请改用可访问的中继，或在设置中填写自建中继地址。

</details>

## 参与贡献 {#contributing}

欢迎报告缺陷、翻译、改进文档和提交拉取请求。

- 彼此尊重；请遵守[行为准则](CODE_OF_CONDUCT.md)。
- 优先提交小而集中的拉取请求。
- 界面文案在 `frontend/src/i18n/locales/` 的 JSON 中修改（`en.json` 是基准）。复制为新语言文件（例如 `pt-BR.json`）后翻译值；保留键名和 `{placeholders}`。

## 开发 {#development}

<details>
<summary><b>从源码构建（开发者）</b></summary>

<br>

**技术栈：** Go（见 `go.mod`）、Node.js、**[Wails v3](https://v3.wails.io)**。模块 `github.com/wailsapp/wails/v3`，版本 **v3.0.0-alpha2.119**（不是 v2 的命令行工具）。

```sh
# Wails v3 CLI（没有 gtk4 的 Linux 使用 -tags gtk3）
go install -tags gtk3 github.com/wailsapp/wails/v3/cmd/wails3@v3.0.0-alpha2.119

# Linux：webkit2gtk 4.1 + GTK3（项目默认）
#   Fedora: webkit2gtk4.1-devel
#   Debian/Ubuntu: libgtk-3-dev libwebkit2gtk-4.1-dev

# 带热重载的桌面开发
WEBKIT_DISABLE_DMABUF_RENDERER=1 wails3 dev -config ./build/config.yml -port 34115

# 生产桌面程序 → bin/croc-desktop
wails3 build -tags gtk3   # macOS / Windows 可省略 -tags gtk3

# 移动端见上文「移动端（iOS · Android）」
wails3 task ios:run
wails3 task android:run

# 测试
npm --prefix frontend ci && npm --prefix frontend run build
go test -tags gtk3 .
npm --prefix frontend test
./e2e/run.sh
```

更多说明见 `e2e/README.md`、`design.md`，以及 `build/` 下的 Taskfile。

**发布：** 推送版本标签后，CI 会构建三个桌面平台、发布 GitHub Release，并更新
[SihanTeng/homebrew-tap](https://github.com/SihanTeng/homebrew-tap) 中的 Homebrew cask
（未设置 `HOMEBREW_TAP_GITHUB_TOKEN` 时会静默跳过）：

```sh
git tag v0.2.1 && git push --tags
```

版本以标签为准：构建时写入二进制、macOS 的 plist、Windows 资源和 MSI，不必再手工改版本文件。

**实现简述：** 应用依赖带钩子层的 croc，进度和确认框是界面事件，由进程内引擎产生。具体锁定见 `go.mod` 的 `replace`。

</details>

## 致谢 {#acknowledgments}

- [@schollz](https://github.com/schollz) 的 [croc](https://github.com/schollz/croc)：协议与传输引擎
- [Wails v3](https://v3.wails.io)：用于桌面和移动端的 Go + Web 界面

## 许可证 {#license}

[MIT](LICENSE) © SihanTeng。croc-desktop 是**非官方**社区项目，与 croc 作者没有隶属或背书关系。`croc` 本身由其各自作者以 MIT 许可发布。

---

<p align="center">
  <sub>给只想把文件送出去的人。</sub>
</p>
