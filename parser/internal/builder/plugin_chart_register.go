package builder

func registerChartPlugins(registry *PluginRegistry) {
	registry.Register(NewChartFilePlugin("barchart", generateBarChart))
	registry.Register(NewChartBlockPlugin("barchart", generateBarChartFromJSON))
	
	registry.Register(NewChartFilePlugin("piechart", generatePieChart))
	registry.Register(NewChartBlockPlugin("piechart", generatePieChartFromJSON))
	
	registry.Register(NewChartFilePlugin("linechart", generateLineChart))
	registry.Register(NewChartBlockPlugin("linechart", generateLineChartFromJSON))

	// The remaining advanced charts do not take colors, but our generic generator func signature expects it.
	// We'll just ignore the colors argument for them.
	registry.Register(NewChartFilePlugin("donutchart", func(sourceDir, file, title, colors string) string { return generateDonutChart(sourceDir, file, title) }))
	registry.Register(NewChartBlockPlugin("donutchart", func(jsonData []byte, title, colors string) string { return generateDonutChartFromJSON(jsonData, title) }))

	registry.Register(NewChartFilePlugin("hbarchart", func(sourceDir, file, title, colors string) string { return generateHBarchart(sourceDir, file, title) }))
	registry.Register(NewChartBlockPlugin("hbarchart", func(jsonData []byte, title, colors string) string { return generateHBarchartFromJSON(jsonData, title) }))

	registry.Register(NewChartFilePlugin("multilinechart", func(sourceDir, file, title, colors string) string { return generateMultiLineChart(sourceDir, file, title) }))
	registry.Register(NewChartBlockPlugin("multilinechart", func(jsonData []byte, title, colors string) string { return generateMultiLineChartFromJSON(jsonData, title) }))

	registry.Register(NewChartFilePlugin("groupedbarchart", func(sourceDir, file, title, colors string) string { return generateGroupedBarChart(sourceDir, file, title) }))
	registry.Register(NewChartBlockPlugin("groupedbarchart", func(jsonData []byte, title, colors string) string { return generateGroupedBarChartFromJSON(jsonData, title) }))

	registry.Register(NewChartFilePlugin("radarchart", func(sourceDir, file, title, colors string) string { return generateRadarChart(sourceDir, file, title) }))
	registry.Register(NewChartBlockPlugin("radarchart", func(jsonData []byte, title, colors string) string { return generateRadarChartFromJSON(jsonData, title) }))
}
