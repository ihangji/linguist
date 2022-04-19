// DO NOT EDIT ** This file was generated with the bake tool ** DO NOT EDIT //

package linguist

var files = map[string]string{
	"data/languages.yml": `# Defines all Languages known to GitHub.
#
# fs_name               - Optional field. Only necessary as a replacement for the sample directory name if the
#                         language name is not a valid filename under the Windows filesystem (e.g., if it
#                         contains an asterisk).
# type                  - Either data, programming, markup, prose, or nil
# aliases               - An Array of additional aliases (implicitly
#                         includes name.downcase)
# ace_mode              - A String name of the Ace Mode used for highlighting whenever
#                         a file is edited. This must match one of the filenames in http://git.io/3XO_Cg.
#                         Use "text" if a mode does not exist.
# codemirror_mode       - A String name of the CodeMirror Mode used for highlighting whenever a file is edited.
#                         This must match a mode from https://git.io/vi9Fx
# codemirror_mime_type  - A String name of the file mime type used for highlighting whenever a file is edited.
#                         This should match the `+"`"+`mime`+"`"+` associated with the mode from https://git.io/f4SoQ
# wrap                  - Boolean wrap to enable line wrapping (default: false)
# extensions            - An Array of associated extensions (the first one is
#                         considered the primary extension, the others should be
#                         listed alphabetically)
# filenames             - An Array of filenames commonly associated with the language
# interpreters          - An Array of associated interpreters
# language_id           - Integer used as a language-name-independent indexed field so that we can rename
#                         languages in Linguist without reindexing all the code on GitHub. Must not be
#                         changed for existing languages without the explicit permission of GitHub staff.
# color                 - CSS hex color to represent the language. Only used if type is "programming" or "markup".
# tm_scope              - The TextMate scope that represents this programming
#                         language. This should match one of the scopes listed in
#                         the grammars.yml file. Use "none" if there is no grammar
#                         for this language.
# group                 - Name of the parent language. Languages in a group are counted
#                         in the statistics as the parent language.
#
# Any additions or modifications (even trivial) should have corresponding
# test changes in `+"`"+`test/test_blob.rb`+"`"+`.
#
# Please keep this list alphabetized. Capitalization comes before lowercase.
---
1C Enterprise:
  type: programming
  color: "#814CCC"
  extensions:
  - ".bsl"
  - ".os"
  tm_scope: source.bsl
  ace_mode: text
  language_id: 0
2-Dimensional Array:
  type: data
  color: "#38761D"
  extensions:
  - ".2da"
  tm_scope: source.2da
  ace_mode: text
  language_id: 387204628
4D:
  type: programming
  color: "#004289"
  extensions:
  - ".4dm"
  tm_scope: source.4dm
  ace_mode: text
  language_id: 577529595
ABAP:
  type: programming
  color: "#E8274B"
  extensions:
  - ".abap"
  tm_scope: source.abap
  ace_mode: abap
  language_id: 1
ABAP CDS:
  type: programming
  color: "#555e25"
  extensions:
  - ".asddls"
  tm_scope: source.abapcds
  language_id: 452681853
  ace_mode: text
ABNF:
  type: data
  ace_mode: text
  extensions:
  - ".abnf"
  tm_scope: source.abnf
  language_id: 429
AGS Script:
  type: programming
  color: "#B9D9FF"
  aliases:
  - ags
  extensions:
  - ".asc"
  - ".ash"
  tm_scope: source.c++
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-c++src
  language_id: 2
AIDL:
  type: programming
  color: "#34EB6B"
  tm_scope: source.aidl
  extensions:
  - ".aidl"
  ace_mode: text
  interpreters:
  - aidl
  language_id: 451700185
AL:
  type: programming
  color: "#3AA2B5"
  extensions:
  - ".al"
  tm_scope: source.al
  ace_mode: text
  language_id: 658971832
AMPL:
  type: programming
  color: "#E6EFBB"
  extensions:
  - ".ampl"
  - ".mod"
  tm_scope: source.ampl
  ace_mode: text
  language_id: 3
ANTLR:
  type: programming
  color: "#9DC3FF"
  extensions:
  - ".g4"
  tm_scope: source.antlr
  ace_mode: text
  language_id: 4
API Blueprint:
  type: markup
  color: "#2ACCA8"
  ace_mode: markdown
  extensions:
  - ".apib"
  tm_scope: text.html.markdown.source.gfm.apib
  language_id: 5
APL:
  type: programming
  color: "#5A8164"
  extensions:
  - ".apl"
  - ".dyalog"
  interpreters:
  - apl
  - aplx
  - dyalog
  tm_scope: source.apl
  ace_mode: text
  codemirror_mode: apl
  codemirror_mime_type: text/apl
  language_id: 6
ASL:
  type: programming
  ace_mode: text
  extensions:
  - ".asl"
  - ".dsl"
  tm_scope: source.asl
  language_id: 124996147
ASN.1:
  type: data
  extensions:
  - ".asn"
  - ".asn1"
  tm_scope: source.asn
  ace_mode: text
  codemirror_mode: asn.1
  codemirror_mime_type: text/x-ttcn-asn
  language_id: 7
ASP.NET:
  type: programming
  tm_scope: text.html.asp
  color: "#9400ff"
  aliases:
  - aspx
  - aspx-vb
  extensions:
  - ".asax"
  - ".ascx"
  - ".ashx"
  - ".asmx"
  - ".aspx"
  - ".axd"
  ace_mode: text
  codemirror_mode: htmlembedded
  codemirror_mime_type: application/x-aspx
  language_id: 564186416
ATS:
  type: programming
  color: "#1ac620"
  aliases:
  - ats2
  extensions:
  - ".dats"
  - ".hats"
  - ".sats"
  tm_scope: source.ats
  ace_mode: ocaml
  language_id: 9
ActionScript:
  type: programming
  tm_scope: source.actionscript.3
  color: "#882B0F"
  aliases:
  - actionscript 3
  - actionscript3
  - as3
  extensions:
  - ".as"
  ace_mode: actionscript
  language_id: 10
Ada:
  type: programming
  color: "#02f88c"
  extensions:
  - ".adb"
  - ".ada"
  - ".ads"
  aliases:
  - ada95
  - ada2005
  tm_scope: source.ada
  ace_mode: ada
  language_id: 11
Adobe Font Metrics:
  type: data
  color: "#fa0f00"
  tm_scope: source.afm
  extensions:
  - ".afm"
  aliases:
  - acfm
  - adobe composite font metrics
  - adobe multiple font metrics
  - amfm
  ace_mode: text
  language_id: 147198098
Agda:
  type: programming
  color: "#315665"
  extensions:
  - ".agda"
  tm_scope: source.agda
  ace_mode: text
  language_id: 12
Alloy:
  type: programming
  color: "#64C800"
  extensions:
  - ".als"
  tm_scope: source.alloy
  ace_mode: text
  language_id: 13
Alpine Abuild:
  type: programming
  color: "#0D597F"
  group: Shell
  aliases:
  - abuild
  - apkbuild
  filenames:
  - APKBUILD
  tm_scope: source.shell
  ace_mode: sh
  codemirror_mode: shell
  codemirror_mime_type: text/x-sh
  language_id: 14
Altium Designer:
  type: data
  color: "#A89663"
  aliases:
  - altium
  extensions:
  - ".OutJob"
  - ".PcbDoc"
  - ".PrjPCB"
  - ".SchDoc"
  tm_scope: source.ini
  ace_mode: ini
  language_id: 187772328
AngelScript:
  type: programming
  color: "#C7D7DC"
  extensions:
  - ".as"
  - ".angelscript"
  tm_scope: source.angelscript
  ace_mode: text
  codemirror_mode: clike
  codemirror_mime_type: text/x-c++src
  language_id: 389477596
Ant Build System:
  type: data
  color: "#A9157E"
  tm_scope: text.xml.ant
  filenames:
  - ant.xml
  - build.xml
  ace_mode: xml
  codemirror_mode: xml
  codemirror_mime_type: application/xml
  language_id: 15
ApacheConf:
  type: data
  color: "#d12127"
  aliases:
  - aconf
  - apache
  extensions:
  - ".apacheconf"
  - ".vhost"
  filenames:
  - ".htaccess"
  - apache2.conf
  - httpd.conf
  tm_scope: source.apache-config
  ace_mode: apache_conf
  language_id: 16
Apex:
  type: programming
  color: "#1797c0"
  extensions:
  - ".cls"
  tm_scope: source.java
  ace_mode: java
  codemirror_mode: clike
  codemirror_mime_type: text/x-java
  language_id: 17
Apollo Guidance Computer:
  type: programming
  color: "#0B3D91"
  group: Assembly
  extensions:
  - ".agc"
  tm_scope: source.agc
  ace_mode: assembly_x86
  language_id: 18
AppleScript:
  type: programming
  aliases:
  - osascript
  extensions:
  - ".applescript"
  - ".scpt"
  interpreters:
  - osascript
  tm_scope: source.applescript
  ace_mode: applescript
  color: "#101F1F"
  language_id: 19
Arc:
  type: programming
  color: "#aa2afe"
  extensions:
  - ".arc"
  tm_scope: none
  ace_mode: text
  language_id: 20
AsciiDoc:
  type: prose
  color: "#73a0c5"
  ace_mode: asciidoc
  wrap: true
  extensions:
  - ".asciidoc"
  - ".adoc"
  - ".asc"
  tm_scope: text.html.asciidoc
  language_id: 22
AspectJ:
  type: programming
  color: "#a957b0"
  extensions:
  - ".aj"
  tm_scope: source.aspectj
  ace_mode: text
  language_id: 23
Assembly:
  type: programming
  color: "#6E4C13"
  aliases:
  - asm
  - nasm
  extensions:
  - ".asm"
  - ".a51"
  - ".i"
  - ".inc"
  - ".nasm"
  tm_scope: source.assembly
  ace_mode: assembly_x86
  language_id: 24
Astro:
  type: markup
  color: "#ff5a03"
  extensions:
  - ".astro"
  tm_scope: source.astro
  ace_mode: html
  codemirror_mode: jsx
  codemirror_mime_type: text/jsx
  language_id: 578209015
Asymptote:
  type: programming
  color: "#ff0000"
  extensions:
  - ".asy"
  interpreters:
  - asy
  tm_scope: source.c++
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-kotlin
  language_id: 591605007
Augeas:
  type: programming
  color: "#9CC134"
  extensions:
  - ".aug"
  tm_scope: none
  ace_mode: text
  language_id: 25
AutoHotkey:
  type: programming
  color: "#6594b9"
  aliases:
  - ahk
  extensions:
  - ".ahk"
  - ".ahkl"
  tm_scope: source.ahk
  ace_mode: autohotkey
  language_id: 26
AutoIt:
  type: programming
  color: "#1C3552"
  aliases:
  - au3
  - AutoIt3
  - AutoItScript
  extensions:
  - ".au3"
  tm_scope: source.autoit
  ace_mode: autohotkey
  language_id: 27
Avro IDL:
  type: data
  color: "#0040FF"
  extensions:
  - ".avdl"
  tm_scope: source.avro
  ace_mode: text
  language_id: 785497837
Awk:
  type: programming
  color: "#c30e9b"
  extensions:
  - ".awk"
  - ".auk"
  - ".gawk"
  - ".mawk"
  - ".nawk"
  interpreters:
  - awk
  - gawk
  - mawk
  - nawk
  tm_scope: source.awk
  ace_mode: text
  language_id: 28
BASIC:
  type: programming
  extensions:
  - ".bas"
  tm_scope: source.basic
  ace_mode: text
  color: "#ff0000"
  language_id: 28923963
Ballerina:
  type: programming
  extensions:
  - ".bal"
  tm_scope: source.ballerina
  ace_mode: text
  color: "#FF5000"
  language_id: 720859680
Batchfile:
  type: programming
  aliases:
  - bat
  - batch
  - dosbatch
  - winbatch
  extensions:
  - ".bat"
  - ".cmd"
  tm_scope: source.batchfile
  ace_mode: batchfile
  color: "#C1F12E"
  language_id: 29
Beef:
  type: programming
  color: "#a52f4e"
  extensions:
  - ".bf"
  tm_scope: source.cs
  ace_mode: csharp
  codemirror_mode: clike
  codemirror_mime_type: text/x-csharp
  language_id: 545626333
Befunge:
  type: programming
  extensions:
  - ".befunge"
  tm_scope: source.befunge
  ace_mode: text
  language_id: 30
Berry:
  type: programming
  extensions:
  - ".be"
  tm_scope: source.berry
  ace_mode: text
  color: "#15A13C"
  aliases:
  - be
  language_id: 121855308
BibTeX:
  type: markup
  color: "#778899"
  group: TeX
  extensions:
  - ".bib"
  - ".bibtex"
  tm_scope: text.bibtex
  ace_mode: tex
  codemirror_mode: stex
  codemirror_mime_type: text/x-stex
  language_id: 982188347
Bicep:
  type: programming
  color: "#519aba"
  extensions:
  - ".bicep"
  tm_scope: source.bicep
  ace_mode: text
  language_id: 321200902
Bison:
  type: programming
  color: "#6A463F"
  group: Yacc
  tm_scope: source.yacc
  extensions:
  - ".bison"
  ace_mode: text
  language_id: 31
BitBake:
  type: programming
  color: "#00bce4"
  tm_scope: none
  extensions:
  - ".bb"
  ace_mode: text
  language_id: 32
Blade:
  type: markup
  color: "#f7523f"
  extensions:
  - ".blade"
  - ".blade.php"
  tm_scope: text.html.php.blade
  ace_mode: text
  language_id: 33
BlitzBasic:
  type: programming
  color: "#00FFAE"
  aliases:
  - b3d
  - blitz3d
  - blitzplus
  - bplus
  extensions:
  - ".bb"
  - ".decls"
  tm_scope: source.blitzmax
  ace_mode: text
  language_id: 34
BlitzMax:
  type: programming
  color: "#cd6400"
  extensions:
  - ".bmx"
  aliases:
  - bmax
  tm_scope: source.blitzmax
  ace_mode: text
  language_id: 35
Bluespec:
  type: programming
  color: "#12223c"
  extensions:
  - ".bsv"
  tm_scope: source.bsv
  ace_mode: verilog
  language_id: 36
Boo:
  type: programming
  color: "#d4bec1"
  extensions:
  - ".boo"
  ace_mode: text
  tm_scope: source.boo
  language_id: 37
Boogie:
  type: programming
  color: "#c80fa0"
  extensions:
  - ".bpl"
  interpreters:
  - boogie
  tm_scope: source.boogie
  ace_mode: text
  language_id: 955017407
Brainfuck:
  type: programming
  color: "#2F2530"
  extensions:
  - ".b"
  - ".bf"
  tm_scope: source.bf
  ace_mode: text
  codemirror_mode: brainfuck
  codemirror_mime_type: text/x-brainfuck
  language_id: 38
Brightscript:
  type: programming
  color: "#662D91"
  extensions:
  - ".brs"
  tm_scope: source.brightscript
  ace_mode: text
  language_id: 39
Browserslist:
  type: data
  color: "#ffd539"
  filenames:
  - ".browserslistrc"
  - browserslist
  tm_scope: text.browserslist
  ace_mode: text
  language_id: 153503348
C:
  type: programming
  color: "#555555"
  extensions:
  - ".c"
  - ".cats"
  - ".h"
  - ".idc"
  interpreters:
  - tcc
  tm_scope: source.c
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-csrc
  language_id: 41
C#:
  type: programming
  ace_mode: csharp
  codemirror_mode: clike
  codemirror_mime_type: text/x-csharp
  tm_scope: source.cs
  color: "#178600"
  aliases:
  - csharp
  - cake
  - cakescript
  extensions:
  - ".cs"
  - ".cake"
  - ".csx"
  - ".linq"
  language_id: 42
C++:
  type: programming
  tm_scope: source.c++
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-c++src
  color: "#f34b7d"
  aliases:
  - cpp
  extensions:
  - ".cpp"
  - ".c++"
  - ".cc"
  - ".cp"
  - ".cxx"
  - ".h"
  - ".h++"
  - ".hh"
  - ".hpp"
  - ".hxx"
  - ".inc"
  - ".inl"
  - ".ino"
  - ".ipp"
  - ".ixx"
  - ".re"
  - ".tcc"
  - ".tpp"
  language_id: 43
C-ObjDump:
  type: data
  extensions:
  - ".c-objdump"
  tm_scope: objdump.x86asm
  ace_mode: assembly_x86
  language_id: 44
C2hs Haskell:
  type: programming
  group: Haskell
  aliases:
  - c2hs
  extensions:
  - ".chs"
  tm_scope: source.haskell
  ace_mode: haskell
  codemirror_mode: haskell
  codemirror_mime_type: text/x-haskell
  language_id: 45
CIL:
  type: data
  tm_scope: source.cil
  extensions:
  - ".cil"
  ace_mode: text
  language_id: 29176339
CLIPS:
  type: programming
  color: "#00A300"
  extensions:
  - ".clp"
  tm_scope: source.clips
  ace_mode: text
  language_id: 46
CMake:
  type: programming
  color: "#DA3434"
  extensions:
  - ".cmake"
  - ".cmake.in"
  filenames:
  - CMakeLists.txt
  tm_scope: source.cmake
  ace_mode: text
  codemirror_mode: cmake
  codemirror_mime_type: text/x-cmake
  language_id: 47
COBOL:
  type: programming
  extensions:
  - ".cob"
  - ".cbl"
  - ".ccp"
  - ".cobol"
  - ".cpy"
  tm_scope: source.cobol
  ace_mode: cobol
  codemirror_mode: cobol
  codemirror_mime_type: text/x-cobol
  language_id: 48
CODEOWNERS:
  type: data
  filenames:
  - CODEOWNERS
  tm_scope: text.codeowners
  ace_mode: gitignore
  language_id: 321684729
COLLADA:
  type: data
  color: "#F1A42B"
  extensions:
  - ".dae"
  tm_scope: text.xml
  ace_mode: xml
  codemirror_mode: xml
  codemirror_mime_type: text/xml
  language_id: 49
CSON:
  type: data
  color: "#244776"
  tm_scope: source.coffee
  ace_mode: coffee
  codemirror_mode: coffeescript
  codemirror_mime_type: text/x-coffeescript
  extensions:
  - ".cson"
  language_id: 424
CSS:
  type: markup
  tm_scope: source.css
  ace_mode: css
  codemirror_mode: css
  codemirror_mime_type: text/css
  color: "#563d7c"
  extensions:
  - ".css"
  language_id: 50
CSV:
  type: data
  color: "#237346"
  ace_mode: text
  tm_scope: none
  extensions:
  - ".csv"
  language_id: 51
CUE:
  type: programming
  extensions:
  - ".cue"
  tm_scope: source.cue
  ace_mode: text
  color: "#5886E1"
  language_id: 356063509
CWeb:
  type: programming
  color: "#00007a"
  extensions:
  - ".w"
  tm_scope: none
  ace_mode: text
  language_id: 657332628
Cabal Config:
  type: data
  color: "#483465"
  aliases:
  - Cabal
  extensions:
  - ".cabal"
  filenames:
  - cabal.config
  - cabal.project
  ace_mode: haskell
  codemirror_mode: haskell
  codemirror_mime_type: text/x-haskell
  tm_scope: source.cabal
  language_id: 677095381
Cadence:
  type: programming
  color: "#00ef8b"
  ace_mode: text
  tm_scope: source.cadence
  extensions:
  - ".cdc"
  language_id: 270184138
Cairo:
  type: programming
  color: "#ff4a48"
  ace_mode: text
  tm_scope: source.cairo
  extensions:
  - ".cairo"
  language_id: 620599567
CameLIGO:
  type: programming
  color: "#3be133"
  extensions:
  - ".mligo"
  tm_scope: source.mligo
  ace_mode: ocaml
  codemirror_mode: mllike
  codemirror_mime_type: text/x-ocaml
  group: LigoLANG
  language_id: 829207807
Cap'n Proto:
  type: programming
  color: "#c42727"
  tm_scope: source.capnp
  extensions:
  - ".capnp"
  ace_mode: text
  language_id: 52
CartoCSS:
  type: programming
  aliases:
  - Carto
  extensions:
  - ".mss"
  ace_mode: text
  tm_scope: source.css.mss
  language_id: 53
Ceylon:
  type: programming
  color: "#dfa535"
  extensions:
  - ".ceylon"
  tm_scope: source.ceylon
  ace_mode: text
  language_id: 54
Chapel:
  type: programming
  color: "#8dc63f"
  aliases:
  - chpl
  extensions:
  - ".chpl"
  tm_scope: source.chapel
  ace_mode: text
  language_id: 55
Charity:
  type: programming
  extensions:
  - ".ch"
  tm_scope: none
  ace_mode: text
  language_id: 56
ChucK:
  type: programming
  color: "#3f8000"
  extensions:
  - ".ck"
  tm_scope: source.java
  ace_mode: java
  codemirror_mode: clike
  codemirror_mime_type: text/x-java
  language_id: 57
Cirru:
  type: programming
  color: "#ccccff"
  tm_scope: source.cirru
  ace_mode: cirru
  extensions:
  - ".cirru"
  language_id: 58
Clarion:
  type: programming
  color: "#db901e"
  ace_mode: text
  extensions:
  - ".clw"
  tm_scope: source.clarion
  language_id: 59
Clarity:
  type: programming
  color: "#5546ff"
  ace_mode: lisp
  extensions:
  - ".clar"
  tm_scope: source.clar
  language_id: 91493841
Classic ASP:
  type: programming
  color: "#6a40fd"
  tm_scope: text.html.asp
  aliases:
  - asp
  extensions:
  - ".asp"
  ace_mode: text
  language_id: 8
Clean:
  type: programming
  color: "#3F85AF"
  extensions:
  - ".icl"
  - ".dcl"
  tm_scope: source.clean
  ace_mode: text
  language_id: 60
Click:
  type: programming
  color: "#E4E6F3"
  extensions:
  - ".click"
  tm_scope: source.click
  ace_mode: text
  language_id: 61
Clojure:
  type: programming
  tm_scope: source.clojure
  ace_mode: clojure
  codemirror_mode: clojure
  codemirror_mime_type: text/x-clojure
  color: "#db5855"
  extensions:
  - ".clj"
  - ".boot"
  - ".cl2"
  - ".cljc"
  - ".cljs"
  - ".cljs.hl"
  - ".cljscm"
  - ".cljx"
  - ".hic"
  filenames:
  - riemann.config
  language_id: 62
Closure Templates:
  type: markup
  color: "#0d948f"
  ace_mode: soy_template
  codemirror_mode: soy
  codemirror_mime_type: text/x-soy
  aliases:
  - soy
  extensions:
  - ".soy"
  tm_scope: text.html.soy
  language_id: 357046146
Cloud Firestore Security Rules:
  type: data
  color: "#FFA000"
  ace_mode: less
  codemirror_mode: css
  codemirror_mime_type: text/css
  tm_scope: source.firestore
  filenames:
  - firestore.rules
  language_id: 407996372
CoNLL-U:
  type: data
  extensions:
  - ".conllu"
  - ".conll"
  tm_scope: text.conllu
  ace_mode: text
  aliases:
  - CoNLL
  - CoNLL-X
  language_id: 421026389
CodeQL:
  type: programming
  color: "#140f46"
  extensions:
  - ".ql"
  - ".qll"
  tm_scope: source.ql
  ace_mode: text
  language_id: 424259634
  aliases:
  - ql
CoffeeScript:
  type: programming
  tm_scope: source.coffee
  ace_mode: coffee
  codemirror_mode: coffeescript
  codemirror_mime_type: text/x-coffeescript
  color: "#244776"
  aliases:
  - coffee
  - coffee-script
  extensions:
  - ".coffee"
  - "._coffee"
  - ".cake"
  - ".cjsx"
  - ".iced"
  filenames:
  - Cakefile
  interpreters:
  - coffee
  language_id: 63
ColdFusion:
  type: programming
  ace_mode: coldfusion
  color: "#ed2cd6"
  aliases:
  - cfm
  - cfml
  - coldfusion html
  extensions:
  - ".cfm"
  - ".cfml"
  tm_scope: text.html.cfm
  language_id: 64
ColdFusion CFC:
  type: programming
  color: "#ed2cd6"
  group: ColdFusion
  ace_mode: coldfusion
  aliases:
  - cfc
  extensions:
  - ".cfc"
  tm_scope: source.cfscript
  language_id: 65
Common Lisp:
  type: programming
  tm_scope: source.lisp
  color: "#3fb68b"
  aliases:
  - lisp
  extensions:
  - ".lisp"
  - ".asd"
  - ".cl"
  - ".l"
  - ".lsp"
  - ".ny"
  - ".podsl"
  - ".sexp"
  interpreters:
  - lisp
  - sbcl
  - ccl
  - clisp
  - ecl
  ace_mode: lisp
  codemirror_mode: commonlisp
  codemirror_mime_type: text/x-common-lisp
  language_id: 66
Common Workflow Language:
  aliases:
  - cwl
  type: programming
  ace_mode: yaml
  codemirror_mode: yaml
  codemirror_mime_type: text/x-yaml
  extensions:
  - ".cwl"
  interpreters:
  - cwl-runner
  color: "#B5314C"
  tm_scope: source.cwl
  language_id: 988547172
Component Pascal:
  type: programming
  color: "#B0CE4E"
  extensions:
  - ".cp"
  - ".cps"
  tm_scope: source.pascal
  ace_mode: pascal
  codemirror_mode: pascal
  codemirror_mime_type: text/x-pascal
  language_id: 67
Cool:
  type: programming
  extensions:
  - ".cl"
  tm_scope: source.cool
  ace_mode: text
  language_id: 68
Coq:
  type: programming
  color: "#d0b68c"
  extensions:
  - ".coq"
  - ".v"
  tm_scope: source.coq
  ace_mode: text
  language_id: 69
Cpp-ObjDump:
  type: data
  extensions:
  - ".cppobjdump"
  - ".c++-objdump"
  - ".c++objdump"
  - ".cpp-objdump"
  - ".cxx-objdump"
  tm_scope: objdump.x86asm
  aliases:
  - c++-objdump
  ace_mode: assembly_x86
  language_id: 70
Creole:
  type: prose
  wrap: true
  extensions:
  - ".creole"
  tm_scope: text.html.creole
  ace_mode: text
  language_id: 71
Crystal:
  type: programming
  color: "#000100"
  extensions:
  - ".cr"
  ace_mode: ruby
  codemirror_mode: crystal
  codemirror_mime_type: text/x-crystal
  tm_scope: source.crystal
  interpreters:
  - crystal
  language_id: 72
Csound:
  type: programming
  color: "#1a1a1a"
  aliases:
  - csound-orc
  extensions:
  - ".orc"
  - ".udo"
  tm_scope: source.csound
  ace_mode: csound_orchestra
  language_id: 73
Csound Document:
  type: programming
  color: "#1a1a1a"
  aliases:
  - csound-csd
  extensions:
  - ".csd"
  tm_scope: source.csound-document
  ace_mode: csound_document
  language_id: 74
Csound Score:
  type: programming
  color: "#1a1a1a"
  aliases:
  - csound-sco
  extensions:
  - ".sco"
  tm_scope: source.csound-score
  ace_mode: csound_score
  language_id: 75
Cuda:
  type: programming
  extensions:
  - ".cu"
  - ".cuh"
  tm_scope: source.cuda-c++
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-c++src
  color: "#3A4E3A"
  language_id: 77
Cue Sheet:
  type: data
  extensions:
  - ".cue"
  tm_scope: source.cuesheet
  ace_mode: text
  language_id: 942714150
Curry:
  type: programming
  color: "#531242"
  extensions:
  - ".curry"
  tm_scope: source.curry
  ace_mode: haskell
  language_id: 439829048
Cycript:
  type: programming
  extensions:
  - ".cy"
  tm_scope: source.js
  ace_mode: javascript
  codemirror_mode: javascript
  codemirror_mime_type: text/javascript
  language_id: 78
Cython:
  type: programming
  color: "#fedf5b"
  extensions:
  - ".pyx"
  - ".pxd"
  - ".pxi"
  aliases:
  - pyrex
  tm_scope: source.cython
  ace_mode: text
  codemirror_mode: python
  codemirror_mime_type: text/x-cython
  language_id: 79
D:
  type: programming
  color: "#ba595e"
  aliases:
  - Dlang
  extensions:
  - ".d"
  - ".di"
  tm_scope: source.d
  ace_mode: d
  codemirror_mode: d
  codemirror_mime_type: text/x-d
  language_id: 80
D-ObjDump:
  type: data
  extensions:
  - ".d-objdump"
  tm_scope: objdump.x86asm
  ace_mode: assembly_x86
  language_id: 81
DIGITAL Command Language:
  type: programming
  aliases:
  - dcl
  extensions:
  - ".com"
  tm_scope: none
  ace_mode: text
  language_id: 82
DM:
  type: programming
  color: "#447265"
  extensions:
  - ".dm"
  aliases:
  - byond
  tm_scope: source.dm
  ace_mode: c_cpp
  language_id: 83
DNS Zone:
  type: data
  extensions:
  - ".zone"
  - ".arpa"
  tm_scope: text.zone_file
  ace_mode: text
  language_id: 84
DTrace:
  type: programming
  aliases:
  - dtrace-script
  extensions:
  - ".d"
  interpreters:
  - dtrace
  tm_scope: source.c
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-csrc
  language_id: 85
Dafny:
  type: programming
  color: "#FFEC25"
  extensions:
  - ".dfy"
  interpreters:
  - dafny
  tm_scope: text.dfy.dafny
  ace_mode: text
  language_id: 969323346
Darcs Patch:
  type: data
  color: "#8eff23"
  aliases:
  - dpatch
  extensions:
  - ".darcspatch"
  - ".dpatch"
  tm_scope: none
  ace_mode: text
  language_id: 86
Dart:
  type: programming
  color: "#00B4AB"
  extensions:
  - ".dart"
  interpreters:
  - dart
  tm_scope: source.dart
  ace_mode: dart
  codemirror_mode: dart
  codemirror_mime_type: application/dart
  language_id: 87
DataWeave:
  type: programming
  color: "#003a52"
  extensions:
  - ".dwl"
  ace_mode: text
  tm_scope: source.data-weave
  language_id: 974514097
Debian Package Control File:
  type: data
  color: "#D70751"
  extensions:
  - ".dsc"
  tm_scope: source.deb-control
  ace_mode: text
  language_id: 527438264
DenizenScript:
  type: programming
  color: "#FBEE96"
  ace_mode: yaml
  codemirror_mode: yaml
  codemirror_mime_type: text/x-yaml
  extensions:
  - ".dsc"
  tm_scope: source.denizenscript
  language_id: 435000929
Dhall:
  type: programming
  color: "#dfafff"
  extensions:
  - ".dhall"
  tm_scope: source.haskell
  ace_mode: haskell
  codemirror_mode: haskell
  codemirror_mime_type: text/x-haskell
  language_id: 793969321
Diff:
  type: data
  extensions:
  - ".diff"
  - ".patch"
  aliases:
  - udiff
  tm_scope: source.diff
  ace_mode: diff
  codemirror_mode: diff
  codemirror_mime_type: text/x-diff
  language_id: 88
DirectX 3D File:
  type: data
  color: "#aace60"
  extensions:
  - ".x"
  ace_mode: text
  tm_scope: none
  language_id: 201049282
Dockerfile:
  type: programming
  aliases:
  - Containerfile
  color: "#384d54"
  tm_scope: source.dockerfile
  extensions:
  - ".dockerfile"
  filenames:
  - Containerfile
  - Dockerfile
  ace_mode: dockerfile
  codemirror_mode: dockerfile
  codemirror_mime_type: text/x-dockerfile
  language_id: 89
Dogescript:
  type: programming
  color: "#cca760"
  extensions:
  - ".djs"
  tm_scope: none
  ace_mode: text
  language_id: 90
Dylan:
  type: programming
  color: "#6c616e"
  extensions:
  - ".dylan"
  - ".dyl"
  - ".intr"
  - ".lid"
  tm_scope: source.dylan
  ace_mode: text
  codemirror_mode: dylan
  codemirror_mime_type: text/x-dylan
  language_id: 91
E:
  type: programming
  color: "#ccce35"
  extensions:
  - ".e"
  interpreters:
  - rune
  tm_scope: none
  ace_mode: text
  language_id: 92
E-mail:
  type: data
  aliases:
  - email
  - eml
  - mail
  - mbox
  extensions:
  - ".eml"
  - ".mbox"
  tm_scope: text.eml.basic
  ace_mode: text
  codemirror_mode: mbox
  codemirror_mime_type: application/mbox
  language_id: 529653389
EBNF:
  type: data
  extensions:
  - ".ebnf"
  tm_scope: source.ebnf
  ace_mode: text
  codemirror_mode: ebnf
  codemirror_mime_type: text/x-ebnf
  language_id: 430
ECL:
  type: programming
  color: "#8a1267"
  extensions:
  - ".ecl"
  - ".eclxml"
  tm_scope: source.ecl
  ace_mode: text
  codemirror_mode: ecl
  codemirror_mime_type: text/x-ecl
  language_id: 93
ECLiPSe:
  type: programming
  color: "#001d9d"
  group: prolog
  extensions:
  - ".ecl"
  tm_scope: source.prolog.eclipse
  ace_mode: prolog
  language_id: 94
EJS:
  type: markup
  color: "#a91e50"
  extensions:
  - ".ejs"
  - ".ect"
  - ".ejs.t"
  - ".jst"
  tm_scope: text.html.js
  ace_mode: ejs
  language_id: 95
EQ:
  type: programming
  color: "#a78649"
  extensions:
  - ".eq"
  tm_scope: source.cs
  ace_mode: csharp
  codemirror_mode: clike
  codemirror_mime_type: text/x-csharp
  language_id: 96
Eagle:
  type: data
  extensions:
  - ".sch"
  - ".brd"
  tm_scope: text.xml
  ace_mode: xml
  codemirror_mode: xml
  codemirror_mime_type: text/xml
  language_id: 97
Earthly:
  type: programming
  aliases:
  - Earthfile
  color: "#2af0ff"
  tm_scope: source.earthfile
  ace_mode: text
  filenames:
  - Earthfile
  language_id: 963512632
Easybuild:
  type: data
  color: "#069406"
  group: Python
  ace_mode: python
  codemirror_mode: python
  codemirror_mime_type: text/x-python
  tm_scope: source.python
  extensions:
  - ".eb"
  language_id: 342840477
Ecere Projects:
  type: data
  color: "#913960"
  group: JavaScript
  extensions:
  - ".epj"
  tm_scope: source.json
  ace_mode: json
  codemirror_mode: javascript
  codemirror_mime_type: application/json
  language_id: 98
EditorConfig:
  type: data
  color: "#fff1f2"
  group: INI
  filenames:
  - ".editorconfig"
  aliases:
  - editor-config
  ace_mode: ini
  codemirror_mode: properties
  codemirror_mime_type: text/x-properties
  tm_scope: source.editorconfig
  language_id: 96139566
Edje Data Collection:
  type: data
  extensions:
  - ".edc"
  tm_scope: source.c++
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-c++src
  language_id: 342840478
Eiffel:
  type: programming
  color: "#4d6977"
  extensions:
  - ".e"
  tm_scope: source.eiffel
  ace_mode: eiffel
  codemirror_mode: eiffel
  codemirror_mime_type: text/x-eiffel
  language_id: 99
Elixir:
  type: programming
  color: "#6e4a7e"
  extensions:
  - ".ex"
  - ".exs"
  tm_scope: source.elixir
  ace_mode: elixir
  filenames:
  - mix.lock
  interpreters:
  - elixir
  language_id: 100
Elm:
  type: programming
  color: "#60B5CC"
  extensions:
  - ".elm"
  tm_scope: source.elm
  ace_mode: elm
  codemirror_mode: elm
  codemirror_mime_type: text/x-elm
  language_id: 101
Emacs Lisp:
  type: programming
  tm_scope: source.emacs.lisp
  color: "#c065db"
  aliases:
  - elisp
  - emacs
  filenames:
  - ".abbrev_defs"
  - ".emacs"
  - ".emacs.desktop"
  - ".gnus"
  - ".spacemacs"
  - ".viper"
  - Cask
  - Project.ede
  - _emacs
  - abbrev_defs
  extensions:
  - ".el"
  - ".emacs"
  - ".emacs.desktop"
  ace_mode: lisp
  codemirror_mode: commonlisp
  codemirror_mime_type: text/x-common-lisp
  language_id: 102
EmberScript:
  type: programming
  color: "#FFF4F3"
  extensions:
  - ".em"
  - ".emberscript"
  tm_scope: source.coffee
  ace_mode: coffee
  codemirror_mode: coffeescript
  codemirror_mime_type: text/x-coffeescript
  language_id: 103
Erlang:
  type: programming
  color: "#B83998"
  extensions:
  - ".erl"
  - ".app.src"
  - ".es"
  - ".escript"
  - ".hrl"
  - ".xrl"
  - ".yrl"
  filenames:
  - Emakefile
  - rebar.config
  - rebar.config.lock
  - rebar.lock
  tm_scope: source.erlang
  ace_mode: erlang
  codemirror_mode: erlang
  codemirror_mime_type: text/x-erlang
  interpreters:
  - escript
  language_id: 104
Euphoria:
  type: programming
  color: "#FF790B"
  extensions:
  - ".e"
  - ".ex"
  interpreters:
  - eui
  - euiw
  ace_mode: text
  tm_scope: source.euphoria
  language_id: 880693982
F#:
  type: programming
  color: "#b845fc"
  aliases:
  - fsharp
  extensions:
  - ".fs"
  - ".fsi"
  - ".fsx"
  tm_scope: source.fsharp
  ace_mode: text
  codemirror_mode: mllike
  codemirror_mime_type: text/x-fsharp
  language_id: 105
F*:
  fs_name: Fstar
  type: programming
  color: "#572e30"
  aliases:
  - fstar
  extensions:
  - ".fst"
  tm_scope: source.fstar
  ace_mode: text
  language_id: 336943375
FIGlet Font:
  type: data
  color: "#FFDDBB"
  aliases:
  - FIGfont
  extensions:
  - ".flf"
  tm_scope: source.figfont
  ace_mode: text
  language_id: 686129783
FLUX:
  type: programming
  color: "#88ccff"
  extensions:
  - ".fx"
  - ".flux"
  tm_scope: none
  ace_mode: text
  language_id: 106
Factor:
  type: programming
  color: "#636746"
  extensions:
  - ".factor"
  filenames:
  - ".factor-boot-rc"
  - ".factor-rc"
  tm_scope: source.factor
  ace_mode: text
  codemirror_mode: factor
  codemirror_mime_type: text/x-factor
  language_id: 108
Fancy:
  type: programming
  color: "#7b9db4"
  extensions:
  - ".fy"
  - ".fancypack"
  filenames:
  - Fakefile
  tm_scope: source.fancy
  ace_mode: text
  language_id: 109
Fantom:
  type: programming
  color: "#14253c"
  extensions:
  - ".fan"
  tm_scope: source.fan
  ace_mode: text
  language_id: 110
Faust:
  type: programming
  color: "#c37240"
  extensions:
  - ".dsp"
  tm_scope: source.faust
  ace_mode: text
  language_id: 622529198
Fennel:
  type: programming
  tm_scope: source.fnl
  ace_mode: text
  color: "#fff3d7"
  interpreters:
  - fennel
  extensions:
  - ".fnl"
  language_id: 239946126
Filebench WML:
  type: programming
  color: "#F6B900"
  extensions:
  - ".f"
  tm_scope: none
  ace_mode: text
  language_id: 111
Filterscript:
  type: programming
  group: RenderScript
  extensions:
  - ".fs"
  tm_scope: none
  ace_mode: text
  language_id: 112
Fluent:
  type: programming
  color: "#ffcc33"
  extensions:
  - ".ftl"
  tm_scope: source.ftl
  ace_mode: text
  language_id: 206353404
Formatted:
  type: data
  extensions:
  - ".for"
  - ".eam.fs"
  tm_scope: none
  ace_mode: text
  language_id: 113
Forth:
  type: programming
  color: "#341708"
  extensions:
  - ".fth"
  - ".4th"
  - ".f"
  - ".for"
  - ".forth"
  - ".fr"
  - ".frt"
  - ".fs"
  tm_scope: source.forth
  ace_mode: forth
  codemirror_mode: forth
  codemirror_mime_type: text/x-forth
  language_id: 114
Fortran:
  group: Fortran
  type: programming
  color: "#4d41b1"
  extensions:
  - ".f"
  - ".f77"
  - ".for"
  - ".fpp"
  tm_scope: source.fortran
  ace_mode: text
  codemirror_mode: fortran
  codemirror_mime_type: text/x-fortran
  language_id: 107
Fortran Free Form:
  group: Fortran
  color: "#4d41b1"
  type: programming
  extensions:
  - ".f90"
  - ".f03"
  - ".f08"
  - ".f95"
  tm_scope: source.fortran.modern
  ace_mode: text
  codemirror_mode: fortran
  codemirror_mime_type: text/x-fortran
  language_id: 761352333
FreeBasic:
  type: programming
  color: "#867db1"
  extensions:
  - ".bi"
  - ".bas"
  tm_scope: source.vbnet
  aliases:
  - fb
  ace_mode: text
  codemirror_mode: vb
  codemirror_mime_type: text/x-vb
  language_id: 472896659
FreeMarker:
  type: programming
  color: "#0050b2"
  aliases:
  - ftl
  extensions:
  - ".ftl"
  tm_scope: text.html.ftl
  ace_mode: ftl
  language_id: 115
Frege:
  type: programming
  color: "#00cafe"
  extensions:
  - ".fr"
  tm_scope: source.haskell
  ace_mode: haskell
  language_id: 116
Futhark:
  type: programming
  color: "#5f021f"
  extensions:
  - ".fut"
  tm_scope: source.futhark
  ace_mode: text
  language_id: 97358117
G-code:
  type: programming
  color: "#D08CF2"
  extensions:
  - ".g"
  - ".cnc"
  - ".gco"
  - ".gcode"
  tm_scope: source.gcode
  ace_mode: gcode
  language_id: 117
GAML:
  type: programming
  color: "#FFC766"
  extensions:
  - ".gaml"
  tm_scope: none
  ace_mode: text
  language_id: 290345951
GAMS:
  type: programming
  color: "#f49a22"
  extensions:
  - ".gms"
  tm_scope: none
  ace_mode: text
  language_id: 118
GAP:
  type: programming
  color: "#0000cc"
  extensions:
  - ".g"
  - ".gap"
  - ".gd"
  - ".gi"
  - ".tst"
  tm_scope: source.gap
  ace_mode: text
  language_id: 119
GCC Machine Description:
  type: programming
  color: "#FFCFAB"
  extensions:
  - ".md"
  tm_scope: source.lisp
  ace_mode: lisp
  codemirror_mode: commonlisp
  codemirror_mime_type: text/x-common-lisp
  language_id: 121
GDB:
  type: programming
  extensions:
  - ".gdb"
  - ".gdbinit"
  tm_scope: source.gdb
  ace_mode: text
  language_id: 122
GDScript:
  type: programming
  color: "#355570"
  extensions:
  - ".gd"
  tm_scope: source.gdscript
  ace_mode: text
  language_id: 123
GEDCOM:
  type: data
  color: "#003058"
  ace_mode: text
  extensions:
  - ".ged"
  tm_scope: source.gedcom
  language_id: 459577965
GLSL:
  type: programming
  color: "#5686a5"
  extensions:
  - ".glsl"
  - ".fp"
  - ".frag"
  - ".frg"
  - ".fs"
  - ".fsh"
  - ".fshader"
  - ".geo"
  - ".geom"
  - ".glslf"
  - ".glslv"
  - ".gs"
  - ".gshader"
  - ".rchit"
  - ".rmiss"
  - ".shader"
  - ".tesc"
  - ".tese"
  - ".vert"
  - ".vrx"
  - ".vsh"
  - ".vshader"
  tm_scope: source.glsl
  ace_mode: glsl
  language_id: 124
GN:
  type: data
  extensions:
  - ".gn"
  - ".gni"
  interpreters:
  - gn
  filenames:
  - ".gn"
  tm_scope: source.gn
  ace_mode: python
  codemirror_mode: python
  codemirror_mime_type: text/x-python
  language_id: 302957008
GSC:
  type: programming
  color: "#FF6800"
  extensions:
  - ".gsc"
  - ".csc"
  - ".gsh"
  tm_scope: source.gsc
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-csrc
  language_id: 257856279
Game Maker Language:
  type: programming
  color: "#71b417"
  extensions:
  - ".gml"
  tm_scope: source.c++
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-c++src
  language_id: 125
Gemfile.lock:
  type: data
  color: "#701516"
  searchable: false
  tm_scope: source.gemfile-lock
  ace_mode: text
  filenames:
  - Gemfile.lock
  language_id: 907065713
Genero:
  type: programming
  color: "#63408e"
  extensions:
  - ".4gl"
  tm_scope: source.genero
  ace_mode: text
  language_id: 986054050
Genero Forms:
  type: markup
  color: "#d8df39"
  extensions:
  - ".per"
  tm_scope: source.genero-forms
  ace_mode: text
  language_id: 902995658
Genie:
  type: programming
  ace_mode: text
  extensions:
  - ".gs"
  color: "#fb855d"
  tm_scope: none
  language_id: 792408528
Genshi:
  type: programming
  color: "#951531"
  extensions:
  - ".kid"
  tm_scope: text.xml.genshi
  aliases:
  - xml+genshi
  - xml+kid
  ace_mode: xml
  codemirror_mode: xml
  codemirror_mime_type: text/xml
  language_id: 126
Gentoo Ebuild:
  type: programming
  color: "#9400ff"
  group: Shell
  extensions:
  - ".ebuild"
  tm_scope: source.shell
  ace_mode: sh
  codemirror_mode: shell
  codemirror_mime_type: text/x-sh
  language_id: 127
Gentoo Eclass:
  type: programming
  color: "#9400ff"
  group: Shell
  extensions:
  - ".eclass"
  tm_scope: source.shell
  ace_mode: sh
  codemirror_mode: shell
  codemirror_mime_type: text/x-sh
  language_id: 128
Gerber Image:
  type: data
  color: "#d20b00"
  aliases:
  - rs-274x
  extensions:
  - ".gbr"
  - ".cmp"
  - ".gbl"
  - ".gbo"
  - ".gbp"
  - ".gbs"
  - ".gko"
  - ".gml"
  - ".gpb"
  - ".gpt"
  - ".gtl"
  - ".gto"
  - ".gtp"
  - ".gts"
  - ".ncl"
  - ".sol"
  interpreters:
  - gerbv
  - gerbview
  tm_scope: source.gerber
  ace_mode: text
  language_id: 404627610
Gettext Catalog:
  type: prose
  aliases:
  - pot
  extensions:
  - ".po"
  - ".pot"
  tm_scope: source.po
  ace_mode: text
  language_id: 129
Gherkin:
  type: programming
  extensions:
  - ".feature"
  - ".story"
  tm_scope: text.gherkin.feature
  aliases:
  - cucumber
  ace_mode: text
  color: "#5B2063"
  language_id: 76
Git Attributes:
  type: data
  color: "#F44D27"
  group: INI
  aliases:
  - gitattributes
  filenames:
  - ".gitattributes"
  tm_scope: source.gitattributes
  ace_mode: gitignore
  codemirror_mode: shell
  codemirror_mime_type: text/x-sh
  language_id: 956324166
Git Config:
  type: data
  color: "#F44D27"
  group: INI
  aliases:
  - gitconfig
  - gitmodules
  extensions:
  - ".gitconfig"
  filenames:
  - ".gitconfig"
  - ".gitmodules"
  ace_mode: ini
  codemirror_mode: properties
  codemirror_mime_type: text/x-properties
  tm_scope: source.gitconfig
  language_id: 807968997
Gleam:
  type: programming
  color: "#ffaff3"
  ace_mode: text
  extensions:
  - ".gleam"
  tm_scope: source.gleam
  language_id: 1054258749
Glyph:
  type: programming
  color: "#c1ac7f"
  extensions:
  - ".glf"
  tm_scope: source.tcl
  ace_mode: tcl
  codemirror_mode: tcl
  codemirror_mime_type: text/x-tcl
  language_id: 130
Glyph Bitmap Distribution Format:
  type: data
  extensions:
  - ".bdf"
  tm_scope: source.bdf
  ace_mode: text
  language_id: 997665271
Gnuplot:
  type: programming
  color: "#f0a9f0"
  extensions:
  - ".gp"
  - ".gnu"
  - ".gnuplot"
  - ".p"
  - ".plot"
  - ".plt"
  interpreters:
  - gnuplot
  tm_scope: source.gnuplot
  ace_mode: text
  language_id: 131
Go:
  type: programming
  color: "#00ADD8"
  aliases:
  - golang
  extensions:
  - ".go"
  tm_scope: source.go
  ace_mode: golang
  codemirror_mode: go
  codemirror_mime_type: text/x-go
  language_id: 132
Go Checksums:
  type: data
  color: "#00ADD8"
  aliases:
  - go.sum
  - go sum
  filenames:
  - go.sum
  tm_scope: go.sum
  ace_mode: text
  language_id: 1054391671
Go Module:
  type: data
  color: "#00ADD8"
  aliases:
  - go.mod
  - go mod
  filenames:
  - go.mod
  tm_scope: go.mod
  ace_mode: text
  language_id: 947461016
Golo:
  type: programming
  color: "#88562A"
  extensions:
  - ".golo"
  tm_scope: source.golo
  ace_mode: text
  language_id: 133
Gosu:
  type: programming
  color: "#82937f"
  extensions:
  - ".gs"
  - ".gst"
  - ".gsx"
  - ".vark"
  tm_scope: source.gosu.2
  ace_mode: text
  language_id: 134
Grace:
  type: programming
  color: "#615f8b"
  extensions:
  - ".grace"
  tm_scope: source.grace
  ace_mode: text
  language_id: 135
Gradle:
  type: data
  color: "#02303a"
  extensions:
  - ".gradle"
  tm_scope: source.groovy.gradle
  ace_mode: text
  language_id: 136
Grammatical Framework:
  type: programming
  aliases:
  - gf
  extensions:
  - ".gf"
  color: "#ff0000"
  tm_scope: source.gf
  ace_mode: haskell
  codemirror_mode: haskell
  codemirror_mime_type: text/x-haskell
  language_id: 137
Graph Modeling Language:
  type: data
  extensions:
  - ".gml"
  tm_scope: none
  ace_mode: text
  language_id: 138
GraphQL:
  type: data
  color: "#e10098"
  extensions:
  - ".graphql"
  - ".gql"
  - ".graphqls"
  tm_scope: source.graphql
  ace_mode: text
  language_id: 139
Graphviz (DOT):
  type: data
  color: "#2596be"
  tm_scope: source.dot
  extensions:
  - ".dot"
  - ".gv"
  ace_mode: text
  language_id: 140
Groovy:
  type: programming
  tm_scope: source.groovy
  ace_mode: groovy
  codemirror_mode: groovy
  codemirror_mime_type: text/x-groovy
  color: "#4298b8"
  extensions:
  - ".groovy"
  - ".grt"
  - ".gtpl"
  - ".gvy"
  interpreters:
  - groovy
  filenames:
  - Jenkinsfile
  language_id: 142
Groovy Server Pages:
  type: programming
  color: "#4298b8"
  group: Groovy
  aliases:
  - gsp
  - java server page
  extensions:
  - ".gsp"
  tm_scope: text.html.jsp
  ace_mode: jsp
  codemirror_mode: htmlembedded
  codemirror_mime_type: application/x-jsp
  language_id: 143
HAProxy:
  type: data
  color: "#106da9"
  extensions:
  - ".cfg"
  filenames:
  - haproxy.cfg
  tm_scope: source.haproxy-config
  ace_mode: text
  language_id: 366607477
HCL:
  type: programming
  extensions:
  - ".hcl"
  - ".nomad"
  - ".tf"
  - ".tfvars"
  - ".workflow"
  aliases:
  - HashiCorp Configuration Language
  - terraform
  ace_mode: ruby
  codemirror_mode: ruby
  codemirror_mime_type: text/x-ruby
  tm_scope: source.terraform
  language_id: 144
HLSL:
  type: programming
  color: "#aace60"
  extensions:
  - ".hlsl"
  - ".cginc"
  - ".fx"
  - ".fxh"
  - ".hlsli"
  ace_mode: text
  tm_scope: source.hlsl
  language_id: 145
HTML:
  type: markup
  tm_scope: text.html.basic
  ace_mode: html
  codemirror_mode: htmlmixed
  codemirror_mime_type: text/html
  color: "#e34c26"
  aliases:
  - xhtml
  extensions:
  - ".html"
  - ".hta"
  - ".htm"
  - ".html.hl"
  - ".inc"
  - ".xht"
  - ".xhtml"
  language_id: 146
HTML+ECR:
  type: markup
  color: "#2e1052"
  tm_scope: text.html.ecr
  group: HTML
  aliases:
  - ecr
  extensions:
  - ".ecr"
  ace_mode: text
  codemirror_mode: htmlmixed
  codemirror_mime_type: text/html
  language_id: 148
HTML+EEX:
  type: markup
  color: "#6e4a7e"
  tm_scope: text.html.elixir
  group: HTML
  aliases:
  - eex
  - heex
  - leex
  extensions:
  - ".eex"
  - ".html.heex"
  - ".html.leex"
  ace_mode: text
  codemirror_mode: htmlmixed
  codemirror_mime_type: text/html
  language_id: 149
HTML+ERB:
  type: markup
  color: "#701516"
  tm_scope: text.html.erb
  group: HTML
  aliases:
  - erb
  - rhtml
  - html+ruby
  extensions:
  - ".erb"
  - ".erb.deface"
  - ".rhtml"
  ace_mode: text
  codemirror_mode: htmlembedded
  codemirror_mime_type: application/x-erb
  language_id: 150
HTML+PHP:
  type: markup
  color: "#4f5d95"
  tm_scope: text.html.php
  group: HTML
  extensions:
  - ".phtml"
  ace_mode: php
  codemirror_mode: php
  codemirror_mime_type: application/x-httpd-php
  language_id: 151
HTML+Razor:
  type: markup
  color: "#512be4"
  tm_scope: text.html.cshtml
  group: HTML
  aliases:
  - razor
  extensions:
  - ".cshtml"
  - ".razor"
  ace_mode: razor
  codemirror_mode: htmlmixed
  codemirror_mime_type: text/html
  language_id: 479039817
HTTP:
  type: data
  color: "#005C9C"
  extensions:
  - ".http"
  tm_scope: source.httpspec
  ace_mode: text
  codemirror_mode: http
  codemirror_mime_type: message/http
  language_id: 152
HXML:
  type: data
  color: "#f68712"
  ace_mode: text
  extensions:
  - ".hxml"
  tm_scope: source.hxml
  language_id: 786683730
Hack:
  type: programming
  ace_mode: php
  codemirror_mode: php
  codemirror_mime_type: application/x-httpd-php
  extensions:
  - ".hack"
  - ".hh"
  - ".hhi"
  - ".php"
  tm_scope: source.hack
  color: "#878787"
  language_id: 153
Haml:
  type: markup
  color: "#ece2a9"
  extensions:
  - ".haml"
  - ".haml.deface"
  tm_scope: text.haml
  ace_mode: haml
  codemirror_mode: haml
  codemirror_mime_type: text/x-haml
  language_id: 154
Handlebars:
  type: markup
  color: "#f7931e"
  aliases:
  - hbs
  - htmlbars
  extensions:
  - ".handlebars"
  - ".hbs"
  tm_scope: text.html.handlebars
  ace_mode: handlebars
  language_id: 155
Harbour:
  type: programming
  color: "#0e60e3"
  extensions:
  - ".hb"
  tm_scope: source.harbour
  ace_mode: text
  language_id: 156
Haskell:
  type: programming
  color: "#5e5086"
  extensions:
  - ".hs"
  - ".hs-boot"
  - ".hsc"
  interpreters:
  - runghc
  - runhaskell
  - runhugs
  tm_scope: source.haskell
  ace_mode: haskell
  codemirror_mode: haskell
  codemirror_mime_type: text/x-haskell
  language_id: 157
Haxe:
  type: programming
  ace_mode: haxe
  codemirror_mode: haxe
  codemirror_mime_type: text/x-haxe
  color: "#df7900"
  extensions:
  - ".hx"
  - ".hxsl"
  tm_scope: source.hx
  language_id: 158
HiveQL:
  type: programming
  extensions:
  - ".q"
  - ".hql"
  color: "#dce200"
  tm_scope: source.hql
  ace_mode: sql
  language_id: 931814087
HolyC:
  type: programming
  color: "#ffefaf"
  extensions:
  - ".hc"
  tm_scope: source.hc
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-csrc
  language_id: 928121743
Hy:
  type: programming
  ace_mode: text
  color: "#7790B2"
  extensions:
  - ".hy"
  interpreters:
  - hy
  aliases:
  - hylang
  tm_scope: source.hy
  language_id: 159
HyPhy:
  type: programming
  ace_mode: text
  extensions:
  - ".bf"
  tm_scope: none
  language_id: 160
IDL:
  type: programming
  color: "#a3522f"
  extensions:
  - ".pro"
  - ".dlm"
  tm_scope: source.idl
  ace_mode: text
  codemirror_mode: idl
  codemirror_mime_type: text/x-idl
  language_id: 161
IGOR Pro:
  type: programming
  color: "#0000cc"
  extensions:
  - ".ipf"
  aliases:
  - igor
  - igorpro
  tm_scope: source.igor
  ace_mode: text
  language_id: 162
INI:
  type: data
  color: "#d1dbe0"
  extensions:
  - ".ini"
  - ".cfg"
  - ".dof"
  - ".lektorproject"
  - ".prefs"
  - ".pro"
  - ".properties"
  - ".url"
  filenames:
  - ".flake8"
  - buildozer.spec
  tm_scope: source.ini
  aliases:
  - dosini
  ace_mode: ini
  codemirror_mode: properties
  codemirror_mime_type: text/x-properties
  language_id: 163
IRC log:
  type: data
  aliases:
  - irc
  - irc logs
  extensions:
  - ".irclog"
  - ".weechatlog"
  tm_scope: none
  ace_mode: text
  codemirror_mode: mirc
  codemirror_mime_type: text/mirc
  language_id: 164
Idris:
  type: programming
  color: "#b30000"
  extensions:
  - ".idr"
  - ".lidr"
  ace_mode: text
  tm_scope: source.idris
  language_id: 165
Ignore List:
  type: data
  color: "#000000"
  group: INI
  aliases:
  - ignore
  - gitignore
  - git-ignore
  extensions:
  - ".gitignore"
  filenames:
  - ".atomignore"
  - ".babelignore"
  - ".bzrignore"
  - ".coffeelintignore"
  - ".cvsignore"
  - ".dockerignore"
  - ".eleventyignore"
  - ".eslintignore"
  - ".gitignore"
  - ".markdownlintignore"
  - ".nodemonignore"
  - ".npmignore"
  - ".prettierignore"
  - ".stylelintignore"
  - ".vercelignore"
  - ".vscodeignore"
  - gitignore-global
  - gitignore_global
  ace_mode: gitignore
  tm_scope: source.gitignore
  codemirror_mode: shell
  codemirror_mime_type: text/x-sh
  language_id: 74444240
ImageJ Macro:
  type: programming
  color: "#99AAFF"
  aliases:
  - ijm
  extensions:
  - ".ijm"
  ace_mode: text
  tm_scope: none
  language_id: 575143428
Inform 7:
  type: programming
  wrap: true
  extensions:
  - ".ni"
  - ".i7x"
  tm_scope: source.inform7
  aliases:
  - i7
  - inform7
  ace_mode: text
  language_id: 166
Inno Setup:
  type: programming
  color: "#264b99"
  extensions:
  - ".iss"
  - ".isl"
  tm_scope: source.inno
  ace_mode: text
  language_id: 167
Io:
  type: programming
  color: "#a9188d"
  extensions:
  - ".io"
  interpreters:
  - io
  tm_scope: source.io
  ace_mode: io
  language_id: 168
Ioke:
  type: programming
  color: "#078193"
  extensions:
  - ".ik"
  interpreters:
  - ioke
  tm_scope: source.ioke
  ace_mode: text
  language_id: 169
Isabelle:
  type: programming
  color: "#FEFE00"
  extensions:
  - ".thy"
  tm_scope: source.isabelle.theory
  ace_mode: text
  language_id: 170
Isabelle ROOT:
  type: programming
  color: "#FEFE00"
  group: Isabelle
  filenames:
  - ROOT
  tm_scope: source.isabelle.root
  ace_mode: text
  language_id: 171
J:
  type: programming
  color: "#9EEDFF"
  extensions:
  - ".ijs"
  interpreters:
  - jconsole
  tm_scope: source.j
  ace_mode: text
  language_id: 172
JAR Manifest:
  type: data
  color: "#b07219"
  filenames:
  - MANIFEST.MF
  tm_scope: source.yaml
  ace_mode: text
  language_id: 447261135
JFlex:
  type: programming
  color: "#DBCA00"
  group: Lex
  extensions:
  - ".flex"
  - ".jflex"
  tm_scope: source.jflex
  ace_mode: text
  language_id: 173
JSON:
  type: data
  color: "#292929"
  tm_scope: source.json
  ace_mode: json
  codemirror_mode: javascript
  codemirror_mime_type: application/json
  extensions:
  - ".json"
  - ".avsc"
  - ".geojson"
  - ".gltf"
  - ".har"
  - ".ice"
  - ".JSON-tmLanguage"
  - ".jsonl"
  - ".mcmeta"
  - ".tfstate"
  - ".tfstate.backup"
  - ".topojson"
  - ".webapp"
  - ".webmanifest"
  - ".yy"
  - ".yyp"
  filenames:
  - ".arcconfig"
  - ".auto-changelog"
  - ".c8rc"
  - ".htmlhintrc"
  - ".imgbotconfig"
  - ".nycrc"
  - ".tern-config"
  - ".tern-project"
  - ".watchmanconfig"
  - Pipfile.lock
  - composer.lock
  - mcmod.info
  language_id: 174
JSON with Comments:
  type: data
  color: "#292929"
  group: JSON
  tm_scope: source.js
  ace_mode: javascript
  codemirror_mode: javascript
  codemirror_mime_type: text/javascript
  aliases:
  - jsonc
  extensions:
  - ".jsonc"
  - ".code-snippets"
  - ".sublime-build"
  - ".sublime-commands"
  - ".sublime-completions"
  - ".sublime-keymap"
  - ".sublime-macro"
  - ".sublime-menu"
  - ".sublime-mousemap"
  - ".sublime-project"
  - ".sublime-settings"
  - ".sublime-theme"
  - ".sublime-workspace"
  - ".sublime_metrics"
  - ".sublime_session"
  filenames:
  - ".babelrc"
  - ".eslintrc.json"
  - ".jscsrc"
  - ".jshintrc"
  - ".jslintrc"
  - api-extractor.json
  - devcontainer.json
  - jsconfig.json
  - language-configuration.json
  - tsconfig.json
  - tslint.json
  language_id: 423
JSON5:
  type: data
  color: "#267CB9"
  extensions:
  - ".json5"
  tm_scope: source.js
  ace_mode: javascript
  codemirror_mode: javascript
  codemirror_mime_type: application/json
  language_id: 175
JSONLD:
  type: data
  color: "#0c479c"
  extensions:
  - ".jsonld"
  tm_scope: source.js
  ace_mode: javascript
  codemirror_mode: javascript
  codemirror_mime_type: application/json
  language_id: 176
JSONiq:
  color: "#40d47e"
  type: programming
  ace_mode: jsoniq
  codemirror_mode: javascript
  codemirror_mime_type: application/json
  extensions:
  - ".jq"
  tm_scope: source.jsoniq
  language_id: 177
Janet:
  type: programming
  color: "#0886a5"
  extensions:
  - ".janet"
  tm_scope: source.janet
  ace_mode: scheme
  codemirror_mode: scheme
  codemirror_mime_type: text/x-scheme
  interpreters:
  - janet
  language_id: 1028705371
Jasmin:
  type: programming
  color: "#d03600"
  ace_mode: java
  extensions:
  - ".j"
  tm_scope: source.jasmin
  language_id: 180
Java:
  type: programming
  tm_scope: source.java
  ace_mode: java
  codemirror_mode: clike
  codemirror_mime_type: text/x-java
  color: "#b07219"
  extensions:
  - ".java"
  - ".jav"
  language_id: 181
Java Properties:
  type: data
  color: "#2A6277"
  extensions:
  - ".properties"
  tm_scope: source.java-properties
  ace_mode: properties
  codemirror_mode: properties
  codemirror_mime_type: text/x-properties
  language_id: 519377561
Java Server Pages:
  type: programming
  color: "#2A6277"
  group: Java
  aliases:
  - jsp
  extensions:
  - ".jsp"
  tm_scope: text.html.jsp
  ace_mode: jsp
  codemirror_mode: htmlembedded
  codemirror_mime_type: application/x-jsp
  language_id: 182
JavaScript:
  type: programming
  tm_scope: source.js
  ace_mode: javascript
  codemirror_mode: javascript
  codemirror_mime_type: text/javascript
  color: "#f1e05a"
  aliases:
  - js
  - node
  extensions:
  - ".js"
  - "._js"
  - ".bones"
  - ".cjs"
  - ".es"
  - ".es6"
  - ".frag"
  - ".gs"
  - ".jake"
  - ".javascript"
  - ".jsb"
  - ".jscad"
  - ".jsfl"
  - ".jslib"
  - ".jsm"
  - ".jspre"
  - ".jss"
  - ".jsx"
  - ".mjs"
  - ".njs"
  - ".pac"
  - ".sjs"
  - ".ssjs"
  - ".xsjs"
  - ".xsjslib"
  filenames:
  - Jakefile
  interpreters:
  - chakra
  - d8
  - gjs
  - js
  - node
  - nodejs
  - qjs
  - rhino
  - v8
  - v8-shell
  language_id: 183
JavaScript+ERB:
  type: programming
  color: "#f1e05a"
  tm_scope: source.js
  group: JavaScript
  extensions:
  - ".js.erb"
  ace_mode: javascript
  codemirror_mode: javascript
  codemirror_mime_type: application/javascript
  language_id: 914318960
Jest Snapshot:
  type: data
  color: "#15c213"
  tm_scope: source.jest.snap
  extensions:
  - ".snap"
  ace_mode: javascript
  codemirror_mode: javascript
  codemirror_mime_type: application/javascript
  language_id: 774635084
Jinja:
  type: markup
  color: "#a52a22"
  aliases:
  - django
  - html+django
  - html+jinja
  - htmldjango
  extensions:
  - ".jinja"
  - ".j2"
  - ".jinja2"
  tm_scope: text.html.django
  ace_mode: django
  codemirror_mode: django
  codemirror_mime_type: text/x-django
  language_id: 147
Jison:
  type: programming
  color: "#56b3cb"
  group: Yacc
  extensions:
  - ".jison"
  tm_scope: source.jison
  ace_mode: text
  language_id: 284531423
Jison Lex:
  type: programming
  color: "#56b3cb"
  group: Lex
  extensions:
  - ".jisonlex"
  tm_scope: source.jisonlex
  ace_mode: text
  language_id: 406395330
Jolie:
  type: programming
  extensions:
  - ".ol"
  - ".iol"
  interpreters:
  - jolie
  color: "#843179"
  ace_mode: text
  tm_scope: source.jolie
  language_id: 998078858
Jsonnet:
  color: "#0064bd"
  type: programming
  ace_mode: text
  extensions:
  - ".jsonnet"
  - ".libsonnet"
  tm_scope: source.jsonnet
  language_id: 664885656
Julia:
  type: programming
  extensions:
  - ".jl"
  interpreters:
  - julia
  color: "#a270ba"
  tm_scope: source.julia
  ace_mode: julia
  codemirror_mode: julia
  codemirror_mime_type: text/x-julia
  language_id: 184
Jupyter Notebook:
  type: markup
  ace_mode: json
  codemirror_mode: javascript
  codemirror_mime_type: application/json
  tm_scope: source.json
  color: "#DA5B0B"
  extensions:
  - ".ipynb"
  filenames:
  - Notebook
  aliases:
  - IPython Notebook
  language_id: 185
KRL:
  type: programming
  color: "#28430A"
  extensions:
  - ".krl"
  tm_scope: none
  ace_mode: text
  language_id: 186
Kaitai Struct:
  type: programming
  aliases:
  - ksy
  ace_mode: yaml
  codemirror_mode: yaml
  codemirror_mime_type: text/x-yaml
  color: "#773b37"
  extensions:
  - ".ksy"
  tm_scope: source.yaml
  language_id: 818804755
KakouneScript:
  type: programming
  color: "#6f8042"
  tm_scope: source.kakscript
  aliases:
  - kak
  - kakscript
  extensions:
  - ".kak"
  filenames:
  - kakrc
  ace_mode: text
  language_id: 603336474
KiCad Layout:
  type: data
  color: "#2f4aab"
  aliases:
  - pcbnew
  extensions:
  - ".kicad_pcb"
  - ".kicad_mod"
  - ".kicad_wks"
  filenames:
  - fp-lib-table
  tm_scope: source.pcb.sexp
  ace_mode: lisp
  codemirror_mode: commonlisp
  codemirror_mime_type: text/x-common-lisp
  language_id: 187
KiCad Legacy Layout:
  type: data
  color: "#2f4aab"
  extensions:
  - ".brd"
  tm_scope: source.pcb.board
  ace_mode: text
  language_id: 140848857
KiCad Schematic:
  type: data
  color: "#2f4aab"
  aliases:
  - eeschema schematic
  extensions:
  - ".sch"
  tm_scope: source.pcb.schematic
  ace_mode: text
  language_id: 622447435
Kit:
  type: markup
  ace_mode: html
  codemirror_mode: htmlmixed
  codemirror_mime_type: text/html
  extensions:
  - ".kit"
  tm_scope: text.html.basic
  language_id: 188
Kotlin:
  type: programming
  color: "#A97BFF"
  extensions:
  - ".kt"
  - ".ktm"
  - ".kts"
  tm_scope: source.kotlin
  ace_mode: text
  codemirror_mode: clike
  codemirror_mime_type: text/x-kotlin
  language_id: 189
Kusto:
  type: data
  extensions:
  - ".csl"
  tm_scope: source.kusto
  ace_mode: text
  language_id: 225697190
LFE:
  type: programming
  color: "#4C3023"
  extensions:
  - ".lfe"
  tm_scope: source.lisp
  ace_mode: lisp
  codemirror_mode: commonlisp
  codemirror_mime_type: text/x-common-lisp
  language_id: 190
LLVM:
  type: programming
  extensions:
  - ".ll"
  tm_scope: source.llvm
  ace_mode: text
  color: "#185619"
  language_id: 191
LOLCODE:
  type: programming
  extensions:
  - ".lol"
  color: "#cc9900"
  tm_scope: none
  ace_mode: text
  language_id: 192
LSL:
  type: programming
  tm_scope: source.lsl
  ace_mode: lsl
  extensions:
  - ".lsl"
  - ".lslp"
  interpreters:
  - lsl
  color: "#3d9970"
  language_id: 193
LTspice Symbol:
  type: data
  extensions:
  - ".asy"
  tm_scope: source.ltspice.symbol
  ace_mode: text
  codemirror_mode: spreadsheet
  codemirror_mime_type: text/x-spreadsheet
  language_id: 1013566805
LabVIEW:
  type: programming
  color: "#fede06"
  extensions:
  - ".lvproj"
  - ".lvlib"
  tm_scope: text.xml
  ace_mode: xml
  codemirror_mode: xml
  codemirror_mime_type: text/xml
  language_id: 194
Lark:
  type: data
  color: "#2980B9"
  extensions:
  - ".lark"
  tm_scope: source.lark
  ace_mode: text
  codemirror_mode: ebnf
  codemirror_mime_type: text/x-ebnf
  language_id: 758480799
Lasso:
  type: programming
  color: "#999999"
  extensions:
  - ".lasso"
  - ".las"
  - ".lasso8"
  - ".lasso9"
  tm_scope: file.lasso
  aliases:
  - lassoscript
  ace_mode: text
  language_id: 195
Latte:
  type: markup
  color: "#f2a542"
  extensions:
  - ".latte"
  tm_scope: text.html.smarty
  ace_mode: smarty
  codemirror_mode: smarty
  codemirror_mime_type: text/x-smarty
  language_id: 196
Lean:
  type: programming
  extensions:
  - ".lean"
  - ".hlean"
  tm_scope: source.lean
  ace_mode: text
  language_id: 197
Less:
  type: markup
  color: "#1d365d"
  aliases:
  - less-css
  extensions:
  - ".less"
  tm_scope: source.css.less
  ace_mode: less
  codemirror_mode: css
  codemirror_mime_type: text/css
  language_id: 198
Lex:
  type: programming
  color: "#DBCA00"
  aliases:
  - flex
  extensions:
  - ".l"
  - ".lex"
  filenames:
  - Lexer.x
  - lexer.x
  tm_scope: source.lex
  ace_mode: text
  language_id: 199
LigoLANG:
  type: programming
  color: "#0e74ff"
  extensions:
  - ".ligo"
  tm_scope: source.ligo
  ace_mode: pascal
  codemirror_mode: pascal
  codemirror_mime_type: text/x-pascal
  group: LigoLANG
  language_id: 1040646257
LilyPond:
  type: programming
  color: "#9ccc7c"
  extensions:
  - ".ly"
  - ".ily"
  tm_scope: source.lilypond
  ace_mode: text
  language_id: 200
Limbo:
  type: programming
  extensions:
  - ".b"
  - ".m"
  tm_scope: none
  ace_mode: text
  language_id: 201
Linker Script:
  type: data
  extensions:
  - ".ld"
  - ".lds"
  - ".x"
  filenames:
  - ld.script
  tm_scope: none
  ace_mode: text
  language_id: 202
Linux Kernel Module:
  type: data
  extensions:
  - ".mod"
  tm_scope: none
  ace_mode: text
  language_id: 203
Liquid:
  type: markup
  color: "#67b8de"
  extensions:
  - ".liquid"
  tm_scope: text.html.liquid
  ace_mode: liquid
  language_id: 204
Literate Agda:
  type: programming
  color: "#315665"
  group: Agda
  extensions:
  - ".lagda"
  tm_scope: none
  ace_mode: text
  language_id: 205
Literate CoffeeScript:
  type: programming
  color: "#244776"
  tm_scope: source.litcoffee
  group: CoffeeScript
  ace_mode: text
  wrap: true
  aliases:
  - litcoffee
  extensions:
  - ".litcoffee"
  - ".coffee.md"
  language_id: 206
Literate Haskell:
  type: programming
  color: "#5e5086"
  group: Haskell
  aliases:
  - lhaskell
  - lhs
  extensions:
  - ".lhs"
  tm_scope: text.tex.latex.haskell
  ace_mode: text
  codemirror_mode: haskell-literate
  codemirror_mime_type: text/x-literate-haskell
  language_id: 207
LiveScript:
  type: programming
  color: "#499886"
  aliases:
  - live-script
  - ls
  extensions:
  - ".ls"
  - "._ls"
  filenames:
  - Slakefile
  tm_scope: source.livescript
  ace_mode: livescript
  codemirror_mode: livescript
  codemirror_mime_type: text/x-livescript
  language_id: 208
Logos:
  type: programming
  extensions:
  - ".xm"
  - ".x"
  - ".xi"
  ace_mode: text
  tm_scope: source.logos
  language_id: 209
Logtalk:
  type: programming
  color: "#295b9a"
  extensions:
  - ".lgt"
  - ".logtalk"
  tm_scope: source.logtalk
  ace_mode: text
  language_id: 210
LookML:
  type: programming
  ace_mode: yaml
  codemirror_mode: yaml
  codemirror_mime_type: text/x-yaml
  color: "#652B81"
  extensions:
  - ".lookml"
  - ".model.lkml"
  - ".view.lkml"
  tm_scope: source.yaml
  language_id: 211
LoomScript:
  type: programming
  extensions:
  - ".ls"
  tm_scope: source.loomscript
  ace_mode: text
  language_id: 212
Lua:
  type: programming
  tm_scope: source.lua
  ace_mode: lua
  codemirror_mode: lua
  codemirror_mime_type: text/x-lua
  color: "#000080"
  extensions:
  - ".lua"
  - ".fcgi"
  - ".nse"
  - ".p8"
  - ".pd_lua"
  - ".rbxs"
  - ".rockspec"
  - ".wlua"
  filenames:
  - ".luacheckrc"
  interpreters:
  - lua
  language_id: 213
M:
  type: programming
  aliases:
  - mumps
  extensions:
  - ".mumps"
  - ".m"
  ace_mode: text
  codemirror_mode: mumps
  codemirror_mime_type: text/x-mumps
  language_id: 214
  tm_scope: none
M4:
  type: programming
  extensions:
  - ".m4"
  - ".mc"
  tm_scope: source.m4
  ace_mode: text
  language_id: 215
M4Sugar:
  type: programming
  group: M4
  aliases:
  - autoconf
  extensions:
  - ".m4"
  filenames:
  - configure.ac
  tm_scope: source.m4
  ace_mode: text
  language_id: 216
MATLAB:
  type: programming
  color: "#e16737"
  aliases:
  - octave
  extensions:
  - ".matlab"
  - ".m"
  tm_scope: source.matlab
  ace_mode: matlab
  codemirror_mode: octave
  codemirror_mime_type: text/x-octave
  language_id: 225
MAXScript:
  type: programming
  color: "#00a6a6"
  extensions:
  - ".ms"
  - ".mcr"
  tm_scope: source.maxscript
  ace_mode: text
  language_id: 217
MLIR:
  type: programming
  color: "#5EC8DB"
  extensions:
  - ".mlir"
  tm_scope: source.mlir
  ace_mode: text
  language_id: 448253929
MQL4:
  type: programming
  color: "#62A8D6"
  extensions:
  - ".mq4"
  - ".mqh"
  tm_scope: source.mql5
  ace_mode: c_cpp
  language_id: 426
MQL5:
  type: programming
  color: "#4A76B8"
  extensions:
  - ".mq5"
  - ".mqh"
  tm_scope: source.mql5
  ace_mode: c_cpp
  language_id: 427
MTML:
  type: markup
  color: "#b7e1f4"
  extensions:
  - ".mtml"
  tm_scope: text.html.basic
  ace_mode: html
  codemirror_mode: htmlmixed
  codemirror_mime_type: text/html
  language_id: 218
MUF:
  type: programming
  group: Forth
  extensions:
  - ".muf"
  - ".m"
  tm_scope: none
  ace_mode: forth
  codemirror_mode: forth
  codemirror_mime_type: text/x-forth
  language_id: 219
Macaulay2:
  type: programming
  extensions:
  - ".m2"
  aliases:
  - m2
  interpreters:
  - M2
  ace_mode: text
  tm_scope: source.m2
  color: "#d8ffff"
  language_id: 34167825
Makefile:
  type: programming
  color: "#427819"
  aliases:
  - bsdmake
  - make
  - mf
  extensions:
  - ".mak"
  - ".d"
  - ".make"
  - ".makefile"
  - ".mk"
  - ".mkfile"
  filenames:
  - BSDmakefile
  - GNUmakefile
  - Kbuild
  - Makefile
  - Makefile.am
  - Makefile.boot
  - Makefile.frag
  - Makefile.in
  - Makefile.inc
  - Makefile.wat
  - makefile
  - makefile.sco
  - mkfile
  interpreters:
  - make
  tm_scope: source.makefile
  ace_mode: makefile
  codemirror_mode: cmake
  codemirror_mime_type: text/x-cmake
  language_id: 220
Mako:
  type: programming
  color: "#7e858d"
  extensions:
  - ".mako"
  - ".mao"
  tm_scope: text.html.mako
  ace_mode: text
  language_id: 221
Markdown:
  type: prose
  color: "#083fa1"
  aliases:
  - pandoc
  ace_mode: markdown
  codemirror_mode: gfm
  codemirror_mime_type: text/x-gfm
  wrap: true
  extensions:
  - ".md"
  - ".markdown"
  - ".mdown"
  - ".mdwn"
  - ".mdx"
  - ".mkd"
  - ".mkdn"
  - ".mkdown"
  - ".ronn"
  - ".scd"
  - ".workbook"
  filenames:
  - contents.lr
  tm_scope: source.gfm
  language_id: 222
Marko:
  type: markup
  color: "#42bff2"
  tm_scope: text.marko
  extensions:
  - ".marko"
  aliases:
  - markojs
  ace_mode: text
  codemirror_mode: htmlmixed
  codemirror_mime_type: text/html
  language_id: 932782397
Mask:
  type: markup
  color: "#f97732"
  ace_mode: mask
  extensions:
  - ".mask"
  tm_scope: source.mask
  language_id: 223
Mathematica:
  type: programming
  color: "#dd1100"
  extensions:
  - ".mathematica"
  - ".cdf"
  - ".m"
  - ".ma"
  - ".mt"
  - ".nb"
  - ".nbp"
  - ".wl"
  - ".wlt"
  aliases:
  - mma
  - wolfram
  - wolfram language
  - wolfram lang
  - wl
  tm_scope: source.mathematica
  ace_mode: text
  codemirror_mode: mathematica
  codemirror_mime_type: text/x-mathematica
  language_id: 224
Maven POM:
  type: data
  group: XML
  tm_scope: text.xml.pom
  filenames:
  - pom.xml
  ace_mode: xml
  codemirror_mode: xml
  codemirror_mime_type: text/xml
  language_id: 226
Max:
  type: programming
  color: "#c4a79c"
  aliases:
  - max/msp
  - maxmsp
  extensions:
  - ".maxpat"
  - ".maxhelp"
  - ".maxproj"
  - ".mxt"
  - ".pat"
  tm_scope: source.json
  ace_mode: json
  codemirror_mode: javascript
  codemirror_mime_type: application/json
  language_id: 227
Mercury:
  type: programming
  color: "#ff2b2b"
  ace_mode: prolog
  interpreters:
  - mmi
  extensions:
  - ".m"
  - ".moo"
  tm_scope: source.mercury
  language_id: 229
Meson:
  type: programming
  color: "#007800"
  filenames:
  - meson.build
  - meson_options.txt
  tm_scope: source.meson
  ace_mode: text
  language_id: 799141244
Metal:
  type: programming
  color: "#8f14e9"
  extensions:
  - ".metal"
  tm_scope: source.c++
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-c++src
  language_id: 230
Microsoft Developer Studio Project:
  type: data
  extensions:
  - ".dsp"
  tm_scope: none
  ace_mode: text
  language_id: 800983837
Microsoft Visual Studio Solution:
  type: data
  extensions:
  - ".sln"
  tm_scope: source.solution
  ace_mode: text
  language_id: 849523096
MiniD:
  type: programming
  extensions:
  - ".minid"
  tm_scope: none
  ace_mode: text
  language_id: 231
MiniYAML:
  type: data
  color: "#ff1111"
  tm_scope: source.miniyaml
  extensions:
  - ".yaml"
  ace_mode: yaml
  codemirror_mode: yaml
  codemirror_mime_type: text/x-yaml
  language_id: 4896465
Mint:
  type: programming
  extensions:
  - ".mint"
  ace_mode: text
  color: "#02b046"
  tm_scope: source.mint
  language_id: 968740319
Mirah:
  type: programming
  color: "#c7a938"
  extensions:
  - ".druby"
  - ".duby"
  - ".mirah"
  tm_scope: source.ruby
  ace_mode: ruby
  codemirror_mode: ruby
  codemirror_mime_type: text/x-ruby
  language_id: 232
Modelica:
  type: programming
  color: "#de1d31"
  extensions:
  - ".mo"
  tm_scope: source.modelica
  ace_mode: text
  codemirror_mode: modelica
  codemirror_mime_type: text/x-modelica
  language_id: 233
Modula-2:
  type: programming
  color: "#10253f"
  extensions:
  - ".mod"
  tm_scope: source.modula2
  ace_mode: text
  language_id: 234
Modula-3:
  type: programming
  extensions:
  - ".i3"
  - ".ig"
  - ".m3"
  - ".mg"
  color: "#223388"
  ace_mode: text
  tm_scope: source.modula-3
  language_id: 564743864
Module Management System:
  type: programming
  extensions:
  - ".mms"
  - ".mmk"
  filenames:
  - descrip.mmk
  - descrip.mms
  tm_scope: none
  ace_mode: text
  language_id: 235
Monkey:
  type: programming
  extensions:
  - ".monkey"
  - ".monkey2"
  ace_mode: text
  tm_scope: source.monkey
  language_id: 236
Monkey C:
  type: programming
  color: "#8D6747"
  extensions:
  - ".mc"
  tm_scope: source.mc
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-csrc
  language_id: 231751931
Moocode:
  type: programming
  extensions:
  - ".moo"
  tm_scope: none
  ace_mode: text
  language_id: 237
MoonScript:
  type: programming
  color: "#ff4585"
  extensions:
  - ".moon"
  interpreters:
  - moon
  tm_scope: source.moonscript
  ace_mode: text
  language_id: 238
Motoko:
  type: programming
  color: "#fbb03b"
  extensions:
  - ".mo"
  tm_scope: source.mo
  ace_mode: text
  language_id: 202937027
Motorola 68K Assembly:
  type: programming
  color: "#005daa"
  group: Assembly
  aliases:
  - m68k
  extensions:
  - ".asm"
  - ".i"
  - ".inc"
  - ".s"
  - ".x68"
  tm_scope: source.m68k
  ace_mode: assembly_x86
  language_id: 477582706
Muse:
  type: prose
  extensions:
  - ".muse"
  tm_scope: text.muse
  ace_mode: text
  wrap: true
  language_id: 474864066
  aliases:
  - amusewiki
  - emacs muse
Mustache:
  type: markup
  color: "#724b3b"
  extensions:
  - ".mustache"
  tm_scope: text.html.smarty
  ace_mode: smarty
  codemirror_mode: smarty
  codemirror_mime_type: text/x-smarty
  language_id: 638334590
Myghty:
  type: programming
  extensions:
  - ".myt"
  tm_scope: none
  ace_mode: text
  language_id: 239
NASL:
  type: programming
  extensions:
  - ".nasl"
  - ".inc"
  tm_scope: source.nasl
  ace_mode: text
  language_id: 171666519
NCL:
  type: programming
  color: "#28431f"
  extensions:
  - ".ncl"
  tm_scope: source.ncl
  ace_mode: text
  language_id: 240
NEON:
  type: data
  extensions:
  - ".neon"
  tm_scope: source.neon
  ace_mode: text
  aliases:
  - nette object notation
  - ne-on
  language_id: 481192983
NL:
  type: data
  extensions:
  - ".nl"
  tm_scope: none
  ace_mode: text
  language_id: 241
NPM Config:
  type: data
  color: "#cb3837"
  group: INI
  aliases:
  - npmrc
  filenames:
  - ".npmrc"
  tm_scope: source.ini.npmrc
  ace_mode: text
  language_id: 685022663
NSIS:
  type: programming
  extensions:
  - ".nsi"
  - ".nsh"
  tm_scope: source.nsis
  ace_mode: text
  codemirror_mode: nsis
  codemirror_mime_type: text/x-nsis
  language_id: 242
NWScript:
  type: programming
  color: "#111522"
  extensions:
  - ".nss"
  tm_scope: source.c.nwscript
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-csrc
  language_id: 731233819
Nearley:
  type: programming
  ace_mode: text
  color: "#990000"
  extensions:
  - ".ne"
  - ".nearley"
  tm_scope: source.ne
  language_id: 521429430
Nemerle:
  type: programming
  color: "#3d3c6e"
  extensions:
  - ".n"
  tm_scope: source.nemerle
  ace_mode: text
  language_id: 243
NetLinx:
  type: programming
  color: "#0aa0ff"
  extensions:
  - ".axs"
  - ".axi"
  tm_scope: source.netlinx
  ace_mode: text
  language_id: 244
NetLinx+ERB:
  type: programming
  color: "#747faa"
  extensions:
  - ".axs.erb"
  - ".axi.erb"
  tm_scope: source.netlinx.erb
  ace_mode: text
  language_id: 245
NetLogo:
  type: programming
  color: "#ff6375"
  extensions:
  - ".nlogo"
  tm_scope: source.lisp
  ace_mode: lisp
  codemirror_mode: commonlisp
  codemirror_mime_type: text/x-common-lisp
  language_id: 246
NewLisp:
  type: programming
  color: "#87AED7"
  extensions:
  - ".nl"
  - ".lisp"
  - ".lsp"
  interpreters:
  - newlisp
  tm_scope: source.lisp
  ace_mode: lisp
  codemirror_mode: commonlisp
  codemirror_mime_type: text/x-common-lisp
  language_id: 247
Nextflow:
  type: programming
  ace_mode: groovy
  tm_scope: source.nextflow
  color: "#3ac486"
  extensions:
  - ".nf"
  filenames:
  - nextflow.config
  interpreters:
  - nextflow
  language_id: 506780613
Nginx:
  type: data
  color: "#009639"
  extensions:
  - ".nginx"
  - ".nginxconf"
  - ".vhost"
  filenames:
  - nginx.conf
  tm_scope: source.nginx
  aliases:
  - nginx configuration file
  ace_mode: text
  codemirror_mode: nginx
  codemirror_mime_type: text/x-nginx-conf
  language_id: 248
Nim:
  type: programming
  color: "#ffc200"
  extensions:
  - ".nim"
  - ".nim.cfg"
  - ".nimble"
  - ".nimrod"
  - ".nims"
  filenames:
  - nim.cfg
  ace_mode: text
  tm_scope: source.nim
  language_id: 249
Ninja:
  type: data
  tm_scope: source.ninja
  extensions:
  - ".ninja"
  ace_mode: text
  language_id: 250
Nit:
  type: programming
  color: "#009917"
  extensions:
  - ".nit"
  tm_scope: source.nit
  ace_mode: text
  language_id: 251
Nix:
  type: programming
  color: "#7e7eff"
  extensions:
  - ".nix"
  aliases:
  - nixos
  tm_scope: source.nix
  ace_mode: nix
  language_id: 252
Nu:
  type: programming
  color: "#c9df40"
  aliases:
  - nush
  extensions:
  - ".nu"
  filenames:
  - Nukefile
  tm_scope: source.nu
  ace_mode: scheme
  codemirror_mode: scheme
  codemirror_mime_type: text/x-scheme
  interpreters:
  - nush
  language_id: 253
NumPy:
  type: programming
  color: "#9C8AF9"
  group: Python
  extensions:
  - ".numpy"
  - ".numpyw"
  - ".numsc"
  tm_scope: none
  ace_mode: text
  codemirror_mode: python
  codemirror_mime_type: text/x-python
  language_id: 254
Nunjucks:
  type: markup
  color: "#3d8137"
  extensions:
  - ".njk"
  aliases:
  - njk
  tm_scope: text.html.nunjucks
  ace_mode: nunjucks
  language_id: 461856962
OCaml:
  type: programming
  ace_mode: ocaml
  codemirror_mode: mllike
  codemirror_mime_type: text/x-ocaml
  color: "#3be133"
  extensions:
  - ".ml"
  - ".eliom"
  - ".eliomi"
  - ".ml4"
  - ".mli"
  - ".mll"
  - ".mly"
  interpreters:
  - ocaml
  - ocamlrun
  - ocamlscript
  tm_scope: source.ocaml
  language_id: 255
ObjDump:
  type: data
  extensions:
  - ".objdump"
  tm_scope: objdump.x86asm
  ace_mode: assembly_x86
  language_id: 256
Object Data Instance Notation:
  type: data
  extensions:
  - ".odin"
  tm_scope: source.odin-ehr
  ace_mode: text
  language_id: 985227236
ObjectScript:
  type: programming
  extensions:
  - ".cls"
  language_id: 202735509
  tm_scope: source.objectscript
  color: "#424893"
  ace_mode: text
Objective-C:
  type: programming
  tm_scope: source.objc
  color: "#438eff"
  aliases:
  - obj-c
  - objc
  - objectivec
  extensions:
  - ".m"
  - ".h"
  ace_mode: objectivec
  codemirror_mode: clike
  codemirror_mime_type: text/x-objectivec
  language_id: 257
Objective-C++:
  type: programming
  tm_scope: source.objc++
  color: "#6866fb"
  aliases:
  - obj-c++
  - objc++
  - objectivec++
  extensions:
  - ".mm"
  ace_mode: objectivec
  codemirror_mode: clike
  codemirror_mime_type: text/x-objectivec
  language_id: 258
Objective-J:
  type: programming
  color: "#ff0c5a"
  aliases:
  - obj-j
  - objectivej
  - objj
  extensions:
  - ".j"
  - ".sj"
  tm_scope: source.js.objj
  ace_mode: text
  language_id: 259
Odin:
  type: programming
  color: "#60AFFE"
  aliases:
  - odinlang
  - odin-lang
  extensions:
  - ".odin"
  tm_scope: source.odin
  ace_mode: text
  language_id: 889244082
Omgrofl:
  type: programming
  extensions:
  - ".omgrofl"
  color: "#cabbff"
  tm_scope: none
  ace_mode: text
  language_id: 260
Opa:
  type: programming
  extensions:
  - ".opa"
  tm_scope: source.opa
  ace_mode: text
  language_id: 261
Opal:
  type: programming
  color: "#f7ede0"
  extensions:
  - ".opal"
  tm_scope: source.opal
  ace_mode: text
  language_id: 262
Open Policy Agent:
  type: programming
  color: "#7d9199"
  ace_mode: text
  extensions:
  - ".rego"
  language_id: 840483232
  tm_scope: source.rego
OpenCL:
  type: programming
  color: "#ed2e2d"
  group: C
  extensions:
  - ".cl"
  - ".opencl"
  tm_scope: source.c
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-csrc
  language_id: 263
OpenEdge ABL:
  type: programming
  color: "#5ce600"
  aliases:
  - progress
  - openedge
  - abl
  extensions:
  - ".p"
  - ".cls"
  - ".w"
  tm_scope: source.abl
  ace_mode: text
  language_id: 264
OpenQASM:
  type: programming
  extensions:
  - ".qasm"
  color: "#AA70FF"
  tm_scope: source.qasm
  ace_mode: text
  language_id: 153739399
OpenRC runscript:
  type: programming
  group: Shell
  aliases:
  - openrc
  interpreters:
  - openrc-run
  tm_scope: source.shell
  ace_mode: sh
  codemirror_mode: shell
  codemirror_mime_type: text/x-sh
  language_id: 265
OpenSCAD:
  type: programming
  color: "#e5cd45"
  extensions:
  - ".scad"
  tm_scope: source.scad
  ace_mode: scad
  language_id: 266
OpenStep Property List:
  type: data
  extensions:
  - ".plist"
  - ".glyphs"
  tm_scope: source.plist
  ace_mode: text
  language_id: 598917541
OpenType Feature File:
  type: data
  aliases:
  - AFDKO
  extensions:
  - ".fea"
  tm_scope: source.opentype
  ace_mode: text
  language_id: 374317347
Org:
  type: prose
  color: "#77aa99"
  wrap: true
  extensions:
  - ".org"
  tm_scope: none
  ace_mode: text
  language_id: 267
Ox:
  type: programming
  extensions:
  - ".ox"
  - ".oxh"
  - ".oxo"
  tm_scope: source.ox
  ace_mode: text
  language_id: 268
Oxygene:
  type: programming
  color: "#cdd0e3"
  extensions:
  - ".oxygene"
  tm_scope: none
  ace_mode: text
  language_id: 269
Oz:
  type: programming
  color: "#fab738"
  extensions:
  - ".oz"
  tm_scope: source.oz
  ace_mode: text
  codemirror_mode: oz
  codemirror_mime_type: text/x-oz
  language_id: 270
P4:
  type: programming
  color: "#7055b5"
  extensions:
  - ".p4"
  tm_scope: source.p4
  ace_mode: text
  language_id: 348895984
PEG.js:
  type: programming
  color: "#234d6b"
  extensions:
  - ".pegjs"
  tm_scope: source.pegjs
  ace_mode: javascript
  codemirror_mode: javascript
  codemirror_mime_type: text/javascript
  language_id: 81442128
PHP:
  type: programming
  tm_scope: text.html.php
  ace_mode: php
  codemirror_mode: php
  codemirror_mime_type: application/x-httpd-php
  color: "#4F5D95"
  extensions:
  - ".php"
  - ".aw"
  - ".ctp"
  - ".fcgi"
  - ".inc"
  - ".php3"
  - ".php4"
  - ".php5"
  - ".phps"
  - ".phpt"
  filenames:
  - ".php"
  - ".php_cs"
  - ".php_cs.dist"
  - Phakefile
  interpreters:
  - php
  aliases:
  - inc
  language_id: 272
PLSQL:
  type: programming
  ace_mode: sql
  codemirror_mode: sql
  codemirror_mime_type: text/x-plsql
  tm_scope: none
  color: "#dad8d8"
  extensions:
  - ".pls"
  - ".bdy"
  - ".ddl"
  - ".fnc"
  - ".pck"
  - ".pkb"
  - ".pks"
  - ".plb"
  - ".plsql"
  - ".prc"
  - ".spc"
  - ".sql"
  - ".tpb"
  - ".tps"
  - ".trg"
  - ".vw"
  language_id: 273
PLpgSQL:
  type: programming
  color: "#336790"
  ace_mode: pgsql
  codemirror_mode: sql
  codemirror_mime_type: text/x-sql
  tm_scope: source.sql
  extensions:
  - ".pgsql"
  - ".sql"
  language_id: 274
POV-Ray SDL:
  type: programming
  color: "#6bac65"
  aliases:
  - pov-ray
  - povray
  extensions:
  - ".pov"
  - ".inc"
  tm_scope: source.pov-ray sdl
  ace_mode: text
  language_id: 275
Pan:
  type: programming
  color: "#cc0000"
  extensions:
  - ".pan"
  tm_scope: source.pan
  ace_mode: text
  language_id: 276
Papyrus:
  type: programming
  color: "#6600cc"
  extensions:
  - ".psc"
  tm_scope: source.papyrus.skyrim
  ace_mode: text
  language_id: 277
Parrot:
  type: programming
  color: "#f3ca0a"
  extensions:
  - ".parrot"
  tm_scope: none
  ace_mode: text
  language_id: 278
Parrot Assembly:
  group: Parrot
  type: programming
  aliases:
  - pasm
  extensions:
  - ".pasm"
  interpreters:
  - parrot
  tm_scope: none
  ace_mode: text
  language_id: 279
Parrot Internal Representation:
  group: Parrot
  tm_scope: source.parrot.pir
  type: programming
  aliases:
  - pir
  extensions:
  - ".pir"
  interpreters:
  - parrot
  ace_mode: text
  language_id: 280
Pascal:
  type: programming
  color: "#E3F171"
  aliases:
  - delphi
  - objectpascal
  extensions:
  - ".pas"
  - ".dfm"
  - ".dpr"
  - ".inc"
  - ".lpr"
  - ".pascal"
  - ".pp"
  interpreters:
  - instantfpc
  tm_scope: source.pascal
  ace_mode: pascal
  codemirror_mode: pascal
  codemirror_mime_type: text/x-pascal
  language_id: 281
Pawn:
  type: programming
  color: "#dbb284"
  extensions:
  - ".pwn"
  - ".inc"
  - ".sma"
  tm_scope: source.pawn
  ace_mode: text
  language_id: 271
Pep8:
  type: programming
  color: "#C76F5B"
  extensions:
  - ".pep"
  ace_mode: text
  tm_scope: source.pep8
  language_id: 840372442
Perl:
  type: programming
  tm_scope: source.perl
  ace_mode: perl
  codemirror_mode: perl
  codemirror_mime_type: text/x-perl
  color: "#0298c3"
  extensions:
  - ".pl"
  - ".al"
  - ".cgi"
  - ".fcgi"
  - ".perl"
  - ".ph"
  - ".plx"
  - ".pm"
  - ".psgi"
  - ".t"
  filenames:
  - Makefile.PL
  - Rexfile
  - ack
  - cpanfile
  interpreters:
  - cperl
  - perl
  aliases:
  - cperl
  language_id: 282
Pic:
  type: markup
  group: Roff
  tm_scope: source.pic
  extensions:
  - ".pic"
  - ".chem"
  ace_mode: text
  codemirror_mode: troff
  codemirror_mime_type: text/troff
  language_id: 425
Pickle:
  type: data
  extensions:
  - ".pkl"
  tm_scope: none
  ace_mode: text
  language_id: 284
PicoLisp:
  type: programming
  color: "#6067af"
  extensions:
  - ".l"
  interpreters:
  - picolisp
  - pil
  tm_scope: source.lisp
  ace_mode: lisp
  language_id: 285
PigLatin:
  type: programming
  color: "#fcd7de"
  extensions:
  - ".pig"
  tm_scope: source.pig_latin
  ace_mode: text
  language_id: 286
Pike:
  type: programming
  color: "#005390"
  extensions:
  - ".pike"
  - ".pmod"
  interpreters:
  - pike
  tm_scope: source.pike
  ace_mode: text
  language_id: 287
PlantUML:
  type: data
  extensions:
  - ".puml"
  - ".iuml"
  - ".plantuml"
  tm_scope: source.wsd
  ace_mode: text
  language_id: 833504686
Pod:
  type: prose
  ace_mode: perl
  codemirror_mode: perl
  codemirror_mime_type: text/x-perl
  wrap: true
  extensions:
  - ".pod"
  interpreters:
  - perl
  tm_scope: none
  language_id: 288
Pod 6:
  type: prose
  ace_mode: perl
  tm_scope: source.raku
  wrap: true
  extensions:
  - ".pod"
  - ".pod6"
  interpreters:
  - perl6
  language_id: 155357471
PogoScript:
  type: programming
  color: "#d80074"
  extensions:
  - ".pogo"
  tm_scope: source.pogoscript
  ace_mode: text
  language_id: 289
Pony:
  type: programming
  extensions:
  - ".pony"
  tm_scope: source.pony
  ace_mode: text
  language_id: 290
PostCSS:
  type: markup
  color: "#dc3a0c"
  tm_scope: source.postcss
  group: CSS
  extensions:
  - ".pcss"
  - ".postcss"
  ace_mode: text
  language_id: 262764437
PostScript:
  type: markup
  color: "#da291c"
  extensions:
  - ".ps"
  - ".eps"
  - ".epsi"
  - ".pfa"
  tm_scope: source.postscript
  aliases:
  - postscr
  ace_mode: text
  language_id: 291
PowerBuilder:
  type: programming
  color: "#8f0f8d"
  extensions:
  - ".pbt"
  - ".sra"
  - ".sru"
  - ".srw"
  tm_scope: none
  ace_mode: text
  language_id: 292
PowerShell:
  type: programming
  color: "#012456"
  tm_scope: source.powershell
  ace_mode: powershell
  codemirror_mode: powershell
  codemirror_mime_type: application/x-powershell
  aliases:
  - posh
  - pwsh
  extensions:
  - ".ps1"
  - ".psd1"
  - ".psm1"
  interpreters:
  - pwsh
  language_id: 293
Prisma:
  type: data
  color: "#0c344b"
  extensions:
  - ".prisma"
  tm_scope: source.prisma
  ace_mode: text
  language_id: 499933428
Processing:
  type: programming
  color: "#0096D8"
  extensions:
  - ".pde"
  tm_scope: source.processing
  ace_mode: text
  language_id: 294
Procfile:
  type: programming
  color: "#3B2F63"
  filenames:
  - Procfile
  tm_scope: source.procfile
  ace_mode: batchfile
  language_id: 305313959
Proguard:
  type: data
  extensions:
  - ".pro"
  tm_scope: none
  ace_mode: text
  language_id: 716513858
Prolog:
  type: programming
  color: "#74283c"
  extensions:
  - ".pl"
  - ".pro"
  - ".prolog"
  - ".yap"
  interpreters:
  - swipl
  - yap
  tm_scope: source.prolog
  ace_mode: prolog
  language_id: 295
Promela:
  type: programming
  color: "#de0000"
  tm_scope: source.promela
  ace_mode: text
  extensions:
  - ".pml"
  language_id: 441858312
Propeller Spin:
  type: programming
  color: "#7fa2a7"
  extensions:
  - ".spin"
  tm_scope: source.spin
  ace_mode: text
  language_id: 296
Protocol Buffer:
  type: data
  aliases:
  - protobuf
  - Protocol Buffers
  extensions:
  - ".proto"
  tm_scope: source.proto
  ace_mode: protobuf
  codemirror_mode: protobuf
  codemirror_mime_type: text/x-protobuf
  language_id: 297
Protocol Buffer Text Format:
  type: data
  aliases:
  - text proto
  - protobuf text format
  extensions:
  - ".textproto"
  - ".pbt"
  - ".pbtxt"
  tm_scope: source.textproto
  ace_mode: text
  language_id: 436568854
Public Key:
  type: data
  extensions:
  - ".asc"
  - ".pub"
  tm_scope: none
  ace_mode: text
  codemirror_mode: asciiarmor
  codemirror_mime_type: application/pgp
  language_id: 298
Pug:
  type: markup
  color: "#a86454"
  extensions:
  - ".jade"
  - ".pug"
  tm_scope: text.jade
  ace_mode: jade
  codemirror_mode: pug
  codemirror_mime_type: text/x-pug
  language_id: 179
Puppet:
  type: programming
  color: "#302B6D"
  extensions:
  - ".pp"
  filenames:
  - Modulefile
  ace_mode: text
  codemirror_mode: puppet
  codemirror_mime_type: text/x-puppet
  tm_scope: source.puppet
  language_id: 299
Pure Data:
  type: data
  extensions:
  - ".pd"
  tm_scope: none
  ace_mode: text
  language_id: 300
PureBasic:
  type: programming
  color: "#5a6986"
  extensions:
  - ".pb"
  - ".pbi"
  tm_scope: none
  ace_mode: text
  language_id: 301
PureScript:
  type: programming
  color: "#1D222D"
  extensions:
  - ".purs"
  tm_scope: source.purescript
  ace_mode: haskell
  codemirror_mode: haskell
  codemirror_mime_type: text/x-haskell
  language_id: 302
Python:
  type: programming
  tm_scope: source.python
  ace_mode: python
  codemirror_mode: python
  codemirror_mime_type: text/x-python
  color: "#3572A5"
  extensions:
  - ".py"
  - ".cgi"
  - ".fcgi"
  - ".gyp"
  - ".gypi"
  - ".lmi"
  - ".py3"
  - ".pyde"
  - ".pyi"
  - ".pyp"
  - ".pyt"
  - ".pyw"
  - ".rpy"
  - ".smk"
  - ".spec"
  - ".tac"
  - ".wsgi"
  - ".xpy"
  filenames:
  - ".gclient"
  - DEPS
  - SConscript
  - SConstruct
  - Snakefile
  - wscript
  interpreters:
  - python
  - python2
  - python3
  aliases:
  - python3
  - rusthon
  language_id: 303
Python console:
  type: programming
  color: "#3572A5"
  group: Python
  aliases:
  - pycon
  tm_scope: text.python.console
  ace_mode: text
  language_id: 428
Python traceback:
  type: data
  color: "#3572A5"
  group: Python
  extensions:
  - ".pytb"
  tm_scope: text.python.traceback
  ace_mode: text
  language_id: 304
Q#:
  type: programming
  extensions:
  - ".qs"
  aliases:
  - qsharp
  color: "#fed659"
  ace_mode: text
  tm_scope: source.qsharp
  language_id: 697448245
QML:
  type: programming
  color: "#44a51c"
  extensions:
  - ".qml"
  - ".qbs"
  tm_scope: source.qml
  ace_mode: text
  language_id: 305
QMake:
  type: programming
  extensions:
  - ".pro"
  - ".pri"
  interpreters:
  - qmake
  tm_scope: source.qmake
  ace_mode: text
  language_id: 306
Qt Script:
  type: programming
  ace_mode: javascript
  codemirror_mode: javascript
  codemirror_mime_type: text/javascript
  extensions:
  - ".qs"
  filenames:
  - installscript.qs
  - toolchain_installscript.qs
  color: "#00b841"
  tm_scope: source.js
  language_id: 558193693
Quake:
  type: programming
  filenames:
  - m3makefile
  - m3overrides
  color: "#882233"
  ace_mode: text
  tm_scope: source.quake
  language_id: 375265331
R:
  type: programming
  color: "#198CE7"
  aliases:
  - R
  - Rscript
  - splus
  extensions:
  - ".r"
  - ".rd"
  - ".rsx"
  filenames:
  - ".Rprofile"
  - expr-dist
  interpreters:
  - Rscript
  tm_scope: source.r
  ace_mode: r
  codemirror_mode: r
  codemirror_mime_type: text/x-rsrc
  language_id: 307
RAML:
  type: markup
  ace_mode: yaml
  codemirror_mode: yaml
  codemirror_mime_type: text/x-yaml
  tm_scope: source.yaml
  color: "#77d9fb"
  extensions:
  - ".raml"
  language_id: 308
RDoc:
  type: prose
  color: "#701516"
  ace_mode: rdoc
  wrap: true
  extensions:
  - ".rdoc"
  tm_scope: text.rdoc
  language_id: 309
REALbasic:
  type: programming
  extensions:
  - ".rbbas"
  - ".rbfrm"
  - ".rbmnu"
  - ".rbres"
  - ".rbtbar"
  - ".rbuistate"
  tm_scope: source.vbnet
  ace_mode: text
  language_id: 310
REXX:
  type: programming
  color: "#d90e09"
  aliases:
  - arexx
  extensions:
  - ".rexx"
  - ".pprx"
  - ".rex"
  interpreters:
  - regina
  - rexx
  tm_scope: source.rexx
  ace_mode: text
  language_id: 311
RMarkdown:
  type: prose
  color: "#198ce7"
  wrap: true
  ace_mode: markdown
  codemirror_mode: gfm
  codemirror_mime_type: text/x-gfm
  extensions:
  - ".qmd"
  - ".rmd"
  tm_scope: source.gfm
  language_id: 313
RPC:
  type: programming
  aliases:
  - rpcgen
  - oncrpc
  - xdr
  ace_mode: c_cpp
  extensions:
  - ".x"
  tm_scope: source.c
  language_id: 1031374237
RPGLE:
  type: programming
  ace_mode: text
  color: "#2BDE21"
  aliases:
  - ile rpg
  - sqlrpgle
  extensions:
  - ".rpgle"
  - ".sqlrpgle"
  tm_scope: source.rpgle
  language_id: 609977990
RPM Spec:
  type: data
  tm_scope: source.rpm-spec
  extensions:
  - ".spec"
  aliases:
  - specfile
  ace_mode: text
  codemirror_mode: rpm
  codemirror_mime_type: text/x-rpm-spec
  language_id: 314
RUNOFF:
  type: markup
  color: "#665a4e"
  extensions:
  - ".rnh"
  - ".rno"
  wrap: true
  tm_scope: text.runoff
  ace_mode: text
  language_id: 315
Racket:
  type: programming
  color: "#3c5caa"
  extensions:
  - ".rkt"
  - ".rktd"
  - ".rktl"
  - ".scrbl"
  interpreters:
  - racket
  tm_scope: source.racket
  ace_mode: lisp
  language_id: 316
Ragel:
  type: programming
  color: "#9d5200"
  extensions:
  - ".rl"
  aliases:
  - ragel-rb
  - ragel-ruby
  tm_scope: none
  ace_mode: text
  language_id: 317
Raku:
  type: programming
  color: "#0000fb"
  extensions:
  - ".6pl"
  - ".6pm"
  - ".nqp"
  - ".p6"
  - ".p6l"
  - ".p6m"
  - ".pl"
  - ".pl6"
  - ".pm"
  - ".pm6"
  - ".raku"
  - ".rakumod"
  - ".t"
  interpreters:
  - perl6
  - raku
  - rakudo
  aliases:
  - perl6
  - perl-6
  tm_scope: source.raku
  ace_mode: perl
  codemirror_mode: perl
  codemirror_mime_type: text/x-perl
  language_id: 283
Rascal:
  type: programming
  color: "#fffaa0"
  extensions:
  - ".rsc"
  tm_scope: source.rascal
  ace_mode: text
  language_id: 173616037
Raw token data:
  type: data
  aliases:
  - raw
  extensions:
  - ".raw"
  tm_scope: none
  ace_mode: text
  language_id: 318
ReScript:
  type: programming
  color: "#ed5051"
  ace_mode: rust
  codemirror_mode: rust
  codemirror_mime_type: text/x-rustsrc
  extensions:
  - ".res"
  interpreters:
  - ocaml
  tm_scope: source.rescript
  language_id: 501875647
Readline Config:
  type: data
  group: INI
  aliases:
  - inputrc
  - readline
  filenames:
  - ".inputrc"
  - inputrc
  tm_scope: source.inputrc
  ace_mode: text
  language_id: 538732839
Reason:
  type: programming
  color: "#ff5847"
  ace_mode: rust
  codemirror_mode: rust
  codemirror_mime_type: text/x-rustsrc
  extensions:
  - ".re"
  - ".rei"
  tm_scope: source.reason
  language_id: 869538413
ReasonLIGO:
  type: programming
  color: "#ff5847"
  ace_mode: rust
  codemirror_mode: rust
  codemirror_mime_type: text/x-rustsrc
  group: LigoLANG
  extensions:
  - ".religo"
  tm_scope: source.religo
  language_id: 319002153
Rebol:
  type: programming
  color: "#358a5b"
  extensions:
  - ".reb"
  - ".r"
  - ".r2"
  - ".r3"
  - ".rebol"
  ace_mode: text
  tm_scope: source.rebol
  language_id: 319
Record Jar:
  type: data
  filenames:
  - language-subtag-registry.txt
  tm_scope: source.record-jar
  codemirror_mode: properties
  codemirror_mime_type: text/x-properties
  ace_mode: text
  color: "#0673ba"
  language_id: 865765202
Red:
  type: programming
  color: "#f50000"
  extensions:
  - ".red"
  - ".reds"
  aliases:
  - red/system
  tm_scope: source.red
  ace_mode: text
  language_id: 320
Redcode:
  type: programming
  extensions:
  - ".cw"
  tm_scope: none
  ace_mode: text
  language_id: 321
Redirect Rules:
  type: data
  aliases:
  - redirects
  filenames:
  - _redirects
  tm_scope: source.redirects
  ace_mode: text
  language_id: 1020148948
Regular Expression:
  type: data
  color: "#009a00"
  extensions:
  - ".regexp"
  - ".regex"
  aliases:
  - regexp
  - regex
  ace_mode: text
  tm_scope: source.regexp
  language_id: 363378884
Ren'Py:
  type: programming
  aliases:
  - renpy
  color: "#ff7f7f"
  extensions:
  - ".rpy"
  tm_scope: source.renpy
  ace_mode: python
  language_id: 322
RenderScript:
  type: programming
  extensions:
  - ".rs"
  - ".rsh"
  tm_scope: none
  ace_mode: text
  language_id: 323
Rich Text Format:
  type: markup
  extensions:
  - ".rtf"
  tm_scope: text.rtf
  ace_mode: text
  language_id: 51601661
Ring:
  type: programming
  color: "#2D54CB"
  extensions:
  - ".ring"
  tm_scope: source.ring
  ace_mode: text
  language_id: 431
Riot:
  type: markup
  color: "#A71E49"
  ace_mode: html
  extensions:
  - ".riot"
  tm_scope: text.html.riot
  language_id: 878396783
RobotFramework:
  type: programming
  color: "#00c0b5"
  extensions:
  - ".robot"
  tm_scope: text.robot
  ace_mode: text
  language_id: 324
Roff:
  type: markup
  color: "#ecdebe"
  extensions:
  - ".roff"
  - ".1"
  - ".1in"
  - ".1m"
  - ".1x"
  - ".2"
  - ".3"
  - ".3in"
  - ".3m"
  - ".3p"
  - ".3pm"
  - ".3qt"
  - ".3x"
  - ".4"
  - ".5"
  - ".6"
  - ".7"
  - ".8"
  - ".9"
  - ".l"
  - ".man"
  - ".mdoc"
  - ".me"
  - ".ms"
  - ".n"
  - ".nr"
  - ".rno"
  - ".tmac"
  filenames:
  - eqnrc
  - mmn
  - mmt
  - troffrc
  - troffrc-end
  tm_scope: text.roff
  aliases:
  - groff
  - man
  - manpage
  - man page
  - man-page
  - mdoc
  - nroff
  - troff
  wrap: true
  ace_mode: text
  codemirror_mode: troff
  codemirror_mime_type: text/troff
  language_id: 141
Roff Manpage:
  type: markup
  color: "#ecdebe"
  group: Roff
  extensions:
  - ".1"
  - ".1in"
  - ".1m"
  - ".1x"
  - ".2"
  - ".3"
  - ".3in"
  - ".3m"
  - ".3p"
  - ".3pm"
  - ".3qt"
  - ".3x"
  - ".4"
  - ".5"
  - ".6"
  - ".7"
  - ".8"
  - ".9"
  - ".man"
  - ".mdoc"
  wrap: true
  tm_scope: text.roff
  ace_mode: text
  codemirror_mode: troff
  codemirror_mime_type: text/troff
  language_id: 612669833
Rouge:
  type: programming
  ace_mode: clojure
  codemirror_mode: clojure
  codemirror_mime_type: text/x-clojure
  color: "#cc0088"
  extensions:
  - ".rg"
  tm_scope: source.clojure
  language_id: 325
Ruby:
  type: programming
  tm_scope: source.ruby
  ace_mode: ruby
  codemirror_mode: ruby
  codemirror_mime_type: text/x-ruby
  color: "#701516"
  aliases:
  - jruby
  - macruby
  - rake
  - rb
  - rbx
  extensions:
  - ".rb"
  - ".builder"
  - ".eye"
  - ".fcgi"
  - ".gemspec"
  - ".god"
  - ".jbuilder"
  - ".mspec"
  - ".pluginspec"
  - ".podspec"
  - ".prawn"
  - ".rabl"
  - ".rake"
  - ".rbi"
  - ".rbuild"
  - ".rbw"
  - ".rbx"
  - ".ru"
  - ".ruby"
  - ".spec"
  - ".thor"
  - ".watchr"
  interpreters:
  - ruby
  - macruby
  - rake
  - jruby
  - rbx
  filenames:
  - ".irbrc"
  - ".pryrc"
  - ".simplecov"
  - Appraisals
  - Berksfile
  - Brewfile
  - Buildfile
  - Capfile
  - Dangerfile
  - Deliverfile
  - Fastfile
  - Gemfile
  - Guardfile
  - Jarfile
  - Mavenfile
  - Podfile
  - Puppetfile
  - Rakefile
  - Snapfile
  - Steepfile
  - Thorfile
  - Vagrantfile
  - buildfile
  language_id: 326
Rust:
  type: programming
  aliases:
  - rs
  color: "#dea584"
  extensions:
  - ".rs"
  - ".rs.in"
  tm_scope: source.rust
  ace_mode: rust
  codemirror_mode: rust
  codemirror_mime_type: text/x-rustsrc
  language_id: 327
SAS:
  type: programming
  color: "#B34936"
  extensions:
  - ".sas"
  tm_scope: source.sas
  ace_mode: text
  codemirror_mode: sas
  codemirror_mime_type: text/x-sas
  language_id: 328
SCSS:
  type: markup
  color: "#c6538c"
  tm_scope: source.css.scss
  ace_mode: scss
  codemirror_mode: css
  codemirror_mime_type: text/x-scss
  extensions:
  - ".scss"
  language_id: 329
SELinux Policy:
  aliases:
  - SELinux Kernel Policy Language
  - sepolicy
  type: data
  tm_scope: source.sepolicy
  extensions:
  - ".te"
  filenames:
  - file_contexts
  - genfs_contexts
  - initial_sids
  - port_contexts
  - security_classes
  ace_mode: text
  language_id: 880010326
SMT:
  type: programming
  extensions:
  - ".smt2"
  - ".smt"
  interpreters:
  - boolector
  - cvc4
  - mathsat5
  - opensmt
  - smtinterpol
  - smt-rat
  - stp
  - verit
  - yices2
  - z3
  tm_scope: source.smt
  ace_mode: text
  language_id: 330
SPARQL:
  type: data
  color: "#0C4597"
  tm_scope: source.sparql
  ace_mode: text
  codemirror_mode: sparql
  codemirror_mime_type: application/sparql-query
  extensions:
  - ".sparql"
  - ".rq"
  language_id: 331
SQF:
  type: programming
  color: "#3F3F3F"
  extensions:
  - ".sqf"
  - ".hqf"
  tm_scope: source.sqf
  ace_mode: text
  language_id: 332
SQL:
  type: data
  color: "#e38c00"
  tm_scope: source.sql
  ace_mode: sql
  codemirror_mode: sql
  codemirror_mime_type: text/x-sql
  extensions:
  - ".sql"
  - ".cql"
  - ".ddl"
  - ".inc"
  - ".mysql"
  - ".prc"
  - ".tab"
  - ".udf"
  - ".viw"
  language_id: 333
SQLPL:
  type: programming
  color: "#e38c00"
  ace_mode: sql
  codemirror_mode: sql
  codemirror_mime_type: text/x-sql
  tm_scope: source.sql
  extensions:
  - ".sql"
  - ".db2"
  language_id: 334
SRecode Template:
  type: markup
  color: "#348a34"
  tm_scope: source.lisp
  ace_mode: lisp
  codemirror_mode: commonlisp
  codemirror_mime_type: text/x-common-lisp
  extensions:
  - ".srt"
  language_id: 335
SSH Config:
  type: data
  group: INI
  filenames:
  - ssh-config
  - ssh_config
  - sshconfig
  - sshconfig.snip
  - sshd-config
  - sshd_config
  ace_mode: text
  tm_scope: source.ssh-config
  language_id: 554920715
STON:
  type: data
  group: Smalltalk
  extensions:
  - ".ston"
  tm_scope: source.smalltalk
  ace_mode: text
  language_id: 336
SVG:
  type: data
  color: "#ff9900"
  extensions:
  - ".svg"
  tm_scope: text.xml.svg
  ace_mode: xml
  codemirror_mode: xml
  codemirror_mime_type: text/xml
  language_id: 337
SWIG:
  type: programming
  extensions:
  - ".i"
  tm_scope: source.c++
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-c++src
  language_id: 1066250075
Sage:
  type: programming
  extensions:
  - ".sage"
  - ".sagews"
  tm_scope: source.python
  ace_mode: python
  codemirror_mode: python
  codemirror_mime_type: text/x-python
  language_id: 338
SaltStack:
  type: programming
  color: "#646464"
  aliases:
  - saltstate
  - salt
  extensions:
  - ".sls"
  tm_scope: source.yaml.salt
  ace_mode: yaml
  codemirror_mode: yaml
  codemirror_mime_type: text/x-yaml
  language_id: 339
Sass:
  type: markup
  color: "#a53b70"
  tm_scope: source.sass
  extensions:
  - ".sass"
  ace_mode: sass
  codemirror_mode: sass
  codemirror_mime_type: text/x-sass
  language_id: 340
Scala:
  type: programming
  tm_scope: source.scala
  ace_mode: scala
  codemirror_mode: clike
  codemirror_mime_type: text/x-scala
  color: "#c22d40"
  extensions:
  - ".scala"
  - ".kojo"
  - ".sbt"
  - ".sc"
  interpreters:
  - scala
  language_id: 341
Scaml:
  type: markup
  color: "#bd181a"
  extensions:
  - ".scaml"
  tm_scope: source.scaml
  ace_mode: text
  language_id: 342
Scheme:
  type: programming
  color: "#1e4aec"
  extensions:
  - ".scm"
  - ".sch"
  - ".sld"
  - ".sls"
  - ".sps"
  - ".ss"
  interpreters:
  - scheme
  - guile
  - bigloo
  - chicken
  - csi
  - gosh
  - r6rs
  tm_scope: source.scheme
  ace_mode: scheme
  codemirror_mode: scheme
  codemirror_mime_type: text/x-scheme
  language_id: 343
Scilab:
  type: programming
  color: "#ca0f21"
  extensions:
  - ".sci"
  - ".sce"
  - ".tst"
  tm_scope: source.scilab
  ace_mode: text
  language_id: 344
Self:
  type: programming
  color: "#0579aa"
  extensions:
  - ".self"
  tm_scope: none
  ace_mode: text
  language_id: 345
ShaderLab:
  type: programming
  color: "#222c37"
  extensions:
  - ".shader"
  ace_mode: text
  tm_scope: source.shaderlab
  language_id: 664257356
Shell:
  type: programming
  color: "#89e051"
  aliases:
  - sh
  - shell-script
  - bash
  - zsh
  extensions:
  - ".sh"
  - ".bash"
  - ".bats"
  - ".cgi"
  - ".command"
  - ".env"
  - ".fcgi"
  - ".ksh"
  - ".sh.in"
  - ".tmux"
  - ".tool"
  - ".zsh"
  - ".zsh-theme"
  filenames:
  - ".bash_aliases"
  - ".bash_history"
  - ".bash_logout"
  - ".bash_profile"
  - ".bashrc"
  - ".cshrc"
  - ".env"
  - ".env.example"
  - ".flaskenv"
  - ".kshrc"
  - ".login"
  - ".profile"
  - ".zlogin"
  - ".zlogout"
  - ".zprofile"
  - ".zshenv"
  - ".zshrc"
  - 9fs
  - PKGBUILD
  - bash_aliases
  - bash_logout
  - bash_profile
  - bashrc
  - cshrc
  - gradlew
  - kshrc
  - login
  - man
  - profile
  - zlogin
  - zlogout
  - zprofile
  - zshenv
  - zshrc
  interpreters:
  - ash
  - bash
  - dash
  - ksh
  - mksh
  - pdksh
  - rc
  - sh
  - zsh
  tm_scope: source.shell
  ace_mode: sh
  codemirror_mode: shell
  codemirror_mime_type: text/x-sh
  language_id: 346
ShellCheck Config:
  type: data
  color: "#cecfcb"
  filenames:
  - ".shellcheckrc"
  aliases:
  - shellcheckrc
  tm_scope: source.shellcheckrc
  ace_mode: ini
  codemirror_mode: properties
  codemirror_mime_type: text/x-properties
  language_id: 687511714
ShellSession:
  type: programming
  extensions:
  - ".sh-session"
  aliases:
  - bash session
  - console
  tm_scope: text.shell-session
  ace_mode: sh
  codemirror_mode: shell
  codemirror_mime_type: text/x-sh
  language_id: 347
Shen:
  type: programming
  color: "#120F14"
  extensions:
  - ".shen"
  tm_scope: source.shen
  ace_mode: text
  language_id: 348
Sieve:
  type: programming
  tm_scope: source.sieve
  ace_mode: text
  extensions:
  - ".sieve"
  codemirror_mode: sieve
  codemirror_mime_type: application/sieve
  language_id: 208976687
Singularity:
  type: programming
  color: "#64E6AD"
  tm_scope: source.singularity
  filenames:
  - Singularity
  ace_mode: text
  language_id: 987024632
Slash:
  type: programming
  color: "#007eff"
  extensions:
  - ".sl"
  tm_scope: text.html.slash
  ace_mode: text
  language_id: 349
Slice:
  type: programming
  color: "#003fa2"
  tm_scope: source.slice
  ace_mode: text
  extensions:
  - ".ice"
  language_id: 894641667
Slim:
  type: markup
  color: "#2b2b2b"
  extensions:
  - ".slim"
  tm_scope: text.slim
  ace_mode: text
  codemirror_mode: slim
  codemirror_mime_type: text/x-slim
  language_id: 350
SmPL:
  type: programming
  extensions:
  - ".cocci"
  aliases:
  - coccinelle
  ace_mode: text
  tm_scope: source.smpl
  color: "#c94949"
  language_id: 164123055
Smali:
  type: programming
  extensions:
  - ".smali"
  ace_mode: text
  tm_scope: source.smali
  language_id: 351
Smalltalk:
  type: programming
  color: "#596706"
  extensions:
  - ".st"
  - ".cs"
  aliases:
  - squeak
  tm_scope: source.smalltalk
  ace_mode: text
  codemirror_mode: smalltalk
  codemirror_mime_type: text/x-stsrc
  language_id: 352
Smarty:
  type: programming
  color: "#f0c040"
  extensions:
  - ".tpl"
  ace_mode: smarty
  codemirror_mode: smarty
  codemirror_mime_type: text/x-smarty
  tm_scope: text.html.smarty
  language_id: 353
Solidity:
  type: programming
  color: "#AA6746"
  ace_mode: text
  tm_scope: source.solidity
  extensions:
  - ".sol"
  language_id: 237469032
Soong:
  type: data
  tm_scope: source.bp
  ace_mode: text
  filenames:
  - Android.bp
  language_id: 222900098
SourcePawn:
  type: programming
  color: "#f69e1d"
  aliases:
  - sourcemod
  extensions:
  - ".sp"
  - ".inc"
  tm_scope: source.sourcepawn
  ace_mode: text
  language_id: 354
Spline Font Database:
  type: data
  extensions:
  - ".sfd"
  tm_scope: text.sfd
  ace_mode: yaml
  language_id: 767169629
Squirrel:
  type: programming
  color: "#800000"
  extensions:
  - ".nut"
  tm_scope: source.nut
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-c++src
  language_id: 355
Stan:
  type: programming
  color: "#b2011d"
  extensions:
  - ".stan"
  ace_mode: text
  tm_scope: source.stan
  language_id: 356
Standard ML:
  type: programming
  color: "#dc566d"
  aliases:
  - sml
  extensions:
  - ".ml"
  - ".fun"
  - ".sig"
  - ".sml"
  tm_scope: source.ml
  ace_mode: text
  codemirror_mode: mllike
  codemirror_mime_type: text/x-ocaml
  language_id: 357
Starlark:
  type: programming
  tm_scope: source.python
  ace_mode: python
  codemirror_mode: python
  codemirror_mime_type: text/x-python
  color: "#76d275"
  extensions:
  - ".bzl"
  filenames:
  - BUCK
  - BUILD
  - BUILD.bazel
  - Tiltfile
  - WORKSPACE
  aliases:
  - bazel
  - bzl
  language_id: 960266174
Stata:
  type: programming
  color: "#1a5f91"
  extensions:
  - ".do"
  - ".ado"
  - ".doh"
  - ".ihlp"
  - ".mata"
  - ".matah"
  - ".sthlp"
  tm_scope: source.stata
  ace_mode: text
  language_id: 358
StringTemplate:
  type: markup
  color: "#3fb34f"
  extensions:
  - ".st"
  tm_scope: source.string-template
  ace_mode: html
  codemirror_mode: htmlmixed
  codemirror_mime_type: text/html
  language_id: 89855901
Stylus:
  type: markup
  color: "#ff6347"
  extensions:
  - ".styl"
  tm_scope: source.stylus
  ace_mode: stylus
  language_id: 359
SubRip Text:
  type: data
  color: "#9e0101"
  extensions:
  - ".srt"
  ace_mode: text
  tm_scope: text.srt
  language_id: 360
SugarSS:
  type: markup
  color: "#2fcc9f"
  tm_scope: source.css.postcss.sugarss
  group: CSS
  extensions:
  - ".sss"
  ace_mode: text
  language_id: 826404698
SuperCollider:
  type: programming
  color: "#46390b"
  extensions:
  - ".sc"
  - ".scd"
  interpreters:
  - sclang
  - scsynth
  tm_scope: source.supercollider
  ace_mode: text
  language_id: 361
Svelte:
  type: markup
  color: "#ff3e00"
  tm_scope: source.svelte
  ace_mode: html
  codemirror_mode: htmlmixed
  codemirror_mime_type: text/html
  extensions:
  - ".svelte"
  language_id: 928734530
Swift:
  type: programming
  color: "#F05138"
  extensions:
  - ".swift"
  tm_scope: source.swift
  ace_mode: text
  codemirror_mode: swift
  codemirror_mime_type: text/x-swift
  language_id: 362
SystemVerilog:
  type: programming
  color: "#DAE1C2"
  extensions:
  - ".sv"
  - ".svh"
  - ".vh"
  tm_scope: source.systemverilog
  ace_mode: verilog
  codemirror_mode: verilog
  codemirror_mime_type: text/x-systemverilog
  language_id: 363
TI Program:
  type: programming
  ace_mode: text
  color: "#A0AA87"
  extensions:
  - ".8xp"
  - ".8xk"
  - ".8xk.txt"
  - ".8xp.txt"
  language_id: 422
  tm_scope: none
TLA:
  type: programming
  color: "#4b0079"
  extensions:
  - ".tla"
  tm_scope: source.tla
  ace_mode: text
  language_id: 364
TOML:
  type: data
  color: "#9c4221"
  extensions:
  - ".toml"
  filenames:
  - Cargo.lock
  - Gopkg.lock
  - Pipfile
  - poetry.lock
  tm_scope: source.toml
  ace_mode: toml
  codemirror_mode: toml
  codemirror_mime_type: text/x-toml
  language_id: 365
TSQL:
  type: programming
  color: "#e38c00"
  extensions:
  - ".sql"
  ace_mode: sql
  tm_scope: source.tsql
  language_id: 918334941
TSV:
  type: data
  color: "#237346"
  ace_mode: text
  tm_scope: source.generic-db
  extensions:
  - ".tsv"
  language_id: 1035892117
TSX:
  type: programming
  color: "#2b7489"
  group: TypeScript
  extensions:
  - ".tsx"
  tm_scope: source.tsx
  ace_mode: javascript
  codemirror_mode: jsx
  codemirror_mime_type: text/jsx
  language_id: 94901924
TXL:
  type: programming
  color: "#0178b8"
  extensions:
  - ".txl"
  tm_scope: source.txl
  ace_mode: text
  language_id: 366
Talon:
  type: programming
  ace_mode: text
  color: "#333333"
  extensions:
  - ".talon"
  tm_scope: source.talon
  language_id: 959889508
Tcl:
  type: programming
  color: "#e4cc98"
  extensions:
  - ".tcl"
  - ".adp"
  - ".tcl.in"
  - ".tm"
  filenames:
  - owh
  - starfield
  interpreters:
  - tclsh
  - wish
  tm_scope: source.tcl
  ace_mode: tcl
  codemirror_mode: tcl
  codemirror_mime_type: text/x-tcl
  language_id: 367
Tcsh:
  type: programming
  group: Shell
  extensions:
  - ".tcsh"
  - ".csh"
  interpreters:
  - tcsh
  - csh
  tm_scope: source.shell
  ace_mode: sh
  codemirror_mode: shell
  codemirror_mime_type: text/x-sh
  language_id: 368
TeX:
  type: markup
  color: "#3D6117"
  ace_mode: tex
  codemirror_mode: stex
  codemirror_mime_type: text/x-stex
  tm_scope: text.tex.latex
  wrap: true
  aliases:
  - latex
  extensions:
  - ".tex"
  - ".aux"
  - ".bbx"
  - ".cbx"
  - ".cls"
  - ".dtx"
  - ".ins"
  - ".lbx"
  - ".ltx"
  - ".mkii"
  - ".mkiv"
  - ".mkvi"
  - ".sty"
  - ".toc"
  language_id: 369
Tea:
  type: markup
  extensions:
  - ".tea"
  tm_scope: source.tea
  ace_mode: text
  language_id: 370
Terra:
  type: programming
  extensions:
  - ".t"
  color: "#00004c"
  tm_scope: source.terra
  ace_mode: lua
  codemirror_mode: lua
  codemirror_mime_type: text/x-lua
  interpreters:
  - lua
  language_id: 371
Texinfo:
  type: prose
  wrap: true
  extensions:
  - ".texinfo"
  - ".texi"
  - ".txi"
  ace_mode: text
  tm_scope: text.texinfo
  interpreters:
  - makeinfo
  language_id: 988020015
Text:
  type: prose
  wrap: true
  aliases:
  - fundamental
  - plain text
  extensions:
  - ".txt"
  - ".fr"
  - ".nb"
  - ".ncl"
  - ".no"
  filenames:
  - CITATION
  - CITATIONS
  - COPYING
  - COPYING.regex
  - COPYRIGHT.regex
  - FONTLOG
  - INSTALL
  - INSTALL.mysql
  - LICENSE
  - LICENSE.mysql
  - NEWS
  - README.me
  - README.mysql
  - README.nss
  - click.me
  - delete.me
  - keep.me
  - package.mask
  - package.use.mask
  - package.use.stable.mask
  - read.me
  - readme.1st
  - test.me
  - use.mask
  - use.stable.mask
  tm_scope: none
  ace_mode: text
  language_id: 372
TextMate Properties:
  type: data
  color: "#df66e4"
  aliases:
  - tm-properties
  filenames:
  - ".tm_properties"
  ace_mode: properties
  codemirror_mode: properties
  codemirror_mime_type: text/x-properties
  tm_scope: source.tm-properties
  language_id: 981795023
Textile:
  type: prose
  color: "#ffe7ac"
  ace_mode: textile
  codemirror_mode: textile
  codemirror_mime_type: text/x-textile
  wrap: true
  extensions:
  - ".textile"
  tm_scope: none
  language_id: 373
Thrift:
  type: programming
  color: "#D12127"
  tm_scope: source.thrift
  extensions:
  - ".thrift"
  ace_mode: text
  language_id: 374
Turing:
  type: programming
  color: "#cf142b"
  extensions:
  - ".t"
  - ".tu"
  tm_scope: source.turing
  ace_mode: text
  language_id: 375
Turtle:
  type: data
  extensions:
  - ".ttl"
  tm_scope: source.turtle
  ace_mode: text
  codemirror_mode: turtle
  codemirror_mime_type: text/turtle
  language_id: 376
Twig:
  type: markup
  color: "#c1d026"
  extensions:
  - ".twig"
  tm_scope: text.html.twig
  ace_mode: twig
  codemirror_mode: twig
  codemirror_mime_type: text/x-twig
  language_id: 377
Type Language:
  type: data
  aliases:
  - tl
  extensions:
  - ".tl"
  tm_scope: source.tl
  ace_mode: text
  language_id: 632765617
TypeScript:
  type: programming
  color: "#2b7489"
  aliases:
  - ts
  interpreters:
  - deno
  - ts-node
  extensions:
  - ".ts"
  tm_scope: source.ts
  ace_mode: typescript
  codemirror_mode: javascript
  codemirror_mime_type: application/typescript
  language_id: 378
Unified Parallel C:
  type: programming
  color: "#4e3617"
  group: C
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-csrc
  extensions:
  - ".upc"
  tm_scope: source.c
  language_id: 379
Unity3D Asset:
  type: data
  color: "#222c37"
  ace_mode: yaml
  codemirror_mode: yaml
  codemirror_mime_type: text/x-yaml
  extensions:
  - ".anim"
  - ".asset"
  - ".mask"
  - ".mat"
  - ".meta"
  - ".prefab"
  - ".unity"
  tm_scope: source.yaml
  language_id: 380
Unix Assembly:
  type: programming
  group: Assembly
  extensions:
  - ".s"
  - ".ms"
  tm_scope: source.x86
  ace_mode: assembly_x86
  language_id: 120
Uno:
  type: programming
  color: "#9933cc"
  extensions:
  - ".uno"
  ace_mode: csharp
  codemirror_mode: clike
  codemirror_mime_type: text/x-csharp
  tm_scope: source.cs
  language_id: 381
UnrealScript:
  type: programming
  color: "#a54c4d"
  extensions:
  - ".uc"
  tm_scope: source.java
  ace_mode: java
  codemirror_mode: clike
  codemirror_mime_type: text/x-java
  language_id: 382
UrWeb:
  type: programming
  color: "#ccccee"
  aliases:
  - Ur/Web
  - Ur
  extensions:
  - ".ur"
  - ".urs"
  tm_scope: source.ur
  ace_mode: text
  language_id: 383
V:
  type: programming
  color: "#4f87c4"
  aliases:
  - vlang
  extensions:
  - ".v"
  tm_scope: source.v
  ace_mode: golang
  codemirror_mode: go
  codemirror_mime_type: text/x-go
  language_id: 603371597
VBA:
  type: programming
  color: "#867db1"
  extensions:
  - ".bas"
  - ".cls"
  - ".frm"
  - ".frx"
  - ".vba"
  tm_scope: source.vbnet
  aliases:
  - vb6
  - visual basic 6
  - visual basic for applications
  ace_mode: text
  codemirror_mode: vb
  codemirror_mime_type: text/x-vb
  language_id: 399230729
VBScript:
  type: programming
  color: "#15dcdc"
  extensions:
  - ".vbs"
  tm_scope: source.vbnet
  ace_mode: text
  codemirror_mode: vbscript
  codemirror_mime_type: text/vbscript
  language_id: 408016005
VCL:
  type: programming
  color: "#148AA8"
  extensions:
  - ".vcl"
  tm_scope: source.varnish.vcl
  ace_mode: text
  language_id: 384
VHDL:
  type: programming
  color: "#adb2cb"
  extensions:
  - ".vhdl"
  - ".vhd"
  - ".vhf"
  - ".vhi"
  - ".vho"
  - ".vhs"
  - ".vht"
  - ".vhw"
  tm_scope: source.vhdl
  ace_mode: vhdl
  codemirror_mode: vhdl
  codemirror_mime_type: text/x-vhdl
  language_id: 385
Vala:
  type: programming
  color: "#a56de2"
  extensions:
  - ".vala"
  - ".vapi"
  tm_scope: source.vala
  ace_mode: vala
  language_id: 386
Valve Data Format:
  type: data
  color: "#f26025"
  aliases:
  - keyvalues
  - vdf
  extensions:
  - ".vdf"
  ace_mode: text
  tm_scope: source.keyvalues
  language_id: 544060961
Verilog:
  type: programming
  color: "#b2b7f8"
  extensions:
  - ".v"
  - ".veo"
  tm_scope: source.verilog
  ace_mode: verilog
  codemirror_mode: verilog
  codemirror_mime_type: text/x-verilog
  language_id: 387
Vim Help File:
  type: prose
  color: "#199f4b"
  aliases:
  - help
  - vimhelp
  extensions:
  - ".txt"
  tm_scope: text.vim-help
  ace_mode: text
  language_id: 508563686
Vim Script:
  type: programming
  color: "#199f4b"
  tm_scope: source.viml
  aliases:
  - vim
  - viml
  - nvim
  extensions:
  - ".vim"
  - ".vba"
  - ".vimrc"
  - ".vmb"
  filenames:
  - ".exrc"
  - ".gvimrc"
  - ".nvimrc"
  - ".vimrc"
  - _vimrc
  - gvimrc
  - nvimrc
  - vimrc
  ace_mode: text
  language_id: 388
Vim Snippet:
  type: markup
  color: "#199f4b"
  aliases:
  - SnipMate
  - UltiSnip
  - UltiSnips
  - NeoSnippet
  extensions:
  - ".snip"
  - ".snippet"
  - ".snippets"
  tm_scope: source.vim-snippet
  ace_mode: text
  language_id: 81265970
Visual Basic .NET:
  type: programming
  color: "#945db7"
  extensions:
  - ".vb"
  - ".vbhtml"
  aliases:
  - visual basic
  - vbnet
  - vb .net
  - vb.net
  tm_scope: source.vbnet
  ace_mode: text
  codemirror_mode: vb
  codemirror_mime_type: text/x-vb
  language_id: 389
Volt:
  type: programming
  color: "#1F1F1F"
  extensions:
  - ".volt"
  tm_scope: source.d
  ace_mode: d
  codemirror_mode: d
  codemirror_mime_type: text/x-d
  language_id: 390
Vue:
  type: markup
  color: "#41b883"
  extensions:
  - ".vue"
  tm_scope: text.html.vue
  ace_mode: html
  language_id: 391
Vyper:
  type: programming
  extensions:
  - ".vy"
  color: "#2980b9"
  ace_mode: text
  tm_scope: source.vyper
  language_id: 1055641948
Wavefront Material:
  type: data
  extensions:
  - ".mtl"
  tm_scope: source.wavefront.mtl
  ace_mode: text
  language_id: 392
Wavefront Object:
  type: data
  extensions:
  - ".obj"
  tm_scope: source.wavefront.obj
  ace_mode: text
  language_id: 393
Web Ontology Language:
  type: data
  color: "#5b70bd"
  extensions:
  - ".owl"
  tm_scope: text.xml
  ace_mode: xml
  language_id: 394
WebAssembly:
  type: programming
  color: "#04133b"
  extensions:
  - ".wast"
  - ".wat"
  aliases:
  - wast
  - wasm
  tm_scope: source.webassembly
  ace_mode: lisp
  codemirror_mode: commonlisp
  codemirror_mime_type: text/x-common-lisp
  language_id: 956556503
WebIDL:
  type: programming
  extensions:
  - ".webidl"
  tm_scope: source.webidl
  ace_mode: text
  codemirror_mode: webidl
  codemirror_mime_type: text/x-webidl
  language_id: 395
WebVTT:
  type: data
  wrap: true
  extensions:
  - ".vtt"
  tm_scope: source.vtt
  ace_mode: text
  language_id: 658679714
Wget Config:
  type: data
  group: INI
  aliases:
  - wgetrc
  filenames:
  - ".wgetrc"
  tm_scope: source.wgetrc
  ace_mode: text
  language_id: 668457123
Wikitext:
  type: prose
  color: "#fc5757"
  wrap: true
  aliases:
  - mediawiki
  - wiki
  extensions:
  - ".mediawiki"
  - ".wiki"
  - ".wikitext"
  tm_scope: text.html.mediawiki
  ace_mode: text
  language_id: 228
Win32 Message File:
  type: data
  extensions:
  - ".mc"
  tm_scope: source.win32-messages
  ace_mode: ini
  codemirror_mode: properties
  codemirror_mime_type: text/x-properties
  language_id: 950967261
Windows Registry Entries:
  type: data
  color: "#52d5ff"
  extensions:
  - ".reg"
  tm_scope: source.reg
  ace_mode: ini
  codemirror_mode: properties
  codemirror_mime_type: text/x-properties
  language_id: 969674868
Witcher Script:
  type: programming
  color: "#ff0000"
  extensions:
  - ".ws"
  ace_mode: text
  tm_scope: source.witcherscript
  language_id: 686821385
Wollok:
  type: programming
  color: "#a23738"
  extensions:
  - ".wlk"
  ace_mode: text
  tm_scope: source.wollok
  language_id: 632745969
World of Warcraft Addon Data:
  type: data
  color: "#f7e43f"
  extensions:
  - ".toc"
  tm_scope: source.toc
  ace_mode: text
  language_id: 396
X BitMap:
  type: data
  group: C
  aliases:
  - xbm
  extensions:
  - ".xbm"
  ace_mode: c_cpp
  tm_scope: source.c
  codemirror_mode: clike
  codemirror_mime_type: text/x-csrc
  language_id: 782911107
X Font Directory Index:
  type: data
  filenames:
  - encodings.dir
  - fonts.alias
  - fonts.dir
  - fonts.scale
  tm_scope: source.fontdir
  ace_mode: text
  language_id: 208700028
X PixMap:
  type: data
  group: C
  aliases:
  - xpm
  extensions:
  - ".xpm"
  - ".pm"
  ace_mode: c_cpp
  tm_scope: source.c
  codemirror_mode: clike
  codemirror_mime_type: text/x-csrc
  language_id: 781846279
X10:
  type: programming
  aliases:
  - xten
  ace_mode: text
  extensions:
  - ".x10"
  color: "#4B6BEF"
  tm_scope: source.x10
  language_id: 397
XC:
  type: programming
  color: "#99DA07"
  extensions:
  - ".xc"
  tm_scope: source.xc
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-csrc
  language_id: 398
XCompose:
  type: data
  filenames:
  - ".XCompose"
  - XCompose
  - xcompose
  tm_scope: config.xcompose
  ace_mode: text
  language_id: 225167241
XML:
  type: data
  color: "#0060ac"
  tm_scope: text.xml
  ace_mode: xml
  codemirror_mode: xml
  codemirror_mime_type: text/xml
  aliases:
  - rss
  - xsd
  - wsdl
  extensions:
  - ".xml"
  - ".adml"
  - ".admx"
  - ".ant"
  - ".axaml"
  - ".axml"
  - ".builds"
  - ".ccproj"
  - ".ccxml"
  - ".clixml"
  - ".cproject"
  - ".cscfg"
  - ".csdef"
  - ".csl"
  - ".csproj"
  - ".ct"
  - ".depproj"
  - ".dita"
  - ".ditamap"
  - ".ditaval"
  - ".dll.config"
  - ".dotsettings"
  - ".filters"
  - ".fsproj"
  - ".fxml"
  - ".glade"
  - ".gml"
  - ".gmx"
  - ".grxml"
  - ".gst"
  - ".hzp"
  - ".iml"
  - ".ivy"
  - ".jelly"
  - ".jsproj"
  - ".kml"
  - ".launch"
  - ".mdpolicy"
  - ".mjml"
  - ".mm"
  - ".mod"
  - ".mxml"
  - ".natvis"
  - ".ncl"
  - ".ndproj"
  - ".nproj"
  - ".nuspec"
  - ".odd"
  - ".osm"
  - ".pkgproj"
  - ".pluginspec"
  - ".proj"
  - ".props"
  - ".ps1xml"
  - ".psc1"
  - ".pt"
  - ".rdf"
  - ".res"
  - ".resx"
  - ".rs"
  - ".rss"
  - ".sch"
  - ".scxml"
  - ".sfproj"
  - ".shproj"
  - ".srdf"
  - ".storyboard"
  - ".sublime-snippet"
  - ".targets"
  - ".tml"
  - ".ts"
  - ".tsx"
  - ".ui"
  - ".urdf"
  - ".ux"
  - ".vbproj"
  - ".vcxproj"
  - ".vsixmanifest"
  - ".vssettings"
  - ".vstemplate"
  - ".vxml"
  - ".wixproj"
  - ".workflow"
  - ".wsdl"
  - ".wsf"
  - ".wxi"
  - ".wxl"
  - ".wxs"
  - ".x3d"
  - ".xacro"
  - ".xaml"
  - ".xib"
  - ".xlf"
  - ".xliff"
  - ".xmi"
  - ".xml.dist"
  - ".xmp"
  - ".xproj"
  - ".xsd"
  - ".xspec"
  - ".xul"
  - ".zcml"
  filenames:
  - ".classpath"
  - ".cproject"
  - ".project"
  - App.config
  - NuGet.config
  - Settings.StyleCop
  - Web.Debug.config
  - Web.Release.config
  - Web.config
  - packages.config
  language_id: 399
XML Property List:
  type: data
  color: "#0060ac"
  group: XML
  extensions:
  - ".plist"
  - ".stTheme"
  - ".tmCommand"
  - ".tmLanguage"
  - ".tmPreferences"
  - ".tmSnippet"
  - ".tmTheme"
  tm_scope: text.xml.plist
  ace_mode: xml
  codemirror_mode: xml
  codemirror_mime_type: text/xml
  language_id: 75622871
XPages:
  type: data
  extensions:
  - ".xsp-config"
  - ".xsp.metadata"
  tm_scope: text.xml
  ace_mode: xml
  codemirror_mode: xml
  codemirror_mime_type: text/xml
  language_id: 400
XProc:
  type: programming
  extensions:
  - ".xpl"
  - ".xproc"
  tm_scope: text.xml
  ace_mode: xml
  codemirror_mode: xml
  codemirror_mime_type: text/xml
  language_id: 401
XQuery:
  type: programming
  color: "#5232e7"
  extensions:
  - ".xquery"
  - ".xq"
  - ".xql"
  - ".xqm"
  - ".xqy"
  ace_mode: xquery
  codemirror_mode: xquery
  codemirror_mime_type: application/xquery
  tm_scope: source.xq
  language_id: 402
XS:
  type: programming
  extensions:
  - ".xs"
  tm_scope: source.c
  ace_mode: c_cpp
  codemirror_mode: clike
  codemirror_mime_type: text/x-csrc
  language_id: 403
XSLT:
  type: programming
  aliases:
  - xsl
  extensions:
  - ".xslt"
  - ".xsl"
  tm_scope: text.xml.xsl
  ace_mode: xml
  codemirror_mode: xml
  codemirror_mime_type: text/xml
  color: "#EB8CEB"
  language_id: 404
Xojo:
  type: programming
  color: "#81bd41"
  extensions:
  - ".xojo_code"
  - ".xojo_menu"
  - ".xojo_report"
  - ".xojo_script"
  - ".xojo_toolbar"
  - ".xojo_window"
  tm_scope: source.xojo
  ace_mode: text
  language_id: 405
Xonsh:
  type: programming
  color: "#285EEF"
  extensions:
  - ".xsh"
  tm_scope: source.python
  ace_mode: text
  codemirror_mode: python
  codemirror_mime_type: text/x-python
  language_id: 614078284
Xtend:
  type: programming
  color: "#24255d"
  extensions:
  - ".xtend"
  tm_scope: source.xtend
  ace_mode: text
  language_id: 406
YAML:
  type: data
  color: "#cb171e"
  tm_scope: source.yaml
  aliases:
  - yml
  extensions:
  - ".yml"
  - ".mir"
  - ".reek"
  - ".rviz"
  - ".sublime-syntax"
  - ".syntax"
  - ".yaml"
  - ".yaml-tmlanguage"
  - ".yaml.sed"
  - ".yml.mysql"
  filenames:
  - ".clang-format"
  - ".clang-tidy"
  - ".gemrc"
  - CITATION.cff
  - glide.lock
  - yarn.lock
  ace_mode: yaml
  codemirror_mode: yaml
  codemirror_mime_type: text/x-yaml
  language_id: 407
YANG:
  type: data
  extensions:
  - ".yang"
  tm_scope: source.yang
  ace_mode: text
  language_id: 408
YARA:
  type: programming
  color: "#220000"
  ace_mode: text
  extensions:
  - ".yar"
  - ".yara"
  tm_scope: source.yara
  language_id: 805122868
YASnippet:
  type: markup
  aliases:
  - snippet
  - yas
  color: "#32AB90"
  extensions:
  - ".yasnippet"
  tm_scope: source.yasnippet
  ace_mode: text
  language_id: 378760102
Yacc:
  type: programming
  extensions:
  - ".y"
  - ".yacc"
  - ".yy"
  tm_scope: source.yacc
  ace_mode: text
  color: "#4B6C4B"
  language_id: 409
ZAP:
  type: programming
  color: "#0d665e"
  extensions:
  - ".zap"
  - ".xzap"
  tm_scope: source.zap
  ace_mode: text
  language_id: 952972794
ZIL:
  type: programming
  color: "#dc75e5"
  extensions:
  - ".zil"
  - ".mud"
  tm_scope: source.zil
  ace_mode: text
  language_id: 973483626
Zeek:
  type: programming
  aliases:
  - bro
  extensions:
  - ".zeek"
  - ".bro"
  tm_scope: source.zeek
  ace_mode: text
  language_id: 40
ZenScript:
  type: programming
  color: "#00BCD1"
  extensions:
  - ".zs"
  tm_scope: source.zenscript
  ace_mode: text
  language_id: 494938890
Zephir:
  type: programming
  color: "#118f9e"
  extensions:
  - ".zep"
  tm_scope: source.php.zephir
  ace_mode: php
  language_id: 410
Zig:
  type: programming
  color: "#ec915c"
  extensions:
  - ".zig"
  tm_scope: source.zig
  ace_mode: text
  language_id: 646424281
Zimpl:
  type: programming
  color: "#d67711"
  extensions:
  - ".zimpl"
  - ".zmpl"
  - ".zpl"
  tm_scope: none
  ace_mode: text
  language_id: 411
cURL Config:
  type: data
  group: INI
  aliases:
  - curlrc
  filenames:
  - ".curlrc"
  - _curlrc
  tm_scope: source.curlrc
  ace_mode: text
  language_id: 992375436
desktop:
  type: data
  extensions:
  - ".desktop"
  - ".desktop.in"
  - ".service"
  tm_scope: source.desktop
  ace_mode: text
  language_id: 412
dircolors:
  type: data
  extensions:
  - ".dircolors"
  filenames:
  - ".dir_colors"
  - ".dircolors"
  - DIR_COLORS
  - _dir_colors
  - _dircolors
  - dir_colors
  tm_scope: source.dircolors
  ace_mode: text
  language_id: 691605112
eC:
  type: programming
  color: "#913960"
  extensions:
  - ".ec"
  - ".eh"
  tm_scope: source.c.ec
  ace_mode: text
  language_id: 413
edn:
  type: data
  ace_mode: clojure
  codemirror_mode: clojure
  codemirror_mime_type: text/x-clojure
  extensions:
  - ".edn"
  tm_scope: source.clojure
  language_id: 414
fish:
  type: programming
  color: "#4aae47"
  group: Shell
  interpreters:
  - fish
  extensions:
  - ".fish"
  tm_scope: source.fish
  ace_mode: text
  language_id: 415
hoon:
  type: programming
  color: "#00b171"
  tm_scope: source.hoon
  ace_mode: text
  extensions:
  - ".hoon"
  language_id: 560883276
jq:
  color: "#c7254e"
  ace_mode: text
  type: programming
  extensions:
  - ".jq"
  tm_scope: source.jq
  language_id: 905371884
kvlang:
  type: markup
  ace_mode: text
  extensions:
  - ".kv"
  color: "#1da6e0"
  tm_scope: source.python.kivy
  language_id: 970675279
mIRC Script:
  type: programming
  color: "#3d57c3"
  extensions:
  - ".mrc"
  tm_scope: source.msl
  ace_mode: text
  language_id: 517654727
mcfunction:
  type: programming
  color: "#E22837"
  extensions:
  - ".mcfunction"
  tm_scope: source.mcfunction
  ace_mode: text
  language_id: 462488745
mupad:
  type: programming
  color: "#244963"
  extensions:
  - ".mu"
  tm_scope: source.mupad
  ace_mode: text
  language_id: 416
nanorc:
  type: data
  color: "#2d004d"
  group: INI
  extensions:
  - ".nanorc"
  filenames:
  - ".nanorc"
  - nanorc
  tm_scope: source.nanorc
  ace_mode: text
  language_id: 775996197
nesC:
  type: programming
  color: "#94B0C7"
  extensions:
  - ".nc"
  ace_mode: text
  tm_scope: source.nesc
  language_id: 417
ooc:
  type: programming
  color: "#b0b77e"
  extensions:
  - ".ooc"
  tm_scope: source.ooc
  ace_mode: text
  language_id: 418
q:
  type: programming
  extensions:
  - ".q"
  tm_scope: source.q
  ace_mode: text
  color: "#0040cd"
  language_id: 970539067
reStructuredText:
  type: prose
  color: "#141414"
  wrap: true
  aliases:
  - rst
  extensions:
  - ".rst"
  - ".rest"
  - ".rest.txt"
  - ".rst.txt"
  tm_scope: text.restructuredtext
  ace_mode: text
  codemirror_mode: rst
  codemirror_mime_type: text/x-rst
  language_id: 419
robots.txt:
  type: data
  aliases:
  - robots
  - robots txt
  filenames:
  - robots.txt
  ace_mode: text
  tm_scope: text.robots-txt
  language_id: 674736065
sed:
  type: programming
  color: "#64b970"
  extensions:
  - ".sed"
  interpreters:
  - gsed
  - minised
  - sed
  - ssed
  ace_mode: text
  tm_scope: source.sed
  language_id: 847830017
wdl:
  type: programming
  color: "#42f1f4"
  extensions:
  - ".wdl"
  tm_scope: source.wdl
  ace_mode: text
  language_id: 374521672
wisp:
  type: programming
  ace_mode: clojure
  codemirror_mode: clojure
  codemirror_mime_type: text/x-clojure
  color: "#7582D1"
  extensions:
  - ".wisp"
  tm_scope: source.clojure
  language_id: 420
xBase:
  type: programming
  color: "#403a40"
  aliases:
  - advpl
  - clipper
  - foxpro
  extensions:
  - ".prg"
  - ".ch"
  - ".prw"
  tm_scope: source.harbour
  ace_mode: text
  language_id: 421
`,

	"data/vendor.yml": `# Vendored files and directories are excluded from language
# statistics.
#
# Lines in this file are Regexps that are matched against the file
# pathname.
#
# Please add additional test coverage to
# `+"`"+`test/test_file_blob.rb#test_vendored`+"`"+` if you make any changes.

## Vendor Conventions ##

# Caches
- (^|/)cache/

# Dependencies
- ^[Dd]ependencies/

# Distributions
- (^|/)dist/

# C deps
- ^deps/
- (^|/)configure$
- (^|/)config\.guess$
- (^|/)config\.sub$

# stuff autogenerated by autoconf - still C deps
- (^|/)aclocal\.m4
- (^|/)libtool\.m4
- (^|/)ltoptions\.m4
- (^|/)ltsugar\.m4
- (^|/)ltversion\.m4
- (^|/)lt~obsolete\.m4

# .NET Core Install Scripts
- (^|/)dotnet-install\.(ps1|sh)$

# Linters
- (^|/)cpplint\.py

# Node dependencies
- (^|/)node_modules/

# Yarn 2
- (^|/)\.yarn/releases/
- (^|/)\.yarn/plugins/
- (^|/)\.yarn/sdks/
- (^|/)\.yarn/versions/
- (^|/)\.yarn/unplugged/

# esy.sh dependencies
- (^|/)_esy$

# Bower Components
- (^|/)bower_components/

# Erlang bundles
- ^rebar$
- (^|/)erlang\.mk

# Go dependencies
- (^|/)Godeps/_workspace/

# Go fixtures
- (^|/)testdata/

# GNU indent profiles
- (^|/)\.indent\.pro

# Minified JavaScript and CSS
- (\.|-)min\.(js|css)$

# Stylesheets imported from packages
- ([^\s]*)import\.(css|less|scss|styl)$

# Bootstrap css and js
- (^|/)bootstrap([^/.]*)\.(js|css|less|scss|styl)$
- (^|/)custom\.bootstrap([^\s]*)(js|css|less|scss|styl)$

# Font Awesome
- (^|/)font-?awesome\.(css|less|scss|styl)$
- (^|/)font-?awesome/.*\.(css|less|scss|styl)$

# Foundation css
- (^|/)foundation\.(css|less|scss|styl)$

# Normalize.css
- (^|/)normalize\.(css|less|scss|styl)$

# Skeleton.css
- (^|/)skeleton\.(css|less|scss|styl)$

# Bourbon css
- (^|/)[Bb]ourbon/.*\.(css|less|scss|styl)$

# Animate.css
- (^|/)animate\.(css|less|scss|styl)$

# Materialize.css
- (^|/)materialize\.(css|less|scss|styl|js)$

# Select2
- (^|/)select2/.*\.(css|scss|js)$

# Bulma css
- (^|/)bulma\.(css|sass|scss)$

# Vendored dependencies
- (3rd|[Tt]hird)[-_]?[Pp]arty/
- (^|/)vendors?/
- (^|/)[Ee]xtern(als?)?/
- (^|/)[Vv]+endor/

# Debian packaging
- ^debian/

# Haxelib projects often contain a neko bytecode file named run.n
- (^|/)run\.n$

# Bootstrap Datepicker
- (^|/)bootstrap-datepicker/

## Commonly Bundled JavaScript frameworks ##

# jQuery
- (^|/)jquery([^.]*)\.js$
- (^|/)jquery\-\d\.\d+(\.\d+)?\.js$

# jQuery UI
- (^|/)jquery\-ui(\-\d\.\d+(\.\d+)?)?(\.\w+)?\.(js|css)$
- (^|/)jquery\.(ui|effects)\.([^.]*)\.(js|css)$

# jQuery Gantt
- (^|/)jquery\.fn\.gantt\.js

# jQuery fancyBox
- (^|/)jquery\.fancybox\.(js|css)

# Fuel UX
- (^|/)fuelux\.js

# jQuery File Upload
- (^|/)jquery\.fileupload(-\w+)?\.js$

# jQuery dataTables
- (^|/)jquery\.dataTables\.js

# bootboxjs
- (^|/)bootbox\.js

# pdf-worker
- (^|/)pdf\.worker\.js

# Slick
- (^|/)slick\.\w+.js$

# Leaflet plugins
- (^|/)Leaflet\.Coordinates-\d+\.\d+\.\d+\.src\.js$
- (^|/)leaflet\.draw-src\.js
- (^|/)leaflet\.draw\.css
- (^|/)Control\.FullScreen\.css
- (^|/)Control\.FullScreen\.js
- (^|/)leaflet\.spin\.js
- (^|/)wicket-leaflet\.js

# Sublime Text workspace files
- (^|/)\.sublime-project
- (^|/)\.sublime-workspace

# VS Code workspace files
- (^|/)\.vscode/

# Prototype
- (^|/)prototype(.*)\.js$
- (^|/)effects\.js$
- (^|/)controls\.js$
- (^|/)dragdrop\.js$

# Typescript definition files
- (.*?)\.d\.ts$

# MooTools
- (^|/)mootools([^.]*)\d+\.\d+.\d+([^.]*)\.js$

# Dojo
- (^|/)dojo\.js$

# MochiKit
- (^|/)MochiKit\.js$

# YUI
- (^|/)yahoo-([^.]*)\.js$
- (^|/)yui([^.]*)\.js$

# WYS editors
- (^|/)ckeditor\.js$
- (^|/)tiny_mce([^.]*)\.js$
- (^|/)tiny_mce/(langs|plugins|themes|utils)

# Ace Editor
- (^|/)ace-builds/

# Fontello CSS files
- (^|/)fontello(.*?)\.css$

# MathJax
- (^|/)MathJax/

# Chart.js
- (^|/)Chart\.js$

# CodeMirror
- (^|/)[Cc]ode[Mm]irror/(\d+\.\d+/)?(lib|mode|theme|addon|keymap|demo)

# SyntaxHighlighter - http://alexgorbatchev.com/
- (^|/)shBrush([^.]*)\.js$
- (^|/)shCore\.js$
- (^|/)shLegacy\.js$

# AngularJS
- (^|/)angular([^.]*)\.js$

# D3.js
- (^|\/)d3(\.v\d+)?([^.]*)\.js$

# React
- (^|/)react(-[^.]*)?\.js$

# flow-typed
- (^|/)flow-typed/.*\.js$

# Modernizr
- (^|/)modernizr\-\d\.\d+(\.\d+)?\.js$
- (^|/)modernizr\.custom\.\d+\.js$

# Knockout
- (^|/)knockout-(\d+\.){3}(debug\.)?js$

## Python ##

# Sphinx
- (^|/)docs?/_?(build|themes?|templates?|static)/

# django
- (^|/)admin_media/
- (^|/)env/

# Fabric
- (^|/)fabfile\.py$

# WAF
- (^|/)waf$

# .osx
- (^|/)\.osx$

## Obj-C ##

# Xcode
### these can be part of a directory name

- \.xctemplate/
- \.imageset/

# Carthage
- (^|/)Carthage/

# Sparkle
- (^|/)Sparkle/

# Crashlytics
- (^|/)Crashlytics\.framework/

# Fabric
- (^|/)Fabric\.framework/

# BuddyBuild
- (^|/)BuddyBuildSDK\.framework/

# Realm
- (^|/)Realm\.framework

# RealmSwift
- (^|/)RealmSwift\.framework

# git config files
- (^|/)\.gitattributes$
- (^|/)\.gitignore$
- (^|/)\.gitmodules$

## Groovy ##

# Gradle
- (^|/)gradlew$
- (^|/)gradlew\.bat$
- (^|/)gradle/wrapper/

## Java ##

# Maven
- (^|/)mvnw$
- (^|/)mvnw\.cmd$
- (^|/)\.mvn/wrapper/

## .NET ##

# Visual Studio IntelliSense
- -vsdoc\.js$
- \.intellisense\.js$

# jQuery validation plugin (MS bundles this with asp.net mvc)
- (^|/)jquery([^.]*)\.validate(\.unobtrusive)?\.js$
- (^|/)jquery([^.]*)\.unobtrusive\-ajax\.js$

# Microsoft Ajax
- (^|/)[Mm]icrosoft([Mm]vc)?([Aa]jax|[Vv]alidation)(\.debug)?\.js$

# NuGet
- (^|/)[Pp]ackages\/.+\.\d+\/

# ExtJS
- (^|/)extjs/.*?\.js$
- (^|/)extjs/.*?\.xml$
- (^|/)extjs/.*?\.txt$
- (^|/)extjs/.*?\.html$
- (^|/)extjs/.*?\.properties$
- (^|/)extjs/\.sencha/
- (^|/)extjs/docs/
- (^|/)extjs/builds/
- (^|/)extjs/cmd/
- (^|/)extjs/examples/
- (^|/)extjs/locale/
- (^|/)extjs/packages/
- (^|/)extjs/plugins/
- (^|/)extjs/resources/
- (^|/)extjs/src/
- (^|/)extjs/welcome/

# Html5shiv
- (^|/)html5shiv\.js$

# Test fixtures
- (^|/)[Tt]ests?/fixtures/
- (^|/)[Ss]pecs?/fixtures/

# PhoneGap/Cordova
- (^|/)cordova([^.]*)\.js$
- (^|/)cordova\-\d\.\d(\.\d)?\.js$

# Foundation js
- (^|/)foundation(\..*)?\.js$

# Vagrant
- (^|/)Vagrantfile$

# .DS_Stores
- (^|/)\.[Dd][Ss]_[Ss]tore$

# R packages
- (^|/)vignettes/
- (^|/)inst/extdata/

# Octicons
- (^|/)octicons\.css
- (^|/)sprockets-octicons\.scss

# Typesafe Activator
- (^|/)activator$
- (^|/)activator\.bat$

# ProGuard
- (^|/)proguard\.pro$
- (^|/)proguard-rules\.pro$

# PuPHPet
- (^|/)puphpet/

# Android Google APIs
- (^|/)\.google_apis/

# Jenkins Pipeline
- (^|/)Jenkinsfile$

# Gitpod
- (^|/)\.gitpod\.Dockerfile$
`,

	"data/documentation.yml": `# Documentation files and directories are excluded from language
# statistics.
#
# Lines in this file are Regexps that are matched against the file
# pathname.
#
# Please add additional test coverage to
# `+"`"+`test/test_blob.rb#test_documentation`+"`"+` if you make any changes.

## Documentation directories ##

- ^[Dd]ocs?/
- (^|/)[Dd]ocumentation/
- (^|/)[Gg]roovydoc/
- (^|/)[Jj]avadoc/
- ^[Mm]an/
- ^[Ee]xamples/
- ^[Dd]emos?/
- (^|/)inst/doc/

## Documentation files ##

- (^|/)CITATION(\.cff|(S)?(\.(bib|md))?)$
- (^|/)CHANGE(S|LOG)?(\.|$)
- (^|/)CONTRIBUTING(\.|$)
- (^|/)COPYING(\.|$)
- (^|/)INSTALL(\.|$)
- (^|/)LICEN[CS]E(\.|$)
- (^|/)[Ll]icen[cs]e(\.|$)
- (^|/)README(\.|$)
- (^|/)[Rr]eadme(\.|$)

# Samples folders
- ^[Ss]amples?/
`,

}
