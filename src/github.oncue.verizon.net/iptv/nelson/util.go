package main

/// had to have a `util` file.
/// Because pragmatism.
/// Because irony.

import (
  "os"
  "time"
  "github.com/parnurzeal/gorequest"
  "github.com/olekukonko/tablewriter"
  "github.com/briandowns/spinner"
)

func AugmentRequest(c *gorequest.SuperAgent, cfg *Config) *gorequest.SuperAgent {
  return c.
    AddCookie(cfg.GetAuthCookie()).
    Set("Content-type","application/json").
    Timeout(15*time.Second).
    SetCurlCommand(false).
    SetDebug(false)
}

func RenderTableToStdout(headers []string, data [][]string){
  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader(headers)
  table.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})
  table.SetHeaderLine(false)
  table.SetRowLine(false)
  table.SetColumnSeparator("")
  table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
  table.SetAlignment(tablewriter.ALIGN_LEFT)
  table.AppendBulk(data) // Add Bulk Data
  table.Render()
}

func JavaEpochToDateStr(long int64) string {
  t := time.Unix(0, long*int64(time.Millisecond))
  return t.Format(time.RFC3339)
}

func ProgressIndicator() *spinner.Spinner {
  s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
  s.Color("green")
  return s
}