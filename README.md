# Neo's Starsiege Launcher

Launches a specified version Starsiege following the proposed `starsiege://` URI

### Quick Install instructions

1. Download `launcher.exe` and place it where your `Starsiege.exe` is located
2. Download [example.reg](example.reg)
3. Open `example.reg` in your favorite text editor
4. Replace all 3 instances of `\\path\\to...` to the folder from step 1, making sure your path separators are two backslashes (`\\`).
5. Save, Exit, and Double-Click on your `.reg` file
6. Click yes on the prompt or prompts
7. Try it out: Github doesn't like to link to non-standard URIs, so pasting the following URI `starsiege://96.126.117.157:29007` into your browser's address bar will work just as well. 
8. Check out the other servers you can now instantly join at [https://starsiegeplayers.com/multiplayer](https://starsiegeplayers.com/multiplayer)

If Starsiege doesn't launch, or you need any additional help, visit us at the [Starsiege Players Discord](https://discord.gg/KA4N6J8)

### Overview

Windows provides a way to register arbitrary protocols with the proper registry configuration.

We can use this mechanism to launch Starsiege, similarly to how [Steam launches apps with the `steam://` URIs](https://developer.valvesoftware.com/wiki/Steam_browser_protocol).
For example pasting the URI `steam://url/StoreAppPage/17080` in your address bar should open a dialog box withing your browser asking to launch steam.
This will take you to the `Tribes: Ascend` Steam store page.

Other URIs such as `rungameid` will go ahead and launch the game in Steam. Pasting the following in your address bar `steam://rungameid/17080` will prompt Steam to launch `Tribes: Ascend`

### Example Registry information (`example.reg`)

Every install of Starsiege is unique, therefore this file provided for use as a template.
The simplest way of using the launcher is to place it alongside `starsiege.exe`

In this case, change all 3 instances of `\\path\\to\\...` to the folder your `starsiege.exe` file is in.

```reg
Windows Registry Editor Version 5.00

[HKEY_CLASSES_ROOT\starsiege]
@="URL:Starsiege Protocol"
"URL Protocol"=""
"DefaultIcon"="C:\\path\\to\\starsiege.exe,1"
"Executable"="C:\\path\\to\\starsiege.exe"

[HKEY_CLASSES_ROOT\starsiege\shell]

[HKEY_CLASSES_ROOT\starsiege\shell\open]

[HKEY_CLASSES_ROOT\starsiege\shell\open\command]
@="\"C:\\path\\to\\launcher.exe\" %1"
```

### Proposed URI Schema Definition

| status               | uri scheme                                  | function                                                                                     |
| -------------------- | ------------------------------------------- | -------------------------------------------------------------------------------------------- |
| :heavy_check_mark:   | `starsiege://{host}:{port}/join`            | Opens`starsiege.exe` using the `+connect IP:{host}:{port}` quick connect function            |
| ------               | `starsiege://{host}:{port}/addrbook[?name]` | Adds the starsiege server specified with`{host}:{port}` and optional `[name]` to addrBook.cs |
| ------               | `starsiege://{host}:{port}/master`          | Adds the master server specified with`{host}:{port}` to addrBook.cs                          |
| ------               | `starsiege://multiplayer`                   | Opens`starsiege.exe` directly to the Multiplayer Server Browser                              |
