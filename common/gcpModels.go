package common

import (
	gcpUtils "cloud.google.com/go/storage"
)

type GCPObjectInfoExtension struct {
	ObjectInfo gcpUtils.ObjectAttrs
}

func (gie *GCPObjectInfoExtension) ContentType() string {
	return gie.ObjectInfo.ContentType
}

func (gie *GCPObjectInfoExtension) CacheControl() string {
	return gie.ObjectInfo.CacheControl
}

func (gie *GCPObjectInfoExtension) ContentDisposition() string {
	return gie.ObjectInfo.ContentDisposition
}

func (gie *GCPObjectInfoExtension) ContentEncoding() string {
	return gie.ObjectInfo.ContentEncoding
}

func (gie *GCPObjectInfoExtension) ContentLanguage() string {
	return gie.ObjectInfo.ContentLanguage
}

func (gie *GCPObjectInfoExtension) ContentMD5() []byte {
	b := gie.ObjectInfo.MD5
	return b
}

//NewCommonMetadata returns a map of user-defined key/value pairs
func (gie *GCPObjectInfoExtension) NewCommonMetadata() Metadata {
	md := Metadata{}
	for k, v := range gie.ObjectInfo.Metadata {
		md[k] = v
	}
	return md
}
