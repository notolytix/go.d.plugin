package hbase

type metrics struct {
	Jvm              *jvmMetrics              `stm:"jvm"`
	Master           *masterMetrics            `stm:"master"`
}

type jvmMetrics struct {
	ProcessName string `json:"tag.ProcessName"`
	HostName    string `json:"tag.Hostname"`
	//MemNonHeapUsedM            float64 `stm:"mem_non_heap_used,1000,1"`
	//MemNonHeapCommittedM       float64 `stm:"mem_non_heap_committed,1000,1"`
	//MemNonHeapMaxM             float64 `stm:"mem_non_heap_max"`
	MemHeapUsedM      float64 `stm:"mem_heap_used,1000,1"`
	MemHeapCommittedM float64 `stm:"mem_heap_committed,1000,1"`
	MemHeapMaxM       float64 `stm:"mem_heap_max"`
	//MemMaxM                    float64 `stm:"mem_max"`
	GcCount                    float64 `stm:"gc_count"`
	GcTimeMillis               float64 `stm:"gc_time_millis"`
	GcNumWarnThresholdExceeded float64 `stm:"gc_num_warn_threshold_exceeded"`
	GcNumInfoThresholdExceeded float64 `stm:"gc_num_info_threshold_exceeded"`
	GcTotalExtraSleepTime      float64 `stm:"gc_total_extra_sleep_time"`
	ThreadsNew                 float64 `stm:"threads_new"`
	ThreadsRunnable            float64 `stm:"threads_runnable"`
	ThreadsBlocked             float64 `stm:"threads_blocked"`
	ThreadsWaiting             float64 `stm:"threads_waiting"`
	ThreadsTimedWaiting        float64 `stm:"threads_timed_waiting"`
	ThreadsTerminated          float64 `stm:"threads_terminated"`
	LogFatal                   float64 `stm:"log_fatal"`
	LogError                   float64 `stm:"log_error"`
	LogWarn                    float64 `stm:"log_warn"`
	LogInfo                    float64 `stm:"log_info"`
}


type masterMetrics struct {
	IsActiveMaster string `json:"tag.isActiveMaster"`
	NumberOfRegions float64 `json:"numRegionServers" stm:"number_of_regions"`
	NumberOfDeadRegions float64 `json:"numDeadRegionServers" stm:"number_of_dead_regions"`
	ClusterRequests float64 `json:"clusterRequests" stm:"cluster_requests"`
}


