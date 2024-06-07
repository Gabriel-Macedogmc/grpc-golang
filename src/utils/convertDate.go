package utils

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func AsTime(x *timestamppb.Timestamp) time.Time {
	return time.Unix(int64(x.GetSeconds()), int64(x.GetNanos())).UTC()
}
