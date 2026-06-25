package builder

func registerChartPlugins(registry *PluginRegistry) {
	registry.Register(NewChartFilePlugin("barchart", generateBarChart))
	registry.Register(NewChartBlockPlugin("barchart", generateBarChartFromJSON))
	
	registry.Register(NewChartFilePlugin("piechart", generatePieChart))
	registry.Register(NewChartBlockPlugin("piechart", generatePieChartFromJSON))
	
	registry.Register(NewChartFilePlugin("linechart", generateLineChart))
	registry.Register(NewChartBlockPlugin("linechart", generateLineChartFromJSON))

	registry.Register(NewChartFilePlugin("donutchart", generateDonutChart))
	registry.Register(NewChartBlockPlugin("donutchart", generateDonutChartFromJSON))

	registry.Register(NewChartFilePlugin("hbarchart", generateHBarchart))
	registry.Register(NewChartBlockPlugin("hbarchart", generateHBarchartFromJSON))

	registry.Register(NewChartFilePlugin("multilinechart", generateMultiLineChart))
	registry.Register(NewChartBlockPlugin("multilinechart", generateMultiLineChartFromJSON))

	registry.Register(NewChartFilePlugin("groupedbarchart", generateGroupedBarChart))
	registry.Register(NewChartBlockPlugin("groupedbarchart", generateGroupedBarChartFromJSON))

	registry.Register(NewChartFilePlugin("radarchart", generateRadarChart))
	registry.Register(NewChartBlockPlugin("radarchart", generateRadarChartFromJSON))
}
