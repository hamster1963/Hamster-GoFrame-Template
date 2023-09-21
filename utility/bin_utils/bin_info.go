package binInfo

var (
	GitTag         = ""
	GitCommitLog   = ""
	GitStatus      = ""
	BuildTime      = ""
	BuildGoVersion = ""
)

// VersionString 返回版本信息
func VersionString() string {
	return "GitTag:" + GitTag + "\n" +
		"GitCommitLog:" + GitCommitLog + "\n" +
		"GitStatus:" + GitStatus + "\n" +
		"BuildTime:" + BuildTime + "\n" +
		"BuildGoVersion:" + BuildGoVersion + "\n"
}
