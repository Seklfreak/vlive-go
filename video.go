package vlive_go

type Video struct {
	Title       string
	Seq         string
	URL         string
	Type        VideoType
	Product     VideoProduct
	Thumbnail   string
	ChannelName string
	ChannelType ChannelType
	ChannelId   string
	ChannelSeq  string
}

type VideoType string

const (
	VideoTypeLive VideoType = "LIVE"
	VideoTypeVOD  VideoType = "VOD"
)

type VideoProduct string

const (
	VideoProductNone VideoProduct = "NONE"
	VideoProductPaid VideoProduct = "PAID"
)

type ChannelType string

const (
	ChannelTypeBasic ChannelType = "BASIC"
)
