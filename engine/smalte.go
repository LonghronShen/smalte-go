package smalte

import "regexp"

type AllowedEnvToReplace struct {
    Pattern *regexp.Regexp
    Repl string
}

type TemplateConfig struct {
    Tmpl string
    Config string
}

