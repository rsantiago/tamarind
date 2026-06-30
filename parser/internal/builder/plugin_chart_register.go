package builder

func init() {
	RegisterDefaultPlugin(func() ShortcodePlugin { return NewChartFilePlugin("barchart", generateBarChart) })
	RegisterDefaultPlugin(func() ShortcodePlugin { return NewChartBlockPlugin("barchart", generateBarChartFromJSON) })

	RegisterDefaultPlugin(func() ShortcodePlugin { return NewChartFilePlugin("piechart", generatePieChart) })
	RegisterDefaultPlugin(func() ShortcodePlugin { return NewChartBlockPlugin("piechart", generatePieChartFromJSON) })

	RegisterDefaultPlugin(func() ShortcodePlugin { return NewChartFilePlugin("linechart", generateLineChart) })
	RegisterDefaultPlugin(func() ShortcodePlugin { return NewChartBlockPlugin("linechart", generateLineChartFromJSON) })

	RegisterDefaultPlugin(func() ShortcodePlugin { return NewChartFilePlugin("donutchart", generateDonutChart) })
	RegisterDefaultPlugin(func() ShortcodePlugin { return NewChartBlockPlugin("donutchart", generateDonutChartFromJSON) })

	RegisterDefaultPlugin(func() ShortcodePlugin { return NewChartFilePlugin("hbarchart", generateHBarchart) })
	RegisterDefaultPlugin(func() ShortcodePlugin { return NewChartBlockPlugin("hbarchart", generateHBarchartFromJSON) })

	RegisterDefaultPlugin(func() ShortcodePlugin { return NewChartFilePlugin("multilinechart", generateMultiLineChart) })
	RegisterDefaultPlugin(func() ShortcodePlugin { return NewChartBlockPlugin("multilinechart", generateMultiLineChartFromJSON) })

	RegisterDefaultPlugin(func() ShortcodePlugin { return NewChartFilePlugin("groupedbarchart", generateGroupedBarChart) })
	RegisterDefaultPlugin(func() ShortcodePlugin { return NewChartBlockPlugin("groupedbarchart", generateGroupedBarChartFromJSON) })

	RegisterDefaultPlugin(func() ShortcodePlugin { return NewChartFilePlugin("radarchart", generateRadarChart) })
	RegisterDefaultPlugin(func() ShortcodePlugin { return NewChartBlockPlugin("radarchart", generateRadarChartFromJSON) })
}
