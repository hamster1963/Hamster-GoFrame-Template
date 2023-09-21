package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"kes-speed-backend/internal/cmd"
	_ "kes-speed-backend/internal/logic"
	_ "kes-speed-backend/internal/packed"
	binInfo "kes-speed-backend/utility/bin_utils"
)

// 初始化为 unknown，如果编译时没有传入这些值，则为 unknown
var (
	GitTag         = "unknown"
	GitCommitLog   = "unknown"
	GitStatus      = "cleanly"
	BuildTime      = "unknown"
	BuildGoVersion = "unknown"
)

func main() {
	// 注入编译时的信息
	binInfo.GitTag = GitTag
	binInfo.GitCommitLog = GitCommitLog
	binInfo.GitStatus = GitStatus
	binInfo.BuildTime = BuildTime
	binInfo.BuildGoVersion = BuildGoVersion
	cmd.Main.Run(gctx.New())
}
