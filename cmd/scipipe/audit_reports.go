package main

import (
	"strings"
	"text/template"
	"time"

	"github.com/scipipe/scipipe"
)

// AuditReport is a container for data to be parsed into an audit report, in
// HTML, TeX or other format
type auditReport struct {
	FileName    string
	SciPipeVer  string
	RunTime     time.Duration
	AuditInfos  []*scipipe.AuditInfo
	ColorDef    string
	ChartHeight string
}

var (
	tplFuncs = template.FuncMap{
		"strrepl":        func(subj string, find string, repl string) string { return strings.Replace(subj, find, repl, -1) },
		"sub":            func(val1 int, val2 int) int { return val1 - val2 },
		"timesub":        func(t1 time.Time, t2 time.Time) time.Duration { return t1.Sub(t2) },
		"durtomillis":    func(exact time.Duration) (rounded time.Duration) { return exact.Truncate(1e6 * time.Nanosecond) },
		"timetomillis":   func(exact time.Time) (rounded time.Time) { return exact.Truncate(1e6 * time.Nanosecond) },
		"durtomillisint": func(exact time.Duration) (millis int) { return int(exact.Nanoseconds() / 1000000) },
		"inc":            func(i int) int { return i + 1 },
		"strmaptoslice": func(m map[string]string) []string {
			sl := []string{}
			for _, s := range m {
				sl = append(sl, s)
			}
			return sl
		},
	}
)

func auditInfoToHTML(inFilePath string, outFilePath string, flatten bool) error {
	_ = "STUB: not implemented"
	return nil
}

func formatTaskHTML(fileName string, auditInfo *scipipe.AuditInfo) (outHTML string) {
	_ = "STUB: not implemented"
	return ""
}

//upStreamHTML := ""
//for filePath, uai := range auditInfo.Upstream {
//	upStreamHTML += formatTaskHTML(filePath, uai)
//}
//if outHTML != "" {
//	outHTML += "<tr><th>Upstreams:</th><td>" + upStreamHTML + "</td></tr>\n"
//}

func auditInfoToBash(inFilePath string, outFilePath string, flatten bool) error {
	_ = "STUB: not implemented"
	return nil
}

func auditInfoToTeX(inFilePath string, outFilePath string, flatten bool) error {
	_ = "STUB: not implemented"
	return nil
}

func extractAuditInfosByID(auditInfo *scipipe.AuditInfo) (auditInfosByID map[string]*scipipe.AuditInfo) {
	_ = "STUB: not implemented"
	return nil
}

func mergeStringAuditInfoMaps(ms ...map[string]*scipipe.AuditInfo) (merged map[string]*scipipe.AuditInfo) {
	_ = "STUB: not implemented"
	return nil
}

func sortAuditInfosByStartTime(auditInfosByID map[string]*scipipe.AuditInfo) []*scipipe.AuditInfo {
	_ = "STUB: not implemented"
	return nil
}

const headHTMLPattern = `<html>
<head>
<style>
	body { font-family: arial, helvetica, sans-serif; }
	table { color: #546E7A; background: #EFF2F5; border: none; width: 960px; margin: 1em 1em 2em 1em; padding: 1.2em; font-size: 10pt; opacity: 1; }
	table:hover { color: black; background: #FFFFEF; }
	th { text-align: right; vertical-align: top; padding: .2em .8em; width: 9em; }
	td { vertical-align: top; }
	.task-title { font-size: 12pt; font-weight: normal; }
	.cmdbox { border: rgb(156, 184, 197) 0px solid; background: #D2DBE0; font-family: 'Ubuntu mono', Monospace, 'Courier New'; padding: .8em 1em; margin: 0.4em 0; font-size: 12pt; }
	table:hover .cmdbox { background: #EFEFCC; }
	.greyout { color: #999; }
	a, a:link, a:visited { color: inherit; text-decoration: none; }
	a:hover { text-decoration: underline; }
</style>
<title>Audit info for: %s</title>
</head>
<body>
`
const bottomHTML = `</body>
</html>`

