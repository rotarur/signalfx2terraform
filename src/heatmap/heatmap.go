package heatmap

import (
   "github.com/doctornkz/signalfx2terraform/src/utils"

   "github.com/hashicorp/hcl2/hcl"
   "github.com/hashicorp/hcl2/hclwrite"
   "github.com/signalfx/signalfx-go/chart"

   "github.com/zclconf/go-cty/cty"
)

// Chart - function for generating heatmap chart
func Chart(f *hclwrite.File, chart *chart.Chart) *hclwrite.Body {

   // program_text wrapper
   programText := utils.ProgramTextProc(chart.ProgramText)

   // tags wrapper
   tags := chart.Tags
   if tags == nil {
      tags = []string{}
   }

   // wrapper around label
   label := utils.LabelProc(chart.Id)

   rootBody := f.Body()
   chartBlock := rootBody.AppendNewBlock("resource", []string{utils.Type[chart.Options.Type], label})
   chartBody := chartBlock.Body()
   chartBody.SetAttributeValue("name", cty.StringVal(chart.Name))
   chartBody.SetAttributeValue("description", cty.StringVal(chart.Description))
   chartBody.SetAttributeTraversal("program_text", hcl.Traversal{
      hcl.TraverseRoot{
         Name: programText,
      },
   })
   chartBody.SetAttributeValue("unit_prefix", cty.StringVal(chart.Options.UnitPrefix))
   chartBody.SetAttributeValue("max_delay", cty.NumberIntVal(utils.MaxDelayProc(chart)))
   chartBody.SetAttributeValue("group_by", utils.GroupByProc(chart))

   if chart.Options.ColorBy == "Range" {
      utils.ColorRangeProc(chart, chartBody)
   } else {
      utils.ColorScale2Proc(chart, chartBody)
   }

   chartBody.SetAttributeValue("minimum_resolution", cty.NumberIntVal(utils.MinResolutionProc(chart)))
   chartBody.SetAttributeValue("disable_sampling", utils.DisableSamplingProc(chart))
   chartBody.SetAttributeValue("refresh_interval", cty.NumberIntVal(utils.RefreshIntervalProc(chart)))
   chartBody.AppendNewline()
   return chartBody
}
