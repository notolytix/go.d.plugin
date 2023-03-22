package hbase

import "github.com/netdata/go.d.plugin/agent/module"

type (
	Charts = module.Charts
	Dims   = module.Dims
	Vars   = module.Vars
)

var jvmCharts = Charts{
	{
		ID:    "jvm_heap_memory",
		Title: "Heap Memory",
		Units: "MiB",
		Fam:   "jvm",
		Ctx:   "hdfs.heap_memory",
		Type:  module.Area,
		Dims: Dims{
			{ID: "jvm_mem_heap_committed", Name: "committed", Div: 1000},
			{ID: "jvm_mem_heap_used", Name: "used", Div: 1000},
		},
		Vars: Vars{
			{ID: "jvm_mem_heap_max"},
		},
	},
	{
		ID:    "jvm_gc_count_total",
		Title: "GC Events",
		Units: "events/s",
		Fam:   "jvm",
		Ctx:   "hdfs.gc_count_total",
		Dims: Dims{
			{ID: "jvm_gc_count", Name: "gc", Algo: module.Incremental},
		},
	},
	{
		ID:    "jvm_gc_time_total",
		Title: "GC Time",
		Units: "ms",
		Fam:   "jvm",
		Ctx:   "hdfs.gc_time_total",
		Dims: Dims{
			{ID: "jvm_gc_time_millis", Name: "time", Algo: module.Incremental},
		},
	},
	{
		ID:    "jvm_gc_threshold",
		Title: "Number of Times That the GC Threshold is Exceeded",
		Units: "events/s",
		Fam:   "jvm",
		Ctx:   "hdfs.gc_threshold",
		Dims: Dims{
			{ID: "jvm_gc_num_info_threshold_exceeded", Name: "info", Algo: module.Incremental},
			{ID: "jvm_gc_num_warn_threshold_exceeded", Name: "warn", Algo: module.Incremental},
		},
	},
	{
		ID:    "jvm_threads",
		Title: "Number of Threads",
		Units: "num",
		Fam:   "jvm",
		Ctx:   "hdfs.threads",
		Type:  module.Stacked,
		Dims: Dims{
			{ID: "jvm_threads_new", Name: "new"},
			{ID: "jvm_threads_runnable", Name: "runnable"},
			{ID: "jvm_threads_blocked", Name: "blocked"},
			{ID: "jvm_threads_waiting", Name: "waiting"},
			{ID: "jvm_threads_timed_waiting", Name: "timed_waiting"},
			{ID: "jvm_threads_terminated", Name: "terminated"},
		},
	},
	{
		ID:    "jvm_logs_total",
		Title: "Number of Logs",
		Units: "logs/s",
		Fam:   "jvm",
		Ctx:   "hdfs.logs_total",
		Type:  module.Stacked,
		Dims: Dims{
			{ID: "jvm_log_info", Name: "info", Algo: module.Incremental},
			{ID: "jvm_log_error", Name: "error", Algo: module.Incremental},
			{ID: "jvm_log_warn", Name: "warn", Algo: module.Incremental},
			{ID: "jvm_log_fatal", Name: "fatal", Algo: module.Incremental},
		},
	},
}

var masterCharts = Charts{
	{
		ID:    "master_regions",
		Title: "Region Servers",
		Units: "Count",
		Fam:   "master",
		Ctx:   "hbase.regions",
		Type:  module.Line,
		Dims: Dims{
			{ID: "master_number_of_regions", Name: "regions"},
			{ID: "master_number_of_dead_regions", Name: "dead_regions"},
		},
	},
	{
		ID:    "cluster_requests",
		Title: "Cluster Requests",
		Units: "Count",
		Fam:   "master",
		Ctx:   "hbase.requests",
		Type:  module.Line,
		Dims: Dims{
			{ID: "master_cluster_requests", Name: "Requests", Algo: module.Incremental},
		},
	},
}

var regionCharts = Charts{
	{
		ID:    "master_regions",
		Title: "Region Servers",
		Units: "Count",
		Fam:   "master",
		Ctx:   "hbase.regions",
		Type:  module.Line,
		Dims: Dims{
			{ID: "master_number_of_regions", Name: "regions"},
			{ID: "master_number_of_dead_regions", Name: "dead_regions"},
		},
	},
	{
		ID:    "cluster_requests",
		Title: "Cluster Requests",
		Units: "Count",
		Fam:   "master",
		Ctx:   "hbase.requests",
		Type:  module.Line,
		Dims: Dims{
			{ID: "master_cluster_requests", Name: "Requests", Algo: module.Incremental},
		},
	},
}

func regionserverCharts() *Charts {
	charts := Charts{}
	panicIfError(charts.Add(*jvmCharts.Copy()...))
	panicIfError(charts.Add(*regionCharts.Copy()...))
	return &charts
}

func masterServerCharts() *Charts {
	charts := Charts{}
	panicIfError(charts.Add(*jvmCharts.Copy()...))
	panicIfError(charts.Add(*masterCharts.Copy()...))
	return &charts
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