// LaTeX code from vision.tex:
const texTemplate = `\documentclass[11pt,oneside,openright]{memoir}

\usepackage{tcolorbox}
\usepackage[scaled]{beramono}
\renewcommand*\familydefault{\ttdefault}
\usepackage[T1]{fontenc}
\usepackage{tabularx}
\usepackage{listings}
\usepackage{graphicx}
\usepackage{tikz}
\usepackage{pgfplots}
\usepackage{pgfplotstable}
\usepackage{xcolor}

{{ .ColorDef }}

% from https://tex.stackexchange.com/a/128040/110842
% filter to only get the current row in \pgfplotsinvokeforeach
\pgfplotsset{
    select row/.style={
        x filter/.code={\ifnum\coordindex=#1\else\def\pgfmathresult{}\fi}
    }
}

\pgfplotstableread[col sep=comma]{
start,end,Name,color
{{ $startTime := (index .AuditInfos 0).StartTime }}
{{ range $i, $v := .AuditInfos }}{{ durtomillisint (timesub $v.StartTime $startTime) }},{{ durtomillisint (timesub $v.FinishTime $startTime) }},{{ strrepl .ProcessName "_" "\\_" }},color{{ $i }}
{{ end }}
}\loadedtable
\pgfplotstablegetrowsof{\loadedtable}
\pgfplotsset{compat=1.13}
\pgfmathsetmacro{\tablerows}{int(\pgfplotsretval-1)}

\begin{document}
\pagestyle{plain}
\noindent
\begin{minipage}{\textwidth}
    \vspace{-8em}\hspace{-8em}
    %\includegraphics[width=9em]{images/scipipe_logo_bluegrey.png}
\end{minipage}

\noindent
{\huge\textbf{SciPipe Audit Report}} \\
{\large\textbf{For file: {{ (strrepl (strrepl .FileName ".audit.json" "") "_" "\\_") }}} \\
\vspace{10pt}

    \begin{tcolorbox}[ title=Summary information ]
    \small
\begin{tabular}{rp{0.72\linewidth}}
SciPipe version: & {{ .SciPipeVer }} \\
Start time:  & {{ timetomillis (index .AuditInfos 0).StartTime }} \\
Finish time: & {{ timetomillis (index .AuditInfos (sub (len .AuditInfos) 1)).FinishTime }} \\
Run time: & {{ durtomillis .RunTime }}  \\
\end{tabular}
    \end{tcolorbox}

\setlength{\fboxsep}{0pt}
\noindent

%\hspace{-0.1725\textwidth}\fbox{\includegraphics[width=1.35\textwidth]{images/cawpre.pdf}}

\section*{Execution timeline}

\begin{tikzpicture}
\begin{axis}[
    xbar, xmin=0,
    y axis line style = { opacity = 0 },
    tickwidth         = 0pt,
	width=10cm,
	height={{ .ChartHeight }}cm,
    % next two lines also from https://tex.stackexchange.com/a/128040/110842,
    ytick={0,...,\tablerows},
    yticklabels from table={\loadedtable}{Name},
    xbar stacked,
    bar shift=0pt,
    y dir=reverse,
    xtick={1, 10, 1000, 60000, 120000, 180000, 240000, 300000, 600000, 900000, 1200000},
    xticklabels={0, 10 ms, 1 s, 1 min, 2 min, 3 min, 4 min, 5 min, 10 min, 15 min, 20 min},
    scaled x ticks=false,
]

\pgfplotsinvokeforeach{0,...,\tablerows}{
    % get color from table, commands defined must be individual for each plot
    % because the color is used in \end{axis} and therefore would otherwise
    % use the last definition
    \pgfplotstablegetelem{#1}{color}\of{\loadedtable}
    \expandafter\edef\csname barcolor.#1\endcsname{\pgfplotsretval}
    \addplot+[color=\csname barcolor.#1\endcsname] table [select row=#1, x expr=\thisrow{end}-\thisrow{start}, y expr=#1]{\loadedtable};
}
\end{axis}
\end{tikzpicture}

\section*{Tasks}
    \lstset{ breaklines=true,
            postbreak=\mbox{\textcolor{red}{$\hookrightarrow$}\space},
            aboveskip=8pt,belowskip=8pt}

{{ range $i, $v := .AuditInfos }}
   \begin{tcolorbox}[ title={{ (strrepl $v.ProcessName "_" "\\_") }},
                      colbacktitle=color{{ $i }}!63!white,
                      colback=color{{ $i }}!37!white,
                      coltitle=black ]
       \small
       \begin{tabular}{rp{0.72\linewidth}}
ID: & {{ $v.ID }} \\
Process: & {{ (strrepl $v.ProcessName "_" "\\_") }} \\
Command: & \begin{lstlisting}
{{ strrepl $v.Command "_" "\\_" }}
\end{lstlisting} \\
Parameters:& {{ range $k, $v := $v.Params }}{{- $k -}}={{- $v -}}{{ end }} \\
Tags: & {{ range $k, $v := $v.Tags }}{{- $k -}}={{- $v -}}{{ end }} \\
Start time:  & {{ timetomillis $v.StartTime }} \\
Finish time: & {{ timetomillis $v.FinishTime }} \\
Execution time: & {{ durtomillis $v.ExecTimeNS }} \\
        \end{tabular}
	\end{tcolorbox}
{{ end }}

\end{document}`

const bashTemplate = `#!/bin/bash
# ------------------------------------------------------------------------
# Bash script for reproducing file: {{ (strrepl .FileName ".audit.json" "") }}
# Generated by: SciPipe {{ .SciPipeVer }}
# SciPipe website: http://scipipe.org
# ------------------------------------------------------------------------
{{ range $i, $v := .AuditInfos }}
proc=$(printf '%-32s' "{{ $v.ProcessName }}")
if [[{{ range $i, $v := (strmaptoslice .OutFiles) }}{{if (gt $i 0)}}||{{end}} -f {{ $v }} {{ end }}]]; then
  echo "$(date '+%Y%m%d %H:%M:%S') | $proc | Some of these output files exist, so skipping:{{ range $i, $v := (strmaptoslice .OutFiles) }}{{if (gt $i 0)}},{{end}} {{ $v }}{{end}}";
else
  echo "$(date '+%Y%m%d %H:%M:%S') | $proc | Executing: {{ strrepl (strrepl $v.Command "../" "") "\"" "\\\"" }}"
  {{ strrepl $v.Command "../" "" }}
  echo "$(date '+%Y%m%d %H:%M:%S') | $proc |  Finished: {{ strrepl (strrepl $v.Command "../" "") "\"" "\\\"" }}"
fi;
{{ end }}
`
