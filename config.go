package jmimg

type Config struct {
	AwsProfile    string `json:"aws_profile"`
	AwsRegion     string `json:"aws_region"`
	BucketName    string `json:"bucket_name"`
	CDNUrl        string `json:"cdn_url"`
	ConvertToJPEG bool   `json:"convert_to_jpeg"`
	MaxWidth      int    `json:"max_width"`
	MaxHeight     int    `json:"max_height"`
	StripMetadata *bool  `json:"strip_metadata"`
}
