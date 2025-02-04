# Text Editors

A text editor is a program used to edit text files (as opposed to binary files, such as images or compiled executables). Text editors come in two main forms, terminal text editors and GUI text editors. Which you prefer is largely down to personal preference, however terminal text editors can be very useful, for example when remotely connecting to a system using SSH. 

This article will go over some examples of terminal and GUI text editors however it is not comprehensive, there are many other options avalible, these are just the editors I have experience with. Notable exclusions are [Vim](https://www.vim.org/), [Emacs](https://www.gnu.org/software/emacs/), [Sublime Text](https://www.sublimetext.com/), [Kate](https://kate-editor.org/en-gb/) and [Atom](https://github.com/atom/atom). 

A term you may see come up often is "IDE", which stands for Integrated Development Environment. An IDE will often include a text editor, as well as features to help with software development, such as an integrated terminal, integrated debugging support and [git version control](git.md) support. As I only have experience with one IDE (VSCode), I have decided to just include IDEs in this article. 

## Terminal Text Editors

### Nano 

GNU Nano is a very old text editor. First released in 1999, it is included by default on almost all Linux distrobutions, and is avalible for Windows and macOS as well. Nano comes with basic syntax highlighting and basic keyboard shortcuts, and it very easy to use for making small edits to files when a GUI text editor is not avalible. 

### Micro

Micro is a modern terminal editor writter in [Golang](golang.md). Designed to be easy to use, it does not use vim bindings, but rather a modifier based system, similar to the programs you are likely used to. It comes with syntax highlighting built in, and support for extensions and theming. 

#### Configuration

Micro's config file is located at `~/.config/micro/settings.json`, and keybindings are located at `~/.config/micro/bindings.json`. A few notable settings are listed below. 

Due to the way terminals work in capturing key input, it is difficult for some keyboard shortcut to be parsed. `Ctrl-I` for example, often used for "Indent Line", cannot be bound within micro. I bind `Ctrl-O` to this instead. 

In order to bind "Ctrl-Backspace" to the standard Ctrl-Backspace behaviour (deleting whole words at once), the following can be used: 

``` json
"OldBackspace": "DeleteWordLeft"
```

The following bindings will bind `Ctrl-W` to close the current tab, or quit if only one tab open, and `Ctrl-Q` to quit all tabs, prompting if unsaved work will be lost. 

``` json
"Ctrl-W": "Quit",
"Ctrl-Q": "QuitAll"
```

#### Editing Highlight Colours

The syntax highlighting colours can be edited by making a colourscheme file such as `~/.config/micro/colorschemes/draculacustom.micro`. [An example of some of the avalible mappings can be found in my personal dotfiles](https://github.com/randomcoder67/XFCE-Laptop-Config/blob/main/configure/home/.config/micro/colorschemes/draculacustom.micro).

The mapping work as follows:

``` yaml
- constant.bool: "\\b(true|false)\\b"
```

In the above snippet, "true" or "false" are matched. The colour they appear as will depends on the theme file, mentioned earlier. The key name above is `constant.bool`. If no colour for `constant.bool` is found, the colour for `constant` will be used instead, so:

``` micro
color-link constant "#ff0000"
```

Would set the colour of "true" or "false" to `#ff0000`, whereas:

``` micro
color-link constant "#ff0000"
color-link constant.bool "#00ffff"
```

Would set the colour to `#00ffff`, as the `constant.bool` colour is the better match for the `constant.bool` key. 

#### Editing Syntax Highlighting 

The above of course only changes the colours of the existing syntax highlighting. If you wish to change what is/isn't highlighted, custom syntax files can be places in `~/.config/micro/syntax/` in files such as `go.yaml`. The default files can be found [in the Micro GitHub repo](https://github.com/zyedidia/micro/tree/master/runtime/syntax). These files use regular expressions to match the text to highlight. The files is read top to bottom, with lower rules overwriting higher ones. This snippet for example: 

``` yaml
- function: "[a-zA-Z0-9_]*[ ]?\\("
- default: "\\("
```

Will highlight the name of a function. The first line selects the function name and also the open bracket. The second line is then used to "undo" the selected of the bracket, by setting all "(" characters to the default colour.

## GUI Text Editors

### VSCode
 
VSCode is an open source code editor/IDE designed by Microsoft (although is is avalible stripped of all Microsoft branding and features as [VSCodium](https://vscodium.com/)). VSCode comes with a number of useful IDE features, such as an inbuilt terminal, error highlighting, inbuilt git support and the ability to edit documents over SSH. It is however not very lightweight, being programmed in Electron. This does make it very extensible however. 

#### Useful Keyboard Shortcuts

`Ctrl+Shift+V` when editing a markdown file will open a new tab showing a preview of how the markdown looks when rendered. 

### Mousepad

Mousepad is the default text editor included in the [Xfce Desktop Environment](xfce.md). It is lightweight and doesn't have any IDE features, but it does come with syntax highlighting and a variety of useful keyboard shortcuts, which can be changed. 

#### Changing Keyboard Shortcuts

Mousepad comes with an inbuilt extension to view and edit keyboard shortcuts. It can be accessed through `Edit -> Preferences -> Plugins -> Shortcuts Editor`. If you wish to automatically load keyboard shortcuts, the file `~/.config/Mousepad/accels.scm` can be used. Note this file is saved upon closing Mousepad, so Mousepad must be closed when editing it, otherwise it will be overwritten with the old file upon closing Mousepad

#### Spell Checking

In order to use spell checking, you need to install a compatible spell checking library. On Arch this can be done using `pacman -S gspell`. The spellchecking is often annoying when editing things that are not just "language text", for example code. A keyboard shortcut can be set to toggle spell checking on and off. I use `Ctrl+Shift+K` for this. This can be set in the aforemention shortcuts editor, or in `accels.scm` as:

```
(gtk_accel_path "<Actions>/app.mousepad-plugin-gspell" "<Primary><Shift>k")
```

#### Editing Colour Scheme

Mousepad uses GTKSourceView for syntax highlighting, so the themes files are located at `/usr/share/gtksourceview-4/styles/`. User created theme files can be put in `~/.local/share/gtksourceview-4/styles/`. 

The theme files are in XML format, each mapping is in a `<style>` tag. The names are all similar to `def:heading`. The `def` part refers to the syntax file where the syntax is defined. So if there is syntax pattern in the `c.lang` file with the style id `function`, you would refer to it as `c:function`. This is covered more below in the Editing Syntax Highlighting section. 

#### Editing Syntax Highlighting

The default syntax highlighting files are located at `/usr/share/gtksourceview-4/language-specs/`. User created files can be put in `~/.local/share/gtksourceview-4/language-specs/`. Files in the user directory will always be prioritised over the defaults. 

The syntax highlighting is in the form of "contexts", using regular expressions to select which areas to highlight. This is an example from my custom `markdown.lang` file. 

``` xml
<context id="3-backticks-code-span-xml" class="no-spell-check">
    <start>^```[ ]?(xml)[ ]?$</start>
    <end>^```$</end>
    <include>
        <context sub-pattern="0" where="start" style-ref="code-block"/>
        <context sub-pattern="0" where="end" style-ref="code-block"/>
        <context ref="xml:xml"/>
    </include>
</context>
```

The above snippet is used to highlight markdown xml code blocks, giving the code inside the syntax highlighting of the xml language. The text matched by the above snippet would be something like (ignore backslash):

``` markdown
``` xml
<text>This is some text</text>
<image src=icon.png alt="This is an image">
\``` 
```

The `<include>` statement means that the contained rules will only be applied to text that is within the `<start>` and `<end>` regex. 

This line: `<context sub-pattern="0" where="start" style-ref="code-block"/>` means the first line "``` xml" will be highlighted in the "code-block" style. 

This line: `<context sub-pattern="0" where="end" style-ref="code-block"/>` means the last line "```" will be highlighted in the "code-block" style.

This line: `<context ref="xml:xml"/>` means that everything inbetween the `<start>` and `<end>` regex will be highlighted in the "xml:xml" style. "xml:xml" means "use every rule in the `xml.lang` file"

Generally, "sub-pattern=0" refers to the whole regex string, with "sub-pattern=1" for example referring to the first regex capture group in the regex string etc. 

### Leafpad 

Leafpad is a very basic text editor, not including syntax highlighting or complex keyboard shortcuts. Not recommended for code editing, however it can be useful for focused writing, if all you need is a blank document and the ability to input text.

### Lite XL

### Gedit

### Geany
