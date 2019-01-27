\NeedsTeXFormat{LaTeX2e}
\ProvidesClass{slipbox}[2019-01-16]

\LoadClass[border={0pt 1cm}]{standalone}

\RequirePackage{fontspec}
\RequirePackage{luaotfload}
\RequirePackage{xifthen}

%\DeclareOption{draft}{\overfullrule=2mm}

\setfontfamily\GaramondCaption{GaramondPremrPro-MedCapt.otf}[
  ItalicFont={GaramondPremrPro-MedItCapt.otf},
%  BoldFont={GaramondPremrPro-BdCapt.otf},
%  BoldItalicFont={GaramondPremrPro-BdItCapt.otf},
  RawFeature={+onum,+pnum,+liga,+clig,+dlig,+hlig,+calt}
]
\setfontfamily\GaramondText{GaramondPremrPro-Med.otf}[
  ItalicFont={GaramondPremrPro-MedIt.otf},
%  BoldFont={GaramondPremrPro-Bd.otf},
%  BoldItalicFont={GaramondPremrPro-BdIt.otf},
  RawFeature={+onum,+pnum,+liga,+clig,+dlig,+hlig,+calt}
]
\setfontfamily\GaramondSubhead{GaramondPremrPro-MedSubh.otf}[
  ItalicFont={GaramondPremrPro-MedItSubh.otf},
%  BoldFont={GaramondPremrPro-BdSubh.otf},
%  BoldItalicFont={GaramondPremrPro-BdItSubh.otf},
  RawFeature={+onum,+pnum,+liga,+clig,+dlig,+hlig,+calt}
]
\setfontfamily\GaramondDisplay{GaramondPremrPro-MedDisp.otf}[
  ItalicFont={GaramondPremrPro-MedItDisp.otf},
%  BoldFont={GaramondPremrPro-BdDisp.otf},
%  BoldItalicFont={GaramondPremrPro-BdItDisp.otf},
  RawFeature={+onum,+pnum,+liga,+clig,+dlig,+hlig,+calt}
]
\setmainfont{GaramondPremrPro-Med.otf}[
  ItalicFont={GaramondPremrPro-MedIt.otf},
%  BoldFont={GaramondPremrPro-Bd.otf},
%  BoldItalicFont={GaramondPremrPro-BdIt.otf},
  RawFeature={+onum,+pnum,+liga,+clig,+dlig,+hlig,+calt}
]
\setsansfont{GillSansMTPro-Medium.otf}[
  ItalicFont={GillSansMTPro-MediumItalic.otf},
%  BoldFont={GillSansMTPro-Bold.otf},
%  BoldItalicFont={GIllSansMTPro-BoldItalic.otf},
  RawFeature={+onum,+pnum}
]
\setmonofont{FiraCode-Medium.otf}[
%  BoldFont={FiraCode-Bold.otf},
  RawFeature={+liga,+calt,+clig,+dlig}
]
\def\tnum#1{{\addfontfeatures{+tnum,-pnum}#1}}
\def\lnum#1{{\addfontfeatures{+lnum,-onum}#1}}
\def\typesize#1#2{\ifnum#1<9\GaramondCaption\else\ifnum#1<15\GaramondText\else\ifnum#1<22\GaramondSubhead\else\GaramondDisplay\fi\fi}
\def\typesize#1#2{\fontsize{#1}{#2}\ifthenelse{#1<9}{\GaramondCaption}{\ifthenelse{#1<15}{\GaramondText}{\ifthenelse{#1<22}{\GaramondSubhead}{\GaramondDisplay}}}\selectfont}

\NewDocumentEnvironment{slip}{}{\noindent\begin{minipage}{48mm}\parskip=0pt\relax\parindent=\baselineskip\relax}{\end{minipage}}

\def\normalsize{\fontsize{10}{12}\selectfont}
\parindent=12pt
\parskip=0pt
\sfcode`:=3000
\sfcode`;=3000
\sfcode`.=3000
\sfcode`?=3000
\sfcode`—=3000
\sloppy

\endinput
