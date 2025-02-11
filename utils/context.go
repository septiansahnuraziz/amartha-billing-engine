package utils

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

func DumpOutGoingContext(c context.Context) string {
	md, _ := metadata.FromOutgoingContext(c)
	return Dump(md)
}

func DumpIncomingContext(c context.Context) string {
	md, _ := metadata.FromIncomingContext(c)
	return Dump(md)
}

func GetTraceID(ctx context.Context) string {
	traceID, ok := ctx.Value("traceID").(string)

	if !ok {
		return ""
	}

	return traceID
}

func GetTraceIDFromCtx(ctx context.Context) string {
	traceID, ok := ctx.Value("traceID").(string)
	if !ok {
		return ""
	}
	return traceID
}

func GetSourceFromCtx(ctx context.Context) string {
	source, ok := ctx.Value("source").(string)
	if !ok {
		return ""
	}

	return source
}

func SetTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, "traceID", traceID)
}

func SetSource(ctx context.Context, source string) context.Context {
	return context.WithValue(ctx, "source", source)
}

func SetDeviceID(ctx context.Context, deviceID string) context.Context {
	return context.WithValue(ctx, "deviceID", deviceID)
}

func GetSourceHeaderValue(context *gin.Context) string {
	source := context.Request.Header.Get("source")

	if source == "" {
		return ""
	}

	return ExpectedString(source)
}

func GetPlatformHeaderValue(ctx *gin.Context) string {
	platform := ctx.Request.Header.Get("platform")

	if platform == "" {
		return ""
	}

	return platform
}

func GetDevicePlatformHeaderValue(ctx *gin.Context) string {
	devicePlatform := ctx.Request.Header.Get("Device-Platform")

	if devicePlatform == "" {
		return ""
	}

	return devicePlatform
}

func GetDeviceIDHeaderValue(ctx *gin.Context) string {
	deviceID := ctx.Request.Header.Get("device-id")

	if deviceID == "" {
		return ""
	}

	return deviceID
}

func GetDeviceIDFromContext(ctx context.Context) string {
	deviceID, isExist := ctx.Value("deviceID").(string)
	if !isExist {
		return ""
	}

	return ExpectedString(deviceID)
}

// GetEpochHeaderValue is method to get epoch header value
func GetEpochHeaderValue(c *gin.Context) string {
	return c.Request.Header.Get("epoch")
}

// GetSignatureHeaderValue is function to get signature header value
func GetSignatureHeaderValue(c *gin.Context) string {
	return c.Request.Header.Get("signature")
}
